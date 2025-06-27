package pkg

import (
	"encoding/json"
	"fmt"
)

type SQLQueryRequest struct {
	Query string `json:"query"`
}

type SQLQueryResponse struct {
	ColumnNames  []string        `json:"column_names"`
	Context      interface{}     `json:"context"`
	Data         [][]interface{} `json:"data"`
	Type         string          `json:"type"`
	ErrorMessage string          `json:"error_message,omitempty"`
}

func (c *MindsDBClient) ExecuteSQL(query string) (*SQLQueryResponse, error) {
	reqBody := SQLQueryRequest{
		Query: query,
	}

	respBody, err := c.makeRequest("POST", "/api/sql/query", reqBody)
	if err != nil {
		return nil, err
	}
	var response SQLQueryResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil || response.Type == "error" || response.ErrorMessage != "" {
		return nil, fmt.Errorf("failed to execute SQL query: %s, error: %v", response.ErrorMessage, err)
	}
	return &response, nil
}
