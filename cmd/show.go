package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/esell/kvpm/util"
	"github.com/spf13/cobra"
)

var UseLocal bool

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the value of an entry",

	Run: func(cmd *cobra.Command, args []string) {
		if UseLocal {
			// get home dir first
			currentUser, err := user.Current()
			if err != nil {
				fmt.Println(err)
			}
			if currentUser.HomeDir == "" {
				fmt.Println("no home directory set, exiting")
				os.Exit(1)
			}
			// prompt for pw
			pass, err := util.ReadPass()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// darn tootin
			localCache, err := ioutil.ReadFile(filepath.Join(currentUser.HomeDir, "kvpm.json"))
			if err != nil {
				fmt.Println(err)
			}
			var localPassList util.PassList
			err = json.Unmarshal(localCache, &localPassList)
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range localPassList.PassEntries {
				if v.PassName == args[0] {
					decryptValue, err := util.DecryptString(v.PassValue, pass)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(decryptValue)
				}
			}
		} else {
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
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&UseLocal, "local", "l", false, "Read from local cache")
}
