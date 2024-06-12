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
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func modifyGuestInDynamoDb(dynamoClient *dynamodb.Client, guest models.Guest, tableName string) error {
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
		logrus.Errorf("failed to update guest in DynamoDB: %v", err)
		return fmt.Errorf("failed to update guest in DynamoDB: %v", err)
	}

	logrus.Printf("Guest %s updated in DynamoDB successfully", guest.Id)
	return nil
}

func ModifyGuest(db *sql.DB, tableName string, dynamoClient *dynamodb.Client) echo.HandlerFunc {
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

		logrus.WithFields(logrus.Fields{
			"Name":         guest.Name,
			"Email":        guest.Email,
			"PhoneNumber":  guest.PhoneNumber,
			"Attending":    guest.Attending,
			"RsvpReceived": guest.RsvpReceived,
			"Note":         guest.Note,
			"GuestId":      guestId,
			"EventId":      eventId,
		}).Info("Updating guest details")

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
		} else {
			err = modifyGuestInDynamoDb(dynamoClient, guest, tableName)
			if err != nil {
				logrus.Errorf("Error modifying guest %s in dynamoDB database %v: ", guest.Name, err)
			}
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Guest details updated successfully"})
	}
}
