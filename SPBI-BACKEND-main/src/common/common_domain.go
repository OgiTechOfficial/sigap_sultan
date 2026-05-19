package common

type ErrorDomain struct {
	Message string      `json:"message"`
	Details interface{} `json:"errors"`
}

type StartEndDateDomain struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
