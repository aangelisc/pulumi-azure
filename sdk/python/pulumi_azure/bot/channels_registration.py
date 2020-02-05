# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ChannelsRegistration(pulumi.CustomResource):
    developer_app_insights_api_key: pulumi.Output[str]
    """
    The Application Insights API Key to associate with the Bot Channels Registration.
    """
    developer_app_insights_application_id: pulumi.Output[str]
    """
    The Application Insights Application ID to associate with the Bot Channels Registration.
    """
    developer_app_insights_key: pulumi.Output[str]
    """
    The Application Insights Key to associate with the Bot Channels Registration.
    """
    display_name: pulumi.Output[str]
    """
    The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
    """
    endpoint: pulumi.Output[str]
    """
    The Bot Channels Registration endpoint.
    """
    location: pulumi.Output[str]
    """
    The supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    microsoft_app_id: pulumi.Output[str]
    """
    The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[str]
    """
    The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, developer_app_insights_api_key=None, developer_app_insights_application_id=None, developer_app_insights_key=None, display_name=None, endpoint=None, location=None, microsoft_app_id=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Bot Channels Registration.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights API Key to associate with the Bot Channels Registration.
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights Application ID to associate with the Bot Channels Registration.
        :param pulumi.Input[str] developer_app_insights_key: The Application Insights Key to associate with the Bot Channels Registration.
        :param pulumi.Input[str] display_name: The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
        :param pulumi.Input[str] endpoint: The Bot Channels Registration endpoint.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] microsoft_app_id: The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channels_registration.html.markdown.
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

            __props__['developer_app_insights_api_key'] = developer_app_insights_api_key
            __props__['developer_app_insights_application_id'] = developer_app_insights_application_id
            __props__['developer_app_insights_key'] = developer_app_insights_key
            __props__['display_name'] = display_name
            __props__['endpoint'] = endpoint
            __props__['location'] = location
            if microsoft_app_id is None:
                raise TypeError("Missing required property 'microsoft_app_id'")
            __props__['microsoft_app_id'] = microsoft_app_id
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
        super(ChannelsRegistration, __self__).__init__(
            'azure:bot/channelsRegistration:ChannelsRegistration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, developer_app_insights_api_key=None, developer_app_insights_application_id=None, developer_app_insights_key=None, display_name=None, endpoint=None, location=None, microsoft_app_id=None, name=None, resource_group_name=None, sku=None, tags=None):
        """
        Get an existing ChannelsRegistration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights API Key to associate with the Bot Channels Registration.
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights Application ID to associate with the Bot Channels Registration.
        :param pulumi.Input[str] developer_app_insights_key: The Application Insights Key to associate with the Bot Channels Registration.
        :param pulumi.Input[str] display_name: The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
        :param pulumi.Input[str] endpoint: The Bot Channels Registration endpoint.
        :param pulumi.Input[str] location: The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] microsoft_app_id: The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channels_registration.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["developer_app_insights_api_key"] = developer_app_insights_api_key
        __props__["developer_app_insights_application_id"] = developer_app_insights_application_id
        __props__["developer_app_insights_key"] = developer_app_insights_key
        __props__["display_name"] = display_name
        __props__["endpoint"] = endpoint
        __props__["location"] = location
        __props__["microsoft_app_id"] = microsoft_app_id
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        return ChannelsRegistration(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

