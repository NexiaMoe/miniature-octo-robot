/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helloion/filter"
	"helloion/function"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Print IP address from httpbin.org",
	Long:  `It will print IP address from httpbin.org.`,

	Run: func(cmd *cobra.Command, args []string) {
		f := new(filter.CLIFilter)

		if cmd.Flag("json").Changed {
			f.Json = true
		}

		statusCode, res, err := function.GetIP(f)
		if err != nil || statusCode != 200 {
			fmt.Println(err)
		}

		fmt.Println(res)

	},
}

func init() {
	f := new(filter.CLIFilter)

	rootCmd.PersistentFlags().BoolVarP(&f.Json, "json", "j", false, "Print JSON format")

	// rootCmd.Flags().BoolVarP(&f.Json, "json", "j", false, "Print JSON format")
	rootCmd.Flag("json")

	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
