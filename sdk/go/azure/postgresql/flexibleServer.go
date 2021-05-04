// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a PostgreSQL Flexible Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/postgresql"
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
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 			ServiceEndpoints: pulumi.StringArray{
// 				pulumi.String("Microsoft.Storage"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("fs"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.DBforPostgreSQL/flexibleServers"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/join/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = postgresql.NewFlexibleServer(ctx, "exampleFlexibleServer", &postgresql.FlexibleServerArgs{
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			Location:              exampleResourceGroup.Location,
// 			Version:               pulumi.String("12"),
// 			DelegatedSubnetId:     exampleSubnet.ID(),
// 			AdministratorLogin:    pulumi.String("psqladminun"),
// 			AdministratorPassword: pulumi.String("H@Sh1CoR3!"),
// 			StorageMb:             pulumi.Int(32768),
// 			SkuName:               pulumi.String("GP_Standard_D4s_v3"),
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
// PostgreSQL Flexible Servers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:postgresql/flexibleServer:FlexibleServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/server1
// ```
type FlexibleServer struct {
	pulumi.CustomResourceState

	// The Administrator Login for the PostgreSQL Flexible Server. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the PostgreSQL Flexible Server. Required when `createMode` is `Default`.
	AdministratorPassword pulumi.StringPtrOutput `pulumi:"administratorPassword"`
	// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntOutput `pulumi:"backupRetentionDays"`
	// The status showing whether the data encryption is enabled with a customer-managed key.
	CmkEnabled pulumi.StringOutput `pulumi:"cmkEnabled"`
	// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
	DelegatedSubnetId pulumi.StringPtrOutput `pulumi:"delegatedSubnetId"`
	// The FQDN of the PostgreSQL Flexible Server.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow FlexibleServerMaintenanceWindowPtrOutput `pulumi:"maintenanceWindow"`
	// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The point in time to restore from `creationSourceServerId` when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc pulumi.StringPtrOutput `pulumi:"pointInTimeRestoreTimeInUtc"`
	// Is public network access enabled?
	PublicNetworkAccessEnabled pulumi.BoolOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	SourceServerId pulumi.StringPtrOutput `pulumi:"sourceServerId"`
	// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
	StorageMb pulumi.IntOutput `pulumi:"storageMb"`
	// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The version of PostgreSQL Flexible Server to use. Possible values are `11` and `12`. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Version pulumi.StringOutput `pulumi:"version"`
	// The availability Zone of the PostgreSQL Flexible Server. Possible values are  `none`, `1`, `2` and `3`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewFlexibleServer registers a new resource with the given unique name, arguments, and options.
func NewFlexibleServer(ctx *pulumi.Context,
	name string, args *FlexibleServerArgs, opts ...pulumi.ResourceOption) (*FlexibleServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource FlexibleServer
	err := ctx.RegisterResource("azure:postgresql/flexibleServer:FlexibleServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleServer gets an existing FlexibleServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleServerState, opts ...pulumi.ResourceOption) (*FlexibleServer, error) {
	var resource FlexibleServer
	err := ctx.ReadResource("azure:postgresql/flexibleServer:FlexibleServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleServer resources.
type flexibleServerState struct {
	// The Administrator Login for the PostgreSQL Flexible Server. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the PostgreSQL Flexible Server. Required when `createMode` is `Default`.
	AdministratorPassword *string `pulumi:"administratorPassword"`
	// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// The status showing whether the data encryption is enabled with a customer-managed key.
	CmkEnabled *string `pulumi:"cmkEnabled"`
	// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	CreateMode *string `pulumi:"createMode"`
	// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
	DelegatedSubnetId *string `pulumi:"delegatedSubnetId"`
	// The FQDN of the PostgreSQL Flexible Server.
	Fqdn *string `pulumi:"fqdn"`
	// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	Location *string `pulumi:"location"`
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow *FlexibleServerMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
	Name *string `pulumi:"name"`
	// The point in time to restore from `creationSourceServerId` when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc *string `pulumi:"pointInTimeRestoreTimeInUtc"`
	// Is public network access enabled?
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
	SkuName *string `pulumi:"skuName"`
	// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	SourceServerId *string `pulumi:"sourceServerId"`
	// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
	StorageMb *int `pulumi:"storageMb"`
	// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
	Tags map[string]string `pulumi:"tags"`
	// The version of PostgreSQL Flexible Server to use. Possible values are `11` and `12`. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Version *string `pulumi:"version"`
	// The availability Zone of the PostgreSQL Flexible Server. Possible values are  `none`, `1`, `2` and `3`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Zone *string `pulumi:"zone"`
}

type FlexibleServerState struct {
	// The Administrator Login for the PostgreSQL Flexible Server. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the PostgreSQL Flexible Server. Required when `createMode` is `Default`.
	AdministratorPassword pulumi.StringPtrInput
	// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput
	// The status showing whether the data encryption is enabled with a customer-managed key.
	CmkEnabled pulumi.StringPtrInput
	// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	CreateMode pulumi.StringPtrInput
	// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
	DelegatedSubnetId pulumi.StringPtrInput
	// The FQDN of the PostgreSQL Flexible Server.
	Fqdn pulumi.StringPtrInput
	// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	Location pulumi.StringPtrInput
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow FlexibleServerMaintenanceWindowPtrInput
	// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
	Name pulumi.StringPtrInput
	// The point in time to restore from `creationSourceServerId` when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc pulumi.StringPtrInput
	// Is public network access enabled?
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
	SkuName pulumi.StringPtrInput
	// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	SourceServerId pulumi.StringPtrInput
	// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
	StorageMb pulumi.IntPtrInput
	// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
	Tags pulumi.StringMapInput
	// The version of PostgreSQL Flexible Server to use. Possible values are `11` and `12`. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Version pulumi.StringPtrInput
	// The availability Zone of the PostgreSQL Flexible Server. Possible values are  `none`, `1`, `2` and `3`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Zone pulumi.StringPtrInput
}

func (FlexibleServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerState)(nil)).Elem()
}

type flexibleServerArgs struct {
	// The Administrator Login for the PostgreSQL Flexible Server. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the PostgreSQL Flexible Server. Required when `createMode` is `Default`.
	AdministratorPassword *string `pulumi:"administratorPassword"`
	// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	CreateMode *string `pulumi:"createMode"`
	// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
	DelegatedSubnetId *string `pulumi:"delegatedSubnetId"`
	// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	Location *string `pulumi:"location"`
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow *FlexibleServerMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
	Name *string `pulumi:"name"`
	// The point in time to restore from `creationSourceServerId` when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc *string `pulumi:"pointInTimeRestoreTimeInUtc"`
	// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
	SkuName *string `pulumi:"skuName"`
	// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	SourceServerId *string `pulumi:"sourceServerId"`
	// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
	StorageMb *int `pulumi:"storageMb"`
	// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
	Tags map[string]string `pulumi:"tags"`
	// The version of PostgreSQL Flexible Server to use. Possible values are `11` and `12`. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Version *string `pulumi:"version"`
	// The availability Zone of the PostgreSQL Flexible Server. Possible values are  `none`, `1`, `2` and `3`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a FlexibleServer resource.
type FlexibleServerArgs struct {
	// The Administrator Login for the PostgreSQL Flexible Server. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the PostgreSQL Flexible Server. Required when `createMode` is `Default`.
	AdministratorPassword pulumi.StringPtrInput
	// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput
	// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	CreateMode pulumi.StringPtrInput
	// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
	DelegatedSubnetId pulumi.StringPtrInput
	// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	Location pulumi.StringPtrInput
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow FlexibleServerMaintenanceWindowPtrInput
	// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
	Name pulumi.StringPtrInput
	// The point in time to restore from `creationSourceServerId` when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc pulumi.StringPtrInput
	// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
	SkuName pulumi.StringPtrInput
	// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `createMode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
	SourceServerId pulumi.StringPtrInput
	// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
	StorageMb pulumi.IntPtrInput
	// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
	Tags pulumi.StringMapInput
	// The version of PostgreSQL Flexible Server to use. Possible values are `11` and `12`. Required when `createMode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Version pulumi.StringPtrInput
	// The availability Zone of the PostgreSQL Flexible Server. Possible values are  `none`, `1`, `2` and `3`. Changing this forces a new PostgreSQL Flexible Server to be created.
	Zone pulumi.StringPtrInput
}

func (FlexibleServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerArgs)(nil)).Elem()
}

type FlexibleServerInput interface {
	pulumi.Input

	ToFlexibleServerOutput() FlexibleServerOutput
	ToFlexibleServerOutputWithContext(ctx context.Context) FlexibleServerOutput
}

func (*FlexibleServer) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexibleServer)(nil))
}

func (i *FlexibleServer) ToFlexibleServerOutput() FlexibleServerOutput {
	return i.ToFlexibleServerOutputWithContext(context.Background())
}

func (i *FlexibleServer) ToFlexibleServerOutputWithContext(ctx context.Context) FlexibleServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerOutput)
}

func (i *FlexibleServer) ToFlexibleServerPtrOutput() FlexibleServerPtrOutput {
	return i.ToFlexibleServerPtrOutputWithContext(context.Background())
}

func (i *FlexibleServer) ToFlexibleServerPtrOutputWithContext(ctx context.Context) FlexibleServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerPtrOutput)
}

type FlexibleServerPtrInput interface {
	pulumi.Input

	ToFlexibleServerPtrOutput() FlexibleServerPtrOutput
	ToFlexibleServerPtrOutputWithContext(ctx context.Context) FlexibleServerPtrOutput
}

type flexibleServerPtrType FlexibleServerArgs

func (*flexibleServerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServer)(nil))
}

func (i *flexibleServerPtrType) ToFlexibleServerPtrOutput() FlexibleServerPtrOutput {
	return i.ToFlexibleServerPtrOutputWithContext(context.Background())
}

func (i *flexibleServerPtrType) ToFlexibleServerPtrOutputWithContext(ctx context.Context) FlexibleServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerPtrOutput)
}

// FlexibleServerArrayInput is an input type that accepts FlexibleServerArray and FlexibleServerArrayOutput values.
// You can construct a concrete instance of `FlexibleServerArrayInput` via:
//
//          FlexibleServerArray{ FlexibleServerArgs{...} }
type FlexibleServerArrayInput interface {
	pulumi.Input

	ToFlexibleServerArrayOutput() FlexibleServerArrayOutput
	ToFlexibleServerArrayOutputWithContext(context.Context) FlexibleServerArrayOutput
}

type FlexibleServerArray []FlexibleServerInput

func (FlexibleServerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FlexibleServer)(nil))
}

func (i FlexibleServerArray) ToFlexibleServerArrayOutput() FlexibleServerArrayOutput {
	return i.ToFlexibleServerArrayOutputWithContext(context.Background())
}

func (i FlexibleServerArray) ToFlexibleServerArrayOutputWithContext(ctx context.Context) FlexibleServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerArrayOutput)
}

// FlexibleServerMapInput is an input type that accepts FlexibleServerMap and FlexibleServerMapOutput values.
// You can construct a concrete instance of `FlexibleServerMapInput` via:
//
//          FlexibleServerMap{ "key": FlexibleServerArgs{...} }
type FlexibleServerMapInput interface {
	pulumi.Input

	ToFlexibleServerMapOutput() FlexibleServerMapOutput
	ToFlexibleServerMapOutputWithContext(context.Context) FlexibleServerMapOutput
}

type FlexibleServerMap map[string]FlexibleServerInput

func (FlexibleServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FlexibleServer)(nil))
}

func (i FlexibleServerMap) ToFlexibleServerMapOutput() FlexibleServerMapOutput {
	return i.ToFlexibleServerMapOutputWithContext(context.Background())
}

func (i FlexibleServerMap) ToFlexibleServerMapOutputWithContext(ctx context.Context) FlexibleServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerMapOutput)
}

type FlexibleServerOutput struct {
	*pulumi.OutputState
}

func (FlexibleServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexibleServer)(nil))
}

func (o FlexibleServerOutput) ToFlexibleServerOutput() FlexibleServerOutput {
	return o
}

func (o FlexibleServerOutput) ToFlexibleServerOutputWithContext(ctx context.Context) FlexibleServerOutput {
	return o
}

func (o FlexibleServerOutput) ToFlexibleServerPtrOutput() FlexibleServerPtrOutput {
	return o.ToFlexibleServerPtrOutputWithContext(context.Background())
}

func (o FlexibleServerOutput) ToFlexibleServerPtrOutputWithContext(ctx context.Context) FlexibleServerPtrOutput {
	return o.ApplyT(func(v FlexibleServer) *FlexibleServer {
		return &v
	}).(FlexibleServerPtrOutput)
}

type FlexibleServerPtrOutput struct {
	*pulumi.OutputState
}

func (FlexibleServerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServer)(nil))
}

func (o FlexibleServerPtrOutput) ToFlexibleServerPtrOutput() FlexibleServerPtrOutput {
	return o
}

func (o FlexibleServerPtrOutput) ToFlexibleServerPtrOutputWithContext(ctx context.Context) FlexibleServerPtrOutput {
	return o
}

type FlexibleServerArrayOutput struct{ *pulumi.OutputState }

func (FlexibleServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FlexibleServer)(nil))
}

func (o FlexibleServerArrayOutput) ToFlexibleServerArrayOutput() FlexibleServerArrayOutput {
	return o
}

func (o FlexibleServerArrayOutput) ToFlexibleServerArrayOutputWithContext(ctx context.Context) FlexibleServerArrayOutput {
	return o
}

func (o FlexibleServerArrayOutput) Index(i pulumi.IntInput) FlexibleServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FlexibleServer {
		return vs[0].([]FlexibleServer)[vs[1].(int)]
	}).(FlexibleServerOutput)
}

type FlexibleServerMapOutput struct{ *pulumi.OutputState }

func (FlexibleServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlexibleServer)(nil))
}

func (o FlexibleServerMapOutput) ToFlexibleServerMapOutput() FlexibleServerMapOutput {
	return o
}

func (o FlexibleServerMapOutput) ToFlexibleServerMapOutputWithContext(ctx context.Context) FlexibleServerMapOutput {
	return o
}

func (o FlexibleServerMapOutput) MapIndex(k pulumi.StringInput) FlexibleServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlexibleServer {
		return vs[0].(map[string]FlexibleServer)[vs[1].(string)]
	}).(FlexibleServerOutput)
}

func init() {
	pulumi.RegisterOutputType(FlexibleServerOutput{})
	pulumi.RegisterOutputType(FlexibleServerPtrOutput{})
	pulumi.RegisterOutputType(FlexibleServerArrayOutput{})
	pulumi.RegisterOutputType(FlexibleServerMapOutput{})
}
