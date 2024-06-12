package models

type Event struct {
	Id        string  `json:"id" dynamodbav:"id"`
	Title     string  `json:"title" dynamodbav:"title"`
	StartTime string  `json:"startTime" dynamodbav:"startTime"`
	EndTime   string  `json:"endTime" dynamodbav:"endTime"`
	Guests    []Guest `json:"guests" dynamodbav:"guests"`
	Venues    []Venue `json:"venues" dynamodbav:"venues"`
}

type Venue struct {
	Id        string `json:"id" dynamodbav:"id"`
	EventId   string `json:"event_id" dynamodbav:"event_id"` /// Foreign Key
	Title     string `json:"title" dynamodbav:"title"`
	Address   string `json:"address" dynamodbav:"address"`
	StartTime string `json:"startTime" dynamodbav:"startTime"`
	EndTime   string `json:"endTime" dynamodbav:"endTime"`
}
