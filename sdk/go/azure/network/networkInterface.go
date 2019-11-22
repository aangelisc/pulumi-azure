// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Network Interface located in a Virtual Network, usually attached to a Virtual Machine.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_interface.html.markdown.
type NetworkInterface struct {
	s *pulumi.ResourceState
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOpt) (*NetworkInterface, error) {
	if args == nil || args.IpConfigurations == nil {
		return nil, errors.New("missing required argument 'IpConfigurations'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appliedDnsServers"] = nil
		inputs["dnsServers"] = nil
		inputs["enableAcceleratedNetworking"] = nil
		inputs["enableIpForwarding"] = nil
		inputs["internalDnsNameLabel"] = nil
		inputs["internalFqdn"] = nil
		inputs["ipConfigurations"] = nil
		inputs["location"] = nil
		inputs["macAddress"] = nil
		inputs["name"] = nil
		inputs["networkSecurityGroupId"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["virtualMachineId"] = nil
	} else {
		inputs["appliedDnsServers"] = args.AppliedDnsServers
		inputs["dnsServers"] = args.DnsServers
		inputs["enableAcceleratedNetworking"] = args.EnableAcceleratedNetworking
		inputs["enableIpForwarding"] = args.EnableIpForwarding
		inputs["internalDnsNameLabel"] = args.InternalDnsNameLabel
		inputs["internalFqdn"] = args.InternalFqdn
		inputs["ipConfigurations"] = args.IpConfigurations
		inputs["location"] = args.Location
		inputs["macAddress"] = args.MacAddress
		inputs["name"] = args.Name
		inputs["networkSecurityGroupId"] = args.NetworkSecurityGroupId
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["virtualMachineId"] = args.VirtualMachineId
	}
	inputs["privateIpAddress"] = nil
	inputs["privateIpAddresses"] = nil
	s, err := ctx.RegisterResource("azure:network/networkInterface:NetworkInterface", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkInterface{s: s}, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkInterfaceState, opts ...pulumi.ResourceOpt) (*NetworkInterface, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appliedDnsServers"] = state.AppliedDnsServers
		inputs["dnsServers"] = state.DnsServers
		inputs["enableAcceleratedNetworking"] = state.EnableAcceleratedNetworking
		inputs["enableIpForwarding"] = state.EnableIpForwarding
		inputs["internalDnsNameLabel"] = state.InternalDnsNameLabel
		inputs["internalFqdn"] = state.InternalFqdn
		inputs["ipConfigurations"] = state.IpConfigurations
		inputs["location"] = state.Location
		inputs["macAddress"] = state.MacAddress
		inputs["name"] = state.Name
		inputs["networkSecurityGroupId"] = state.NetworkSecurityGroupId
		inputs["privateIpAddress"] = state.PrivateIpAddress
		inputs["privateIpAddresses"] = state.PrivateIpAddresses
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["virtualMachineId"] = state.VirtualMachineId
	}
	s, err := ctx.ReadResource("azure:network/networkInterface:NetworkInterface", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkInterface{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NetworkInterface) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NetworkInterface) ID() pulumi.IDOutput {
	return r.s.ID()
}

// If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set
func (r *NetworkInterface) AppliedDnsServers() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["appliedDnsServers"])
}

// List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list
func (r *NetworkInterface) DnsServers() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["dnsServers"])
}

// Enables Azure Accelerated Networking using SR-IOV. Only certain VM instance sizes are supported. Refer to [Create a Virtual Machine with Accelerated Networking](https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli). Defaults to `false`.
func (r *NetworkInterface) EnableAcceleratedNetworking() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableAcceleratedNetworking"])
}

// Enables IP Forwarding on the NIC. Defaults to `false`.
func (r *NetworkInterface) EnableIpForwarding() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableIpForwarding"])
}

// Relative DNS name for this NIC used for internal communications between VMs in the same VNet
func (r *NetworkInterface) InternalDnsNameLabel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internalDnsNameLabel"])
}

func (r *NetworkInterface) InternalFqdn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internalFqdn"])
}

// One or more `ipConfiguration` associated with this NIC as documented below.
func (r *NetworkInterface) IpConfigurations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ipConfigurations"])
}

// The location/region where the network interface is created. Changing this forces a new resource to be created.
func (r *NetworkInterface) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The media access control (MAC) address of the network interface.
func (r *NetworkInterface) MacAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["macAddress"])
}

// The name of the network interface. Changing this forces a new resource to be created.
func (r *NetworkInterface) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the Network Security Group to associate with the network interface.
func (r *NetworkInterface) NetworkSecurityGroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["networkSecurityGroupId"])
}

// The first private IP address of the network interface.
func (r *NetworkInterface) PrivateIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["privateIpAddress"])
}

// The private IP addresses of the network interface.
func (r *NetworkInterface) PrivateIpAddresses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["privateIpAddresses"])
}

// The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.
func (r *NetworkInterface) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *NetworkInterface) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Reference to a VM with which this NIC has been associated.
func (r *NetworkInterface) VirtualMachineId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["virtualMachineId"])
}

// Input properties used for looking up and filtering NetworkInterface resources.
type NetworkInterfaceState struct {
	// If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set
	AppliedDnsServers interface{}
	// List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list
	DnsServers interface{}
	// Enables Azure Accelerated Networking using SR-IOV. Only certain VM instance sizes are supported. Refer to [Create a Virtual Machine with Accelerated Networking](https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli). Defaults to `false`.
	EnableAcceleratedNetworking interface{}
	// Enables IP Forwarding on the NIC. Defaults to `false`.
	EnableIpForwarding interface{}
	// Relative DNS name for this NIC used for internal communications between VMs in the same VNet
	InternalDnsNameLabel interface{}
	InternalFqdn interface{}
	// One or more `ipConfiguration` associated with this NIC as documented below.
	IpConfigurations interface{}
	// The location/region where the network interface is created. Changing this forces a new resource to be created.
	Location interface{}
	// The media access control (MAC) address of the network interface.
	MacAddress interface{}
	// The name of the network interface. Changing this forces a new resource to be created.
	Name interface{}
	// The ID of the Network Security Group to associate with the network interface.
	NetworkSecurityGroupId interface{}
	// The first private IP address of the network interface.
	PrivateIpAddress interface{}
	// The private IP addresses of the network interface.
	PrivateIpAddresses interface{}
	// The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Reference to a VM with which this NIC has been associated.
	VirtualMachineId interface{}
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set
	AppliedDnsServers interface{}
	// List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list
	DnsServers interface{}
	// Enables Azure Accelerated Networking using SR-IOV. Only certain VM instance sizes are supported. Refer to [Create a Virtual Machine with Accelerated Networking](https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli). Defaults to `false`.
	EnableAcceleratedNetworking interface{}
	// Enables IP Forwarding on the NIC. Defaults to `false`.
	EnableIpForwarding interface{}
	// Relative DNS name for this NIC used for internal communications between VMs in the same VNet
	InternalDnsNameLabel interface{}
	InternalFqdn interface{}
	// One or more `ipConfiguration` associated with this NIC as documented below.
	IpConfigurations interface{}
	// The location/region where the network interface is created. Changing this forces a new resource to be created.
	Location interface{}
	// The media access control (MAC) address of the network interface.
	MacAddress interface{}
	// The name of the network interface. Changing this forces a new resource to be created.
	Name interface{}
	// The ID of the Network Security Group to associate with the network interface.
	NetworkSecurityGroupId interface{}
	// The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Reference to a VM with which this NIC has been associated.
	VirtualMachineId interface{}
}
