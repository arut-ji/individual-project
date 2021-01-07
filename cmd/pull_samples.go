package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
)

func NewPullSamplesCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "pull-samples",
		Short: "Pull a new set of samples from Github.",
		Long:  `Pull a new set of samples from Github.`,
		Run: func(cmd *cobra.Command, args []string) {
			pullSamples(ctx)
		},
	}
}

func pullSamples(ctx context.Context) {
	log.Println("Pulling samples from Github ...")
}
