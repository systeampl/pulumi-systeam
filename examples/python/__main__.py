import pulumi
import pulumi_systeam as systeam

# Provider configured via env vars: SYSTEAM_API_URL, SYSTEAM_API_TOKEN

# Look up organization
org = systeam.get_organization()

# Create a project
project = systeam.Project("my-project",
    name="Production Services",
    organization_id=org.id,
)

# Create a check
check = systeam.Check("api-health",
    name="API Health Check",
    project_id=project.id,
    check_type="http",
    target="https://api.example.com/health",
    interval=60,
    timeout=10,
)

# Create notification channel
slack_channel = systeam.NotificationChannel("slack-alerts",
    name="Slack Alerts",
    organization_id=org.id,
    channel_type="slack",
    config={
        "webhook_url": "https://hooks.slack.com/services/xxx",
    },
)

# Create SLO
slo = systeam.CheckSlo("api-slo",
    check_id=check.id,
    target_uptime=99.9,
    period_days=30,
)

# Export check UUID for ping URL
pulumi.export("check_uuid", check.uuid)
pulumi.export("project_id", project.id)
