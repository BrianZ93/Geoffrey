package routes

import (
	"backend/api"
	models "backend/models/events"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddGuest(db *sql.DB, tableName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		eventId := c.Param("eventId")
		if eventId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Event ID is required"})
		}

		var guest models.Guest
		if err := c.Bind(&guest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Generate a new UUID for the guest
		guest.Id = uuid.New().String()
		guest.EventId = eventId

		exists, err := api.TableExists(db, tableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if !exists {
			err = api.CreateTableFromStruct(db, tableName, models.Guest{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Insert the guest into the Events_Guests table
		_, err = db.Exec("INSERT INTO Events_Guests (Id, EventId, Name, Email, PhoneNumber, Attending, RsvpReceived, Note) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			guest.Id, guest.EventId, guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Guest added successfully"})
	}
}
