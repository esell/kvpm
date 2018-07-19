package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/esell/kvpm/util"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Save all entries locally",

	Run: func(cmd *cobra.Command, args []string) {
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
		secList := make([]util.PassEntry, 1)
		for _, secret := range secretList.Values() {
			secretResp, err := basicClient.GetSecret(context.Background(), "https://"+vaultName+".vault.azure.net", path.Base(*secret.ID), "")
			if err != nil {
				fmt.Printf("unable to get value for secret: %v\n", err)
				os.Exit(1)
			}
			encryptValue, err := util.EncryptString(*secretResp.Value, pass)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tempPassEntry := util.PassEntry{PassName: path.Base(*secret.ID), PassValue: encryptValue}
			secList = append(secList, tempPassEntry)
		}
		blah := util.PassList{PassEntries: secList}
		b, err := json.Marshal(blah)
		if err != nil {
			fmt.Println(err)
		}

		// write to file
		f, err := os.Create(filepath.Join(currentUser.HomeDir, "kvpm.json"))
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write(b)
		if err != nil {
			fmt.Println(err)
		}
		f.Sync()
		f.Close()
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
