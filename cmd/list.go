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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
