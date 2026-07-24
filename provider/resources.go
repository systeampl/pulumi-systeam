package provider

import (
	_ "embed"

	"github.com/pawel-cygal/terraform-provider-systeam/shim"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

const (
	mainPkg = "systeam"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-systeam/bridge-metadata.json
var bridgeMetadata []byte

// All resources use the TF "id" (computed) as the Pulumi resource ID.
var idField = tfbridge.DelegateIDField("id", "systeam", "https://github.com/pawel-cygal/healtchecks")

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:            pfbridge.ShimProvider(shim.New()),
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
		Name:         "systeam",
		DisplayName:  "SysTeam Monitoring",
		Publisher:    "pawel-cygal",
		Description:  "A Pulumi provider for SysTeam Monitoring",
		Keywords:     []string{"pulumi", "systeam", "monitoring", "healthchecks"},
		License:      "Apache-2.0",
		Homepage:     "https://checks.systeam.pl",
		GitHubOrg:    "pawel-cygal",
		Config:       map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"systeam_project":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project"), ComputeID: idField},
			"systeam_check":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Check"), ComputeID: idField},
			"systeam_check_slo":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CheckSlo"), ComputeID: idField},
			"systeam_notification_channel":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NotificationChannel"), ComputeID: idField},
			"systeam_maintenance_window":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MaintenanceWindow"), ComputeID: idField},
			"systeam_escalation_policy":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EscalationPolicy"), ComputeID: idField},
			"systeam_oncall_schedule":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OncallSchedule"), ComputeID: idField},
			"systeam_status_page":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StatusPage"), ComputeID: idField},
			"systeam_integration_key":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IntegrationKey"), ComputeID: idField},
			"systeam_agent_registration_token": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AgentRegistrationToken"), ComputeID: idField},
			"systeam_team":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team"), ComputeID: idField},
			"systeam_service":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Service"), ComputeID: idField},
			"systeam_lifecycle_watch":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LifecycleWatch"), ComputeID: idField},
			"systeam_contact_method":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContactMethod"), ComputeID: idField},
			"systeam_playbook":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Playbook"), ComputeID: idField},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"systeam_organization": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrganization")},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_systeam",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule(mainPkg, mainMod, tokens.MakeStandard(mainPkg)))

	return prov
}
