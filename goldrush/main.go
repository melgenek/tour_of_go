package main

import (
	"fmt"
	"goldrush/client"
	"os"
	"sync"
	"time"
)

var (
	errorOccurrences = map[string]int{}
	errorMutex       = sync.RWMutex{}
)

func printErr(e error) {
	if e != nil {
		message := e.Error()
		errorMutex.Lock()
		defer errorMutex.Unlock()
		errorOccurrences[message]++
		if errorOccurrences[message] < 5 || errorOccurrences[message]%100 == 0 {
			fmt.Println(e, errorOccurrences[message])
		}
	}
}

func main() {
	address, isSet := os.LookupEnv("ADDRESS")
	mineClient := client.MineClient{Host: "localhost"}
	if isSet {
		mineClient = client.MineClient{Host: address}
	}

	licenseChan := make(chan int, 30)

	go func() {
		for {
			currentLicenses := len(licenseChan)
			if currentLicenses < 30 && currentLicenses%3 == 0 {
				license, licenseErr := mineClient.IssueLicense()
				if licenseErr != nil {
					printErr(licenseErr)
				} else {
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
			exploreRes, exploreErr := mineClient.Explore(i, j)
			printErr(exploreErr)

			left := exploreRes.Amount
			for k := 1; k <= 10 && left > 0; {
				licenseId := <-licenseChan
				digResult, digErr := mineClient.Dig(i, j, k, licenseId)
				if digErr != nil {
					printErr(digErr)
					licenseChan <- licenseId
				} else {
					k++
					for g := 0; g < len(digResult); g++ {
						left--
						_, goldErr := mineClient.Cash(digResult[g])
						printErr(goldErr)
					}
				}
			}
		}
	}
}
