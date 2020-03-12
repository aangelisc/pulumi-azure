// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Extension for a Virtual Machine Scale Set.
//
// > **NOTE:** This resource is not intended to be used with the `compute.ScaleSet` resource - instead it's intended for this to be used with the `compute.LinuxVirtualMachineScaleSet` and `compute.WindowsVirtualMachineScaleSet` resources.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine_scale_set_extension.html.markdown.
type VirtualMachineScaleSetExtension struct {
	pulumi.CustomResourceState

	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings pulumi.StringPtrOutput `pulumi:"protectedSettings"`
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions pulumi.StringArrayOutput `pulumi:"provisionAfterExtensions"`
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher pulumi.StringOutput `pulumi:"publisher"`
	// A JSON String which specifies Settings for the Extension.
	Settings pulumi.StringPtrOutput `pulumi:"settings"`
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the Script Handler which should be used.
	TypeHandlerVersion pulumi.StringOutput `pulumi:"typeHandlerVersion"`
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringOutput `pulumi:"virtualMachineScaleSetId"`
}

// NewVirtualMachineScaleSetExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	if args == nil || args.Publisher == nil {
		return nil, errors.New("missing required argument 'Publisher'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.TypeHandlerVersion == nil {
		return nil, errors.New("missing required argument 'TypeHandlerVersion'")
	}
	if args == nil || args.VirtualMachineScaleSetId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineScaleSetId'")
	}
	if args == nil {
		args = &VirtualMachineScaleSetExtensionArgs{}
	}
	var resource VirtualMachineScaleSetExtension
	err := ctx.RegisterResource("azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetExtension gets an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	var resource VirtualMachineScaleSetExtension
	err := ctx.ReadResource("azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetExtension resources.
type virtualMachineScaleSetExtensionState struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings *string `pulumi:"protectedSettings"`
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions []string `pulumi:"provisionAfterExtensions"`
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher *string `pulumi:"publisher"`
	// A JSON String which specifies Settings for the Extension.
	Settings *string `pulumi:"settings"`
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
	// Specifies the version of the Script Handler which should be used.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId *string `pulumi:"virtualMachineScaleSetId"`
}

type VirtualMachineScaleSetExtensionState struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings pulumi.StringPtrInput
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions pulumi.StringArrayInput
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher pulumi.StringPtrInput
	// A JSON String which specifies Settings for the Extension.
	Settings pulumi.StringPtrInput
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
	// Specifies the version of the Script Handler which should be used.
	TypeHandlerVersion pulumi.StringPtrInput
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringPtrInput
}

func (VirtualMachineScaleSetExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetExtensionArgs struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings *string `pulumi:"protectedSettings"`
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions []string `pulumi:"provisionAfterExtensions"`
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher string `pulumi:"publisher"`
	// A JSON String which specifies Settings for the Extension.
	Settings *string `pulumi:"settings"`
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
	// Specifies the version of the Script Handler which should be used.
	TypeHandlerVersion string `pulumi:"typeHandlerVersion"`
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId string `pulumi:"virtualMachineScaleSetId"`
}

// The set of arguments for constructing a VirtualMachineScaleSetExtension resource.
type VirtualMachineScaleSetExtensionArgs struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings pulumi.StringPtrInput
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions pulumi.StringArrayInput
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher pulumi.StringInput
	// A JSON String which specifies Settings for the Extension.
	Settings pulumi.StringPtrInput
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type pulumi.StringInput
	// Specifies the version of the Script Handler which should be used.
	TypeHandlerVersion pulumi.StringInput
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringInput
}

func (VirtualMachineScaleSetExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionArgs)(nil)).Elem()
}
