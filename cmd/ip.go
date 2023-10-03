/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helloion/filter"
	"helloion/function"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
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

		// Check if periodically flag is true, if true then it will run in daemon mode

		if cmd.Flag("periodically").Value.String() == "true" {
			s := gocron.NewScheduler(time.UTC)
			second, err := strconv.Atoi(cmd.Flag("Second").Value.String())

			if err != nil {
				fmt.Println(err)
			}

			// Schedule a job for every 5 seconds.
			s.Every(second).Seconds().Do(function.GetIP, f)

			s.StartAsync()

			// We need to add context background here to prevent exit from function until we press ctrl+c or program exit

			c := make(chan os.Signal, 1)

			signal.Notify(c, os.Interrupt)

			is_shutdown := false
			is_running := false

		OUTTER:
			for {
				if is_shutdown {
					break
				}

				select {
				case s, ok := <-c:
					if ok {
						fmt.Println("Got signal:", s)
						is_shutdown = true

						c = nil

						continue OUTTER
					}
				default:
					if !is_running {
						fmt.Println("IP Daemon Running")
						is_running = true
					}

				}
			}

			fmt.Println("IP Daemon Stopped")

			os.Exit(0)

		}

		statusCode, res, err := function.GetIP(f)
		if err != nil || statusCode != 200 {
			fmt.Println(err)
		}

		fmt.Println(res)

	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	f := new(filter.CLIFilter)

	ipCmd.Flags().BoolVarP(&f.Json, "json", "j", false, "Print JSON format")

	ipCmd.Flags().BoolVarP(&f.Periodically, "periodically", "p", false, "Get IP address periodically")

	ipCmd.Flags().IntVarP(&f.Second, "Second", "s", 5, "Get IP address periodically")

	// rootCmd.Flags().BoolVarP(&f.Json, "json", "j", false, "Print JSON format")
	// rootCmd.Flag("json")
	// rootCmd.Flag("periodically")
	// rootCmd.Flag("Second")

	// rootCmd.HasLocalFlags()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
