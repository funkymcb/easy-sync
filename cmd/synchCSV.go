/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// synchCSVCmd represents the synchCSV command
var synchCSVCmd = &cobra.Command{
	Use:   "synchCSV",
	Short: "synchs a csv file with the specified platform",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("synchCSV called")
	},
}

func init() {
	rootCmd.AddCommand(synchCSVCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// synchCSVCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// synchCSVCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
