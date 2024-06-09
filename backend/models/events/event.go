package models

type Event struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	Guests    []Guest `json:"guests"`
	Venues    []Venue `json:"venues"`
}

type Venue struct {
	Id        string `json:"id"`
	EventId   string `json:"event_id"` // Foreign key to Event
	Title     string `json:"title"`
	Address   string `json:"address"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
