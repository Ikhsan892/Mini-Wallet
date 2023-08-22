package command

import (
	"github.com/spf13/cobra"
	"wallet"
	"wallet/adapter"
	"wallet/adapter/web"
)

func NewServer(core wallet.App) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start API server",
		Run: func(cmd *cobra.Command, args []string) {
			adapter.RunAdapter(web.NewEcho(core))
		},
	}
}
