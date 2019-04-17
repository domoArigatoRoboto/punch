package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release [subdomain]",
	Short: "Release subdomain",
	Long:  `Release a subdomain you have reserved`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subdomain = args[0]
		release(subdomain)
	},
}

func init() {
	subdomainCmd.AddCommand(releaseCmd)
}

func release(subdomain string) {
	if !checkSubdomain(subdomain) {
		fmt.Fprintf(os.Stderr, "Invalid Subdomain\n")
		os.Exit(1)
	}
	err := restAPI.ReleaseSubdomainAPI(subdomain)
	if err != nil {
		printError(err)
		os.Exit(1)
	}
	fmt.Println("Successfully released subdomain")
}
