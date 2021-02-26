package main

import (
	"goldrush/client"
	"os"
	"time"
)

func main() {
	address, isSet := os.LookupEnv("ADDRESS")
	if !isSet {
		address = "localhost"
	}
	mineClient := client.NewMineClient(address)

	go func() {
		for {
			time.Sleep(5 * time.Minute)
			mineClient.ReportMetrics()
		}
	}()

	licenseChan := make(chan int, 30)

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
}
