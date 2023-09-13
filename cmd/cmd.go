package cmd

import (
	"fmt"
	"os"

	api "github.com/BearTS/go-echo-template/api/cmd"
	mailer "github.com/BearTS/go-echo-template/mailer/cmd"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

// Execute - starts the CLI
func init() {
	cmd.AddCommand(api.RootCmd())
	cmd.AddCommand(mailer.RootCmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
