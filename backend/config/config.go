package config

import (
	"flag"
	"log"

	"github.com/alex-ant/envs"
)

var (
	Environment = flag.String("Environment", "local", "Current Environment")
)

func init() {
	// Parse flags if not parsed already.
	if !flag.Parsed() {
		flag.Parse()
	}

	// Determine and read environment variables.
	flagsErr := envs.GetAllFlags()
	if flagsErr != nil {
		log.Fatalln(flagsErr)
	}
}
