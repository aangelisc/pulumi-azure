// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Version Set within a API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_api_version_set.html.markdown.
type ApiVersionSet struct {
	s *pulumi.ResourceState
}

// NewApiVersionSet registers a new resource with the given unique name, arguments, and options.
func NewApiVersionSet(ctx *pulumi.Context,
	name string, args *ApiVersionSetArgs, opts ...pulumi.ResourceOpt) (*ApiVersionSet, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VersioningScheme == nil {
		return nil, errors.New("missing required argument 'VersioningScheme'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["description"] = nil
		inputs["displayName"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["versionHeaderName"] = nil
		inputs["versionQueryName"] = nil
		inputs["versioningScheme"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["description"] = args.Description
		inputs["displayName"] = args.DisplayName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["versionHeaderName"] = args.VersionHeaderName
		inputs["versionQueryName"] = args.VersionQueryName
		inputs["versioningScheme"] = args.VersioningScheme
	}
	s, err := ctx.RegisterResource("azure:apimanagement/apiVersionSet:ApiVersionSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiVersionSet{s: s}, nil
}

// GetApiVersionSet gets an existing ApiVersionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiVersionSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApiVersionSetState, opts ...pulumi.ResourceOpt) (*ApiVersionSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["description"] = state.Description
		inputs["displayName"] = state.DisplayName
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["versionHeaderName"] = state.VersionHeaderName
		inputs["versionQueryName"] = state.VersionQueryName
		inputs["versioningScheme"] = state.VersioningScheme
	}
	s, err := ctx.ReadResource("azure:apimanagement/apiVersionSet:ApiVersionSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiVersionSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApiVersionSet) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApiVersionSet) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
func (r *ApiVersionSet) ApiManagementName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// The description of API Version Set.
func (r *ApiVersionSet) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The display name of this API Version Set.
func (r *ApiVersionSet) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// The name of the API Version Set. Changing this forces a new resource to be created.
func (r *ApiVersionSet) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
func (r *ApiVersionSet) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The name of the Header which should be read from Inbound Requests which defines the API Version.
func (r *ApiVersionSet) VersionHeaderName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["versionHeaderName"])
}

// The name of the Query String which should be read from Inbound Requests which defines the API Version.
func (r *ApiVersionSet) VersionQueryName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["versionQueryName"])
}

// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
func (r *ApiVersionSet) VersioningScheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["versioningScheme"])
}

// Input properties used for looking up and filtering ApiVersionSet resources.
type ApiVersionSetState struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The description of API Version Set.
	Description interface{}
	// The display name of this API Version Set.
	DisplayName interface{}
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName interface{}
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName interface{}
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme interface{}
}

// The set of arguments for constructing a ApiVersionSet resource.
type ApiVersionSetArgs struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The description of API Version Set.
	Description interface{}
	// The display name of this API Version Set.
	DisplayName interface{}
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName interface{}
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName interface{}
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme interface{}
}
