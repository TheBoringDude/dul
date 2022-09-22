/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/TheBoringDude/dul/lib"
	simplefiletest "github.com/TheBoringDude/simple-filetest"
	"github.com/deta/deta-go/service/drive"
	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	projectKey string
	driveName  string
	cfgFile    string
	files      []string
)

var homePathDir string = lib.HomeDir()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dul",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if projectKey == "" {
			projectKey = viper.GetString("projectkey")
		}
		if driveName == "" {
			driveName = viper.GetString("drivename")
		}

		if len(files) == 0 {
			fmt.Println("No files specified.")
			return
		}

		myDrive := lib.InitDrive(projectKey, driveName)

		for _, v := range files {
			if !simplefiletest.FileExists(v) {
				log.Fatalf("File: %s does not exist.\n", v)
			}

			file, err := os.Open(v)
			if err != nil {
				log.Fatalf("Error reading file: %s\nerror: %v", v, err)
			}

			mtype, err := mimetype.DetectFile(v)
			if err != nil {
				log.Fatalf("Failed to detect MIME Type of %s (err: %v)\n", v, err)
			}

			name := filepath.Base(v)

			put, err := myDrive.Put(&drive.PutInput{
				Name:        name,
				Body:        bufio.NewReader(file),
				ContentType: mtype.String(),
			})
			if err != nil {
				log.Fatalf("Failed to put file in Drive: %s\nerror: %v", v, err)
			}

			fmt.Printf("  [i] Succesfully put file: %s\n", put)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&projectKey, "project-key", "", "Your Deta project key. (default: config.projectkey)")
	rootCmd.PersistentFlags().StringVarP(&driveName, "drive", "d", "", "Drive name to store the files. (default: config.drivename)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: $HOME/.dul.yaml)")

	rootCmd.Flags().StringSliceVarP(&files, "files", "f", []string{}, "Array of files you want to upload.")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(homePathDir)
		viper.SetConfigName(".dul")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found, create the file

			if cfgFile == "" {
				cfgFile = path.Join(homePathDir, ".dul.yaml")
			}

			if err = viper.WriteConfigAs(cfgFile); err != nil {
				log.Fatalln(err)
			}
		} else {
			log.Fatalf("Can't read config: %v", err)
		}
	}

}
