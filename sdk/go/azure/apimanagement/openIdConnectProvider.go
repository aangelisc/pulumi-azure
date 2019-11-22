// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an OpenID Connect Provider within a API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_openid_connect_provider.html.markdown.
type OpenIdConnectProvider struct {
	s *pulumi.ResourceState
}

// NewOpenIdConnectProvider registers a new resource with the given unique name, arguments, and options.
func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOpt) (*OpenIdConnectProvider, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.MetadataEndpoint == nil {
		return nil, errors.New("missing required argument 'MetadataEndpoint'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["clientId"] = nil
		inputs["clientSecret"] = nil
		inputs["description"] = nil
		inputs["displayName"] = nil
		inputs["metadataEndpoint"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["clientId"] = args.ClientId
		inputs["clientSecret"] = args.ClientSecret
		inputs["description"] = args.Description
		inputs["displayName"] = args.DisplayName
		inputs["metadataEndpoint"] = args.MetadataEndpoint
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OpenIdConnectProvider{s: s}, nil
}

// GetOpenIdConnectProvider gets an existing OpenIdConnectProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOpt) (*OpenIdConnectProvider, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["clientId"] = state.ClientId
		inputs["clientSecret"] = state.ClientSecret
		inputs["description"] = state.Description
		inputs["displayName"] = state.DisplayName
		inputs["metadataEndpoint"] = state.MetadataEndpoint
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OpenIdConnectProvider{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OpenIdConnectProvider) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OpenIdConnectProvider) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
func (r *OpenIdConnectProvider) ApiManagementName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// The Client ID used for the Client Application.
func (r *OpenIdConnectProvider) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The Client Secret used for the Client Application.
func (r *OpenIdConnectProvider) ClientSecret() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientSecret"])
}

// A description of this OpenID Connect Provider.
func (r *OpenIdConnectProvider) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// A user-friendly name for this OpenID Connect Provider.
func (r *OpenIdConnectProvider) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// The URI of the Metadata endpoint.
func (r *OpenIdConnectProvider) MetadataEndpoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["metadataEndpoint"])
}

// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
func (r *OpenIdConnectProvider) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
func (r *OpenIdConnectProvider) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering OpenIdConnectProvider resources.
type OpenIdConnectProviderState struct {
	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The Client ID used for the Client Application.
	ClientId interface{}
	// The Client Secret used for the Client Application.
	ClientSecret interface{}
	// A description of this OpenID Connect Provider.
	Description interface{}
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName interface{}
	// The URI of the Metadata endpoint.
	MetadataEndpoint interface{}
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}

// The set of arguments for constructing a OpenIdConnectProvider resource.
type OpenIdConnectProviderArgs struct {
	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The Client ID used for the Client Application.
	ClientId interface{}
	// The Client Secret used for the Client Application.
	ClientSecret interface{}
	// A description of this OpenID Connect Provider.
	Description interface{}
	// A user-friendly name for this OpenID Connect Provider.
	DisplayName interface{}
	// The URI of the Metadata endpoint.
	MetadataEndpoint interface{}
	// the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}
