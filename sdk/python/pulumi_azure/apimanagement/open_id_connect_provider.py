# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OpenIdConnectProviderArgs', 'OpenIdConnectProvider']

@pulumi.input_type
class OpenIdConnectProviderArgs:
    def __init__(__self__, *,
                 api_management_name: pulumi.Input[str],
                 client_id: pulumi.Input[str],
                 client_secret: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 metadata_endpoint: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OpenIdConnectProvider resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID used for the Client Application.
        :param pulumi.Input[str] client_secret: The Client Secret used for the Client Application.
        :param pulumi.Input[str] display_name: A user-friendly name for this OpenID Connect Provider.
        :param pulumi.Input[str] metadata_endpoint: The URI of the Metadata endpoint.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: A description of this OpenID Connect Provider.
        :param pulumi.Input[str] name: the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "api_management_name", api_management_name)
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "client_secret", client_secret)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "metadata_endpoint", metadata_endpoint)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Input[str]:
        """
        The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        The Client ID used for the Client Application.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Input[str]:
        """
        The Client Secret used for the Client Application.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        A user-friendly name for this OpenID Connect Provider.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="metadataEndpoint")
    def metadata_endpoint(self) -> pulumi.Input[str]:
        """
        The URI of the Metadata endpoint.
        """
        return pulumi.get(self, "metadata_endpoint")

    @metadata_endpoint.setter
    def metadata_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "metadata_endpoint", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of this OpenID Connect Provider.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OpenIdConnectProviderState:
    def __init__(__self__, *,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 metadata_endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OpenIdConnectProvider resources.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID used for the Client Application.
        :param pulumi.Input[str] client_secret: The Client Secret used for the Client Application.
        :param pulumi.Input[str] description: A description of this OpenID Connect Provider.
        :param pulumi.Input[str] display_name: A user-friendly name for this OpenID Connect Provider.
        :param pulumi.Input[str] metadata_endpoint: The URI of the Metadata endpoint.
        :param pulumi.Input[str] name: the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        if api_management_name is not None:
            pulumi.set(__self__, "api_management_name", api_management_name)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if metadata_endpoint is not None:
            pulumi.set(__self__, "metadata_endpoint", metadata_endpoint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Client ID used for the Client Application.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The Client Secret used for the Client Application.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of this OpenID Connect Provider.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-friendly name for this OpenID Connect Provider.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="metadataEndpoint")
    def metadata_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the Metadata endpoint.
        """
        return pulumi.get(self, "metadata_endpoint")

    @metadata_endpoint.setter
    def metadata_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)


class OpenIdConnectProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 metadata_endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an OpenID Connect Provider within a API Management Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="My Company",
            publisher_email="company@exmaple.com",
            sku_name="Developer_1")
        example_open_id_connect_provider = azure.apimanagement.OpenIdConnectProvider("exampleOpenIdConnectProvider",
            api_management_name=example_service.name,
            resource_group_name=example_resource_group.name,
            client_id="00001111-2222-3333-4444-555566667777",
            display_name="Example Provider",
            metadata_endpoint="https://example.com/example")
        ```

        ## Import

        API Management OpenID Connect Providers can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/openidConnectProviders/provider1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID used for the Client Application.
        :param pulumi.Input[str] client_secret: The Client Secret used for the Client Application.
        :param pulumi.Input[str] description: A description of this OpenID Connect Provider.
        :param pulumi.Input[str] display_name: A user-friendly name for this OpenID Connect Provider.
        :param pulumi.Input[str] metadata_endpoint: The URI of the Metadata endpoint.
        :param pulumi.Input[str] name: the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OpenIdConnectProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an OpenID Connect Provider within a API Management Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="My Company",
            publisher_email="company@exmaple.com",
            sku_name="Developer_1")
        example_open_id_connect_provider = azure.apimanagement.OpenIdConnectProvider("exampleOpenIdConnectProvider",
            api_management_name=example_service.name,
            resource_group_name=example_resource_group.name,
            client_id="00001111-2222-3333-4444-555566667777",
            display_name="Example Provider",
            metadata_endpoint="https://example.com/example")
        ```

        ## Import

        API Management OpenID Connect Providers can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/openidConnectProviders/provider1
        ```

        :param str resource_name: The name of the resource.
        :param OpenIdConnectProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OpenIdConnectProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 metadata_endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OpenIdConnectProviderArgs.__new__(OpenIdConnectProviderArgs)

            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__.__dict__["api_management_name"] = api_management_name
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if client_secret is None and not opts.urn:
                raise TypeError("Missing required property 'client_secret'")
            __props__.__dict__["client_secret"] = client_secret
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if metadata_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'metadata_endpoint'")
            __props__.__dict__["metadata_endpoint"] = metadata_endpoint
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
        super(OpenIdConnectProvider, __self__).__init__(
            'azure:apimanagement/openIdConnectProvider:OpenIdConnectProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            metadata_endpoint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'OpenIdConnectProvider':
        """
        Get an existing OpenIdConnectProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] client_id: The Client ID used for the Client Application.
        :param pulumi.Input[str] client_secret: The Client Secret used for the Client Application.
        :param pulumi.Input[str] description: A description of this OpenID Connect Provider.
        :param pulumi.Input[str] display_name: A user-friendly name for this OpenID Connect Provider.
        :param pulumi.Input[str] metadata_endpoint: The URI of the Metadata endpoint.
        :param pulumi.Input[str] name: the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OpenIdConnectProviderState.__new__(_OpenIdConnectProviderState)

        __props__.__dict__["api_management_name"] = api_management_name
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["metadata_endpoint"] = metadata_endpoint
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        return OpenIdConnectProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The Client ID used for the Client Application.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        The Client Secret used for the Client Application.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of this OpenID Connect Provider.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A user-friendly name for this OpenID Connect Provider.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="metadataEndpoint")
    def metadata_endpoint(self) -> pulumi.Output[str]:
        """
        The URI of the Metadata endpoint.
        """
        return pulumi.get(self, "metadata_endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        the Name of the OpenID Connect Provider which should be created within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

