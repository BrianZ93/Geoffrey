package aws

import (
	"backend/api/aws/sync"
	routes "backend/api/routes/events"
	models "backend/models/events"
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	logrus "github.com/sirupsen/logrus"
)

// SyncEvents syncs events from SQLite to DynamoDB
func SyncEventsFromSqLite3ToDynamoDB(db *sql.DB, svc *dynamodb.Client, eventsTableName string, guestsTableName string, venuesTableName string) error {
	// Fetch events from SQLite
	sqliteEvents, err := routes.GetEvents(db, eventsTableName)
	if err != nil {
		logrus.Errorf("error fetching events from SQLite: %v", err)
		return fmt.Errorf("error fetching events from SQLite: %v", err)
	}

	// Fetch events from DynamoDB
	dynamoEvents, err := sync.FetchEventsFromDynamoDB(svc)
	if err != nil {
		logrus.Errorf("error fetching events from DynamoDB: %v", err)
		return fmt.Errorf("error fetching events from DynamoDB: %v", err)
	}

	// Create a map of DynamoDB events for easy lookup
	dynamoEventMap := make(map[string]models.Event)
	for _, event := range dynamoEvents {
		dynamoEventMap[event.Id] = event
	}

	// Compare and add missing events, guests, and venues to DynamoDB
	for _, event := range sqliteEvents {
		if _, exists := dynamoEventMap[event.Id]; !exists {
			logrus.Infof("Event %s found in local Sqlite3 database but not in dynamoDB table %s, adding event...", event.Title, eventsTableName)
			err := sync.WriteEventToDynamoDB(svc, event, eventsTableName)
			if err != nil {
				logrus.Errorf("error writing event to DynamoDB: %v", err)
				return fmt.Errorf("error writing event to DynamoDB: %v", err)
			}
		}

		// Check if all Guests exist and add them if they do not
		for _, guest := range event.Guests {
			exists, err := sync.CheckGuestExists(svc, guest.Id, guestsTableName)
			if err != nil {
				logrus.Errorf("error checking if guest %s exists in DynamoDB: %v", guest.Id, err)
				return fmt.Errorf("error checking if guest %s exists in DynamoDB: %v", guest.Id, err)
			}
			if !exists {
				logrus.Infof("Guest %s for event %s does not exist in DynamoDB. Adding guest...", guest.Name, event.Title)
				err := sync.WriteGuestToDynamoDB(svc, guest, guestsTableName)
				if err != nil {
					logrus.Errorf("error writing guest to DynamoDB: %v", err)
					return fmt.Errorf("error writing guest to DynamoDB: %v", err)
				}
			}
		}

		logrus.Info("Guest Data Validated")

		// Check if all Venues exist and add them if they do not
		for _, venue := range event.Venues {
			exists, err := sync.CheckVenueExists(svc, venue.Id, venuesTableName)
			if err != nil {
				logrus.Errorf("error checking if venue %s exists in DynamoDB: %v", venue.Id, err)
				return fmt.Errorf("error checking if venue %s exists in DynamoDB: %v", venue.Id, err)
			}
			if !exists {
				logrus.Infof("Venue %s for event %s does not exist in DynamoDB. Adding venue...", venue.Title, event.Title)
				err := sync.WriteVenueToDynamoDB(svc, venue, venuesTableName)
				if err != nil {
					logrus.Errorf("error writing venue to DynamoDB: %v", err)
					return fmt.Errorf("error writing venue to DynamoDB: %v", err)
				}
			}
		}

		logrus.Info("Venues Data Validated")
	}

	return nil
}
