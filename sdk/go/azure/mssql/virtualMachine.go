// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Microsoft SQL Virtual Machine
//
// ## Example Usage
//
// This example provisions a brief Managed MsSql Virtual Machine.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVirtualMachine, err := compute.LookupVirtualMachine(ctx, &compute.LookupVirtualMachineArgs{
// 			Name:              "example-vm",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewVirtualMachine(ctx, "exampleMssql_virtualMachineVirtualMachine", &mssql.VirtualMachineArgs{
// 			VirtualMachineId:              pulumi.String(exampleVirtualMachine.Id),
// 			SqlLicenseType:                pulumi.String("PAYG"),
// 			RServicesEnabled:              pulumi.Bool(true),
// 			SqlConnectivityPort:           pulumi.Int(1433),
// 			SqlConnectivityType:           pulumi.String("PRIVATE"),
// 			SqlConnectivityUpdatePassword: pulumi.String("Password1234!"),
// 			SqlConnectivityUpdateUsername: pulumi.String("sqllogin"),
// 			AutoPatching: &mssql.VirtualMachineAutoPatchingArgs{
// 				DayOfWeek:                          pulumi.String("Sunday"),
// 				MaintenanceWindowDurationInMinutes: pulumi.Int(60),
// 				MaintenanceWindowStartingHour:      pulumi.Int(2),
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
// Sql Virtual Machines can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mssql/virtualMachine:VirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
// ```
type VirtualMachine struct {
	pulumi.CustomResourceState

	// An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
	AutoBackup VirtualMachineAutoBackupPtrOutput `pulumi:"autoBackup"`
	// An `autoPatching` block as defined below.
	AutoPatching VirtualMachineAutoPatchingPtrOutput `pulumi:"autoPatching"`
	// (Optional) An `keyVaultCredential` block as defined below.
	KeyVaultCredential VirtualMachineKeyVaultCredentialPtrOutput `pulumi:"keyVaultCredential"`
	// Should R Services be enabled?
	RServicesEnabled pulumi.BoolPtrOutput `pulumi:"rServicesEnabled"`
	// The SQL Server port. Defaults to `1433`.
	SqlConnectivityPort pulumi.IntPtrOutput `pulumi:"sqlConnectivityPort"`
	// The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
	SqlConnectivityType pulumi.StringPtrOutput `pulumi:"sqlConnectivityType"`
	// The SQL Server sysadmin login password.
	SqlConnectivityUpdatePassword pulumi.StringPtrOutput `pulumi:"sqlConnectivityUpdatePassword"`
	// The SQL Server sysadmin login to create.
	SqlConnectivityUpdateUsername pulumi.StringPtrOutput `pulumi:"sqlConnectivityUpdateUsername"`
	// The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
	SqlLicenseType pulumi.StringOutput `pulumi:"sqlLicenseType"`
	// An `storageConfiguration` block as defined below.
	StorageConfiguration VirtualMachineStorageConfigurationPtrOutput `pulumi:"storageConfiguration"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SqlLicenseType == nil {
		return nil, errors.New("invalid value for required argument 'SqlLicenseType'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	var resource VirtualMachine
	err := ctx.RegisterResource("azure:mssql/virtualMachine:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachine gets an existing VirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure:mssql/virtualMachine:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachine resources.
type virtualMachineState struct {
	// An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
	AutoBackup *VirtualMachineAutoBackup `pulumi:"autoBackup"`
	// An `autoPatching` block as defined below.
	AutoPatching *VirtualMachineAutoPatching `pulumi:"autoPatching"`
	// (Optional) An `keyVaultCredential` block as defined below.
	KeyVaultCredential *VirtualMachineKeyVaultCredential `pulumi:"keyVaultCredential"`
	// Should R Services be enabled?
	RServicesEnabled *bool `pulumi:"rServicesEnabled"`
	// The SQL Server port. Defaults to `1433`.
	SqlConnectivityPort *int `pulumi:"sqlConnectivityPort"`
	// The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
	SqlConnectivityType *string `pulumi:"sqlConnectivityType"`
	// The SQL Server sysadmin login password.
	SqlConnectivityUpdatePassword *string `pulumi:"sqlConnectivityUpdatePassword"`
	// The SQL Server sysadmin login to create.
	SqlConnectivityUpdateUsername *string `pulumi:"sqlConnectivityUpdateUsername"`
	// The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
	SqlLicenseType *string `pulumi:"sqlLicenseType"`
	// An `storageConfiguration` block as defined below.
	StorageConfiguration *VirtualMachineStorageConfiguration `pulumi:"storageConfiguration"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type VirtualMachineState struct {
	// An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
	AutoBackup VirtualMachineAutoBackupPtrInput
	// An `autoPatching` block as defined below.
	AutoPatching VirtualMachineAutoPatchingPtrInput
	// (Optional) An `keyVaultCredential` block as defined below.
	KeyVaultCredential VirtualMachineKeyVaultCredentialPtrInput
	// Should R Services be enabled?
	RServicesEnabled pulumi.BoolPtrInput
	// The SQL Server port. Defaults to `1433`.
	SqlConnectivityPort pulumi.IntPtrInput
	// The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
	SqlConnectivityType pulumi.StringPtrInput
	// The SQL Server sysadmin login password.
	SqlConnectivityUpdatePassword pulumi.StringPtrInput
	// The SQL Server sysadmin login to create.
	SqlConnectivityUpdateUsername pulumi.StringPtrInput
	// The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
	SqlLicenseType pulumi.StringPtrInput
	// An `storageConfiguration` block as defined below.
	StorageConfiguration VirtualMachineStorageConfigurationPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringPtrInput
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	// An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
	AutoBackup *VirtualMachineAutoBackup `pulumi:"autoBackup"`
	// An `autoPatching` block as defined below.
	AutoPatching *VirtualMachineAutoPatching `pulumi:"autoPatching"`
	// (Optional) An `keyVaultCredential` block as defined below.
	KeyVaultCredential *VirtualMachineKeyVaultCredential `pulumi:"keyVaultCredential"`
	// Should R Services be enabled?
	RServicesEnabled *bool `pulumi:"rServicesEnabled"`
	// The SQL Server port. Defaults to `1433`.
	SqlConnectivityPort *int `pulumi:"sqlConnectivityPort"`
	// The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
	SqlConnectivityType *string `pulumi:"sqlConnectivityType"`
	// The SQL Server sysadmin login password.
	SqlConnectivityUpdatePassword *string `pulumi:"sqlConnectivityUpdatePassword"`
	// The SQL Server sysadmin login to create.
	SqlConnectivityUpdateUsername *string `pulumi:"sqlConnectivityUpdateUsername"`
	// The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
	SqlLicenseType string `pulumi:"sqlLicenseType"`
	// An `storageConfiguration` block as defined below.
	StorageConfiguration *VirtualMachineStorageConfiguration `pulumi:"storageConfiguration"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
	AutoBackup VirtualMachineAutoBackupPtrInput
	// An `autoPatching` block as defined below.
	AutoPatching VirtualMachineAutoPatchingPtrInput
	// (Optional) An `keyVaultCredential` block as defined below.
	KeyVaultCredential VirtualMachineKeyVaultCredentialPtrInput
	// Should R Services be enabled?
	RServicesEnabled pulumi.BoolPtrInput
	// The SQL Server port. Defaults to `1433`.
	SqlConnectivityPort pulumi.IntPtrInput
	// The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
	SqlConnectivityType pulumi.StringPtrInput
	// The SQL Server sysadmin login password.
	SqlConnectivityUpdatePassword pulumi.StringPtrInput
	// The SQL Server sysadmin login to create.
	SqlConnectivityUpdateUsername pulumi.StringPtrInput
	// The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
	SqlLicenseType pulumi.StringInput
	// An `storageConfiguration` block as defined below.
	StorageConfiguration VirtualMachineStorageConfigurationPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}

type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput
}

func (*VirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil))
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

func (i *VirtualMachine) ToVirtualMachinePtrOutput() VirtualMachinePtrOutput {
	return i.ToVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachinePtrOutputWithContext(ctx context.Context) VirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePtrOutput)
}

type VirtualMachinePtrInput interface {
	pulumi.Input

	ToVirtualMachinePtrOutput() VirtualMachinePtrOutput
	ToVirtualMachinePtrOutputWithContext(ctx context.Context) VirtualMachinePtrOutput
}

type virtualMachinePtrType VirtualMachineArgs

func (*virtualMachinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil))
}

func (i *virtualMachinePtrType) ToVirtualMachinePtrOutput() VirtualMachinePtrOutput {
	return i.ToVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *virtualMachinePtrType) ToVirtualMachinePtrOutputWithContext(ctx context.Context) VirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePtrOutput)
}

// VirtualMachineArrayInput is an input type that accepts VirtualMachineArray and VirtualMachineArrayOutput values.
// You can construct a concrete instance of `VirtualMachineArrayInput` via:
//
//          VirtualMachineArray{ VirtualMachineArgs{...} }
type VirtualMachineArrayInput interface {
	pulumi.Input

	ToVirtualMachineArrayOutput() VirtualMachineArrayOutput
	ToVirtualMachineArrayOutputWithContext(context.Context) VirtualMachineArrayOutput
}

type VirtualMachineArray []VirtualMachineInput

func (VirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VirtualMachine)(nil))
}

func (i VirtualMachineArray) ToVirtualMachineArrayOutput() VirtualMachineArrayOutput {
	return i.ToVirtualMachineArrayOutputWithContext(context.Background())
}

func (i VirtualMachineArray) ToVirtualMachineArrayOutputWithContext(ctx context.Context) VirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineArrayOutput)
}

// VirtualMachineMapInput is an input type that accepts VirtualMachineMap and VirtualMachineMapOutput values.
// You can construct a concrete instance of `VirtualMachineMapInput` via:
//
//          VirtualMachineMap{ "key": VirtualMachineArgs{...} }
type VirtualMachineMapInput interface {
	pulumi.Input

	ToVirtualMachineMapOutput() VirtualMachineMapOutput
	ToVirtualMachineMapOutputWithContext(context.Context) VirtualMachineMapOutput
}

type VirtualMachineMap map[string]VirtualMachineInput

func (VirtualMachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VirtualMachine)(nil))
}

func (i VirtualMachineMap) ToVirtualMachineMapOutput() VirtualMachineMapOutput {
	return i.ToVirtualMachineMapOutputWithContext(context.Background())
}

func (i VirtualMachineMap) ToVirtualMachineMapOutputWithContext(ctx context.Context) VirtualMachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineMapOutput)
}

type VirtualMachineOutput struct {
	*pulumi.OutputState
}

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil))
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachinePtrOutput() VirtualMachinePtrOutput {
	return o.ToVirtualMachinePtrOutputWithContext(context.Background())
}

func (o VirtualMachineOutput) ToVirtualMachinePtrOutputWithContext(ctx context.Context) VirtualMachinePtrOutput {
	return o.ApplyT(func(v VirtualMachine) *VirtualMachine {
		return &v
	}).(VirtualMachinePtrOutput)
}

type VirtualMachinePtrOutput struct {
	*pulumi.OutputState
}

func (VirtualMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil))
}

func (o VirtualMachinePtrOutput) ToVirtualMachinePtrOutput() VirtualMachinePtrOutput {
	return o
}

func (o VirtualMachinePtrOutput) ToVirtualMachinePtrOutputWithContext(ctx context.Context) VirtualMachinePtrOutput {
	return o
}

type VirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachine)(nil))
}

func (o VirtualMachineArrayOutput) ToVirtualMachineArrayOutput() VirtualMachineArrayOutput {
	return o
}

func (o VirtualMachineArrayOutput) ToVirtualMachineArrayOutputWithContext(ctx context.Context) VirtualMachineArrayOutput {
	return o
}

func (o VirtualMachineArrayOutput) Index(i pulumi.IntInput) VirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachine {
		return vs[0].([]VirtualMachine)[vs[1].(int)]
	}).(VirtualMachineOutput)
}

type VirtualMachineMapOutput struct{ *pulumi.OutputState }

func (VirtualMachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualMachine)(nil))
}

func (o VirtualMachineMapOutput) ToVirtualMachineMapOutput() VirtualMachineMapOutput {
	return o
}

func (o VirtualMachineMapOutput) ToVirtualMachineMapOutputWithContext(ctx context.Context) VirtualMachineMapOutput {
	return o
}

func (o VirtualMachineMapOutput) MapIndex(k pulumi.StringInput) VirtualMachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualMachine {
		return vs[0].(map[string]VirtualMachine)[vs[1].(string)]
	}).(VirtualMachineOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
	pulumi.RegisterOutputType(VirtualMachinePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineMapOutput{})
}
