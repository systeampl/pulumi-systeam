# Publishing the Syschecks Pulumi provider

This document describes how `pulumi-systeam` (the Pulumi provider for
**Syschecks** / SysTeam Monitoring) is built, versioned and released. The whole
release is automated in [`.github/workflows/release.yml`](.github/workflows/release.yml)
and driven by a git tag.

## What gets published, and where

| Artifact | Destination | Produced by |
| --- | --- | --- |
| `pulumi-resource-systeam` plugin binaries (linux/darwin/windows × amd64/arm64) | **GitHub Release** of this repo | GoReleaser (`.goreleaser.yaml`) |
| `pulumi_systeam` Python SDK | **PyPI** | `twine upload` in the release workflow |
| Registry listing + API docs | **registry.pulumi.com** | one-time submission to `pulumi/registry` (see below) |

Pulumi resolves the plugin from this repo's GitHub Releases because
`GitHubOrg: "systeampl"` is set in [`provider/resources.go`](provider/resources.go).

## Dependency on the Terraform provider

This is a bridged provider. It wraps
[`systeampl/terraform-provider-systeam`](https://github.com/systeampl/terraform-provider-systeam)
via the Pulumi-Terraform Bridge, importing its `shim` package as a **normal,
tagged Go module** (see `go.mod`):

```
require github.com/systeampl/terraform-provider-systeam vX.Y.Z
```

There is **no** relative `replace` directive — the repo builds standalone. When
the Terraform provider changes, bump that `require` to the new tag and
regenerate (`make generate_sdk`).

## One-time setup

1. **Repository secret** (Settings → Secrets and variables → Actions):
   - `PYPI_API_TOKEN` — a PyPI API token (the token value; the workflow uses
     `__token__` as the username). Scope it to the `pulumi-systeam` project once
     it exists, or account-wide for the first publish.
   - `GITHUB_TOKEN` is provided automatically — no action needed.
2. **Pulumi Registry listing** (for public docs at pulumi.com/registry): open a
   PR against [`pulumi/registry`](https://github.com/pulumi/registry) adding a
   package reference for `systeampl/pulumi-systeam`. This is only needed once;
   subsequent releases are picked up automatically.

## Cutting a release

Versioning is `vX.Y.Z` (semver). The version is injected into the plugin binary
via ldflags (`provider/pkg/version.Version`) and stamped into the Python SDK's
`setup.py` at build time.

```bash
# 1. (if the Terraform provider changed) bump the require in go.mod, then:
make generate_sdk        # regenerate schema.json + sdk/python
go build ./...           # sanity check

# 2. commit, then tag and push — this triggers the release workflow:
git tag v0.1.0
git push origin v0.1.0
```

The workflow then:

1. builds the plugin binaries and attaches them to the GitHub Release, and
2. builds and uploads the `pulumi_systeam` Python SDK to PyPI.

## Local development

```bash
make build               # build plugin binary into bin/ (version = dev unless VERSION=… passed)
make generate_sdk        # regenerate schema + Python SDK
make install             # install the plugin into ~/.pulumi/plugins for local testing
make VERSION=0.1.0 build # stamp a specific version locally
```

## Consuming the provider

Once published:

```bash
pip install pulumi_systeam
```

```python
import pulumi_systeam as systeam
project = systeam.Project("production", name="Production", organization_id=1)
```

Configuration (env or `pulumi config`):

```bash
export SYSTEAM_API_URL=https://syschecks.com    # base URL, no /api suffix
export SYSTEAM_API_TOKEN=pat_xxxxx
```

> **Note:** even without this branded package, the underlying Terraform provider
> can be used from Pulumi today via
> [`pulumi package add terraform-provider systeampl/systeam`](https://www.pulumi.com/registry/packages/terraform-provider/)
> (the "Any Terraform Provider" bridge), because it is already live on the
> Terraform Registry. This branded SDK exists for a first-class
> `pip install pulumi_systeam` experience and registry docs.
