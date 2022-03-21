package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// synchJSONCmd represents the synchFile command
var synchJSONCmd = &cobra.Command{
	Use:   "synchJSON",
	Short: "reads JSON file and synchs it with the specified platform",
	Long: `reads a JSON file and synchs it with the platform specified
eg.: easy-sync synchJSON member-list.json --platform easyverein
run: easy-sync synchJSON --help for more information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("synchFile called")
	},
}

func init() {
	rootCmd.AddCommand(synchJSONCmd)

	synchJSONCmd.Flags().StringVarP(
		&platform,
		"platform",
		"p",
		"",
		"the platform the json file shall be synched with.\nvalid platforms: [easyverein (easy), wordpress (wp)]",
	)
	synchJSONCmd.MarkFlagRequired("platform")
}
