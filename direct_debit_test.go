package bri

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/stretchr/testify/assert"
)

// generateSha1Timestamp will generate sha1 hash from UnixNano timestamp
func generateSha1Timestamp(salt string) string {
	key := fmt.Sprintf("%s-%d", salt, time.Now().UnixNano())

	h := sha1.New()
	h.Write([]byte(key))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (bri *BriSanguTestSuite) TestDirectDebit_01_CreateCardToken() {
	// modify url for unit test
	urlCreateCardTokenOTP = fmt.Sprintf("/sandbox%s", urlCreateCardTokenOTP)
	urlCreateCardTokenOTPVerify = fmt.Sprintf("/sandbox%s", urlCreateCardTokenOTPVerify)
	urlCreatePaymentChargeOTP = fmt.Sprintf("/sandbox%s", urlCreatePaymentChargeOTP)
	urlCreatePaymentChargeOTPVerify = fmt.Sprintf("/sandbox%s", urlCreatePaymentChargeOTPVerify)
	urlDeleteCardToken = fmt.Sprintf("/sandbox%s", urlDeleteCardToken)

	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := CardTokenOTPRequest{
		Body: CardTokenOTPRequestData{
			CardPan:      bri.cardPan,
			PhoneNumber:  bri.phoneNumber,
			Email:        bri.email,
			OtpBriStatus: "YES",
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreateCardTokenOTP(token, req)

	assert.Equal(bri.T(), "PENDING_USER_VERIFICATION", resp.Body.Status)
	assert.Equal(bri.T(), nil, err)

	bri.registrationCardToken = resp.Body.Token
}

func (bri *BriSanguTestSuite) TestDirectDebit_02_VerifyCreateCardToken() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := CardTokenOTPVerifyRequest{
		Body: CardTokenOTPVerifyRequestData{
			RegistrationToken: bri.registrationCardToken,
			Passcode:          "999999", // default dev otp code
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreateCardTokenOTPVerify(token, req)

	assert.Equal(bri.T(), "0000", resp.Body.Status)
	assert.Equal(bri.T(), nil, err)

	bri.cardToken = resp.Body.CardToken
}

func (bri *BriSanguTestSuite) TestDirectDebit_03_ChargePaymentNoOTP() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	idempotencyKey := generateSha1Timestamp("03_ChargePaymentOTP")
	req := PaymentChargeOTPRequest{
		Body: PaymentChargeOTPRequestData{
			CardToken:    bri.cardToken,
			Amount:       "89000.00",
			Currency:     "IDR",
			Remarks:      "testing remarks from unit test",
			OtpBriStatus: "NO",
			Metadata: map[string]interface{}{
				"": nil,
			},
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreatePaymentChargeOTP(token, idempotencyKey, req)

	assert.Equal(bri.T(), "0000", resp.Body.Status)
	assert.Equal(bri.T(), "SUCCESS", resp.Body.PaymentStatus)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestDirectDebit_04_ChargePaymentOTP() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	idempotencyKey := generateSha1Timestamp("04_ChargePaymentOTP")
	req := PaymentChargeOTPRequest{
		Body: PaymentChargeOTPRequestData{
			CardToken:    bri.cardToken,
			Amount:       "63000.00",
			Currency:     "IDR",
			Remarks:      "testing remarks from unit test again",
			OtpBriStatus: "YES",
			Metadata: map[string]interface{}{
				"": nil,
			},
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreatePaymentChargeOTP(token, idempotencyKey, req)

	assert.Equal(bri.T(), "PENDING_USER_VERIFICATION", resp.Body.Status)
	assert.Equal(bri.T(), nil, err)

	bri.chargeToken = resp.Body.ChargeToken
}

func (bri *BriSanguTestSuite) TestDirectDebit_05_VerifyChargePaymentOTP() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := PaymentChargeOTPVerifyRequest{
		Body: PaymentChargeOTPVerifyRequestData{
			CardToken:   bri.cardToken,
			ChargeToken: bri.chargeToken,
			Passcode:    "999999", // default dev otp code
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreatePaymentChargeOTPVerify(token, req)

	assert.Equal(bri.T(), "0000", resp.Body.Status)
	assert.Equal(bri.T(), "SUCCESS", resp.Body.PaymentStatus)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestDirectDebit_06_DeleteCardToken() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := DeleteCardTokenRequest{
		Body: DeleteCardTokenRequestData{
			CardToken: bri.cardToken,
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.DeleteCardToken(token, req)

	assert.Equal(bri.T(), "0000", resp.Body.Status)
	assert.Equal(bri.T(), nil, err)
}
