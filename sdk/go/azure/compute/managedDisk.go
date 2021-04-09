// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a managed disk.
//
// ## Example Usage
// ### With Create Empty
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		_, err = compute.NewManagedDisk(ctx, "exampleManagedDisk", &compute.ManagedDiskArgs{
// 			Location:           pulumi.String("West US 2"),
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			StorageAccountType: pulumi.String("Standard_LRS"),
// 			CreateOption:       pulumi.String("Empty"),
// 			DiskSizeGb:         pulumi.Int(1),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Create Copy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		source, err := compute.NewManagedDisk(ctx, "source", &compute.ManagedDiskArgs{
// 			Location:           pulumi.String("West US 2"),
// 			ResourceGroupName:  example.Name,
// 			StorageAccountType: pulumi.String("Standard_LRS"),
// 			CreateOption:       pulumi.String("Empty"),
// 			DiskSizeGb:         pulumi.Int(1),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewManagedDisk(ctx, "copy", &compute.ManagedDiskArgs{
// 			Location:           pulumi.String("West US 2"),
// 			ResourceGroupName:  example.Name,
// 			StorageAccountType: pulumi.String("Standard_LRS"),
// 			CreateOption:       pulumi.String("Copy"),
// 			SourceResourceId:   source.ID(),
// 			DiskSizeGb:         pulumi.Int(1),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
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
// Managed Disks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/managedDisk:ManagedDisk example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1
// ```
type ManagedDisk struct {
	pulumi.CustomResourceState

	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
	CreateOption pulumi.StringOutput `pulumi:"createOption"`
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
	DiskEncryptionSetId pulumi.StringPtrOutput `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite pulumi.IntOutput `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite pulumi.IntOutput `pulumi:"diskMbpsReadWrite"`
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
	DiskSizeGb pulumi.IntOutput `pulumi:"diskSizeGb"`
	// A `encryptionSettings` block as defined below.
	EncryptionSettings ManagedDiskEncryptionSettingsPtrOutput `pulumi:"encryptionSettings"`
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId pulumi.StringPtrOutput `pulumi:"imageReferenceId"`
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId pulumi.StringPtrOutput `pulumi:"sourceResourceId"`
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri pulumi.StringOutput `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType pulumi.StringOutput `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewManagedDisk registers a new resource with the given unique name, arguments, and options.
func NewManagedDisk(ctx *pulumi.Context,
	name string, args *ManagedDiskArgs, opts ...pulumi.ResourceOption) (*ManagedDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreateOption == nil {
		return nil, errors.New("invalid value for required argument 'CreateOption'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountType == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountType'")
	}
	var resource ManagedDisk
	err := ctx.RegisterResource("azure:compute/managedDisk:ManagedDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedDisk gets an existing ManagedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedDiskState, opts ...pulumi.ResourceOption) (*ManagedDisk, error) {
	var resource ManagedDisk
	err := ctx.ReadResource("azure:compute/managedDisk:ManagedDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedDisk resources.
type managedDiskState struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
	CreateOption *string `pulumi:"createOption"`
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite *int `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite *int `pulumi:"diskMbpsReadWrite"`
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// A `encryptionSettings` block as defined below.
	EncryptionSettings *ManagedDiskEncryptionSettings `pulumi:"encryptionSettings"`
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId *string `pulumi:"imageReferenceId"`
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType *string `pulumi:"osType"`
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri *string `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones *string `pulumi:"zones"`
}

type ManagedDiskState struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
	CreateOption pulumi.StringPtrInput
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
	DiskEncryptionSetId pulumi.StringPtrInput
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite pulumi.IntPtrInput
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite pulumi.IntPtrInput
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
	DiskSizeGb pulumi.IntPtrInput
	// A `encryptionSettings` block as defined below.
	EncryptionSettings ManagedDiskEncryptionSettingsPtrInput
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId pulumi.StringPtrInput
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri pulumi.StringPtrInput
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones pulumi.StringPtrInput
}

func (ManagedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDiskState)(nil)).Elem()
}

type managedDiskArgs struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
	CreateOption string `pulumi:"createOption"`
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite *int `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite *int `pulumi:"diskMbpsReadWrite"`
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// A `encryptionSettings` block as defined below.
	EncryptionSettings *ManagedDiskEncryptionSettings `pulumi:"encryptionSettings"`
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId *string `pulumi:"imageReferenceId"`
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType *string `pulumi:"osType"`
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri *string `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType string `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a ManagedDisk resource.
type ManagedDiskArgs struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
	CreateOption pulumi.StringInput
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
	DiskEncryptionSetId pulumi.StringPtrInput
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite pulumi.IntPtrInput
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite pulumi.IntPtrInput
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
	DiskSizeGb pulumi.IntPtrInput
	// A `encryptionSettings` block as defined below.
	EncryptionSettings ManagedDiskEncryptionSettingsPtrInput
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName pulumi.StringInput
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId pulumi.StringPtrInput
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri pulumi.StringPtrInput
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones pulumi.StringPtrInput
}

func (ManagedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDiskArgs)(nil)).Elem()
}

type ManagedDiskInput interface {
	pulumi.Input

	ToManagedDiskOutput() ManagedDiskOutput
	ToManagedDiskOutputWithContext(ctx context.Context) ManagedDiskOutput
}

func (*ManagedDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDisk)(nil))
}

func (i *ManagedDisk) ToManagedDiskOutput() ManagedDiskOutput {
	return i.ToManagedDiskOutputWithContext(context.Background())
}

func (i *ManagedDisk) ToManagedDiskOutputWithContext(ctx context.Context) ManagedDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskOutput)
}

func (i *ManagedDisk) ToManagedDiskPtrOutput() ManagedDiskPtrOutput {
	return i.ToManagedDiskPtrOutputWithContext(context.Background())
}

func (i *ManagedDisk) ToManagedDiskPtrOutputWithContext(ctx context.Context) ManagedDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskPtrOutput)
}

type ManagedDiskPtrInput interface {
	pulumi.Input

	ToManagedDiskPtrOutput() ManagedDiskPtrOutput
	ToManagedDiskPtrOutputWithContext(ctx context.Context) ManagedDiskPtrOutput
}

type managedDiskPtrType ManagedDiskArgs

func (*managedDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDisk)(nil))
}

func (i *managedDiskPtrType) ToManagedDiskPtrOutput() ManagedDiskPtrOutput {
	return i.ToManagedDiskPtrOutputWithContext(context.Background())
}

func (i *managedDiskPtrType) ToManagedDiskPtrOutputWithContext(ctx context.Context) ManagedDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskPtrOutput)
}

// ManagedDiskArrayInput is an input type that accepts ManagedDiskArray and ManagedDiskArrayOutput values.
// You can construct a concrete instance of `ManagedDiskArrayInput` via:
//
//          ManagedDiskArray{ ManagedDiskArgs{...} }
type ManagedDiskArrayInput interface {
	pulumi.Input

	ToManagedDiskArrayOutput() ManagedDiskArrayOutput
	ToManagedDiskArrayOutputWithContext(context.Context) ManagedDiskArrayOutput
}

type ManagedDiskArray []ManagedDiskInput

func (ManagedDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ManagedDisk)(nil))
}

func (i ManagedDiskArray) ToManagedDiskArrayOutput() ManagedDiskArrayOutput {
	return i.ToManagedDiskArrayOutputWithContext(context.Background())
}

func (i ManagedDiskArray) ToManagedDiskArrayOutputWithContext(ctx context.Context) ManagedDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskArrayOutput)
}

// ManagedDiskMapInput is an input type that accepts ManagedDiskMap and ManagedDiskMapOutput values.
// You can construct a concrete instance of `ManagedDiskMapInput` via:
//
//          ManagedDiskMap{ "key": ManagedDiskArgs{...} }
type ManagedDiskMapInput interface {
	pulumi.Input

	ToManagedDiskMapOutput() ManagedDiskMapOutput
	ToManagedDiskMapOutputWithContext(context.Context) ManagedDiskMapOutput
}

type ManagedDiskMap map[string]ManagedDiskInput

func (ManagedDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ManagedDisk)(nil))
}

func (i ManagedDiskMap) ToManagedDiskMapOutput() ManagedDiskMapOutput {
	return i.ToManagedDiskMapOutputWithContext(context.Background())
}

func (i ManagedDiskMap) ToManagedDiskMapOutputWithContext(ctx context.Context) ManagedDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskMapOutput)
}

type ManagedDiskOutput struct {
	*pulumi.OutputState
}

func (ManagedDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDisk)(nil))
}

func (o ManagedDiskOutput) ToManagedDiskOutput() ManagedDiskOutput {
	return o
}

func (o ManagedDiskOutput) ToManagedDiskOutputWithContext(ctx context.Context) ManagedDiskOutput {
	return o
}

func (o ManagedDiskOutput) ToManagedDiskPtrOutput() ManagedDiskPtrOutput {
	return o.ToManagedDiskPtrOutputWithContext(context.Background())
}

func (o ManagedDiskOutput) ToManagedDiskPtrOutputWithContext(ctx context.Context) ManagedDiskPtrOutput {
	return o.ApplyT(func(v ManagedDisk) *ManagedDisk {
		return &v
	}).(ManagedDiskPtrOutput)
}

type ManagedDiskPtrOutput struct {
	*pulumi.OutputState
}

func (ManagedDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDisk)(nil))
}

func (o ManagedDiskPtrOutput) ToManagedDiskPtrOutput() ManagedDiskPtrOutput {
	return o
}

func (o ManagedDiskPtrOutput) ToManagedDiskPtrOutputWithContext(ctx context.Context) ManagedDiskPtrOutput {
	return o
}

type ManagedDiskArrayOutput struct{ *pulumi.OutputState }

func (ManagedDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedDisk)(nil))
}

func (o ManagedDiskArrayOutput) ToManagedDiskArrayOutput() ManagedDiskArrayOutput {
	return o
}

func (o ManagedDiskArrayOutput) ToManagedDiskArrayOutputWithContext(ctx context.Context) ManagedDiskArrayOutput {
	return o
}

func (o ManagedDiskArrayOutput) Index(i pulumi.IntInput) ManagedDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedDisk {
		return vs[0].([]ManagedDisk)[vs[1].(int)]
	}).(ManagedDiskOutput)
}

type ManagedDiskMapOutput struct{ *pulumi.OutputState }

func (ManagedDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedDisk)(nil))
}

func (o ManagedDiskMapOutput) ToManagedDiskMapOutput() ManagedDiskMapOutput {
	return o
}

func (o ManagedDiskMapOutput) ToManagedDiskMapOutputWithContext(ctx context.Context) ManagedDiskMapOutput {
	return o
}

func (o ManagedDiskMapOutput) MapIndex(k pulumi.StringInput) ManagedDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedDisk {
		return vs[0].(map[string]ManagedDisk)[vs[1].(string)]
	}).(ManagedDiskOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedDiskOutput{})
	pulumi.RegisterOutputType(ManagedDiskPtrOutput{})
	pulumi.RegisterOutputType(ManagedDiskArrayOutput{})
	pulumi.RegisterOutputType(ManagedDiskMapOutput{})
}
