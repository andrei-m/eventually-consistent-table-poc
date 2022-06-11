package server

import "time"

// A basic data model for automobiles and sale events. For the sake of the POC, assume that Autos are maintained in an OLTP database and Sales are maintained separately in an OLAP database. Costs are modeled as whole dollar amounts for simplicity.

type Auto struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	ModelName string `json:"modelName"`
}

type Sale struct {
	AutoID   int       `json:"autoId"`
	Time     time.Time `json:"timestamp"`
	MsrpUSD  int       `json:"msrpUsd"`
	PriceUSD int       `json:"priceUsd"`
}
