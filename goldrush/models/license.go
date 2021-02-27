package models

import (
	"fmt"
	"math/rand"
	"sync"
)

type License struct {
	Id         int `json:"id"`
	DigAllowed int `json:"digAllowed"`
	DigUsed    int `json:"digUsed"`
}

func (license *License) IsUsed() bool {
	return license.DigAllowed == license.DigUsed
}

func (license *License) UseOnce() {
	license.DigUsed++
}

type Licenses struct {
	licenses map[int]License
	mu       sync.Mutex
}

func NewLicenses() *Licenses {
	return &Licenses{licenses: make(map[int]License)}
}

func (licenses *Licenses) Generate() License {
	licenses.mu.Lock()
	defer licenses.mu.Unlock()

	newLicense := License{rand.Intn(1000), 3, 0}
	licenses.licenses[newLicense.Id] = newLicense

	return newLicense
}

func (licenses *Licenses) Use(id int) License {
	licenses.mu.Lock()
	defer licenses.mu.Unlock()

	license := licenses.licenses[id]
	license.DigUsed += 1

	return license
}

func (licenses *Licenses) String() string {
	var values []License
	for _, v := range licenses.licenses {
		values = append(values, v)
	}
	return fmt.Sprintf("%+v", values)
}
