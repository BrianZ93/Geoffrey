package types

import (
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

// FetchVenueFromDynamoDB fetches a venue by ID from the DynamoDB table
func FetchVenueFromDynamoDB(dynamoClient *dynamodb.Client, venueID string, venuesTableName string) (*models.Venue, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(venuesTableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: venueID},
		},
	}

	result, err := dynamoClient.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var venue models.Venue
	err = attributevalue.UnmarshalMap(result.Item, &venue)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal venue: %v", err)
	}

	return &venue, nil
}

// UpdateVenueInSQLite updates the venue details in the SQLite database
func UpdateVenueInSQLite(db *sql.DB, venue models.Venue) error {
	query := `UPDATE Events_Venues SET Title = ?, Address = ?, StartTime = ?, EndTime = ? WHERE Id = ?`
	_, err := db.Exec(query, venue.Title, venue.Address, venue.StartTime, venue.EndTime, venue.Id)
	return err
}

// CompareVenueDetails compares individual venue details between SQLite and DynamoDB
func CompareVenueDetails(db *sql.DB, sqliteVenue models.Venue, dynamoVenue models.Venue) error {
	var needsUpdating bool

	if sqliteVenue.Title != dynamoVenue.Title {
		logrus.Infof("Venue title mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteVenue.Id, sqliteVenue.Title, dynamoVenue.Title)
		needsUpdating = true
	}
	if sqliteVenue.Address != dynamoVenue.Address {
		logrus.Infof("Venue address mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteVenue.Id, sqliteVenue.Address, dynamoVenue.Address)
		needsUpdating = true
	}
	if sqliteVenue.StartTime != dynamoVenue.StartTime {
		logrus.Infof("Venue start time mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteVenue.Id, sqliteVenue.StartTime, dynamoVenue.StartTime)
		needsUpdating = true
	}
	if sqliteVenue.EndTime != dynamoVenue.EndTime {
		logrus.Infof("Venue end time mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteVenue.Id, sqliteVenue.EndTime, dynamoVenue.EndTime)
		needsUpdating = true
	}

	if needsUpdating {
		logrus.Infof("Venue details mismatch for ID %s. Updating SQLite to match DynamoDB.", sqliteVenue.Id)
		err := UpdateVenueInSQLite(db, dynamoVenue)
		if err != nil {
			return fmt.Errorf("error updating venue %s in SQLite: %v", dynamoVenue.Id, err)
		}
		logrus.Infof("Venue %s updated successfully in SQLite.", dynamoVenue.Id)
	}

	return nil
}

// CheckVenueExists checks if a venue exists in the DynamoDB table
func CheckVenueExists(dynamoClient *dynamodb.Client, venueID string, venuesTableName string) (bool, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(venuesTableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: venueID},
		},
	}

	result, err := dynamoClient.GetItem(context.TODO(), input)
	if err != nil {
		return false, err
	}

	return result.Item != nil, nil
}

// WriteVenueToDynamoDB writes the given venue to the DynamoDB table
func WriteVenueToDynamoDB(dynamoClient *dynamodb.Client, venue models.Venue, venuesTableName string) error {
	// Marshal the venue into a map of DynamoDB attribute values
	av, err := attributevalue.MarshalMap(venue)
	if err != nil {
		logrus.Errorf("failed to marshal venue: %v", err)
		return fmt.Errorf("failed to marshal venue: %v", err)
	}

	// Log the marshaled venue for debugging
	logrus.Debugf("Marshaled venue: %v", av)

	// Create the PutItem input for the venue
	input := &dynamodb.PutItemInput{
		TableName: aws.String(venuesTableName),
		Item:      av,
	}

	// Log the PutItem input for debugging
	logrus.Debugf("PutItem input: %v", input)

	// Put the item into the DynamoDB table
	_, err = dynamoClient.PutItem(context.TODO(), input)
	if err != nil {
		logrus.Errorf("failed to put venue in DynamoDB: %v", err)
		return fmt.Errorf("failed to put venue in DynamoDB: %v", err)
	}

	logrus.Printf("Venue %s written to DynamoDB successfully", venue.Id)
	return nil
}
