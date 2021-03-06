// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Network Watcher.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		_, err = network.NewNetworkWatcher(ctx, "exampleNetworkWatcher", &network.NetworkWatcherArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// Network Watchers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkWatcher:NetworkWatcher watcher1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1
// ```
type NetworkWatcher struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Network Watcher. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNetworkWatcher registers a new resource with the given unique name, arguments, and options.
func NewNetworkWatcher(ctx *pulumi.Context,
	name string, args *NetworkWatcherArgs, opts ...pulumi.ResourceOption) (*NetworkWatcher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NetworkWatcher
	err := ctx.RegisterResource("azure:network/networkWatcher:NetworkWatcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkWatcher gets an existing NetworkWatcher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkWatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkWatcherState, opts ...pulumi.ResourceOption) (*NetworkWatcher, error) {
	var resource NetworkWatcher
	err := ctx.ReadResource("azure:network/networkWatcher:NetworkWatcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkWatcher resources.
type networkWatcherState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Network Watcher. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type NetworkWatcherState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Network Watcher. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkWatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherState)(nil)).Elem()
}

type networkWatcherArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Network Watcher. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkWatcher resource.
type NetworkWatcherArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Network Watcher. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkWatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherArgs)(nil)).Elem()
}

type NetworkWatcherInput interface {
	pulumi.Input

	ToNetworkWatcherOutput() NetworkWatcherOutput
	ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput
}

func (*NetworkWatcher) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkWatcher)(nil))
}

func (i *NetworkWatcher) ToNetworkWatcherOutput() NetworkWatcherOutput {
	return i.ToNetworkWatcherOutputWithContext(context.Background())
}

func (i *NetworkWatcher) ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherOutput)
}

func (i *NetworkWatcher) ToNetworkWatcherPtrOutput() NetworkWatcherPtrOutput {
	return i.ToNetworkWatcherPtrOutputWithContext(context.Background())
}

func (i *NetworkWatcher) ToNetworkWatcherPtrOutputWithContext(ctx context.Context) NetworkWatcherPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherPtrOutput)
}

type NetworkWatcherPtrInput interface {
	pulumi.Input

	ToNetworkWatcherPtrOutput() NetworkWatcherPtrOutput
	ToNetworkWatcherPtrOutputWithContext(ctx context.Context) NetworkWatcherPtrOutput
}

type networkWatcherPtrType NetworkWatcherArgs

func (*networkWatcherPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcher)(nil))
}

func (i *networkWatcherPtrType) ToNetworkWatcherPtrOutput() NetworkWatcherPtrOutput {
	return i.ToNetworkWatcherPtrOutputWithContext(context.Background())
}

func (i *networkWatcherPtrType) ToNetworkWatcherPtrOutputWithContext(ctx context.Context) NetworkWatcherPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherPtrOutput)
}

// NetworkWatcherArrayInput is an input type that accepts NetworkWatcherArray and NetworkWatcherArrayOutput values.
// You can construct a concrete instance of `NetworkWatcherArrayInput` via:
//
//          NetworkWatcherArray{ NetworkWatcherArgs{...} }
type NetworkWatcherArrayInput interface {
	pulumi.Input

	ToNetworkWatcherArrayOutput() NetworkWatcherArrayOutput
	ToNetworkWatcherArrayOutputWithContext(context.Context) NetworkWatcherArrayOutput
}

type NetworkWatcherArray []NetworkWatcherInput

func (NetworkWatcherArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkWatcher)(nil))
}

func (i NetworkWatcherArray) ToNetworkWatcherArrayOutput() NetworkWatcherArrayOutput {
	return i.ToNetworkWatcherArrayOutputWithContext(context.Background())
}

func (i NetworkWatcherArray) ToNetworkWatcherArrayOutputWithContext(ctx context.Context) NetworkWatcherArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherArrayOutput)
}

// NetworkWatcherMapInput is an input type that accepts NetworkWatcherMap and NetworkWatcherMapOutput values.
// You can construct a concrete instance of `NetworkWatcherMapInput` via:
//
//          NetworkWatcherMap{ "key": NetworkWatcherArgs{...} }
type NetworkWatcherMapInput interface {
	pulumi.Input

	ToNetworkWatcherMapOutput() NetworkWatcherMapOutput
	ToNetworkWatcherMapOutputWithContext(context.Context) NetworkWatcherMapOutput
}

type NetworkWatcherMap map[string]NetworkWatcherInput

func (NetworkWatcherMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkWatcher)(nil))
}

func (i NetworkWatcherMap) ToNetworkWatcherMapOutput() NetworkWatcherMapOutput {
	return i.ToNetworkWatcherMapOutputWithContext(context.Background())
}

func (i NetworkWatcherMap) ToNetworkWatcherMapOutputWithContext(ctx context.Context) NetworkWatcherMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherMapOutput)
}

type NetworkWatcherOutput struct {
	*pulumi.OutputState
}

func (NetworkWatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkWatcher)(nil))
}

func (o NetworkWatcherOutput) ToNetworkWatcherOutput() NetworkWatcherOutput {
	return o
}

func (o NetworkWatcherOutput) ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput {
	return o
}

func (o NetworkWatcherOutput) ToNetworkWatcherPtrOutput() NetworkWatcherPtrOutput {
	return o.ToNetworkWatcherPtrOutputWithContext(context.Background())
}

func (o NetworkWatcherOutput) ToNetworkWatcherPtrOutputWithContext(ctx context.Context) NetworkWatcherPtrOutput {
	return o.ApplyT(func(v NetworkWatcher) *NetworkWatcher {
		return &v
	}).(NetworkWatcherPtrOutput)
}

type NetworkWatcherPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkWatcherPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcher)(nil))
}

func (o NetworkWatcherPtrOutput) ToNetworkWatcherPtrOutput() NetworkWatcherPtrOutput {
	return o
}

func (o NetworkWatcherPtrOutput) ToNetworkWatcherPtrOutputWithContext(ctx context.Context) NetworkWatcherPtrOutput {
	return o
}

type NetworkWatcherArrayOutput struct{ *pulumi.OutputState }

func (NetworkWatcherArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkWatcher)(nil))
}

func (o NetworkWatcherArrayOutput) ToNetworkWatcherArrayOutput() NetworkWatcherArrayOutput {
	return o
}

func (o NetworkWatcherArrayOutput) ToNetworkWatcherArrayOutputWithContext(ctx context.Context) NetworkWatcherArrayOutput {
	return o
}

func (o NetworkWatcherArrayOutput) Index(i pulumi.IntInput) NetworkWatcherOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkWatcher {
		return vs[0].([]NetworkWatcher)[vs[1].(int)]
	}).(NetworkWatcherOutput)
}

type NetworkWatcherMapOutput struct{ *pulumi.OutputState }

func (NetworkWatcherMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkWatcher)(nil))
}

func (o NetworkWatcherMapOutput) ToNetworkWatcherMapOutput() NetworkWatcherMapOutput {
	return o
}

func (o NetworkWatcherMapOutput) ToNetworkWatcherMapOutputWithContext(ctx context.Context) NetworkWatcherMapOutput {
	return o
}

func (o NetworkWatcherMapOutput) MapIndex(k pulumi.StringInput) NetworkWatcherOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkWatcher {
		return vs[0].(map[string]NetworkWatcher)[vs[1].(string)]
	}).(NetworkWatcherOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkWatcherOutput{})
	pulumi.RegisterOutputType(NetworkWatcherPtrOutput{})
	pulumi.RegisterOutputType(NetworkWatcherArrayOutput{})
	pulumi.RegisterOutputType(NetworkWatcherMapOutput{})
}
