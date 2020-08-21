# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['IoTHub']


class IoTHub(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubEndpointArgs']]]]] = None,
                 event_hub_partition_count: Optional[pulumi.Input[float]] = None,
                 event_hub_retention_in_days: Optional[pulumi.Input[float]] = None,
                 fallback_route: Optional[pulumi.Input[pulumi.InputType['IoTHubFallbackRouteArgs']]] = None,
                 file_upload: Optional[pulumi.Input[pulumi.InputType['IoTHubFileUploadArgs']]] = None,
                 ip_filter_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubIpFilterRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 routes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubRouteArgs']]]]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['IoTHubSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an IotHub

        > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.

        > **NOTE:** Routes can be defined either directly on the `iot.IoTHub` resource, or using the `iot.Route` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.

        > **NOTE:** Fallback route can be defined either directly on the `iot.IoTHub` resource, or using the `iot.FallbackRoute` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="Canada Central")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku="Basic")
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            resource_group_name=example_resource_group.name,
            namespace_name=example_event_hub_namespace.name,
            partition_count=2,
            message_retention=1)
        example_authorization_rule = azure.eventhub.AuthorizationRule("exampleAuthorizationRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            send=True)
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.iot.IoTHubSkuArgs(
                name="S1",
                capacity=1,
            ),
            endpoints=[
                azure.iot.IoTHubEndpointArgs(
                    type="AzureIotHub.StorageContainer",
                    connection_string=example_account.primary_blob_connection_string,
                    name="export",
                    batch_frequency_in_seconds=60,
                    max_chunk_size_in_bytes=10485760,
                    container_name=example_container.name,
                    encoding="Avro",
                    file_name_format="{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}",
                ),
                azure.iot.IoTHubEndpointArgs(
                    type="AzureIotHub.EventHub",
                    connection_string=example_authorization_rule.primary_connection_string,
                    name="export2",
                ),
            ],
            routes=[
                azure.iot.IoTHubRouteArgs(
                    name="export",
                    source="DeviceMessages",
                    condition="true",
                    endpoint_names=["export"],
                    enabled=True,
                ),
                azure.iot.IoTHubRouteArgs(
                    name="export2",
                    source="DeviceMessages",
                    condition="true",
                    endpoint_names=["export2"],
                    enabled=True,
                ),
            ],
            tags={
                "purpose": "testing",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubEndpointArgs']]]] endpoints: An `endpoint` block as defined below.
        :param pulumi.Input[float] event_hub_partition_count: The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
        :param pulumi.Input[float] event_hub_retention_in_days: The event hub retention to use in days. Must be between `1` and `7`.
        :param pulumi.Input[pulumi.InputType['IoTHubFallbackRouteArgs']] fallback_route: A `fallback_route` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
        :param pulumi.Input[pulumi.InputType['IoTHubFileUploadArgs']] file_upload: A `file_upload` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubIpFilterRuleArgs']]]] ip_filter_rules: One or more `ip_filter_rule` blocks as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubRouteArgs']]]] routes: A `route` block as defined below.
        :param pulumi.Input[pulumi.InputType['IoTHubSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['endpoints'] = endpoints
            __props__['event_hub_partition_count'] = event_hub_partition_count
            __props__['event_hub_retention_in_days'] = event_hub_retention_in_days
            __props__['fallback_route'] = fallback_route
            __props__['file_upload'] = file_upload
            __props__['ip_filter_rules'] = ip_filter_rules
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['routes'] = routes
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['event_hub_events_endpoint'] = None
            __props__['event_hub_events_path'] = None
            __props__['event_hub_operations_endpoint'] = None
            __props__['event_hub_operations_path'] = None
            __props__['hostname'] = None
            __props__['shared_access_policies'] = None
            __props__['type'] = None
        super(IoTHub, __self__).__init__(
            'azure:iot/ioTHub:IoTHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubEndpointArgs']]]]] = None,
            event_hub_events_endpoint: Optional[pulumi.Input[str]] = None,
            event_hub_events_path: Optional[pulumi.Input[str]] = None,
            event_hub_operations_endpoint: Optional[pulumi.Input[str]] = None,
            event_hub_operations_path: Optional[pulumi.Input[str]] = None,
            event_hub_partition_count: Optional[pulumi.Input[float]] = None,
            event_hub_retention_in_days: Optional[pulumi.Input[float]] = None,
            fallback_route: Optional[pulumi.Input[pulumi.InputType['IoTHubFallbackRouteArgs']]] = None,
            file_upload: Optional[pulumi.Input[pulumi.InputType['IoTHubFileUploadArgs']]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ip_filter_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubIpFilterRuleArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            routes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubRouteArgs']]]]] = None,
            shared_access_policies: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubSharedAccessPolicyArgs']]]]] = None,
            sku: Optional[pulumi.Input[pulumi.InputType['IoTHubSkuArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'IoTHub':
        """
        Get an existing IoTHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubEndpointArgs']]]] endpoints: An `endpoint` block as defined below.
        :param pulumi.Input[str] event_hub_events_endpoint: The EventHub compatible endpoint for events data
        :param pulumi.Input[str] event_hub_events_path: The EventHub compatible path for events data
        :param pulumi.Input[str] event_hub_operations_endpoint: The EventHub compatible endpoint for operational data
        :param pulumi.Input[str] event_hub_operations_path: The EventHub compatible path for operational data
        :param pulumi.Input[float] event_hub_partition_count: The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
        :param pulumi.Input[float] event_hub_retention_in_days: The event hub retention to use in days. Must be between `1` and `7`.
        :param pulumi.Input[pulumi.InputType['IoTHubFallbackRouteArgs']] fallback_route: A `fallback_route` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
        :param pulumi.Input[pulumi.InputType['IoTHubFileUploadArgs']] file_upload: A `file_upload` block as defined below.
        :param pulumi.Input[str] hostname: The hostname of the IotHub Resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubIpFilterRuleArgs']]]] ip_filter_rules: One or more `ip_filter_rule` blocks as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubRouteArgs']]]] routes: A `route` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IoTHubSharedAccessPolicyArgs']]]] shared_access_policies: One or more `shared_access_policy` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['IoTHubSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["endpoints"] = endpoints
        __props__["event_hub_events_endpoint"] = event_hub_events_endpoint
        __props__["event_hub_events_path"] = event_hub_events_path
        __props__["event_hub_operations_endpoint"] = event_hub_operations_endpoint
        __props__["event_hub_operations_path"] = event_hub_operations_path
        __props__["event_hub_partition_count"] = event_hub_partition_count
        __props__["event_hub_retention_in_days"] = event_hub_retention_in_days
        __props__["fallback_route"] = fallback_route
        __props__["file_upload"] = file_upload
        __props__["hostname"] = hostname
        __props__["ip_filter_rules"] = ip_filter_rules
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["routes"] = routes
        __props__["shared_access_policies"] = shared_access_policies
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["type"] = type
        return IoTHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoints(self) -> List['outputs.IoTHubEndpoint']:
        """
        An `endpoint` block as defined below.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="eventHubEventsEndpoint")
    def event_hub_events_endpoint(self) -> str:
        """
        The EventHub compatible endpoint for events data
        """
        return pulumi.get(self, "event_hub_events_endpoint")

    @property
    @pulumi.getter(name="eventHubEventsPath")
    def event_hub_events_path(self) -> str:
        """
        The EventHub compatible path for events data
        """
        return pulumi.get(self, "event_hub_events_path")

    @property
    @pulumi.getter(name="eventHubOperationsEndpoint")
    def event_hub_operations_endpoint(self) -> str:
        """
        The EventHub compatible endpoint for operational data
        """
        return pulumi.get(self, "event_hub_operations_endpoint")

    @property
    @pulumi.getter(name="eventHubOperationsPath")
    def event_hub_operations_path(self) -> str:
        """
        The EventHub compatible path for operational data
        """
        return pulumi.get(self, "event_hub_operations_path")

    @property
    @pulumi.getter(name="eventHubPartitionCount")
    def event_hub_partition_count(self) -> float:
        """
        The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
        """
        return pulumi.get(self, "event_hub_partition_count")

    @property
    @pulumi.getter(name="eventHubRetentionInDays")
    def event_hub_retention_in_days(self) -> float:
        """
        The event hub retention to use in days. Must be between `1` and `7`.
        """
        return pulumi.get(self, "event_hub_retention_in_days")

    @property
    @pulumi.getter(name="fallbackRoute")
    def fallback_route(self) -> 'outputs.IoTHubFallbackRoute':
        """
        A `fallback_route` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
        """
        return pulumi.get(self, "fallback_route")

    @property
    @pulumi.getter(name="fileUpload")
    def file_upload(self) -> Optional['outputs.IoTHubFileUpload']:
        """
        A `file_upload` block as defined below.
        """
        return pulumi.get(self, "file_upload")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        The hostname of the IotHub Resource.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="ipFilterRules")
    def ip_filter_rules(self) -> Optional[List['outputs.IoTHubIpFilterRule']]:
        """
        One or more `ip_filter_rule` blocks as defined below.
        """
        return pulumi.get(self, "ip_filter_rules")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def routes(self) -> List['outputs.IoTHubRoute']:
        """
        A `route` block as defined below.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter(name="sharedAccessPolicies")
    def shared_access_policies(self) -> List['outputs.IoTHubSharedAccessPolicy']:
        """
        One or more `shared_access_policy` blocks as defined below.
        """
        return pulumi.get(self, "shared_access_policies")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.IoTHubSku':
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

