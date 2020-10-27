package lambda

import (
	"fmt"
	"github.com/amoriartyCH/accounts-statistics-tool/aws"
	"github.com/amoriartyCH/accounts-statistics-tool/config"
	"github.com/amoriartyCH/accounts-statistics-tool/service"
)

type Lambda struct {
	Service      service.Service
}

type jsonBody struct {}

// New returns a new Lambda using the provided configs
func New(cfg *config.Config) *Lambda {

	return &Lambda{
		Service:      service.NewService(cfg),
	}
}

// Execute handles lambda execution
func (lambda *Lambda) Execute(j *jsonBody) error {

	srCSV := lambda.Service.GetStatisticsReport("CIC report and full accounts")

	err := aws.GenerateEmail(srCSV)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}