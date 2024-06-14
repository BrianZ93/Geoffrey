import { defineFunction } from "@aws-amplify/backend";

export const pullFromDyanamoDB = defineFunction({
  name: "my-first-function",
  entry: "./handler.ts",
});
