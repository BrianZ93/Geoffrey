package routes

import (
	"database/sql"
	"net/http"

	models "backend/models/cookbook"

	"github.com/labstack/echo/v4"
)

func GetAllRecipes(
	db *sql.DB,
	recipesTableName string,
	recipeStepsTableName string,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var recipes []models.Recipe

		// Query for all recipes
		rows, err := db.Query("SELECT ID, Name, TotalTimeSeconds, TotalTimePrepSeconds, Country, Note FROM " + recipesTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer rows.Close()

		for rows.Next() {
			var recipe models.Recipe
			if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.TotalTimeSeconds, &recipe.TotalTimePrepSeconds, &recipe.Country, &recipe.Note); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			// Query for recipe steps for each recipe
			stepRows, err := db.Query("SELECT ID, Name, ApplianceUsed, CookingAction, TimeSeconds, PrepTimeSeconds, Description, RecipeStage FROM "+recipeStepsTableName+" WHERE RecipeID = ?", recipe.ID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer stepRows.Close()

			var steps []models.RecipeStep
			for stepRows.Next() {
				var step models.RecipeStep
				if err := stepRows.Scan(&step.ID, &step.Name, &step.ApplianceUsed, &step.CookingAction, &step.TimeSeconds, &step.PrepTimeSeconds, &step.Description, &step.RecipeStage); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}

				// Query for ingredients for each recipe step
				ingredientRows, err := db.Query("SELECT ID, Name FROM "+recipeStepsTableName+"_Ingredients WHERE RecipeStepID = ?", step.ID)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
				defer ingredientRows.Close()

				var stepIngredients []models.Ingredient
				for ingredientRows.Next() {
					var ingredient models.Ingredient
					if err := ingredientRows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
						return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
					}
					stepIngredients = append(stepIngredients, ingredient)
				}
				step.Ingredients = stepIngredients

				steps = append(steps, step)
			}
			recipe.Steps = steps

			recipes = append(recipes, recipe)
		}

		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, recipes)
	}
}
