package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/esell/kvpm/util"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a password",

	Run: func(cmd *cobra.Command, args []string) {

		vaultName := os.Getenv("KVAULT")

		basicClient, err := util.GetBasicClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = basicClient.DeleteSecret(context.Background(), "https://"+vaultName+".vault.azure.net", args[0])
		if err != nil {
			fmt.Printf("error deleting secret: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(args[0] + " deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
