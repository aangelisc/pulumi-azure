// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

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
	case "azure:servicefabric/cluster:Cluster":
		r, err = NewCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:servicefabric/meshApplication:MeshApplication":
		r, err = NewMeshApplication(ctx, name, nil, pulumi.URN_(urn))
	case "azure:servicefabric/meshLocalNetwork:MeshLocalNetwork":
		r, err = NewMeshLocalNetwork(ctx, name, nil, pulumi.URN_(urn))
	case "azure:servicefabric/meshSecret:MeshSecret":
		r, err = NewMeshSecret(ctx, name, nil, pulumi.URN_(urn))
	case "azure:servicefabric/meshSecretValue:MeshSecretValue":
		r, err = NewMeshSecretValue(ctx, name, nil, pulumi.URN_(urn))
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
		"servicefabric/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicefabric/meshApplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicefabric/meshLocalNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicefabric/meshSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"servicefabric/meshSecretValue",
		&module{version},
	)
}
