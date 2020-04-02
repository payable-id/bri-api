package bri

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	urlCreateCardTokenOTP           = "/v1/directdebit/tokens"         // POST
	urlCreateCardTokenOTPVerify     = "/v1/directdebit/tokens"         // PATCH
	urlDeleteCardToken              = "/v1/directdebit/tokens"         // DELETE
	urlCreatePaymentChargeOTP       = "/v1/directdebit/charges"        // POST
	urlCreatePaymentChargeOTPVerify = "/v1/directdebit/charges/verify" // POST
)

// CreateCardTokenOTP verifies that the information provided by the customers matches the bank data.
// This API will alse send OTP code confirmation to user if user phonenumber is valid.
func (g *CoreGateway) CreateCardTokenOTP(token string, req CardTokenOTPRequest) (res CardTokenOTPResponse, err error) {
	req.Body.OtpBriStatus = "YES"

	token = "Bearer " + token
	method := http.MethodPost
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(urlCreateCardTokenOTP, method, token, timestamp, string(body), g.Client.ClientSecret)

	headers := map[string]string{
		"Authorization":   token,
		"BRI-Timestamp":   timestamp,
		"X-BRI-Signature": signature,
		"Content-Type":    "application/json",
		"X-BRI-Api-Key":   g.Client.APIKey,
	}

	err = g.CallDirectDebit(method, urlCreateCardTokenOTP, headers, strings.NewReader(string(body)), &res)
	return
}

// CreateCardTokenOTPVerify is used to verify OTP from create card token OTP url.
func (g *CoreGateway) CreateCardTokenOTPVerify(token string, req CardTokenOTPVerifyRequest) (res CardTokenOTPVerifyResponse, err error) {
	token = "Bearer " + token
	method := http.MethodPatch
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(urlCreateCardTokenOTPVerify, method, token, timestamp, string(body), g.Client.ClientSecret)

	headers := map[string]string{
		"Authorization":   token,
		"BRI-Timestamp":   timestamp,
		"X-BRI-Signature": signature,
		"Content-Type":    "application/json",
		"X-BRI-Api-Key":   g.Client.APIKey,
	}

	err = g.CallDirectDebit(method, urlCreateCardTokenOTPVerify, headers, strings.NewReader(string(body)), &res)
	return
}

// DeleteCardToken is used to unbind user's direct debit card token
func (g *CoreGateway) DeleteCardToken(token string, req DeleteCardTokenRequest) (res DeleteCardTokenResponse, err error) {
	token = "Bearer " + token
	method := http.MethodDelete
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(urlDeleteCardToken, method, token, timestamp, string(body), g.Client.ClientSecret)

	headers := map[string]string{
		"Authorization":   token,
		"BRI-Timestamp":   timestamp,
		"X-BRI-Signature": signature,
		"Content-Type":    "application/json",
		"X-BRI-Api-Key":   g.Client.APIKey,
	}

	err = g.CallDirectDebit(method, urlDeleteCardToken, headers, strings.NewReader(string(body)), &res)
	return
}

// CreatePaymentChargeOTP is used for payment of direct link transactions based on card number via card_token acquired from binding process (create a card token).
// This API will alse send OTP code confirmation to user if user phonenumber is valid.
func (g *CoreGateway) CreatePaymentChargeOTP(token, idempotencyKey string, req PaymentChargeOTPRequest) (res PaymentChargeResponse, err error) {
	token = "Bearer " + token
	method := http.MethodPost
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(urlCreatePaymentChargeOTP, method, token, timestamp, string(body), g.Client.ClientSecret)

	headers := map[string]string{
		"Authorization":   token,
		"BRI-Timestamp":   timestamp,
		"X-BRI-Signature": signature,
		"Content-Type":    "application/json",
		"Idempotency-Key": idempotencyKey,
		"X-BRI-Api-Key":   g.Client.APIKey,
	}

	err = g.CallDirectDebit(method, urlCreatePaymentChargeOTP, headers, strings.NewReader(string(body)), &res)
	return
}

// CreatePaymentChargeOTPVerify is used to verify OTP from create payment charge OTP url.
func (g *CoreGateway) CreatePaymentChargeOTPVerify(token string, req PaymentChargeOTPVerifyRequest) (res PaymentChargeResponse, err error) {
	token = "Bearer " + token
	method := http.MethodPost
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(urlCreatePaymentChargeOTPVerify, method, token, timestamp, string(body), g.Client.ClientSecret)

	headers := map[string]string{
		"Authorization":   token,
		"BRI-Timestamp":   timestamp,
		"X-BRI-Signature": signature,
		"Content-Type":    "application/json",
		"X-BRI-Api-Key":   g.Client.APIKey,
	}

	err = g.CallDirectDebit(method, urlCreatePaymentChargeOTPVerify, headers, strings.NewReader(string(body)), &res)
	return
}
