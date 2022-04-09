package writer

import (
	"encoding/json"
	"os"

	"github.com/funkymcb/easy-sync/pkg/models"
)

// WriteJSONFile writes a slice of members to a json file
func WriteJSONFile(members []models.Member, outFile string) error {
	file, err := json.MarshalIndent(members, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(outFile, file, 0644); err != nil {
		return err
	}

	return nil
}
