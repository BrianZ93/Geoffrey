package routes

import (
	models "backend/models/events"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Guests struct {
	Guests []models.Guest `yaml:"guests"`
}

func LoadGuests(filename string) ([]models.Guest, error) {
	// Read the YAML file
	logrus.Infof("Reading guests from %s", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		logrus.Errorf("Could not read %s file: %v", filename, err)
		return nil, err
	}

	// Unmarshal the YAML data into the Guest struct
	var guests Guests
	err = yaml.Unmarshal(data, &guests)
	if err != nil {
		logrus.Errorf("Could not unmarshal YAML data: %v", err)
		return nil, err
	}

	return guests.Guests, nil
}
