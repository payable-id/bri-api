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

// Location defines location data
type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// CardTokenOTPRequest defines payload for direct debit - create card token OTP
type CardTokenOTPRequest struct {
	Body CardTokenOTPRequestData `json:"body"`
}

// CardTokenOTPRequestData defines item data payload for direct debit - create card token OTP
type CardTokenOTPRequestData struct {
	CardPan      string `json:"card_pan"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	OtpBriStatus string `json:"otp_bri_status"`
}

// CardTokenOTPVerifyRequest defines payload for direct debit - create card token OTP verify
type CardTokenOTPVerifyRequest struct {
	Body CardTokenOTPVerifyRequestData `json:"body"`
}

// CardTokenOTPVerifyRequestData defines item data payload for direct debit - create card token OTP verify
type CardTokenOTPVerifyRequestData struct {
	RegistrationToken string `json:"registration_token"`
	Passcode          string `json:"passcode"`
}

// PaymentChargeOTPRequest defines payload for direct debit - create payment charge OTP
type PaymentChargeOTPRequest struct {
	Body PaymentChargeOTPRequestData `json:"body"`
}

// PaymentChargeOTPRequestData defines data payload for direct debit - create payment charge OTP
type PaymentChargeOTPRequestData struct {
	CardToken    string                 `json:"card_token"`
	Amount       string                 `json:"amount"`
	Currency     string                 `json:"currency"`
	Remarks      string                 `json:"remarks"`
	OtpBriStatus string                 `json:"otp_bri_status"`
	Metadata     map[string]interface{} `json:"metadata"`
}

// PaymentChargeOTPVerifyRequest defines payload for direct debit - create payment charge OTP verify
type PaymentChargeOTPVerifyRequest struct {
	Body PaymentChargeOTPVerifyRequestData `json:"body"`
}

// PaymentChargeOTPVerifyRequestData defines data payload for direct debit - create payment charge OTP
type PaymentChargeOTPVerifyRequestData struct {
	CardToken   string `json:"card_token"`
	ChargeToken string `json:"charge_token"`
	Passcode    string `json:"passcode"`
}

// DeleteCardTokenRequest defines payload for direct debit - delete card token
type DeleteCardTokenRequest struct {
	Body DeleteCardTokenRequestData `json:"body"`
}

// DeleteCardTokenRequestData defines data payload for direct debit - delete card token
type DeleteCardTokenRequestData struct {
	CardToken string `json:"card_token"`
}

// ChargeDetailRequest defines payload for direct debit - charge detail
type ChargeDetailRequest struct {
	Body ChargeDetailRequestData `json:"body"`
}

// ChargeDetailRequestData defines data payload for direct debit - charge detail
type ChargeDetailRequestData struct {
	PaymentID string                 `json:"payment_id"`
	Remarks   string                 `json:"remarks"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// RefundRequest defines payload for direct debit - refund
type RefundRequest struct {
	Body RefundRequestData `json:"body"`
}

// RefundRequestData defines data payload for direct debit - refund
type RefundRequestData struct {
	CardToken string                 `json:"card_token"`
	Amount    string                 `json:"amount"`
	PaymentID string                 `json:"payment_id"`
	Currency  string                 `json:"currency"`
	Reason    string                 `json:"reason"`
	Metadata  map[string]interface{} `json:"metadata"`
}

type GetMutationRequest struct {
	AccountNumber string `json:"accountNumber"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
}
