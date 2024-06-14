import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { ddbDocClient } from "./pages/aws-config";
import { ScanCommand } from "@aws-sdk/lib-dynamodb";

function App() {
  const getAllItems = async (tableName: string) => {
    try {
      const params = {
        TableName: tableName,
      };
      const data = await ddbDocClient.send(new ScanCommand(params));
      console.log(data);
      return data.Items;
    } catch (err) {
      console.error("Error", err);
      throw new Error("Could not retrieve items");
    }
  };

  getAllItems("Events");

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
