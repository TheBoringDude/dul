package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/TheBoringDude/dul/lib"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files from drive.",
	Long:  `List all the files that are stored in drive in a table format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectKey == "" {
			projectKey = viper.GetString("projectkey")
		}
		if driveName == "" {
			driveName = viper.GetString("drivename")
		}

		myDrive := lib.InitDrive(projectKey, driveName)

		list, err := myDrive.List(1000, "", "")
		if err != nil {
			log.Fatalf("Failed to list the files from Drive. (err: %v)\n", err)
		}

		data := [][]string{}

		for i, v := range list.Names {
			data = append(data, []string{strconv.Itoa(i + 1), v})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "Files"})

		for _, v := range data {
			table.Append(v)
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
