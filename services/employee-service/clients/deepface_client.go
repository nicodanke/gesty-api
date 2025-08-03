package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeepfaceClient struct {
	BaseURL string
	Client  *http.Client
}

func (d *DeepfaceClient) Get(endpoint string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", d.BaseURL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return d.Client.Do(req)
}

// POST method with JSON body
func (d *DeepfaceClient) PostJSON(endpoint string, payload any) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", d.BaseURL, endpoint)

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return d.Client.Do(req)
}
