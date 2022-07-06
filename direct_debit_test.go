package bri

import (
	"github.com/stretchr/testify/assert"
)

func (bri *BriSanguTestSuite) TestDirectDebit_01_CreateCardToken() {
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

	bri.paymentID = resp.Body.PaymentID
	bri.amount = resp.Body.Amount
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

	bri.paymentID = resp.Body.PaymentID
	bri.amount = resp.Body.Amount
}

func (bri *BriSanguTestSuite) TestDirectDebit_06_GetChargeDetail_Found() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := ChargeDetailRequest{
		Body: ChargeDetailRequestData{
			PaymentID: bri.paymentID,
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.GetChargeDetail(token, req)

	assert.Equal(bri.T(), "0000", resp.Body.Status)
	assert.Equal(bri.T(), "SUCCESS", resp.Body.PaymentStatus)
	assert.Equal(bri.T(), bri.paymentID, resp.Body.PaymentID)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestDirectDebit_07_GetChargeDetail_NotFound() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := ChargeDetailRequest{
		Body: ChargeDetailRequestData{
			PaymentID: "3435",
		},
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.GetChargeDetail(token, req)

	assert.Equal(bri.T(), 400, resp.ErrorResponse.StatusCode)
	assert.Equal(bri.T(), "0301", resp.Error.Code)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestDirectDebit_08_Refund() {
	coreGateway := CoreGateway{
		Client: bri.client,
	}

	tokenResp, _ := coreGateway.GetToken()
	req := RefundRequest{
		Body: RefundRequestData{
			CardToken: bri.cardToken,
			Amount:    bri.amount,
			PaymentID: bri.paymentID,
			Currency:  "IDR",
			Reason:    "test refund",
			Metadata: map[string]interface{}{
				"": nil,
			},
		},
	}

	token := tokenResp.AccessToken
	idempotencyKey := generateSha1Timestamp("08_Refund")
	resp, err := coreGateway.RefundDirectDebit(token, idempotencyKey, req)

	assert.Equal(bri.T(), "0000", resp.Body.Status)
	assert.Equal(bri.T(), "SUCCESS", resp.Body.RefundStatus)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestDirectDebit_09_DeleteCardToken() {
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
