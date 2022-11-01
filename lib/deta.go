package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/drive"
)

// initializes new deta drive instance
func InitDrive(projectKey string, driveName string) *drive.Drive {
	if projectKey == "" {
		fmt.Println("  [!] No project key specified, please set your project key by adding `--project-key` flag in command.")
		os.Exit(1)
	}

	d, err := deta.New(deta.WithProjectKey(projectKey))

	if err != nil {
		log.Fatalf("Failed to initialize new Deta instance. Maybe the project key is wrong? \nError: %v\n", err)
	}

	drive, err := drive.New(d, driveName)
	if err != nil {
		log.Fatalf("Failed to initialize Deta Drive instance.\nError: %v\n", err)
	}

	return drive
}
