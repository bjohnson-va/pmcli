# Group golang SDK
- [Installation](#installation)
- [Usage](#usage)
- [Contact](#contact)

## Installation

* Using dep? `dep ensure -update github.com/vendasta/gosdks`

## Usage

To initialize a client:

```golang
client, err := group.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
if err != nil {
    logging.Criticalf(ctx, "Failed to initialize group client because %s", err.Error())
    os.Exit(-1)
}
```

Example ListMembers call:

```golang
members, totalMembers, nextCursor, hasMore, err = client.ListMembers(ctx, Path{Nodes: []string{"G-123", "G-456"}}, PageSize(10), Cursor("AOjdu8a"))
```

## Contact

Teamy McTeamface

