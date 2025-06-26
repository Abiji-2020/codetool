package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Abiji-2020/codetool/config"
)

type MindsDBClient struct {
	APIKey  string
	BaseUrl string
	Client  *http.Client
}

func NewMindsDBClient(apiKey string) *MindsDBClient {
	return &MindsDBClient{
		APIKey:  apiKey,
		BaseUrl: config.BASE_MINDSDB_URL,
		Client:  &http.Client{},
	}
}

func (c *MindsDBClient) makeRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}
	req, err := http.NewRequest(method, c.BaseUrl+endpoint, reqBody)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Warning: failed to close response body:", closeErr)
		}
	}()
	return io.ReadAll(resp.Body)

}
