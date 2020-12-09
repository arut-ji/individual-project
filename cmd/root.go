package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "latom",
		Short: "Latom is a CLI application aimed to aid Kubernetes code smells research.",
		Long: `Latom is a CLI application aimed to aid Kubernetes code smells research.
This CLI app provides several commands to pull Kubernetes files from Github, 
datasource inspection, and statically analyze the scripts.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Latom ...")
		},
	}
}

func Execute() {
	ctx := context.Background()

	rootCmd := newRootCmd()
	rootCmd.AddCommand(NewPullSamplesCmd())
	rootCmd.AddCommand(NewPlaygroundCmd(ctx))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
