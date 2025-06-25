package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abiji-2020/codetool/config"
)

func ConnectToDatabase() error {
	query := `CREATE DATABASE IF NOT EXISTS codetool with 
	engine = 'pgvector', 
	parameters = {
	"host": "pgvector",
	"port": 5432,
	"user":"pgvector_user",
	"password":"pgvector_password",
	"distance": "cosine",
	"database": "pgvector_db"
  };`
	httpClient := &http.Client{}
	endpoint := "/api/sql/query"

	payload := map[string]any{
		"query": query,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}
	req, err := http.NewRequest("POST", config.BASE_MINDSDB_URL+endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to connect to database: %s", resp.Status)
	}

	return nil
}

func CreateTable() error {
	query := `CREATE TABLE codetool.code_snippets (
	id SERIAL PRIMARY KEY,
	language VARCHAR(50) NOT NULL,
	snippet TEXT NOT NULL,
	repo TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	url TEXT NOT NULL
	);`
	httpClient := &http.Client{}
	endpoint := "/api/sql/query"
	payload := map[string]any{
		"query": query,
	}
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}
	req, err := http.NewRequest("POST", config.BASE_MINDSDB_URL+endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create table: %s", resp.Status)
	}

	return nil
}
