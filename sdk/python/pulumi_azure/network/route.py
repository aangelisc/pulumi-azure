# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Route(pulumi.CustomResource):
    address_prefix: pulumi.Output[str]
    """
    The destination CIDR to which the route applies, such as `10.1.0.0/16`
    """
    name: pulumi.Output[str]
    """
    The name of the route. Changing this forces a new resource to be created.
    """
    next_hop_in_ip_address: pulumi.Output[str]
    """
    Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
    """
    next_hop_type: pulumi.Output[str]
    """
    The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the route. Changing this forces a new resource to be created.
    """
    route_table_name: pulumi.Output[str]
    """
    The name of the route table within which create the route. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, address_prefix=None, name=None, next_hop_in_ip_address=None, next_hop_type=None, resource_group_name=None, route_table_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Route within a Route Table.

        > **NOTE on Route Tables and Routes:** This provider currently
        provides both a standalone Route resource, and allows for Routes to be defined in-line within the Route Table resource.
        At this time you cannot use a Route Table with in-line Routes in conjunction with any Route resources. Doing so will cause a conflict of Route configurations and will overwrite Routes.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/route.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The destination CIDR to which the route applies, such as `10.1.0.0/16`
        :param pulumi.Input[str] name: The name of the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] next_hop_in_ip_address: Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        :param pulumi.Input[str] next_hop_type: The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_table_name: The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if address_prefix is None:
                raise TypeError("Missing required property 'address_prefix'")
            __props__['address_prefix'] = address_prefix
            __props__['name'] = name
            __props__['next_hop_in_ip_address'] = next_hop_in_ip_address
            if next_hop_type is None:
                raise TypeError("Missing required property 'next_hop_type'")
            __props__['next_hop_type'] = next_hop_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if route_table_name is None:
                raise TypeError("Missing required property 'route_table_name'")
            __props__['route_table_name'] = route_table_name
        super(Route, __self__).__init__(
            'azure:network/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address_prefix=None, name=None, next_hop_in_ip_address=None, next_hop_type=None, resource_group_name=None, route_table_name=None):
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: The destination CIDR to which the route applies, such as `10.1.0.0/16`
        :param pulumi.Input[str] name: The name of the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] next_hop_in_ip_address: Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        :param pulumi.Input[str] next_hop_type: The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_table_name: The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_prefix"] = address_prefix
        __props__["name"] = name
        __props__["next_hop_in_ip_address"] = next_hop_in_ip_address
        __props__["next_hop_type"] = next_hop_type
        __props__["resource_group_name"] = resource_group_name
        __props__["route_table_name"] = route_table_name
        return Route(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

