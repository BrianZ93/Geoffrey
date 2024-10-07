package main

import (
	"backend/api"
	"backend/api/aws"
	cookbookroutes "backend/api/routes/cookbook"
	eventroutes "backend/api/routes/events"
	financesroutes "backend/api/routes/finances"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	// Create the logger and set formatting
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting backend")

	// Create the sql database connection for events
	dbPath := "./data/database.sqlite3"
	eventsDb, err := api.InitDb(dbPath)
	if err != nil {
		logrus.Errorf("Error creating database at %s:", dbPath)
	}

	// Create the sql database connection for finances
	dbPath = "./data/finances.sqlite3"
	financesDb, err := api.InitDb(dbPath)
	if err != nil {
		logrus.Errorf("Error creating database at %s:", dbPath)
	}

	// Create the sql database connection for cookbook
	dbPath = "./data/cookbook.sqlite3"
	cookbookDB, err := api.InitDb(dbPath)
	if err != nil {
		logrus.Errorf("Error creating database at %s:", dbPath)
	}

	// Setup Echo
	e := echo.New()

	// Enable CORS for frontend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:9000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Setup API Routes
	// Table Names
	// Events
	eventsTableName := "Events"
	guestsTableName := "Events_Guests"
	venuesTableName := "Events_Venues"

	// Finances
	// financesTableName := "Finances"
	propertiesTableName := "Properties"
	mortgagesTableName := "Mortgages"
	tenantsTableName := "Tenants"
	incomeTableName := "Income"
	expensesTableName := "Expenses"

	// Cookbook
	recipesTableName := "Recipes"
	ingredientsTableName := "Ingredients"
	recipeStepsTableName := "RecipeSteps"

	// Create dynamoDB interface
	dynamoClient := aws.InitDynamoDBClient()
	// Create events and events_guests table if it does not exist **ONLY FOR AWS DB**
	err = aws.CreateTable(dynamoClient, eventsTableName)
	if err != nil {
		logrus.Fatalf("Error creating table: %v", err)
	}
	err = aws.CreateTable(dynamoClient, guestsTableName)
	if err != nil {
		logrus.Fatalf("Error creating table: %v", err)
	}
	err = aws.CreateTable(dynamoClient, venuesTableName)
	if err != nil {
		logrus.Fatalf("Error creating table: %v", err)
	}

	// Grab any events that exist locally and write them to DynamoDB
	// *** This should really only happen if the dynamoDB table gets wiped ***
	aws.SyncEventsFromSqLite3ToDynamoDB(eventsDb, dynamoClient, eventsTableName, guestsTableName, venuesTableName)

	// EVENT ROUTES
	e.POST("/create-event", eventroutes.CreateEvent(eventsDb, eventsTableName, dynamoClient))
	e.GET("/events", eventroutes.GetEventsHandler(eventsDb, eventsTableName))
	e.PUT("/events/:eventId", eventroutes.UpdateEvent(eventsDb, eventsTableName, dynamoClient))

	// Guests Routes
	e.POST("/events/:eventId/guests", eventroutes.AddGuest(eventsDb, guestsTableName, dynamoClient))
	e.PUT("/events/:eventId/guests/:guestId", eventroutes.ModifyGuest(eventsDb, guestsTableName, dynamoClient))
	e.DELETE("/events/:eventId/guests/:guestId", eventroutes.DeleteGuest(eventsDb, eventsTableName, dynamoClient))

	// FINANCES ROUTES

	// Property Routes
	e.POST("/finances/properties", financesroutes.CreateProperty(financesDb, propertiesTableName, mortgagesTableName, tenantsTableName, incomeTableName, expensesTableName))

	// COOKBOOK ROUTES

	// Recipe Routes
	e.POST("/cookbook/add-recipe", cookbookroutes.CreateRecipe(cookbookDB, recipesTableName, ingredientsTableName, recipeStepsTableName))
	e.GET("/cookbook/recipes", cookbookroutes.GetAllRecipes(cookbookDB, recipesTableName, recipeStepsTableName))

	// Starting the backend server
	logrus.Info("Starting server on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
