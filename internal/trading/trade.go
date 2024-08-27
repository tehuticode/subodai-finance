package trading

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	WooXBaseURL = "https://api.woo.org"
)

type ExchangeClient struct {
	APIKey    string
	SecretKey string
	Client    *http.Client
	BaseURL   string
}

func NewExchangeClient(apiKey, secretKey string) *ExchangeClient {
	return &ExchangeClient{
		APIKey:    apiKey,
		SecretKey: secretKey,
		Client:    &http.Client{},
		BaseURL:   WooXBaseURL,
	}
}

func (e *ExchangeClient) PlaceOrder(symbol string, side string, orderType string, price float64, quantity float64) error {
	endpoint := "/v1/order"
	url := e.BaseURL + endpoint

	payload := map[string]interface{}{
		"symbol":     symbol,
		"side":       side,
		"order_type": orderType,
		"price":      price,
		"quantity":   quantity,
		"timestamp":  time.Now().UnixNano() / int64(time.Millisecond),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %w", err)
	}

	signature := e.generateSignature(string(jsonPayload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", e.APIKey)
	req.Header.Set("x-api-signature", signature)

	resp, err := e.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error response from exchange: %s", resp.Status)
	}

	return nil
}

func (e *ExchangeClient) generateSignature(payload string) string {
	h := hmac.New(sha256.New, []byte(e.SecretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
