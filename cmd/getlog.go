/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helloion/filter"
	"helloion/function"
	"helloion/utils"

	"github.com/spf13/cobra"
)

// getlogCmd represents the getlog command
var getlogCmd = &cobra.Command{
	Use:   "getlog",
	Short: "Get log from database",
	Long:  `Get log from database`,
	Run: func(cmd *cobra.Command, args []string) {
		f := new(filter.CLIFilter)
		_, ip, err := function.GetIPFromDB(f)

		utils.CheckErr(err)

		fmt.Println(ip)

	},
}

func init() {
	rootCmd.AddCommand(getlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
