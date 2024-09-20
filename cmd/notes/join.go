/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package notes

import (
	"DnFreddie/goseq/lib"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// joinCmd represents the join command
var period string
var JoinCmd = &cobra.Command{
	Use:   "join",
	Short: "Joins notes in one ",
	Long:  `Join notes any changes to the notes will be applaied to the notes (by defult from one week last 7 notes) `,
	Run: func(cmd *cobra.Command, args []string) {

		home := viper.GetString("HOME")
		p := parseDateRange(period)

		dirs, _ := os.ReadDir(path.Join(home, lib.AGENDA))
		err := lib.JoinNotes(&dirs, p)
		if err != nil {
			log.Fatal(err)
		}
		lib.ScanEverything()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	JoinCmd.Flags().StringVarP(&period, "range", "r", "week", "Date range (day, week, month, year, all)")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func parseDateRange(input string) lib.DateRange {
	switch strings.ToLower(input) {
	case "day":
		return lib.Day
	case "week":
		return lib.Week
	case "month":
		return lib.Month
	case "year":
		return lib.Year
	case "all":
		return lib.All
	default:
		fmt.Printf("Invalid date range: %s. Defaulting to 'all'.\n", input)
		return lib.Week
	}
}
