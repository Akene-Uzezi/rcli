package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Define local variables to hold the flag values
	var url string
	var method string
	var verbose bool

	// define the root command
	rootCmd := &cobra.Command{
		Use:   "httpcli",
		Short: "A simple http utility CLI tool",
		Long:  `An HTTP client CLI built using the spf13/cobra framework to parse flags cleanly`,
		// The run function executes when the command is invoked successfully
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("--- Parsed Go (Cobra) Flags")
			fmt.Printf("URL: %s\n", url)
			fmt.Printf("Method: %s\n", method)
			fmt.Printf("Verbose: %t\n", verbose)
		},
	}

	// Bind Flags to variables
	// StringVarP takes: (pointer)
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "The target url for the request")
	rootCmd.Flags().StringVarP(&method, "method", "m", "GET", "the http method to use")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "turn verbose on")

	// tell cobra that the url flag is mandatory
	rootCmd.MarkFlagRequired("url")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
