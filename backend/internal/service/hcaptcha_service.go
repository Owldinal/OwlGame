package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"owl-backend/internal/config"
	"strings"
)

type hCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"` // optional
}

func VerifyCaptcha(hcaptcha string) (bool, error) {
	postData := url.Values{
		"response": {hcaptcha},
		"secret":   {config.C.HCaptchaKey},
	}

	reqBody := strings.NewReader(postData.Encode())
	resp, err := http.Post("https://api.hcaptcha.com/siteverify", "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		return false, fmt.Errorf("http.Post failed: %v", err)
	}
	defer resp.Body.Close()

	var hcResp hCaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&hcResp); err != nil {
		return false, fmt.Errorf("json.Decode failed: %v", err)
	}

	if !hcResp.Success {
		return false, fmt.Errorf("hCaptcha verification failed: %v", hcResp.ErrorCodes)
	}

	return true, nil
}
