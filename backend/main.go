package main

import (
	"backend/api"
	eventroutes "backend/api/routes/events"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	// Create the logger and set formatting
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Starting backend")

	// Create the database connection
	dbPath := "./data/database.sqlite3"
	db, err := api.InitDb(dbPath)
	if err != nil {
		logrus.Errorf("Error creating database at %s:", dbPath)
	}

	// Setup Echo
	e := echo.New()

	// Enable CORS for frontend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:9000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Setup API Routes
	// Event Routes
	eventsTableName := "Events"
	guestsTableName := "Events_Guests"
	e.POST("/create-event", eventroutes.CreateEvent(db, eventsTableName))
	e.GET("/events", eventroutes.GetEvents(db, eventsTableName))
	e.PUT("/events/:eventId", eventroutes.UpdateEvent(db, eventsTableName))

	// Guests Routes
	e.POST("/events/:eventId/guests", eventroutes.AddGuest(db, guestsTableName))

	// Starting the backend server
	logrus.Info("Starting server on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
