import { ddbDocClient } from "./aws-config";
import { ScanCommand, UpdateCommand, UpdateCommandInput } from "@aws-sdk/lib-dynamodb";

export const getAllItems = async (tableName: string) => {
    try {
      const params = {
        TableName: tableName,
      };
      const data = await ddbDocClient.send(new ScanCommand(params));
      console.log(data);
      return data.Items || [];
    } catch (err) {
      console.error("Error", err);
      throw new Error("Could not retrieve items");
    }
  };

/**
 * Update a guest in the DynamoDB table.
 * 
 * @param {string} tableName - The name of the DynamoDB table.
 * @param {string} guestId - The ID of the guest to update.
 * @param {Object} updates - The updates to apply to the guest.
 * @returns {Promise} - A promise that resolves when the update is complete.
 */
export async function updateGuest(tableName: string, guestId: string, updates: { attending: boolean }) {
    const params: UpdateCommandInput = {
      TableName: tableName,
      Key: {
        id: guestId
      },
      UpdateExpression: 'set #attending = :attending',
      ExpressionAttributeNames: {
        '#attending': 'attending'
      },
      ExpressionAttributeValues: {
        ':attending': updates.attending
      },
      ReturnValues: 'UPDATED_NEW'
    };
  
    try {
      const command = new UpdateCommand(params);
      const result = await ddbDocClient.send(command);
      console.log('Update succeeded:', result);
      return result;
    } catch (error) {
      console.error('Update failed:', error);
      throw error;
    }
  }