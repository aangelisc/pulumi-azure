// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a ServiceBus Topic.
// 
// **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_topic.html.markdown.
type Topic struct {
	s *pulumi.ResourceState
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOpt) (*Topic, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoDeleteOnIdle"] = nil
		inputs["defaultMessageTtl"] = nil
		inputs["duplicateDetectionHistoryTimeWindow"] = nil
		inputs["enableBatchedOperations"] = nil
		inputs["enableExpress"] = nil
		inputs["enableFilteringMessagesBeforePublishing"] = nil
		inputs["enablePartitioning"] = nil
		inputs["location"] = nil
		inputs["maxSizeInMegabytes"] = nil
		inputs["name"] = nil
		inputs["namespaceName"] = nil
		inputs["requiresDuplicateDetection"] = nil
		inputs["resourceGroupName"] = nil
		inputs["status"] = nil
		inputs["supportOrdering"] = nil
	} else {
		inputs["autoDeleteOnIdle"] = args.AutoDeleteOnIdle
		inputs["defaultMessageTtl"] = args.DefaultMessageTtl
		inputs["duplicateDetectionHistoryTimeWindow"] = args.DuplicateDetectionHistoryTimeWindow
		inputs["enableBatchedOperations"] = args.EnableBatchedOperations
		inputs["enableExpress"] = args.EnableExpress
		inputs["enableFilteringMessagesBeforePublishing"] = args.EnableFilteringMessagesBeforePublishing
		inputs["enablePartitioning"] = args.EnablePartitioning
		inputs["location"] = args.Location
		inputs["maxSizeInMegabytes"] = args.MaxSizeInMegabytes
		inputs["name"] = args.Name
		inputs["namespaceName"] = args.NamespaceName
		inputs["requiresDuplicateDetection"] = args.RequiresDuplicateDetection
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["status"] = args.Status
		inputs["supportOrdering"] = args.SupportOrdering
	}
	s, err := ctx.RegisterResource("azure:servicebus/topic:Topic", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Topic{s: s}, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TopicState, opts ...pulumi.ResourceOpt) (*Topic, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoDeleteOnIdle"] = state.AutoDeleteOnIdle
		inputs["defaultMessageTtl"] = state.DefaultMessageTtl
		inputs["duplicateDetectionHistoryTimeWindow"] = state.DuplicateDetectionHistoryTimeWindow
		inputs["enableBatchedOperations"] = state.EnableBatchedOperations
		inputs["enableExpress"] = state.EnableExpress
		inputs["enableFilteringMessagesBeforePublishing"] = state.EnableFilteringMessagesBeforePublishing
		inputs["enablePartitioning"] = state.EnablePartitioning
		inputs["location"] = state.Location
		inputs["maxSizeInMegabytes"] = state.MaxSizeInMegabytes
		inputs["name"] = state.Name
		inputs["namespaceName"] = state.NamespaceName
		inputs["requiresDuplicateDetection"] = state.RequiresDuplicateDetection
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["status"] = state.Status
		inputs["supportOrdering"] = state.SupportOrdering
	}
	s, err := ctx.ReadResource("azure:servicebus/topic:Topic", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Topic{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Topic) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Topic) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ISO 8601 timespan duration of the idle interval after which the
// Topic is automatically deleted, minimum of 5 minutes.
func (r *Topic) AutoDeleteOnIdle() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["autoDeleteOnIdle"])
}

// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
// TTL value is set on the message itself.
func (r *Topic) DefaultMessageTtl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultMessageTtl"])
}

// The ISO 8601 timespan duration during which
// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
func (r *Topic) DuplicateDetectionHistoryTimeWindow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["duplicateDetectionHistoryTimeWindow"])
}

// Boolean flag which controls if server-side
// batched operations are enabled. Defaults to false.
func (r *Topic) EnableBatchedOperations() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableBatchedOperations"])
}

// Boolean flag which controls whether Express Entities
// are enabled. An express topic holds a message in memory temporarily before writing
// it to persistent storage. Defaults to false.
func (r *Topic) EnableExpress() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableExpress"])
}

func (r *Topic) EnableFilteringMessagesBeforePublishing() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableFilteringMessagesBeforePublishing"])
}

// Boolean flag which controls whether to enable
// the topic to be partitioned across multiple message brokers. Defaults to false.
// Changing this forces a new resource to be created.
func (r *Topic) EnablePartitioning() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enablePartitioning"])
}

// Specifies the supported Azure location where the resource exists.
// Changing this forces a new resource to be created.
func (r *Topic) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Integer value which controls the size of
// memory allocated for the topic. For supported values see the "Queue/topic size"
// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
func (r *Topic) MaxSizeInMegabytes() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxSizeInMegabytes"])
}

// Specifies the name of the ServiceBus Topic resource. Changing this forces a
// new resource to be created.
func (r *Topic) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the ServiceBus Namespace to create
// this topic in. Changing this forces a new resource to be created.
func (r *Topic) NamespaceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namespaceName"])
}

// Boolean flag which controls whether
// the Topic requires duplicate detection. Defaults to false. Changing this forces
// a new resource to be created.
func (r *Topic) RequiresDuplicateDetection() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["requiresDuplicateDetection"])
}

// The name of the resource group in which to
// create the namespace. Changing this forces a new resource to be created.
func (r *Topic) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
func (r *Topic) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// Boolean flag which controls whether the Topic
// supports ordering. Defaults to false.
func (r *Topic) SupportOrdering() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["supportOrdering"])
}

// Input properties used for looking up and filtering Topic resources.
type TopicState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle interface{}
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl interface{}
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow interface{}
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations interface{}
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress interface{}
	EnableFilteringMessagesBeforePublishing interface{}
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning interface{}
	// Specifies the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location interface{}
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes interface{}
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name interface{}
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName interface{}
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection interface{}
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status interface{}
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering interface{}
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle interface{}
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl interface{}
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow interface{}
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations interface{}
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress interface{}
	EnableFilteringMessagesBeforePublishing interface{}
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning interface{}
	// Specifies the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location interface{}
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes interface{}
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name interface{}
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName interface{}
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection interface{}
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status interface{}
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering interface{}
}
