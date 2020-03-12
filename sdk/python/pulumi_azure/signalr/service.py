# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Service(pulumi.CustomResource):
    cors: pulumi.Output[list]
    """
    A `cors` block as documented below.

      * `allowedOrigins` (`list`)
    """
    features: pulumi.Output[list]
    """
    A `features` block as documented below.

      * `flag` (`str`)
      * `value` (`str`)
    """
    hostname: pulumi.Output[str]
    """
    The FQDN of the SignalR service.
    """
    ip_address: pulumi.Output[str]
    """
    The publicly accessible IP of the SignalR service.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the SignalR service. Changing this forces a new resource to be created.
    """
    primary_access_key: pulumi.Output[str]
    """
    The primary access key for the SignalR service.
    """
    primary_connection_string: pulumi.Output[str]
    """
    The primary connection string for the SignalR service.
    """
    public_port: pulumi.Output[float]
    """
    The publicly accessible port of the SignalR service which is designed for browser/client use.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
    """
    secondary_access_key: pulumi.Output[str]
    """
    The secondary access key for the SignalR service.
    """
    secondary_connection_string: pulumi.Output[str]
    """
    The secondary connection string for the SignalR service.
    """
    server_port: pulumi.Output[float]
    """
    The publicly accessible port of the SignalR service which is designed for customer server side use.
    """
    sku: pulumi.Output[dict]
    """
    A `sku` block as documented below.

      * `capacity` (`float`)
      * `name` (`str`) - The name of the SignalR service. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, cors=None, features=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure SignalR service.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/signalr_service.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] cors: A `cors` block as documented below.
        :param pulumi.Input[list] features: A `features` block as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] sku: A `sku` block as documented below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **cors** object supports the following:

          * `allowedOrigins` (`pulumi.Input[list]`)

        The **features** object supports the following:

          * `flag` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`)
          * `name` (`pulumi.Input[str]`) - The name of the SignalR service. Changing this forces a new resource to be created.
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

            __props__['cors'] = cors
            __props__['features'] = features
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['hostname'] = None
            __props__['ip_address'] = None
            __props__['primary_access_key'] = None
            __props__['primary_connection_string'] = None
            __props__['public_port'] = None
            __props__['secondary_access_key'] = None
            __props__['secondary_connection_string'] = None
            __props__['server_port'] = None
        super(Service, __self__).__init__(
            'azure:signalr/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cors=None, features=None, hostname=None, ip_address=None, location=None, name=None, primary_access_key=None, primary_connection_string=None, public_port=None, resource_group_name=None, secondary_access_key=None, secondary_connection_string=None, server_port=None, sku=None, tags=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] cors: A `cors` block as documented below.
        :param pulumi.Input[list] features: A `features` block as documented below.
        :param pulumi.Input[str] hostname: The FQDN of the SignalR service.
        :param pulumi.Input[str] ip_address: The publicly accessible IP of the SignalR service.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The primary access key for the SignalR service.
        :param pulumi.Input[str] primary_connection_string: The primary connection string for the SignalR service.
        :param pulumi.Input[float] public_port: The publicly accessible port of the SignalR service which is designed for browser/client use.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The secondary access key for the SignalR service.
        :param pulumi.Input[str] secondary_connection_string: The secondary connection string for the SignalR service.
        :param pulumi.Input[float] server_port: The publicly accessible port of the SignalR service which is designed for customer server side use.
        :param pulumi.Input[dict] sku: A `sku` block as documented below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **cors** object supports the following:

          * `allowedOrigins` (`pulumi.Input[list]`)

        The **features** object supports the following:

          * `flag` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`)
          * `name` (`pulumi.Input[str]`) - The name of the SignalR service. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cors"] = cors
        __props__["features"] = features
        __props__["hostname"] = hostname
        __props__["ip_address"] = ip_address
        __props__["location"] = location
        __props__["name"] = name
        __props__["primary_access_key"] = primary_access_key
        __props__["primary_connection_string"] = primary_connection_string
        __props__["public_port"] = public_port
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["secondary_connection_string"] = secondary_connection_string
        __props__["server_port"] = server_port
        __props__["sku"] = sku
        __props__["tags"] = tags
        return Service(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

