// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notificationhub

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
	case "azure:notificationhub/authorizationRule:AuthorizationRule":
		r, err = NewAuthorizationRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure:notificationhub/hub:Hub":
		r, err = NewHub(ctx, name, nil, pulumi.URN_(urn))
	case "azure:notificationhub/namespace:Namespace":
		r, err = NewNamespace(ctx, name, nil, pulumi.URN_(urn))
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
		"notificationhub/authorizationRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"notificationhub/hub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"notificationhub/namespace",
		&module{version},
	)
}
