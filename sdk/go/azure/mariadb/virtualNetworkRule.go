// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a MariaDB Virtual Network Rule.
//
// > **NOTE:** MariaDB Virtual Network Rules [can only be used with SKU Tiers of `GeneralPurpose` or `MemoryOptimized`](https://docs.microsoft.com/en-us/azure/mariadb/concepts-data-access-security-vnet)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mariadb"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 				pulumi.String("10.7.29.0/29"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		internal, err := network.NewSubnet(ctx, "internal", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.7.29.0/29"),
// 			},
// 			ServiceEndpoints: pulumi.StringArray{
// 				pulumi.String("Microsoft.Sql"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServer, err := mariadb.NewServer(ctx, "exampleServer", &mariadb.ServerArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			AdministratorLogin:         pulumi.String("mariadbadminun"),
// 			AdministratorLoginPassword: pulumi.String("H@Sh1CoR3!"),
// 			Version:                    pulumi.String("5.7"),
// 			SslEnforcement:             pulumi.String("Enabled"),
// 			SkuName:                    pulumi.String("GP_Gen5_2"),
// 			StorageProfile: &mariadb.ServerStorageProfileArgs{
// 				StorageMb:           pulumi.Int(5120),
// 				BackupRetentionDays: pulumi.Int(7),
// 				GeoRedundantBackup:  pulumi.String("Disabled"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mariadb.NewVirtualNetworkRule(ctx, "exampleVirtualNetworkRule", &mariadb.VirtualNetworkRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			SubnetId:          internal.ID(),
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
// MariaDB Virtual Network Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mariadb/virtualNetworkRule:VirtualNetworkRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/MariaDB/servers/myserver/virtualNetworkRules/vnetrulename
// ```
type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	// The name of the MariaDB Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group where the MariaDB server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this MariaDB virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The ID of the subnet that the MariaDB server will be connected to.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource VirtualNetworkRule
	err := ctx.RegisterResource("azure:mariadb/virtualNetworkRule:VirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkRule gets an existing VirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	var resource VirtualNetworkRule
	err := ctx.ReadResource("azure:mariadb/virtualNetworkRule:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type virtualNetworkRuleState struct {
	// The name of the MariaDB Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group where the MariaDB server resides. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this MariaDB virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// The ID of the subnet that the MariaDB server will be connected to.
	SubnetId *string `pulumi:"subnetId"`
}

type VirtualNetworkRuleState struct {
	// The name of the MariaDB Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group where the MariaDB server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the SQL Server to which this MariaDB virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// The ID of the subnet that the MariaDB server will be connected to.
	SubnetId pulumi.StringPtrInput
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	// The name of the MariaDB Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group where the MariaDB server resides. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server to which this MariaDB virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// The ID of the subnet that the MariaDB server will be connected to.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// The name of the MariaDB Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group where the MariaDB server resides. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the SQL Server to which this MariaDB virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// The ID of the subnet that the MariaDB server will be connected to.
	SubnetId pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}

type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput
}

func (*VirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil))
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}

func (i *VirtualNetworkRule) ToVirtualNetworkRulePtrOutput() VirtualNetworkRulePtrOutput {
	return i.ToVirtualNetworkRulePtrOutputWithContext(context.Background())
}

func (i *VirtualNetworkRule) ToVirtualNetworkRulePtrOutputWithContext(ctx context.Context) VirtualNetworkRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRulePtrOutput)
}

type VirtualNetworkRulePtrInput interface {
	pulumi.Input

	ToVirtualNetworkRulePtrOutput() VirtualNetworkRulePtrOutput
	ToVirtualNetworkRulePtrOutputWithContext(ctx context.Context) VirtualNetworkRulePtrOutput
}

type virtualNetworkRulePtrType VirtualNetworkRuleArgs

func (*virtualNetworkRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRule)(nil))
}

func (i *virtualNetworkRulePtrType) ToVirtualNetworkRulePtrOutput() VirtualNetworkRulePtrOutput {
	return i.ToVirtualNetworkRulePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkRulePtrType) ToVirtualNetworkRulePtrOutputWithContext(ctx context.Context) VirtualNetworkRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRulePtrOutput)
}

// VirtualNetworkRuleArrayInput is an input type that accepts VirtualNetworkRuleArray and VirtualNetworkRuleArrayOutput values.
// You can construct a concrete instance of `VirtualNetworkRuleArrayInput` via:
//
//          VirtualNetworkRuleArray{ VirtualNetworkRuleArgs{...} }
type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VirtualNetworkRule)(nil))
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

// VirtualNetworkRuleMapInput is an input type that accepts VirtualNetworkRuleMap and VirtualNetworkRuleMapOutput values.
// You can construct a concrete instance of `VirtualNetworkRuleMapInput` via:
//
//          VirtualNetworkRuleMap{ "key": VirtualNetworkRuleArgs{...} }
type VirtualNetworkRuleMapInput interface {
	pulumi.Input

	ToVirtualNetworkRuleMapOutput() VirtualNetworkRuleMapOutput
	ToVirtualNetworkRuleMapOutputWithContext(context.Context) VirtualNetworkRuleMapOutput
}

type VirtualNetworkRuleMap map[string]VirtualNetworkRuleInput

func (VirtualNetworkRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VirtualNetworkRule)(nil))
}

func (i VirtualNetworkRuleMap) ToVirtualNetworkRuleMapOutput() VirtualNetworkRuleMapOutput {
	return i.ToVirtualNetworkRuleMapOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleMap) ToVirtualNetworkRuleMapOutputWithContext(ctx context.Context) VirtualNetworkRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleMapOutput)
}

type VirtualNetworkRuleOutput struct {
	*pulumi.OutputState
}

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil))
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRulePtrOutput() VirtualNetworkRulePtrOutput {
	return o.ToVirtualNetworkRulePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRulePtrOutputWithContext(ctx context.Context) VirtualNetworkRulePtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *VirtualNetworkRule {
		return &v
	}).(VirtualNetworkRulePtrOutput)
}

type VirtualNetworkRulePtrOutput struct {
	*pulumi.OutputState
}

func (VirtualNetworkRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRule)(nil))
}

func (o VirtualNetworkRulePtrOutput) ToVirtualNetworkRulePtrOutput() VirtualNetworkRulePtrOutput {
	return o
}

func (o VirtualNetworkRulePtrOutput) ToVirtualNetworkRulePtrOutputWithContext(ctx context.Context) VirtualNetworkRulePtrOutput {
	return o
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil))
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleMapOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualNetworkRule)(nil))
}

func (o VirtualNetworkRuleMapOutput) ToVirtualNetworkRuleMapOutput() VirtualNetworkRuleMapOutput {
	return o
}

func (o VirtualNetworkRuleMapOutput) ToVirtualNetworkRuleMapOutputWithContext(ctx context.Context) VirtualNetworkRuleMapOutput {
	return o
}

func (o VirtualNetworkRuleMapOutput) MapIndex(k pulumi.StringInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].(map[string]VirtualNetworkRule)[vs[1].(string)]
	}).(VirtualNetworkRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRulePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleMapOutput{})
}
