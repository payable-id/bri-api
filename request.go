package bri

type CreateVaRequest struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Name            string `json:"nama"`
	Amount          string `json:"amount"`
	Description     string `json:"keterangan"`
	ExpiredDate     string `json:"expiredDate"`
}

type GetReportVaRequest struct {
	InstitutionCode string
	BrivaNo         string
	StartDate       string
	EndDate         string
}
