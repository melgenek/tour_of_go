package main

import (
	"fmt"
	"goldrush/client"
	"goldrush/models"
	"goldrush/utils"
	"os"
	"runtime"
	"time"
)

type Coordinates struct {
	posX int
	posY int
}

const maxLicenses = 10

type useLicense func(callback func(int))

func main() {
	address, isRemote := os.LookupEnv("ADDRESS")
	if !isRemote {
		address = "localhost"
	}
	mineClient := client.NewMineClient(address)

	exploreChan := make(chan Coordinates, 1000)
	digChan := make(chan models.ExploreResp, 1000)
	goldChan := make(chan string, 1000)
	cashChan := make(chan int, 5000)

	getLicenseLease := issueLicense(mineClient, cashChan, isRemote)

	cpus := runtime.NumCPU()
	explorers := cpus * 2
	diggers := cpus * 2
	cashiers := 1

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

	go func() {
		for {
			fmt.Printf("Explore: %d. Dig: %d. Gold: %d. Cash %d\n", len(exploreChan), len(digChan), len(goldChan), len(cashChan))
			if isRemote {
				time.Sleep(5 * time.Minute)
			} else {
				time.Sleep(10 * time.Second)
			}
		}
	}()

	for i := 0; i < 3500; i++ {
		for j := 0; j < 3500; j++ {
			exploreChan <- Coordinates{posX: i, posY: j}
			processed := 3500*i + j
			if processed%20000 == 0 {
				fmt.Printf("Processed %d\n", processed)
			}
		}
	}
}

//func allInOne(mineClient *client.MineClient, exploreChan chan Coordinates,
//	digChan chan models.ExploreResp, goldChan chan string, cashChan chan int) {
//
//}

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
			fmt.Printf("Liceses: %d\n", len(licenseIdChannel))
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

func explore(mineClient *client.MineClient, exploreChan chan Coordinates, digChan chan models.ExploreResp) {
	for coords := range exploreChan {
		for {
			exploreRes, exploreErr := mineClient.Explore(coords.posX, coords.posY)
			if exploreErr == nil {
				if exploreRes.Amount > 0 {
					digChan <- exploreRes
				}
				break
			}
		}
	}
}

func reportMetrics(mineClient *client.MineClient, isRemote bool) {
	for {
		//if isRemote {
		//	time.Sleep(5*time.Minute - 5*time.Second)
		//} else {
		//	time.Sleep(1*time.Minute - 5*time.Second)
		//}

		fmt.Println("----------")
		utils.PrintMemoryUsage()
		utils.PrintCpuUsage()
		//mineClient.PrintMetrics()
		fmt.Println("----------")
		time.Sleep(30 * time.Second)
	}
}
