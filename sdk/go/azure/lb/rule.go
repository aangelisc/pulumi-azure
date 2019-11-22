// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Load Balancer Rule.
// 
// > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb_rule.html.markdown.
type Rule struct {
	s *pulumi.ResourceState
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOpt) (*Rule, error) {
	if args == nil || args.BackendPort == nil {
		return nil, errors.New("missing required argument 'BackendPort'")
	}
	if args == nil || args.FrontendIpConfigurationName == nil {
		return nil, errors.New("missing required argument 'FrontendIpConfigurationName'")
	}
	if args == nil || args.FrontendPort == nil {
		return nil, errors.New("missing required argument 'FrontendPort'")
	}
	if args == nil || args.LoadbalancerId == nil {
		return nil, errors.New("missing required argument 'LoadbalancerId'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backendAddressPoolId"] = nil
		inputs["backendPort"] = nil
		inputs["disableOutboundSnat"] = nil
		inputs["enableFloatingIp"] = nil
		inputs["frontendIpConfigurationName"] = nil
		inputs["frontendPort"] = nil
		inputs["idleTimeoutInMinutes"] = nil
		inputs["loadDistribution"] = nil
		inputs["loadbalancerId"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["probeId"] = nil
		inputs["protocol"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["backendAddressPoolId"] = args.BackendAddressPoolId
		inputs["backendPort"] = args.BackendPort
		inputs["disableOutboundSnat"] = args.DisableOutboundSnat
		inputs["enableFloatingIp"] = args.EnableFloatingIp
		inputs["frontendIpConfigurationName"] = args.FrontendIpConfigurationName
		inputs["frontendPort"] = args.FrontendPort
		inputs["idleTimeoutInMinutes"] = args.IdleTimeoutInMinutes
		inputs["loadDistribution"] = args.LoadDistribution
		inputs["loadbalancerId"] = args.LoadbalancerId
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["probeId"] = args.ProbeId
		inputs["protocol"] = args.Protocol
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	inputs["frontendIpConfigurationId"] = nil
	s, err := ctx.RegisterResource("azure:lb/rule:Rule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Rule{s: s}, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RuleState, opts ...pulumi.ResourceOpt) (*Rule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backendAddressPoolId"] = state.BackendAddressPoolId
		inputs["backendPort"] = state.BackendPort
		inputs["disableOutboundSnat"] = state.DisableOutboundSnat
		inputs["enableFloatingIp"] = state.EnableFloatingIp
		inputs["frontendIpConfigurationId"] = state.FrontendIpConfigurationId
		inputs["frontendIpConfigurationName"] = state.FrontendIpConfigurationName
		inputs["frontendPort"] = state.FrontendPort
		inputs["idleTimeoutInMinutes"] = state.IdleTimeoutInMinutes
		inputs["loadDistribution"] = state.LoadDistribution
		inputs["loadbalancerId"] = state.LoadbalancerId
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["probeId"] = state.ProbeId
		inputs["protocol"] = state.Protocol
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:lb/rule:Rule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Rule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Rule) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Rule) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
func (r *Rule) BackendAddressPoolId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backendAddressPoolId"])
}

// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
func (r *Rule) BackendPort() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["backendPort"])
}

// Indicates whether outbound snat is disabled or enabled. Default false.
func (r *Rule) DisableOutboundSnat() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["disableOutboundSnat"])
}

// Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.
func (r *Rule) EnableFloatingIp() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableFloatingIp"])
}

func (r *Rule) FrontendIpConfigurationId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["frontendIpConfigurationId"])
}

// The name of the frontend IP configuration to which the rule is associated.
func (r *Rule) FrontendIpConfigurationName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["frontendIpConfigurationName"])
}

// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
func (r *Rule) FrontendPort() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["frontendPort"])
}

// Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.
func (r *Rule) IdleTimeoutInMinutes() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["idleTimeoutInMinutes"])
}

// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
func (r *Rule) LoadDistribution() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadDistribution"])
}

// The ID of the Load Balancer in which to create the Rule.
func (r *Rule) LoadbalancerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadbalancerId"])
}

func (r *Rule) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the LB Rule.
func (r *Rule) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A reference to a Probe used by this Load Balancing Rule.
func (r *Rule) ProbeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["probeId"])
}

// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
func (r *Rule) Protocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocol"])
}

// The name of the resource group in which to create the resource.
func (r *Rule) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering Rule resources.
type RuleState struct {
	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId interface{}
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort interface{}
	// Indicates whether outbound snat is disabled or enabled. Default false.
	DisableOutboundSnat interface{}
	// Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.
	EnableFloatingIp interface{}
	FrontendIpConfigurationId interface{}
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName interface{}
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort interface{}
	// Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.
	IdleTimeoutInMinutes interface{}
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution interface{}
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId interface{}
	Location interface{}
	// Specifies the name of the LB Rule.
	Name interface{}
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId interface{}
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol interface{}
	// The name of the resource group in which to create the resource.
	ResourceGroupName interface{}
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// A reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolId interface{}
	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	BackendPort interface{}
	// Indicates whether outbound snat is disabled or enabled. Default false.
	DisableOutboundSnat interface{}
	// Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.
	EnableFloatingIp interface{}
	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIpConfigurationName interface{}
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	FrontendPort interface{}
	// Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.
	IdleTimeoutInMinutes interface{}
	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
	LoadDistribution interface{}
	// The ID of the Load Balancer in which to create the Rule.
	LoadbalancerId interface{}
	Location interface{}
	// Specifies the name of the LB Rule.
	Name interface{}
	// A reference to a Probe used by this Load Balancing Rule.
	ProbeId interface{}
	// The transport protocol for the external endpoint. Possible values are `Tcp`, `Udp` or `All`.
	Protocol interface{}
	// The name of the resource group in which to create the resource.
	ResourceGroupName interface{}
}
