package synch

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/funkymcb/easy-sync/pkg/models"
)

var Members []models.Member

// JSONtoPlatform reads a json file parses it into the member model
// and synchs it with the specified platform
func JSONtoPlatform(inputFile, platform string) error {
	if err := parseJSONFile(inputFile); err != nil {
		return err
	}
	// TODO write functions that call apis based on given platform
	return nil
}

// parseJSONFile reads json file and parses into member struct
func parseJSONFile(inputFile string) error {
	jsonData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("could not read json file: %v", err)
	}
	if err := json.Unmarshal(jsonData, &Members); err != nil {
		return fmt.Errorf("could not unmarshal json data to member struct: %v", err)
	}
	return nil
}
