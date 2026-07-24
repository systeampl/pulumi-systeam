# Pulumi Provider for Syschecks (SysTeam)

A [Pulumi](https://www.pulumi.com/) provider for **Syschecks** (SysTeam
Healthchecks), wrapping the official
[Terraform provider](https://github.com/systeampl/terraform-provider-systeam)
via the Pulumi-Terraform Bridge. **Full parity — all 15 resources**, identical
behaviour to Terraform.

## Installation

Install the Python SDK:

```bash
pip install pulumi_systeam
```

## Configuration

Via environment variables:

```bash
export SYSTEAM_API_URL=https://syschecks.com    # base URL (no /api suffix)
export SYSTEAM_API_TOKEN=pat_xxxxx              # a Personal Access Token
```

Or in Pulumi config:

```bash
pulumi config set systeam:apiUrl https://syschecks.com
pulumi config set --secret systeam:apiToken pat_xxxxx
```

## Resources

`Project`, `Check`, `CheckSlo`, `NotificationChannel`, `EscalationPolicy`,
`OncallSchedule`, `StatusPage`, `MaintenanceWindow`, `IntegrationKey`, `Team`,
`Service`, `Playbook`, `LifecycleWatch`, `ContactMethod`,
`AgentRegistrationToken`.

Data source: `get_organization()`.

## Example

```python
import pulumi_systeam as systeam

project = systeam.Project("production", name="Production", organization_id=1)

systeam.Check("api-health",
    name="API Health",
    type="uptime",
    project_id=project.id,
    url="https://api.example.com/healthz",
    interval=60)
```

See [examples/python/](examples/python/) for a complete example.

## How it works

This provider is generated from the Terraform provider's schema via the
Pulumi-Terraform Bridge, so every resource behaves exactly as its Terraform
counterpart. To regenerate after a Terraform provider change:

```bash
make generate_sdk
```

## License

Apache 2.0.
