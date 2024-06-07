package api

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func TableExists(db *sql.DB, tableName string) (bool, error) {
	query := fmt.Sprintf(`SELECT name FROM sqlite_master WHERE type='table' AND name='%s';`, tableName)
	var name string
	err := db.QueryRow(query).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result += sep + s
	}
	return result
}

func mapGoTypeToSQLType(t reflect.Type) (string, error) {
	switch t.Kind() {
	case reflect.String:
		return "TEXT", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER", nil
	case reflect.Bool:
		return "BOOLEAN", nil
	case reflect.Float32, reflect.Float64:
		return "REAL", nil
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return "DATETIME", nil
		}
	}
	return "", fmt.Errorf("unsupported Go type: %s", t.Kind())
}

func generateCreateTableSQL(tableName string, db *sql.DB, t interface{}) (string, string, error) {
	typ := reflect.TypeOf(t)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return "", "", fmt.Errorf("expected a struct type, got %s", typ.Kind())
	}

	var columns []string
	foreignKeyColumns := make(map[string]struct{})

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldType, err := mapGoTypeToSQLType(field.Type)
		if err != nil {
			// Handle slice type
			if field.Type.Kind() == reflect.Slice {
				sliceElem := field.Type.Elem()
				if sliceElem.Kind() == reflect.Struct {
					sliceTableName := tableName + "_" + fieldName
					sliceStruct := reflect.New(sliceElem).Interface()
					// Create table for the slice
					err := CreateTableFromStruct(db, sliceTableName, sliceStruct)
					if err != nil {
						return "", "", err
					}
					// Add foreign key column
					foreignKeyColumnName := tableName + "_id"
					if _, exists := foreignKeyColumns[foreignKeyColumnName]; !exists {
						columns = append(columns, fmt.Sprintf("%s TEXT", foreignKeyColumnName))
						foreignKeyColumns[foreignKeyColumnName] = struct{}{}
					}
				}
				continue
			}
		}

		columns = append(columns, fmt.Sprintf("%s %s", fieldName, fieldType))
	}

	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, join(columns, ", "))
	return tableName, createTableSQL, nil
}

func CreateTableFromStruct(db *sql.DB, tableName string, t interface{}) error {
	_, createTableSQL, err := generateCreateTableSQL(tableName, db, t)
	if err != nil {
		return err
	}

	logrus.Infof("Creating table %s with SQL: %s\n", tableName, createTableSQL)
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", tableName, err)
	}
	return nil
}
