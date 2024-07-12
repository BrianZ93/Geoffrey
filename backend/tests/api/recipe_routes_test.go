package tests

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	routes "backend/api/routes/cookbook"
	models "backend/models/cookbook"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDb() (*sql.DB, error) {
	// Create an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create the tables for testing
	_, err = db.Exec(`
	CREATE TABLE Recipes (
		ID TEXT PRIMARY KEY,
		Name TEXT,
		TotalTimeSeconds INTEGER,
		TotalTimePrepSeconds INTEGER,
		Country TEXT,
		Note TEXT
	);
	CREATE TABLE Ingredients (
		ID TEXT PRIMARY KEY,
		RecipeID TEXT,
		Name TEXT,
		FOREIGN KEY(RecipeID) REFERENCES Recipes(ID)
	);
	CREATE TABLE RecipeSteps (
		ID TEXT PRIMARY KEY,
		RecipeID TEXT,
		Name TEXT,
		ApplianceUsed INTEGER,
		CookingAction INTEGER,
		TimeSeconds INTEGER,
		PrepTimeSeconds INTEGER,
		Description TEXT,
		RecipeStage INTEGER,
		FOREIGN KEY(RecipeID) REFERENCES Recipes(ID)
	);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateRecipe(t *testing.T) {
	// Setup
	e := echo.New()
	db, err := setupTestDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	recipesTableName := "Recipes"
	ingredientsTableName := "Ingredients"
	recipeStepsTableName := "RecipeSteps"

	// Mock request body
	recipe := models.Recipe{
		Name:                 "Test Recipe",
		TotalTimeSeconds:     3600,
		TotalTimePrepSeconds: 1800,
		Country:              "USA",
		Steps: []models.RecipeStep{
			{
				Name:          "Step 1",
				ApplianceUsed: models.Stove,
				CookingAction: models.Fry,
				TimeSeconds:   300,
				Description:   "Fry the ingredients",
				RecipeStage:   models.Cooking,
				Ingredients: []models.Ingredient{
					{Name: "Ingredient 1"},
					{Name: "Ingredient 2"},
				},
			},
		},
		FlavorProfile: models.FlavorProfile{
			Sweet:  3,
			Salty:  2,
			Spicy:  4,
			Sour:   1,
			Savory: 5,
		},
		Note: "Test note",
	}

	jsonRecipe, _ := json.Marshal(recipe)
	req := httptest.NewRequest(http.MethodPost, "/recipes", strings.NewReader(string(jsonRecipe)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Handler
	h := routes.CreateRecipe(db, recipesTableName, ingredientsTableName, recipeStepsTableName)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"message":"Recipe created successfully"}`, strings.TrimSpace(rec.Body.String()))
	}

	// Verify
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Recipes").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM Ingredients").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	err = db.QueryRow("SELECT COUNT(*) FROM RecipeSteps").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}
