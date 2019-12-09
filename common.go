package bri

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

func getTimestamp(format string) (timestamp string) {
	dt := time.Now().UTC()
	timestamp = dt.Format(format)
	return
}

func generateSignature(path string, method string, token string, timestamp string, body string, secret string) (sig string) {
	payload := "path=" + path +
		"&verb=" + method +
		"&token=" + token +
		"&timestamp=" + timestamp +
		"&body=" + body

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))

	sig = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}
