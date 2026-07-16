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
		Use:   "httpcli [url]",
		Short: "httpcli is a minimal HTTP client cli",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error making request: %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()
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
		Short: "make a post request",
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

	putCmd := &cobra.Command{
		Use:   "put [url]",
		Short: "make a put request",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]

			reqBody := strings.NewReader(data)

			req, err := http.NewRequest("PUT", url, reqBody)
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
				fmt.Fprintf(os.Stderr, "error executing put request: %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to read put response: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("status: %s\n", resp.Status)
			fmt.Println(string(body))
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "del [url]",
		Short: "make a delete request",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	postCmd.Flags().StringSliceVarP(&headers, "headers", "H", []string{}, "HTTP headers to pass")
	postCmd.Flags().StringVarP(&data, "data", "d", "", "the raw string data payload for the POST request")
	deleteCmd.Flags().StringSliceVarP(&headers, "headers", "H", []string{}, "HTTP headers to pass")
	deleteCmd.Flags().StringVarP(&data, "data", "d", "", "the raw string data payload for the DELETE request")
	putCmd.Flags().StringSliceVarP(&headers, "headers", "H", []string{}, "HTTP headers to pass")
	putCmd.Flags().StringVarP(&data, "data", "d", "", "the raw string data to be payload for the PUT request")
	// add get command to root command and execute
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(putCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
