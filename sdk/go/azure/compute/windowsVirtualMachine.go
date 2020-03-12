// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type WindowsVirtualMachine struct {
	pulumi.CustomResourceState

	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities WindowsVirtualMachineAdditionalCapabilitiesPtrOutput `pulumi:"additionalCapabilities"`
	// One or more `additionalUnattendContent` blocks as defined below. Changing this forces a new resource to be created.
	AdditionalUnattendContents WindowsVirtualMachineAdditionalUnattendContentArrayOutput `pulumi:"additionalUnattendContents"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringOutput `pulumi:"adminPassword"`
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// Should Extension Operations be allowed on this Virtual Machine? Changing this forces a new resource to be created.
	AllowExtensionOperations pulumi.BoolPtrOutput `pulumi:"allowExtensionOperations"`
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrOutput `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics WindowsVirtualMachineBootDiagnosticsPtrOutput `pulumi:"bootDiagnostics"`
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. Changing this forces a new resource to be created.
	ComputerName pulumi.StringOutput `pulumi:"computerName"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData pulumi.StringPtrOutput `pulumi:"customData"`
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId pulumi.StringPtrOutput `pulumi:"dedicatedHostId"`
	// Specifies if Automatic Updates are Enabled for the Windows Virtual Machine. Changing this forces a new resource to be created.
	EnableAutomaticUpdates pulumi.BoolPtrOutput `pulumi:"enableAutomaticUpdates"`
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// An `identity` block as defined below.
	Identity WindowsVirtualMachineIdentityPtrOutput `pulumi:"identity"`
	// Specifies the type of on-premise license (also known as [Azure Hybrid Use Benefit](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing)) which should be used for this Virtual Machine. Possible values are `None`, `Windows_Client` and `Windows_Server`. Changing this forces a new resource to be created.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// The Azure location where the Windows Virtual Machine should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrOutput `pulumi:"maxBidPrice"`
	// The name of the Windows Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// A `osDisk` block as defined below.
	OsDisk WindowsVirtualMachineOsDiskOutput `pulumi:"osDisk"`
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan WindowsVirtualMachinePlanPtrOutput `pulumi:"plan"`
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrOutput `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIpAddress pulumi.StringOutput `pulumi:"publicIpAddress"`
	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIpAddresses pulumi.StringArrayOutput `pulumi:"publicIpAddresses"`
	// The name of the Resource Group in which the Windows Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `secret` blocks as defined below.
	Secrets WindowsVirtualMachineSecretArrayOutput `pulumi:"secrets"`
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size pulumi.StringOutput `pulumi:"size"`
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId pulumi.StringPtrOutput `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference WindowsVirtualMachineSourceImageReferencePtrOutput `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Time Zone which should be used by the Virtual Machine, [the possible values are defined here](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
	// One or more `winrmListener` blocks as defined below.
	WinrmListeners WindowsVirtualMachineWinrmListenerArrayOutput `pulumi:"winrmListeners"`
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewWindowsVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewWindowsVirtualMachine(ctx *pulumi.Context,
	name string, args *WindowsVirtualMachineArgs, opts ...pulumi.ResourceOption) (*WindowsVirtualMachine, error) {
	if args == nil || args.AdminPassword == nil {
		return nil, errors.New("missing required argument 'AdminPassword'")
	}
	if args == nil || args.AdminUsername == nil {
		return nil, errors.New("missing required argument 'AdminUsername'")
	}
	if args == nil || args.NetworkInterfaceIds == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceIds'")
	}
	if args == nil || args.OsDisk == nil {
		return nil, errors.New("missing required argument 'OsDisk'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &WindowsVirtualMachineArgs{}
	}
	var resource WindowsVirtualMachine
	err := ctx.RegisterResource("azure:compute/windowsVirtualMachine:WindowsVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWindowsVirtualMachine gets an existing WindowsVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWindowsVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WindowsVirtualMachineState, opts ...pulumi.ResourceOption) (*WindowsVirtualMachine, error) {
	var resource WindowsVirtualMachine
	err := ctx.ReadResource("azure:compute/windowsVirtualMachine:WindowsVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WindowsVirtualMachine resources.
type windowsVirtualMachineState struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities *WindowsVirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// One or more `additionalUnattendContent` blocks as defined below. Changing this forces a new resource to be created.
	AdditionalUnattendContents []WindowsVirtualMachineAdditionalUnattendContent `pulumi:"additionalUnattendContents"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword *string `pulumi:"adminPassword"`
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername *string `pulumi:"adminUsername"`
	// Should Extension Operations be allowed on this Virtual Machine? Changing this forces a new resource to be created.
	AllowExtensionOperations *bool `pulumi:"allowExtensionOperations"`
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics *WindowsVirtualMachineBootDiagnostics `pulumi:"bootDiagnostics"`
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. Changing this forces a new resource to be created.
	ComputerName *string `pulumi:"computerName"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData *string `pulumi:"customData"`
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId *string `pulumi:"dedicatedHostId"`
	// Specifies if Automatic Updates are Enabled for the Windows Virtual Machine. Changing this forces a new resource to be created.
	EnableAutomaticUpdates *bool `pulumi:"enableAutomaticUpdates"`
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// An `identity` block as defined below.
	Identity *WindowsVirtualMachineIdentity `pulumi:"identity"`
	// Specifies the type of on-premise license (also known as [Azure Hybrid Use Benefit](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing)) which should be used for this Virtual Machine. Possible values are `None`, `Windows_Client` and `Windows_Server`. Changing this forces a new resource to be created.
	LicenseType *string `pulumi:"licenseType"`
	// The Azure location where the Windows Virtual Machine should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Windows Virtual Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// A `osDisk` block as defined below.
	OsDisk *WindowsVirtualMachineOsDisk `pulumi:"osDisk"`
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan *WindowsVirtualMachinePlan `pulumi:"plan"`
	Priority *string `pulumi:"priority"`
	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent *bool `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIpAddress *string `pulumi:"publicIpAddress"`
	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIpAddresses []string `pulumi:"publicIpAddresses"`
	// The name of the Resource Group in which the Windows Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `secret` blocks as defined below.
	Secrets []WindowsVirtualMachineSecret `pulumi:"secrets"`
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size *string `pulumi:"size"`
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference *WindowsVirtualMachineSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Time Zone which should be used by the Virtual Machine, [the possible values are defined here](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone *string `pulumi:"timezone"`
	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
	// One or more `winrmListener` blocks as defined below.
	WinrmListeners []WindowsVirtualMachineWinrmListener `pulumi:"winrmListeners"`
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone *string `pulumi:"zone"`
}

type WindowsVirtualMachineState struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities WindowsVirtualMachineAdditionalCapabilitiesPtrInput
	// One or more `additionalUnattendContent` blocks as defined below. Changing this forces a new resource to be created.
	AdditionalUnattendContents WindowsVirtualMachineAdditionalUnattendContentArrayInput
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringPtrInput
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringPtrInput
	// Should Extension Operations be allowed on this Virtual Machine? Changing this forces a new resource to be created.
	AllowExtensionOperations pulumi.BoolPtrInput
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrInput
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics WindowsVirtualMachineBootDiagnosticsPtrInput
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. Changing this forces a new resource to be created.
	ComputerName pulumi.StringPtrInput
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData pulumi.StringPtrInput
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId pulumi.StringPtrInput
	// Specifies if Automatic Updates are Enabled for the Windows Virtual Machine. Changing this forces a new resource to be created.
	EnableAutomaticUpdates pulumi.BoolPtrInput
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity WindowsVirtualMachineIdentityPtrInput
	// Specifies the type of on-premise license (also known as [Azure Hybrid Use Benefit](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing)) which should be used for this Virtual Machine. Possible values are `None`, `Windows_Client` and `Windows_Server`. Changing this forces a new resource to be created.
	LicenseType pulumi.StringPtrInput
	// The Azure location where the Windows Virtual Machine should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Windows Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayInput
	// A `osDisk` block as defined below.
	OsDisk WindowsVirtualMachineOsDiskPtrInput
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan WindowsVirtualMachinePlanPtrInput
	Priority pulumi.StringPtrInput
	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIpAddress pulumi.StringPtrInput
	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIpAddresses pulumi.StringArrayInput
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrInput
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIpAddress pulumi.StringPtrInput
	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIpAddresses pulumi.StringArrayInput
	// The name of the Resource Group in which the Windows Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `secret` blocks as defined below.
	Secrets WindowsVirtualMachineSecretArrayInput
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size pulumi.StringPtrInput
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference WindowsVirtualMachineSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags pulumi.StringMapInput
	// Specifies the Time Zone which should be used by the Virtual Machine, [the possible values are defined here](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone pulumi.StringPtrInput
	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineId pulumi.StringPtrInput
	// One or more `winrmListener` blocks as defined below.
	WinrmListeners WindowsVirtualMachineWinrmListenerArrayInput
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone pulumi.StringPtrInput
}

func (WindowsVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsVirtualMachineState)(nil)).Elem()
}

type windowsVirtualMachineArgs struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities *WindowsVirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// One or more `additionalUnattendContent` blocks as defined below. Changing this forces a new resource to be created.
	AdditionalUnattendContents []WindowsVirtualMachineAdditionalUnattendContent `pulumi:"additionalUnattendContents"`
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword string `pulumi:"adminPassword"`
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername string `pulumi:"adminUsername"`
	// Should Extension Operations be allowed on this Virtual Machine? Changing this forces a new resource to be created.
	AllowExtensionOperations *bool `pulumi:"allowExtensionOperations"`
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics *WindowsVirtualMachineBootDiagnostics `pulumi:"bootDiagnostics"`
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. Changing this forces a new resource to be created.
	ComputerName *string `pulumi:"computerName"`
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData *string `pulumi:"customData"`
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId *string `pulumi:"dedicatedHostId"`
	// Specifies if Automatic Updates are Enabled for the Windows Virtual Machine. Changing this forces a new resource to be created.
	EnableAutomaticUpdates *bool `pulumi:"enableAutomaticUpdates"`
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// An `identity` block as defined below.
	Identity *WindowsVirtualMachineIdentity `pulumi:"identity"`
	// Specifies the type of on-premise license (also known as [Azure Hybrid Use Benefit](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing)) which should be used for this Virtual Machine. Possible values are `None`, `Windows_Client` and `Windows_Server`. Changing this forces a new resource to be created.
	LicenseType *string `pulumi:"licenseType"`
	// The Azure location where the Windows Virtual Machine should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// The name of the Windows Virtual Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// A `osDisk` block as defined below.
	OsDisk WindowsVirtualMachineOsDisk `pulumi:"osDisk"`
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan *WindowsVirtualMachinePlan `pulumi:"plan"`
	Priority *string `pulumi:"priority"`
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent *bool `pulumi:"provisionVmAgent"`
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the Resource Group in which the Windows Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `secret` blocks as defined below.
	Secrets []WindowsVirtualMachineSecret `pulumi:"secrets"`
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size string `pulumi:"size"`
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId *string `pulumi:"sourceImageId"`
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference *WindowsVirtualMachineSourceImageReference `pulumi:"sourceImageReference"`
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Time Zone which should be used by the Virtual Machine, [the possible values are defined here](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone *string `pulumi:"timezone"`
	// One or more `winrmListener` blocks as defined below.
	WinrmListeners []WindowsVirtualMachineWinrmListener `pulumi:"winrmListeners"`
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a WindowsVirtualMachine resource.
type WindowsVirtualMachineArgs struct {
	// A `additionalCapabilities` block as defined below.
	AdditionalCapabilities WindowsVirtualMachineAdditionalCapabilitiesPtrInput
	// One or more `additionalUnattendContent` blocks as defined below. Changing this forces a new resource to be created.
	AdditionalUnattendContents WindowsVirtualMachineAdditionalUnattendContentArrayInput
	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	AdminPassword pulumi.StringInput
	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	AdminUsername pulumi.StringInput
	// Should Extension Operations be allowed on this Virtual Machine? Changing this forces a new resource to be created.
	AllowExtensionOperations pulumi.BoolPtrInput
	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrInput
	// A `bootDiagnostics` block as defined below.
	BootDiagnostics WindowsVirtualMachineBootDiagnosticsPtrInput
	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. Changing this forces a new resource to be created.
	ComputerName pulumi.StringPtrInput
	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	CustomData pulumi.StringPtrInput
	// The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
	DedicatedHostId pulumi.StringPtrInput
	// Specifies if Automatic Updates are Enabled for the Windows Virtual Machine. Changing this forces a new resource to be created.
	EnableAutomaticUpdates pulumi.BoolPtrInput
	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity WindowsVirtualMachineIdentityPtrInput
	// Specifies the type of on-premise license (also known as [Azure Hybrid Use Benefit](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing)) which should be used for this Virtual Machine. Possible values are `None`, `Windows_Client` and `Windows_Server`. Changing this forces a new resource to be created.
	LicenseType pulumi.StringPtrInput
	// The Azure location where the Windows Virtual Machine should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
	MaxBidPrice pulumi.Float64PtrInput
	// The name of the Windows Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayInput
	// A `osDisk` block as defined below.
	OsDisk WindowsVirtualMachineOsDiskInput
	// A `plan` block as defined below. Changing this forces a new resource to be created.
	Plan WindowsVirtualMachinePlanPtrInput
	Priority pulumi.StringPtrInput
	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
	ProvisionVmAgent pulumi.BoolPtrInput
	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the Resource Group in which the Windows Virtual Machine should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `secret` blocks as defined below.
	Secrets WindowsVirtualMachineSecretArrayInput
	// The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
	Size pulumi.StringInput
	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
	SourceImageId pulumi.StringPtrInput
	// A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
	SourceImageReference WindowsVirtualMachineSourceImageReferencePtrInput
	// A mapping of tags which should be assigned to this Virtual Machine.
	Tags pulumi.StringMapInput
	// Specifies the Time Zone which should be used by the Virtual Machine, [the possible values are defined here](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone pulumi.StringPtrInput
	// One or more `winrmListener` blocks as defined below.
	WinrmListeners WindowsVirtualMachineWinrmListenerArrayInput
	// The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
	Zone pulumi.StringPtrInput
}

func (WindowsVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsVirtualMachineArgs)(nil)).Elem()
}

