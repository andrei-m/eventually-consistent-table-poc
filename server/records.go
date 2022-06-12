package server

import (
	"math/rand"
	"sync"
	"time"
)

// A basic data model for automobiles and sale events. For the sake of the POC, assume that Autos are maintained in an OLTP database and Sales are maintained separately in an OLAP database. Costs are modeled as whole dollar amounts for simplicity.

type Auto struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	ModelName string `json:"modelName"`
}

type AutoDB struct {
	lock  sync.Mutex
	autos []Auto
}

func (a *AutoDB) NewRandomizedAuto() Auto {
	a.lock.Lock()
	defer a.lock.Unlock()
	auto := Auto{
		ID:        len(a.autos) + 1,
		Brand:     randStringBytes(8),
		ModelName: randStringBytes(8),
	}
	a.autos = append(a.autos, auto)
	return auto
}

func (a *AutoDB) GetAutos() []Auto {
	return a.autos
}

type Sale struct {
	AutoID   int       `json:"autoId"`
	Time     time.Time `json:"timestamp"`
	MsrpUSD  int       `json:"msrpUsd"`
	PriceUSD int       `json:"priceUsd"`
}

type SaleDB struct {
	lock  sync.Mutex
	sales []Sale
}

func (s *SaleDB) NewRandomizedSale(autoID int) Sale {
	s.lock.Lock()
	defer s.lock.Unlock()
	sale := Sale{
		AutoID:   autoID,
		Time:     time.Now(),
		MsrpUSD:  rand.Intn(10000),
		PriceUSD: rand.Intn(10000),
	}
	s.sales = append(s.sales, sale)
	return sale

}
