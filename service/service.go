package service

import (
	"fmt"
	"github.com/amoriartyCH/accounts-statistics-tool/models"
	"os"
	"strconv"
	"time"

	"github.com/amoriartyCH/accounts-statistics-tool/config"
	"github.com/amoriartyCH/accounts-statistics-tool/db"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	GetStatisticsReport(dataDescription string) *models.StatisticsReport
}

type Impl struct {
	transactionClient db.TransactionClient
}

func NewService(cfg *config.Config) Service {

	return &Impl{
		transactionClient: db.NewTransactionDatabaseClient(cfg),
	}
}

func (s *Impl) GetStatisticsReport(dataDescription string) *models.StatisticsReport {

	// Retrieve all transactions which are closed and match our dataDescription string.
	transactions, err := s.transactionClient.GetAccountsTransactions(dataDescription)
	if err != nil {
		log.Error(fmt.Sprintf("Error when retrieving transactions: %s", err))
		os.Exit(1)
	}

	// Store our stats inside a StatisticsReport struct.
	sr := sortTransactionsPerMonth(transactions)

	// Print the struct cleanly for the user to view.
	// (in future I suggest we look to output to CSV or some sort of document store).
	// Payments-reconciler example of storing to CSV.
	printStatisticsReport(sr)

	return sr
}

/*
	This function will take a slice of transactions and sort them:
		-Firstly they will be grouped by a status of "accepted" or "rejected"
		-Secondly they will be grouped by year of filing
		-Finally they will be grouped by month of filing
 */
func sortTransactionsPerMonth(transactions *[]models.Transaction) *models.StatisticsReport {

	// Initialise our statisticsReport, this will be used to hold all stats needed.
	sr := models.NewStatisticsReport()

	// Instantly set our ClosedTransactions counter to the length of the slice passed in.
	sr.FirstYearAcceptedMonthlyFilings["ClosedTransactions"] = len(*transactions)

	// Define a time of 1 year ago using today's date.
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	// Loop over the found transactions and sort them per year and month.
	for _, t := range *transactions {

		// Retrieve the status of the transactions filing from each transaction.
		accepted := t.Data.Filings[t.ID + "-1"].Status == "accepted"
		rejected := t.Data.Filings[t.ID + "-1"].Status == "rejected"

		// If our status was accepted then we are interested in logging which year/month it happened.
		if accepted {

			// If our transaction was closed within a year from today, then its added to our FirstYearFilings map.
			if t.Data.ClosedAt.After(oneYearAgo) {
				sr.FirstYearAcceptedMonthlyFilings[strconv.Itoa(int(t.Data.ClosedAt.Month()))]++
			}

			// Increase our accepted transactions by 1 each loop if we reach this point.
			sr.FirstYearAcceptedMonthlyFilings["AcceptedTransactions"]++

		} else if rejected {
			//Alternatively if the filing we rejected then we increase our rejected filings by 1.
			sr.FirstYearAcceptedMonthlyFilings["RejectedTransactions"]++
		}
	}

	// Return our fully populated statistics report.
	return sr
}

func printStatisticsReport(sr *models.StatisticsReport) {

	// Filings for the first year, printed per month.
	log.Info(fmt.Sprintf("--- Statistics Report Tool ---"))
	log.Info(fmt.Sprintf("--- Within 12 months Filings (Per Month) ---"))

	for i, f := range sr.FirstYearAcceptedMonthlyFilings {
		month, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		log.Info(fmt.Sprintf("%v Filings: %d", time.Month(month).String(), f))
	}

	log.Info(fmt.Sprintf("--- Total: %d ---", sr.FirstYearAcceptedMonthlyFilings["ClosedTransactions"]))
	log.Info(fmt.Sprintf("-------------------"))


	// Total filings printed per status.
	log.Info(fmt.Sprintf("--- Filings grouped by status ---"))
	log.Info(fmt.Sprintf("Closed transactions: %d", sr.FirstYearAcceptedMonthlyFilings["ClosedTransactions"]))
	log.Info(fmt.Sprintf("Accepted transactions: %d", sr.FirstYearAcceptedMonthlyFilings["AcceptedTransactions"]))
	log.Info(fmt.Sprintf("Rejected transactions: %d", sr.FirstYearAcceptedMonthlyFilings["RejectedTransactions"]))
	log.Info(fmt.Sprintf("-------------------"))
}