// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Stream Analytics Output to a ServiceBus Topic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			ResourceGroupName:  pulumi.String(exampleResourceGroup.Name),
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewOutputServicebusTopic(ctx, "exampleOutputServicebusTopic", &streamanalytics.OutputServicebusTopicArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			TopicName:              exampleTopic.Name,
// 			ServicebusNamespace:    exampleNamespace.Name,
// 			SharedAccessPolicyKey:  exampleNamespace.DefaultPrimaryKey,
// 			SharedAccessPolicyName: pulumi.String("RootManageSharedAccessKey"),
// 			Serialization: &streamanalytics.OutputServicebusTopicSerializationArgs{
// 				Format: pulumi.String("Avro"),
// 			},
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
// Stream Analytics Output ServiceBus Topic's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
// ```
type OutputServicebusTopic struct {
	pulumi.CustomResourceState

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationOutput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewOutputServicebusTopic registers a new resource with the given unique name, arguments, and options.
func NewOutputServicebusTopic(ctx *pulumi.Context,
	name string, args *OutputServicebusTopicArgs, opts ...pulumi.ResourceOption) (*OutputServicebusTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Serialization == nil {
		return nil, errors.New("invalid value for required argument 'Serialization'")
	}
	if args.ServicebusNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServicebusNamespace'")
	}
	if args.SharedAccessPolicyKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyKey'")
	}
	if args.SharedAccessPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyName'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	var resource OutputServicebusTopic
	err := ctx.RegisterResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputServicebusTopic gets an existing OutputServicebusTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputServicebusTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputServicebusTopicState, opts ...pulumi.ResourceOption) (*OutputServicebusTopic, error) {
	var resource OutputServicebusTopic
	err := ctx.ReadResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputServicebusTopic resources.
type outputServicebusTopicState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *OutputServicebusTopicSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace *string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName *string `pulumi:"topicName"`
}

type OutputServicebusTopicState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationPtrInput
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
	// The name of the Service Bus Topic.
	TopicName pulumi.StringPtrInput
}

func (OutputServicebusTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputServicebusTopicState)(nil)).Elem()
}

type outputServicebusTopicArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a OutputServicebusTopic resource.
type OutputServicebusTopicArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationInput
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
	// The name of the Service Bus Topic.
	TopicName pulumi.StringInput
}

func (OutputServicebusTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputServicebusTopicArgs)(nil)).Elem()
}

type OutputServicebusTopicInput interface {
	pulumi.Input

	ToOutputServicebusTopicOutput() OutputServicebusTopicOutput
	ToOutputServicebusTopicOutputWithContext(ctx context.Context) OutputServicebusTopicOutput
}

func (*OutputServicebusTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputServicebusTopic)(nil))
}

func (i *OutputServicebusTopic) ToOutputServicebusTopicOutput() OutputServicebusTopicOutput {
	return i.ToOutputServicebusTopicOutputWithContext(context.Background())
}

func (i *OutputServicebusTopic) ToOutputServicebusTopicOutputWithContext(ctx context.Context) OutputServicebusTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputServicebusTopicOutput)
}

func (i *OutputServicebusTopic) ToOutputServicebusTopicPtrOutput() OutputServicebusTopicPtrOutput {
	return i.ToOutputServicebusTopicPtrOutputWithContext(context.Background())
}

func (i *OutputServicebusTopic) ToOutputServicebusTopicPtrOutputWithContext(ctx context.Context) OutputServicebusTopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputServicebusTopicPtrOutput)
}

type OutputServicebusTopicPtrInput interface {
	pulumi.Input

	ToOutputServicebusTopicPtrOutput() OutputServicebusTopicPtrOutput
	ToOutputServicebusTopicPtrOutputWithContext(ctx context.Context) OutputServicebusTopicPtrOutput
}

type outputServicebusTopicPtrType OutputServicebusTopicArgs

func (*outputServicebusTopicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputServicebusTopic)(nil))
}

func (i *outputServicebusTopicPtrType) ToOutputServicebusTopicPtrOutput() OutputServicebusTopicPtrOutput {
	return i.ToOutputServicebusTopicPtrOutputWithContext(context.Background())
}

func (i *outputServicebusTopicPtrType) ToOutputServicebusTopicPtrOutputWithContext(ctx context.Context) OutputServicebusTopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputServicebusTopicPtrOutput)
}

// OutputServicebusTopicArrayInput is an input type that accepts OutputServicebusTopicArray and OutputServicebusTopicArrayOutput values.
// You can construct a concrete instance of `OutputServicebusTopicArrayInput` via:
//
//          OutputServicebusTopicArray{ OutputServicebusTopicArgs{...} }
type OutputServicebusTopicArrayInput interface {
	pulumi.Input

	ToOutputServicebusTopicArrayOutput() OutputServicebusTopicArrayOutput
	ToOutputServicebusTopicArrayOutputWithContext(context.Context) OutputServicebusTopicArrayOutput
}

type OutputServicebusTopicArray []OutputServicebusTopicInput

func (OutputServicebusTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OutputServicebusTopic)(nil))
}

func (i OutputServicebusTopicArray) ToOutputServicebusTopicArrayOutput() OutputServicebusTopicArrayOutput {
	return i.ToOutputServicebusTopicArrayOutputWithContext(context.Background())
}

func (i OutputServicebusTopicArray) ToOutputServicebusTopicArrayOutputWithContext(ctx context.Context) OutputServicebusTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputServicebusTopicArrayOutput)
}

// OutputServicebusTopicMapInput is an input type that accepts OutputServicebusTopicMap and OutputServicebusTopicMapOutput values.
// You can construct a concrete instance of `OutputServicebusTopicMapInput` via:
//
//          OutputServicebusTopicMap{ "key": OutputServicebusTopicArgs{...} }
type OutputServicebusTopicMapInput interface {
	pulumi.Input

	ToOutputServicebusTopicMapOutput() OutputServicebusTopicMapOutput
	ToOutputServicebusTopicMapOutputWithContext(context.Context) OutputServicebusTopicMapOutput
}

type OutputServicebusTopicMap map[string]OutputServicebusTopicInput

func (OutputServicebusTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OutputServicebusTopic)(nil))
}

func (i OutputServicebusTopicMap) ToOutputServicebusTopicMapOutput() OutputServicebusTopicMapOutput {
	return i.ToOutputServicebusTopicMapOutputWithContext(context.Background())
}

func (i OutputServicebusTopicMap) ToOutputServicebusTopicMapOutputWithContext(ctx context.Context) OutputServicebusTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputServicebusTopicMapOutput)
}

type OutputServicebusTopicOutput struct {
	*pulumi.OutputState
}

func (OutputServicebusTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputServicebusTopic)(nil))
}

func (o OutputServicebusTopicOutput) ToOutputServicebusTopicOutput() OutputServicebusTopicOutput {
	return o
}

func (o OutputServicebusTopicOutput) ToOutputServicebusTopicOutputWithContext(ctx context.Context) OutputServicebusTopicOutput {
	return o
}

func (o OutputServicebusTopicOutput) ToOutputServicebusTopicPtrOutput() OutputServicebusTopicPtrOutput {
	return o.ToOutputServicebusTopicPtrOutputWithContext(context.Background())
}

func (o OutputServicebusTopicOutput) ToOutputServicebusTopicPtrOutputWithContext(ctx context.Context) OutputServicebusTopicPtrOutput {
	return o.ApplyT(func(v OutputServicebusTopic) *OutputServicebusTopic {
		return &v
	}).(OutputServicebusTopicPtrOutput)
}

type OutputServicebusTopicPtrOutput struct {
	*pulumi.OutputState
}

func (OutputServicebusTopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputServicebusTopic)(nil))
}

func (o OutputServicebusTopicPtrOutput) ToOutputServicebusTopicPtrOutput() OutputServicebusTopicPtrOutput {
	return o
}

func (o OutputServicebusTopicPtrOutput) ToOutputServicebusTopicPtrOutputWithContext(ctx context.Context) OutputServicebusTopicPtrOutput {
	return o
}

type OutputServicebusTopicArrayOutput struct{ *pulumi.OutputState }

func (OutputServicebusTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputServicebusTopic)(nil))
}

func (o OutputServicebusTopicArrayOutput) ToOutputServicebusTopicArrayOutput() OutputServicebusTopicArrayOutput {
	return o
}

func (o OutputServicebusTopicArrayOutput) ToOutputServicebusTopicArrayOutputWithContext(ctx context.Context) OutputServicebusTopicArrayOutput {
	return o
}

func (o OutputServicebusTopicArrayOutput) Index(i pulumi.IntInput) OutputServicebusTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputServicebusTopic {
		return vs[0].([]OutputServicebusTopic)[vs[1].(int)]
	}).(OutputServicebusTopicOutput)
}

type OutputServicebusTopicMapOutput struct{ *pulumi.OutputState }

func (OutputServicebusTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputServicebusTopic)(nil))
}

func (o OutputServicebusTopicMapOutput) ToOutputServicebusTopicMapOutput() OutputServicebusTopicMapOutput {
	return o
}

func (o OutputServicebusTopicMapOutput) ToOutputServicebusTopicMapOutputWithContext(ctx context.Context) OutputServicebusTopicMapOutput {
	return o
}

func (o OutputServicebusTopicMapOutput) MapIndex(k pulumi.StringInput) OutputServicebusTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputServicebusTopic {
		return vs[0].(map[string]OutputServicebusTopic)[vs[1].(string)]
	}).(OutputServicebusTopicOutput)
}

func init() {
	pulumi.RegisterOutputType(OutputServicebusTopicOutput{})
	pulumi.RegisterOutputType(OutputServicebusTopicPtrOutput{})
	pulumi.RegisterOutputType(OutputServicebusTopicArrayOutput{})
	pulumi.RegisterOutputType(OutputServicebusTopicMapOutput{})
}
