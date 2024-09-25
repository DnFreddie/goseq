/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package git

import (
	"DnFreddie/goseq/lib"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var iname bool
var regex bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var re lib.GrepFlag
		var insencitive lib.GrepFlag

		period := lib.Period{
			Range:  lib.All,
			Amount: 0,
		}


		projects, err := NewDRetriver().GetNotes(period)
		if err != nil {
			fmt.Printf("Faield to retrive projects: %v", err)
			os.Exit(1)
		}
		if iname {
			insencitive = lib.ToLower

		}
		if regex {
			re = lib.Regex
		}

		if err := lib.Search(projects, strings.Join(args, " "), re|insencitive); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	GitCmd.AddCommand(searchCmd)

	searchCmd.Flags().BoolVarP(&iname, "iname", "i", false, "Case Insensitive Search")
	searchCmd.Flags().BoolVarP(&regex, "regex", "E", false, "Accepts Posix Regex Search")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
