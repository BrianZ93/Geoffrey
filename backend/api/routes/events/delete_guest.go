package routes

import (
	"backend/api"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteGuest(db *sql.DB, tableName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		guestId := c.Param("guestId")
		if guestId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Guest ID is required"})
		}

		// Check if the table exists
		exists, err := api.TableExists(db, tableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if !exists {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Table does not exist"})
		}

		// Delete the guest from the Events_Guests table
		_, err = db.Exec("DELETE FROM Events_Guests WHERE Id = ?", guestId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Guest deleted successfully"})
	}
}
