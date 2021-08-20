package main

import (
	"github.com/spf13/cobra"
	"log"
)

// todo file system type should be given: NTFS, FAT32, etc
var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "Normalize file name for windows",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := Normalize(args[0]); err != nil {
			log.Fatalf("error normalizing file name (%s): %s\n", args[0], err)
		}
	},
}

func Normalize(name string) (string, error) {
	return name, nil
}
