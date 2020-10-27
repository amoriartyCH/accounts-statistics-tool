package main

import (
	"fmt"
	"github.com/amoriartyCH/accounts-statistics-tool/config"
	"github.com/amoriartyCH/accounts-statistics-tool/lambda"
	. "github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
)

// Main entry point for application, will create config from ENV variables and start the lambda.
func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Error(fmt.Sprintf("Error when establishing config: %s", err))
		os.Exit(1)
	}

	config.SetLogLevel(cfg)

	l := lambda.New(cfg)
	Start(l.Execute)
}
