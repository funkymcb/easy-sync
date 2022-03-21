package reader

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/models"
)

var Members []models.Member

// ReadFile reads a csv member-list and parses it into the user struct
// which can then be used for synching between the web services
func ReadFile(inFile, outFile, delimiter string) {
	fmt.Println("read member-list from file:", inFile)
	records, err := readData(inFile, delimiter)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over each record and parse it into member struct
	for _, record := range records {
		member := models.Member{
			Ausw:              record[0],
			MemberID:          record[1],
			Salution:          record[2],
			FirstName:         record[3],
			SecondName:        record[4],
			Company:           record[5],
			ShippingMethod:    record[6],
			AddressSupplement: record[7],
			Street:            record[8],
			PostalCode:        record[9],
			Town:              record[10],
			Telefone:          record[11],
			Country:           record[12],
			IBAN:              record[13],
			BIC:               record[14],
			BirthDate:         record[15],
			BusinessTelefone:  record[16],
			Fax:               record[17],
			Mobile:            record[18],
			EMail:             record[19],
			Nationality:       record[20],
			Gender:            record[21],
			MartialStatus:     record[22],
			Job:               record[23],
			Status:            record[24],
			BankDesignation:   record[25],
			Entry:             record[26],
			Termination:       record[27],
			Department:        record[28],
			Function:          record[29],
			MandateNumber:     record[30],
			Title:             record[31],
		}
		Members = append(Members, member)
	}

	fmt.Println("writing output to file:", outFile)
	if err := writeJSONFile(Members, outFile); err != nil {
		log.Fatal("could not write members to json file", err)
	}
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

	r := csv.NewReader(f)
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

// writeJSONFile writes a slice of members to a json file
func writeJSONFile(members []models.Member, outFile string) error {
	file, err := json.MarshalIndent(members, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(outFile, file, 0644); err != nil {
		return err
	}

	return nil
}
