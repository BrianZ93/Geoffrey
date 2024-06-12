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

func addGuestToDynamoDB(dynamoClient *dynamodb.Client, guest models.Guest, tableName string) error {
	// Marshal the guest into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(guest)
	if err != nil {
		logrus.Errorf("failed to marshal guest: %v", err)
		return fmt.Errorf("failed to marshal guest: %v", err)
	}

	// Log the marshaled guest for debugging
	logrus.Debugf("Marshaled guest: %v", av)

	// Create the PutItem input for the guest
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	// Log the PutItem input for debugging
	logrus.Debugf("PutItem input: %v", input)

	// Put the item into the DynamoDB table
	_, err = dynamoClient.PutItem(context.TODO(), input)
	if err != nil {
		logrus.Errorf("failed to put guest in DynamoDB: %v", err)
		return fmt.Errorf("failed to put guest in DynamoDB: %v", err)
	}

	logrus.Printf("Guest %s written to DynamoDB successfully", guest.Id)
	return nil

}

func AddGuest(db *sql.DB, tableName string, dynamoClient *dynamodb.Client) echo.HandlerFunc {
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
		} else {
			// Add the guest into dynamoDB if it was successfully added into sqlite3
			err = addGuestToDynamoDB(dynamoClient, guest, tableName)
			if err != nil {
				logrus.Errorf("Error adding guest to dynamoDB table: %v", err)
			}
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Guest added successfully"})
	}
}
