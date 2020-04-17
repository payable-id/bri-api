package bri

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BriSanguTestSuite struct {
	suite.Suite
	client Client

	// property for direct debit
	cardPan               string
	phoneNumber           string
	email                 string
	registrationCardToken string
	cardToken             string
	chargeToken           string
}

type credentials struct {
	BaseUrl            string
	DirectDebitBaseURL string
	ClientId           string
	ClientSecret       string
	APIKey             string
	IsSandbox          bool
	CardPan            string
	PhoneNumber        string
	Email              string
}

func TestBriSanguTestSuite(t *testing.T) {
	suite.Run(t, new(BriSanguTestSuite))
}

func (bri *BriSanguTestSuite) SetupTest() {
	theToml, err := ioutil.ReadFile("credential_test.toml")
	if err != nil {
		bri.T().Log(err)
		bri.T().FailNow()
	}

	var cred credentials
	if _, err := toml.Decode(string(theToml), &cred); err != nil {
		bri.T().Log(err)
		bri.T().FailNow()
	}

	bri.client = NewClient()
	bri.client.BaseUrl = cred.BaseUrl
	bri.client.DirectDebitBaseURL = cred.DirectDebitBaseURL
	bri.client.ClientId = cred.ClientId
	bri.client.ClientSecret = cred.ClientSecret
	bri.client.APIKey = cred.APIKey
	bri.client.DirectDebitHostUseSandboxPrefix(cred.IsSandbox)

	bri.cardPan = cred.CardPan
	bri.phoneNumber = cred.PhoneNumber
	bri.email = cred.Email
}

func (bri *BriSanguTestSuite) TestGetTokenSuccess() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	resp, _, err := coreGateway.GetToken()
	containsProduct := false
	for _, v := range resp.ProductList {
		if strings.Contains(v, "briva") {
			containsProduct = true
			break
		}
	}

	assert.NotNil(bri.T(), resp.AccessToken)
	assert.Equal(bri.T(), "179999", resp.ExpiredTime)
	assert.Equal(bri.T(), true, containsProduct)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetTokenFailedInvalidKeySecret() {
	bri.client.ClientSecret = "123"

	coreGateway := CoreGateway{
		Client: bri.client,
	}

	resp, _, err := coreGateway.GetToken()

	assert.Equal(bri.T(), "", resp.AccessToken)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetTokenFailedInvalidUrl() {
	bri.client.BaseUrl = "https://sandbox.partner.api.bri.co.id/v1"

	coreGateway := CoreGateway{
		Client: bri.client,
	}

	resp, _, err := coreGateway.GetToken()

	assert.NotNil(bri.T(), err)
	assert.Equal(bri.T(), "", resp.AccessToken)
}

func (bri *BriSanguTestSuite) TestGetTokenFailedInvalidProduct() {
	bri.client.ClientId = "QYTaZS7Rw3JztRWhXAYrXjKUAg13AvRa"

	coreGateway := CoreGateway{
		Client: bri.client,
	}

	resp, _, err := coreGateway.GetToken()

	containsProduct := false
	for _, v := range resp.ProductList {
		if strings.Contains(v, "briva") {
			containsProduct = true
			break
		}
	}

	assert.Equal(bri.T(), false, containsProduct)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestCreateVaSuccess() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 0)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "123123" + random,
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.CreateVA(token, req)

	assert.Equal(bri.T(), true, resp.Status)
	assert.Equal(bri.T(), "00", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestCreateVaFailedDuplicate() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 0)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "123123" + random,
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.CreateVA(token, req)

	// create second request
	resp, _, err = coreGateway.CreateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "13", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestCreateVaFailedExpiredMoreThanThreeMonths() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 1)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "123123" + random,
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.CreateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "12", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestUpdateVaSuccess() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 0)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "1231233313",
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.UpdateVA(token, req)

	assert.Equal(bri.T(), true, resp.Status)
	assert.Equal(bri.T(), "00", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestUpdateVaFailedCustomerNotFound() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 1)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "1231233313555",
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.UpdateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "14", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestUpdateVaFailedExpiredMoreThanThreeMonths() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 1)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "1231233313",
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.UpdateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "12", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetReportVaSuccess() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	req := GetReportVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		StartDate:       "20190918",
		EndDate:         "20190918",
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.GetReportVA(token, req)

	assert.Equal(bri.T(), true, resp.Status)
	assert.Equal(bri.T(), "00", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetReportVaFailedNoTransaction() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	req := GetReportVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		StartDate:       "20190919",
		EndDate:         "20190919",
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.GetReportVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "41", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetReportVaFailedInvalidDateRange() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}
	tokenResp, _, err := coreGateway.GetToken()

	req := GetReportVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		StartDate:       "20190916",
		EndDate:         "20190918",
	}

	token := tokenResp.AccessToken
	resp, _, err := coreGateway.GetReportVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "42", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}
