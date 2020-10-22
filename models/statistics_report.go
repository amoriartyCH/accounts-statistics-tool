package models

type StatisticsReport struct {

	FirstYearAcceptedMonthlyFilings 	map[string]int
	SecondYearAcceptedMonthlyFilings 	map[int]int
}

/*
	CONSTRUCTOR
	This function will return a newly constructed StatisticsReport with default values.
 */
func NewStatisticsReport() *StatisticsReport {
	return &StatisticsReport{
		FirstYearAcceptedMonthlyFilings:  make(map[string]int, 0),
		SecondYearAcceptedMonthlyFilings: make(map[int]int, 0),
	}
}

