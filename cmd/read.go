/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/funkymcb/easy-sync/pkg/reader"
	"github.com/spf13/cobra"
)

var outputFile string
var csvDelimiter string

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "reads a member-list and parses it to a JSON file",
	Long: `read reads a member-list and parses it to a API readable JSON file

valid file formats are .csv (more to follow)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf(`file to be read needs to be specified
eg.: easy-sync read member-list.csv
run: easy-sync read --help for more information
`)
			os.Exit(1)
		}
		if len(csvDelimiter) > 1 {
			fmt.Printf(`delimiter of csv needs to be a single character
eg.: easy-sync read member-list.csv -d ';'
run: easy-sync read --help for more information
`)
			os.Exit(1)
		}

		reader.ReadFile(args[0], outputFile, csvDelimiter)
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
	readCmd.Flags().StringVarP(
		&csvDelimiter,
		"delimiter",
		"d",
		",",
		"delimiter used to seperate csv fields. Needs to be a single character",
	)
}
