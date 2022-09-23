package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TheBoringDude/dul/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getFile string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get --file [file]",
	Short: "Download file from drive.",
	Long:  `Downloads a file from deta drive.`,
	Run: func(cmd *cobra.Command, args []string) {
		if getFile == "" {
			fmt.Println("No file name specified to download.")
			return
		}

		if projectKey == "" {
			projectKey = viper.GetString("projectkey")
		}
		if driveName == "" {
			driveName = viper.GetString("drivename")
		}

		myDrive := lib.InitDrive(projectKey, driveName)
		f, err := myDrive.Get(getFile)
		if err != nil {
			log.Fatalf("Error getting file. (err: %v)\n", err)
		}

		defer f.Close()

		outFile, err := os.Create(getFile)
		if err != nil {
			log.Fatalf("Failed creating new file: %s (err: %v)\n", getFile, err)
		}

		defer outFile.Close()

		_, err = io.Copy(outFile, f)
		if err != nil {
			log.Fatalf("Failed copying contents to file. (err: %v)\n", err)
		}

		fmt.Printf("\n  [i] Successfully downloaded file: %s\n", getFile)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&getFile, "file", "f", "", "Name of file to get or download.")
}
