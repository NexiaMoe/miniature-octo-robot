/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helloion/database"
	"helloion/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "helloion",
		Short: "CLI for testing REST Client to httpbin.org",
		Long:  `A test for use Rest API Client on golang with CLI.`,

		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	// if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {

}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.helloion.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// Add flag Config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".", "config file (default is $HOME/app.env)")

	fmt.Println(cfgFile)

}

func initConfig() {
	fmt.Println(cfgFile)
	config, err := utils.LoadConfig(cfgFile)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Initialize Database
	database.Init(&config)

}
