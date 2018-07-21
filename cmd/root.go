package cmd

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/esell/kvpm/util"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kvpm",
	Short: "A toy password manager backed by Azure Key Vault",
	Run: func(cmd *cobra.Command, args []string) {

		vaultName := os.Getenv("KVAULT")

		basicClient, err := util.GetBasicClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		secretList, err := basicClient.GetSecrets(context.Background(), "https://"+vaultName+".vault.azure.net", nil)
		if err != nil {
			fmt.Printf("unable to get list of secrets: %v\n", err)
			os.Exit(1)
		}

		// group by ContentType
		secWithType := make(map[string][]string)
		secWithoutType := make([]string, 1)
		for _, secret := range secretList.Values() {
			if secret.ContentType != nil {
				_, exists := secWithType[*secret.ContentType]
				if exists {
					secWithType[*secret.ContentType] = append(secWithType[*secret.ContentType], path.Base(*secret.ID))
				} else {
					tempSlice := make([]string, 1)
					tempSlice[0] = path.Base(*secret.ID)
					secWithType[*secret.ContentType] = tempSlice
				}
			} else {
				secWithoutType = append(secWithoutType, path.Base(*secret.ID))
			}
		}
		showIndent := "   "
		for k, v := range secWithType {
			// split
			splitOut := strings.Split(k, "/")
			if len(splitOut) > 1 {
				fmt.Printf(" |-- ")
				for i := 0; i < len(splitOut); i++ {
					if i == 0 {
						fmt.Println(splitOut[i])
					} else {
						fmt.Println(" |" + buildIndent(i) + "|-- " + splitOut[i])
					}
				}
				for _, sec := range v {
					fmt.Println(" |" + buildIndent(len(splitOut)) + " |-- " + sec)
				}
			} else {
				fmt.Println(" |-- " + k)
				for _, sec := range v {
					fmt.Println(" |" + showIndent + "|-- " + sec)
				}
			}
		}
		for _, wov := range secWithoutType {
			fmt.Println(" |-- " + wov)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func buildIndent(level int) string {
	var spaces strings.Builder
	for i := 0; i < level; i++ {
		fmt.Fprintf(&spaces, "%s", "   ")
	}
	return spaces.String()
}
