# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FirewallPolicy(pulumi.CustomResource):
    custom_block_response_body: pulumi.Output[str]
    """
    If a `custom_rule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
    """
    custom_block_response_status_code: pulumi.Output[float]
    """
    If a `custom_rule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
    """
    custom_rules: pulumi.Output[list]
    """
    One or more `custom_rule` blocks as defined below.

      * `action` (`str`)
      * `enabled` (`bool`) - Is the policy a enabled state or disabled state. Defaults to `true`.
      * `matchConditions` (`list`)
        * `matchValues` (`list`)
        * `matchVariable` (`str`)
        * `negationCondition` (`bool`)
        * `operator` (`str`)
        * `selector` (`str`)
        * `transforms` (`list`)

      * `name` (`str`) - The name of the policy. Changing this forces a new resource to be created.
      * `priority` (`float`)
      * `rateLimitDurationInMinutes` (`float`)
      * `rateLimitThreshold` (`float`)
      * `type` (`str`)
    """
    enabled: pulumi.Output[bool]
    """
    Is the policy a enabled state or disabled state. Defaults to `true`.
    """
    frontend_endpoint_ids: pulumi.Output[list]
    """
    the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    managed_rules: pulumi.Output[list]
    """
    One or more `managed_rule` blocks as defined below.

      * `exclusions` (`list`)
        * `matchVariable` (`str`)
        * `operator` (`str`)
        * `selector` (`str`)

      * `overrides` (`list`)
        * `exclusions` (`list`)
          * `matchVariable` (`str`)
          * `operator` (`str`)
          * `selector` (`str`)

        * `ruleGroupName` (`str`)
        * `rules` (`list`)
          * `action` (`str`)
          * `enabled` (`bool`) - Is the policy a enabled state or disabled state. Defaults to `true`.
          * `exclusions` (`list`)
            * `matchVariable` (`str`)
            * `operator` (`str`)
            * `selector` (`str`)

          * `rule_id` (`str`)

      * `type` (`str`)
      * `version` (`str`)
    """
    mode: pulumi.Output[str]
    """
    The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
    """
    name: pulumi.Output[str]
    """
    The name of the policy. Changing this forces a new resource to be created.
    """
    redirect_url: pulumi.Output[str]
    """
    If action type is redirect, this field represents redirect URL for the client.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the Web Application Firewall Policy.
    """
    def __init__(__self__, resource_name, opts=None, custom_block_response_body=None, custom_block_response_status_code=None, custom_rules=None, enabled=None, managed_rules=None, mode=None, name=None, redirect_url=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Front Door Web Application Firewall Policy instance.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/frontdoor_firewall_policy.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_block_response_body: If a `custom_rule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
        :param pulumi.Input[float] custom_block_response_status_code: If a `custom_rule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
        :param pulumi.Input[list] custom_rules: One or more `custom_rule` blocks as defined below.
        :param pulumi.Input[bool] enabled: Is the policy a enabled state or disabled state. Defaults to `true`.
        :param pulumi.Input[list] managed_rules: One or more `managed_rule` blocks as defined below.
        :param pulumi.Input[str] mode: The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] redirect_url: If action type is redirect, this field represents redirect URL for the client.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Web Application Firewall Policy.

        The **custom_rules** object supports the following:

          * `action` (`pulumi.Input[str]`)
          * `enabled` (`pulumi.Input[bool]`) - Is the policy a enabled state or disabled state. Defaults to `true`.
          * `matchConditions` (`pulumi.Input[list]`)
            * `matchValues` (`pulumi.Input[list]`)
            * `matchVariable` (`pulumi.Input[str]`)
            * `negationCondition` (`pulumi.Input[bool]`)
            * `operator` (`pulumi.Input[str]`)
            * `selector` (`pulumi.Input[str]`)
            * `transforms` (`pulumi.Input[list]`)

          * `name` (`pulumi.Input[str]`) - The name of the policy. Changing this forces a new resource to be created.
          * `priority` (`pulumi.Input[float]`)
          * `rateLimitDurationInMinutes` (`pulumi.Input[float]`)
          * `rateLimitThreshold` (`pulumi.Input[float]`)
          * `type` (`pulumi.Input[str]`)

        The **managed_rules** object supports the following:

          * `exclusions` (`pulumi.Input[list]`)
            * `matchVariable` (`pulumi.Input[str]`)
            * `operator` (`pulumi.Input[str]`)
            * `selector` (`pulumi.Input[str]`)

          * `overrides` (`pulumi.Input[list]`)
            * `exclusions` (`pulumi.Input[list]`)
              * `matchVariable` (`pulumi.Input[str]`)
              * `operator` (`pulumi.Input[str]`)
              * `selector` (`pulumi.Input[str]`)

            * `ruleGroupName` (`pulumi.Input[str]`)
            * `rules` (`pulumi.Input[list]`)
              * `action` (`pulumi.Input[str]`)
              * `enabled` (`pulumi.Input[bool]`) - Is the policy a enabled state or disabled state. Defaults to `true`.
              * `exclusions` (`pulumi.Input[list]`)
                * `matchVariable` (`pulumi.Input[str]`)
                * `operator` (`pulumi.Input[str]`)
                * `selector` (`pulumi.Input[str]`)

              * `rule_id` (`pulumi.Input[str]`)

          * `type` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)
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

            __props__['custom_block_response_body'] = custom_block_response_body
            __props__['custom_block_response_status_code'] = custom_block_response_status_code
            __props__['custom_rules'] = custom_rules
            __props__['enabled'] = enabled
            __props__['managed_rules'] = managed_rules
            __props__['mode'] = mode
            __props__['name'] = name
            __props__['redirect_url'] = redirect_url
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['frontend_endpoint_ids'] = None
            __props__['location'] = None
        super(FirewallPolicy, __self__).__init__(
            'azure:frontdoor/firewallPolicy:FirewallPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, custom_block_response_body=None, custom_block_response_status_code=None, custom_rules=None, enabled=None, frontend_endpoint_ids=None, location=None, managed_rules=None, mode=None, name=None, redirect_url=None, resource_group_name=None, tags=None):
        """
        Get an existing FirewallPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_block_response_body: If a `custom_rule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
        :param pulumi.Input[float] custom_block_response_status_code: If a `custom_rule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
        :param pulumi.Input[list] custom_rules: One or more `custom_rule` blocks as defined below.
        :param pulumi.Input[bool] enabled: Is the policy a enabled state or disabled state. Defaults to `true`.
        :param pulumi.Input[list] frontend_endpoint_ids: the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[list] managed_rules: One or more `managed_rule` blocks as defined below.
        :param pulumi.Input[str] mode: The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] redirect_url: If action type is redirect, this field represents redirect URL for the client.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Web Application Firewall Policy.

        The **custom_rules** object supports the following:

          * `action` (`pulumi.Input[str]`)
          * `enabled` (`pulumi.Input[bool]`) - Is the policy a enabled state or disabled state. Defaults to `true`.
          * `matchConditions` (`pulumi.Input[list]`)
            * `matchValues` (`pulumi.Input[list]`)
            * `matchVariable` (`pulumi.Input[str]`)
            * `negationCondition` (`pulumi.Input[bool]`)
            * `operator` (`pulumi.Input[str]`)
            * `selector` (`pulumi.Input[str]`)
            * `transforms` (`pulumi.Input[list]`)

          * `name` (`pulumi.Input[str]`) - The name of the policy. Changing this forces a new resource to be created.
          * `priority` (`pulumi.Input[float]`)
          * `rateLimitDurationInMinutes` (`pulumi.Input[float]`)
          * `rateLimitThreshold` (`pulumi.Input[float]`)
          * `type` (`pulumi.Input[str]`)

        The **managed_rules** object supports the following:

          * `exclusions` (`pulumi.Input[list]`)
            * `matchVariable` (`pulumi.Input[str]`)
            * `operator` (`pulumi.Input[str]`)
            * `selector` (`pulumi.Input[str]`)

          * `overrides` (`pulumi.Input[list]`)
            * `exclusions` (`pulumi.Input[list]`)
              * `matchVariable` (`pulumi.Input[str]`)
              * `operator` (`pulumi.Input[str]`)
              * `selector` (`pulumi.Input[str]`)

            * `ruleGroupName` (`pulumi.Input[str]`)
            * `rules` (`pulumi.Input[list]`)
              * `action` (`pulumi.Input[str]`)
              * `enabled` (`pulumi.Input[bool]`) - Is the policy a enabled state or disabled state. Defaults to `true`.
              * `exclusions` (`pulumi.Input[list]`)
                * `matchVariable` (`pulumi.Input[str]`)
                * `operator` (`pulumi.Input[str]`)
                * `selector` (`pulumi.Input[str]`)

              * `rule_id` (`pulumi.Input[str]`)

          * `type` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_block_response_body"] = custom_block_response_body
        __props__["custom_block_response_status_code"] = custom_block_response_status_code
        __props__["custom_rules"] = custom_rules
        __props__["enabled"] = enabled
        __props__["frontend_endpoint_ids"] = frontend_endpoint_ids
        __props__["location"] = location
        __props__["managed_rules"] = managed_rules
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["redirect_url"] = redirect_url
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return FirewallPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

