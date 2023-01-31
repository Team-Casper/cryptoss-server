package sms

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/team-casper/cryptoss-server/config"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Send(conf *config.Config, to, msgContent string) error {
	body := CreateSMSMsg(conf.FromNumber, to, msgContent)

	dataParams, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal sms request body: %w", err)
	}
	buf := bytes.NewBuffer(dataParams)

	req, err := http.NewRequest(http.MethodPost, conf.SMSEndpoint, buf)

	if err := SetHeader(req, conf); err != nil {
		return fmt.Errorf("error occurs while setting request header for sms: %w", err)
	}

	client := http.Client{}
	if _, err := client.Do(req); err != nil {
		return fmt.Errorf("error occurs while sending sms: %w", err)
	}

	return nil
}

func SetHeader(req *http.Request, conf *config.Config) error {
	timestamp := time.Now()
	// set header
	timestampStr := strconv.FormatInt(timestamp.UnixMilli(), 10)
	sig, err := Sign(conf, timestampStr)
	if err != nil {
		return fmt.Errorf("failed to make signature: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ncp-apigw-timestamp", timestampStr)
	req.Header.Set("x-ncp-iam-access-key", conf.AccessKeyId)
	req.Header.Set("x-ncp-apigw-signature-v2", sig)

	return nil
}

func Sign(conf *config.Config, timestamp string) (string, error) {
	const space = " "
	const newLine = "\n"

	reqUrl := conf.SMSEndpoint

	u, err := url.Parse(reqUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse url(%s): %w", reqUrl, err)
	}

	h := hmac.New(sha256.New, []byte(conf.SecretKey))
	h.Write([]byte("POST"))
	h.Write([]byte(space))
	h.Write([]byte(u.RequestURI()))
	h.Write([]byte(newLine))
	h.Write([]byte(timestamp))
	h.Write([]byte(newLine))
	h.Write([]byte(conf.AccessKeyId))
	rawSignature := h.Sum(nil)

	base64signature := base64.StdEncoding.EncodeToString(rawSignature)

	return base64signature, nil
}
