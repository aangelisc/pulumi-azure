// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventgrid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EventGrid Event Subscription
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventgrid_event_subscription.html.markdown.
type EventSubscription struct {
	pulumi.CustomResourceState

	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
	EventDeliverySchema pulumi.StringPtrOutput `pulumi:"eventDeliverySchema"`
	// A `eventhubEndpoint` block as defined below.
	EventhubEndpoint EventSubscriptionEventhubEndpointPtrOutput `pulumi:"eventhubEndpoint"`
	// A `hybridConnectionEndpoint` block as defined below.
	HybridConnectionEndpoint EventSubscriptionHybridConnectionEndpointPtrOutput `pulumi:"hybridConnectionEndpoint"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayOutput `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `retryPolicy` block as defined below.
	RetryPolicy EventSubscriptionRetryPolicyOutput `pulumi:"retryPolicy"`
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination EventSubscriptionStorageBlobDeadLetterDestinationPtrOutput `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint EventSubscriptionStorageQueueEndpointPtrOutput `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter EventSubscriptionSubjectFilterPtrOutput `pulumi:"subjectFilter"`
	// Specifies the name of the topic to associate with the event subscription.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint EventSubscriptionWebhookEndpointPtrOutput `pulumi:"webhookEndpoint"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &EventSubscriptionArgs{}
	}
	var resource EventSubscription
	err := ctx.RegisterResource("azure:eventgrid/eventSubscription:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("azure:eventgrid/eventSubscription:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// A `eventhubEndpoint` block as defined below.
	EventhubEndpoint *EventSubscriptionEventhubEndpoint `pulumi:"eventhubEndpoint"`
	// A `hybridConnectionEndpoint` block as defined below.
	HybridConnectionEndpoint *EventSubscriptionHybridConnectionEndpoint `pulumi:"hybridConnectionEndpoint"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels []string `pulumi:"labels"`
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `retryPolicy` block as defined below.
	RetryPolicy *EventSubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination *EventSubscriptionStorageBlobDeadLetterDestination `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint *EventSubscriptionStorageQueueEndpoint `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter *EventSubscriptionSubjectFilter `pulumi:"subjectFilter"`
	// Specifies the name of the topic to associate with the event subscription.
	TopicName *string `pulumi:"topicName"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint *EventSubscriptionWebhookEndpoint `pulumi:"webhookEndpoint"`
}

type EventSubscriptionState struct {
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
	EventDeliverySchema pulumi.StringPtrInput
	// A `eventhubEndpoint` block as defined below.
	EventhubEndpoint EventSubscriptionEventhubEndpointPtrInput
	// A `hybridConnectionEndpoint` block as defined below.
	HybridConnectionEndpoint EventSubscriptionHybridConnectionEndpointPtrInput
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayInput
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayInput
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `retryPolicy` block as defined below.
	RetryPolicy EventSubscriptionRetryPolicyPtrInput
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination EventSubscriptionStorageBlobDeadLetterDestinationPtrInput
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint EventSubscriptionStorageQueueEndpointPtrInput
	// A `subjectFilter` block as defined below.
	SubjectFilter EventSubscriptionSubjectFilterPtrInput
	// Specifies the name of the topic to associate with the event subscription.
	TopicName pulumi.StringPtrInput
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint EventSubscriptionWebhookEndpointPtrInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// A `eventhubEndpoint` block as defined below.
	EventhubEndpoint *EventSubscriptionEventhubEndpoint `pulumi:"eventhubEndpoint"`
	// A `hybridConnectionEndpoint` block as defined below.
	HybridConnectionEndpoint *EventSubscriptionHybridConnectionEndpoint `pulumi:"hybridConnectionEndpoint"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels []string `pulumi:"labels"`
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `retryPolicy` block as defined below.
	RetryPolicy *EventSubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination *EventSubscriptionStorageBlobDeadLetterDestination `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint *EventSubscriptionStorageQueueEndpoint `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter *EventSubscriptionSubjectFilter `pulumi:"subjectFilter"`
	// Specifies the name of the topic to associate with the event subscription.
	TopicName *string `pulumi:"topicName"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint *EventSubscriptionWebhookEndpoint `pulumi:"webhookEndpoint"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
	EventDeliverySchema pulumi.StringPtrInput
	// A `eventhubEndpoint` block as defined below.
	EventhubEndpoint EventSubscriptionEventhubEndpointPtrInput
	// A `hybridConnectionEndpoint` block as defined below.
	HybridConnectionEndpoint EventSubscriptionHybridConnectionEndpointPtrInput
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayInput
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayInput
	// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `retryPolicy` block as defined below.
	RetryPolicy EventSubscriptionRetryPolicyPtrInput
	// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination EventSubscriptionStorageBlobDeadLetterDestinationPtrInput
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint EventSubscriptionStorageQueueEndpointPtrInput
	// A `subjectFilter` block as defined below.
	SubjectFilter EventSubscriptionSubjectFilterPtrInput
	// Specifies the name of the topic to associate with the event subscription.
	TopicName pulumi.StringPtrInput
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint EventSubscriptionWebhookEndpointPtrInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

