// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Load Balancer NAT Rule.
//
// > **NOTE:** This resource cannot be used with with virtual machine scale sets, instead use the `lb.NatPool` resource.
//
// > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/lb"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FrontendIpConfigurations: lb.LoadBalancerFrontendIpConfigurationArray{
// 				&lb.LoadBalancerFrontendIpConfigurationArgs{
// 					Name:              pulumi.String("PublicIPAddress"),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewNatRule(ctx, "exampleNatRule", &lb.NatRuleArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			LoadbalancerId:              exampleLoadBalancer.ID(),
// 			Protocol:                    pulumi.String("Tcp"),
// 			FrontendPort:                pulumi.Int(3389),
// 			BackendPort:                 pulumi.Int(3389),
// 			FrontendIpConfigurationName: pulumi.String("PublicIPAddress"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Load Balancer NAT Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:lb/natRule:NatRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
// ```
type NatRule struct {
	pulumi.CustomResourceState

	BackendIpConfigurationId pulumi.StringOutput `pulumi:"backendIpConfigurationId"`
	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort pulumi.IntOutput `pulumi:"backendPort"`
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp pulumi.BoolOutput `pulumi:"enableFloatingIp"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset            pulumi.BoolPtrOutput `pulumi:"enableTcpReset"`
	FrontendIpConfigurationId pulumi.StringOutput  `pulumi:"frontendIpConfigurationId"`
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName pulumi.StringOutput `pulumi:"frontendIpConfigurationName"`
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort pulumi.IntOutput `pulumi:"frontendPort"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes pulumi.IntOutput `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the NAT Rule.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the NAT Rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewNatRule registers a new resource with the given unique name, arguments, and options.
func NewNatRule(ctx *pulumi.Context,
	name string, args *NatRuleArgs, opts ...pulumi.ResourceOption) (*NatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendPort == nil {
		return nil, errors.New("invalid value for required argument 'BackendPort'")
	}
	if args.FrontendIpConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'FrontendIpConfigurationName'")
	}
	if args.FrontendPort == nil {
		return nil, errors.New("invalid value for required argument 'FrontendPort'")
	}
	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NatRule
	err := ctx.RegisterResource("azure:lb/natRule:NatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatRule gets an existing NatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatRuleState, opts ...pulumi.ResourceOption) (*NatRule, error) {
	var resource NatRule
	err := ctx.ReadResource("azure:lb/natRule:NatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatRule resources.
type natRuleState struct {
	BackendIpConfigurationId *string `pulumi:"backendIpConfigurationId"`
	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort *int `pulumi:"backendPort"`
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp *bool `pulumi:"enableFloatingIp"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset            *bool   `pulumi:"enableTcpReset"`
	FrontendIpConfigurationId *string `pulumi:"frontendIpConfigurationId"`
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName *string `pulumi:"frontendIpConfigurationName"`
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort *int `pulumi:"frontendPort"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the NAT Rule.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the NAT Rule.
	Name *string `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type NatRuleState struct {
	BackendIpConfigurationId pulumi.StringPtrInput
	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort pulumi.IntPtrInput
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp pulumi.BoolPtrInput
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset            pulumi.BoolPtrInput
	FrontendIpConfigurationId pulumi.StringPtrInput
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName pulumi.StringPtrInput
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort pulumi.IntPtrInput
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The ID of the Load Balancer in which to create the NAT Rule.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the NAT Rule.
	Name pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringPtrInput
}

func (NatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*natRuleState)(nil)).Elem()
}

type natRuleArgs struct {
	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort int `pulumi:"backendPort"`
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp *bool `pulumi:"enableFloatingIp"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName string `pulumi:"frontendIpConfigurationName"`
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort int `pulumi:"frontendPort"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the NAT Rule.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the NAT Rule.
	Name *string `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NatRule resource.
type NatRuleArgs struct {
	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort pulumi.IntInput
	// Are the Floating IPs enabled for this Load Balncer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	EnableFloatingIp pulumi.BoolPtrInput
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	EnableTcpReset pulumi.BoolPtrInput
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName pulumi.StringInput
	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort pulumi.IntInput
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30` minutes. Defaults to `4` minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The ID of the Load Balancer in which to create the NAT Rule.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the NAT Rule.
	Name pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
	Protocol pulumi.StringInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
}

func (NatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natRuleArgs)(nil)).Elem()
}

type NatRuleInput interface {
	pulumi.Input

	ToNatRuleOutput() NatRuleOutput
	ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput
}

func (*NatRule) ElementType() reflect.Type {
	return reflect.TypeOf((*NatRule)(nil))
}

func (i *NatRule) ToNatRuleOutput() NatRuleOutput {
	return i.ToNatRuleOutputWithContext(context.Background())
}

func (i *NatRule) ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRuleOutput)
}

func (i *NatRule) ToNatRulePtrOutput() NatRulePtrOutput {
	return i.ToNatRulePtrOutputWithContext(context.Background())
}

func (i *NatRule) ToNatRulePtrOutputWithContext(ctx context.Context) NatRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRulePtrOutput)
}

type NatRulePtrInput interface {
	pulumi.Input

	ToNatRulePtrOutput() NatRulePtrOutput
	ToNatRulePtrOutputWithContext(ctx context.Context) NatRulePtrOutput
}

type natRulePtrType NatRuleArgs

func (*natRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NatRule)(nil))
}

func (i *natRulePtrType) ToNatRulePtrOutput() NatRulePtrOutput {
	return i.ToNatRulePtrOutputWithContext(context.Background())
}

func (i *natRulePtrType) ToNatRulePtrOutputWithContext(ctx context.Context) NatRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRulePtrOutput)
}

// NatRuleArrayInput is an input type that accepts NatRuleArray and NatRuleArrayOutput values.
// You can construct a concrete instance of `NatRuleArrayInput` via:
//
//          NatRuleArray{ NatRuleArgs{...} }
type NatRuleArrayInput interface {
	pulumi.Input

	ToNatRuleArrayOutput() NatRuleArrayOutput
	ToNatRuleArrayOutputWithContext(context.Context) NatRuleArrayOutput
}

type NatRuleArray []NatRuleInput

func (NatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NatRule)(nil))
}

func (i NatRuleArray) ToNatRuleArrayOutput() NatRuleArrayOutput {
	return i.ToNatRuleArrayOutputWithContext(context.Background())
}

func (i NatRuleArray) ToNatRuleArrayOutputWithContext(ctx context.Context) NatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRuleArrayOutput)
}

// NatRuleMapInput is an input type that accepts NatRuleMap and NatRuleMapOutput values.
// You can construct a concrete instance of `NatRuleMapInput` via:
//
//          NatRuleMap{ "key": NatRuleArgs{...} }
type NatRuleMapInput interface {
	pulumi.Input

	ToNatRuleMapOutput() NatRuleMapOutput
	ToNatRuleMapOutputWithContext(context.Context) NatRuleMapOutput
}

type NatRuleMap map[string]NatRuleInput

func (NatRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NatRule)(nil))
}

func (i NatRuleMap) ToNatRuleMapOutput() NatRuleMapOutput {
	return i.ToNatRuleMapOutputWithContext(context.Background())
}

func (i NatRuleMap) ToNatRuleMapOutputWithContext(ctx context.Context) NatRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRuleMapOutput)
}

type NatRuleOutput struct {
	*pulumi.OutputState
}

func (NatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatRule)(nil))
}

func (o NatRuleOutput) ToNatRuleOutput() NatRuleOutput {
	return o
}

func (o NatRuleOutput) ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput {
	return o
}

func (o NatRuleOutput) ToNatRulePtrOutput() NatRulePtrOutput {
	return o.ToNatRulePtrOutputWithContext(context.Background())
}

func (o NatRuleOutput) ToNatRulePtrOutputWithContext(ctx context.Context) NatRulePtrOutput {
	return o.ApplyT(func(v NatRule) *NatRule {
		return &v
	}).(NatRulePtrOutput)
}

type NatRulePtrOutput struct {
	*pulumi.OutputState
}

func (NatRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatRule)(nil))
}

func (o NatRulePtrOutput) ToNatRulePtrOutput() NatRulePtrOutput {
	return o
}

func (o NatRulePtrOutput) ToNatRulePtrOutputWithContext(ctx context.Context) NatRulePtrOutput {
	return o
}

type NatRuleArrayOutput struct{ *pulumi.OutputState }

func (NatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NatRule)(nil))
}

func (o NatRuleArrayOutput) ToNatRuleArrayOutput() NatRuleArrayOutput {
	return o
}

func (o NatRuleArrayOutput) ToNatRuleArrayOutputWithContext(ctx context.Context) NatRuleArrayOutput {
	return o
}

func (o NatRuleArrayOutput) Index(i pulumi.IntInput) NatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NatRule {
		return vs[0].([]NatRule)[vs[1].(int)]
	}).(NatRuleOutput)
}

type NatRuleMapOutput struct{ *pulumi.OutputState }

func (NatRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NatRule)(nil))
}

func (o NatRuleMapOutput) ToNatRuleMapOutput() NatRuleMapOutput {
	return o
}

func (o NatRuleMapOutput) ToNatRuleMapOutputWithContext(ctx context.Context) NatRuleMapOutput {
	return o
}

func (o NatRuleMapOutput) MapIndex(k pulumi.StringInput) NatRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NatRule {
		return vs[0].(map[string]NatRule)[vs[1].(string)]
	}).(NatRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(NatRuleOutput{})
	pulumi.RegisterOutputType(NatRulePtrOutput{})
	pulumi.RegisterOutputType(NatRuleArrayOutput{})
	pulumi.RegisterOutputType(NatRuleMapOutput{})
}
