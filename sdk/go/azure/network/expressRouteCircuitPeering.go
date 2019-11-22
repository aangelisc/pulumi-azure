// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an ExpressRoute Circuit Peering.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/express_route_circuit_peering.html.markdown.
type ExpressRouteCircuitPeering struct {
	s *pulumi.ResourceState
}

// NewExpressRouteCircuitPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringArgs, opts ...pulumi.ResourceOpt) (*ExpressRouteCircuitPeering, error) {
	if args == nil || args.ExpressRouteCircuitName == nil {
		return nil, errors.New("missing required argument 'ExpressRouteCircuitName'")
	}
	if args == nil || args.PeeringType == nil {
		return nil, errors.New("missing required argument 'PeeringType'")
	}
	if args == nil || args.PrimaryPeerAddressPrefix == nil {
		return nil, errors.New("missing required argument 'PrimaryPeerAddressPrefix'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SecondaryPeerAddressPrefix == nil {
		return nil, errors.New("missing required argument 'SecondaryPeerAddressPrefix'")
	}
	if args == nil || args.VlanId == nil {
		return nil, errors.New("missing required argument 'VlanId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["expressRouteCircuitName"] = nil
		inputs["microsoftPeeringConfig"] = nil
		inputs["peerAsn"] = nil
		inputs["peeringType"] = nil
		inputs["primaryPeerAddressPrefix"] = nil
		inputs["resourceGroupName"] = nil
		inputs["secondaryPeerAddressPrefix"] = nil
		inputs["sharedKey"] = nil
		inputs["vlanId"] = nil
	} else {
		inputs["expressRouteCircuitName"] = args.ExpressRouteCircuitName
		inputs["microsoftPeeringConfig"] = args.MicrosoftPeeringConfig
		inputs["peerAsn"] = args.PeerAsn
		inputs["peeringType"] = args.PeeringType
		inputs["primaryPeerAddressPrefix"] = args.PrimaryPeerAddressPrefix
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["secondaryPeerAddressPrefix"] = args.SecondaryPeerAddressPrefix
		inputs["sharedKey"] = args.SharedKey
		inputs["vlanId"] = args.VlanId
	}
	inputs["azureAsn"] = nil
	inputs["primaryAzurePort"] = nil
	inputs["secondaryAzurePort"] = nil
	s, err := ctx.RegisterResource("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExpressRouteCircuitPeering{s: s}, nil
}

// GetExpressRouteCircuitPeering gets an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExpressRouteCircuitPeeringState, opts ...pulumi.ResourceOpt) (*ExpressRouteCircuitPeering, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["azureAsn"] = state.AzureAsn
		inputs["expressRouteCircuitName"] = state.ExpressRouteCircuitName
		inputs["microsoftPeeringConfig"] = state.MicrosoftPeeringConfig
		inputs["peerAsn"] = state.PeerAsn
		inputs["peeringType"] = state.PeeringType
		inputs["primaryAzurePort"] = state.PrimaryAzurePort
		inputs["primaryPeerAddressPrefix"] = state.PrimaryPeerAddressPrefix
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["secondaryAzurePort"] = state.SecondaryAzurePort
		inputs["secondaryPeerAddressPrefix"] = state.SecondaryPeerAddressPrefix
		inputs["sharedKey"] = state.SharedKey
		inputs["vlanId"] = state.VlanId
	}
	s, err := ctx.ReadResource("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExpressRouteCircuitPeering{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ExpressRouteCircuitPeering) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ExpressRouteCircuitPeering) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ASN used by Azure.
func (r *ExpressRouteCircuitPeering) AzureAsn() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["azureAsn"])
}

// The name of the ExpressRoute Circuit in which to create the Peering.
func (r *ExpressRouteCircuitPeering) ExpressRouteCircuitName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["expressRouteCircuitName"])
}

// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
func (r *ExpressRouteCircuitPeering) MicrosoftPeeringConfig() pulumi.Output {
	return r.s.State["microsoftPeeringConfig"]
}

// The Either a 16-bit or a 32-bit ASN. Can either be public or private..
func (r *ExpressRouteCircuitPeering) PeerAsn() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["peerAsn"])
}

// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
func (r *ExpressRouteCircuitPeering) PeeringType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["peeringType"])
}

// The Primary Port used by Azure for this Peering.
func (r *ExpressRouteCircuitPeering) PrimaryAzurePort() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["primaryAzurePort"])
}

// A `/30` subnet for the primary link.
func (r *ExpressRouteCircuitPeering) PrimaryPeerAddressPrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["primaryPeerAddressPrefix"])
}

// The name of the resource group in which to
// create the Express Route Circuit Peering. Changing this forces a new resource to be created.
func (r *ExpressRouteCircuitPeering) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The Secondary Port used by Azure for this Peering.
func (r *ExpressRouteCircuitPeering) SecondaryAzurePort() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secondaryAzurePort"])
}

// A `/30` subnet for the secondary link.
func (r *ExpressRouteCircuitPeering) SecondaryPeerAddressPrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secondaryPeerAddressPrefix"])
}

// The shared key. Can be a maximum of 25 characters.
func (r *ExpressRouteCircuitPeering) SharedKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedKey"])
}

// A valid VLAN ID to establish this peering on.
func (r *ExpressRouteCircuitPeering) VlanId() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["vlanId"])
}

// Input properties used for looking up and filtering ExpressRouteCircuitPeering resources.
type ExpressRouteCircuitPeeringState struct {
	// The ASN used by Azure.
	AzureAsn interface{}
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName interface{}
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig interface{}
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private..
	PeerAsn interface{}
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType interface{}
	// The Primary Port used by Azure for this Peering.
	PrimaryAzurePort interface{}
	// A `/30` subnet for the primary link.
	PrimaryPeerAddressPrefix interface{}
	// The name of the resource group in which to
	// create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The Secondary Port used by Azure for this Peering.
	SecondaryAzurePort interface{}
	// A `/30` subnet for the secondary link.
	SecondaryPeerAddressPrefix interface{}
	// The shared key. Can be a maximum of 25 characters.
	SharedKey interface{}
	// A valid VLAN ID to establish this peering on.
	VlanId interface{}
}

// The set of arguments for constructing a ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitPeeringArgs struct {
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName interface{}
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig interface{}
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private..
	PeerAsn interface{}
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType interface{}
	// A `/30` subnet for the primary link.
	PrimaryPeerAddressPrefix interface{}
	// The name of the resource group in which to
	// create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `/30` subnet for the secondary link.
	SecondaryPeerAddressPrefix interface{}
	// The shared key. Can be a maximum of 25 characters.
	SharedKey interface{}
	// A valid VLAN ID to establish this peering on.
	VlanId interface{}
}
