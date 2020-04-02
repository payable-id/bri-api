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
	ErrorResponse
}

// CardTokenOTPResponseData defines data response for direct debit - create card token OTP
type CardTokenOTPResponseData struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

// CardTokenOTPVerifyResponse defines response for direct debit - create card token OTP verify
type CardTokenOTPVerifyResponse struct {
	Body CardTokenOTPVerifyResponseData `json:"body"`
	ErrorResponse
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

// ErrorResponse defines response data if error request.Example:
// {
//     "error": {
//         "code": "0920",
//         "message": "Expired OTP"
//     },
//     "status_code": 400,
//     "status": {
//         "code": "0602",
//         "desc": "Invalid signature"
//     },
//     "recorded_at": "2020-03-12T03:58:46Z"
// }
type ErrorResponse struct {
	Error      ErrorDetail `json:"error"`
	StatusCode int         `json:"status_code"`
	Status     ErrorStatus `json:"status"`
}

// ErrorDetail defines response error detail. Example:
// {
//     "error": {
//         "code": "0920",
//         "message": "Expired OTP"
//     }
// }
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorStatus defines error data if unauthorized. Example:
// {
//     "status": {
//         "code": "0602",
//         "desc": "Invalid signature"
//     }
// }
type ErrorStatus struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

// PaymentChargeResponse defines response for direct debit - create payment charge [using OTP or not]
type PaymentChargeResponse struct {
	Body PaymentChargeResponseData `json:"body"`
	ErrorResponse
}

// PaymentChargeResponseData defines data response for direct debit - create payment charge [using OTP or not]
type PaymentChargeResponseData struct {
	Status        string                 `json:"status"`
	ChargeToken   string                 `json:"charge_token"`
	PaymentID     string                 `json:"payment_id"`
	Amount        string                 `json:"amount"`
	Currency      string                 `json:"currency"`
	Remarks       string                 `json:"remarks"`
	DeviceID      string                 `json:"device_id"`
	PaymentStatus string                 `json:"payment_status"`
	Location      Location               `json:"location"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// DeleteCardTokenResponse defines response for direct debit - delete card token
type DeleteCardTokenResponse struct {
	Body DeleteCardTokenResponseData `json:"body"`
	ErrorResponse
}

// DeleteCardTokenResponseData defines data response for direct debit - delete card token
type DeleteCardTokenResponseData struct {
	Status string `json:"status"`
}
