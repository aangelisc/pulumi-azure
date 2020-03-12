# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetPublicIpPrefixResult:
    """
    A collection of values returned by getPublicIpPrefix.
    """
    def __init__(__self__, id=None, ip_prefix=None, location=None, name=None, prefix_length=None, resource_group_name=None, sku=None, tags=None, zones=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ip_prefix and not isinstance(ip_prefix, str):
            raise TypeError("Expected argument 'ip_prefix' to be a str")
        __self__.ip_prefix = ip_prefix
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The supported Azure location where the resource exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the Public IP prefix resource.
        """
        if prefix_length and not isinstance(prefix_length, float):
            raise TypeError("Expected argument 'prefix_length' to be a float")
        __self__.prefix_length = prefix_length
        """
        The number of bits of the prefix.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the public IP.
        """
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        __self__.sku = sku
        """
        The SKU of the Public IP Prefix.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assigned to the resource.
        """
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones
class AwaitableGetPublicIpPrefixResult(GetPublicIpPrefixResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublicIpPrefixResult(
            id=self.id,
            ip_prefix=self.ip_prefix,
            location=self.location,
            name=self.name,
            prefix_length=self.prefix_length,
            resource_group_name=self.resource_group_name,
            sku=self.sku,
            tags=self.tags,
            zones=self.zones)

def get_public_ip_prefix(name=None,resource_group_name=None,zones=None,opts=None):
    """
    Use this data source to access information about an existing Public IP Prefix.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/public_ip_prefix.html.markdown.


    :param str name: Specifies the name of the public IP prefix.
    :param str resource_group_name: Specifies the name of the resource group.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['zones'] = zones
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getPublicIpPrefix:getPublicIpPrefix', __args__, opts=opts).value

    return AwaitableGetPublicIpPrefixResult(
        id=__ret__.get('id'),
        ip_prefix=__ret__.get('ipPrefix'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        prefix_length=__ret__.get('prefixLength'),
        resource_group_name=__ret__.get('resourceGroupName'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        zones=__ret__.get('zones'))
