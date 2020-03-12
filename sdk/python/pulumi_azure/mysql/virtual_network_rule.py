# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VirtualNetworkRule(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
    """
    server_name: pulumi.Output[str]
    """
    The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
    """
    subnet_id: pulumi.Output[str]
    """
    The ID of the subnet that the MySQL server will be connected to.
    """
    def __init__(__self__, resource_name, opts=None, name=None, resource_group_name=None, server_name=None, subnet_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a MySQL Virtual Network Rule.

        > **NOTE:** MySQL Virtual Network Rules [can only be used with SKU Tiers of `GeneralPurpose` or `MemoryOptimized`](https://docs.microsoft.com/en-us/azure/mysql/concepts-data-access-and-security-vnet)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/mysql_virtual_network_rule.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet that the MySQL server will be connected to.
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

            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
        super(VirtualNetworkRule, __self__).__init__(
            'azure:mysql/virtualNetworkRule:VirtualNetworkRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, resource_group_name=None, server_name=None, subnet_id=None):
        """
        Get an existing VirtualNetworkRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_name: The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet that the MySQL server will be connected to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["server_name"] = server_name
        __props__["subnet_id"] = subnet_id
        return VirtualNetworkRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

