package main

import (
	"context"
	_ "embed"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/systeampl/pulumi-systeam/provider"
)

// schema is the generated Pulumi package schema (produced by `make generate_schema`
// into this directory). The bridge requires it at runtime to serve the provider.
//
//go:embed schema.json
var schema []byte

func main() {
	meta := pfbridge.ProviderMetadata{PackageSchema: schema}
	pfbridge.Main(context.Background(), "systeam", provider.Provider(), meta)
}
