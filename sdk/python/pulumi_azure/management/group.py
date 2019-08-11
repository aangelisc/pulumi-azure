# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Group(pulumi.CustomResource):
    display_name: pulumi.Output[str]
    """
    A friendly name for this Management Group. If not specified, this'll be the same as the `group_id`.
    """
    group_id: pulumi.Output[str]
    """
    The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
    """
    parent_management_group_id: pulumi.Output[str]
    """
    The ID of the Parent Management Group. Changing this forces a new resource to be created.
    """
    subscription_ids: pulumi.Output[list]
    """
    A list of Subscription GUIDs which should be assigned to the Management Group.
    """
    def __init__(__self__, resource_name, opts=None, display_name=None, group_id=None, parent_management_group_id=None, subscription_ids=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Management Group.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: A friendly name for this Management Group. If not specified, this'll be the same as the `group_id`.
        :param pulumi.Input[str] group_id: The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
        :param pulumi.Input[str] parent_management_group_id: The ID of the Parent Management Group. Changing this forces a new resource to be created.
        :param pulumi.Input[list] subscription_ids: A list of Subscription GUIDs which should be assigned to the Management Group.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/management_group.html.markdown.
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

            __props__['display_name'] = display_name
            __props__['group_id'] = group_id
            __props__['parent_management_group_id'] = parent_management_group_id
            __props__['subscription_ids'] = subscription_ids
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:managementgroups/managementGroup:ManagementGroup")])
        opts = alias_opts if opts is None else opts.merge(alias_opts)
        super(Group, __self__).__init__(
            'azure:management/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, display_name=None, group_id=None, parent_management_group_id=None, subscription_ids=None):
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: A friendly name for this Management Group. If not specified, this'll be the same as the `group_id`.
        :param pulumi.Input[str] group_id: The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
        :param pulumi.Input[str] parent_management_group_id: The ID of the Parent Management Group. Changing this forces a new resource to be created.
        :param pulumi.Input[list] subscription_ids: A list of Subscription GUIDs which should be assigned to the Management Group.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/management_group.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["display_name"] = display_name
        __props__["group_id"] = group_id
        __props__["parent_management_group_id"] = parent_management_group_id
        __props__["subscription_ids"] = subscription_ids
        return Group(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

