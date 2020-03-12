# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RegistryWebook(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chart_push`, `chart_delete`
    """
    custom_headers: pulumi.Output[dict]
    """
    Custom headers that will be added to the webhook notifications request.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
    """
    registry_name: pulumi.Output[str]
    """
    The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
    """
    scope: pulumi.Output[str]
    """
    Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
    """
    service_uri: pulumi.Output[str]
    """
    Specifies the service URI for the Webhook to post notifications.
    """
    status: pulumi.Output[str]
    """
    Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
    """
    tags: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, actions=None, custom_headers=None, location=None, name=None, registry_name=None, resource_group_name=None, scope=None, service_uri=None, status=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Container Registry Webhook.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_registry_webhook.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chart_push`, `chart_delete`
        :param pulumi.Input[dict] custom_headers: Custom headers that will be added to the webhook notifications request.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[str] registry_name: The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope: Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
        :param pulumi.Input[str] service_uri: Specifies the service URI for the Webhook to post notifications.
        :param pulumi.Input[str] status: Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
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

            if actions is None:
                raise TypeError("Missing required property 'actions'")
            __props__['actions'] = actions
            __props__['custom_headers'] = custom_headers
            __props__['location'] = location
            __props__['name'] = name
            if registry_name is None:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scope'] = scope
            if service_uri is None:
                raise TypeError("Missing required property 'service_uri'")
            __props__['service_uri'] = service_uri
            __props__['status'] = status
            __props__['tags'] = tags
        super(RegistryWebook, __self__).__init__(
            'azure:containerservice/registryWebook:RegistryWebook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions=None, custom_headers=None, location=None, name=None, registry_name=None, resource_group_name=None, scope=None, service_uri=None, status=None, tags=None):
        """
        Get an existing RegistryWebook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chart_push`, `chart_delete`
        :param pulumi.Input[dict] custom_headers: Custom headers that will be added to the webhook notifications request.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[str] registry_name: The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope: Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
        :param pulumi.Input[str] service_uri: Specifies the service URI for the Webhook to post notifications.
        :param pulumi.Input[str] status: Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions"] = actions
        __props__["custom_headers"] = custom_headers
        __props__["location"] = location
        __props__["name"] = name
        __props__["registry_name"] = registry_name
        __props__["resource_group_name"] = resource_group_name
        __props__["scope"] = scope
        __props__["service_uri"] = service_uri
        __props__["status"] = status
        __props__["tags"] = tags
        return RegistryWebook(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

