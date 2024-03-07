package model

type person struct {
	Name              string `json:"name"`
	Height            int    `json:"height"`
	Gender            string `json:"gender"`
	WantedDatesNumber int    `json:"wantedDatesNumber"`
}
