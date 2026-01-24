package main

import (
	"fmt"
	"os"

	"github.com/Ivan-Martins-DevProjects/PayHub/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	var PayHubRoot = &cobra.Command{
		Use:   "payhub",
		Short: "Gateway de Pagamentos",
		Long:  "Todos as suas formas de pagamento concentradas em um único lugar",
	}

	var PayHubInit = &cobra.Command{
		Use:   "init",
		Short: "Inicialize sua configuração a partir do seus arquivos YAML",
		Run: func(cmd *cobra.Command, args []string) {
			err := commands.FuncInit()
			if err != nil {
				fmt.Printf("Erro com a aplicação: %v", err)
				return
			}

			fmt.Printf("Criação dos arquivos de configuração concluída:")
		},
	}

	PayHubRoot.AddCommand(PayHubInit)

	if err := PayHubRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
