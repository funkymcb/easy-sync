package reader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/writer"
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
	if err := writer.WriteJSONFile(members, outFile); err != nil {
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
