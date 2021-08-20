package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var bollCmd = &cobra.Command{
		Use:   "boll",
		Short: "Boll is a filesystem utility tool, aimed to work with different fs types.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	bollCmd.AddCommand(copyCmd)
	bollCmd.AddCommand(normalizeCmd)

	if err := bollCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
