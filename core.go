package bri

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

const (
	TOKEN_PATH      = "/oauth/client_credential/accesstoken?grant_type=client_credentials"
	VA_PATH         = "/v1/briva"
	VA_REPORT_PATH  = "/v1/briva/report"
	BRI_TIME_FORMAT = "2006-01-02T15:04:05.999Z"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) GetToken() (res TokenResponse, err error) {
	data := url.Values{}
	data.Set("client_id", gateway.Client.ClientId)
	data.Set("client_secret", gateway.Client.ClientSecret)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", TOKEN_PATH, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) CreateVA(token string, req CreateVaRequest) (res VaResponse, err error) {
	token = "Bearer " + token
	method := "POST"
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(VA_PATH, method, token, timestamp, string(body), gateway.Client.ClientSecret)

	headers := map[string]string{
		"Authorization": token,
		"BRI-Timestamp": timestamp,
		"BRI-Signature": signature,
		"Content-Type":  "application/json",
	}

	err = gateway.Call(method, VA_PATH, headers, strings.NewReader(string(body)), &res)

	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) UpdateVA(token string, req CreateVaRequest) (res VaResponse, err error) {
	token = "Bearer " + token
	method := "PUT"
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(VA_PATH, method, token, timestamp, string(body), gateway.Client.ClientSecret)

	headers := map[string]string{
		"Authorization": token,
		"BRI-Timestamp": timestamp,
		"BRI-Signature": signature,
		"Content-Type":  "application/json",
	}

	err = gateway.Call(method, VA_PATH, headers, strings.NewReader(string(body)), &res)

	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) GetReportVA(token string, req GetReportVaRequest) (res VaReportResponse, err error) {
	token = "Bearer " + token
	method := "GET"
	body := ""
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	path := VA_REPORT_PATH + "/" + req.InstitutionCode + "/" + req.BrivaNo + "/" + req.StartDate + "/" + req.EndDate
	signature := generateSignature(path, method, token, timestamp, string(body), gateway.Client.ClientSecret)

	headers := map[string]string{
		"Authorization": token,
		"BRI-Timestamp": timestamp,
		"BRI-Signature": signature,
	}

	err = gateway.Call(method, path, headers, strings.NewReader(string(body)), &res)

	if err != nil {
		return
	}

	return
}
