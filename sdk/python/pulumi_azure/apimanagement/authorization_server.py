# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AuthorizationServer(pulumi.CustomResource):
    api_management_name: pulumi.Output[str]
    """
    The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
    """
    authorization_endpoint: pulumi.Output[str]
    """
    The OAUTH Authorization Endpoint.
    """
    authorization_methods: pulumi.Output[list]
    """
    The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
    """
    bearer_token_sending_methods: pulumi.Output[list]
    """
    The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
    """
    client_authentication_methods: pulumi.Output[list]
    """
    The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
    """
    client_id: pulumi.Output[str]
    """
    The Client/App ID registered with this Authorization Server.
    """
    client_registration_endpoint: pulumi.Output[str]
    """
    The URI of page where Client/App Registration is performed for this Authorization Server.
    """
    client_secret: pulumi.Output[str]
    """
    The Client/App Secret registered with this Authorization Server.
    """
    default_scope: pulumi.Output[str]
    """
    The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
    """
    description: pulumi.Output[str]
    """
    A description of the Authorization Server, which may contain HTML formatting tags.
    """
    display_name: pulumi.Output[str]
    """
    The user-friendly name of this Authorization Server.
    """
    grant_types: pulumi.Output[list]
    """
    Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
    """
    name: pulumi.Output[str]
    """
    The name of this Authorization Server. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
    """
    resource_owner_password: pulumi.Output[str]
    """
    The password associated with the Resource Owner.
    """
    resource_owner_username: pulumi.Output[str]
    """
    The username associated with the Resource Owner.
    """
    support_state: pulumi.Output[bool]
    """
    Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
    """
    token_body_parameters: pulumi.Output[list]
    token_endpoint: pulumi.Output[str]
    """
    The OAUTH Token Endpoint.
    """
    def __init__(__self__, resource_name, opts=None, api_management_name=None, authorization_endpoint=None, authorization_methods=None, bearer_token_sending_methods=None, client_authentication_methods=None, client_id=None, client_registration_endpoint=None, client_secret=None, default_scope=None, description=None, display_name=None, grant_types=None, name=None, resource_group_name=None, resource_owner_password=None, resource_owner_username=None, support_state=None, token_body_parameters=None, token_endpoint=None, __name__=None, __opts__=None):
        """
        Manages an Authorization Server within an API Management Service.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] authorization_endpoint: The OAUTH Authorization Endpoint.
        :param pulumi.Input[list] authorization_methods: The HTTP Verbs supported by the Authorization Endpoint. Possible values are `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT` and `TRACE`.
        :param pulumi.Input[list] bearer_token_sending_methods: The mechanism by which Access Tokens are passed to the API. Possible values are `authorizationHeader` and `query`.
        :param pulumi.Input[list] client_authentication_methods: The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are `Basic` and `Body`.
        :param pulumi.Input[str] client_id: The Client/App ID registered with this Authorization Server.
        :param pulumi.Input[str] client_registration_endpoint: The URI of page where Client/App Registration is performed for this Authorization Server.
        :param pulumi.Input[str] client_secret: The Client/App Secret registered with this Authorization Server.
        :param pulumi.Input[str] default_scope: The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
        :param pulumi.Input[str] description: A description of the Authorization Server, which may contain HTML formatting tags.
        :param pulumi.Input[str] display_name: The user-friendly name of this Authorization Server.
        :param pulumi.Input[list] grant_types: Form of Authorization Grants required when requesting an Access Token. Possible values are `authorizationCode`, `clientCredentials`, `implicit` and `resourceOwnerPassword`.
        :param pulumi.Input[str] name: The name of this Authorization Server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_owner_password: The password associated with the Resource Owner.
        :param pulumi.Input[str] resource_owner_username: The username associated with the Resource Owner.
        :param pulumi.Input[bool] support_state: Does this Authorization Server support State? If this is set to `true` the client may use the state parameter to raise protocol security.
        :param pulumi.Input[str] token_endpoint: The OAUTH Token Endpoint.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_authorization_server.html.markdown.
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

        if api_management_name is None:
            raise TypeError("Missing required property 'api_management_name'")
        __props__['api_management_name'] = api_management_name

        if authorization_endpoint is None:
            raise TypeError("Missing required property 'authorization_endpoint'")
        __props__['authorization_endpoint'] = authorization_endpoint

        if authorization_methods is None:
            raise TypeError("Missing required property 'authorization_methods'")
        __props__['authorization_methods'] = authorization_methods

        __props__['bearer_token_sending_methods'] = bearer_token_sending_methods

        __props__['client_authentication_methods'] = client_authentication_methods

        if client_id is None:
            raise TypeError("Missing required property 'client_id'")
        __props__['client_id'] = client_id

        if client_registration_endpoint is None:
            raise TypeError("Missing required property 'client_registration_endpoint'")
        __props__['client_registration_endpoint'] = client_registration_endpoint

        __props__['client_secret'] = client_secret

        __props__['default_scope'] = default_scope

        __props__['description'] = description

        if display_name is None:
            raise TypeError("Missing required property 'display_name'")
        __props__['display_name'] = display_name

        if grant_types is None:
            raise TypeError("Missing required property 'grant_types'")
        __props__['grant_types'] = grant_types

        __props__['name'] = name

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['resource_owner_password'] = resource_owner_password

        __props__['resource_owner_username'] = resource_owner_username

        __props__['support_state'] = support_state

        __props__['token_body_parameters'] = token_body_parameters

        __props__['token_endpoint'] = token_endpoint

        super(AuthorizationServer, __self__).__init__(
            'azure:apimanagement/authorizationServer:AuthorizationServer',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

