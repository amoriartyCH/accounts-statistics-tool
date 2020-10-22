package models

type StatisticsReport struct {

	ClosedTransactions					int
	AcceptedTransactions				int
	RejectedTransactions				int

	FirstYearAcceptedMonthlyFilings 	map[int]int
	SecondYearAcceptedMonthlyFilings 	map[int]int
}

/*
	CONSTRUCTOR
	This function will return a newly constructed StatisticsReport with default values.
 */
func NewStatisticsReport() *StatisticsReport {
	return &StatisticsReport{
		ClosedTransactions: 0,
		AcceptedTransactions: 0,
		RejectedTransactions: 0,
		FirstYearAcceptedMonthlyFilings:  make(map[int]int, 0),
		SecondYearAcceptedMonthlyFilings: make(map[int]int, 0),
	}
}