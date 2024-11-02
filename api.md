# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Allowance">Allowance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Balance">Balance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#BidList">BidList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#LatestBlock">LatestBlock</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Provider">Provider</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#SessionList">SessionList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#TokenSupply">TokenSupply</a>

# Blockchain

Methods:

- <code title="post /blockchain/approve">client.Blockchain.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainApproveParams">BlockchainApproveParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Eth

Methods:

- <code title="post /blockchain/send/eth">client.Blockchain.Eth.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainEthService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainEthSendParams">BlockchainEthSendParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Mor

Methods:

- <code title="post /blockchain/send/mor">client.Blockchain.Mor.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainMorService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainMorSendParams">BlockchainMorSendParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Stats">Stats</a>

Methods:

- <code title="post /blockchain/models">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelNewParams">BlockchainModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/models/{id}">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /blockchain/models/{id}/session">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.Session">Session</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelSessionParams">BlockchainModelSessionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bids

Methods:

- <code title="get /blockchain/models/{id}/bids">client.Blockchain.Models.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidListParams">BlockchainModelBidListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models/{id}/bids/active">client.Blockchain.Models.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidService.Active">Active</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models/{id}/bids/rated">client.Blockchain.Models.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidService.Rated">Rated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Stats

Methods:

- <code title="get /blockchain/models/{id}/stats">client.Blockchain.Models.Stats.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelStatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Stats">Stats</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bids

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Bid">Bid</a>

Methods:

- <code title="post /blockchain/bids">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidNewParams">BlockchainBidNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/bids/{id}">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/bids/{id}">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /blockchain/bids/{id}/session">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.Session">Session</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidSessionParams">BlockchainBidSessionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sessions

Methods:

- <code title="post /blockchain/sessions">client.Blockchain.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionNewParams">BlockchainSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/sessions/{id}">client.Blockchain.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/sessions/{id}/close">client.Blockchain.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Budget

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Budget">Budget</a>

Methods:

- <code title="get /blockchain/sessions/budget">client.Blockchain.Sessions.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionBudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### User

Methods:

- <code title="get /blockchain/sessions/user">client.Blockchain.Sessions.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionUserListParams">BlockchainSessionUserListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#SessionList">SessionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Provider

Methods:

- <code title="get /blockchain/sessions/provider">client.Blockchain.Sessions.Provider.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionProviderListParams">BlockchainSessionProviderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#SessionList">SessionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Providers

Methods:

- <code title="post /blockchain/providers">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderNewParams">BlockchainProviderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Provider">Provider</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/providers">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Provider">Provider</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/providers/{id}">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Bids

Methods:

- <code title="get /blockchain/providers/{id}/bids">client.Blockchain.Providers.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidListParams">BlockchainProviderBidListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/providers/{id}/bids/active">client.Blockchain.Providers.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidService.Active">Active</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Balance

Methods:

- <code title="get /blockchain/balance">client.Blockchain.Balance.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Allowance

Methods:

- <code title="get /blockchain/allowance">client.Blockchain.Allowance.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceGetParams">BlockchainAllowanceGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Allowance">Allowance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LatestBlock

Methods:

- <code title="get /blockchain/latestBlock">client.Blockchain.LatestBlock.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainLatestBlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#LatestBlock">LatestBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Token

### Supply

Methods:

- <code title="get /blockchain/token/supply">client.Blockchain.Token.Supply.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainTokenSupplyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#TokenSupply">TokenSupply</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#TransactionList">TransactionList</a>

Methods:

- <code title="get /blockchain/transactions">client.Blockchain.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#TransactionList">TransactionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Proxy

## Sessions

Methods:

- <code title="post /proxy/sessions/initiate">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionInitiateParams">ProxySessionInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /proxy/sessions/{id}/providerClaim">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionService.ProviderClaim">ProviderClaim</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimParams">ProxySessionProviderClaimParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ClaimableBalance">ClaimableBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ProviderClaimableBalance

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ClaimableBalance">ClaimableBalance</a>

Methods:

- <code title="get /proxy/sessions/{id}/providerClaimableBalance">client.Proxy.Sessions.ProviderClaimableBalance.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimableBalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ClaimableBalance">ClaimableBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
