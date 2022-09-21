package lib

import (
	"log"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/drive"
)

// initializes new deta drive instance
func InitDrive(projectKey string, driveName string) *drive.Drive {
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
