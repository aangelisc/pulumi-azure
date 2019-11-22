// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Stream Analytics Output to a ServiceBus Topic.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_topic.html.markdown.
type OutputServicebusTopic struct {
	s *pulumi.ResourceState
}

// NewOutputServicebusTopic registers a new resource with the given unique name, arguments, and options.
func NewOutputServicebusTopic(ctx *pulumi.Context,
	name string, args *OutputServicebusTopicArgs, opts ...pulumi.ResourceOpt) (*OutputServicebusTopic, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Serialization == nil {
		return nil, errors.New("missing required argument 'Serialization'")
	}
	if args == nil || args.ServicebusNamespace == nil {
		return nil, errors.New("missing required argument 'ServicebusNamespace'")
	}
	if args == nil || args.SharedAccessPolicyKey == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyKey'")
	}
	if args == nil || args.SharedAccessPolicyName == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyName'")
	}
	if args == nil || args.StreamAnalyticsJobName == nil {
		return nil, errors.New("missing required argument 'StreamAnalyticsJobName'")
	}
	if args == nil || args.TopicName == nil {
		return nil, errors.New("missing required argument 'TopicName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serialization"] = nil
		inputs["servicebusNamespace"] = nil
		inputs["sharedAccessPolicyKey"] = nil
		inputs["sharedAccessPolicyName"] = nil
		inputs["streamAnalyticsJobName"] = nil
		inputs["topicName"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serialization"] = args.Serialization
		inputs["servicebusNamespace"] = args.ServicebusNamespace
		inputs["sharedAccessPolicyKey"] = args.SharedAccessPolicyKey
		inputs["sharedAccessPolicyName"] = args.SharedAccessPolicyName
		inputs["streamAnalyticsJobName"] = args.StreamAnalyticsJobName
		inputs["topicName"] = args.TopicName
	}
	s, err := ctx.RegisterResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OutputServicebusTopic{s: s}, nil
}

// GetOutputServicebusTopic gets an existing OutputServicebusTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputServicebusTopic(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OutputServicebusTopicState, opts ...pulumi.ResourceOpt) (*OutputServicebusTopic, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serialization"] = state.Serialization
		inputs["servicebusNamespace"] = state.ServicebusNamespace
		inputs["sharedAccessPolicyKey"] = state.SharedAccessPolicyKey
		inputs["sharedAccessPolicyName"] = state.SharedAccessPolicyName
		inputs["streamAnalyticsJobName"] = state.StreamAnalyticsJobName
		inputs["topicName"] = state.TopicName
	}
	s, err := ctx.ReadResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OutputServicebusTopic{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OutputServicebusTopic) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OutputServicebusTopic) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the Stream Output. Changing this forces a new resource to be created.
func (r *OutputServicebusTopic) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
func (r *OutputServicebusTopic) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `serialization` block as defined below.
func (r *OutputServicebusTopic) Serialization() pulumi.Output {
	return r.s.State["serialization"]
}

// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
func (r *OutputServicebusTopic) ServicebusNamespace() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["servicebusNamespace"])
}

// The shared access policy key for the specified shared access policy.
func (r *OutputServicebusTopic) SharedAccessPolicyKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedAccessPolicyKey"])
}

// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
func (r *OutputServicebusTopic) SharedAccessPolicyName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedAccessPolicyName"])
}

// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
func (r *OutputServicebusTopic) StreamAnalyticsJobName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["streamAnalyticsJobName"])
}

// The name of the Service Bus Topic.
func (r *OutputServicebusTopic) TopicName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["topicName"])
}

// Input properties used for looking up and filtering OutputServicebusTopic resources.
type OutputServicebusTopicState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `serialization` block as defined below.
	Serialization interface{}
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace interface{}
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey interface{}
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName interface{}
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName interface{}
	// The name of the Service Bus Topic.
	TopicName interface{}
}

// The set of arguments for constructing a OutputServicebusTopic resource.
type OutputServicebusTopicArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `serialization` block as defined below.
	Serialization interface{}
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace interface{}
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey interface{}
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName interface{}
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName interface{}
	// The name of the Service Bus Topic.
	TopicName interface{}
}
