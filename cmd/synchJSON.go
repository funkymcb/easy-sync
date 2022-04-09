package cmd

import (
	"log"

	"github.com/funkymcb/easy-sync/pkg/synch"
	"github.com/spf13/cobra"
)

// synchJSONCmd represents the synchFile command
var synchJSONCmd = &cobra.Command{
	Use:   "synchJSON",
	Short: "reads JSON file and synchs it with the specified platform",
	Long: `reads a JSON file and synchs it with the platform specified
eg.:
	easy-sync synchJSON --input member-list.json --platform easyverein

if you just want to show what would be synched use the dryrun flag:
	easy-sync synchJSON --input member-list.json --platform easyverein --dryrun
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := synch.JSONtoPlatform(inputFile, platform); err != nil {
			log.Fatal(err)
		}
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

	synchJSONCmd.Flags().StringVarP(
		&inputFile,
		"input",
		"i",
		"",
		"the path to the JSON file you want to synch",
	)
	synchJSONCmd.MarkFlagRequired("input")

	synchJSONCmd.Flags().BoolVar(
		&dryrunFlag,
		"dryrun",
		false,
		"shows output of command without performing actions",
	)
}
