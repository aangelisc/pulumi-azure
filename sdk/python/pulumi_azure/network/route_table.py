# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RouteTable(pulumi.CustomResource):
    disable_bgp_route_propagation: pulumi.Output[bool]
    """
    Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the route.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
    """
    routes: pulumi.Output[list]
    """
    [List of objects](https://www.terraform.io/docs/configuration/attr-as-blocks.html) representing routes. Each object accepts the arguments documented below.
    """
    subnets: pulumi.Output[list]
    """
    The collection of Subnets associated with this route table.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, disable_bgp_route_propagation=None, location=None, name=None, resource_group_name=None, routes=None, tags=None, __name__=None, __opts__=None):
        """
        Manages a Route Table
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_bgp_route_propagation: Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the route.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
        :param pulumi.Input[list] routes: [List of objects](https://www.terraform.io/docs/configuration/attr-as-blocks.html) representing routes. Each object accepts the arguments documented below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/route_table.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['disable_bgp_route_propagation'] = disable_bgp_route_propagation

        __props__['location'] = location

        __props__['name'] = name

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['routes'] = routes

        __props__['tags'] = tags

        __props__['subnets'] = None

        super(RouteTable, __self__).__init__(
            'azure:network/routeTable:RouteTable',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

