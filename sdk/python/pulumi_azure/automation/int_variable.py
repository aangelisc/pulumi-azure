# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IntVariable(pulumi.CustomResource):
    automation_account_name: pulumi.Output[str]
    """
    The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
    """
    description: pulumi.Output[str]
    """
    The description of the Automation Variable.
    """
    encrypted: pulumi.Output[bool]
    """
    Specifies if the Automation Variable is encrypted. Defaults to `false`.
    """
    name: pulumi.Output[str]
    """
    The name of the Automation Variable. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
    """
    value: pulumi.Output[float]
    """
    The value of the Automation Variable as a `integer`.
    """
    def __init__(__self__, resource_name, opts=None, automation_account_name=None, description=None, encrypted=None, name=None, resource_group_name=None, value=None, __name__=None, __opts__=None):
        """
        Manages a integer variable in Azure Automation
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The description of the Automation Variable.
        :param pulumi.Input[bool] encrypted: Specifies if the Automation Variable is encrypted. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the Automation Variable. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
        :param pulumi.Input[float] value: The value of the Automation Variable as a `integer`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_variable_int.html.markdown.
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

        if automation_account_name is None:
            raise TypeError("Missing required property 'automation_account_name'")
        __props__['automation_account_name'] = automation_account_name

        __props__['description'] = description

        __props__['encrypted'] = encrypted

        __props__['name'] = name

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['value'] = value

        super(IntVariable, __self__).__init__(
            'azure:automation/intVariable:IntVariable',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

