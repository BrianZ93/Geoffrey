package routes

import (
	"database/sql"
	"net/http"

	"backend/api"
	models "backend/models/finances"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateProperty(
	db *sql.DB,
	propertiesTableName string,
	mortgagesTableName string,
	tenantsTableName string,
	incomeTableName string,
	expensesTableName string,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var property models.Property
		if err := c.Bind(&property); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Generate a new UUID for the property, ignoring any ID received from the frontend
		property.ID = uuid.New().String()

		// Ensure the properties table exists
		exists, err := api.TableExists(db, propertiesTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, propertiesTableName, models.Property{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the mortgages table exists
		exists, err = api.TableExists(db, mortgagesTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, mortgagesTableName, models.Mortgage{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the tenants table exists
		exists, err = api.TableExists(db, tenantsTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, tenantsTableName, models.Tenant{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the income table exists
		exists, err = api.TableExists(db, incomeTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, incomeTableName, models.Income{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Ensure the expenses table exists
		exists, err = api.TableExists(db, expensesTableName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			err = api.CreateTableFromStruct(db, expensesTableName, models.Expense{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}

		// Insert the property into the Properties table
		_, err = db.Exec("INSERT INTO "+propertiesTableName+" (ID, PurchaseDate, PurchasePrice, StreetNumber, StreetName, City, State, Zip, PropertyUse, PropertyType) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			property.ID, property.PurchaseDate, property.PurchasePrice, property.Address.StreetNumber, property.Address.StreetName, property.Address.City, property.Address.State, property.Address.Zip, property.PropertyUse, property.PropertyType)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Insert mortgages into the Mortgages table if not empty
		if len(property.Mortgages) > 0 {
			for _, mortgage := range property.Mortgages {
				mortgage.PropertyId = property.ID // Set the foreign key
				if _, err := db.Exec("INSERT INTO "+mortgagesTableName+" (PropertyID, ID, DebtType, OriginationDate, FirstPaymentDate, Maturity, InterestRate, LienholderPosition, Servicer, Address, Term, Escrow) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
					mortgage.PropertyId, mortgage.ID, mortgage.DebtType, mortgage.OriginationDate, mortgage.FirstPaymentDate, mortgage.Maturity, mortgage.InterestRate, mortgage.LienholderPosition, mortgage.Servicer, mortgage.Address, mortgage.Term, mortgage.Escrow); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		// Insert tenants into the Tenants table if not empty
		if len(property.Tenants) > 0 {
			for _, tenant := range property.Tenants {
				tenant.PropertyId = property.ID // Set the foreign key
				if _, err := db.Exec("INSERT INTO "+tenantsTableName+" (PropertyID, ID, FirstName, LastName, StreetNumber, StreetName, City, State, Zip, LeaseType, LeaseStartDate, LeaseEndDate, Rent, SecurityDeposit, OtherIncome) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
					tenant.PropertyId, tenant.ID, tenant.FirstName, tenant.LastName, tenant.Address.StreetNumber, tenant.Address.StreetName, tenant.Address.City, tenant.Address.State, tenant.Address.Zip, tenant.LeaseType, tenant.LeaseStartDate, tenant.LeaseEndDate, tenant.Rent, tenant.SecurityDeposit, tenant.OtherIncome); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		// Insert incomes into the Income table if not empty
		if len(property.Income) > 0 {
			for _, income := range property.Income {
				if _, err := db.Exec("INSERT INTO "+incomeTableName+" (ID, PropertyID, Description, Memo, Amount, Date) VALUES (?, ?, ?, ?, ?, ?)",
					income.ID, property.ID, income.Description, income.Memo, income.Amount, income.Date); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		// Insert expenses into the Expenses table if not empty
		if len(property.Expenses) > 0 {
			for _, expense := range property.Expenses {
				if _, err := db.Exec("INSERT INTO "+expensesTableName+" (ID, PropertyID, Description, Memo, Amount, Date) VALUES (?, ?, ?, ?, ?, ?)",
					expense.ID, property.ID, expense.Description, expense.Memo, expense.Amount, expense.Date); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
			}
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Property created successfully"})
	}
}
