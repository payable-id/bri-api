package bri

type TokenResponse struct {
	AccessToken string   `json:"access_token"`
	ExpiredTime string   `json:"expires_in"`
	ProductList []string `json:"api_product_list_json"`
}

type VaResponse struct {
	Status              bool   `json:"status"`
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
	ErrDesc             string `json:"errDesc"`
	Data                VaData `json:"data"`
}

type VaData struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Name            string `json:"nama"`
	Amount          string `json:"amount"`
	Description     string `json:"keterangan"`
	ExpiredDate     string `json:"expiredDate"`
}

type VaReportResponse struct {
	Status       bool           `json:"status"`
	ResponseCode string         `json:"responseCode"`
	Description  string         `json:"responseDescription"`
	ErrDesc      string         `json:"errDesc"`
	Data         []VaReportData `json:"data"`
}

type VaReportData struct {
	BrivaNo     string `json:"brivaNo"`
	CustCode    string `json:"custCode"`
	Nama        string `json:"nama"`
	Amount      string `json:"amount"`
	Description string `json:"keterangan"`
	PaymentDate string `json:"paymentDate"`
	TellerId    string `json:"tellerid"`
	AccountNo   string `json:"no_rek"`
}

// CardTokenOTPResponse defines response for direct debit - create card token OTP
type CardTokenOTPResponse struct {
	Body CardTokenOTPResponseData `json:"body"`
}

// CardTokenOTPResponseData defines data response for direct debit - create card token OTP
type CardTokenOTPResponseData struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

// CardTokenOTPVerifyResponse defines response for direct debit - create card token OTP verify
type CardTokenOTPVerifyResponse struct {
	Body CardTokenOTPVerifyResponseData `json:"body"`
}

// CardTokenOTPVerifyResponseData defines data response for direct debit - create card token OTP verify
type CardTokenOTPVerifyResponseData struct {
	Status           string                 `json:"status"`
	PhoneNumber      string                 `json:"phone_number"`
	DeviceID         string                 `json:"device_id"`
	CardToken        string                 `json:"card_token"`
	Location         Location               `json:"location"`
	Last4            string                 `json:"last4"`
	Email            string                 `json:"email"`
	CardType         string                 `json:"card_type"`
	LimitTransaction string                 `json:"limit_transaction"`
	Metadata         map[string]interface{} `json:"metadata"`
}

// PaymentChargeOTPResponse defines response for direct debit - create payment charge OTP
type PaymentChargeOTPResponse struct {
	Body PaymentChargeOTPResponseData `json:"body"`
}

// PaymentChargeOTPResponseData defines data response for direct debit - create payment charge OTP
type PaymentChargeOTPResponseData struct {
	ChargeToken string `json:"charge_token"`
	Status      string `json:"status"`
}

// PaymentChargeOTPVerifyResponse defines response for direct debit - create payment charge OTP verify
type PaymentChargeOTPVerifyResponse struct {
	Body PaymentChargeOTPResponseData `json:"body"`
}

// PaymentChargeOTPVerifyResponseData defines data response for direct debit - create payment charge OTP verify
type PaymentChargeOTPVerifyResponseData struct {
	Status        string                 `json:"status"`
	PaymentID     string                 `json:"payment_id"`
	Amount        string                 `json:"amount"`
	Currency      string                 `json:"currency"`
	Remarks       string                 `json:"remarks"`
	DeviceID      string                 `json:"device_id"`
	PaymentStatus string                 `json:"payment_status"`
	Location      Location               `json:"location"`
	Metadata      map[string]interface{} `json:"metadata"`
}
