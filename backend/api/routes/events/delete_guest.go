package routes

import (
	"backend/api"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func deleteGuestFromDynamoDB(dynamoClient *dynamodb.Client, guestID string, tableName string) error {
	// Create the DeleteItem input
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: guestID},
		},
	}

	// Log the DeleteItem input for debugging
	logrus.Debugf("DeleteItem input: %v", input)

	// Delete the item from the DynamoDB table
	_, err := dynamoClient.DeleteItem(context.TODO(), input)
	if err != nil {
		logrus.Errorf("failed to delete guest from DynamoDB: %v", err)
		return fmt.Errorf("failed to delete guest from DynamoDB: %v", err)
	}

	logrus.Printf("Guest %s deleted from DynamoDB successfully", guestID)
	return nil
}

func DeleteGuest(db *sql.DB, tableName string, dynamoClient *dynamodb.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		guestId := c.Param("guestId")
		if guestId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Guest ID is required"})
		}

		// Check if the table exists
		exists, err := api.TableExists(db, tableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if !exists {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Table does not exist"})
		}

		// Delete the guest from the Events_Guests table
		_, err = db.Exec("DELETE FROM Events_Guests WHERE Id = ?", guestId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		} else {
			err = deleteGuestFromDynamoDB(dynamoClient, guestId, tableName)
			if err != nil {
				logrus.Errorf("Error deleting guest from dynamoDB: %v", err)
			}
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Guest deleted successfully"})
	}
}
