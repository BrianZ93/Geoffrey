package routes

import (
	"backend/api"
	models "backend/models/events"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func addEventToDynamoDb(dynamoClient *dynamodb.Client, event models.Event, tableName string) error {
	// Marshall the event into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(event)
	if err != nil {
		logrus.Errorf("failed to marshal event: %v", err)
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	// Log the marshaled guest for debugging
	logrus.Debugf("Marshaled event: %v", av)

	// Create the PutItem input for the event
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	// Log the PutItem input for debugging
	logrus.Debugf("PutItem input: %v", input)

	// Put the item into the DynamoDB table
	_, err = dynamoClient.PutItem(context.TODO(), input)
	if err != nil {
		logrus.Errorf("failed to put event in DynamoDB: %v", err)
		return fmt.Errorf("failed to put event in DynamoDB: %v", err)
	}

	logrus.Printf("Event %s written to DynamoDB successfully", event.Id)
	return nil

}

func CreateEvent(db *sql.DB, tableName string, dynamoClient *dynamodb.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var event models.Event
		if err := c.Bind(&event); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Generate a new UUID for the event, ignoring any ID received from the frontend
		event.Id = uuid.New().String()

		exists, err := api.TableExists(db, tableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if !exists {
			err = api.CreateTableFromStruct(db, tableName, models.Event{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Insert the event into the Events table
		_, err = db.Exec("INSERT INTO Events (Id, Title, StartTime, EndTime) VALUES (?, ?, ?, ?)",
			event.Id, event.Title, event.StartTime, event.EndTime)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Insert guests into the Guests table
		for _, guest := range event.Guests {
			guest.EventId = event.Id // Set the foreign key
			if _, err := db.Exec("INSERT INTO Events_Guests (EventId, Id, Name, Email, PhoneNumber, Attending, RsvpReceived, Note) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
				guest.EventId, guest.Id, guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Insert venues into the Venues table
		for _, venue := range event.Venues {
			venue.EventId = event.Id // Set the foreign key
			if _, err := db.Exec("INSERT INTO Events_Venues (EventId, Id, Title, Address, StartTime, EndTime) VALUES (?, ?, ?, ?, ?, ?)",
				venue.EventId, venue.Id, venue.Title, venue.Address, venue.StartTime, venue.EndTime); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		err = addEventToDynamoDb(dynamoClient, event, tableName)
		if err != nil {
			logrus.Errorf("Error adding event to dynamoDB table: %v", err)
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Event created successfully"})
	}
}
