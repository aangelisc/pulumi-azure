# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class BackendAddressPool(pulumi.CustomResource):
    """
    Create a LoadBalancer Backend Address Pool.
    
    ~> **NOTE:** When using this resource, the LoadBalancer needs to have a FrontEnd IP Configuration Attached
    """
    def __init__(__self__, __name__, __opts__=None, loadbalancer_id=None, location=None, name=None, resource_group_name=None):
        """Create a BackendAddressPool resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not loadbalancer_id:
            raise TypeError('Missing required property loadbalancer_id')
        elif not isinstance(loadbalancer_id, basestring):
            raise TypeError('Expected property loadbalancer_id to be a basestring')
        __self__.loadbalancer_id = loadbalancer_id
        """
        The ID of the LoadBalancer in which to create the Backend Address Pool.
        """
        __props__['loadbalancerId'] = loadbalancer_id

        if location and not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the Backend Address Pool.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the resource.
        """
        __props__['resourceGroupName'] = resource_group_name

        __self__.backend_ip_configurations = pulumi.runtime.UNKNOWN
        __self__.load_balancing_rules = pulumi.runtime.UNKNOWN

        super(BackendAddressPool, __self__).__init__(
            'azure:lb/backendAddressPool:BackendAddressPool',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'backendIpConfigurations' in outs:
            self.backend_ip_configurations = outs['backendIpConfigurations']
        if 'loadBalancingRules' in outs:
            self.load_balancing_rules = outs['loadBalancingRules']
        if 'loadbalancerId' in outs:
            self.loadbalancer_id = outs['loadbalancerId']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
