# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 data_location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[str] data_location: The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        :param pulumi.Input[str] name: The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Communication Service.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if data_location is not None:
            pulumi.set(__self__, "data_location", data_location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        """
        return pulumi.get(self, "data_location")

    @data_location.setter
    def data_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the Communication Service.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ServiceState:
    def __init__(__self__, *,
                 data_location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Service resources.
        :param pulumi.Input[str] data_location: The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        :param pulumi.Input[str] name: The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Communication Service.
        """
        if data_location is not None:
            pulumi.set(__self__, "data_location", data_location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        """
        return pulumi.get(self, "data_location")

    @data_location.setter
    def data_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the Communication Service.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a Communication Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.communication.Service("exampleService",
            resource_group_name=example_resource_group.name,
            data_location="United States")
        ```

        ## Import

        Communication Services can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:communication/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/CommunicationServices/communicationService1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_location: The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        :param pulumi.Input[str] name: The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Communication Service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Communication Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.communication.Service("exampleService",
            resource_group_name=example_resource_group.name,
            data_location="United States")
        ```

        ## Import

        Communication Services can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:communication/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/CommunicationServices/communicationService1
        ```

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["data_location"] = data_location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
        super(Service, __self__).__init__(
            'azure:communication/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_location: The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        :param pulumi.Input[str] name: The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Communication Service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceState.__new__(_ServiceState)

        __props__.__dict__["data_location"] = data_location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> pulumi.Output[Optional[str]]:
        """
        The location where the Communication service stores its data at rest. Possible values are `Asia Pacific`, `Australia`, `Europe`, `UK` and `United States`. Defaults to `United States`.
        """
        return pulumi.get(self, "data_location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Communication Service resource. Changing this forces a new Communication Service to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Communication Service.
        """
        return pulumi.get(self, "tags")
