package models

type Guest struct {
	Id           string `json:"id" dynamodbav:"id"`
	EventId      string `json:"event_id" dynamodbav:"event_id"` // Foreign key
	Name         string `json:"name" dynamodbav:"name"`
	Email        string `json:"email" dynamodbav:"email"`
	PhoneNumber  string `json:"phoneNumber" dynamodbav:"phoneNumber"`
	Attending    bool   `json:"attending" dynamodbav:"attending"`
	RsvpReceived bool   `json:"rsvpReceived" dynamodbav:"rsvpReceived"`
	Note         string `json:"note" dynamodbav:"note"`
}
