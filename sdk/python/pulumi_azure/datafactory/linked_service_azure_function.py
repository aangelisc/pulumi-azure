# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LinkedServiceAzureFunctionArgs', 'LinkedServiceAzureFunction']

@pulumi.input_type
class LinkedServiceAzureFunctionArgs:
    def __init__(__self__, *,
                 data_factory_name: pulumi.Input[str],
                 key: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a LinkedServiceAzureFunction resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] key: The system key of the Azure Function.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] url: The url of the Azure Function.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
               factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        """
        pulumi.set(__self__, "data_factory_name", data_factory_name)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "url", url)
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if integration_runtime_name is not None:
            pulumi.set(__self__, "integration_runtime_name", integration_runtime_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Input[str]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The system key of the Azure Function.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The url of the Azure Function.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Data Factory Linked Service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> Optional[pulumi.Input[str]]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "integration_runtime_name")

    @integration_runtime_name.setter
    def integration_runtime_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_runtime_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
        factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of parameters to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class _LinkedServiceAzureFunctionState:
    def __init__(__self__, *,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LinkedServiceAzureFunction resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] key: The system key of the Azure Function.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
               factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] url: The url of the Azure Function.
        """
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if data_factory_name is not None:
            pulumi.set(__self__, "data_factory_name", data_factory_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if integration_runtime_name is not None:
            pulumi.set(__self__, "integration_runtime_name", integration_runtime_name)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Data Factory Linked Service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> Optional[pulumi.Input[str]]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "integration_runtime_name")

    @integration_runtime_name.setter
    def integration_runtime_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_runtime_name", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The system key of the Azure Function.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
        factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of parameters to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The url of the Azure Function.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class LinkedServiceAzureFunction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_function_app = example_resource_group.name.apply(lambda name: azure.appservice.get_function_app(name="test-azure-functions",
            resource_group_name=name))
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_azure_function = azure.datafactory.LinkedServiceAzureFunction("exampleLinkedServiceAzureFunction",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            url=example_function_app.apply(lambda example_function_app: f"https://{example_function_app.default_hostname}"),
            key="foo")
        ```

        ## Import

        Data Factory Linked Service's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] key: The system key of the Azure Function.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
               factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] url: The url of the Azure Function.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkedServiceAzureFunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_function_app = example_resource_group.name.apply(lambda name: azure.appservice.get_function_app(name="test-azure-functions",
            resource_group_name=name))
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_azure_function = azure.datafactory.LinkedServiceAzureFunction("exampleLinkedServiceAzureFunction",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            url=example_function_app.apply(lambda example_function_app: f"https://{example_function_app.default_hostname}"),
            key="foo")
        ```

        ## Import

        Data Factory Linked Service's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
        ```

        :param str resource_name: The name of the resource.
        :param LinkedServiceAzureFunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkedServiceAzureFunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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
            __props__ = LinkedServiceAzureFunctionArgs.__new__(LinkedServiceAzureFunctionArgs)

            __props__.__dict__["additional_properties"] = additional_properties
            __props__.__dict__["annotations"] = annotations
            if data_factory_name is None and not opts.urn:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__.__dict__["data_factory_name"] = data_factory_name
            __props__.__dict__["description"] = description
            __props__.__dict__["integration_runtime_name"] = integration_runtime_name
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
        super(LinkedServiceAzureFunction, __self__).__init__(
            'azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            integration_runtime_name: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'LinkedServiceAzureFunction':
        """
        Get an existing LinkedServiceAzureFunction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] key: The system key of the Azure Function.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
               factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        :param pulumi.Input[str] url: The url of the Azure Function.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LinkedServiceAzureFunctionState.__new__(_LinkedServiceAzureFunctionState)

        __props__.__dict__["additional_properties"] = additional_properties
        __props__.__dict__["annotations"] = annotations
        __props__.__dict__["data_factory_name"] = data_factory_name
        __props__.__dict__["description"] = description
        __props__.__dict__["integration_runtime_name"] = integration_runtime_name
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["url"] = url
        return LinkedServiceAzureFunction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Output[str]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the Data Factory Linked Service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> pulumi.Output[Optional[str]]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "integration_runtime_name")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The system key of the Azure Function.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
        factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of parameters to associate with the Data Factory Linked Service.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The url of the Azure Function.
        """
        return pulumi.get(self, "url")

