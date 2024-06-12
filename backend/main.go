package main

import (
	"backend/api"
	"backend/api/aws"
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

	logrus.SetLevel(logrus.DebugLevel)
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
	// Table Names
	eventsTableName := "Events"
	guestsTableName := "Events_Guests"
	venuesTableName := "Events_Venues"

	// Create dynamoDB interface
	dynamoClient := aws.InitDynamoDBClient()
	// Create events and events_guests table if it does not exist
	err = aws.CreateTable(dynamoClient, eventsTableName)
	if err != nil {
		logrus.Fatalf("Error creating table: %v", err)
	}
	err = aws.CreateTable(dynamoClient, guestsTableName)
	if err != nil {
		logrus.Fatalf("Error creating table: %v", err)
	}
	err = aws.CreateTable(dynamoClient, venuesTableName)
	if err != nil {
		logrus.Fatalf("Error creating table: %v", err)
	}

	// Grab any events that exist locally and write them to DynamoDB
	// *** This should really only happen if the dynamoDB table gets wiped ***
	aws.SyncEventsFromSqLite3ToDynamoDB(db, dynamoClient, eventsTableName, guestsTableName, venuesTableName)

	// Event Routes
	e.POST("/create-event", eventroutes.CreateEvent(db, eventsTableName, dynamoClient))
	e.GET("/events", eventroutes.GetEventsHandler(db, eventsTableName))
	e.PUT("/events/:eventId", eventroutes.UpdateEvent(db, eventsTableName, dynamoClient))

	// Guests Routes
	e.POST("/events/:eventId/guests", eventroutes.AddGuest(db, guestsTableName, dynamoClient))
	e.PUT("/events/:eventId/guests/:guestId", eventroutes.ModifyGuest(db, guestsTableName, dynamoClient))
	e.DELETE("/events/:eventId/guests/:guestId", eventroutes.DeleteGuest(db, eventsTableName, dynamoClient))

	// Starting the backend server
	logrus.Info("Starting server on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
