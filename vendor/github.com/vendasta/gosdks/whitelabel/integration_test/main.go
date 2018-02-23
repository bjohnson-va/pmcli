package main

import (
	"log"

	"context"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/whitelabel"
)

func main() {

	client := whitelabel.BuildWhiteLabelClient("VCONFIG", "UcEsxWXatyq9vW8dtfxY", config.Prod)

	d, err := client.Get(context.Background(), "VUNI", "")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("DATA %#v", d)
}
