// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage Private DNS zone Virtual Network Links. These Links enable DNS resolution and registration inside Azure Virtual Networks using Azure Private DNS.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/private_dns_zone_virtual_network_link.html.markdown.
type ZoneVirtualNetworkLink struct {
	s *pulumi.ResourceState
}

// NewZoneVirtualNetworkLink registers a new resource with the given unique name, arguments, and options.
func NewZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *ZoneVirtualNetworkLinkArgs, opts ...pulumi.ResourceOpt) (*ZoneVirtualNetworkLink, error) {
	if args == nil || args.PrivateDnsZoneName == nil {
		return nil, errors.New("missing required argument 'PrivateDnsZoneName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkId == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["privateDnsZoneName"] = nil
		inputs["registrationEnabled"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["virtualNetworkId"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["privateDnsZoneName"] = args.PrivateDnsZoneName
		inputs["registrationEnabled"] = args.RegistrationEnabled
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["virtualNetworkId"] = args.VirtualNetworkId
	}
	s, err := ctx.RegisterResource("azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ZoneVirtualNetworkLink{s: s}, nil
}

// GetZoneVirtualNetworkLink gets an existing ZoneVirtualNetworkLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ZoneVirtualNetworkLinkState, opts ...pulumi.ResourceOpt) (*ZoneVirtualNetworkLink, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["privateDnsZoneName"] = state.PrivateDnsZoneName
		inputs["registrationEnabled"] = state.RegistrationEnabled
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["virtualNetworkId"] = state.VirtualNetworkId
	}
	s, err := ctx.ReadResource("azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ZoneVirtualNetworkLink{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ZoneVirtualNetworkLink) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ZoneVirtualNetworkLink) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
func (r *ZoneVirtualNetworkLink) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
func (r *ZoneVirtualNetworkLink) PrivateDnsZoneName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["privateDnsZoneName"])
}

// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
func (r *ZoneVirtualNetworkLink) RegistrationEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["registrationEnabled"])
}

// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
func (r *ZoneVirtualNetworkLink) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *ZoneVirtualNetworkLink) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
func (r *ZoneVirtualNetworkLink) VirtualNetworkId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["virtualNetworkId"])
}

// Input properties used for looking up and filtering ZoneVirtualNetworkLink resources.
type ZoneVirtualNetworkLinkState struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName interface{}
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId interface{}
}

// The set of arguments for constructing a ZoneVirtualNetworkLink resource.
type ZoneVirtualNetworkLinkArgs struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName interface{}
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Resource ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId interface{}
}
