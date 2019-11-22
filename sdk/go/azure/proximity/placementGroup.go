// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proximity

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a proximity placement group for virtual machines, virtual machine scale sets and availability sets.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/proximity_placement_group.html.markdown.
type PlacementGroup struct {
	s *pulumi.ResourceState
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOpt) (*PlacementGroup, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
	}
	s, err := ctx.RegisterResource("azure:proximity/placementGroup:PlacementGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PlacementGroup{s: s}, nil
}

// GetPlacementGroup gets an existing PlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PlacementGroupState, opts ...pulumi.ResourceOpt) (*PlacementGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:proximity/placementGroup:PlacementGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PlacementGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PlacementGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PlacementGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *PlacementGroup) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the availability set. Changing this forces a new resource to be created.
func (r *PlacementGroup) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
func (r *PlacementGroup) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *PlacementGroup) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering PlacementGroup resources.
type PlacementGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
