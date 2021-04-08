# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SubscriptionRuleArgs', 'SubscriptionRule']

@pulumi.input_type
class SubscriptionRuleArgs:
    def __init__(__self__, *,
                 filter_type: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 subscription_name: pulumi.Input[str],
                 topic_name: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 correlation_filter: Optional[pulumi.Input['SubscriptionRuleCorrelationFilterArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sql_filter: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubscriptionRule resource.
        :param pulumi.Input[str] filter_type: Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subscription_name: The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] action: Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        :param pulumi.Input['SubscriptionRuleCorrelationFilterArgs'] correlation_filter: A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_filter: Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        """
        pulumi.set(__self__, "filter_type", filter_type)
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "subscription_name", subscription_name)
        pulumi.set(__self__, "topic_name", topic_name)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if correlation_filter is not None:
            pulumi.set(__self__, "correlation_filter", correlation_filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sql_filter is not None:
            pulumi.set(__self__, "sql_filter", sql_filter)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> pulumi.Input[str]:
        """
        Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Input[str]:
        """
        The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="correlationFilter")
    def correlation_filter(self) -> Optional[pulumi.Input['SubscriptionRuleCorrelationFilterArgs']]:
        """
        A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        """
        return pulumi.get(self, "correlation_filter")

    @correlation_filter.setter
    def correlation_filter(self, value: Optional[pulumi.Input['SubscriptionRuleCorrelationFilterArgs']]):
        pulumi.set(self, "correlation_filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sqlFilter")
    def sql_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        """
        return pulumi.get(self, "sql_filter")

    @sql_filter.setter
    def sql_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_filter", value)


@pulumi.input_type
class _SubscriptionRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 correlation_filter: Optional[pulumi.Input['SubscriptionRuleCorrelationFilterArgs']] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sql_filter: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SubscriptionRule resources.
        :param pulumi.Input[str] action: Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        :param pulumi.Input['SubscriptionRuleCorrelationFilterArgs'] correlation_filter: A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        :param pulumi.Input[str] filter_type: Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_filter: Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        :param pulumi.Input[str] subscription_name: The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if correlation_filter is not None:
            pulumi.set(__self__, "correlation_filter", correlation_filter)
        if filter_type is not None:
            pulumi.set(__self__, "filter_type", filter_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if sql_filter is not None:
            pulumi.set(__self__, "sql_filter", sql_filter)
        if subscription_name is not None:
            pulumi.set(__self__, "subscription_name", subscription_name)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="correlationFilter")
    def correlation_filter(self) -> Optional[pulumi.Input['SubscriptionRuleCorrelationFilterArgs']]:
        """
        A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        """
        return pulumi.get(self, "correlation_filter")

    @correlation_filter.setter
    def correlation_filter(self, value: Optional[pulumi.Input['SubscriptionRuleCorrelationFilterArgs']]):
        pulumi.set(self, "correlation_filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="sqlFilter")
    def sql_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        """
        return pulumi.get(self, "sql_filter")

    @sql_filter.setter
    def sql_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_filter", value)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)


class SubscriptionRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 correlation_filter: Optional[pulumi.Input[pulumi.InputType['SubscriptionRuleCorrelationFilterArgs']]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sql_filter: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a ServiceBus Subscription Rule.

        ## Example Usage
        ### SQL Filter)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_subscription = azure.servicebus.Subscription("exampleSubscription",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            max_delivery_count=1)
        example_subscription_rule = azure.servicebus.SubscriptionRule("exampleSubscriptionRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            subscription_name=example_subscription.name,
            filter_type="SqlFilter",
            sql_filter="colour = 'red'")
        ```
        ### Correlation Filter)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_subscription = azure.servicebus.Subscription("exampleSubscription",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            max_delivery_count=1)
        example_subscription_rule = azure.servicebus.SubscriptionRule("exampleSubscriptionRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            subscription_name=example_subscription.name,
            filter_type="CorrelationFilter",
            correlation_filter=azure.servicebus.SubscriptionRuleCorrelationFilterArgs(
                correlation_id="high",
                label="red",
                properties={
                    "customProperty": "value",
                },
            ))
        ```

        ## Import

        Service Bus Subscription Rule can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:servicebus/subscriptionRule:SubscriptionRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        :param pulumi.Input[pulumi.InputType['SubscriptionRuleCorrelationFilterArgs']] correlation_filter: A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        :param pulumi.Input[str] filter_type: Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_filter: Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        :param pulumi.Input[str] subscription_name: The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubscriptionRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a ServiceBus Subscription Rule.

        ## Example Usage
        ### SQL Filter)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_subscription = azure.servicebus.Subscription("exampleSubscription",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            max_delivery_count=1)
        example_subscription_rule = azure.servicebus.SubscriptionRule("exampleSubscriptionRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            subscription_name=example_subscription.name,
            filter_type="SqlFilter",
            sql_filter="colour = 'red'")
        ```
        ### Correlation Filter)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_topic = azure.servicebus.Topic("exampleTopic",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_subscription = azure.servicebus.Subscription("exampleSubscription",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            max_delivery_count=1)
        example_subscription_rule = azure.servicebus.SubscriptionRule("exampleSubscriptionRule",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            topic_name=example_topic.name,
            subscription_name=example_subscription.name,
            filter_type="CorrelationFilter",
            correlation_filter=azure.servicebus.SubscriptionRuleCorrelationFilterArgs(
                correlation_id="high",
                label="red",
                properties={
                    "customProperty": "value",
                },
            ))
        ```

        ## Import

        Service Bus Subscription Rule can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:servicebus/subscriptionRule:SubscriptionRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
        ```

        :param str resource_name: The name of the resource.
        :param SubscriptionRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubscriptionRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 correlation_filter: Optional[pulumi.Input[pulumi.InputType['SubscriptionRuleCorrelationFilterArgs']]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sql_filter: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = SubscriptionRuleArgs.__new__(SubscriptionRuleArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["correlation_filter"] = correlation_filter
            if filter_type is None and not opts.urn:
                raise TypeError("Missing required property 'filter_type'")
            __props__.__dict__["filter_type"] = filter_type
            __props__.__dict__["name"] = name
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["sql_filter"] = sql_filter
            if subscription_name is None and not opts.urn:
                raise TypeError("Missing required property 'subscription_name'")
            __props__.__dict__["subscription_name"] = subscription_name
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:eventhub/subscriptionRule:SubscriptionRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SubscriptionRule, __self__).__init__(
            'azure:servicebus/subscriptionRule:SubscriptionRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            correlation_filter: Optional[pulumi.Input[pulumi.InputType['SubscriptionRuleCorrelationFilterArgs']]] = None,
            filter_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sql_filter: Optional[pulumi.Input[str]] = None,
            subscription_name: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'SubscriptionRule':
        """
        Get an existing SubscriptionRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        :param pulumi.Input[pulumi.InputType['SubscriptionRuleCorrelationFilterArgs']] correlation_filter: A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        :param pulumi.Input[str] filter_type: Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_filter: Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        :param pulumi.Input[str] subscription_name: The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubscriptionRuleState.__new__(_SubscriptionRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["correlation_filter"] = correlation_filter
        __props__.__dict__["filter_type"] = filter_type
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["sql_filter"] = sql_filter
        __props__.__dict__["subscription_name"] = subscription_name
        __props__.__dict__["topic_name"] = topic_name
        return SubscriptionRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[str]]:
        """
        Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="correlationFilter")
    def correlation_filter(self) -> pulumi.Output[Optional['outputs.SubscriptionRuleCorrelationFilter']]:
        """
        A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        """
        return pulumi.get(self, "correlation_filter")

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> pulumi.Output[str]:
        """
        Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        """
        return pulumi.get(self, "filter_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the ServiceBus Namespace in which the ServiceBus Topic exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sqlFilter")
    def sql_filter(self) -> pulumi.Output[Optional[str]]:
        """
        Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        """
        return pulumi.get(self, "sql_filter")

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Output[str]:
        """
        The name of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subscription_name")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        The name of the ServiceBus Topic in which the ServiceBus Subscription exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "topic_name")

