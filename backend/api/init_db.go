package api

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	_ "path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func InitDb(filepath string) (*sql.DB, error) {
	// Get the directory from the filepath
	dir := path.Dir(filepath)

	// Check if the directory exists, create it if it doesn't
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logrus.Info("Directory does not exist, creating a new one...")

		// Create the directory and all necessary parents
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logrus.Errorf("Failed to create directory: %v", err)
			return nil, fmt.Errorf("failed to create directory: %w", err)
		} else {
			logrus.Info("Directory was successfully created")
		}
	}

	// Check if the database file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		logrus.Info("Database file does not exist, creating a new one...")

		// Create the database file
		dbFile, err := os.Create(filepath)
		if err != nil {
			logrus.Errorf("Failed to create database file: %v", err)
			return nil, fmt.Errorf("failed to create database file: %w", err)
		} else {
			logrus.Info("Database was successfully created")
		}
		dbFile.Close()
	}

	// Open the database
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		logrus.Errorf("failed to open database: %v", err)
		return nil, fmt.Errorf("failed to open database: %w", err)
	} else {
		logrus.Info("Database connected")
	}

	return db, nil
}
