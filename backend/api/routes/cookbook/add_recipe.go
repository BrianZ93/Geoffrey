package routes

import (
	"database/sql"
	"net/http"

	"backend/api"
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
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Generate a new UUID for the recipe, ignoring any ID received from the frontend
		recipe.Id = uuid.New().String()

		// Ensure the recipes table exists
		exists, err := api.TableExists(db, recipesTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, recipesTableName, models.Recipe{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the ingredients table exists
		exists, err = api.TableExists(db, ingredientsTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, ingredientsTableName, models.Ingredient{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the recipe steps table exists
		exists, err = api.TableExists(db, recipeStepsTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, recipeStepsTableName, models.RecipeStep{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Insert the recipe into the Recipes table
		_, err = db.Exec("INSERT INTO "+recipesTableName+" (ID, Name, TotalTimeSeconds, TotalTimePrepSeconds, Country, Note) VALUES (?, ?, ?, ?, ?, ?)",
			recipe.Id, recipe.Name, recipe.TotalTimeSeconds, recipe.TotalTimePrepSeconds, recipe.Country, recipe.Note)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Insert ingredients into the Ingredients table if not empty
		if len(recipe.Ingredients) > 0 {
			for _, ingredient := range recipe.Ingredients {
				ingredient.Id = uuid.New().String() // Generate a new UUID for the ingredient
				if _, err := db.Exec("INSERT INTO "+ingredientsTableName+" (ID, Name) VALUES (?, ?)",
					ingredient.Id, ingredient.Name); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		// Insert recipe steps into the RecipeSteps table if not empty
		if len(recipe.Steps) > 0 {
			for _, step := range recipe.Steps {
				step.Id = uuid.New().String() // Generate a new UUID for the step
				if _, err := db.Exec("INSERT INTO "+recipeStepsTableName+" (ID, RecipeID, Name, ApplianceUsed, CookingAction, TimeSeconds, PrepTimeSeconds, Description, RecipeStage) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
					step.Id, recipe.Id, step.Name, step.ApplianceUsed, step.CookingAction, step.TimeSeconds, step.PrepTimeSeconds, step.Description, step.RecipeStage); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}

				// Insert step ingredients into the Ingredients table
				if len(step.Ingredients) > 0 {
					for _, ingredient := range step.Ingredients {
						ingredient.Id = uuid.New().String() // Generate a new UUID for the ingredient
						if _, err := db.Exec("INSERT INTO "+ingredientsTableName+" (ID, Name) VALUES (?, ?)",
							ingredient.Id, ingredient.Name); err != nil {
							return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
						}
					}
				}
			}
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Recipe created successfully"})
	}
}
