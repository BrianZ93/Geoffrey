// aws-config.ts
import { DynamoDBClient } from "@aws-sdk/client-dynamodb";
import { DynamoDBDocumentClient } from "@aws-sdk/lib-dynamodb";

// Environment variables for local development
const REGION = process.env.REACT_APP_AWS_REGION;
const ACCESS_KEY_ID = process.env.REACT_APP_AWS_ACCESS_KEY_ID;
const SECRET_ACCESS_KEY = process.env.REACT_APP_AWS_SECRET_ACCESS_KEY;

// Function to create DynamoDB client
const createDynamoDBClient = (): DynamoDBClient => {
  if (REGION && ACCESS_KEY_ID && SECRET_ACCESS_KEY) {
    // Local environment
    return new DynamoDBClient({
      region: REGION,
      credentials: {
        accessKeyId: ACCESS_KEY_ID!,
        secretAccessKey: SECRET_ACCESS_KEY!,
      },
    });
  } else {
    // Amplify environment
    return new DynamoDBClient({
      region: "us-east-2", // Amplify sets this automatically
    });
  }
};

// DynamoDB Client
const ddbClient = createDynamoDBClient();
const ddbDocClient = DynamoDBDocumentClient.from(ddbClient);

export { ddbDocClient };
