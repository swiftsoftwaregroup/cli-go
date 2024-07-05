package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cli-golang",
		Short: "A simple CLI project",
		Long:  `A simple CLI project to serve as a starter template for Go CLI applications`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from cli-golang!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
