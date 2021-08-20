package main

import (
	"github.com/spf13/cobra"
	"log"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Copy(args[0], args[1]); err != nil {
			log.Fatalf("error copying from %s to %s: %s\n", args[0], args[1], err)
		}
	},
}

func Copy(src string, dst string) error {
	return nil
}
