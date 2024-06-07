package routes

import (
	models "backend/models/events"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateEvent(db *sql.DB, tableName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var event models.Event
		if err := c.Bind(&event); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Check if the event exists
		var eventExists bool
		err := db.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE Id = ?)", tableName), event.Id).Scan(&eventExists)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !eventExists {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Event not found"})
		}

		// Update the event details
		_, err = db.Exec(fmt.Sprintf("UPDATE %s SET Title = ?, StartTime = ?, EndTime = ? WHERE Id = ?", tableName),
			event.Title, event.StartTime, event.EndTime, event.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Update guests if not null or empty
		if event.Guests != nil && len(event.Guests) > 0 {
			for _, guest := range event.Guests {
				guest.EventId = event.Id // Ensure the foreign key is set
				_, err = db.Exec("REPLACE INTO Events_Guests (EventId, Id, Name, Email, PhoneNumber, Attending, RsvpReceived, Note) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
					guest.EventId, guest.Id, guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		// Update venues if not null or empty
		if event.Venues != nil && len(event.Venues) > 0 {
			for _, venue := range event.Venues {
				venue.EventId = event.Id // Ensure the foreign key is set
				_, err = db.Exec("REPLACE INTO Events_Venues (EventId, Id, Title, Address, StartTime, EndTime) VALUES (?, ?, ?, ?, ?, ?)",
					venue.EventId, venue.Id, venue.Title, venue.Address, venue.StartTime, venue.EndTime)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Event updated successfully"})
	}
}
