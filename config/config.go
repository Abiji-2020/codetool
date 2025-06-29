package config

const (
	BASE_MINDSDB_URL = "http://localhost:47334"
)

var (
	Database               = "codetool"
	Host                   = "pgvector"
	Port                   = 5432
	User                   = "pgvector_user"
	Password               = "pgvector_password"
	MainDatabase           = "pgvector_db"
	CodeDatabase           = "codetool"
	Distance               = "cosine"
	Table                  = "code_snippets"
	EmbeddingModelProvider = "ollama"
	EmbeddingModelName     = "bge-m3:latest"
	EmbeddingBaseUrl       = "http://ollama-intel-gpu:11434"
	OllamaBaseUrl          = "http://ollama-intel-gpu:11434"
	RerankingModelProvider = "ollama"
	RerankingModelName     = "llama3.2:1b"
	RerankingBaseUrl       = "http://ollama-intel-gpu:11434"
	StorageSchema          = "kb_snippet_storage"
	AgentModelName         = "llama3.2:1b"
	AgentProvider          = "ollama"
	AgentKnowledgeBase     = "mindsdb.snippets_kb"
	AgentName              = "codetool_agent"
	CustomModelName        = "codetool_custom_model"
)
