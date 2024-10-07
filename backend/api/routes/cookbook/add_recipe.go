package routes

import (
	"database/sql"
	"log"
	"net/http"

	"backend/api"
	mainModels "backend/models"
	models "backend/models/cookbook"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateRecipe(
	db *sql.DB,
	recipesTableName string,
	ingredientsTableName string,
	recipeStepsTableName string,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var recipe models.Recipe
		if err := c.Bind(&recipe); err != nil {
			log.Printf("Error binding recipe: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Log the incoming recipe data for debugging
		log.Printf("Received recipe: %+v", recipe)

		// Convert the country value to a string
		countryStr, err := mainModels.Country(recipe.Country).String()
		if err != nil {
			log.Printf("Error converting country: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid country value"})
		}

		// Generate a new UUID for the recipe, ignoring any ID received from the frontend
		recipe.ID = uuid.New().String()

		// Generate new UUIDs for each step and ignore any IDs received from the frontend
		for i := range recipe.Steps {
			recipe.Steps[i].ID = uuid.New().String()

			// Generate new UUIDs for each ingredient in the step
			for j := range recipe.Steps[i].Ingredients {
				recipe.Steps[i].Ingredients[j].ID = uuid.New().String()
			}
		}

		// Ensure the recipes table exists
		exists, err := api.TableExists(db, recipesTableName)
		if err != nil {
			log.Printf("Error checking recipes table existence: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, recipesTableName, models.Recipe{})
			if err != nil {
				log.Printf("Error creating recipes table: %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the ingredients table exists
		exists, err = api.TableExists(db, ingredientsTableName)
		if err != nil {
			log.Printf("Error checking ingredients table existence: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, ingredientsTableName, models.Ingredient{})
			if err != nil {
				log.Printf("Error creating ingredients table: %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the recipe steps table exists
		exists, err = api.TableExists(db, recipeStepsTableName)
		if err != nil {
			log.Printf("Error checking recipe steps table existence: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, recipeStepsTableName, models.RecipeStep{})
			if err != nil {
				log.Printf("Error creating recipe steps table: %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Insert the recipe into the Recipes table
		_, err = db.Exec("INSERT INTO "+recipesTableName+" (ID, Name, TotalTimeSeconds, TotalTimePrepSeconds, Country, Note) VALUES (?, ?, ?, ?, ?, ?)",
			recipe.ID, recipe.Name, recipe.TotalTimeSeconds, recipe.TotalTimePrepSeconds, countryStr, recipe.Note) // Use countryStr here
		if err != nil {
			log.Printf("Error inserting recipe: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Insert recipe steps into the RecipeSteps table if not empty
		if len(recipe.Steps) > 0 {
			for _, step := range recipe.Steps {
				if _, err := db.Exec("INSERT INTO "+recipeStepsTableName+" (ID, RecipeID, Name, ApplianceUsed, CookingAction, TimeSeconds, PrepTimeSeconds, Description, RecipeStage) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
					step.ID, recipe.ID, step.Name, step.ApplianceUsed, step.CookingAction, step.TimeSeconds, step.PrepTimeSeconds, step.Description, step.RecipeStage); err != nil {
					log.Printf("Error inserting recipe step: %v", err)
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}

				// Insert step ingredients into the Ingredients table
				if len(step.Ingredients) > 0 {
					for _, ingredient := range step.Ingredients {
						if _, err := db.Exec("INSERT INTO "+ingredientsTableName+" (ID, Name) VALUES (?, ?)",
							ingredient.ID, ingredient.Name); err != nil {
							log.Printf("Error inserting step ingredient: %v", err)
							return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
						}
					}
				}
			}
		}

		log.Printf("Recipe created successfully: %+v", recipe)
		return c.JSON(http.StatusOK, map[string]string{"message": "Recipe created successfully"})
	}
}
