// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Private Endpoint.
//
// Azure Private Endpoint is a network interface that connects you privately and securely to a service powered by Azure Private Link. Private Endpoint uses a private IP address from your VNet, effectively bringing the service into your VNet. The service could be an Azure service such as Azure Storage, SQL, etc. or your own Private Link Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/lb"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatedns"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatelink"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		service, err := network.NewSubnet(ctx, "service", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 			EnforcePrivateLinkServiceNetworkPolicies: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		endpoint, err := network.NewSubnet(ctx, "endpoint", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 			EnforcePrivateLinkEndpointNetworkPolicies: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Sku:               pulumi.String("Standard"),
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Sku:               pulumi.String("Standard"),
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FrontendIpConfigurations: lb.LoadBalancerFrontendIpConfigurationArray{
// 				&lb.LoadBalancerFrontendIpConfigurationArgs{
// 					Name:              examplePublicIp.Name,
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLinkService, err := privatedns.NewLinkService(ctx, "exampleLinkService", &privatedns.LinkServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NatIpConfigurations: privatedns.LinkServiceNatIpConfigurationArray{
// 				&privatedns.LinkServiceNatIpConfigurationArgs{
// 					Name:     examplePublicIp.Name,
// 					Primary:  pulumi.Bool(true),
// 					SubnetId: service.ID(),
// 				},
// 			},
// 			LoadBalancerFrontendIpConfigurationIds: pulumi.StringArray{
// 				pulumi.String(exampleLoadBalancer.FrontendIpConfigurations.ApplyT(func(frontendIpConfigurations []lb.LoadBalancerFrontendIpConfiguration) (string, error) {
// 					return frontendIpConfigurations[0].Id, nil
// 				}).(pulumi.StringOutput)),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatelink.NewEndpoint(ctx, "exampleEndpoint", &privatelink.EndpointArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SubnetId:          endpoint.ID(),
// 			PrivateServiceConnection: &privatelink.EndpointPrivateServiceConnectionArgs{
// 				Name:                        pulumi.String("example-privateserviceconnection"),
// 				PrivateConnectionResourceId: exampleLinkService.ID(),
// 				IsManualConnection:          pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Using a Private Link Service Alias with existing resources:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatelink"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vnet, err := network.LookupVirtualNetwork(ctx, &network.LookupVirtualNetworkArgs{
// 			Name:              "example-network",
// 			ResourceGroupName: rg.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		subnet, err := network.LookupSubnet(ctx, &network.LookupSubnetArgs{
// 			Name:               "default",
// 			VirtualNetworkName: vnet.Name,
// 			ResourceGroupName:  rg.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatelink.NewEndpoint(ctx, "example", &privatelink.EndpointArgs{
// 			Location:          pulumi.String(rg.Location),
// 			ResourceGroupName: pulumi.String(rg.Name),
// 			SubnetId:          pulumi.String(subnet.Id),
// 			PrivateServiceConnection: &privatelink.EndpointPrivateServiceConnectionArgs{
// 				Name:                           pulumi.String("example-privateserviceconnection"),
// 				PrivateConnectionResourceAlias: pulumi.String("example-privatelinkservice.d20286c8-4ea5-11eb-9584-8f53157226c6.centralus.azure.privatelinkservice"),
// 				IsManualConnection:             pulumi.Bool(true),
// 				RequestMessage:                 pulumi.String("PL"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Private Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatelink/endpoint:Endpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/privateEndpoints/endpoint1
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	CustomDnsConfigs EndpointCustomDnsConfigArrayOutput `pulumi:"customDnsConfigs"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name                  pulumi.StringOutput                     `pulumi:"name"`
	PrivateDnsZoneConfigs EndpointPrivateDnsZoneConfigArrayOutput `pulumi:"privateDnsZoneConfigs"`
	// A `privateDnsZoneGroup` block as defined below.
	PrivateDnsZoneGroup EndpointPrivateDnsZoneGroupPtrOutput `pulumi:"privateDnsZoneGroup"`
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnectionOutput `pulumi:"privateServiceConnection"`
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateServiceConnection == nil {
		return nil, errors.New("invalid value for required argument 'PrivateServiceConnection'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("azure:privatelink/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure:privatelink/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	CustomDnsConfigs []EndpointCustomDnsConfig `pulumi:"customDnsConfigs"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name                  *string                        `pulumi:"name"`
	PrivateDnsZoneConfigs []EndpointPrivateDnsZoneConfig `pulumi:"privateDnsZoneConfigs"`
	// A `privateDnsZoneGroup` block as defined below.
	PrivateDnsZoneGroup *EndpointPrivateDnsZoneGroup `pulumi:"privateDnsZoneGroup"`
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection *EndpointPrivateServiceConnection `pulumi:"privateServiceConnection"`
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type EndpointState struct {
	CustomDnsConfigs EndpointCustomDnsConfigArrayInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name                  pulumi.StringPtrInput
	PrivateDnsZoneConfigs EndpointPrivateDnsZoneConfigArrayInput
	// A `privateDnsZoneGroup` block as defined below.
	PrivateDnsZoneGroup EndpointPrivateDnsZoneGroupPtrInput
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnectionPtrInput
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `privateDnsZoneGroup` block as defined below.
	PrivateDnsZoneGroup *EndpointPrivateDnsZoneGroup `pulumi:"privateDnsZoneGroup"`
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnection `pulumi:"privateServiceConnection"`
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `privateDnsZoneGroup` block as defined below.
	PrivateDnsZoneGroup EndpointPrivateDnsZoneGroupPtrInput
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnectionInput
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

func (i *Endpoint) ToEndpointPtrOutput() EndpointPtrOutput {
	return i.ToEndpointPtrOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPtrOutput)
}

type EndpointPtrInput interface {
	pulumi.Input

	ToEndpointPtrOutput() EndpointPtrOutput
	ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput
}

type endpointPtrType EndpointArgs

func (*endpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil))
}

func (i *endpointPtrType) ToEndpointPtrOutput() EndpointPtrOutput {
	return i.ToEndpointPtrOutputWithContext(context.Background())
}

func (i *endpointPtrType) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPtrOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//          EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Endpoint)(nil))
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//          EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Endpoint)(nil))
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct {
	*pulumi.OutputState
}

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointPtrOutput() EndpointPtrOutput {
	return o.ToEndpointPtrOutputWithContext(context.Background())
}

func (o EndpointOutput) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return o.ApplyT(func(v Endpoint) *Endpoint {
		return &v
	}).(EndpointPtrOutput)
}

type EndpointPtrOutput struct {
	*pulumi.OutputState
}

func (EndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil))
}

func (o EndpointPtrOutput) ToEndpointPtrOutput() EndpointPtrOutput {
	return o
}

func (o EndpointPtrOutput) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return o
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoint)(nil))
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Endpoint {
		return vs[0].([]Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Endpoint)(nil))
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Endpoint {
		return vs[0].(map[string]Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointPtrOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
