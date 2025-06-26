package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abiji-2020/codetool/config"
)

func ConnectToDatabase() error {
	query := fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS codetool WITH 
    engine = 'pgvector', 
    parameters = {
        "host": "%s",
        "port": %d,
        "user": "%s",
        "password": "%s",
        "distance": "cosine",
        "database": "%s"
    };`, config.Host, config.Port, config.User, config.Password, config.MainDatabase)

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
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to connect to database: %s", resp.Status)
	}

	return nil
}

func CreateTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
        id SERIAL PRIMARY KEY,
        language VARCHAR(50) NOT NULL,
        snippet TEXT NOT NULL,
        repo TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        url TEXT NOT NULL
    );`, config.Database, config.Table)

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
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create table: %s", resp.Status)
	}

	return nil
}
