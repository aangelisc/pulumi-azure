# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Group(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, containers=None, dns_name_label=None, ip_address_type=None, location=None, name=None, os_type=None, resource_group_name=None, restart_policy=None, tags=None):
        """Create a Group resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not containers:
            raise TypeError('Missing required property containers')
        elif not isinstance(containers, list):
            raise TypeError('Expected property containers to be a list')
        __self__.containers = containers
        __props__['containers'] = containers

        if dns_name_label and not isinstance(dns_name_label, basestring):
            raise TypeError('Expected property dns_name_label to be a basestring')
        __self__.dns_name_label = dns_name_label
        __props__['dnsNameLabel'] = dns_name_label

        if ip_address_type and not isinstance(ip_address_type, basestring):
            raise TypeError('Expected property ip_address_type to be a basestring')
        __self__.ip_address_type = ip_address_type
        __props__['ipAddressType'] = ip_address_type

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        __props__['name'] = name

        if not os_type:
            raise TypeError('Missing required property os_type')
        elif not isinstance(os_type, basestring):
            raise TypeError('Expected property os_type to be a basestring')
        __self__.os_type = os_type
        __props__['osType'] = os_type

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        __props__['resourceGroupName'] = resource_group_name

        if restart_policy and not isinstance(restart_policy, basestring):
            raise TypeError('Expected property restart_policy to be a basestring')
        __self__.restart_policy = restart_policy
        __props__['restartPolicy'] = restart_policy

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        __props__['tags'] = tags

        __self__.fqdn = pulumi.runtime.UNKNOWN
        __self__.ip_address = pulumi.runtime.UNKNOWN

        super(Group, __self__).__init__(
            'azure:containerservice/group:Group',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'containers' in outs:
            self.containers = outs['containers']
        if 'dnsNameLabel' in outs:
            self.dns_name_label = outs['dnsNameLabel']
        if 'fqdn' in outs:
            self.fqdn = outs['fqdn']
        if 'ipAddress' in outs:
            self.ip_address = outs['ipAddress']
        if 'ipAddressType' in outs:
            self.ip_address_type = outs['ipAddressType']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'osType' in outs:
            self.os_type = outs['osType']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'restartPolicy' in outs:
            self.restart_policy = outs['restartPolicy']
        if 'tags' in outs:
            self.tags = outs['tags']
