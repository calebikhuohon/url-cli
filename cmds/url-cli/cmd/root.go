package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:           "url-cli",
		Short:         "CLI that visits provided Urls and returns their response body size",
		SilenceErrors: false,
		SilenceUsage:  true,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	urlBodySizeListCmd.Flags().StringSliceP("urls", "u", nil, "-u <url1, url2, url3>,...")
	_ = urlBodySizeListCmd.MarkFlagRequired("urls")
}
