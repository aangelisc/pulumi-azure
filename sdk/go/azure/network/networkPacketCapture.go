// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures Network Packet Capturing against a Virtual Machine using a Network Watcher.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_packet_capture.html.markdown.
type NetworkPacketCapture struct {
	s *pulumi.ResourceState
}

// NewNetworkPacketCapture registers a new resource with the given unique name, arguments, and options.
func NewNetworkPacketCapture(ctx *pulumi.Context,
	name string, args *NetworkPacketCaptureArgs, opts ...pulumi.ResourceOpt) (*NetworkPacketCapture, error) {
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageLocation == nil {
		return nil, errors.New("missing required argument 'StorageLocation'")
	}
	if args == nil || args.TargetResourceId == nil {
		return nil, errors.New("missing required argument 'TargetResourceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["filters"] = nil
		inputs["maximumBytesPerPacket"] = nil
		inputs["maximumBytesPerSession"] = nil
		inputs["maximumCaptureDuration"] = nil
		inputs["name"] = nil
		inputs["networkWatcherName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["storageLocation"] = nil
		inputs["targetResourceId"] = nil
	} else {
		inputs["filters"] = args.Filters
		inputs["maximumBytesPerPacket"] = args.MaximumBytesPerPacket
		inputs["maximumBytesPerSession"] = args.MaximumBytesPerSession
		inputs["maximumCaptureDuration"] = args.MaximumCaptureDuration
		inputs["name"] = args.Name
		inputs["networkWatcherName"] = args.NetworkWatcherName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["storageLocation"] = args.StorageLocation
		inputs["targetResourceId"] = args.TargetResourceId
	}
	s, err := ctx.RegisterResource("azure:network/networkPacketCapture:NetworkPacketCapture", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkPacketCapture{s: s}, nil
}

// GetNetworkPacketCapture gets an existing NetworkPacketCapture resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPacketCapture(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkPacketCaptureState, opts ...pulumi.ResourceOpt) (*NetworkPacketCapture, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["filters"] = state.Filters
		inputs["maximumBytesPerPacket"] = state.MaximumBytesPerPacket
		inputs["maximumBytesPerSession"] = state.MaximumBytesPerSession
		inputs["maximumCaptureDuration"] = state.MaximumCaptureDuration
		inputs["name"] = state.Name
		inputs["networkWatcherName"] = state.NetworkWatcherName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["storageLocation"] = state.StorageLocation
		inputs["targetResourceId"] = state.TargetResourceId
	}
	s, err := ctx.ReadResource("azure:network/networkPacketCapture:NetworkPacketCapture", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkPacketCapture{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NetworkPacketCapture) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NetworkPacketCapture) ID() pulumi.IDOutput {
	return r.s.ID()
}

// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) Filters() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["filters"])
}

// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) MaximumBytesPerPacket() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maximumBytesPerPacket"])
}

// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) MaximumBytesPerSession() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maximumBytesPerSession"])
}

// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) MaximumCaptureDuration() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maximumCaptureDuration"])
}

// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Network Watcher. Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) NetworkWatcherName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["networkWatcherName"])
}

// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) StorageLocation() pulumi.Output {
	return r.s.State["storageLocation"]
}

// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
func (r *NetworkPacketCapture) TargetResourceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["targetResourceId"])
}

// Input properties used for looking up and filtering NetworkPacketCapture resources.
type NetworkPacketCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters interface{}
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket interface{}
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession interface{}
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration interface{}
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName interface{}
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation interface{}
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId interface{}
}

// The set of arguments for constructing a NetworkPacketCapture resource.
type NetworkPacketCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters interface{}
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket interface{}
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession interface{}
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration interface{}
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName interface{}
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation interface{}
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId interface{}
}
