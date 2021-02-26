package main

import (
	"goldrush/client"
	"goldrush/models"
	"os"
	"time"
)

type Coordinates struct {
	posX int
	posY int
}

func main() {
	address, isSet := os.LookupEnv("ADDRESS")
	if !isSet {
		address = "localhost"
	}
	mineClient := client.NewMineClient(address)

	licenseChan := make(chan int, 30)
	exploreChan := make(chan Coordinates, 200)
	digChan := make(chan models.ExploreResp, 50)
	goldChan := make(chan []string, 10)

	for w := 1; w <= 200; w++ {
		go explore(mineClient, exploreChan, digChan)
	}
	for w := 1; w <= 50; w++ {
		go dig(mineClient, digChan, licenseChan, goldChan)
	}
	for w := 1; w <= 10; w++ {
		go cash(mineClient, goldChan)
	}
	go reportMetrics(mineClient)

	go func() {
		for {
			currentLicenses := len(licenseChan)
			if currentLicenses < 30 && currentLicenses%3 == 0 {
				license, licenseErr := mineClient.IssueLicense()
				if licenseErr == nil {
					for l := 0; l < license.DigAllowed; l++ {
						licenseChan <- license.Id
					}
				}
			} else {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < 3500; i++ {
		for j := 0; j < 3500; j++ {
			exploreChan <- Coordinates{posX: i, posY: j}
		}
	}
}

func cash(mineClient *client.MineClient, goldChan chan []string) {
	for gold := range goldChan {
		for g := 0; g < len(gold); g++ {
			for {
				_, cashErr := mineClient.Cash(gold[g])
				if cashErr == nil {
					break
				}
			}
		}
	}
}

func dig(mineClient *client.MineClient, digChan chan models.ExploreResp, licenseChan chan int, goldChan chan []string) {
	for exploreRes := range digChan {
		left := exploreRes.Amount
		for k := 1; k <= 10 && left > 0; {
			licenseId := <-licenseChan
			for {
				digResult, digErr := mineClient.Dig(exploreRes.Area.PosX, exploreRes.Area.PosY, k, licenseId)
				if digErr == nil {
					goldChan <- digResult
					k++
					left -= len(digResult)
					break
				}
			}
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

func reportMetrics(mineClient *client.MineClient) {
	for {
		time.Sleep(1*time.Minute - 20*time.Millisecond)
		mineClient.ReportMetrics()
	}
}

/**

for i := 0; i < 3500; i++ {
	for j := 0; j < 3500; j++ {
		exploreRes, _ := mineClient.Explore(i, j)
		left := exploreRes.Amount
		for k := 1; k <= 10 && left > 0; {
			licenseId := <-licenseChan
			digResult, digErr := mineClient.Dig(i, j, k, licenseId)
			if digErr != nil {
				licenseChan <- licenseId
			} else {
				k++
				for g := 0; g < len(digResult); g++ {
					left--
					_, _ = mineClient.Cash(digResult[g])
				}
			}
		}
	}
}
*/
