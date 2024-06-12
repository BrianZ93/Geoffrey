package routes

import (
	"backend/api"
	models "backend/models/events"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetEventByID(db *sql.DB, eventID string) (models.Event, error) {
	var event models.Event

	// Fetch the event details
	query := `SELECT Id, Title, StartTime, EndTime FROM Events WHERE Id = ?`
	err := db.QueryRow(query, eventID).Scan(&event.Id, &event.Title, &event.StartTime, &event.EndTime)
	if err != nil {
		return event, err
	}

	// Fetch the guests for the event
	guestsQuery := `SELECT Id, EventId, Name, Email, PhoneNumber, Attending, RsvpReceived, Note FROM Events_Guests WHERE EventId = ?`
	guestsRows, err := db.Query(guestsQuery, eventID)
	if err != nil {
		return event, err
	}
	defer guestsRows.Close()

	for guestsRows.Next() {
		var guest models.Guest
		if err := guestsRows.Scan(&guest.Id, &guest.EventId, &guest.Name, &guest.Email, &guest.PhoneNumber, &guest.Attending, &guest.RsvpReceived, &guest.Note); err != nil {
			return event, err
		}
		event.Guests = append(event.Guests, guest)
	}

	// Fetch the venues for the event
	venuesQuery := `SELECT Id, EventId, Title, Address, StartTime, EndTime FROM Events_Venues WHERE EventId = ?`
	venuesRows, err := db.Query(venuesQuery, eventID)
	if err != nil {
		return event, err
	}
	defer venuesRows.Close()

	for venuesRows.Next() {
		var venue models.Venue
		if err := venuesRows.Scan(&venue.Id, &venue.EventId, &venue.Title, &venue.Address, &venue.StartTime, &venue.EndTime); err != nil {
			return event, err
		}
		event.Venues = append(event.Venues, venue)
	}

	return event, nil
}

func fetchGuestsByEventID(db *sql.DB, eventID string) ([]models.Guest, error) {
	rows, err := db.Query("SELECT Id, EventId, Name, Email, PhoneNumber, Attending, RsvpReceived, Note FROM Events_Guests WHERE EventId = ?", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []models.Guest
	for rows.Next() {
		var guest models.Guest
		if err := rows.Scan(&guest.Id, &guest.EventId, &guest.Name, &guest.Email, &guest.PhoneNumber, &guest.Attending, &guest.RsvpReceived, &guest.Note); err != nil {
			return nil, err
		}
		guests = append(guests, guest)
	}

	return guests, nil
}

func fetchVenuesByEventID(db *sql.DB, eventID string) ([]models.Venue, error) {
	rows, err := db.Query("SELECT Id, EventId, Title, Address, StartTime, EndTime FROM Events_Venues WHERE EventId = ?", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var venues []models.Venue
	for rows.Next() {
		var venue models.Venue
		if err := rows.Scan(&venue.Id, &venue.EventId, &venue.Title, &venue.Address, &venue.StartTime, &venue.EndTime); err != nil {
			return nil, err
		}
		venues = append(venues, venue)
	}

	return venues, nil
}

func fetchAllEvents(db *sql.DB, tableName string) ([]models.Event, error) {
	query := fmt.Sprintf("SELECT Id, Title, StartTime, EndTime FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.Id, &event.Title, &event.StartTime, &event.EndTime); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvents(db *sql.DB, tableName string) ([]models.Event, error) {
	// First check if the table exists
	exists, err := api.TableExists(db, tableName)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no such table exists: %s", tableName)
	}

	// Fetch all events
	events, err := fetchAllEvents(db, tableName)
	if err != nil {
		return nil, err
	}

	// Fetch related guests and venues for each event
	for i := range events {
		events[i].Guests, err = fetchGuestsByEventID(db, events[i].Id)
		if err != nil {
			return nil, err
		}

		events[i].Venues, err = fetchVenuesByEventID(db, events[i].Id)
		if err != nil {
			return nil, err
		}
	}

	return events, nil
}

func GetEventsHandler(db *sql.DB, tableName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := GetEvents(db, tableName)
		if err != nil {
			logrus.Errorf("error: %v %v", http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, events)
	}
}
