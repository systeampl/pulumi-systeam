package main

import (
	"context"

	"github.com/pawel-cygal/pulumi-systeam/provider"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
)

func main() {
	meta := pfbridge.ProviderMetadata{PackageSchema: nil}
	pfbridge.Main(context.Background(), "systeam", provider.Provider(), meta)
}
