// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Load Balancer Resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb.html.markdown.
type LoadBalancer struct {
	s *pulumi.ResourceState
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["frontendIpConfigurations"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["frontendIpConfigurations"] = args.FrontendIpConfigurations
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["privateIpAddress"] = nil
	inputs["privateIpAddresses"] = nil
	s, err := ctx.RegisterResource("azure:lb/loadBalancer:LoadBalancer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerState, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["frontendIpConfigurations"] = state.FrontendIpConfigurations
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["privateIpAddress"] = state.PrivateIpAddress
		inputs["privateIpAddresses"] = state.PrivateIpAddresses
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:lb/loadBalancer:LoadBalancer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LoadBalancer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LoadBalancer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// One or multiple `frontendIpConfiguration` blocks as documented below.
func (r *LoadBalancer) FrontendIpConfigurations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["frontendIpConfigurations"])
}

// Specifies the supported Azure Region where the Load Balancer should be created.
func (r *LoadBalancer) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the frontend ip configuration.
func (r *LoadBalancer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
func (r *LoadBalancer) PrivateIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["privateIpAddress"])
}

// The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
func (r *LoadBalancer) PrivateIpAddresses() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["privateIpAddresses"])
}

// The name of the Resource Group in which to create the Load Balancer.
func (r *LoadBalancer) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
func (r *LoadBalancer) Sku() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sku"])
}

// A mapping of tags to assign to the resource.
func (r *LoadBalancer) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering LoadBalancer resources.
type LoadBalancerState struct {
	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations interface{}
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location interface{}
	// Specifies the name of the frontend ip configuration.
	Name interface{}
	// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
	PrivateIpAddress interface{}
	// The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
	PrivateIpAddresses interface{}
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName interface{}
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// One or multiple `frontendIpConfiguration` blocks as documented below.
	FrontendIpConfigurations interface{}
	// Specifies the supported Azure Region where the Load Balancer should be created.
	Location interface{}
	// Specifies the name of the frontend ip configuration.
	Name interface{}
	// The name of the Resource Group in which to create the Load Balancer.
	ResourceGroupName interface{}
	// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
