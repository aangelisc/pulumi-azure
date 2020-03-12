// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an ExpressRoute circuit.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/express_route_circuit.html.markdown.
type ExpressRouteCircuit struct {
	pulumi.CustomResourceState

	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations pulumi.BoolPtrOutput `pulumi:"allowClassicOperations"`
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps pulumi.IntOutput `pulumi:"bandwidthInMbps"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation pulumi.StringOutput `pulumi:"peeringLocation"`
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey pulumi.StringOutput `pulumi:"serviceKey"`
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName pulumi.StringOutput `pulumi:"serviceProviderName"`
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState pulumi.StringOutput `pulumi:"serviceProviderProvisioningState"`
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewExpressRouteCircuit registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuit(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	if args == nil || args.BandwidthInMbps == nil {
		return nil, errors.New("missing required argument 'BandwidthInMbps'")
	}
	if args == nil || args.PeeringLocation == nil {
		return nil, errors.New("missing required argument 'PeeringLocation'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceProviderName == nil {
		return nil, errors.New("missing required argument 'ServiceProviderName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ExpressRouteCircuitArgs{}
	}
	var resource ExpressRouteCircuit
	err := ctx.RegisterResource("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuit gets an existing ExpressRouteCircuit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	var resource ExpressRouteCircuit
	err := ctx.ReadResource("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuit resources.
type expressRouteCircuitState struct {
	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps *int `pulumi:"bandwidthInMbps"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey *string `pulumi:"serviceKey"`
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName *string `pulumi:"serviceProviderName"`
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState *string `pulumi:"serviceProviderProvisioningState"`
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku *ExpressRouteCircuitSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ExpressRouteCircuitState struct {
	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations pulumi.BoolPtrInput
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps pulumi.IntPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation pulumi.StringPtrInput
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey pulumi.StringPtrInput
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName pulumi.StringPtrInput
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState pulumi.StringPtrInput
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ExpressRouteCircuitState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitState)(nil)).Elem()
}

type expressRouteCircuitArgs struct {
	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps int `pulumi:"bandwidthInMbps"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation string `pulumi:"peeringLocation"`
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName string `pulumi:"serviceProviderName"`
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ExpressRouteCircuit resource.
type ExpressRouteCircuitArgs struct {
	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations pulumi.BoolPtrInput
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps pulumi.IntInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation pulumi.StringInput
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName pulumi.StringInput
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ExpressRouteCircuitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitArgs)(nil)).Elem()
}

