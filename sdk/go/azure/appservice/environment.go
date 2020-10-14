// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an App Service Environment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("westeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ase, err := network.NewSubnet(ctx, "ase", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnet(ctx, "gateway", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewEnvironment(ctx, "exampleEnvironment", &appservice.EnvironmentArgs{
// 			SubnetId:                  ase.ID(),
// 			PricingTier:               pulumi.String("I2"),
// 			FrontEndScaleFactor:       pulumi.Int(10),
// 			InternalLoadBalancingMode: pulumi.String("Web, Publishing"),
// 			AllowedUserIpCidrs: pulumi.StringArray{
// 				pulumi.String("11.22.33.44/32"),
// 				pulumi.String("55.66.77.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Environment struct {
	pulumi.CustomResourceState

	// Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
	AllowedUserIpCidrs pulumi.StringArrayOutput `pulumi:"allowedUserIpCidrs"`
	// Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
	FrontEndScaleFactor pulumi.IntPtrOutput `pulumi:"frontEndScaleFactor"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
	InternalLoadBalancingMode pulumi.StringPtrOutput `pulumi:"internalLoadBalancingMode"`
	// The location where the App Service Environment exists.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
	PricingTier pulumi.StringPtrOutput `pulumi:"pricingTier"`
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format
	UserWhitelistedIpRanges pulumi.StringArrayOutput `pulumi:"userWhitelistedIpRanges"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &EnvironmentArgs{}
	}
	var resource Environment
	err := ctx.RegisterResource("azure:appservice/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azure:appservice/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
	AllowedUserIpCidrs []string `pulumi:"allowedUserIpCidrs"`
	// Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
	FrontEndScaleFactor *int `pulumi:"frontEndScaleFactor"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// The location where the App Service Environment exists.
	Location *string `pulumi:"location"`
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
	PricingTier *string `pulumi:"pricingTier"`
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format
	UserWhitelistedIpRanges []string `pulumi:"userWhitelistedIpRanges"`
}

type EnvironmentState struct {
	// Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
	AllowedUserIpCidrs pulumi.StringArrayInput
	// Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
	FrontEndScaleFactor pulumi.IntPtrInput
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
	InternalLoadBalancingMode pulumi.StringPtrInput
	// The location where the App Service Environment exists.
	Location pulumi.StringPtrInput
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
	PricingTier pulumi.StringPtrInput
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
	// Deprecated: this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format
	UserWhitelistedIpRanges pulumi.StringArrayInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
	AllowedUserIpCidrs []string `pulumi:"allowedUserIpCidrs"`
	// Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
	FrontEndScaleFactor *int `pulumi:"frontEndScaleFactor"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
	PricingTier *string `pulumi:"pricingTier"`
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format
	UserWhitelistedIpRanges []string `pulumi:"userWhitelistedIpRanges"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Allowed user added IP ranges on the ASE database. Use the addresses you want to set as the explicit egress address ranges.
	AllowedUserIpCidrs pulumi.StringArrayInput
	// Scale factor for front end instances. Possible values are between `5` and `15`. Defaults to `15`.
	FrontEndScaleFactor pulumi.IntPtrInput
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None`, `Web`, `Publishing` and combined value `"Web, Publishing"`. Defaults to `None`.
	InternalLoadBalancingMode pulumi.StringPtrInput
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Pricing tier for the front end instances. Possible values are `I1`, `I2` and `I3`. Defaults to `I1`.
	PricingTier pulumi.StringPtrInput
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
	// Deprecated: this property has been renamed to `allowed_user_ip_cidrs` better reflect the expected ip range format
	UserWhitelistedIpRanges pulumi.StringArrayInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}
