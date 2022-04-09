package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/reader"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "readCSV",
	Short: "reads a member-list csv and parses it to a JSON file",
	Long: `readCSV reads a csv member-list and parses it to a JSON file

example:
	easy-sync readCSV member-list.csv --delimiter ';' --output result.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf(`file to be read needs to be specified
eg.: easy-sync readCSV member-list.csv -o result.json
run: easy-sync readCSV --help for more information
`)
			os.Exit(1)
		}

		if err := reader.ReadCSV(args[0], outputFile, csvDelimiter); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	readCmd.Flags().StringVarP(
		&outputFile,
		"output",
		"o",
		"out/member-list.json",
		"name of the output file",
	)
	readCmd.MarkFlagRequired("output")

	readCmd.Flags().StringVarP(
		&csvDelimiter,
		"delimiter",
		"d",
		",",
		"delimiter used to seperate csv fields. Needs to be a single character",
	)
}
