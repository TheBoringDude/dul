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
	Use:   "delete [file1] [file2] [file3] ...",
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
}
