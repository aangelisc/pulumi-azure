// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Virtual Network Gateway to establish secure, cross-premises connectivity.
// 
// > **Note:** Please be aware that provisioning a Virtual Network Gateway takes a long time (between 30 minutes and 1 hour)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_network_gateway.html.markdown.
type VirtualNetworkGateway struct {
	s *pulumi.ResourceState
}

// NewVirtualNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGateway(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayArgs, opts ...pulumi.ResourceOpt) (*VirtualNetworkGateway, error) {
	if args == nil || args.IpConfigurations == nil {
		return nil, errors.New("missing required argument 'IpConfigurations'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["activeActive"] = nil
		inputs["bgpSettings"] = nil
		inputs["defaultLocalNetworkGatewayId"] = nil
		inputs["enableBgp"] = nil
		inputs["ipConfigurations"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
		inputs["type"] = nil
		inputs["vpnClientConfiguration"] = nil
		inputs["vpnType"] = nil
	} else {
		inputs["activeActive"] = args.ActiveActive
		inputs["bgpSettings"] = args.BgpSettings
		inputs["defaultLocalNetworkGatewayId"] = args.DefaultLocalNetworkGatewayId
		inputs["enableBgp"] = args.EnableBgp
		inputs["ipConfigurations"] = args.IpConfigurations
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
		inputs["type"] = args.Type
		inputs["vpnClientConfiguration"] = args.VpnClientConfiguration
		inputs["vpnType"] = args.VpnType
	}
	s, err := ctx.RegisterResource("azure:network/virtualNetworkGateway:VirtualNetworkGateway", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkGateway{s: s}, nil
}

// GetVirtualNetworkGateway gets an existing VirtualNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualNetworkGatewayState, opts ...pulumi.ResourceOpt) (*VirtualNetworkGateway, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["activeActive"] = state.ActiveActive
		inputs["bgpSettings"] = state.BgpSettings
		inputs["defaultLocalNetworkGatewayId"] = state.DefaultLocalNetworkGatewayId
		inputs["enableBgp"] = state.EnableBgp
		inputs["ipConfigurations"] = state.IpConfigurations
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
		inputs["type"] = state.Type
		inputs["vpnClientConfiguration"] = state.VpnClientConfiguration
		inputs["vpnType"] = state.VpnType
	}
	s, err := ctx.ReadResource("azure:network/virtualNetworkGateway:VirtualNetworkGateway", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkGateway{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualNetworkGateway) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualNetworkGateway) ID() pulumi.IDOutput {
	return r.s.ID()
}

// If `true`, an active-active Virtual Network Gateway
// will be created. An active-active gateway requires a `HighPerformance` or an
// `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
// Defaults to `false`.
func (r *VirtualNetworkGateway) ActiveActive() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["activeActive"])
}

func (r *VirtualNetworkGateway) BgpSettings() pulumi.Output {
	return r.s.State["bgpSettings"]
}

// The ID of the local network gateway
// through which outbound Internet traffic from the virtual network in which the
// gateway is created will be routed (*forced tunneling*). Refer to the
// [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
// If not specified, forced tunneling is disabled.
func (r *VirtualNetworkGateway) DefaultLocalNetworkGatewayId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultLocalNetworkGatewayId"])
}

// If `true`, BGP (Border Gateway Protocol) will be enabled
// for this Virtual Network Gateway. Defaults to `false`.
func (r *VirtualNetworkGateway) EnableBgp() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableBgp"])
}

// One or two `ipConfiguration` blocks documented below.
// An active-standby gateway requires exactly one `ipConfiguration` block whereas
// an active-active gateway requires exactly two `ipConfiguration` blocks.
func (r *VirtualNetworkGateway) IpConfigurations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ipConfigurations"])
}

// The location/region where the Virtual Network Gateway is
// located. Changing the location/region forces a new resource to be created.
func (r *VirtualNetworkGateway) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The name of the Virtual Network Gateway. Changing the name
// forces a new resource to be created.
func (r *VirtualNetworkGateway) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to
// create the Virtual Network Gateway. Changing the resource group name forces
// a new resource to be created.
func (r *VirtualNetworkGateway) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Configuration of the size and capacity of the virtual network
// gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
// `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ`, and `VpnGw3AZ`
// and depend on the `type` and `vpnType` arguments.
// A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
// sku is only supported by an `ExpressRoute` gateway.
func (r *VirtualNetworkGateway) Sku() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sku"])
}

// A mapping of tags to assign to the resource.
func (r *VirtualNetworkGateway) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The type of the Virtual Network Gateway. Valid options are
// `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
func (r *VirtualNetworkGateway) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// A `vpnClientConfiguration` block which
// is documented below. In this block the Virtual Network Gateway can be configured
// to accept IPSec point-to-site connections.
func (r *VirtualNetworkGateway) VpnClientConfiguration() pulumi.Output {
	return r.s.State["vpnClientConfiguration"]
}

// The routing type of the Virtual Network Gateway. Valid
// options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
func (r *VirtualNetworkGateway) VpnType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpnType"])
}

// Input properties used for looking up and filtering VirtualNetworkGateway resources.
type VirtualNetworkGatewayState struct {
	// If `true`, an active-active Virtual Network Gateway
	// will be created. An active-active gateway requires a `HighPerformance` or an
	// `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
	// Defaults to `false`.
	ActiveActive interface{}
	BgpSettings interface{}
	// The ID of the local network gateway
	// through which outbound Internet traffic from the virtual network in which the
	// gateway is created will be routed (*forced tunneling*). Refer to the
	// [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
	// If not specified, forced tunneling is disabled.
	DefaultLocalNetworkGatewayId interface{}
	// If `true`, BGP (Border Gateway Protocol) will be enabled
	// for this Virtual Network Gateway. Defaults to `false`.
	EnableBgp interface{}
	// One or two `ipConfiguration` blocks documented below.
	// An active-standby gateway requires exactly one `ipConfiguration` block whereas
	// an active-active gateway requires exactly two `ipConfiguration` blocks.
	IpConfigurations interface{}
	// The location/region where the Virtual Network Gateway is
	// located. Changing the location/region forces a new resource to be created.
	Location interface{}
	// The name of the Virtual Network Gateway. Changing the name
	// forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to
	// create the Virtual Network Gateway. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName interface{}
	// Configuration of the size and capacity of the virtual network
	// gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
	// `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ`, and `VpnGw3AZ`
	// and depend on the `type` and `vpnType` arguments.
	// A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
	// sku is only supported by an `ExpressRoute` gateway.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The type of the Virtual Network Gateway. Valid options are
	// `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
	Type interface{}
	// A `vpnClientConfiguration` block which
	// is documented below. In this block the Virtual Network Gateway can be configured
	// to accept IPSec point-to-site connections.
	VpnClientConfiguration interface{}
	// The routing type of the Virtual Network Gateway. Valid
	// options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
	VpnType interface{}
}

// The set of arguments for constructing a VirtualNetworkGateway resource.
type VirtualNetworkGatewayArgs struct {
	// If `true`, an active-active Virtual Network Gateway
	// will be created. An active-active gateway requires a `HighPerformance` or an
	// `UltraPerformance` sku. If `false`, an active-standby gateway will be created.
	// Defaults to `false`.
	ActiveActive interface{}
	BgpSettings interface{}
	// The ID of the local network gateway
	// through which outbound Internet traffic from the virtual network in which the
	// gateway is created will be routed (*forced tunneling*). Refer to the
	// [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
	// If not specified, forced tunneling is disabled.
	DefaultLocalNetworkGatewayId interface{}
	// If `true`, BGP (Border Gateway Protocol) will be enabled
	// for this Virtual Network Gateway. Defaults to `false`.
	EnableBgp interface{}
	// One or two `ipConfiguration` blocks documented below.
	// An active-standby gateway requires exactly one `ipConfiguration` block whereas
	// an active-active gateway requires exactly two `ipConfiguration` blocks.
	IpConfigurations interface{}
	// The location/region where the Virtual Network Gateway is
	// located. Changing the location/region forces a new resource to be created.
	Location interface{}
	// The name of the Virtual Network Gateway. Changing the name
	// forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to
	// create the Virtual Network Gateway. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName interface{}
	// Configuration of the size and capacity of the virtual network
	// gateway. Valid options are `Basic`, `Standard`, `HighPerformance`, `UltraPerformance`,
	// `ErGw1AZ`, `ErGw2AZ`, `ErGw3AZ`, `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ`, and `VpnGw3AZ`
	// and depend on the `type` and `vpnType` arguments.
	// A `PolicyBased` gateway only supports the `Basic` sku. Further, the `UltraPerformance`
	// sku is only supported by an `ExpressRoute` gateway.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The type of the Virtual Network Gateway. Valid options are
	// `Vpn` or `ExpressRoute`. Changing the type forces a new resource to be created.
	Type interface{}
	// A `vpnClientConfiguration` block which
	// is documented below. In this block the Virtual Network Gateway can be configured
	// to accept IPSec point-to-site connections.
	VpnClientConfiguration interface{}
	// The routing type of the Virtual Network Gateway. Valid
	// options are `RouteBased` or `PolicyBased`. Defaults to `RouteBased`.
	VpnType interface{}
}
