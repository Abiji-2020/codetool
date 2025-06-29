package pkg

import (
	"fmt"

	"github.com/Abiji-2020/codetool/config"
)

func (c *MindsDBClient) CreateEngine() error {
	query := fmt.Sprintf(`
	CREATE  ML_ENGINE IF NOT EXISTS ollama_engine
	FROM ollama
	USING 
		base_url = '%s'
	;`, config.OllamaBaseUrl)

	response, err := c.ExecuteSQL(query)
	if err != nil {
		return fmt.Errorf("failed to create engine: %v", err)
	}
	// Check if the response contains an error message
	if response.Type == "error" {
		return fmt.Errorf("engine creation failed: %s", response.Data[0][0])
	}

	return nil
}

func (c *MindsDBClient) CreateModel() error {
	query := fmt.Sprintf(`
	 CREATE MODEL IF NOT EXISTS %s
	 PREDICT answer 
	 USING 
	 	engine = ollama_engine,
		model_name = %s,
		prompt_template = 'Answer based on the given context. If the answer is not found in the context, say I dont know.'
		`, config.CustomModelName, config.AgentModelName)
	response, err := c.ExecuteSQL(query)
	if err != nil {
		return fmt.Errorf("failed to create model: %v", err)
	}
	// Check if the response contains an error message
	if response.Type == "error" {
		return fmt.Errorf("model creation failed: %s", response.Data[0][0])
	}
	return nil
}
