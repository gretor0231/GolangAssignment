package model

type Gender string

type Person struct {
	Name        string `json:"name"`
	Height      int    `json:"height"`
	Gender      Gender `json:"gender"`
	WantedDates int    `json:"wantedDates"`
}
