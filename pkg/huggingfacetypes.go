package pkg

type CodeRow struct {
	RepositoryName        string `json:"repository_name"`
	FunctionPathInRepo    string `json:"func_path_in_repository"`
	FunctionName          string `json:"func_name"`
	CompleteFunction      string `json:"whole_func_string"`
	Language              string `json:"language"`
	FuncCodeString        string `json:"func_code_string"`
	FunctionDocumentation string `json:"func_documentation_string"`
	FunctionUrl           string `json:"func_code_url"`
}

type RowWrapper struct {
	Row    CodeRow `json:"row"`
	RowIdx int     `json:"row_idx"`
}

type HFResponse struct {
	Rows           []RowWrapper `json:"rows"`
	NumRowsTotal   int          `json:"num_rows_total"`
	NumRowsPerPage int          `json:"num_rows_per_page"`
	Partial        bool         `json:"partial"`
}

type LanguageRange struct {
	Name   string
	Start  int
	End    int
	Target int
}
