// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Network Rule Collection within an Azure Firewall.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/firewall_network_rule_collection.html.markdown.
type FirewallNetworkRuleCollection struct {
	s *pulumi.ResourceState
}

// NewFirewallNetworkRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewFirewallNetworkRuleCollection(ctx *pulumi.Context,
	name string, args *FirewallNetworkRuleCollectionArgs, opts ...pulumi.ResourceOpt) (*FirewallNetworkRuleCollection, error) {
	if args == nil || args.Action == nil {
		return nil, errors.New("missing required argument 'Action'")
	}
	if args == nil || args.AzureFirewallName == nil {
		return nil, errors.New("missing required argument 'AzureFirewallName'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["action"] = nil
		inputs["azureFirewallName"] = nil
		inputs["name"] = nil
		inputs["priority"] = nil
		inputs["resourceGroupName"] = nil
		inputs["rules"] = nil
	} else {
		inputs["action"] = args.Action
		inputs["azureFirewallName"] = args.AzureFirewallName
		inputs["name"] = args.Name
		inputs["priority"] = args.Priority
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["rules"] = args.Rules
	}
	s, err := ctx.RegisterResource("azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirewallNetworkRuleCollection{s: s}, nil
}

// GetFirewallNetworkRuleCollection gets an existing FirewallNetworkRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallNetworkRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FirewallNetworkRuleCollectionState, opts ...pulumi.ResourceOpt) (*FirewallNetworkRuleCollection, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["action"] = state.Action
		inputs["azureFirewallName"] = state.AzureFirewallName
		inputs["name"] = state.Name
		inputs["priority"] = state.Priority
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["rules"] = state.Rules
	}
	s, err := ctx.ReadResource("azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirewallNetworkRuleCollection{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FirewallNetworkRuleCollection) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FirewallNetworkRuleCollection) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
func (r *FirewallNetworkRuleCollection) Action() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["action"])
}

// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
func (r *FirewallNetworkRuleCollection) AzureFirewallName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["azureFirewallName"])
}

// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
func (r *FirewallNetworkRuleCollection) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
func (r *FirewallNetworkRuleCollection) Priority() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["priority"])
}

// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
func (r *FirewallNetworkRuleCollection) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// One or more `rule` blocks as defined below.
func (r *FirewallNetworkRuleCollection) Rules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["rules"])
}

// Input properties used for looking up and filtering FirewallNetworkRuleCollection resources.
type FirewallNetworkRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action interface{}
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName interface{}
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority interface{}
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// One or more `rule` blocks as defined below.
	Rules interface{}
}

// The set of arguments for constructing a FirewallNetworkRuleCollection resource.
type FirewallNetworkRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action interface{}
	// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName interface{}
	// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority interface{}
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// One or more `rule` blocks as defined below.
	Rules interface{}
}
