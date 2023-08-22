package command

import (
	"github.com/spf13/cobra"
	"wallet"
)

func NewMigration(core wallet.App) *cobra.Command {
	return &cobra.Command{
		Use:   "migration",
		Short: "Run Migration",
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "up" {
				err := core.Migration().Up()
				if err != nil {
					core.ZapLogger().Error(err.Error())
				}
			}
		},
	}
}
