# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ActionHttp(pulumi.CustomResource):
    body: pulumi.Output[str]
    """
    Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
    """
    headers: pulumi.Output[dict]
    """
    Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
    """
    logic_app_id: pulumi.Output[str]
    """
    Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
    """
    method: pulumi.Output[str]
    """
    Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
    """
    uri: pulumi.Output[str]
    """
    Specifies the URI which will be called when this HTTP Action is triggered.
    """
    def __init__(__self__, resource_name, opts=None, body=None, headers=None, logic_app_id=None, method=None, name=None, uri=None, __name__=None, __opts__=None):
        """
        Manages an HTTP Action within a Logic App Workflow
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[dict] headers: Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
        :param pulumi.Input[str] logic_app_id: Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] method: Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
        :param pulumi.Input[str] name: Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[str] uri: Specifies the URI which will be called when this HTTP Action is triggered.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/logic_app_action_http.html.markdown.
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

        __props__['body'] = body

        __props__['headers'] = headers

        if logic_app_id is None:
            raise TypeError("Missing required property 'logic_app_id'")
        __props__['logic_app_id'] = logic_app_id

        if method is None:
            raise TypeError("Missing required property 'method'")
        __props__['method'] = method

        __props__['name'] = name

        if uri is None:
            raise TypeError("Missing required property 'uri'")
        __props__['uri'] = uri

        super(ActionHttp, __self__).__init__(
            'azure:logicapps/actionHttp:ActionHttp',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

