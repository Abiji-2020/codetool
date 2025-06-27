package pkg

import (
	"fmt"
	"strings"

	"github.com/Abiji-2020/codetool/config"
)

type KBRecord struct {
	ID       string                 `json:"id"`
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (c *MindsDBClient) CreateKnowledgeBase(name string) error {
	query := fmt.Sprintf(`
	CREATE KNOWLEDGE_BASE IF NOT EXISTS %s 
	USING
		embedding_model = {
			"provider": "%s",
			"model_name": "%s",
			"base_url": "%s"
		},
		reranking_model = {
			"provider" : "%s",
			"model_name" : "%s",
			"base_url" : "%s"
		},
		storage = %s.%s, 
		metadata_columns = ["language", "repo", "url"],
		content_columns = ["code", "documentation"],
		id_column = "id";	
	`, name,
		config.EmbeddingModelProvider,
		config.EmbeddingModelName,
		config.EmbeddingBaseUrl,
		config.RerankingModelProvider,
		config.RerankingModelName,
		config.RerankingBaseUrl,
		config.CodeDatabase,
		config.StorageSchema,
	)
	_, err := c.ExecuteSQL(query)
	return err
}

func (c *MindsDBClient) InsertIntoKnowledgeBase(kbName string, data []KBRecord) error {
	for _, record := range data {
		query := fmt.Sprintf(`
		INSERT INTO %s (id, content, %s)
		VALUES ('%s', '%s', %s)
		`, kbName, formatMetadataColumns(record.Metadata),
			record.ID,
			record.Content,
			formatMetadataValues(record.Metadata),
		)
		_, err := c.ExecuteSQL(query)
		if err != nil {
			return fmt.Errorf("failed to insert record %s into knowledge base %s: %v", record.ID, kbName, err)
		}
	}
	return nil
}

func formatMetadataColumns(metadata map[string]interface{}) string {
	var columns []string
	for key := range metadata {
		columns = append(columns, key)
	}
	return fmt.Sprintf("'%s'", fmt.Sprint(strings.Join(columns, "','")))
}

func formatMetadataValues(metadata map[string]interface{}) string {
	var values []string
	for _, value := range metadata {
		values = append(values, fmt.Sprintf("'%v'", value))
	}
	return strings.Join(values, ", ")
}

func (c *MindsDBClient) SearchKnowledgeBase(kbName, searchContent string, relevance float32) (*SQLQueryResponse, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %s 
		WHERE content = '%s'
		AND relevance >= %f`,
		kbName,
		searchContent,
		relevance)
	return c.ExecuteSQL(query)
}

func (c *MindsDBClient) SearchKBWithMetadata(kbName, searchQuery string, metadata map[string]interface{}, relevance float32) (*SQLQueryResponse, error) {
	whereClause := fmt.Sprintf("content = '%s' AND relevance >= %f", searchQuery, relevance)
	for key, value := range metadata {
		whereClause += fmt.Sprintf(" AND %s = '%v'", key, value)
	}
	query := fmt.Sprintf(
		`SELECT * FROM %s
  		WHERE %s`,
		kbName,
		whereClause,
	)
	return c.ExecuteSQL(query)
}
