package command

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Run Webserver",
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	// no need to provide the default cobra completion command
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
