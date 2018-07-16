# About

This is a toy password manager that uses [Azure Key Vault](https://docs.microsoft.com/en-us/azure/key-vault/) as the backend to store your secrets/passwords. It supports the basic CRUD features.


# Install

`go get github.com/esell/kv-pass-manager` :)


# Usage

You first need to set four environment variables before using the app:


`AZURE_TENANT_ID`: Your Azure tenant ID

`AZURE_CLIENT_ID`: Your Azure client ID. This will be an app ID from your AAD.

`AZURE_CLIENT_SECRET`: The secret for the client ID above.

`KVAULT`: The name of your vault (just the name, not the full URL/path)



List the secrets currently in the vault (not the values though):
`kv-pass`

Get the value for a secret in the vault:
`kv-pass YOUR_SECRETS_NAME`

Add or Update a secret in the vault:
`kv-pass -edit YOUR_NEW_VALUE YOUR_SECRETS_NAME`

Delete a secret in the vault:
`kv-pass -delete YOUR_SECRETS_NAME`


# Hacking

If you want to hack on this, you'll need to clone/fork the repo and then use [dep](https://github.com/golang/dep) to install the dependencies.