package main

import (
	"fmt"
	"goldrush/client"
	"goldrush/models"
	"goldrush/utils"
	"os"
	"runtime"
	"sync"
	"time"
)

type Coordinates struct {
	posX int
	posY int
}

type AreaStats struct {
	mu    sync.Mutex
	avg   float32
	n     int
	total int
}

func (stats *AreaStats) observe(amount int) {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.n++
	stats.total += amount
	stats.avg = float32(stats.total) / float32(stats.n)
}

const maxLicenses = 10

type useLicense func(callback func(int))

const queueSize = 100

func main() {
	address, isRemote := os.LookupEnv("ADDRESS")
	if !isRemote {
		address = "localhost"
	}
	mineClient := client.NewMineClient(address)

	areaChan := make(chan models.Area)
	exploreChan := make(chan Coordinates, queueSize)
	digChan := make(chan models.ExploreResp, queueSize)
	goldChan := make(chan string, queueSize)
	cashChan := make(chan int, 5000)

	getLicenseLease := issueLicense(mineClient, cashChan, isRemote)

	cpus := runtime.NumCPU()
	fmt.Println("Cpus: ", cpus)
	areaExplorers := cpus
	explorers := cpus * 2
	diggers := cpus * 2
	cashiers := 1

	stats := AreaStats{}
	for w := 1; w <= areaExplorers; w++ {
		go exploreArea(mineClient, areaChan, exploreChan, &stats)
	}
	for w := 1; w <= explorers; w++ {
		go explore(mineClient, exploreChan, digChan)
	}
	for w := 1; w <= diggers; w++ {
		go dig(mineClient, digChan, getLicenseLease, goldChan)
	}
	for w := 1; w <= cashiers; w++ {
		go cash(mineClient, goldChan, cashChan)
	}

	go reportMetrics(mineClient, isRemote)
	go reportUsage()

	processed := 0
	go func() {
		for {
			fmt.Printf("Processed: %d. Area explore: %d. Explore: %d. Dig: %d. Gold: %d. Cash %d. Stats: %v\n", processed, len(areaChan), len(exploreChan), len(digChan), len(goldChan), len(cashChan), &stats)
			if isRemote {
				time.Sleep(5 * time.Minute)
			} else {
				time.Sleep(10 * time.Second)
			}
		}
	}()

	step := 5
	for i := 0; i < 3500; i += step {
		for j := 0; j < 3500; j += step {
			areaChan <- models.Area{PosX: i, PosY: j, SizeX: step, SizeY: step}
			processed = 3500*i + j
		}
	}
}

func issueLicense(mineClient *client.MineClient, cashChan chan int, isRemote bool) func(callback func(int)) {
	licenseIdChannel := make(chan int, 50)
	licenseIdAckChannel := make(chan int, 50)
	licenses := make(map[int]*models.License)

	go func() {
		for {
			select {
			case licenseId := <-licenseIdAckChannel:
				licenses[licenseId].DigUsed++
				if licenses[licenseId].DigAllowed == licenses[licenseId].DigUsed {
					delete(licenses, licenseId)
				}
			default:
				if len(licenses) < maxLicenses {
					var cashList []int
					select {
					case cash := <-cashChan:
						cashList = []int{cash}
					default:
						cashList = []int{}
					}

					for {
						license, licenseErr := mineClient.IssueLicense(cashList)
						if licenseErr == nil {
							if licenses[license.Id] == nil {
								licenses[license.Id] = &license
							} else {
								licenses[license.Id].DigAllowed += license.DigAllowed
							}
							for i := 0; i < license.DigAllowed; i++ {
								licenseIdChannel <- license.Id
							}
							break
						}
					}
				} else {
					time.Sleep(10 * time.Millisecond)
				}
			}
		}
	}()

	go func() {
		for {
			fmt.Printf("Licenses: %d\n", len(licenseIdChannel))
			if isRemote {
				time.Sleep(5 * time.Minute)
			} else {
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return func(callback func(int)) {
		licenseId := <-licenseIdChannel
		callback(licenseId)
		licenseIdAckChannel <- licenseId
	}
}

func cash(mineClient *client.MineClient, goldChan chan string, cashChan chan int) {
	for gold := range goldChan {
		for {
			cash, cashErr := mineClient.Cash(gold)
			if cashErr == nil {
				for i := 0; i < len(cash); i++ {
					select {
					case cashChan <- cash[i]:
					default:
					}
				}
				break
			}
		}
	}
}

func dig(mineClient *client.MineClient, digChan chan models.ExploreResp, useLicense useLicense, goldChan chan string) {
	for exploreRes := range digChan {
		left := exploreRes.Amount
		for k := 1; k <= 10 && left > 0; {
			useLicense(func(licenseId int) {
				for {
					digResult, digErr := mineClient.Dig(exploreRes.Area.PosX, exploreRes.Area.PosY, k, licenseId)
					if digErr == nil {
						for i := 0; i < len(digResult); i++ {
							goldChan <- digResult[i]
						}
						k++
						left -= len(digResult)
						break
					}
				}
			})
		}
	}
}

func exploreArea(mineClient *client.MineClient, areaChan chan models.Area, exploreChan chan Coordinates, stats *AreaStats) {
	for area := range areaChan {
		for {
			exploreRes, exploreErr := mineClient.Explore(&area)
			if exploreErr == nil {
				if exploreRes.Amount > 0 {
					if stats.avg <= float32(exploreRes.Amount) || len(exploreChan) < queueSize {
						stats.observe(exploreRes.Amount)
						for i := exploreRes.Area.PosX; i < exploreRes.Area.PosX+exploreRes.Area.SizeX; i++ {
							for j := exploreRes.Area.PosY; j < exploreRes.Area.PosY+exploreRes.Area.SizeY; j++ {
								exploreChan <- Coordinates{posX: i, posY: j}
							}
						}
					}
				}
				break
			}
		}
	}
}

func explore(mineClient *client.MineClient, exploreChan chan Coordinates, digChan chan models.ExploreResp) {
	for coords := range exploreChan {
		for {
			exploreRes, exploreErr := mineClient.Explore(&models.Area{PosX: coords.posX, PosY: coords.posY, SizeX: 1, SizeY: 1})
			if exploreErr == nil {
				if exploreRes.Amount > 0 {
					digChan <- exploreRes
				}
				break
			}
		}
	}
}

func reportUsage() {
	for {
		fmt.Println("----------")
		utils.PrintMemoryUsage()
		utils.PrintCpuUsage()
		fmt.Println("----------")
		time.Sleep(60 * time.Second)
	}
}

func reportMetrics(mineClient *client.MineClient, isRemote bool) {
	if isRemote {
		time.Sleep(15*time.Minute - 20*time.Second)
	} else {
		time.Sleep(2*time.Minute - 5*time.Second)
	}
	mineClient.PrintMetrics()
}
