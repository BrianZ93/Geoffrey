import AWS from "aws-sdk";

// Configure AWS SDK
AWS.config.update({
  region: "us-east-2",
  credentials: new AWS.CognitoIdentityCredentials({
    IdentityPoolId: "us-east-2:71602eb7-6fff-4ef3-b546-c66e06a558a3",
  }),
});

// Create DynamoDB service object
const docClient = new AWS.DynamoDB.DocumentClient();

export const fetchItems = async () => {
  const params = {
    TableName: "Events",
  };

  try {
    const data = await docClient.scan(params).promise();
    console.log(data.Items);
    return data.Items;
  } catch (error) {
    console.error(
      "Unable to scan the table. Error JSON:",
      JSON.stringify(error, null, 2)
    );
  }
};
