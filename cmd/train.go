package cmd

import (
	"fmt"
	"sync"

	"github.com/Abiji-2020/codetool/pkg"
	"github.com/Abiji-2020/codetool/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var count int32
var trainCmd = &cobra.Command{
	Use:     "train",
	Short:   "Train a model with the given dataset",
	Long:    `Train a model with the count given in the argument --count and defaults to 1000.`,
	Aliases: []string{"t"},
	Run: func(cmd *cobra.Command, args []string) {
		println("Training model with count:", count)
		langs := []pkg.LanguageRange{
			{Name: "python", Start: 0, End: 412177, Target: int(count)},
			{Name: "java", Start: 412178, End: 866628, Target: int(count)},
			{Name: "javascript", Start: 866629, End: 990517, Target: int(count)},
			{Name: "go", Start: 990518, End: 1308349, Target: int(count)},
			{Name: "ruby", Start: 1308350, End: 1357140, Target: int(count)},
			{Name: "php", Start: 1357141, End: 1889853, Target: int(count)},
		}

		var wg sync.WaitGroup

		ch := make(chan pkg.ResultCodeSnippets, len(langs))
		progressCh := make(chan views.ProgressMessage, 100)
		langTargets := make(map[string]int)
		for _, lang := range langs {
			langTargets[lang.Name] = lang.Target
		}

		model := views.NewModel(langTargets)
		p := tea.NewProgram(model)

		go func() {
			for msg := range progressCh {
				p.Send(msg) // ✅ forward progress updates into the TUI
			}
		}()

		for _, lang := range langs {
			wg.Add(1)
			go pkg.RetriveCodeSnippets(lang, &wg, ch, progressCh)
		}

		go func() {
			wg.Wait()
			close(ch)
			close(progressCh)
		}()
		if _, err := p.Run(); err != nil {
			fmt.Println("TUI error:", err)
		}

		languageData := make(map[string][]pkg.CodeSnippets)
		for result := range ch {
			languageData[result.Language] = result.Data
		}
	},
}

func init() {
	trainCmd.Flags().Int32VarP(&count, "count", "c", 1000, "Number of training iterations")
}
