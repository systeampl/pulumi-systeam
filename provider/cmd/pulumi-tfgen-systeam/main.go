package main

import (
	"github.com/pawel-cygal/pulumi-systeam/provider"
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"
)

func main() {
	// Plugin Framework provider → use the pf tfgen entry point (not the runtime
	// tfbridge.Main, which lacks a real provider for `schema` and nil-panics).
	tfgen.Main("systeam", provider.Provider())
}
