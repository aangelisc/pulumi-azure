# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SystemTopicEventSubscription']


class SystemTopicEventSubscription(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_filter: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAdvancedFilterArgs']]] = None,
                 azure_function_endpoint: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAzureFunctionEndpointArgs']]] = None,
                 event_delivery_schema: Optional[pulumi.Input[str]] = None,
                 eventhub_endpoint_id: Optional[pulumi.Input[str]] = None,
                 expiration_time_utc: Optional[pulumi.Input[str]] = None,
                 hybrid_connection_endpoint_id: Optional[pulumi.Input[str]] = None,
                 included_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retry_policy: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionRetryPolicyArgs']]] = None,
                 service_bus_queue_endpoint_id: Optional[pulumi.Input[str]] = None,
                 service_bus_topic_endpoint_id: Optional[pulumi.Input[str]] = None,
                 storage_blob_dead_letter_destination: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationArgs']]] = None,
                 storage_queue_endpoint: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageQueueEndpointArgs']]] = None,
                 subject_filter: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionSubjectFilterArgs']]] = None,
                 system_topic: Optional[pulumi.Input[str]] = None,
                 webhook_endpoint: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionWebhookEndpointArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an EventGrid System Topic Event Subscription.

        ## Import

        EventGrid System Topic Event Subscriptions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/systemTopics/topic1/eventSubscriptions/subscription1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAdvancedFilterArgs']] advanced_filter: A `advanced_filter` block as defined below.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAzureFunctionEndpointArgs']] azure_function_endpoint: An `azure_function_endpoint` block as defined below.
        :param pulumi.Input[str] event_delivery_schema: Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eventhub_endpoint_id: Specifies the id where the Event Hub is located.
        :param pulumi.Input[str] expiration_time_utc: Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        :param pulumi.Input[str] hybrid_connection_endpoint_id: Specifies the id where the Hybrid Connection is located.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_event_types: A list of applicable event types that need to be part of the event subscription.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: A list of labels to assign to the event subscription.
        :param pulumi.Input[str] name: The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionRetryPolicyArgs']] retry_policy: A `retry_policy` block as defined below.
        :param pulumi.Input[str] service_bus_queue_endpoint_id: Specifies the id where the Service Bus Queue is located.
        :param pulumi.Input[str] service_bus_topic_endpoint_id: Specifies the id where the Service Bus Topic is located.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationArgs']] storage_blob_dead_letter_destination: A `storage_blob_dead_letter_destination` block as defined below.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageQueueEndpointArgs']] storage_queue_endpoint: A `storage_queue_endpoint` block as defined below.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionSubjectFilterArgs']] subject_filter: A `subject_filter` block as defined below.
        :param pulumi.Input[str] system_topic: The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionWebhookEndpointArgs']] webhook_endpoint: A `webhook_endpoint` block as defined below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['advanced_filter'] = advanced_filter
            __props__['azure_function_endpoint'] = azure_function_endpoint
            __props__['event_delivery_schema'] = event_delivery_schema
            __props__['eventhub_endpoint_id'] = eventhub_endpoint_id
            __props__['expiration_time_utc'] = expiration_time_utc
            __props__['hybrid_connection_endpoint_id'] = hybrid_connection_endpoint_id
            __props__['included_event_types'] = included_event_types
            __props__['labels'] = labels
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retry_policy'] = retry_policy
            __props__['service_bus_queue_endpoint_id'] = service_bus_queue_endpoint_id
            __props__['service_bus_topic_endpoint_id'] = service_bus_topic_endpoint_id
            __props__['storage_blob_dead_letter_destination'] = storage_blob_dead_letter_destination
            __props__['storage_queue_endpoint'] = storage_queue_endpoint
            __props__['subject_filter'] = subject_filter
            if system_topic is None and not opts.urn:
                raise TypeError("Missing required property 'system_topic'")
            __props__['system_topic'] = system_topic
            __props__['webhook_endpoint'] = webhook_endpoint
        super(SystemTopicEventSubscription, __self__).__init__(
            'azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            advanced_filter: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAdvancedFilterArgs']]] = None,
            azure_function_endpoint: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAzureFunctionEndpointArgs']]] = None,
            event_delivery_schema: Optional[pulumi.Input[str]] = None,
            eventhub_endpoint_id: Optional[pulumi.Input[str]] = None,
            expiration_time_utc: Optional[pulumi.Input[str]] = None,
            hybrid_connection_endpoint_id: Optional[pulumi.Input[str]] = None,
            included_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            retry_policy: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionRetryPolicyArgs']]] = None,
            service_bus_queue_endpoint_id: Optional[pulumi.Input[str]] = None,
            service_bus_topic_endpoint_id: Optional[pulumi.Input[str]] = None,
            storage_blob_dead_letter_destination: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationArgs']]] = None,
            storage_queue_endpoint: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageQueueEndpointArgs']]] = None,
            subject_filter: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionSubjectFilterArgs']]] = None,
            system_topic: Optional[pulumi.Input[str]] = None,
            webhook_endpoint: Optional[pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionWebhookEndpointArgs']]] = None) -> 'SystemTopicEventSubscription':
        """
        Get an existing SystemTopicEventSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAdvancedFilterArgs']] advanced_filter: A `advanced_filter` block as defined below.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionAzureFunctionEndpointArgs']] azure_function_endpoint: An `azure_function_endpoint` block as defined below.
        :param pulumi.Input[str] event_delivery_schema: Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eventhub_endpoint_id: Specifies the id where the Event Hub is located.
        :param pulumi.Input[str] expiration_time_utc: Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        :param pulumi.Input[str] hybrid_connection_endpoint_id: Specifies the id where the Hybrid Connection is located.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_event_types: A list of applicable event types that need to be part of the event subscription.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: A list of labels to assign to the event subscription.
        :param pulumi.Input[str] name: The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionRetryPolicyArgs']] retry_policy: A `retry_policy` block as defined below.
        :param pulumi.Input[str] service_bus_queue_endpoint_id: Specifies the id where the Service Bus Queue is located.
        :param pulumi.Input[str] service_bus_topic_endpoint_id: Specifies the id where the Service Bus Topic is located.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationArgs']] storage_blob_dead_letter_destination: A `storage_blob_dead_letter_destination` block as defined below.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionStorageQueueEndpointArgs']] storage_queue_endpoint: A `storage_queue_endpoint` block as defined below.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionSubjectFilterArgs']] subject_filter: A `subject_filter` block as defined below.
        :param pulumi.Input[str] system_topic: The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
        :param pulumi.Input[pulumi.InputType['SystemTopicEventSubscriptionWebhookEndpointArgs']] webhook_endpoint: A `webhook_endpoint` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["advanced_filter"] = advanced_filter
        __props__["azure_function_endpoint"] = azure_function_endpoint
        __props__["event_delivery_schema"] = event_delivery_schema
        __props__["eventhub_endpoint_id"] = eventhub_endpoint_id
        __props__["expiration_time_utc"] = expiration_time_utc
        __props__["hybrid_connection_endpoint_id"] = hybrid_connection_endpoint_id
        __props__["included_event_types"] = included_event_types
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["retry_policy"] = retry_policy
        __props__["service_bus_queue_endpoint_id"] = service_bus_queue_endpoint_id
        __props__["service_bus_topic_endpoint_id"] = service_bus_topic_endpoint_id
        __props__["storage_blob_dead_letter_destination"] = storage_blob_dead_letter_destination
        __props__["storage_queue_endpoint"] = storage_queue_endpoint
        __props__["subject_filter"] = subject_filter
        __props__["system_topic"] = system_topic
        __props__["webhook_endpoint"] = webhook_endpoint
        return SystemTopicEventSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="advancedFilter")
    def advanced_filter(self) -> pulumi.Output[Optional['outputs.SystemTopicEventSubscriptionAdvancedFilter']]:
        """
        A `advanced_filter` block as defined below.
        """
        return pulumi.get(self, "advanced_filter")

    @property
    @pulumi.getter(name="azureFunctionEndpoint")
    def azure_function_endpoint(self) -> pulumi.Output[Optional['outputs.SystemTopicEventSubscriptionAzureFunctionEndpoint']]:
        """
        An `azure_function_endpoint` block as defined below.
        """
        return pulumi.get(self, "azure_function_endpoint")

    @property
    @pulumi.getter(name="eventDeliverySchema")
    def event_delivery_schema(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "event_delivery_schema")

    @property
    @pulumi.getter(name="eventhubEndpointId")
    def eventhub_endpoint_id(self) -> pulumi.Output[str]:
        """
        Specifies the id where the Event Hub is located.
        """
        return pulumi.get(self, "eventhub_endpoint_id")

    @property
    @pulumi.getter(name="expirationTimeUtc")
    def expiration_time_utc(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
        """
        return pulumi.get(self, "expiration_time_utc")

    @property
    @pulumi.getter(name="hybridConnectionEndpointId")
    def hybrid_connection_endpoint_id(self) -> pulumi.Output[str]:
        """
        Specifies the id where the Hybrid Connection is located.
        """
        return pulumi.get(self, "hybrid_connection_endpoint_id")

    @property
    @pulumi.getter(name="includedEventTypes")
    def included_event_types(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of applicable event types that need to be part of the event subscription.
        """
        return pulumi.get(self, "included_event_types")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of labels to assign to the event subscription.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="retryPolicy")
    def retry_policy(self) -> pulumi.Output['outputs.SystemTopicEventSubscriptionRetryPolicy']:
        """
        A `retry_policy` block as defined below.
        """
        return pulumi.get(self, "retry_policy")

    @property
    @pulumi.getter(name="serviceBusQueueEndpointId")
    def service_bus_queue_endpoint_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the id where the Service Bus Queue is located.
        """
        return pulumi.get(self, "service_bus_queue_endpoint_id")

    @property
    @pulumi.getter(name="serviceBusTopicEndpointId")
    def service_bus_topic_endpoint_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the id where the Service Bus Topic is located.
        """
        return pulumi.get(self, "service_bus_topic_endpoint_id")

    @property
    @pulumi.getter(name="storageBlobDeadLetterDestination")
    def storage_blob_dead_letter_destination(self) -> pulumi.Output[Optional['outputs.SystemTopicEventSubscriptionStorageBlobDeadLetterDestination']]:
        """
        A `storage_blob_dead_letter_destination` block as defined below.
        """
        return pulumi.get(self, "storage_blob_dead_letter_destination")

    @property
    @pulumi.getter(name="storageQueueEndpoint")
    def storage_queue_endpoint(self) -> pulumi.Output[Optional['outputs.SystemTopicEventSubscriptionStorageQueueEndpoint']]:
        """
        A `storage_queue_endpoint` block as defined below.
        """
        return pulumi.get(self, "storage_queue_endpoint")

    @property
    @pulumi.getter(name="subjectFilter")
    def subject_filter(self) -> pulumi.Output[Optional['outputs.SystemTopicEventSubscriptionSubjectFilter']]:
        """
        A `subject_filter` block as defined below.
        """
        return pulumi.get(self, "subject_filter")

    @property
    @pulumi.getter(name="systemTopic")
    def system_topic(self) -> pulumi.Output[str]:
        """
        The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
        """
        return pulumi.get(self, "system_topic")

    @property
    @pulumi.getter(name="webhookEndpoint")
    def webhook_endpoint(self) -> pulumi.Output[Optional['outputs.SystemTopicEventSubscriptionWebhookEndpoint']]:
        """
        A `webhook_endpoint` block as defined below.
        """
        return pulumi.get(self, "webhook_endpoint")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
