// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Traffic Manager Profile to which multiple endpoints can be attached.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/traffic_manager_profile.html.markdown.
type TrafficManagerProfile struct {
	pulumi.CustomResourceState

	// This block specifies the DNS configuration of the
	// Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfigOutput `pulumi:"dnsConfig"`
	// The FQDN of the created Profile.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// This block specifies the Endpoint monitoring
	// configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfigOutput `pulumi:"monitorConfig"`
	// The name of the virtual network. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the profile, can be set to either
	// `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringOutput `pulumi:"profileStatus"`
	// The name of the resource group in which to
	// create the virtual network.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
	// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
	// - `Performance` - Traffic is routed via the User's closest Endpoint
	// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
	// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
	// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
	TrafficRoutingMethod pulumi.StringOutput `pulumi:"trafficRoutingMethod"`
}

// NewTrafficManagerProfile registers a new resource with the given unique name, arguments, and options.
func NewTrafficManagerProfile(ctx *pulumi.Context,
	name string, args *TrafficManagerProfileArgs, opts ...pulumi.ResourceOption) (*TrafficManagerProfile, error) {
	if args == nil || args.DnsConfig == nil {
		return nil, errors.New("missing required argument 'DnsConfig'")
	}
	if args == nil || args.MonitorConfig == nil {
		return nil, errors.New("missing required argument 'MonitorConfig'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TrafficRoutingMethod == nil {
		return nil, errors.New("missing required argument 'TrafficRoutingMethod'")
	}
	if args == nil {
		args = &TrafficManagerProfileArgs{}
	}
	var resource TrafficManagerProfile
	err := ctx.RegisterResource("azure:network/trafficManagerProfile:TrafficManagerProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficManagerProfile gets an existing TrafficManagerProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficManagerProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficManagerProfileState, opts ...pulumi.ResourceOption) (*TrafficManagerProfile, error) {
	var resource TrafficManagerProfile
	err := ctx.ReadResource("azure:network/trafficManagerProfile:TrafficManagerProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficManagerProfile resources.
type trafficManagerProfileState struct {
	// This block specifies the DNS configuration of the
	// Profile, it supports the fields documented below.
	DnsConfig *TrafficManagerProfileDnsConfig `pulumi:"dnsConfig"`
	// The FQDN of the created Profile.
	Fqdn *string `pulumi:"fqdn"`
	// This block specifies the Endpoint monitoring
	// configuration for the Profile, it supports the fields documented below.
	MonitorConfig *TrafficManagerProfileMonitorConfig `pulumi:"monitorConfig"`
	// The name of the virtual network. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The status of the profile, can be set to either
	// `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group in which to
	// create the virtual network.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
	// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
	// - `Performance` - Traffic is routed via the User's closest Endpoint
	// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
	// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
	// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
}

type TrafficManagerProfileState struct {
	// This block specifies the DNS configuration of the
	// Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfigPtrInput
	// The FQDN of the created Profile.
	Fqdn pulumi.StringPtrInput
	// This block specifies the Endpoint monitoring
	// configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfigPtrInput
	// The name of the virtual network. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The status of the profile, can be set to either
	// `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the virtual network.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the algorithm used to route traffic, possible values are:
	// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
	// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
	// - `Performance` - Traffic is routed via the User's closest Endpoint
	// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
	// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
	// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
	TrafficRoutingMethod pulumi.StringPtrInput
}

func (TrafficManagerProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerProfileState)(nil)).Elem()
}

type trafficManagerProfileArgs struct {
	// This block specifies the DNS configuration of the
	// Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfig `pulumi:"dnsConfig"`
	// This block specifies the Endpoint monitoring
	// configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfig `pulumi:"monitorConfig"`
	// The name of the virtual network. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The status of the profile, can be set to either
	// `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group in which to
	// create the virtual network.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
	// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
	// - `Performance` - Traffic is routed via the User's closest Endpoint
	// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
	// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
	// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
	TrafficRoutingMethod string `pulumi:"trafficRoutingMethod"`
}

// The set of arguments for constructing a TrafficManagerProfile resource.
type TrafficManagerProfileArgs struct {
	// This block specifies the DNS configuration of the
	// Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfigInput
	// This block specifies the Endpoint monitoring
	// configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfigInput
	// The name of the virtual network. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The status of the profile, can be set to either
	// `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the virtual network.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the algorithm used to route traffic, possible values are:
	// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
	// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
	// - `Performance` - Traffic is routed via the User's closest Endpoint
	// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
	// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
	// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
	TrafficRoutingMethod pulumi.StringInput
}

func (TrafficManagerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerProfileArgs)(nil)).Elem()
}

