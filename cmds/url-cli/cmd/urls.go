package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"time"
	processor "url-cli/pkg/urls-processor"
)

var urlBodySizeListCmd = &cobra.Command{
	Use:   "list -u <urls>",
	Short: "list the urls sorted by the size of the response body",
	RunE: func(cmd *cobra.Command, args []string) error {
		urls, err := cmd.Flags().GetStringSlice("urls")
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()

		pairs := processor.VisitUrls(ctx, urls)

		b, err := json.MarshalIndent(pairs, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(urlBodySizeListCmd)
}
