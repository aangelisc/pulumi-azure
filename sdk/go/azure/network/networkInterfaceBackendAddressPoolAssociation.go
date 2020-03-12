// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages the association between a Network Interface and a Load Balancer's Backend Address Pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_interface_backend_address_pool_association.html.markdown.
type NetworkInterfaceBackendAddressPoolAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringOutput `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringOutput `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceBackendAddressPoolAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceBackendAddressPoolAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceBackendAddressPoolAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceBackendAddressPoolAssociation, error) {
	if args == nil || args.BackendAddressPoolId == nil {
		return nil, errors.New("missing required argument 'BackendAddressPoolId'")
	}
	if args == nil || args.IpConfigurationName == nil {
		return nil, errors.New("missing required argument 'IpConfigurationName'")
	}
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil {
		args = &NetworkInterfaceBackendAddressPoolAssociationArgs{}
	}
	var resource NetworkInterfaceBackendAddressPoolAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceBackendAddressPoolAssociation gets an existing NetworkInterfaceBackendAddressPoolAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceBackendAddressPoolAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceBackendAddressPoolAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceBackendAddressPoolAssociation, error) {
	var resource NetworkInterfaceBackendAddressPoolAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceBackendAddressPoolAssociation resources.
type networkInterfaceBackendAddressPoolAssociationState struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId *string `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName *string `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
}

type NetworkInterfaceBackendAddressPoolAssociationState struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringPtrInput
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringPtrInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
}

func (NetworkInterfaceBackendAddressPoolAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceBackendAddressPoolAssociationState)(nil)).Elem()
}

type networkInterfaceBackendAddressPoolAssociationArgs struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId string `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName string `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceBackendAddressPoolAssociation resource.
type NetworkInterfaceBackendAddressPoolAssociationArgs struct {
	// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringInput
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceBackendAddressPoolAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceBackendAddressPoolAssociationArgs)(nil)).Elem()
}

