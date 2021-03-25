// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Firewall Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := network.LookupFirewallPolicy(ctx, &network.LookupFirewallPolicyArgs{
// 			Name:              "existing",
// 			ResourceGroupName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupFirewallPolicy(ctx *pulumi.Context, args *LookupFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyResult, error) {
	var rv LookupFirewallPolicyResult
	err := ctx.Invoke("azure:network/getFirewallPolicy:getFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallPolicy.
type LookupFirewallPolicyArgs struct {
	// The name of this Firewall Policy.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Firewall Policy exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getFirewallPolicy.
type LookupFirewallPolicyResult struct {
	BasePolicyId  string                `pulumi:"basePolicyId"`
	ChildPolicies []string              `pulumi:"childPolicies"`
	Dns           []GetFirewallPolicyDn `pulumi:"dns"`
	Firewalls     []string              `pulumi:"firewalls"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string   `pulumi:"id"`
	Location             string   `pulumi:"location"`
	Name                 string   `pulumi:"name"`
	ResourceGroupName    string   `pulumi:"resourceGroupName"`
	RuleCollectionGroups []string `pulumi:"ruleCollectionGroups"`
	// A mapping of tags assigned to the Firewall Policy.
	Tags                         map[string]string                              `pulumi:"tags"`
	ThreatIntelligenceAllowlists []GetFirewallPolicyThreatIntelligenceAllowlist `pulumi:"threatIntelligenceAllowlists"`
	ThreatIntelligenceMode       string                                         `pulumi:"threatIntelligenceMode"`
}
