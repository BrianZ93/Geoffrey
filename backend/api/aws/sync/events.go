package sync

import (
	routes "backend/api/routes/events"
	models "backend/models/events"
	"context"
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
)

// CheckGuestExists checks if a guest exists in the DynamoDB table
func CheckGuestExists(svc *dynamodb.Client, guestID string, guestsTableName string) (bool, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(guestsTableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: guestID},
		},
	}

	result, err := svc.GetItem(context.TODO(), input)
	if err != nil {
		return false, err
	}

	return result.Item != nil, nil
}

// CompareEventDetails compares guests and venues for a given event between SQLite and DynamoDB
func CompareEventDetails(db *sql.DB, svc *dynamodb.Client, eventID string, guestsTableName string, venuesTableName string) error {
	// Fetch event details from SQLite
	sqliteEvent, err := routes.GetEventByID(db, eventID)
	if err != nil {
		logrus.Errorf("error fetching event from SQLite: %v", err)
		return fmt.Errorf("error fetching event from SQLite: %v", err)
	}

	// Check each guest in SQLite event against DynamoDB
	for _, sqliteGuest := range sqliteEvent.Guests {
		dynamoGuest, err := FetchGuestFromDynamoDB(svc, sqliteGuest.Id, guestsTableName)
		if err != nil {
			logrus.Errorf("error fetching guest %s from DynamoDB: %v", sqliteGuest.Id, err)
			return fmt.Errorf("error fetching guest %s from DynamoDB: %v", sqliteGuest.Id, err)
		}
		if dynamoGuest == nil {
			logrus.Infof("Guest %s for event %s does not exist in DynamoDB.", sqliteGuest.Name, sqliteEvent.Title)
		} else {
			CompareGuestDetails(db, sqliteGuest, *dynamoGuest)
		}
	}

	// Check each venue in SQLite event against DynamoDB
	for _, sqliteVenue := range sqliteEvent.Venues {
		dynamoVenue, err := FetchVenueFromDynamoDB(svc, sqliteVenue.Id, venuesTableName)
		if err != nil {
			logrus.Errorf("error fetching venue %s from DynamoDB: %v", sqliteVenue.Id, err)
			return fmt.Errorf("error fetching venue %s from DynamoDB: %v", sqliteVenue.Id, err)
		}
		if dynamoVenue == nil {
			logrus.Infof("Venue %s for event %s does not exist in DynamoDB.", sqliteVenue.Title, sqliteEvent.Title)
		} else {
			CompareVenueDetails(db, sqliteVenue, *dynamoVenue)
		}
	}

	return nil
}

// WriteEventToDynamoDB writes the given event to the DynamoDB table
func WriteEventToDynamoDB(svc *dynamodb.Client, event models.Event, eventsTableName string) error {
	// Marshal the event into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(event)
	if err != nil {
		logrus.Errorf("failed to marshal event: %v", err)
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	// Check if 'id' field is present
	if _, ok := av["id"]; !ok {
		logrus.Errorf("missing 'id' field in marshaled event: %v", av)
		logrus.Errorf("original event: %+v", event)
		return fmt.Errorf("missing 'id' field in event: %v", event)
	}

	// Log the marshaled event for debugging
	logrus.Debugf("Marshaled event: %v", av)

	// Create the PutItem input
	input := &dynamodb.PutItemInput{
		TableName: aws.String(eventsTableName),
		Item:      av,
	}

	// Log the PutItem input for debugging
	logrus.Debugf("PutItem input: %v", input)

	// Put the item into the DynamoDB table
	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		logrus.Errorf("failed to put item in DynamoDB: %v", err)
		return fmt.Errorf("failed to put item in DynamoDB: %v", err)
	}

	logrus.Printf("Event %s written to DynamoDB successfully", event.Id)

	return nil
}

// FetchEventsFromDynamoDB fetches all events from the DynamoDB table
func FetchEventsFromDynamoDB(svc *dynamodb.Client) ([]models.Event, error) {
	var events []models.Event

	// Create the Scan input
	input := &dynamodb.ScanInput{
		TableName: aws.String("Events"),
	}

	// Scan the DynamoDB table
	result, err := svc.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %v", err)
	}

	// Unmarshal the results
	for _, item := range result.Items {
		var event models.Event
		err = attributevalue.UnmarshalMap(item, &event)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal item: %v", err)
		}
		events = append(events, event)
	}

	return events, nil
}
