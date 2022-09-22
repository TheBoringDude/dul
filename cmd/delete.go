/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/TheBoringDude/dul/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete files from drive.",
	Long:  `Delete files from drive.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No filenames specified.")
			return
		}

		if projectKey == "" {
			projectKey = viper.GetString("projectkey")
		}
		if driveName == "" {
			driveName = viper.GetString("drivename")
		}

		myDrive := lib.InitDrive(projectKey, driveName)

		for _, v := range args {
			if _, err := myDrive.Delete(v); err != nil {
				log.Fatalf("Failed to delete file from drive: %s\n", v)
			}

			fmt.Printf("  [i] Successfully deleted file from drive: %s\n", v)
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
