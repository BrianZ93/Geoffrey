package routes

import (
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

func updateEventInDynamoDb(dynamoClient *dynamodb.Client, event models.Event, tableName string) error {
	// Marshal the event into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(event)
	if err != nil {
		logrus.Errorf("failed to marshal event: %v", err)
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	// Log the marshaled event for debugging
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
		logrus.Errorf("failed to update event in DynamoDB: %v", err)
		return fmt.Errorf("failed to update event in DynamoDB: %v", err)
	}

	logrus.Printf("Event %s updated in DynamoDB successfully", event.Id)
	return nil
}

func UpdateEvent(db *sql.DB, eventsTableName string, dynamoClient *dynamodb.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var event models.Event
		if err := c.Bind(&event); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Log the received event to check date format
		logrus.Infof("Received event: %+v", event)

		// Check if the event exists
		var eventExists bool
		err := db.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE Id = ?)", eventsTableName), event.Id).Scan(&eventExists)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !eventExists {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Event not found"})
		}

		// Update the event details
		logrus.Infof("Updating event: %+v", event)

		_, err = db.Exec(fmt.Sprintf("UPDATE %s SET Title = ?, StartTime = ?, EndTime = ? WHERE Id = ?", eventsTableName),
			event.Title, event.StartTime, event.EndTime, event.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Update guests if not null or empty
		if event.Guests != nil && len(event.Guests) > 0 {
			for _, guest := range event.Guests {
				guest.EventId = event.Id // Ensure the foreign key is set

				// Check if the guest exists
				var guestExists bool
				err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM Events_Guests WHERE Id = ?)", guest.Id).Scan(&guestExists)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}

				if guestExists {
					_, err = db.Exec("UPDATE Events_Guests SET Name = ?, Email = ?, PhoneNumber = ?, Attending = ?, RsvpReceived = ?, Note = ? WHERE Id = ?",
						guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note, guest.Id)
				} else {
					_, err = db.Exec("INSERT INTO Events_Guests (EventId, Id, Name, Email, PhoneNumber, Attending, RsvpReceived, Note) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
						guest.EventId, guest.Id, guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note)
				}

				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		// Update venues if not null or empty
		if event.Venues != nil && len(event.Venues) > 0 {
			for _, venue := range event.Venues {
				venue.EventId = event.Id // Ensure the foreign key is set

				// Check if the venue exists
				var venueExists bool
				err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM Events_Venues WHERE Id = ?)", venue.Id).Scan(&venueExists)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}

				if venueExists {
					_, err = db.Exec("UPDATE Events_Venues SET Title = ?, Address = ?, StartTime = ?, EndTime = ? WHERE Id = ?",
						venue.Title, venue.Address, venue.StartTime, venue.EndTime, venue.Id)
				} else {
					_, err = db.Exec("INSERT INTO Events_Venues (EventId, Id, Title, Address, StartTime, EndTime) VALUES (?, ?, ?, ?, ?, ?)",
						venue.EventId, venue.Id, venue.Title, venue.Address, venue.StartTime, venue.EndTime)
				}

				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		logrus.Infof("Event updated successfully: %+v", event)
		logrus.Info("Updating event in dynamoDB...")
		updateEventInDynamoDb(dynamoClient, event, eventsTableName)

		return c.JSON(http.StatusOK, map[string]string{"message": "Event updated successfully"})
	}
}
