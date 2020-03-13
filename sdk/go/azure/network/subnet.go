// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.
//
// > **NOTE on Virtual Networks and Subnet's:** This provider currently
// provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
// At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/subnet.html.markdown.
type Subnet struct {
	pulumi.CustomResourceState

	// The address prefix to use for the subnet.
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// One or more `delegation` blocks as defined below.
	Delegations SubnetDelegationArrayOutput `pulumi:"delegations"`
	// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
	EnforcePrivateLinkEndpointNetworkPolicies pulumi.BoolPtrOutput `pulumi:"enforcePrivateLinkEndpointNetworkPolicies"`
	// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
	EnforcePrivateLinkServiceNetworkPolicies pulumi.BoolPtrOutput `pulumi:"enforcePrivateLinkServiceNetworkPolicies"`
	// The name of the subnet. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
	ServiceEndpoints pulumi.StringArrayOutput `pulumi:"serviceEndpoints"`
	// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
	VirtualNetworkName pulumi.StringOutput `pulumi:"virtualNetworkName"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil || args.AddressPrefix == nil {
		return nil, errors.New("missing required argument 'AddressPrefix'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkName == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkName'")
	}
	if args == nil {
		args = &SubnetArgs{}
	}
	var resource Subnet
	err := ctx.RegisterResource("azure:network/subnet:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("azure:network/subnet:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
	// The address prefix to use for the subnet.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// One or more `delegation` blocks as defined below.
	Delegations []SubnetDelegation `pulumi:"delegations"`
	// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
	EnforcePrivateLinkEndpointNetworkPolicies *bool `pulumi:"enforcePrivateLinkEndpointNetworkPolicies"`
	// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
	EnforcePrivateLinkServiceNetworkPolicies *bool `pulumi:"enforcePrivateLinkServiceNetworkPolicies"`
	// The name of the subnet. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
	ServiceEndpoints []string `pulumi:"serviceEndpoints"`
	// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
	VirtualNetworkName *string `pulumi:"virtualNetworkName"`
}

type SubnetState struct {
	// The address prefix to use for the subnet.
	AddressPrefix pulumi.StringPtrInput
	// One or more `delegation` blocks as defined below.
	Delegations SubnetDelegationArrayInput
	// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
	EnforcePrivateLinkEndpointNetworkPolicies pulumi.BoolPtrInput
	// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
	EnforcePrivateLinkServiceNetworkPolicies pulumi.BoolPtrInput
	// The name of the subnet. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
	ServiceEndpoints pulumi.StringArrayInput
	// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
	VirtualNetworkName pulumi.StringPtrInput
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// The address prefix to use for the subnet.
	AddressPrefix string `pulumi:"addressPrefix"`
	// One or more `delegation` blocks as defined below.
	Delegations []SubnetDelegation `pulumi:"delegations"`
	// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
	EnforcePrivateLinkEndpointNetworkPolicies *bool `pulumi:"enforcePrivateLinkEndpointNetworkPolicies"`
	// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
	EnforcePrivateLinkServiceNetworkPolicies *bool `pulumi:"enforcePrivateLinkServiceNetworkPolicies"`
	// The name of the subnet. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
	ServiceEndpoints []string `pulumi:"serviceEndpoints"`
	// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// The address prefix to use for the subnet.
	AddressPrefix pulumi.StringInput
	// One or more `delegation` blocks as defined below.
	Delegations SubnetDelegationArrayInput
	// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
	EnforcePrivateLinkEndpointNetworkPolicies pulumi.BoolPtrInput
	// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforcePrivateLinkEndpointNetworkPolicies`.
	EnforcePrivateLinkServiceNetworkPolicies pulumi.BoolPtrInput
	// The name of the subnet. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
	ServiceEndpoints pulumi.StringArrayInput
	// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
	VirtualNetworkName pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

