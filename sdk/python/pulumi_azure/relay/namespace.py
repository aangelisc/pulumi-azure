# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Namespace(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
    """
    metric_id: pulumi.Output[str]
    """
    The Identifier for Azure Insights metrics.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
    """
    primary_connection_string: pulumi.Output[str]
    """
    The primary connection string for the authorization rule `RootManageSharedAccessKey`.
    """
    primary_key: pulumi.Output[str]
    """
    The primary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Azure Relay Namespace.
    """
    secondary_connection_string: pulumi.Output[str]
    """
    The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
    """
    secondary_key: pulumi.Output[str]
    """
    The secondary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    sku_name: pulumi.Output[str]
    """
    The name of the SKU to use. At this time the only supported value is `Standard`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Relay Namespace.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/relay_namespace.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Namespace.
        :param pulumi.Input[str] sku_name: The name of the SKU to use. At this time the only supported value is `Standard`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['metric_id'] = None
            __props__['primary_connection_string'] = None
            __props__['primary_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['secondary_key'] = None
        super(Namespace, __self__).__init__(
            'azure:relay/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, metric_id=None, name=None, primary_connection_string=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_key=None, sku_name=None, tags=None):
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metric_id: The Identifier for Azure Insights metrics.
        :param pulumi.Input[str] name: Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_connection_string: The primary connection string for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] primary_key: The primary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Azure Relay Namespace.
        :param pulumi.Input[str] secondary_connection_string: The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] secondary_key: The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] sku_name: The name of the SKU to use. At this time the only supported value is `Standard`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["metric_id"] = metric_id
        __props__["name"] = name
        __props__["primary_connection_string"] = primary_connection_string
        __props__["primary_key"] = primary_key
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["secondary_key"] = secondary_key
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        return Namespace(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

