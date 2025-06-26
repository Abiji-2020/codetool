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
	Distance               = "cosine"
	Table                  = "code_snippets"
	EmbeddingModelProvider = "ollama"
	EmbeddingModelName     = ""
	EmbeddingBaseUrl       = ""
	RerankingModelProvider = "ollama"
	RerankingModelName     = ""
	RerankingBaseUrl       = ""
	StorageSchema          = ""
)
