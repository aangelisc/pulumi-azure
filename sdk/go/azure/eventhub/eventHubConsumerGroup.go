// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub_consumer_group_legacy.html.markdown.
type EventHubConsumerGroup struct {
	s *pulumi.ResourceState
}

// NewEventHubConsumerGroup registers a new resource with the given unique name, arguments, and options.
func NewEventHubConsumerGroup(ctx *pulumi.Context,
	name string, args *EventHubConsumerGroupArgs, opts ...pulumi.ResourceOpt) (*EventHubConsumerGroup, error) {
	if args == nil || args.EventhubName == nil {
		return nil, errors.New("missing required argument 'EventhubName'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["eventhubName"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["namespaceName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["userMetadata"] = nil
	} else {
		inputs["eventhubName"] = args.EventhubName
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["namespaceName"] = args.NamespaceName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["userMetadata"] = args.UserMetadata
	}
	s, err := ctx.RegisterResource("azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventHubConsumerGroup{s: s}, nil
}

// GetEventHubConsumerGroup gets an existing EventHubConsumerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EventHubConsumerGroupState, opts ...pulumi.ResourceOpt) (*EventHubConsumerGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["eventhubName"] = state.EventhubName
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["namespaceName"] = state.NamespaceName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["userMetadata"] = state.UserMetadata
	}
	s, err := ctx.ReadResource("azure:eventhub/eventHubConsumerGroup:EventHubConsumerGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventHubConsumerGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EventHubConsumerGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EventHubConsumerGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the name of the EventHub. Changing this forces a new resource to be created.
func (r *EventHubConsumerGroup) EventhubName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventhubName"])
}

func (r *EventHubConsumerGroup) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
func (r *EventHubConsumerGroup) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
func (r *EventHubConsumerGroup) NamespaceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namespaceName"])
}

// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
func (r *EventHubConsumerGroup) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the user metadata.
func (r *EventHubConsumerGroup) UserMetadata() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userMetadata"])
}

// Input properties used for looking up and filtering EventHubConsumerGroup resources.
type EventHubConsumerGroupState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName interface{}
	Location interface{}
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName interface{}
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the user metadata.
	UserMetadata interface{}
}

// The set of arguments for constructing a EventHubConsumerGroup resource.
type EventHubConsumerGroupArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName interface{}
	Location interface{}
	// Specifies the name of the EventHub Consumer Group resource. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName interface{}
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the user metadata.
	UserMetadata interface{}
}
