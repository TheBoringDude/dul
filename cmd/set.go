/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configProjectKey string
	configDriveName  string
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set default config keys and values.",
	Long: `Set default config keys and values.

Setting the values of 'project-key' and 'drive' won't affect in here.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configProjectKey == "" && configDriveName == "" {
			return
		}

		viper.SetConfigName("dul")
		viper.SetConfigType("yaml")

		if configProjectKey != "" {
			viper.Set("projectKey", configProjectKey)
		}

		if configDriveName != "" {
			viper.Set("driveName", configDriveName)
		}

		if cfgFile == "" {
			cfgFile = path.Join(homePathDir, ".dul.yaml")
		}

		viper.WriteConfigAs(cfgFile)

		fmt.Println("\n  [i] Successfully updated the default config.")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVar(&configProjectKey, "projectKey", "", "The default project key to use.")
	setCmd.Flags().StringVar(&configDriveName, "driveName", "", "The default drive name to use.")
}
