package pkg

import (
	"fmt"

	"github.com/Abiji-2020/codetool/config"
)

func (c *MindsDBClient) CreateAgent(name string) error {
	query := fmt.Sprintf(`
	CREATE AGENT IF NOT EXISTS %s
	USING 
		model = '%s',
		include_knowledge_bases = ['%s'],
		prompt_template = '
		   You are a helpful coding assistant with access to a knowledge base containing code snippets and documentation.
        
        The knowledge base contains:
        - Code snippets in various programming languages
        - Documentation explaining what each code snippet does
        - Repository information and URLs for reference
        
        When answering questions:
        1. Search the knowledge base for relevant code examples
        2. Provide the code snippet if available
        3. Explain what the code does based on the documentation
        4. Include language and repository information when helpful
        
        Question: {{question}}';
	`, name, config.CustomModelName, config.AgentKnowledgeBase)
	_, err := c.ExecuteSQL(query)
	if err != nil {
		return fmt.Errorf("failed to create agent: %v", err)
	}
	return nil
}
