// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Stream Analytics Stream Input EventHub.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_eventhub.html.markdown.
type StreamInputEventHub struct {
	s *pulumi.ResourceState
}

// NewStreamInputEventHub registers a new resource with the given unique name, arguments, and options.
func NewStreamInputEventHub(ctx *pulumi.Context,
	name string, args *StreamInputEventHubArgs, opts ...pulumi.ResourceOpt) (*StreamInputEventHub, error) {
	if args == nil || args.EventhubConsumerGroupName == nil {
		return nil, errors.New("missing required argument 'EventhubConsumerGroupName'")
	}
	if args == nil || args.EventhubName == nil {
		return nil, errors.New("missing required argument 'EventhubName'")
	}
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
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["eventhubConsumerGroupName"] = nil
		inputs["eventhubName"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serialization"] = nil
		inputs["servicebusNamespace"] = nil
		inputs["sharedAccessPolicyKey"] = nil
		inputs["sharedAccessPolicyName"] = nil
		inputs["streamAnalyticsJobName"] = nil
	} else {
		inputs["eventhubConsumerGroupName"] = args.EventhubConsumerGroupName
		inputs["eventhubName"] = args.EventhubName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serialization"] = args.Serialization
		inputs["servicebusNamespace"] = args.ServicebusNamespace
		inputs["sharedAccessPolicyKey"] = args.SharedAccessPolicyKey
		inputs["sharedAccessPolicyName"] = args.SharedAccessPolicyName
		inputs["streamAnalyticsJobName"] = args.StreamAnalyticsJobName
	}
	s, err := ctx.RegisterResource("azure:streamanalytics/streamInputEventHub:StreamInputEventHub", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StreamInputEventHub{s: s}, nil
}

// GetStreamInputEventHub gets an existing StreamInputEventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamInputEventHub(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StreamInputEventHubState, opts ...pulumi.ResourceOpt) (*StreamInputEventHub, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["eventhubConsumerGroupName"] = state.EventhubConsumerGroupName
		inputs["eventhubName"] = state.EventhubName
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serialization"] = state.Serialization
		inputs["servicebusNamespace"] = state.ServicebusNamespace
		inputs["sharedAccessPolicyKey"] = state.SharedAccessPolicyKey
		inputs["sharedAccessPolicyName"] = state.SharedAccessPolicyName
		inputs["streamAnalyticsJobName"] = state.StreamAnalyticsJobName
	}
	s, err := ctx.ReadResource("azure:streamanalytics/streamInputEventHub:StreamInputEventHub", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StreamInputEventHub{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *StreamInputEventHub) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *StreamInputEventHub) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
func (r *StreamInputEventHub) EventhubConsumerGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventhubConsumerGroupName"])
}

// The name of the Event Hub.
func (r *StreamInputEventHub) EventhubName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventhubName"])
}

// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
func (r *StreamInputEventHub) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
func (r *StreamInputEventHub) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `serialization` block as defined below.
func (r *StreamInputEventHub) Serialization() pulumi.Output {
	return r.s.State["serialization"]
}

// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
func (r *StreamInputEventHub) ServicebusNamespace() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["servicebusNamespace"])
}

// The shared access policy key for the specified shared access policy.
func (r *StreamInputEventHub) SharedAccessPolicyKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedAccessPolicyKey"])
}

// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
func (r *StreamInputEventHub) SharedAccessPolicyName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedAccessPolicyName"])
}

// The name of the Stream Analytics Job. Changing this forces a new resource to be created. 
func (r *StreamInputEventHub) StreamAnalyticsJobName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["streamAnalyticsJobName"])
}

// Input properties used for looking up and filtering StreamInputEventHub resources.
type StreamInputEventHubState struct {
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName interface{}
	// The name of the Event Hub.
	EventhubName interface{}
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `serialization` block as defined below.
	Serialization interface{}
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace interface{}
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey interface{}
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName interface{}
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created. 
	StreamAnalyticsJobName interface{}
}

// The set of arguments for constructing a StreamInputEventHub resource.
type StreamInputEventHubArgs struct {
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName interface{}
	// The name of the Event Hub.
	EventhubName interface{}
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `serialization` block as defined below.
	Serialization interface{}
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace interface{}
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey interface{}
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName interface{}
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created. 
	StreamAnalyticsJobName interface{}
}
