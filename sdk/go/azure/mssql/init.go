// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:mssql/database:Database":
		r = &Database{}
	case "azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy":
		r = &DatabaseExtendedAuditingPolicy{}
	case "azure:mssql/databaseVulnerabilityAssessmentRuleBaseline:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure:mssql/elasticPool:ElasticPool":
		r = &ElasticPool{}
	case "azure:mssql/firewallRule:FirewallRule":
		r = &FirewallRule{}
	case "azure:mssql/server:Server":
		r = &Server{}
	case "azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy":
		r = &ServerExtendedAuditingPolicy{}
	case "azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure:mssql/serverVulnerabilityAssessment:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure:mssql/virtualMachine:VirtualMachine":
		r = &VirtualMachine{}
	case "azure:mssql/virtualNetworkRule:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/databaseExtendedAuditingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/databaseVulnerabilityAssessmentRuleBaseline",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/elasticPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/serverExtendedAuditingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/serverSecurityAlertPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/serverVulnerabilityAssessment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/virtualMachine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mssql/virtualNetworkRule",
		&module{version},
	)
}
