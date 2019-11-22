// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Container within an Azure Storage Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_container.html.markdown.
type Container struct {
	s *pulumi.ResourceState
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOpt) (*Container, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["containerAccessType"] = nil
		inputs["metadata"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["storageAccountName"] = nil
	} else {
		inputs["containerAccessType"] = args.ContainerAccessType
		inputs["metadata"] = args.Metadata
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["storageAccountName"] = args.StorageAccountName
	}
	inputs["hasImmutabilityPolicy"] = nil
	inputs["hasLegalHold"] = nil
	inputs["properties"] = nil
	s, err := ctx.RegisterResource("azure:storage/container:Container", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Container{s: s}, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ContainerState, opts ...pulumi.ResourceOpt) (*Container, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["containerAccessType"] = state.ContainerAccessType
		inputs["hasImmutabilityPolicy"] = state.HasImmutabilityPolicy
		inputs["hasLegalHold"] = state.HasLegalHold
		inputs["metadata"] = state.Metadata
		inputs["name"] = state.Name
		inputs["properties"] = state.Properties
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["storageAccountName"] = state.StorageAccountName
	}
	s, err := ctx.ReadResource("azure:storage/container:Container", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Container{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Container) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Container) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
func (r *Container) ContainerAccessType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["containerAccessType"])
}

// Is there an Immutability Policy configured on this Storage Container?
func (r *Container) HasImmutabilityPolicy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["hasImmutabilityPolicy"])
}

// Is there a Legal Hold configured on this Storage Container?
func (r *Container) HasLegalHold() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["hasLegalHold"])
}

// A mapping of MetaData for this Container.
func (r *Container) Metadata() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["metadata"])
}

// The name of the Container which should be created within the Storage Account.
func (r *Container) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// (**Deprecated**) Key-value definition of additional properties associated to the Storage Container
func (r *Container) Properties() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["properties"])
}

// The name of the resource group in which to create the storage container. This field is no longer used and will be removed in 2.0. 
func (r *Container) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The name of the Storage Account where the Container should be created.
func (r *Container) StorageAccountName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageAccountName"])
}

// Input properties used for looking up and filtering Container resources.
type ContainerState struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType interface{}
	// Is there an Immutability Policy configured on this Storage Container?
	HasImmutabilityPolicy interface{}
	// Is there a Legal Hold configured on this Storage Container?
	HasLegalHold interface{}
	// A mapping of MetaData for this Container.
	Metadata interface{}
	// The name of the Container which should be created within the Storage Account.
	Name interface{}
	// (**Deprecated**) Key-value definition of additional properties associated to the Storage Container
	Properties interface{}
	// The name of the resource group in which to create the storage container. This field is no longer used and will be removed in 2.0. 
	ResourceGroupName interface{}
	// The name of the Storage Account where the Container should be created.
	StorageAccountName interface{}
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
	ContainerAccessType interface{}
	// A mapping of MetaData for this Container.
	Metadata interface{}
	// The name of the Container which should be created within the Storage Account.
	Name interface{}
	// The name of the resource group in which to create the storage container. This field is no longer used and will be removed in 2.0. 
	ResourceGroupName interface{}
	// The name of the Storage Account where the Container should be created.
	StorageAccountName interface{}
}
