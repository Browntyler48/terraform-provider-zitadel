package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/zitadel/terraform-provider-zitadel/v2/zitadel"
)

func main() {
	ctx := context.Background()
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/zitadel/zitadel",
		Debug:   debug,
	}

	if err := providerserver.Serve(ctx, zitadel.NewProviderPV6, opts); err != nil {
		log.Fatal(err)
	}
}
