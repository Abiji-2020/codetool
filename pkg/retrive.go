package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	view "github.com/Abiji-2020/codetool/views"
)

type CodeSnippets struct {
	Code string
	Repo string
	URL  string
}

type ResultCodeSnippets struct {
	Language string
	Count    int
	Data     []CodeSnippets
}

func RetriveCodeSnippets(lang LanguageRange, wg *sync.WaitGroup, ch chan<- ResultCodeSnippets, progress chan<- view.ProgressMessage) {
	defer wg.Done()

	offset := lang.Start
	collected := []CodeSnippets{}
	batchSize := 100

	for offset < lang.End && len(collected) < lang.Target {
		url := fmt.Sprintf("https://datasets-server.huggingface.co/rows?dataset=claudios/code_search_net&config=all&split=train&offset=%d&length=%d", offset, batchSize)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching data for %s: %v\n", lang.Name, err)
			time.Sleep(1 * time.Second)
			continue
		}

		var res HFResponse
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			fmt.Printf("Error decoding response for %s: %v\n", lang.Name, err)
			err = resp.Body.Close()
			if err != nil {
				fmt.Println("Error closing response body:", err)
			}
			break
		}
		defer func() {
			if closeErr := resp.Body.Close(); closeErr != nil {
				fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
			}
		}()

		if len(res.Rows) == 0 {
			break
		}

		for _, row := range res.Rows {
			if row.Row.Language == lang.Name {
				code := row.Row.CompleteFunction + "\n" + row.Row.FunctionDocumentation
				url := row.Row.FunctionUrl
				repo := row.Row.RepositoryName
				collected = append(collected, CodeSnippets{
					Repo: repo,
					URL:  url,
					Code: code,
				})
				if len(collected) >= lang.Target {
					break
				}
			}
		}

		progress <- view.ProgressMessage{
			Language: lang.Name,
			Current:  len(collected),
			Total:    lang.Target,
			Done:     len(collected) >= lang.Target,
		}
		offset += batchSize
		time.Sleep(200 * time.Millisecond)

	}

	ch <- ResultCodeSnippets{
		Language: lang.Name,
		Count:    len(collected),
		Data:     collected,
	}
}
