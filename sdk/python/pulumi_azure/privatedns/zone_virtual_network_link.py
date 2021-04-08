# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ZoneVirtualNetworkLinkArgs', 'ZoneVirtualNetworkLink']

@pulumi.input_type
class ZoneVirtualNetworkLinkArgs:
    def __init__(__self__, *,
                 private_dns_zone_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 virtual_network_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 registration_enabled: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ZoneVirtualNetworkLink resource.
        :param pulumi.Input[str] private_dns_zone_name: The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_network_id: The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] registration_enabled: Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "private_dns_zone_name", private_dns_zone_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "virtual_network_id", virtual_network_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if registration_enabled is not None:
            pulumi.set(__self__, "registration_enabled", registration_enabled)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="privateDnsZoneName")
    def private_dns_zone_name(self) -> pulumi.Input[str]:
        """
        The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "private_dns_zone_name")

    @private_dns_zone_name.setter
    def private_dns_zone_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_dns_zone_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_id")

    @virtual_network_id.setter
    def virtual_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_network_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="registrationEnabled")
    def registration_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        """
        return pulumi.get(self, "registration_enabled")

    @registration_enabled.setter
    def registration_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "registration_enabled", value)

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
class _ZoneVirtualNetworkLinkState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 private_dns_zone_name: Optional[pulumi.Input[str]] = None,
                 registration_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ZoneVirtualNetworkLink resources.
        :param pulumi.Input[str] name: The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_dns_zone_name: The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        :param pulumi.Input[bool] registration_enabled: Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_network_id: The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_dns_zone_name is not None:
            pulumi.set(__self__, "private_dns_zone_name", private_dns_zone_name)
        if registration_enabled is not None:
            pulumi.set(__self__, "registration_enabled", registration_enabled)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_network_id is not None:
            pulumi.set(__self__, "virtual_network_id", virtual_network_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateDnsZoneName")
    def private_dns_zone_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "private_dns_zone_name")

    @private_dns_zone_name.setter
    def private_dns_zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_dns_zone_name", value)

    @property
    @pulumi.getter(name="registrationEnabled")
    def registration_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        """
        return pulumi.get(self, "registration_enabled")

    @registration_enabled.setter
    def registration_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "registration_enabled", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
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
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_id")

    @virtual_network_id.setter
    def virtual_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_network_id", value)


class ZoneVirtualNetworkLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_dns_zone_name: Optional[pulumi.Input[str]] = None,
                 registration_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Enables you to manage Private DNS zone Virtual Network Links. These Links enable DNS resolution and registration inside Azure Virtual Networks using Azure Private DNS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_zone = azure.privatedns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        example_zone_virtual_network_link = azure.privatedns.ZoneVirtualNetworkLink("exampleZoneVirtualNetworkLink",
            resource_group_name=example_resource_group.name,
            private_dns_zone_name=example_zone.name,
            virtual_network_id=azurerm_virtual_network["example"]["id"])
        ```

        ## Import

        Private DNS Zone Virtual Network Links can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink link1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1.com/virtualNetworkLinks/myVnetLink1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_dns_zone_name: The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        :param pulumi.Input[bool] registration_enabled: Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_network_id: The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneVirtualNetworkLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enables you to manage Private DNS zone Virtual Network Links. These Links enable DNS resolution and registration inside Azure Virtual Networks using Azure Private DNS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_zone = azure.privatedns.Zone("exampleZone", resource_group_name=example_resource_group.name)
        example_zone_virtual_network_link = azure.privatedns.ZoneVirtualNetworkLink("exampleZoneVirtualNetworkLink",
            resource_group_name=example_resource_group.name,
            private_dns_zone_name=example_zone.name,
            virtual_network_id=azurerm_virtual_network["example"]["id"])
        ```

        ## Import

        Private DNS Zone Virtual Network Links can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink link1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1.com/virtualNetworkLinks/myVnetLink1
        ```

        :param str resource_name: The name of the resource.
        :param ZoneVirtualNetworkLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneVirtualNetworkLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_dns_zone_name: Optional[pulumi.Input[str]] = None,
                 registration_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ZoneVirtualNetworkLinkArgs.__new__(ZoneVirtualNetworkLinkArgs)

            __props__.__dict__["name"] = name
            if private_dns_zone_name is None and not opts.urn:
                raise TypeError("Missing required property 'private_dns_zone_name'")
            __props__.__dict__["private_dns_zone_name"] = private_dns_zone_name
            __props__.__dict__["registration_enabled"] = registration_enabled
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if virtual_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_network_id'")
            __props__.__dict__["virtual_network_id"] = virtual_network_id
        super(ZoneVirtualNetworkLink, __self__).__init__(
            'azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_dns_zone_name: Optional[pulumi.Input[str]] = None,
            registration_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            virtual_network_id: Optional[pulumi.Input[str]] = None) -> 'ZoneVirtualNetworkLink':
        """
        Get an existing ZoneVirtualNetworkLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_dns_zone_name: The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        :param pulumi.Input[bool] registration_enabled: Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_network_id: The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneVirtualNetworkLinkState.__new__(_ZoneVirtualNetworkLinkState)

        __props__.__dict__["name"] = name
        __props__.__dict__["private_dns_zone_name"] = private_dns_zone_name
        __props__.__dict__["registration_enabled"] = registration_enabled
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["virtual_network_id"] = virtual_network_id
        return ZoneVirtualNetworkLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateDnsZoneName")
    def private_dns_zone_name(self) -> pulumi.Output[str]:
        """
        The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "private_dns_zone_name")

    @property
    @pulumi.getter(name="registrationEnabled")
    def registration_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
        """
        return pulumi.get(self, "registration_enabled")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
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
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_id")

