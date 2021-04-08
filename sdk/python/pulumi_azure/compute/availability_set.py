# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AvailabilitySetArgs', 'AvailabilitySet']

@pulumi.input_type
class AvailabilitySetArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 platform_update_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AvailabilitySet resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] managed: Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        :param pulumi.Input[str] name: Specifies the name of the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_update_domain_count: Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed is not None:
            pulumi.set(__self__, "managed", managed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_fault_domain_count is not None:
            pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        if platform_update_domain_count is not None:
            pulumi.set(__self__, "platform_update_domain_count", platform_update_domain_count)
        if proximity_placement_group_id is not None:
            pulumi.set(__self__, "proximity_placement_group_id", proximity_placement_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def managed(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the availability set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_fault_domain_count", value)

    @property
    @pulumi.getter(name="platformUpdateDomainCount")
    def platform_update_domain_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_update_domain_count")

    @platform_update_domain_count.setter
    def platform_update_domain_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_update_domain_count", value)

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @proximity_placement_group_id.setter
    def proximity_placement_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proximity_placement_group_id", value)

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


@pulumi.input_type
class _AvailabilitySetState:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 platform_update_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AvailabilitySet resources.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] managed: Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        :param pulumi.Input[str] name: Specifies the name of the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_update_domain_count: Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed is not None:
            pulumi.set(__self__, "managed", managed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_fault_domain_count is not None:
            pulumi.set(__self__, "platform_fault_domain_count", platform_fault_domain_count)
        if platform_update_domain_count is not None:
            pulumi.set(__self__, "platform_update_domain_count", platform_update_domain_count)
        if proximity_placement_group_id is not None:
            pulumi.set(__self__, "proximity_placement_group_id", proximity_placement_group_id)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def managed(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the availability set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @platform_fault_domain_count.setter
    def platform_fault_domain_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_fault_domain_count", value)

    @property
    @pulumi.getter(name="platformUpdateDomainCount")
    def platform_update_domain_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_update_domain_count")

    @platform_update_domain_count.setter
    def platform_update_domain_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_update_domain_count", value)

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @proximity_placement_group_id.setter
    def proximity_placement_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proximity_placement_group_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
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


class AvailabilitySet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 platform_update_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Availability Set for Virtual Machines.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_availability_set = azure.compute.AvailabilitySet("exampleAvailabilitySet",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        Availability Sets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/availabilitySet:AvailabilitySet group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] managed: Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        :param pulumi.Input[str] name: Specifies the name of the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_update_domain_count: Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AvailabilitySetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Availability Set for Virtual Machines.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_availability_set = azure.compute.AvailabilitySet("exampleAvailabilitySet",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        Availability Sets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/availabilitySet:AvailabilitySet group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
        ```

        :param str resource_name: The name of the resource.
        :param AvailabilitySetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AvailabilitySetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
                 platform_update_domain_count: Optional[pulumi.Input[int]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = AvailabilitySetArgs.__new__(AvailabilitySetArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["managed"] = managed
            __props__.__dict__["name"] = name
            __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
            __props__.__dict__["platform_update_domain_count"] = platform_update_domain_count
            __props__.__dict__["proximity_placement_group_id"] = proximity_placement_group_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
        super(AvailabilitySet, __self__).__init__(
            'azure:compute/availabilitySet:AvailabilitySet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[str]] = None,
            managed: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_fault_domain_count: Optional[pulumi.Input[int]] = None,
            platform_update_domain_count: Optional[pulumi.Input[int]] = None,
            proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'AvailabilitySet':
        """
        Get an existing AvailabilitySet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] managed: Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        :param pulumi.Input[str] name: Specifies the name of the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_fault_domain_count: Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] platform_update_domain_count: Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] proximity_placement_group_id: The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AvailabilitySetState.__new__(_AvailabilitySetState)

        __props__.__dict__["location"] = location
        __props__.__dict__["managed"] = managed
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_fault_domain_count"] = platform_fault_domain_count
        __props__.__dict__["platform_update_domain_count"] = platform_update_domain_count
        __props__.__dict__["proximity_placement_group_id"] = proximity_placement_group_id
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        return AvailabilitySet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def managed(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the availability set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_fault_domain_count")

    @property
    @pulumi.getter(name="platformUpdateDomainCount")
    def platform_update_domain_count(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "platform_update_domain_count")

    @property
    @pulumi.getter(name="proximityPlacementGroupId")
    def proximity_placement_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

