package main

import "time"

type nameChange struct {
	DateTime time.Time `json:"DateTime"`
	NewName  string    `json:"newName"`
}

type userMessage struct {
	DateTime time.Time `json:"DateTime"`
	Content  string    `json:"Content"`
}

type message struct {
	DateTime time.Time `json:"DateTime"`
	Content  string    `json:"Content"`
	User     string    `json:"User"`
}

type dateMessage struct {
	Content string `json:"Content"`
	User    string `json:"User"`
}

type barChartOutput struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type lineChartOutput struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type daySummaryInfo struct {
	NumberOfMesages           int     `json:"NumberOfMessages"`
	AverageMessageLength      float64 `json:"AverageMessageLengthWords"`
	UserWhoSentLongestMessage string  `json:"UserWhoSentLongestMessage"`
	LengthOfLongestMessage    int     `json:"LengthOfLongestMessageWords"`
}

type overallSummaryInfo struct {
	MostPopularMessage               string `json:"MostPopularMessage"`
	MostPopularMessageCount          int    `json:"MostPopularMessageCount"`
	NumberOfMessagesSent             int    `json:"NumberOfMessagesSent"`
	NumberOfDaysWithActivity         int    `json:"NumberOfDaysWithActivity"`
	NumberOfDaysAnalysed             int    `json:"NumberOfDaysAnalysed"`
	MostActiveUser                   string `json:"MostActiveUser"`
	NumberOfMessagesByMostActiveUser int    `json:"NumberOfMessagesByMostActiveUser"`
}
