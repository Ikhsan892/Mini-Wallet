package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"wallet"
	"wallet/cmd/command"
)

func main() {
	//http://patorjk.com/software/taag/#p=display&f=Small%20Slant&t=Echo
	fmt.Println(`	
		  
 /$$      /$$           /$$ /$$             /$$    
| $$  /$ | $$          | $$| $$            | $$    
| $$ /$$$| $$  /$$$$$$ | $$| $$  /$$$$$$  /$$$$$$  
| $$/$$ $$ $$ |____  $$| $$| $$ /$$__  $$|_  $$_/  
| $$$$_  $$$$  /$$$$$$$| $$| $$| $$$$$$$$  | $$    
| $$$/ \  $$$ /$$__  $$| $$| $$| $$_____/  | $$ /$$
| $$/   \  $$|  $$$$$$$| $$| $$|  $$$$$$$  |  $$$$/
|__/     \__/ \_______/|__/|__/ \_______/   \___/  
                                                   

		Julo Programming Test
	
	
	`)

	app := wallet.New()
	app.Start()

	command.RootCmd.AddCommand(command.NewServer(app))
	command.RootCmd.AddCommand(command.NewMigration(app))
	command.RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	err := command.RootCmd.Execute()
	if err != nil {
		app.ZapLogger().Error(err.Error())
	}

	app.Shutdown()

}
