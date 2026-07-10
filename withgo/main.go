package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Define the root command
	rootCmd := &cobra.Command{
		Use:   "httpcli",
		Short: "httpcli is a minimal HTTP client cli",
	}

	// Define the get subcommand
	getCmd := &cobra.Command{
		Use:   "get [url]",
		Short: "fetch content from a url",
		Args:  cobra.ExactArgs(1), // requires exactly one argument
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			// perform the http get requires
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error making request %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			// read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading response %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("status: %s\n", resp.Status)
			fmt.Println(string(body))
		},
	}

	// add get command to root command and execute
	rootCmd.AddCommand(getCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
