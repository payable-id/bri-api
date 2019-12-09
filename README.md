# BRI Sangu

## Usage blueprint

1. There is a type named `Client` (`bri.Client`) that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used. The gateway type need a Client instance.
3. Any activity (token request) is done in the gateway level.

## Example

```go
    briClient := bri.NewClient()
    briClient.BaseUrl = "BRI_BASE_URL"
    briClient.ClientId = "BRI_CLIENT_ID"
    briClient.ClientSecret = "BRI_CLIENT_SECRET"

    coreGateway := bri.CoreGateway{
        Client: briClient,
    }

    res, _ := coreGateway.GetToken()
```