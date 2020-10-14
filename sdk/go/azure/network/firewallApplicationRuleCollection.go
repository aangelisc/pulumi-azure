// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application Rule Collection within an Azure Firewall.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("North Europe"),
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
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFirewall, err := network.NewFirewall(ctx, "exampleFirewall", &network.FirewallArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.FirewallIpConfigurationArray{
// 				&network.FirewallIpConfigurationArgs{
// 					Name:              pulumi.String("configuration"),
// 					SubnetId:          exampleSubnet.ID(),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewFirewallApplicationRuleCollection(ctx, "exampleFirewallApplicationRuleCollection", &network.FirewallApplicationRuleCollectionArgs{
// 			AzureFirewallName: exampleFirewall.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Priority:          pulumi.Int(100),
// 			Action:            pulumi.String("Allow"),
// 			Rules: network.FirewallApplicationRuleCollectionRuleArray{
// 				&network.FirewallApplicationRuleCollectionRuleArgs{
// 					Name: pulumi.String("testrule"),
// 					SourceAddresses: pulumi.StringArray{
// 						pulumi.String("10.0.0.0/16"),
// 					},
// 					TargetFqdns: pulumi.StringArray{
// 						pulumi.String("*.google.com"),
// 					},
// 					Protocols: network.FirewallApplicationRuleCollectionRuleProtocolArray{
// 						&network.FirewallApplicationRuleCollectionRuleProtocolArgs{
// 							Port: pulumi.Int(443),
// 							Type: pulumi.String("Https"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FirewallApplicationRuleCollection struct {
	pulumi.CustomResourceState

	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringOutput `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRuleArrayOutput `pulumi:"rules"`
}

// NewFirewallApplicationRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewFirewallApplicationRuleCollection(ctx *pulumi.Context,
	name string, args *FirewallApplicationRuleCollectionArgs, opts ...pulumi.ResourceOption) (*FirewallApplicationRuleCollection, error) {
	if args == nil || args.Action == nil {
		return nil, errors.New("missing required argument 'Action'")
	}
	if args == nil || args.AzureFirewallName == nil {
		return nil, errors.New("missing required argument 'AzureFirewallName'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	if args == nil {
		args = &FirewallApplicationRuleCollectionArgs{}
	}
	var resource FirewallApplicationRuleCollection
	err := ctx.RegisterResource("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallApplicationRuleCollection gets an existing FirewallApplicationRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallApplicationRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallApplicationRuleCollectionState, opts ...pulumi.ResourceOption) (*FirewallApplicationRuleCollection, error) {
	var resource FirewallApplicationRuleCollection
	err := ctx.ReadResource("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallApplicationRuleCollection resources.
type firewallApplicationRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action *string `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName *string `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority *int `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules []FirewallApplicationRuleCollectionRule `pulumi:"rules"`
}

type FirewallApplicationRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringPtrInput
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringPtrInput
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntPtrInput
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRuleArrayInput
}

func (FirewallApplicationRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallApplicationRuleCollectionState)(nil)).Elem()
}

type firewallApplicationRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action string `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName string `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority int `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules []FirewallApplicationRuleCollectionRule `pulumi:"rules"`
}

// The set of arguments for constructing a FirewallApplicationRuleCollection resource.
type FirewallApplicationRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringInput
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringInput
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntInput
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRuleArrayInput
}

func (FirewallApplicationRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallApplicationRuleCollectionArgs)(nil)).Elem()
}
