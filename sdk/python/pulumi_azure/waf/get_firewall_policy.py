# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetFirewallPolicyResult',
    'AwaitableGetFirewallPolicyResult',
    'get_firewall_policy',
]

@pulumi.output_type
class GetFirewallPolicyResult:
    """
    A collection of values returned by getFirewallPolicy.
    """
    def __init__(__self__, id=None, location=None, name=None, resource_group_name=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")


class AwaitableGetFirewallPolicyResult(GetFirewallPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicyResult(
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_firewall_policy(name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        tags: Optional[Mapping[str, str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallPolicyResult:
    """
    Use this data source to access information about an existing Web Application Firewall Policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.waf.get_firewall_policy(resource_group_name="existing",
        name="existing")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of the Web Application Firewall Policy
    :param str resource_group_name: The name of the Resource Group where the Web Application Firewall Policy exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:waf/getFirewallPolicy:getFirewallPolicy', __args__, opts=opts, typ=GetFirewallPolicyResult).value

    return AwaitableGetFirewallPolicyResult(
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags)
