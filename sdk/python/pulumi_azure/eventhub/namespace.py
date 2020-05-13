# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("azure.eventhub.Namespace has been deprecated in favor of azure.servicebus.Namespace", DeprecationWarning)
class Namespace(pulumi.CustomResource):
    capacity: pulumi.Output[float]
    """
    Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
    """
    default_primary_connection_string: pulumi.Output[str]
    """
    The primary connection string for the authorization
    rule `RootManageSharedAccessKey`.
    """
    default_primary_key: pulumi.Output[str]
    """
    The primary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    default_secondary_connection_string: pulumi.Output[str]
    """
    The secondary connection string for the
    authorization rule `RootManageSharedAccessKey`.
    """
    default_secondary_key: pulumi.Output[str]
    """
    The secondary access key for the authorization rule `RootManageSharedAccessKey`.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Namespace resource . Changing this forces a
    new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the namespace.
    """
    sku: pulumi.Output[str]
    """
    Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    zone_redundant: pulumi.Output[bool]
    """
    Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
    """
    warnings.warn("azure.eventhub.Namespace has been deprecated in favor of azure.servicebus.Namespace", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, capacity=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, zone_redundant=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a ServiceBus Namespace.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        ```


        Deprecated: azure.eventhub.Namespace has been deprecated in favor of azure.servicebus.Namespace

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] capacity: Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Namespace resource . Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the namespace.
        :param pulumi.Input[str] sku: Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] zone_redundant: Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
        """
        pulumi.log.warn("Namespace is deprecated: azure.eventhub.Namespace has been deprecated in favor of azure.servicebus.Namespace")
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

            __props__['capacity'] = capacity
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zone_redundant'] = zone_redundant
            __props__['default_primary_connection_string'] = None
            __props__['default_primary_key'] = None
            __props__['default_secondary_connection_string'] = None
            __props__['default_secondary_key'] = None
        super(Namespace, __self__).__init__(
            'azure:eventhub/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, capacity=None, default_primary_connection_string=None, default_primary_key=None, default_secondary_connection_string=None, default_secondary_key=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, zone_redundant=None):
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] capacity: Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
        :param pulumi.Input[str] default_primary_connection_string: The primary connection string for the authorization
               rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_primary_key: The primary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_secondary_connection_string: The secondary connection string for the
               authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_secondary_key: The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Namespace resource . Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the namespace.
        :param pulumi.Input[str] sku: Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] zone_redundant: Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["capacity"] = capacity
        __props__["default_primary_connection_string"] = default_primary_connection_string
        __props__["default_primary_key"] = default_primary_key
        __props__["default_secondary_connection_string"] = default_secondary_connection_string
        __props__["default_secondary_key"] = default_secondary_key
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["zone_redundant"] = zone_redundant
        return Namespace(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

