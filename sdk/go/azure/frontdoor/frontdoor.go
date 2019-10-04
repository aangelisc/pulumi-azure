// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package frontdoor

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Front Door instance.
// 
// Azure Front Door Service is Microsoft's highly available and scalable web application acceleration platform and global HTTP(s) load balancer. It provides built-in DDoS protection and application layer security and caching. Front Door enables you to build applications that maximize and automate high-availability and performance for your end-users. Use Front Door with Azure services including Web/Mobile Apps, Cloud Services and Virtual Machines – or combine it with on-premises services for hybrid deployments and smooth cloud migration.
// 
// Below are some of the key scenarios that Azure Front Door Service addresses: 
// * Use Front Door to improve application scale and availability with instant multi-region failover
// * Use Front Door to improve application performance with SSL offload and routing requests to the fastest available application backend.
// * Use Front Door for application layer security and DDoS protection for your application.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/frontdoor.html.markdown.
type Frontdoor struct {
	s *pulumi.ResourceState
}

// NewFrontdoor registers a new resource with the given unique name, arguments, and options.
func NewFrontdoor(ctx *pulumi.Context,
	name string, args *FrontdoorArgs, opts ...pulumi.ResourceOpt) (*Frontdoor, error) {
	if args == nil || args.BackendPools == nil {
		return nil, errors.New("missing required argument 'BackendPools'")
	}
	if args == nil || args.BackendPoolHealthProbes == nil {
		return nil, errors.New("missing required argument 'BackendPoolHealthProbes'")
	}
	if args == nil || args.BackendPoolLoadBalancings == nil {
		return nil, errors.New("missing required argument 'BackendPoolLoadBalancings'")
	}
	if args == nil || args.EnforceBackendPoolsCertificateNameCheck == nil {
		return nil, errors.New("missing required argument 'EnforceBackendPoolsCertificateNameCheck'")
	}
	if args == nil || args.FrontendEndpoints == nil {
		return nil, errors.New("missing required argument 'FrontendEndpoints'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RoutingRules == nil {
		return nil, errors.New("missing required argument 'RoutingRules'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backendPools"] = nil
		inputs["backendPoolHealthProbes"] = nil
		inputs["backendPoolLoadBalancings"] = nil
		inputs["enforceBackendPoolsCertificateNameCheck"] = nil
		inputs["friendlyName"] = nil
		inputs["frontendEndpoints"] = nil
		inputs["loadBalancerEnabled"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["routingRules"] = nil
		inputs["tags"] = nil
	} else {
		inputs["backendPools"] = args.BackendPools
		inputs["backendPoolHealthProbes"] = args.BackendPoolHealthProbes
		inputs["backendPoolLoadBalancings"] = args.BackendPoolLoadBalancings
		inputs["enforceBackendPoolsCertificateNameCheck"] = args.EnforceBackendPoolsCertificateNameCheck
		inputs["friendlyName"] = args.FriendlyName
		inputs["frontendEndpoints"] = args.FrontendEndpoints
		inputs["loadBalancerEnabled"] = args.LoadBalancerEnabled
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["routingRules"] = args.RoutingRules
		inputs["tags"] = args.Tags
	}
	inputs["cname"] = nil
	s, err := ctx.RegisterResource("azure:frontdoor/frontdoor:Frontdoor", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Frontdoor{s: s}, nil
}

// GetFrontdoor gets an existing Frontdoor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFrontdoor(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FrontdoorState, opts ...pulumi.ResourceOpt) (*Frontdoor, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backendPools"] = state.BackendPools
		inputs["backendPoolHealthProbes"] = state.BackendPoolHealthProbes
		inputs["backendPoolLoadBalancings"] = state.BackendPoolLoadBalancings
		inputs["cname"] = state.Cname
		inputs["enforceBackendPoolsCertificateNameCheck"] = state.EnforceBackendPoolsCertificateNameCheck
		inputs["friendlyName"] = state.FriendlyName
		inputs["frontendEndpoints"] = state.FrontendEndpoints
		inputs["loadBalancerEnabled"] = state.LoadBalancerEnabled
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["routingRules"] = state.RoutingRules
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:frontdoor/frontdoor:Frontdoor", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Frontdoor{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Frontdoor) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Frontdoor) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A `backendPool` block as defined below.
func (r *Frontdoor) BackendPools() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backendPools"])
}

// A `backendPoolHealthProbe` block as defined below.
func (r *Frontdoor) BackendPoolHealthProbes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backendPoolHealthProbes"])
}

// A `backendPoolLoadBalancing` block as defined below.
func (r *Frontdoor) BackendPoolLoadBalancings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backendPoolLoadBalancings"])
}

// The host that each frontendEndpoint must CNAME to.
func (r *Frontdoor) Cname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cname"])
}

// Whether to enforce certificate name check on HTTPS requests to all backend pools. No effect on non-HTTPS requests. Permitted values are `true` or `false`.
func (r *Frontdoor) EnforceBackendPoolsCertificateNameCheck() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enforceBackendPoolsCertificateNameCheck"])
}

// A friendly name for the Front Door service.
func (r *Frontdoor) FriendlyName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["friendlyName"])
}

// A `frontendEndpoint` block as defined below.
func (r *Frontdoor) FrontendEndpoints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["frontendEndpoints"])
}

// Operational status of the Front Door load balancer. Permitted values are `true` or `false` Defaults to `true`.
func (r *Frontdoor) LoadBalancerEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["loadBalancerEnabled"])
}

// Resource location. Changing this forces a new resource to be created.
func (r *Frontdoor) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Name of the Front Door which is globally unique. Changing this forces a new resource to be created.
func (r *Frontdoor) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Name of the Resource group within the Azure subscription. Changing this forces a new resource to be created.
func (r *Frontdoor) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `routingRule` block as defined below.
func (r *Frontdoor) RoutingRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["routingRules"])
}

// Resource tags.
func (r *Frontdoor) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Frontdoor resources.
type FrontdoorState struct {
	// A `backendPool` block as defined below.
	BackendPools interface{}
	// A `backendPoolHealthProbe` block as defined below.
	BackendPoolHealthProbes interface{}
	// A `backendPoolLoadBalancing` block as defined below.
	BackendPoolLoadBalancings interface{}
	// The host that each frontendEndpoint must CNAME to.
	Cname interface{}
	// Whether to enforce certificate name check on HTTPS requests to all backend pools. No effect on non-HTTPS requests. Permitted values are `true` or `false`.
	EnforceBackendPoolsCertificateNameCheck interface{}
	// A friendly name for the Front Door service.
	FriendlyName interface{}
	// A `frontendEndpoint` block as defined below.
	FrontendEndpoints interface{}
	// Operational status of the Front Door load balancer. Permitted values are `true` or `false` Defaults to `true`.
	LoadBalancerEnabled interface{}
	// Resource location. Changing this forces a new resource to be created.
	Location interface{}
	// Name of the Front Door which is globally unique. Changing this forces a new resource to be created.
	Name interface{}
	// Name of the Resource group within the Azure subscription. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `routingRule` block as defined below.
	RoutingRules interface{}
	// Resource tags.
	Tags interface{}
}

// The set of arguments for constructing a Frontdoor resource.
type FrontdoorArgs struct {
	// A `backendPool` block as defined below.
	BackendPools interface{}
	// A `backendPoolHealthProbe` block as defined below.
	BackendPoolHealthProbes interface{}
	// A `backendPoolLoadBalancing` block as defined below.
	BackendPoolLoadBalancings interface{}
	// Whether to enforce certificate name check on HTTPS requests to all backend pools. No effect on non-HTTPS requests. Permitted values are `true` or `false`.
	EnforceBackendPoolsCertificateNameCheck interface{}
	// A friendly name for the Front Door service.
	FriendlyName interface{}
	// A `frontendEndpoint` block as defined below.
	FrontendEndpoints interface{}
	// Operational status of the Front Door load balancer. Permitted values are `true` or `false` Defaults to `true`.
	LoadBalancerEnabled interface{}
	// Resource location. Changing this forces a new resource to be created.
	Location interface{}
	// Name of the Front Door which is globally unique. Changing this forces a new resource to be created.
	Name interface{}
	// Name of the Resource group within the Azure subscription. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `routingRule` block as defined below.
	RoutingRules interface{}
	// Resource tags.
	Tags interface{}
}
