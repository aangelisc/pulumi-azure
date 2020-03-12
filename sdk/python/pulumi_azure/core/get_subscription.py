# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSubscriptionResult:
    """
    A collection of values returned by getSubscription.
    """
    def __init__(__self__, display_name=None, id=None, location_placement_id=None, quota_id=None, spending_limit=None, state=None, subscription_id=None, tenant_id=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        The subscription display name.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if location_placement_id and not isinstance(location_placement_id, str):
            raise TypeError("Expected argument 'location_placement_id' to be a str")
        __self__.location_placement_id = location_placement_id
        """
        The subscription location placement ID.
        """
        if quota_id and not isinstance(quota_id, str):
            raise TypeError("Expected argument 'quota_id' to be a str")
        __self__.quota_id = quota_id
        """
        The subscription quota ID.
        """
        if spending_limit and not isinstance(spending_limit, str):
            raise TypeError("Expected argument 'spending_limit' to be a str")
        __self__.spending_limit = spending_limit
        """
        The subscription spending limit.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
        """
        if subscription_id and not isinstance(subscription_id, str):
            raise TypeError("Expected argument 'subscription_id' to be a str")
        __self__.subscription_id = subscription_id
        """
        The subscription GUID.
        """
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        __self__.tenant_id = tenant_id
        """
        The subscription tenant ID.
        """
class AwaitableGetSubscriptionResult(GetSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubscriptionResult(
            display_name=self.display_name,
            id=self.id,
            location_placement_id=self.location_placement_id,
            quota_id=self.quota_id,
            spending_limit=self.spending_limit,
            state=self.state,
            subscription_id=self.subscription_id,
            tenant_id=self.tenant_id)

def get_subscription(subscription_id=None,opts=None):
    """
    Use this data source to access information about an existing Subscription.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subscription.html.markdown.


    :param str subscription_id: Specifies the ID of the subscription. If this argument is omitted, the subscription ID of the current Azure Resource Manager provider is used.
    """
    __args__ = dict()


    __args__['subscriptionId'] = subscription_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:core/getSubscription:getSubscription', __args__, opts=opts).value

    return AwaitableGetSubscriptionResult(
        display_name=__ret__.get('displayName'),
        id=__ret__.get('id'),
        location_placement_id=__ret__.get('locationPlacementId'),
        quota_id=__ret__.get('quotaId'),
        spending_limit=__ret__.get('spendingLimit'),
        state=__ret__.get('state'),
        subscription_id=__ret__.get('subscriptionId'),
        tenant_id=__ret__.get('tenantId'))
