# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetVirtualNetworkResult(object):
    """
    A collection of values returned by getVirtualNetwork.
    """
    def __init__(__self__, address_spaces=None, dns_servers=None, subnets=None, vnet_peerings=None):
        if not address_spaces:
            raise TypeError('Missing required argument address_spaces')
        elif not isinstance(address_spaces, list):
            raise TypeError('Expected argument address_spaces to be a list')
        __self__.address_spaces = address_spaces
        """
        The list of address spaces used by the virtual network.
        """
        if not dns_servers:
            raise TypeError('Missing required argument dns_servers')
        elif not isinstance(dns_servers, list):
            raise TypeError('Expected argument dns_servers to be a list')
        __self__.dns_servers = dns_servers
        """
        The list of DNS servers used by the virtual network.
        """
        if not subnets:
            raise TypeError('Missing required argument subnets')
        elif not isinstance(subnets, list):
            raise TypeError('Expected argument subnets to be a list')
        __self__.subnets = subnets
        """
        The list of name of the subnets that are attached to this virtual network.
        """
        if not vnet_peerings:
            raise TypeError('Missing required argument vnet_peerings')
        elif not isinstance(vnet_peerings, dict):
            raise TypeError('Expected argument vnet_peerings to be a dict')
        __self__.vnet_peerings = vnet_peerings
        """
        A mapping of name - virtual network id of the virtual network peerings.
        """

def get_virtual_network(name=None, resource_group_name=None):
    """
    Use this data source to access the properties of an Azure Virtual Network.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __ret__ = pulumi.runtime.invoke('azure:network/getVirtualNetwork:getVirtualNetwork', __args__)

    return GetVirtualNetworkResult(
        address_spaces=__ret__['addressSpaces'],
        dns_servers=__ret__['dnsServers'],
        subnets=__ret__['subnets'],
        vnet_peerings=__ret__['vnetPeerings'])
