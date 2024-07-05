package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cli-go",
		Short: "A simple CLI project",
		Long:  `A simple CLI project to serve as a starter template for Go CLI applications`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from cli-go!")
		},
	}

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var language string

var greetCmd = &cobra.Command{
	Use:   "greet [file]",
	Short: "Greets a person",
	Long:  `This subcommand greets a person whose name is in the specified file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name, err := readFile(args[0])
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		greeting := getGreeting(language)
		fmt.Printf("%s, %s!\n", greeting, name)
	},
}

func init() {
	greetCmd.Flags().StringVarP(&language, "language", "l", "en", "Language for the greeting (en, es, bg)")
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func getGreeting(lang string) string {
	switch lang {
	case "bg":
		return "Здравей"
	case "es":
		return "Hola"
	default:
		return "Hello"
	}
}
