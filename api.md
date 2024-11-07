# Shared Response Types

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Balance">Balance</a>
- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#BidList">BidList</a>
- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Session">Session</a>

# Blockchains

Methods:

- <code title="post /blockchain/send/eth">client.Blockchains.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainService.SendEth">SendEth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSendEthParams">BlockchainSendEthParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/send/mor">client.Blockchains.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainService.SendMor">SendMor</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSendMorParams">BlockchainSendMorParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Models

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Model">Model</a>

Methods:

- <code title="post /blockchain/models">client.Blockchains.Models.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelNewParams">BlockchainModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/models">client.Blockchains.Models.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/models/{id}">client.Blockchains.Models.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Sessions

Methods:

- <code title="post /blockchain/models/{id}/session">client.Blockchains.Models.Sessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelSessionNewParams">BlockchainModelSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Bids

Methods:

- <code title="get /blockchain/models/{id}/bids">client.Blockchains.Models.Bids.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelBidService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelBidListParams">BlockchainModelBidListParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Active

Methods:

- <code title="get /blockchain/models/{id}/bids/active">client.Blockchains.Models.Bids.Active.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelBidActiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Rated

Methods:

- <code title="get /blockchain/models/{id}/bids/rated">client.Blockchains.Models.Bids.Rated.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelBidRatedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Stats

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ModelStats">ModelStats</a>

Methods:

- <code title="get /blockchain/models/{id}/stats">client.Blockchains.Models.Stats.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainModelStatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ModelStats">ModelStats</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BlockchainBids

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Bid">Bid</a>

Methods:

- <code title="post /blockchain/bids">client.BlockchainBids.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBidService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBidNewParams">BlockchainBidNewParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/bids/{id}">client.BlockchainBids.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBidService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Bid">Bid</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/bids/{id}">client.BlockchainBids.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBidService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /blockchain/bids/{id}/session">client.BlockchainBids.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBidService.Session">Session</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBidSessionParams">BlockchainBidSessionParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BlockchainSessions

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Budget">Budget</a>
- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#SessionList">SessionList</a>

Methods:

- <code title="post /blockchain/sessions">client.BlockchainSessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionNewParams">BlockchainSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/sessions/{id}">client.BlockchainSessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/sessions/budget">client.BlockchainSessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionService.Budget">Budget</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /blockchain/sessions/{id}/close">client.BlockchainSessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /blockchain/sessions/provider">client.BlockchainSessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionService.Provider">Provider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionProviderParams">BlockchainSessionProviderParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#SessionList">SessionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/sessions/user">client.BlockchainSessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionService.User">User</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainSessionUserParams">BlockchainSessionUserParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#SessionList">SessionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Blockchain

Methods:

- <code title="post /blockchain/approve">client.Blockchain.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainApproveParams">BlockchainApproveParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Provider">Provider</a>

Methods:

- <code title="post /blockchain/providers">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderNewParams">BlockchainProviderNewParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Provider">Provider</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /blockchain/providers">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Provider">Provider</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /blockchain/providers/{id}">client.Blockchain.Providers.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Bids

Methods:

- <code title="get /blockchain/providers/{id}/bids">client.Blockchain.Providers.Bids.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderBidService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderBidListParams">BlockchainProviderBidListParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Active

Methods:

- <code title="get /blockchain/providers/{id}/bids/active">client.Blockchain.Providers.Bids.Active.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainProviderBidActiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#BidList">BidList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Balance

Methods:

- <code title="get /blockchain/balance">client.Blockchain.Balance.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainBalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Allowance

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Allowance">Allowance</a>

Methods:

- <code title="get /blockchain/allowance">client.Blockchain.Allowance.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainAllowanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainAllowanceGetParams">BlockchainAllowanceGetParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#Allowance">Allowance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LatestBlock

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#LatestBlock">LatestBlock</a>

Methods:

- <code title="get /blockchain/latestBlock">client.Blockchain.LatestBlock.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainLatestBlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#LatestBlock">LatestBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Token

### Supply

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#TokenSupply">TokenSupply</a>

Methods:

- <code title="get /blockchain/token/supply">client.Blockchain.Token.Supply.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainTokenSupplyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#TokenSupply">TokenSupply</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#TransactionList">TransactionList</a>

Methods:

- <code title="get /blockchain/transactions">client.Blockchain.Transactions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#BlockchainTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#TransactionList">TransactionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Proxy

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ClaimableBalance">ClaimableBalance</a>

Methods:

- <code title="post /proxy/sessions/initiate">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ProxySessionService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ProxySessionInitiateParams">ProxySessionInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go/shared#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /proxy/sessions/{id}/providerClaim">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ProxySessionService.ProviderClaim">ProviderClaim</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ProxySessionProviderClaimParams">ProxySessionProviderClaimParams</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ClaimableBalance">ClaimableBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /proxy/sessions/{id}/providerClaimableBalance">client.Proxy.Sessions.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ProxySessionService.ProviderClaimableBalance">ProviderClaimableBalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go">morpheusmarketplace</a>.<a href="https://pkg.go.dev/github.com/srt0422/morpheus-marketplace-go#ClaimableBalance">ClaimableBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
