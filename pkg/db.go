package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/Abiji-2020/codetool/config"
)

func (c *MindsDBClient) ConnectToDatabase() error {
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

	payload := map[string]any{
		"query": query,
	}

	response, err := c.makeRequest("POST", "/api/sql/query", payload)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Parse response to check for errors
	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there's an error in the response
	if errorMsg, exists := result["error"]; exists {
		return fmt.Errorf("database connection failed: %v", errorMsg)
	}

	return nil
}

func (c *MindsDBClient) CreateTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
        id SERIAL PRIMARY KEY,
        language VARCHAR(50) NOT NULL,
        snippet TEXT NOT NULL,
        repo TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        url TEXT NOT NULL
    );`, config.Database, config.Table)

	payload := map[string]any{
		"query": query,
	}

	response, err := c.makeRequest("POST", "/api/sql/query", payload)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	// Parse response to check for errors
	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there's an error in the response
	if errorMsg, exists := result["error_message"]; exists {
		return fmt.Errorf("table creation failed: %v", errorMsg)
	}

	return nil
}

func (c *MindsDBClient) InsertCodeSnippets(languageData map[string][]CodeSnippets) error {
	for language, snippets := range languageData {
		for _, snippet := range snippets {
			query := fmt.Sprintf(`INSERT INTO %s.%s (language, snippet, repo, url) VALUES ('%s', '%s', '%s', '%s');`,
				config.Database, config.Table, language, snippet.Code, snippet.Repo, snippet.URL)

			payload := map[string]any{
				"query": query,
			}

			response, err := c.makeRequest("POST", "/api/sql/query", payload)
			if err != nil {
				return fmt.Errorf("failed to insert snippet for %s: %v", language, err)
			}

			// Parse response to check for errors
			var result map[string]interface{}
			if err := json.Unmarshal(response, &result); err != nil {
				return fmt.Errorf("failed to parse insert response: %v", err)
			}

			if errorMsg, exists := result["error"]; exists {
				return fmt.Errorf("insert failed for %s: %v", language, errorMsg)
			}
		}
	}

	return nil
}
