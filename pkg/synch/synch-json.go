package synch

import (
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/collector"
	"github.com/funkymcb/easy-sync/pkg/comparator"
	"github.com/funkymcb/easy-sync/pkg/models"
	"github.com/funkymcb/easy-sync/pkg/reader"
)

var JSONMembers []models.Member

// JSONtoPlatform reads a json file parses it into the member model
// and synchs it with the specified platform
func JSONtoPlatform(inputFile, platform string) error {
	fmt.Printf("\n################### initialize new synch ####################\n\n")
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

	return nil
}

// JSONtoEasy synchs slice of members with easyverein
func JSONtoEasy() error {
	log.Println("get members of easyverein:")
	// get members from easy-verein for comparison
	easyMembers, err := collector.GetEasyMembers()
	if err != nil {
		return err
	}

	log.Printf("fetched %d members from easyverein", len(easyMembers))

	membersNotInEasy, err := comparator.GetMembersDiff(
		JSONMembers,
		easyMembers,
	)
	if err != nil {
		return fmt.Errorf("could not compare members of json and easyverein %v", err)
	}
	if len(membersNotInEasy) == 0 {
		log.Println("everything is up to date")
		os.Exit(0)
	}
	log.Printf("%d members missing in easyverein", len(membersNotInEasy))

	// TODO check comparator for accuracy
	// TODO implement function to add members to easyverein

	return nil
}

// JSONtoWordpress synchs slice of members with wordpress
func JSONtoWordpress() error {
	return nil
}
