package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

func ListCredentials() {
	// open credentials file
	path := filepath.Join(os.Getenv("HOME"), ".aws/credentials")
	log.Infof("Open %s", path)
	credentials, err := ini.Load(path)
	if err != nil {
		log.WithError(err).Fatal("Error occured with open credentials")
		os.Exit(1)
	}

	// read all sections
	log.Info("Read all sections")
	sections := credentials.Sections()
	for _, s := range sections {
		if s.Name() == "DEFAULT" {
			continue
		}
		fmt.Println(s.Name())
	}
	log.Info("Done")
}
