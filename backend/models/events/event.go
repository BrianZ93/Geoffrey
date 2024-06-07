package models

import "time"

type Event struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Guests    []Guest   `json:"guests"`
	Venues    []Venue   `json:"venues"`
}

type Venue struct {
	Id        string    `json:"id"`
	EventId   string    `json:"event_id"` // Foreign key to Event
	Title     string    `json:"title"`
	Address   string    `json:"address"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
