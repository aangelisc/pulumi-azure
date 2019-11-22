// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notificationhub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Notification Hub Namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/notification_hub_namespace.html.markdown.
type Namespace struct {
	s *pulumi.ResourceState
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOpt) (*Namespace, error) {
	if args == nil || args.NamespaceType == nil {
		return nil, errors.New("missing required argument 'NamespaceType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["enabled"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["namespaceType"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["skuName"] = nil
	} else {
		inputs["enabled"] = args.Enabled
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["namespaceType"] = args.NamespaceType
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["skuName"] = args.SkuName
	}
	inputs["servicebusEndpoint"] = nil
	s, err := ctx.RegisterResource("azure:notificationhub/namespace:Namespace", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Namespace{s: s}, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NamespaceState, opts ...pulumi.ResourceOpt) (*Namespace, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["enabled"] = state.Enabled
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["namespaceType"] = state.NamespaceType
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["servicebusEndpoint"] = state.ServicebusEndpoint
		inputs["sku"] = state.Sku
		inputs["skuName"] = state.SkuName
	}
	s, err := ctx.ReadResource("azure:notificationhub/namespace:Namespace", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Namespace{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Namespace) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Namespace) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Is this Notification Hub Namespace enabled? Defaults to `true`.
func (r *Namespace) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// The Azure Region in which this Notification Hub Namespace should be created.
func (r *Namespace) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
func (r *Namespace) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
func (r *Namespace) NamespaceType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namespaceType"])
}

// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
func (r *Namespace) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The ServiceBus Endpoint for this Notification Hub Namespace.
func (r *Namespace) ServicebusEndpoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["servicebusEndpoint"])
}

// ) A `sku` block as described below.
func (r *Namespace) Sku() pulumi.Output {
	return r.s.State["sku"]
}

// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
func (r *Namespace) SkuName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["skuName"])
}

// Input properties used for looking up and filtering Namespace resources.
type NamespaceState struct {
	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled interface{}
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location interface{}
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name interface{}
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType interface{}
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The ServiceBus Endpoint for this Notification Hub Namespace.
	ServicebusEndpoint interface{}
	// ) A `sku` block as described below.
	Sku interface{}
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName interface{}
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled interface{}
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location interface{}
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name interface{}
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType interface{}
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// ) A `sku` block as described below.
	Sku interface{}
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName interface{}
}
