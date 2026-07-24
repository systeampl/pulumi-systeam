# Pulumi provider for Syschecks (SysTeam)

Manage your **[Syschecks](https://syschecks.com)** (SysTeam Monitoring) stack as
code with [Pulumi](https://www.pulumi.com/): monitoring checks, projects, status
pages, on-call schedules, escalation policies, SLOs, integration keys, teams,
services, contact methods, lifecycle watches and multi-stage incident
**playbooks**.

Generated from the official
[Terraform provider](https://github.com/systeampl/terraform-provider-systeam)
via the Pulumi-Terraform Bridge — **full parity, identical behaviour** across
all 15 resources.

## Installation

```bash
pip install pulumi_systeam
```

## Configuration

Via environment variables:

```bash
export SYSTEAM_API_URL=https://syschecks.com    # base URL, no /api suffix
export SYSTEAM_API_TOKEN=pat_xxxxx              # a Personal Access Token
```

Or with Pulumi config:

```bash
pulumi config set systeam:apiUrl https://syschecks.com
pulumi config set --secret systeam:apiToken pat_xxxxx
```

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

## Resources

`Project`, `Check`, `CheckSlo`, `NotificationChannel`, `EscalationPolicy`,
`OncallSchedule`, `StatusPage`, `MaintenanceWindow`, `IntegrationKey`, `Team`,
`Service`, `Playbook`, `LifecycleWatch`, `ContactMethod`,
`AgentRegistrationToken`.

Data source: `get_organization()`.

## Links

- Homepage: <https://syschecks.com>
- Source & issues: <https://github.com/systeampl/pulumi-systeam>
- License: Apache-2.0
