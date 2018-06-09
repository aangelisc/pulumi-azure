# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class VirtualNetwork(pulumi.CustomResource):
    """
    Manages a new virtual network including any configured subnets. Each subnet can
    optionally be configured with a security group to be associated with the subnet.
    
    ~> **NOTE on Virtual Networks and Subnet's:** Terraform currently
    provides both a standalone [Subnet resource](subnet.html), and allows for Subnets to be defined in-line within the [Virtual Network resource](virtual_network.html).
    At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
    """
    def __init__(__self__, __name__, __opts__=None, address_spaces=None, dns_servers=None, location=None, name=None, resource_group_name=None, subnets=None, tags=None):
        """Create a VirtualNetwork resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not address_spaces:
            raise TypeError('Missing required property address_spaces')
        elif not isinstance(address_spaces, list):
            raise TypeError('Expected property address_spaces to be a list')
        __self__.address_spaces = address_spaces
        """
        The address space that is used the virtual
        network. You can supply more than one address space. Changing this forces
        a new resource to be created.
        """
        __props__['addressSpaces'] = address_spaces

        if dns_servers and not isinstance(dns_servers, list):
            raise TypeError('Expected property dns_servers to be a list')
        __self__.dns_servers = dns_servers
        """
        List of IP addresses of DNS servers
        """
        __props__['dnsServers'] = dns_servers

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The location/region where the virtual network is
        created. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the subnet.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to
        create the virtual network.
        """
        __props__['resourceGroupName'] = resource_group_name

        if subnets and not isinstance(subnets, list):
            raise TypeError('Expected property subnets to be a list')
        __self__.subnets = subnets
        """
        Can be specified multiple times to define multiple
        subnets. Each `subnet` block supports fields documented below.
        """
        __props__['subnets'] = subnets

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        super(VirtualNetwork, __self__).__init__(
            'azure:network/virtualNetwork:VirtualNetwork',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'addressSpaces' in outs:
            self.address_spaces = outs['addressSpaces']
        if 'dnsServers' in outs:
            self.dns_servers = outs['dnsServers']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'subnets' in outs:
            self.subnets = outs['subnets']
        if 'tags' in outs:
            self.tags = outs['tags']
