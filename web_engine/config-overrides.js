const CopyWebpackPlugin = require("copy-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const path = require("path");
const fs = require("fs");

module.exports = function override(config, env) {
  // Determine which website to use based on environment variable
  const website = process.env.REACT_APP_WEBSITE || "bachelor_party";

  // Define paths for each website
  const websiteConfigs = {
    bachelor_party: [
      {
        from: path.resolve(
          __dirname,
          "public/bachelor_party/public_html/bachelor_party.html"
        ),
        to: path.resolve(__dirname, "build/index.html"),
      },
      {
        from: path.resolve(
          __dirname,
          "public/bachelor_party/public_html/assets"
        ),
        to: path.resolve(__dirname, "build/assets"),
      },
      // Add other directories if needed
    ],
    happily_ever_after_party: [
      {
        from: path.resolve(
          __dirname,
          "public/happily_ever_after_party/happily_ever_after_party.html"
        ),
        to: path.resolve(__dirname, "build/index.html"),
      },
      {
        from: path.resolve(__dirname, "public/happily_ever_after_party/css"),
        to: path.resolve(__dirname, "build/css"),
      },
      {
        from: path.resolve(__dirname, "public/happily_ever_after_party/img"),
        to: path.resolve(__dirname, "build/img"),
      },
      {
        from: path.resolve(__dirname, "public/happily_ever_after_party/js"),
        to: path.resolve(__dirname, "build/js"),
      },
      // Add other directories if needed
    ],
  };

  // Select the correct configuration based on the website
  const selectedConfig = websiteConfigs[website];

  // Remove the default HtmlWebpackPlugin instance
  config.plugins = config.plugins.filter(
    (plugin) => !(plugin instanceof HtmlWebpackPlugin)
  );

  // Add the CopyWebpackPlugin to the plugins array
  config.plugins.push(
    new CopyWebpackPlugin({
      patterns: selectedConfig,
    })
  );

  return config;
};
