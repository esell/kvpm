package util

import (
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
)

func GetBasicClient() (keyvault.BaseClient, error) {
	authorizer, err := kvauth.NewAuthorizerFromEnvironment()
	if err != nil {
		return keyvault.New(), err
	}

	basicClient := keyvault.New()
	basicClient.Authorizer = authorizer

	return basicClient, nil
}
