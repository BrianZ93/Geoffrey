package sync

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

// FetchGuestFromDynamoDB fetches a guest by ID from the DynamoDB table
func FetchGuestFromDynamoDB(svc *dynamodb.Client, guestID string, guestsTableName string) (*models.Guest, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(guestsTableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: guestID},
		},
	}

	result, err := svc.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var guest models.Guest
	err = attributevalue.UnmarshalMap(result.Item, &guest)
	if err != nil {
		logrus.Errorf("failed to unmarshal guest: %v", err)
		return nil, fmt.Errorf("failed to unmarshal guest: %v", err)
	}

	return &guest, nil
}

// UpdateGuestInSQLite updates the guest details in the SQLite database
func UpdateGuestInSQLite(db *sql.DB, guest models.Guest) error {
	query := `UPDATE Events_Guests SET Name = ?, Email = ?, PhoneNumber = ?, Attending = ?, RsvpReceived = ?, Note = ? WHERE Id = ?`
	_, err := db.Exec(query, guest.Name, guest.Email, guest.PhoneNumber, guest.Attending, guest.RsvpReceived, guest.Note, guest.Id)
	return err
}

// CompareGuestDetails compares individual guest details between SQLite and DynamoDB
func CompareGuestDetails(db *sql.DB, sqliteGuest models.Guest, dynamoGuest models.Guest) error {
	var needsUpdating bool

	if sqliteGuest.Name != dynamoGuest.Name {
		logrus.Infof("Guest name mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteGuest.Id, sqliteGuest.Name, dynamoGuest.Name)
		needsUpdating = true
	}
	if sqliteGuest.Email != dynamoGuest.Email {
		logrus.Infof("Guest email mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteGuest.Id, sqliteGuest.Email, dynamoGuest.Email)
		needsUpdating = true
	}
	if sqliteGuest.PhoneNumber != dynamoGuest.PhoneNumber {
		logrus.Infof("Guest phone number mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteGuest.Id, sqliteGuest.PhoneNumber, dynamoGuest.PhoneNumber)
		needsUpdating = true
	}
	if sqliteGuest.Attending != dynamoGuest.Attending {
		logrus.Infof("Guest attending status mismatch for ID %s: SQLite = %t, DynamoDB = %t", sqliteGuest.Id, sqliteGuest.Attending, dynamoGuest.Attending)
		needsUpdating = true
	}
	if sqliteGuest.RsvpReceived != dynamoGuest.RsvpReceived {
		logrus.Infof("Guest RSVP received status mismatch for ID %s: SQLite = %t, DynamoDB = %t", sqliteGuest.Id, sqliteGuest.RsvpReceived, dynamoGuest.RsvpReceived)
		needsUpdating = true
	}
	if sqliteGuest.Note != dynamoGuest.Note {
		logrus.Infof("Guest note mismatch for ID %s: SQLite = %s, DynamoDB = %s", sqliteGuest.Id, sqliteGuest.Note, dynamoGuest.Note)
		needsUpdating = true
	}

	if needsUpdating {
		logrus.Infof("Guest details mismatch for ID %s. Updating SQLite to match DynamoDB.", sqliteGuest.Id)
		err := UpdateGuestInSQLite(db, dynamoGuest)
		if err != nil {
			logrus.Errorf("error updating guest %s in SQLite: %v", dynamoGuest.Id, err)
			return fmt.Errorf("error updating guest %s in SQLite: %v", dynamoGuest.Id, err)
		}
		logrus.Infof("Guest %s updated successfully in SQLite.", dynamoGuest.Id)
	}

	return nil
}

// WriteGuestToDynamoDB writes the given guest to the DynamoDB table
func WriteGuestToDynamoDB(svc *dynamodb.Client, guest models.Guest, tableName string) error {
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
	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		logrus.Errorf("failed to put guest in DynamoDB: %v", err)
		return fmt.Errorf("failed to put guest in DynamoDB: %v", err)
	}

	logrus.Printf("Guest %s written to DynamoDB successfully", guest.Id)
	return nil
}
