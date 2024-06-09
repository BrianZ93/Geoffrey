package routes

import (
	"backend/api"
	models "backend/models/events"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ModifyGuest(db *sql.DB, tableName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		eventId := c.Param("eventId")
		if eventId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Event ID is required"})
		}

		guestId := c.Param("guestId")
		if guestId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Guest ID is required"})
		}

		var guest models.Guest
		if err := c.Bind(&guest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		exists, err := api.TableExists(db, tableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if !exists {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Table not found"})
		}

		// Update the guest in the Events_Guests table
		_, err = db.Exec("UPDATE Events_Guests SET Name = ?, Email = ?, PhoneNumber = ?, Attending = ?, RsvpReceived = ?, Note = ? WHERE Id = ? AND EventId = ?",
			guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note, guestId, eventId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Guest details updated successfully"})
	}
}
