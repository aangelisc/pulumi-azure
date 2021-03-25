// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a ServiceBus Topic.
//
// **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
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
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"source": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
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
// Service Bus Topics can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicebus/topic:Topic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1
// ```
type Topic struct {
	pulumi.CustomResourceState

	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringOutput `pulumi:"autoDeleteOnIdle"`
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl pulumi.StringOutput `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress pulumi.BoolPtrOutput `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning pulumi.BoolPtrOutput `pulumi:"enablePartitioning"`
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntOutput `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection pulumi.BoolPtrOutput `pulumi:"requiresDuplicateDetection"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering pulumi.BoolPtrOutput `pulumi:"supportOrdering"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/topic:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure:servicebus/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure:servicebus/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status *string `pulumi:"status"`
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering *bool `pulumi:"supportOrdering"`
}

type TopicState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning pulumi.BoolPtrInput
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrInput
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering pulumi.BoolPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status *string `pulumi:"status"`
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering *bool `pulumi:"supportOrdering"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning pulumi.BoolPtrInput
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrInput
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering pulumi.BoolPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *Topic) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPtrOutput)
}

type TopicPtrInput interface {
	pulumi.Input

	ToTopicPtrOutput() TopicPtrOutput
	ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput
}

type topicPtrType TopicArgs

func (*topicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (i *topicPtrType) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *topicPtrType) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPtrOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//          TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Topic)(nil))
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//          TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Topic)(nil))
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct {
	*pulumi.OutputState
}

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func (o TopicOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o.ToTopicPtrOutputWithContext(context.Background())
}

func (o TopicOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o.ApplyT(func(v Topic) *Topic {
		return &v
	}).(TopicPtrOutput)
}

type TopicPtrOutput struct {
	*pulumi.OutputState
}

func (TopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (o TopicPtrOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o
}

func (o TopicPtrOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Topic)(nil))
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Topic {
		return vs[0].([]Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Topic)(nil))
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Topic {
		return vs[0].(map[string]Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicPtrOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
