// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A CDN Endpoint is the entity within a CDN Profile containing configuration information regarding caching behaviors and origins. The CDN Endpoint is exposed using the URL format <endpointname>.azureedge.net.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cdn_endpoint.html.markdown.
type Endpoint struct {
	s *pulumi.ResourceState
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOpt) (*Endpoint, error) {
	if args == nil || args.Origins == nil {
		return nil, errors.New("missing required argument 'Origins'")
	}
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["contentTypesToCompresses"] = nil
		inputs["geoFilters"] = nil
		inputs["isCompressionEnabled"] = nil
		inputs["isHttpAllowed"] = nil
		inputs["isHttpsAllowed"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["optimizationType"] = nil
		inputs["origins"] = nil
		inputs["originHostHeader"] = nil
		inputs["originPath"] = nil
		inputs["probePath"] = nil
		inputs["profileName"] = nil
		inputs["querystringCachingBehaviour"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["contentTypesToCompresses"] = args.ContentTypesToCompresses
		inputs["geoFilters"] = args.GeoFilters
		inputs["isCompressionEnabled"] = args.IsCompressionEnabled
		inputs["isHttpAllowed"] = args.IsHttpAllowed
		inputs["isHttpsAllowed"] = args.IsHttpsAllowed
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["optimizationType"] = args.OptimizationType
		inputs["origins"] = args.Origins
		inputs["originHostHeader"] = args.OriginHostHeader
		inputs["originPath"] = args.OriginPath
		inputs["probePath"] = args.ProbePath
		inputs["profileName"] = args.ProfileName
		inputs["querystringCachingBehaviour"] = args.QuerystringCachingBehaviour
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
	}
	inputs["hostName"] = nil
	s, err := ctx.RegisterResource("azure:cdn/endpoint:Endpoint", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Endpoint{s: s}, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EndpointState, opts ...pulumi.ResourceOpt) (*Endpoint, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["contentTypesToCompresses"] = state.ContentTypesToCompresses
		inputs["geoFilters"] = state.GeoFilters
		inputs["hostName"] = state.HostName
		inputs["isCompressionEnabled"] = state.IsCompressionEnabled
		inputs["isHttpAllowed"] = state.IsHttpAllowed
		inputs["isHttpsAllowed"] = state.IsHttpsAllowed
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["optimizationType"] = state.OptimizationType
		inputs["origins"] = state.Origins
		inputs["originHostHeader"] = state.OriginHostHeader
		inputs["originPath"] = state.OriginPath
		inputs["probePath"] = state.ProbePath
		inputs["profileName"] = state.ProfileName
		inputs["querystringCachingBehaviour"] = state.QuerystringCachingBehaviour
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:cdn/endpoint:Endpoint", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Endpoint{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Endpoint) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Endpoint) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
func (r *Endpoint) ContentTypesToCompresses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["contentTypesToCompresses"])
}

// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
func (r *Endpoint) GeoFilters() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["geoFilters"])
}

func (r *Endpoint) HostName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostName"])
}

// Indicates whether compression is to be enabled. Defaults to false.
func (r *Endpoint) IsCompressionEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isCompressionEnabled"])
}

// Defaults to `true`.
func (r *Endpoint) IsHttpAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isHttpAllowed"])
}

// Defaults to `true`.
func (r *Endpoint) IsHttpsAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isHttpsAllowed"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *Endpoint) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the CDN Endpoint. Changing this forces a new resource to be created.
func (r *Endpoint) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
func (r *Endpoint) OptimizationType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["optimizationType"])
}

// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
func (r *Endpoint) Origins() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["origins"])
}

// The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
func (r *Endpoint) OriginHostHeader() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["originHostHeader"])
}

// The path used at for origin requests.
func (r *Endpoint) OriginPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["originPath"])
}

// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
func (r *Endpoint) ProbePath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["probePath"])
}

// The CDN Profile to which to attach the CDN Endpoint.
func (r *Endpoint) ProfileName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["profileName"])
}

// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. Defaults to `IgnoreQueryString`.
func (r *Endpoint) QuerystringCachingBehaviour() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["querystringCachingBehaviour"])
}

// The name of the resource group in which to create the CDN Endpoint.
func (r *Endpoint) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *Endpoint) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Endpoint resources.
type EndpointState struct {
	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses interface{}
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters interface{}
	HostName interface{}
	// Indicates whether compression is to be enabled. Defaults to false.
	IsCompressionEnabled interface{}
	// Defaults to `true`.
	IsHttpAllowed interface{}
	// Defaults to `true`.
	IsHttpsAllowed interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the CDN Endpoint. Changing this forces a new resource to be created.
	Name interface{}
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType interface{}
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins interface{}
	// The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
	OriginHostHeader interface{}
	// The path used at for origin requests.
	OriginPath interface{}
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath interface{}
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName interface{}
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour interface{}
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses interface{}
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters interface{}
	// Indicates whether compression is to be enabled. Defaults to false.
	IsCompressionEnabled interface{}
	// Defaults to `true`.
	IsHttpAllowed interface{}
	// Defaults to `true`.
	IsHttpsAllowed interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the CDN Endpoint. Changing this forces a new resource to be created.
	Name interface{}
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType interface{}
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins interface{}
	// The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
	OriginHostHeader interface{}
	// The path used at for origin requests.
	OriginPath interface{}
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath interface{}
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName interface{}
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour interface{}
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
