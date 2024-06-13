// Load environment variables from .env file in development
if (window.location.hostname === "localhost") {
  const dotenv = require("dotenv");
  dotenv.config({ path: "../../../.env" });
}

// Determine the environment
const isDevelopment = window.location.hostname === "localhost";

// Set AWS configuration
AWS.config.update({
  region: isDevelopment ? process.env.AWS_REGION : "YOUR_PROD_AWS_REGION",
  credentials: new AWS.Credentials(
    isDevelopment
      ? process.env.AWS_ACCESS_KEY_ID
      : "YOUR_PROD_AWS_ACCESS_KEY_ID",
    isDevelopment
      ? process.env.AWS_SECRET_ACCESS_KEY
      : "YOUR_PROD_AWS_SECRET_ACCESS_KEY"
  ),
});

// Initialize the DynamoDB Document Client
const docClient = new AWS.DynamoDB.DocumentClient();

// Function to fetch events from DynamoDB
async function fetchEvents() {
  const params = {
    TableName: "Events",
  };

  try {
    const data = await docClient.scan(params).promise();
    return data.Items;
  } catch (error) {
    console.error("Error fetching events:", error);
    return [];
  }
}

// Export the fetchEvents function
module.exports = {
  fetchEvents,
};
