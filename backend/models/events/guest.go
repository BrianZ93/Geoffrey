package models

type Guest struct {
	Id           string `json:"id"`
	EventId      string `json:"event_id"` // Foreign key to Event
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Attending    bool   `json:"attending"`
	RsvpReceived bool   `json:"rsvpReceived"`
	Note         string `json:"note"`
}
