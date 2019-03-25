package models

import (
	"time"
)

//CurrencyState struct
type CurrencyState struct {
	SourceID      string         `json:"sourceId"`
	UpdateDate    time.Time      `json:"date"`
	Organizations []Organization `json:"organizations"`
}

//Organization for currency
type Organization struct {
	id         string                       `json:"id"`
	Title      string                       `json:"title"`
	RegionID   string                       `json:"regionId"`
	CityID     string                       `json:"cityId"`
	Phone      string                       `json:"phone"`
	Address    string                       `json:"address"`
	Link       string                       `json:"link"`
	Currencies map[string]map[string]string `json:"currencies"`
}
