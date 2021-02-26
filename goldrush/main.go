package main

import (
	"fmt"
	"goldrush/client"
	"goldrush/models"
	"goldrush/utils"
	"os"
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

	exploreChan := make(chan Coordinates)
	digChan := make(chan models.ExploreResp, 200)
	goldChan := make(chan string, 100)
	cashChan := make(chan int, 1000)

	getLicenseLease := issueLicense(cashChan, mineClient)

	for w := 1; w <= 200; w++ {
		go explore(mineClient, exploreChan, digChan)
	}
	for w := 1; w <= 100; w++ {
		go dig(mineClient, digChan, getLicenseLease, goldChan)
	}
	for w := 1; w <= 20; w++ {
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
			if processed%10000 == 0 {
				fmt.Printf("Processed %d\n", processed)
			}
		}
	}
}

func issueLicense(cashChan chan int, mineClient *client.MineClient) func(callback func(int)) {
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
		if isRemote {
			time.Sleep(5*time.Minute - 10*time.Second)
		} else {
			time.Sleep(1*time.Minute - 10*time.Second)
		}

		fmt.Println("----------")
		utils.PrintMemoryUsage()
		utils.PrintCpuUsage()
		utils.PrintAvgUsage()
		mineClient.PrintMetrics()
		fmt.Println("----------")
	}
}
