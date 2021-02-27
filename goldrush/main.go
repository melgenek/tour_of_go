package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"goldrush/client"
	"goldrush/models"
	"goldrush/utils"
	"os"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var totalCash uint64 = 0

const maxLicenses = 10

func main() {
	cpus := runtime.NumCPU()
	start := time.Now()
	fmt.Println("Cpus: ", cpus)

	address, isRemote := os.LookupEnv("ADDRESS")
	if !isRemote {
		address = "localhost"
	}
	mineClient := client.NewMineClient(address)

	go reportMetrics(mineClient, isRemote)
	go reportUsage()

	diggers := maxLicenses
	cashChan := make(chan int, diggers*10)
	go func() {
		for {
			fmt.Printf("Minutes: %.1f. Cash %d. Cash queue %d.\n", time.Since(start).Minutes(), totalCash, len(cashChan))
			if isRemote {
				time.Sleep(5 * time.Minute)
			} else {
				time.Sleep(10 * time.Second)
			}
		}
	}()

	const width = 3500
	const step = 25
	const areasN = width * width / step / step
	var areas [areasN]models.ExploreResp

	var wg sync.WaitGroup
	var areaExplorers = cpus
	for w := 0; w < areaExplorers; w++ {
		wg.Add(1)
		go func(currentId int) {
			for i := currentId; i < areasN; i += areaExplorers {
				x := ((i * step) / width) * step
				y := (i * step) % width
				for {
					exploreRes, exploreErr := mineClient.Explore(&models.Area{PosX: x, PosY: y, SizeX: step, SizeY: step})
					if exploreErr == nil {
						areas[i] = exploreRes
						break
					}
				}
			}
			wg.Done()
		}(w)
	}
	wg.Wait()
	sort.SliceStable(areas[:], func(i, j int) bool {
		return areas[i].Amount > areas[j].Amount
	})

	var amounts [areasN]float64
	for i, value := range areas {
		amounts[i] = float64(value.Amount)
	}
	max, _ := stats.Max(amounts[:])
	min, _ := stats.Min(amounts[:])
	mean, _ := stats.Mean(amounts[:])
	fmt.Printf("Max = %0.2f. Min = %0.2f. Avg = %0.2f\n", max, min, mean)

	for w := 0; w < diggers; w++ {
		go explore2(mineClient, w, diggers, areas[:], cashChan)
	}

	select {}
}

func explore2(mineClient *client.MineClient, id int, total int, explores []models.ExploreResp, cashChan chan int) {
	var license *models.License
	for idx := id; idx < len(explores); idx += total {
		area := explores[idx]
		totalLeft := area.Amount
		for i := area.Area.PosX; i < area.Area.PosX+area.Area.SizeX; i++ {
			for j := area.Area.PosY; j < area.Area.PosY+area.Area.SizeY; j++ {
				for {
					exploreRes, exploreErr := mineClient.Explore(&models.Area{PosX: i, PosY: j, SizeX: 1, SizeY: 1})
					if exploreErr == nil {
						left := exploreRes.Amount
						totalLeft -= left
						for k := 1; k <= 10 && left > 0; {
							if license == nil || license.IsUsed() {
								license = issueNewLicense(mineClient, cashChan)
							}
							for {
								digResults, digErr := mineClient.Dig(exploreRes.Area.PosX, exploreRes.Area.PosY, k, license.Id)
								if digErr == nil {
									for _, gold := range digResults {
										cash(mineClient, gold, cashChan)
									}
									k++
									left -= len(digResults)
									license.UseOnce()
									break
								}
							}
						}
						break
					}
				}
			}
		}
	}
}

func issueNewLicense(mineClient *client.MineClient, cashChan chan int) *models.License {
	for {
		var cashList []int
		select {
		case cash := <-cashChan:
			cashList = []int{cash}
		default:
			cashList = []int{}
		}
		license, licenseErr := mineClient.IssueLicense(cashList)
		if licenseErr == nil {
			return &license
		}
	}
}

func cash(mineClient *client.MineClient, gold string, cashChan chan int) {
	for {
		cash, cashErr := mineClient.Cash(gold)
		if cashErr == nil {
			atomic.AddUint64(&totalCash, uint64(len(cash)))
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
