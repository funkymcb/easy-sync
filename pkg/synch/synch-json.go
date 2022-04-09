package synch

import (
	"log"

	"github.com/funkymcb/easy-sync/pkg/collector"
	"github.com/funkymcb/easy-sync/pkg/models"
	"github.com/funkymcb/easy-sync/pkg/reader"
)

var JSONMembers []models.Member

// JSONtoPlatform reads a json file parses it into the member model
// and synchs it with the specified platform
func JSONtoPlatform(inputFile, platform string) error {
	if err := reader.ParseJSONtoMembers(inputFile, &JSONMembers); err != nil {
		return err
	}

	if platform == "easyverein" || platform == "easy" {
		if err := JSONtoEasy(); err != nil {
			return err
		}
	}
	if platform == "wordpress" || platform == "wp" {
		if err := JSONtoWordpress(); err != nil {
			return err
		}
	}
	// 2. parse them into member struct
	// 3. compare jsonMembers with easyMembers
	// 4. POST diff to easyVerein

	return nil
}

// JSONtoEasy synchs slice of members with easyverein
func JSONtoEasy() error {
	log.Println("get members of easyverein")
	// get members from easy-verein for comparison
	easyMembers, err := collector.GetEasyMembers()
	if err != nil {
		return err
	}

	// DEBUG output
	log.Printf("fetched %d members from easyverein", len(easyMembers))

	return nil
}

// JSONtoWordpress synchs slice of members with wordpress
func JSONtoWordpress() error {
	return nil
}
