package lib

import (
	"log"

	"github.com/mitchellh/go-homedir"
)

func HomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	return home
}
