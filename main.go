package main

import (
	"fmt"
	"os"

	"github.com/BSaunders95/accounts-statistics-tool/config"
	"github.com/BSaunders95/accounts-statistics-tool/service"
	log "github.com/sirupsen/logrus"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Error(fmt.Sprintf("Error when establishing config: %s", err))
		os.Exit(1)
	}

	dataDescription := "CIC report and full accounts"
	if len(os.Args) > 1 {
		dataDescription = os.Args[1]
	}

	svc := service.NewService(cfg)

	svc.GetNumberOfCICReports(dataDescription)
}
