package models

import . "time"

type StatisticsReport struct {
	ClosedTransactions   int
	AcceptedTransactions int
	RejectedTransactions int

	FirstYearAcceptedMonthlyFilings  map[Month]int
	SecondYearAcceptedMonthlyFilings map[Month]int
}

/*
	CONSTRUCTOR
	This function will return a newly constructed StatisticsReport with default values.
*/
func NewStatisticsReport() *StatisticsReport {
	return &StatisticsReport{
		ClosedTransactions:               0,
		AcceptedTransactions:             0,
		RejectedTransactions:             0,
		FirstYearAcceptedMonthlyFilings:  initialiseMap(),
		SecondYearAcceptedMonthlyFilings: initialiseMap(),
	}
}

// Returns a map with months mapped to 0 values ready to be used.
func initialiseMap() map[Month]int {
	return map[Month]int{
		January:   0,
		February:  0,
		March:     0,
		April:     0,
		May:       0,
		June:      0,
		July:      0,
		August:    0,
		September: 0,
		October:   0,
		November:  0,
		December:  0,
	}
}
