# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetNetworkSecurityGroupResult(object):
    """
    A collection of values returned by getNetworkSecurityGroup.
    """
    def __init__(__self__, location=None, security_rules=None, tags=None):
        if not location:
            raise TypeError('Missing required argument location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected argument location to be a basestring')
        __self__.location = location
        """
        The supported Azure location where the resource exists.
        """
        if not security_rules:
            raise TypeError('Missing required argument security_rules')
        elif not isinstance(security_rules, list):
            raise TypeError('Expected argument security_rules to be a list')
        __self__.security_rules = security_rules
        """
        One or more `security_rule` blocks as defined below.
        """
        if not tags:
            raise TypeError('Missing required argument tags')
        elif not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """

def get_network_security_group(name=None, resource_group_name=None):
    """
    Use this data source to access the properties of a Network Security Group.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __ret__ = pulumi.runtime.invoke('azure:network/getNetworkSecurityGroup:getNetworkSecurityGroup', __args__)

    return GetNetworkSecurityGroupResult(
        location=__ret__['location'],
        security_rules=__ret__['securityRules'],
        tags=__ret__['tags'])
