// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

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
	case "azure:mariadb/configuration:Configuration":
		r, err = NewConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure:mariadb/database:Database":
		r, err = NewDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure:mariadb/firewallRule:FirewallRule":
		r, err = NewFirewallRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure:mariadb/server:Server":
		r, err = NewServer(ctx, name, nil, pulumi.URN_(urn))
	case "azure:mariadb/virtualNetworkRule:VirtualNetworkRule":
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
		"mariadb/configuration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mariadb/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mariadb/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mariadb/server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"mariadb/virtualNetworkRule",
		&module{version},
	)
}
