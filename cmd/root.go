/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/funkymcb/easy-sync/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version   string
	GitCommit string
)

var versionFlag bool
var cfgFile string
var Cfg config.EasySyncConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "easy-sync",
	Short: "cli-tool for synching members of easyverein and wordpress",
	Long: `cli-tool for synching members of easyverein and wordpress

for guidance run: easy-sync --help

for more information visit:	https://github.com/funkymcb/easy-sync`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Printf("easy-sync version: %s commit: %s\n", Version, GitCommit)
			os.Exit(0)
		}
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"configs/easy-sync.yaml",
		"path to config file",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&versionFlag,
		"version",
		"v",
		false,
		"print version of easy-sync",
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
	if !versionFlag {
		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

			// unmarshal config file into struct
			if err := viper.Unmarshal(&Cfg); err != nil {
				log.Fatalf("unable to decode into config struct %v", err)
			}
		} else {
			fmt.Printf("no config file found under path: %s\n", cfgFile)
			fmt.Printf("for more information run:\n\n")
			fmt.Printf("	easy-sync --help\n\n")
			os.Exit(1)
		}
	}
}
