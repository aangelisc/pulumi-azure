# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SubscriptionArgs', 'Subscription']

@pulumi.input_type
class SubscriptionArgs:
    def __init__(__self__, *,
                 subscription_name: pulumi.Input[str],
                 alias: Optional[pulumi.Input[str]] = None,
                 billing_scope_id: Optional[pulumi.Input[str]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 workload: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Subscription resource.
        :param pulumi.Input[str] subscription_name: The Name of the Subscription. This is the Display Name in the portal.
        :param pulumi.Input[str] alias: The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_id: The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] workload: The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        pulumi.set(__self__, "subscription_name", subscription_name)
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if billing_scope_id is not None:
            pulumi.set(__self__, "billing_scope_id", billing_scope_id)
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)
        if workload is not None:
            pulumi.set(__self__, "workload", workload)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Input[str]:
        """
        The Name of the Subscription. This is the Display Name in the portal.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="billingScopeId")
    def billing_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "billing_scope_id")

    @billing_scope_id.setter
    def billing_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_scope_id", value)

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "subscription_id")

    @subscription_id.setter
    def subscription_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_id", value)

    @property
    @pulumi.getter
    def workload(self) -> Optional[pulumi.Input[str]]:
        """
        The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "workload")

    @workload.setter
    def workload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workload", value)


@pulumi.input_type
class _SubscriptionState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 billing_scope_id: Optional[pulumi.Input[str]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 workload: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Subscription resources.
        :param pulumi.Input[str] alias: The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_id: The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_name: The Name of the Subscription. This is the Display Name in the portal.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant to which the subscription belongs.
        :param pulumi.Input[str] workload: The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if billing_scope_id is not None:
            pulumi.set(__self__, "billing_scope_id", billing_scope_id)
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)
        if subscription_name is not None:
            pulumi.set(__self__, "subscription_name", subscription_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if workload is not None:
            pulumi.set(__self__, "workload", workload)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="billingScopeId")
    def billing_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "billing_scope_id")

    @billing_scope_id.setter
    def billing_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_scope_id", value)

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "subscription_id")

    @subscription_id.setter
    def subscription_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_id", value)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the Subscription. This is the Display Name in the portal.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Tenant to which the subscription belongs.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def workload(self) -> Optional[pulumi.Input[str]]:
        """
        The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "workload")

    @workload.setter
    def workload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workload", value)


class Subscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 billing_scope_id: Optional[pulumi.Input[str]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 workload: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Alias for a Subscription - which adds an Alias to an existing Subscription, allowing it to be managed in the provider - or create a new Subscription with a new Alias.

        > **NOTE:** Destroying a Subscription controlled by this resource will place the Subscription into a cancelled state. It is possible to re-activate a subscription within 90-days of cancellation, after which time the Subscription is irrevocably deleted, and the Subscription ID cannot be re-used. For further information see [here](https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/cancel-azure-subscription#what-happens-after-subscription-cancellation). Users can optionally delete a Subscription once 72 hours have passed, however, this functionality is not suitable for this provider. A `Deleted` subscription cannot be reactivated.

        > **NOTE:** It is not possible to destroy (cancel) a subscription if it contains resources. If resources are present that are not managed by the provider then these will need to be removed before the Subscription can be destroyed.

        > **NOTE:** Azure supports Multiple Aliases per Subscription, however, to reliably manage this resource in this provider only a single Alias is supported.

        ## Example Usage
        ### Creating A New Alias And Subscription For An Enrollment Account

        ```python
        import pulumi
        import pulumi_azure as azure

        example_enrollment_account_scope = azure.billing.get_enrollment_account_scope(billing_account_name="1234567890",
            enrollment_account_name="0123456")
        example_subscription = azure.core.Subscription("exampleSubscription",
            subscription_name="My Example EA Subscription",
            billing_scope_id=example_enrollment_account_scope.id)
        ```
        ### Creating A New Alias And Subscription For A Microsoft Customer Account

        ```python
        import pulumi
        import pulumi_azure as azure

        example_mca_account_scope = azure.billing.get_mca_account_scope(billing_account_name="e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
            billing_profile_name="PE2Q-NOIT-BG7-TGB",
            invoice_section_name="MTT4-OBS7-PJA-TGB")
        example_subscription = azure.core.Subscription("exampleSubscription",
            subscription_name="My Example MCA Subscription",
            billing_scope_id=example_mca_account_scope.id)
        ```
        ### Adding An Alias To An Existing Subscription

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.Subscription("example",
            alias="examplesub",
            subscription_id="12345678-12234-5678-9012-123456789012",
            subscription_name="My Example Subscription")
        ```

        ## Import

        Subscriptions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:core/subscription:Subscription example "/providers/Microsoft.Subscription/aliases/subscription1"
        ```

         In this scenario, the `subscription_id` property can be completed and the provider will assume control of the existing subscription by creating an Alias. e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_id: The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_name: The Name of the Subscription. This is the Display Name in the portal.
        :param pulumi.Input[str] workload: The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Alias for a Subscription - which adds an Alias to an existing Subscription, allowing it to be managed in the provider - or create a new Subscription with a new Alias.

        > **NOTE:** Destroying a Subscription controlled by this resource will place the Subscription into a cancelled state. It is possible to re-activate a subscription within 90-days of cancellation, after which time the Subscription is irrevocably deleted, and the Subscription ID cannot be re-used. For further information see [here](https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/cancel-azure-subscription#what-happens-after-subscription-cancellation). Users can optionally delete a Subscription once 72 hours have passed, however, this functionality is not suitable for this provider. A `Deleted` subscription cannot be reactivated.

        > **NOTE:** It is not possible to destroy (cancel) a subscription if it contains resources. If resources are present that are not managed by the provider then these will need to be removed before the Subscription can be destroyed.

        > **NOTE:** Azure supports Multiple Aliases per Subscription, however, to reliably manage this resource in this provider only a single Alias is supported.

        ## Example Usage
        ### Creating A New Alias And Subscription For An Enrollment Account

        ```python
        import pulumi
        import pulumi_azure as azure

        example_enrollment_account_scope = azure.billing.get_enrollment_account_scope(billing_account_name="1234567890",
            enrollment_account_name="0123456")
        example_subscription = azure.core.Subscription("exampleSubscription",
            subscription_name="My Example EA Subscription",
            billing_scope_id=example_enrollment_account_scope.id)
        ```
        ### Creating A New Alias And Subscription For A Microsoft Customer Account

        ```python
        import pulumi
        import pulumi_azure as azure

        example_mca_account_scope = azure.billing.get_mca_account_scope(billing_account_name="e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
            billing_profile_name="PE2Q-NOIT-BG7-TGB",
            invoice_section_name="MTT4-OBS7-PJA-TGB")
        example_subscription = azure.core.Subscription("exampleSubscription",
            subscription_name="My Example MCA Subscription",
            billing_scope_id=example_mca_account_scope.id)
        ```
        ### Adding An Alias To An Existing Subscription

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.Subscription("example",
            alias="examplesub",
            subscription_id="12345678-12234-5678-9012-123456789012",
            subscription_name="My Example Subscription")
        ```

        ## Import

        Subscriptions can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:core/subscription:Subscription example "/providers/Microsoft.Subscription/aliases/subscription1"
        ```

         In this scenario, the `subscription_id` property can be completed and the provider will assume control of the existing subscription by creating an Alias. e.g.

        :param str resource_name: The name of the resource.
        :param SubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 billing_scope_id: Optional[pulumi.Input[str]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 workload: Optional[pulumi.Input[str]] = None,
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
            __props__ = SubscriptionArgs.__new__(SubscriptionArgs)

            __props__.__dict__["alias"] = alias
            __props__.__dict__["billing_scope_id"] = billing_scope_id
            __props__.__dict__["subscription_id"] = subscription_id
            if subscription_name is None and not opts.urn:
                raise TypeError("Missing required property 'subscription_name'")
            __props__.__dict__["subscription_name"] = subscription_name
            __props__.__dict__["workload"] = workload
            __props__.__dict__["tags"] = None
            __props__.__dict__["tenant_id"] = None
        super(Subscription, __self__).__init__(
            'azure:core/subscription:Subscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            billing_scope_id: Optional[pulumi.Input[str]] = None,
            subscription_id: Optional[pulumi.Input[str]] = None,
            subscription_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            workload: Optional[pulumi.Input[str]] = None) -> 'Subscription':
        """
        Get an existing Subscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_id: The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        :param pulumi.Input[str] subscription_name: The Name of the Subscription. This is the Display Name in the portal.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant to which the subscription belongs.
        :param pulumi.Input[str] workload: The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubscriptionState.__new__(_SubscriptionState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["billing_scope_id"] = billing_scope_id
        __props__.__dict__["subscription_id"] = subscription_id
        __props__.__dict__["subscription_name"] = subscription_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["workload"] = workload
        return Subscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="billingScopeId")
    def billing_scope_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "billing_scope_id")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[str]:
        """
        The ID of the Subscription. Cannot be specified with `billing_account`, `billing_profile`, `enrollment_account`, or `invoice_section` Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Output[str]:
        """
        The Name of the Subscription. This is the Display Name in the portal.
        """
        return pulumi.get(self, "subscription_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The ID of the Tenant to which the subscription belongs.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def workload(self) -> pulumi.Output[Optional[str]]:
        """
        The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        """
        return pulumi.get(self, "workload")

