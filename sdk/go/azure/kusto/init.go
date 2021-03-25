// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

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
	case "azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration":
		r, err = NewAttachedDatabaseConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/cluster:Cluster":
		r, err = NewCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/clusterCustomerManagedKey:ClusterCustomerManagedKey":
		r, err = NewClusterCustomerManagedKey(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment":
		r, err = NewClusterPrincipalAssignment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/database:Database":
		r, err = NewDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/databasePrincipal:DatabasePrincipal":
		r, err = NewDatabasePrincipal(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment":
		r, err = NewDatabasePrincipalAssignment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/eventGridDataConnection:EventGridDataConnection":
		r, err = NewEventGridDataConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/eventhubDataConnection:EventhubDataConnection":
		r, err = NewEventhubDataConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure:kusto/iotHubDataConnection:IotHubDataConnection":
		r, err = NewIotHubDataConnection(ctx, name, nil, pulumi.URN_(urn))
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
		"kusto/attachedDatabaseConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/clusterCustomerManagedKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/clusterPrincipalAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/databasePrincipal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/databasePrincipalAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/eventGridDataConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/eventhubDataConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/iotHubDataConnection",
		&module{version},
	)
}
