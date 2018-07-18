package main

import (
	"fmt"
	"os"

	"github.com/esell/kvpm/cmd"
)

func main() {
	if os.Getenv("AZURE_TENANT_ID") == "" || os.Getenv("AZURE_CLIENT_ID") == "" || os.Getenv("AZURE_CLIENT_SECRET") == "" || os.Getenv("KVAULT") == "" {
		fmt.Println("env vars not set, exiting...")
		os.Exit(1)
	}

	cmd.Execute()
}
