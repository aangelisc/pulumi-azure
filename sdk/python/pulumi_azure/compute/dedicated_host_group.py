# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DedicatedHostGroupArgs', 'DedicatedHostGroup']

@pulumi.input_type
class DedicatedHostGroupArgs:
    def __init__(__self__, *,
                 platform_fault_domain_count: pulumi.Input[int],
                 resource_group_name: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DedicatedHostGroup resource.
        :param pulumi.Input[int] platform_fault_domain_count: The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Input[int]:
        """
        The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "platform_fault_domain_count", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[str]]:
        """
        A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zones", value)


@pulumi.input_type
class _DedicatedHostGroupState:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DedicatedHostGroup resources.
        :param pulumi.Input[str] location: The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_fault_domain_count is not None:
            pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_fault_domain_count", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zones(self) -> Optional[pulumi.Input[str]]:
        """
        A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

    @zones.setter
    def zones(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zones", value)


class DedicatedHostGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manage a Dedicated Host Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_dedicated_host_group = azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            platform_fault_domain_count=1)
        ```

        ## Import

        Dedicated Host Group can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/dedicatedHostGroup:DedicatedHostGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/hostGroups/group1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedHostGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a Dedicated Host Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_dedicated_host_group = azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            platform_fault_domain_count=1)
        ```

        ## Import

        Dedicated Host Group can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/dedicatedHostGroup:DedicatedHostGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/hostGroups/group1
        ```

        :param str resource_name: The name of the resource.
        :param DedicatedHostGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedHostGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[str]] = None,
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
            __props__ = DedicatedHostGroupArgs.__new__(DedicatedHostGroupArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if platform_fault_domain_count is None and not opts.urn:
                raise TypeError("Missing required property 'platform_fault_domain_count'")
            __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zones"] = zones
        super(DedicatedHostGroup, __self__).__init__(
            'azure:compute/dedicatedHostGroup:DedicatedHostGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            zones: Optional[pulumi.Input[str]] = None) -> 'DedicatedHostGroup':
        """
        Get an existing DedicatedHostGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zones: A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedHostGroupState.__new__(_DedicatedHostGroupState)

        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zones"] = zones
        return DedicatedHostGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Output[int]:
        """
        The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[str]]:
        """
        A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "zones")

