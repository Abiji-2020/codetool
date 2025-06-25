package cmd

import (
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
	},
}

func init() {
	trainCmd.Flags().Int32VarP(&count, "count", "c", 1000, "Number of training iterations")
}
