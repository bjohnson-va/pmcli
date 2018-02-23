# Tiny ID

Tiny ID provides a safe library that generates short human friendly identifiers at the expense of an RPC and
the cost of a vStore transaction. (5ms read + 5ms write = 10ms RTT @ 50th percentile).  Please use sparingly
and consider using UUIDs which are less expensive.  If you are expecting to have a high number of entities, it is
strongly recommended to not use this library for your IDs.

# Usage

```golang
// create a vstore client
vstoreAPI, err := vstore.New()
if err != nil {
    log.Fatalf("Error initializing vStore client. %s", err.Error())
}
// you should share a namespace with your other entities
yourAppsVStoreNamespace := "account-group"
yourAppsServiceAccount := []string{"account-group-prod@repcore-prod.iam.gserviceaccount.com"}
tinyIDPrefix := "AG"

gen, err := tinyid.NewVStoreGenerator(context.Background(), vstoreAPI, yourAppsVStoreNamespace, yourAppsServiceAccount, tinyIDPrefix, Length(6))
if err != nil {
    log.Fatalf("Error initializing tiny id generator. %s", err.Error())
}

tinyID, err := gen.GenerateTinyID()
if err != nil {
    log.Fatalf("Error generating tiny id. %s", err.Error())
}
log.Printf("Tiny ID %s", tinyID) // will output `AG-PFRTXQ` eventually...

```
