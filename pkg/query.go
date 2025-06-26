package pkg

import "encoding/json"

type SQLQueryRequest struct {
	Query string `json:"query"`
}

type SQLQueryResponse struct {
	ColumnNames []string        `json:"column_names"`
	Context     interface{}     `json:"context"`
	Data        [][]interface{} `json:"data"`
	Type        string          `json:"type"`
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
	if err != nil {
		return nil, err
	}
	return &response, nil
}
