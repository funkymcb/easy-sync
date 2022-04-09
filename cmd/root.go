/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version   string
	GitCommit string
)

var (
	cfgFile string
	cfg     models.EasySyncConfig
)

var (
	outputFile   string
	inputFile    string
	csvDelimiter string
	platform     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "easy-sync",
	Short: "cli-tool for synching members of easyverein and wordpress",
	Long: `cli-tool for synching members of easyverein and wordpress

for guidance run: easy-sync --help

for more information visit:	https://github.com/funkymcb/easy-sync`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Help())
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Version = fmt.Sprintf(
		"easy-sync version: %s commit: %s\n",
		Version,
		GitCommit,
	)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"configs/easy-sync.yaml",
		"path to config file",
	)
	rootCmd.PersistentFlags().BoolVar(
		&models.VerboseFlag,
		"verbose",
		false,
		"add more verbose output to command execution",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		configDir, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".easy-sync" (without extension).
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".easy-sync")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// unmarshal config file into struct
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("unable to decode into config struct %v", err)
		}
		models.SetConfig(cfg)
	} else {
		fmt.Printf("no config file found under path: %s\n", cfgFile)
		fmt.Printf("for more information run:\n\n")
		fmt.Printf("	easy-sync --help\n\n")
		os.Exit(1)
	}
}
