// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management.html.markdown.
type Service struct {
	s *pulumi.ResourceState
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOpt) (*Service, error) {
	if args == nil || args.PublisherEmail == nil {
		return nil, errors.New("missing required argument 'PublisherEmail'")
	}
	if args == nil || args.PublisherName == nil {
		return nil, errors.New("missing required argument 'PublisherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["additionalLocations"] = nil
		inputs["certificates"] = nil
		inputs["hostnameConfiguration"] = nil
		inputs["identity"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["notificationSenderEmail"] = nil
		inputs["policy"] = nil
		inputs["publisherEmail"] = nil
		inputs["publisherName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["security"] = nil
		inputs["signIn"] = nil
		inputs["signUp"] = nil
		inputs["sku"] = nil
		inputs["skuName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["additionalLocations"] = args.AdditionalLocations
		inputs["certificates"] = args.Certificates
		inputs["hostnameConfiguration"] = args.HostnameConfiguration
		inputs["identity"] = args.Identity
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["notificationSenderEmail"] = args.NotificationSenderEmail
		inputs["policy"] = args.Policy
		inputs["publisherEmail"] = args.PublisherEmail
		inputs["publisherName"] = args.PublisherName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["security"] = args.Security
		inputs["signIn"] = args.SignIn
		inputs["signUp"] = args.SignUp
		inputs["sku"] = args.Sku
		inputs["skuName"] = args.SkuName
		inputs["tags"] = args.Tags
	}
	inputs["gatewayRegionalUrl"] = nil
	inputs["gatewayUrl"] = nil
	inputs["managementApiUrl"] = nil
	inputs["portalUrl"] = nil
	inputs["publicIpAddresses"] = nil
	inputs["scmUrl"] = nil
	s, err := ctx.RegisterResource("azure:apimanagement/service:Service", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Service{s: s}, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServiceState, opts ...pulumi.ResourceOpt) (*Service, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["additionalLocations"] = state.AdditionalLocations
		inputs["certificates"] = state.Certificates
		inputs["gatewayRegionalUrl"] = state.GatewayRegionalUrl
		inputs["gatewayUrl"] = state.GatewayUrl
		inputs["hostnameConfiguration"] = state.HostnameConfiguration
		inputs["identity"] = state.Identity
		inputs["location"] = state.Location
		inputs["managementApiUrl"] = state.ManagementApiUrl
		inputs["name"] = state.Name
		inputs["notificationSenderEmail"] = state.NotificationSenderEmail
		inputs["policy"] = state.Policy
		inputs["portalUrl"] = state.PortalUrl
		inputs["publicIpAddresses"] = state.PublicIpAddresses
		inputs["publisherEmail"] = state.PublisherEmail
		inputs["publisherName"] = state.PublisherName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["scmUrl"] = state.ScmUrl
		inputs["security"] = state.Security
		inputs["signIn"] = state.SignIn
		inputs["signUp"] = state.SignUp
		inputs["sku"] = state.Sku
		inputs["skuName"] = state.SkuName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:apimanagement/service:Service", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Service{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Service) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Service) ID() pulumi.IDOutput {
	return r.s.ID()
}

// One or more `additionalLocation` blocks as defined below.
func (r *Service) AdditionalLocations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["additionalLocations"])
}

// One or more (up to 10) `certificate` blocks as defined below.
func (r *Service) Certificates() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["certificates"])
}

// The URL of the Regional Gateway for the API Management Service in the specified region.
func (r *Service) GatewayRegionalUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["gatewayRegionalUrl"])
}

// The URL of the Gateway for the API Management Service.
func (r *Service) GatewayUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["gatewayUrl"])
}

// A `hostnameConfiguration` block as defined below.
func (r *Service) HostnameConfiguration() pulumi.Output {
	return r.s.State["hostnameConfiguration"]
}

// An `identity` block is documented below.
func (r *Service) Identity() pulumi.Output {
	return r.s.State["identity"]
}

// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
func (r *Service) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The URL for the Management API associated with this API Management service.
func (r *Service) ManagementApiUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["managementApiUrl"])
}

// The name of the API Management Service. Changing this forces a new resource to be created.
func (r *Service) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Email address from which the notification will be sent.
func (r *Service) NotificationSenderEmail() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["notificationSenderEmail"])
}

// A `policy` block as defined below.
func (r *Service) Policy() pulumi.Output {
	return r.s.State["policy"]
}

// The URL for the Publisher Portal associated with this API Management service.
func (r *Service) PortalUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["portalUrl"])
}

// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
func (r *Service) PublicIpAddresses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["publicIpAddresses"])
}

// The email of publisher/company.
func (r *Service) PublisherEmail() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["publisherEmail"])
}

// The name of publisher/company.
func (r *Service) PublisherName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["publisherName"])
}

// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
func (r *Service) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
func (r *Service) ScmUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["scmUrl"])
}

// A `security` block as defined below.
func (r *Service) Security() pulumi.Output {
	return r.s.State["security"]
}

// A `signIn` block as defined below.
func (r *Service) SignIn() pulumi.Output {
	return r.s.State["signIn"]
}

// A `signUp` block as defined below.
func (r *Service) SignUp() pulumi.Output {
	return r.s.State["signUp"]
}

// A `sku` block as documented below
func (r *Service) Sku() pulumi.Output {
	return r.s.State["sku"]
}

// `skuName` is a string consisting of two parts separated by an underscore(\_). The fist part is the `name`, valid values include: `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
func (r *Service) SkuName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["skuName"])
}

// A mapping of tags assigned to the resource.
func (r *Service) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Service resources.
type ServiceState struct {
	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations interface{}
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates interface{}
	// The URL of the Regional Gateway for the API Management Service in the specified region.
	GatewayRegionalUrl interface{}
	// The URL of the Gateway for the API Management Service.
	GatewayUrl interface{}
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration interface{}
	// An `identity` block is documented below.
	Identity interface{}
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location interface{}
	// The URL for the Management API associated with this API Management service.
	ManagementApiUrl interface{}
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name interface{}
	// Email address from which the notification will be sent.
	NotificationSenderEmail interface{}
	// A `policy` block as defined below.
	Policy interface{}
	// The URL for the Publisher Portal associated with this API Management service.
	PortalUrl interface{}
	// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
	PublicIpAddresses interface{}
	// The email of publisher/company.
	PublisherEmail interface{}
	// The name of publisher/company.
	PublisherName interface{}
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
	ScmUrl interface{}
	// A `security` block as defined below.
	Security interface{}
	// A `signIn` block as defined below.
	SignIn interface{}
	// A `signUp` block as defined below.
	SignUp interface{}
	// A `sku` block as documented below
	Sku interface{}
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The fist part is the `name`, valid values include: `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations interface{}
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates interface{}
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration interface{}
	// An `identity` block is documented below.
	Identity interface{}
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name interface{}
	// Email address from which the notification will be sent.
	NotificationSenderEmail interface{}
	// A `policy` block as defined below.
	Policy interface{}
	// The email of publisher/company.
	PublisherEmail interface{}
	// The name of publisher/company.
	PublisherName interface{}
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `security` block as defined below.
	Security interface{}
	// A `signIn` block as defined below.
	SignIn interface{}
	// A `signUp` block as defined below.
	SignUp interface{}
	// A `sku` block as documented below
	Sku interface{}
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The fist part is the `name`, valid values include: `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
}
