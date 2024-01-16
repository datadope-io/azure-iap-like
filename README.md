# Create an Azure function to protect internal endpoints with Authentication

Clients will see an Authentication web against the identity provider configured.
Once logged, it will reach this app (`handler.go`) that will make the requests to the configured URL.

The Azure function OS should be windows.
The App Service Plan should be "Function premium" (to be able to connect to internal virtual networks).
The cheapest (EP1) is ~160€/month.

Modify `handler.go` to point to the internal endpoint.
Should be reacheable from the function subnetwork.

Compile the app:
```
GOOS=windows GOARCH=amd64 go build handler.go
```

Zip the files and upload to the function:
```
zip function.zip -r *
az functionapp deployment source config-zip --resource-group RG_NAME--name FUNCTION_NAME --src function.zip
```

Enable Authentication in the function.
