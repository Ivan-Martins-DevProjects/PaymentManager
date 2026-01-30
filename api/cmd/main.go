package main

import (
	"fmt"
	"os"

	payhubInit "github.com/Ivan-Martins-DevProjects/PayHub/internal/commands/init"
	"github.com/spf13/cobra"
)

func main() {
	var PayHubRoot = &cobra.Command{
		Use:   "payhub",
		Short: "Gateway de Pagamentos",
		Long:  "Todos as suas formas de pagamento concentradas em um único lugar",
	}

	PayHubRoot.AddCommand(payhubInit.PayHubInit)
	// PayHubRoot.AddCommand(payhubToken.PayHubToken)

	if err := PayHubRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
