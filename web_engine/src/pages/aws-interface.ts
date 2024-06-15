import { ddbDocClient } from "./aws-config";
import { ScanCommand } from "@aws-sdk/lib-dynamodb";

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