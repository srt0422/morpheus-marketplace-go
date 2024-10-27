# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Bid">Bid</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>

# Blockchain

## Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelListResponse">BlockchainModelListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelDeleteResponse">BlockchainModelDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelExistsResponse">BlockchainModelExistsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelResetstatsResponse">BlockchainModelResetstatsResponse</a>

Methods:

- <code title="post /blockchain/models">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelNewParams">BlockchainModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelListResponse">BlockchainModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/models/{id}">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelDeleteResponse">BlockchainModelDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models/{id}/exists">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.Exists">Exists</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelExistsResponse">BlockchainModelExistsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/models/{id}/resetstats">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.Resetstats">Resetstats</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelResetstatsResponse">BlockchainModelResetstatsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/models/{id}/session">client.Blockchain.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelService.Session">Session</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelSessionParams">BlockchainModelSessionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bids

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidListResponse">BlockchainModelBidListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidActiveResponse">BlockchainModelBidActiveResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidRatedResponse">BlockchainModelBidRatedResponse</a>

Methods:

- <code title="get /blockchain/models/{id}/bids">client.Blockchain.Models.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidListParams">BlockchainModelBidListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models/{id}/bids/active">client.Blockchain.Models.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidService.Active">Active</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models/{id}/bids/rated">client.Blockchain.Models.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelBidService.Rated">Rated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Minstake

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#MinStake">MinStake</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelMinstakeSetResponse">BlockchainModelMinstakeSetResponse</a>

Methods:

- <code title="get /blockchain/models/minstake">client.Blockchain.Models.Minstake.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelMinstakeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#MinStake">MinStake</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/models/minstake">client.Blockchain.Models.Minstake.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelMinstakeService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelMinstakeSetParams">BlockchainModelMinstakeSetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelMinstakeSetResponse">BlockchainModelMinstakeSetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Stats

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ModelStats">ModelStats</a>

Methods:

- <code title="get /blockchain/models/{id}/stats">client.Blockchain.Models.Stats.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainModelStatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ModelStats">ModelStats</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bids

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidNewResponse">BlockchainBidNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidGetResponse">BlockchainBidGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidDeleteResponse">BlockchainBidDeleteResponse</a>

Methods:

- <code title="post /blockchain/bids">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidNewParams">BlockchainBidNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidNewResponse">BlockchainBidNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/bids/{id}">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidGetResponse">BlockchainBidGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/bids/{id}">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidDeleteResponse">BlockchainBidDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/bids/{id}/session">client.Blockchain.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidService.Session">Session</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBidSessionParams">BlockchainBidSessionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionGetResponse">BlockchainSessionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionCloseResponse">BlockchainSessionCloseResponse</a>

Methods:

- <code title="post /blockchain/sessions">client.Blockchain.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionNewParams">BlockchainSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/sessions/{id}">client.Blockchain.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionGetResponse">BlockchainSessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/sessions/{id}/close">client.Blockchain.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionCloseResponse">BlockchainSessionCloseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### User

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionUserListResponse">BlockchainSessionUserListResponse</a>

Methods:

- <code title="get /blockchain/sessions/user">client.Blockchain.Sessions.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionUserListParams">BlockchainSessionUserListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionUserListResponse">BlockchainSessionUserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Provider

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionProviderListResponse">BlockchainSessionProviderListResponse</a>

Methods:

- <code title="get /blockchain/sessions/provider">client.Blockchain.Sessions.Provider.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionProviderListParams">BlockchainSessionProviderListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionProviderListResponse">BlockchainSessionProviderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Budget

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionBudgetListResponse">BlockchainSessionBudgetListResponse</a>

Methods:

- <code title="get /blockchain/sessions/budget">client.Blockchain.Sessions.Budget.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionBudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSessionBudgetListResponse">BlockchainSessionBudgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Provider">Provider</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderListResponse">BlockchainProviderListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderDeleteResponse">BlockchainProviderDeleteResponse</a>

Methods:

- <code title="post /blockchain/providers">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderNewParams">BlockchainProviderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Provider">Provider</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/providers">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderListResponse">BlockchainProviderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/providers/{id}">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderDeleteResponse">BlockchainProviderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bids

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidListResponse">BlockchainProviderBidListResponse</a>

Methods:

- <code title="get /blockchain/providers/{id}/bids">client.Blockchain.Providers.Bids.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidListParams">BlockchainProviderBidListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Active

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidActiveListResponse">BlockchainProviderBidActiveListResponse</a>

Methods:

- <code title="get /blockchain/providers/{id}/bids/active">client.Blockchain.Providers.Bids.Active.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainProviderBidActiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Balance

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Balance">Balance</a>

Methods:

- <code title="get /blockchain/balance">client.Blockchain.Balance.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainBalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Allowance

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Allowance">Allowance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceApproveResponse">BlockchainAllowanceApproveResponse</a>

Methods:

- <code title="get /blockchain/allowance">client.Blockchain.Allowance.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceGetParams">BlockchainAllowanceGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#Allowance">Allowance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/allowance">client.Blockchain.Allowance.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceApproveParams">BlockchainAllowanceApproveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainAllowanceApproveResponse">BlockchainAllowanceApproveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Send

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendEthResponse">BlockchainSendEthResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendMorResponse">BlockchainSendMorResponse</a>

Methods:

- <code title="post /blockchain/send/eth">client.Blockchain.Send.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendService.Eth">Eth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendEthParams">BlockchainSendEthParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendEthResponse">BlockchainSendEthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/send/mor">client.Blockchain.Send.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendService.Mor">Mor</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendMorParams">BlockchainSendMorParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainSendMorResponse">BlockchainSendMorResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LatestBlock

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#LatestBlock">LatestBlock</a>

Methods:

- <code title="get /blockchain/latestBlock">client.Blockchain.LatestBlock.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainLatestBlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#LatestBlock">LatestBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Token

### Supply

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#TokenSupply">TokenSupply</a>

Methods:

- <code title="get /blockchain/token/supply">client.Blockchain.Token.Supply.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainTokenSupplyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#TokenSupply">TokenSupply</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#TransactionList">TransactionList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainTransactionListResponse">BlockchainTransactionListResponse</a>

Methods:

- <code title="get /blockchain/transactions">client.Blockchain.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#BlockchainTransactionListParams">BlockchainTransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#TransactionList">TransactionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Proxy

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimResponse">ProxySessionProviderClaimResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimableBalanceResponse">ProxySessionProviderClaimableBalanceResponse</a>

Methods:

- <code title="post /proxy/sessions/initiate">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionInitiateParams">ProxySessionInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /proxy/sessions/{id}/providerClaim">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionService.ProviderClaim">ProviderClaim</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimParams">ProxySessionProviderClaimParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimResponse">ProxySessionProviderClaimResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /proxy/sessions/{id}/providerClaimableBalance">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionService.ProviderClaimableBalance">ProviderClaimableBalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/morpheus-marketplace-go#ProxySessionProviderClaimableBalanceResponse">ProxySessionProviderClaimableBalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
