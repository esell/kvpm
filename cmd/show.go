package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/esell/kvpm/util"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the value of an entry",

	Run: func(cmd *cobra.Command, args []string) {

		vaultName := os.Getenv("KVAULT")

		basicClient, err := util.GetBasicClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		secretResp, err := basicClient.GetSecret(context.Background(), "https://"+vaultName+".vault.azure.net", args[0], "")
		if err != nil {
			fmt.Printf("unable to get value for secret: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(*secretResp.Value)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
