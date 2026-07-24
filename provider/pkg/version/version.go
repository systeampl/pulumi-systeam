package version

// Version is the semver of the provider, injected at build time via -ldflags.
// Defaults to a dev value for local builds.
var Version string = "0.0.0-dev"
