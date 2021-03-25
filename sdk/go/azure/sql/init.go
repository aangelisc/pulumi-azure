// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

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
	case "azure:sql/activeDirectoryAdministrator:ActiveDirectoryAdministrator":
		r, err = NewActiveDirectoryAdministrator(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sql/database:Database":
		r, err = NewDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sql/elasticPool:ElasticPool":
		r, err = NewElasticPool(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sql/failoverGroup:FailoverGroup":
		r, err = NewFailoverGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sql/firewallRule:FirewallRule":
		r, err = NewFirewallRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sql/sqlServer:SqlServer":
		r, err = NewSqlServer(ctx, name, nil, pulumi.URN_(urn))
	case "azure:sql/virtualNetworkRule:VirtualNetworkRule":
		r, err = NewVirtualNetworkRule(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"sql/activeDirectoryAdministrator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sql/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sql/elasticPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sql/failoverGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sql/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sql/sqlServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"sql/virtualNetworkRule",
		&module{version},
	)
}
