package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var headers map[string]any

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

	// define variable to capture the flag values
	var headers []string
	var data string

	postCmd := &cobra.Command{
		Use:   "post [url]",
		Short: "make a post command",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]

			reqBody := strings.NewReader(data)

			req, err := http.NewRequest("POST", url, reqBody)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to initialize request: %v\n", err)
				os.Exit(1)
			}

			for _, h := range headers {
				parts := strings.SplitN(h, ":", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					req.Header.Add(key, value)
				}
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error executing post request: %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to read post response: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("status: %s\n", resp.Status)
			fmt.Println(string(body))
		},
	}

	postCmd.Flags().StringSliceVarP(&headers, "headers", "H", []string{}, "HTTP headers to pass")
	postCmd.Flags().StringVarP(&data, "data", "d", "", "the raw string data payload for the POST body")

	// add get command to root command and execute
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(postCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
