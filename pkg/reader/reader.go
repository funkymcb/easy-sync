package reader

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/models"
	"golang.org/x/text/encoding/charmap"
)

// ReadFile reads a csv member-list and parses it into the user struct
// which can then be used for synching between the web services
func ReadFile(inFile, outFile, delimiter string) error {
	// check if delimiter is single char
	if len(delimiter) > 1 {
		return fmt.Errorf("delimiter of csv needs to be a single character")
	}

	log.Println("read member-list from file:", inFile)
	records, err := readData(inFile, delimiter)
	if err != nil {
		return fmt.Errorf(
			"could not read csv file. make sure to specify the correct csv delimiter using the --delimiter flag. error: %v",
			err,
		)
	}

	members := parseMembers(records)

	log.Println("writing output to file:", outFile)
	if err := writeJSONFile(members, outFile); err != nil {
		return fmt.Errorf("could not write members to json file %v", err)
	}
	return nil
}

// readData parses the actual csv content to a 2d sting slice
// outer slice represents a single line,
// inner slice represents each value of the line
func readData(fileName, delimiter string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// provide io.reader to csv.reader to decode in utf-8
	r := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(f))
	delimiterSlice := []rune(delimiter)
	r.Comma = delimiterSlice[0]

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

// parseMembers reads the csv records and parses them to the members struct
func parseMembers(records [][]string) []models.Member {
	var members []models.Member
	// iterate over each record and parse it into member struct
	for _, record := range records {
		username := models.GenerateUsername(record[3], record[4])

		member := models.Member{
			Salution:    record[2],
			FirstName:   record[3],
			FamilyName:  record[4],
			Company:     record[5],
			Street:      record[8],
			PostalCode:  record[9],
			Town:        record[10],
			Telefone:    record[11],
			IBAN:        record[13],
			BIC:         record[14],
			DateOfBirth: record[15],
			Mobile:      record[18],
			EMail:       record[19],
			Title:       record[31],
			CustomFields: models.CustomFields{
				Gender:     record[21],
				BankName:   record[25],
				Username:   username,
				Fax:        record[17],
				Department: record[28],
			},
			// Ausw:			  record[0],
			// MemberID:          record[1],
			// ShippingMethod:    record[6],
			// AddressSupplement: record[7],
			// Country:           record[12],
			// BusinessTelefone:  record[16],
			// Nationality:       record[20],
			// MartialStatus:     record[22],
			// Job:               record[23],
			// Status:            record[24],
			// Entry:             record[26],
			// Termination:       record[27],
			// Department:        record[28],
			// Function:          record[29],
			// MandateNumber:     record[30],
		}
		members = append(members, member)
	}
	return members
}

// writeJSONFile writes a slice of members to a json file
func writeJSONFile(members []models.Member, outFile string) error {
	file, err := json.MarshalIndent(members, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(outFile, file, 0644); err != nil {
		return err
	}

	return nil
}
