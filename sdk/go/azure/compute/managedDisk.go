// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage a managed disk.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/managed_disk.html.markdown.
type ManagedDisk struct {
	s *pulumi.ResourceState
}

// NewManagedDisk registers a new resource with the given unique name, arguments, and options.
func NewManagedDisk(ctx *pulumi.Context,
	name string, args *ManagedDiskArgs, opts ...pulumi.ResourceOpt) (*ManagedDisk, error) {
	if args == nil || args.CreateOption == nil {
		return nil, errors.New("missing required argument 'CreateOption'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageAccountType == nil {
		return nil, errors.New("missing required argument 'StorageAccountType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["createOption"] = nil
		inputs["diskSizeGb"] = nil
		inputs["encryptionSettings"] = nil
		inputs["imageReferenceId"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["osType"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sourceResourceId"] = nil
		inputs["sourceUri"] = nil
		inputs["storageAccountType"] = nil
		inputs["tags"] = nil
		inputs["zones"] = nil
	} else {
		inputs["createOption"] = args.CreateOption
		inputs["diskSizeGb"] = args.DiskSizeGb
		inputs["encryptionSettings"] = args.EncryptionSettings
		inputs["imageReferenceId"] = args.ImageReferenceId
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["osType"] = args.OsType
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sourceResourceId"] = args.SourceResourceId
		inputs["sourceUri"] = args.SourceUri
		inputs["storageAccountType"] = args.StorageAccountType
		inputs["tags"] = args.Tags
		inputs["zones"] = args.Zones
	}
	s, err := ctx.RegisterResource("azure:compute/managedDisk:ManagedDisk", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ManagedDisk{s: s}, nil
}

// GetManagedDisk gets an existing ManagedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedDisk(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ManagedDiskState, opts ...pulumi.ResourceOpt) (*ManagedDisk, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["createOption"] = state.CreateOption
		inputs["diskSizeGb"] = state.DiskSizeGb
		inputs["encryptionSettings"] = state.EncryptionSettings
		inputs["imageReferenceId"] = state.ImageReferenceId
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["osType"] = state.OsType
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sourceResourceId"] = state.SourceResourceId
		inputs["sourceUri"] = state.SourceUri
		inputs["storageAccountType"] = state.StorageAccountType
		inputs["tags"] = state.Tags
		inputs["zones"] = state.Zones
	}
	s, err := ctx.ReadResource("azure:compute/managedDisk:ManagedDisk", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ManagedDisk{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ManagedDisk) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ManagedDisk) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The method to use when creating the managed disk. Possible values include:
func (r *ManagedDisk) CreateOption() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createOption"])
}

// Specifies the size of the managed disk to create in gigabytes.
// If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
func (r *ManagedDisk) DiskSizeGb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["diskSizeGb"])
}

// an `encryption_settings` block as defined below.
func (r *ManagedDisk) EncryptionSettings() *pulumi.Output {
	return r.s.State["encryptionSettings"]
}

// ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
func (r *ManagedDisk) ImageReferenceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["imageReferenceId"])
}

// Specified the supported Azure location where the resource exists.
// Changing this forces a new resource to be created.
func (r *ManagedDisk) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the managed disk. Changing this forces a
// new resource to be created.
func (r *ManagedDisk) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specify a value when the source of an `Import` or `Copy`
// operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`
func (r *ManagedDisk) OsType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["osType"])
}

// The name of the resource group in which to create
// the managed disk.
func (r *ManagedDisk) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// ID of an existing managed disk to copy `create_option` is `Copy`
// or the recovery point to restore when `create_option` is `Restore`
func (r *ManagedDisk) SourceResourceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceResourceId"])
}

// URI to a valid VHD file to be used when `create_option` is `Import`.
func (r *ManagedDisk) SourceUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceUri"])
}

// The type of storage to use for the managed disk.
// Allowable values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
func (r *ManagedDisk) StorageAccountType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageAccountType"])
}

// A mapping of tags to assign to the resource.
func (r *ManagedDisk) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// A collection containing the availability zone to allocate the Managed Disk in.
func (r *ManagedDisk) Zones() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zones"])
}

// Input properties used for looking up and filtering ManagedDisk resources.
type ManagedDiskState struct {
	// The method to use when creating the managed disk. Possible values include:
	CreateOption interface{}
	// Specifies the size of the managed disk to create in gigabytes.
	// If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb interface{}
	// an `encryption_settings` block as defined below.
	EncryptionSettings interface{}
	// ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
	ImageReferenceId interface{}
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the managed disk. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Specify a value when the source of an `Import` or `Copy`
	// operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`
	OsType interface{}
	// The name of the resource group in which to create
	// the managed disk.
	ResourceGroupName interface{}
	// ID of an existing managed disk to copy `create_option` is `Copy`
	// or the recovery point to restore when `create_option` is `Restore`
	SourceResourceId interface{}
	// URI to a valid VHD file to be used when `create_option` is `Import`.
	SourceUri interface{}
	// The type of storage to use for the managed disk.
	// Allowable values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones interface{}
}

// The set of arguments for constructing a ManagedDisk resource.
type ManagedDiskArgs struct {
	// The method to use when creating the managed disk. Possible values include:
	CreateOption interface{}
	// Specifies the size of the managed disk to create in gigabytes.
	// If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb interface{}
	// an `encryption_settings` block as defined below.
	EncryptionSettings interface{}
	// ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
	ImageReferenceId interface{}
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the managed disk. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Specify a value when the source of an `Import` or `Copy`
	// operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`
	OsType interface{}
	// The name of the resource group in which to create
	// the managed disk.
	ResourceGroupName interface{}
	// ID of an existing managed disk to copy `create_option` is `Copy`
	// or the recovery point to restore when `create_option` is `Restore`
	SourceResourceId interface{}
	// URI to a valid VHD file to be used when `create_option` is `Import`.
	SourceUri interface{}
	// The type of storage to use for the managed disk.
	// Allowable values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones interface{}
}
