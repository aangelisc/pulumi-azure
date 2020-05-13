# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("azure.eventhub.EventHubAuthorizationRule has been deprecated in favor of azure.eventhub.AuthorizationRule", DeprecationWarning)
class EventHubAuthorizationRule(pulumi.CustomResource):
    eventhub_name: pulumi.Output[str]
    """
    Specifies the name of the EventHub. Changing this forces a new resource to be created.
    """
    listen: pulumi.Output[bool]
    """
    Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
    """
    manage: pulumi.Output[bool]
    """
    Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
    """
    namespace_name: pulumi.Output[str]
    """
    Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
    """
    primary_connection_string: pulumi.Output[str]
    """
    The Primary Connection String for the Event Hubs authorization Rule.
    """
    primary_connection_string_alias: pulumi.Output[str]
    """
    The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
    """
    primary_key: pulumi.Output[str]
    """
    The Primary Key for the Event Hubs authorization Rule.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
    """
    secondary_connection_string: pulumi.Output[str]
    """
    The Secondary Connection String for the Event Hubs Authorization Rule.
    """
    secondary_connection_string_alias: pulumi.Output[str]
    """
    The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
    """
    secondary_key: pulumi.Output[str]
    """
    The Secondary Key for the Event Hubs Authorization Rule.
    """
    send: pulumi.Output[bool]
    """
    Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
    """
    warnings.warn("azure.eventhub.EventHubAuthorizationRule has been deprecated in favor of azure.eventhub.AuthorizationRule", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, eventhub_name=None, listen=None, manage=None, name=None, namespace_name=None, resource_group_name=None, send=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Event Hubs authorization Rule within an Event Hub.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_event_hub_namespace = azure.eventhub.EventHubNamespace("exampleEventHubNamespace",
            location="West US",
            resource_group_name=example_resource_group.name,
            sku="Basic",
            capacity=2,
            tags={
                "environment": "Production",
            })
        example_event_hub = azure.eventhub.EventHub("exampleEventHub",
            namespace_name=example_event_hub_namespace.name,
            resource_group_name=example_resource_group.name,
            partition_count=2,
            message_retention=2)
        example_authorization_rule = azure.eventhub.AuthorizationRule("exampleAuthorizationRule",
            namespace_name=example_event_hub_namespace.name,
            eventhub_name=example_event_hub.name,
            resource_group_name=example_resource_group.name,
            listen=True,
            send=False,
            manage=False)
        ```


        Deprecated: azure.eventhub.EventHubAuthorizationRule has been deprecated in favor of azure.eventhub.AuthorizationRule

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] send: Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        pulumi.log.warn("EventHubAuthorizationRule is deprecated: azure.eventhub.EventHubAuthorizationRule has been deprecated in favor of azure.eventhub.AuthorizationRule")
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if eventhub_name is None:
                raise TypeError("Missing required property 'eventhub_name'")
            __props__['eventhub_name'] = eventhub_name
            __props__['listen'] = listen
            __props__['manage'] = manage
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['send'] = send
            __props__['primary_connection_string'] = None
            __props__['primary_connection_string_alias'] = None
            __props__['primary_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['secondary_connection_string_alias'] = None
            __props__['secondary_key'] = None
        super(EventHubAuthorizationRule, __self__).__init__(
            'azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, eventhub_name=None, listen=None, manage=None, name=None, namespace_name=None, primary_connection_string=None, primary_connection_string_alias=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_connection_string_alias=None, secondary_key=None, send=None):
        """
        Get an existing EventHubAuthorizationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the EventHub. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The Primary Connection String for the Event Hubs authorization Rule.
        :param pulumi.Input[str] primary_connection_string_alias: The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] primary_key: The Primary Key for the Event Hubs authorization Rule.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string: The Secondary Connection String for the Event Hubs Authorization Rule.
        :param pulumi.Input[str] secondary_connection_string_alias: The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
        :param pulumi.Input[str] secondary_key: The Secondary Key for the Event Hubs Authorization Rule.
        :param pulumi.Input[bool] send: Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["eventhub_name"] = eventhub_name
        __props__["listen"] = listen
        __props__["manage"] = manage
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["primary_connection_string"] = primary_connection_string
        __props__["primary_connection_string_alias"] = primary_connection_string_alias
        __props__["primary_key"] = primary_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["secondary_connection_string_alias"] = secondary_connection_string_alias
        __props__["secondary_key"] = secondary_key
        __props__["send"] = send
        return EventHubAuthorizationRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

