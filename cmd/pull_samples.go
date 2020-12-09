package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func NewPullSamplesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pull-samples",
		Short: "Pull a new set of samples from Github.",
		Long:  `Pull a new set of samples from Github.`,
		Run: func(cmd *cobra.Command, args []string) {
			pullSamples()
		},
	}
}

func pullSamples() {
	log.Println("Pulling samples from Github ...")
}
