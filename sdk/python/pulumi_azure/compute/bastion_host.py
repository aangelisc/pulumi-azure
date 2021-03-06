# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BastionHostArgs', 'BastionHost']

@pulumi.input_type
class BastionHostArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 ip_configuration: Optional[pulumi.Input['BastionHostIpConfigurationArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a BastionHost resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bastion Host.
        :param pulumi.Input['BastionHostIpConfigurationArgs'] ip_configuration: A `ip_configuration` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        :param pulumi.Input[str] name: Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if ip_configuration is not None:
            pulumi.set(__self__, "ip_configuration", ip_configuration)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Bastion Host.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="ipConfiguration")
    def ip_configuration(self) -> Optional[pulumi.Input['BastionHostIpConfigurationArgs']]:
        """
        A `ip_configuration` block as defined below.
        """
        return pulumi.get(self, "ip_configuration")

    @ip_configuration.setter
    def ip_configuration(self, value: Optional[pulumi.Input['BastionHostIpConfigurationArgs']]):
        pulumi.set(self, "ip_configuration", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
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


@pulumi.input_type
class _BastionHostState:
    def __init__(__self__, *,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 ip_configuration: Optional[pulumi.Input['BastionHostIpConfigurationArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering BastionHost resources.
        :param pulumi.Input[str] dns_name: The FQDN for the Bastion Host.
        :param pulumi.Input['BastionHostIpConfigurationArgs'] ip_configuration: A `ip_configuration` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        :param pulumi.Input[str] name: Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bastion Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if dns_name is not None:
            pulumi.set(__self__, "dns_name", dns_name)
        if ip_configuration is not None:
            pulumi.set(__self__, "ip_configuration", ip_configuration)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The FQDN for the Bastion Host.
        """
        return pulumi.get(self, "dns_name")

    @dns_name.setter
    def dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_name", value)

    @property
    @pulumi.getter(name="ipConfiguration")
    def ip_configuration(self) -> Optional[pulumi.Input['BastionHostIpConfigurationArgs']]:
        """
        A `ip_configuration` block as defined below.
        """
        return pulumi.get(self, "ip_configuration")

    @ip_configuration.setter
    def ip_configuration(self, value: Optional[pulumi.Input['BastionHostIpConfigurationArgs']]):
        pulumi.set(self, "ip_configuration", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Bastion Host.
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


class BastionHost(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_configuration: Optional[pulumi.Input[pulumi.InputType['BastionHostIpConfigurationArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a Bastion Host.

        ## Example Usage

        This example deploys an Azure Bastion Host Instance to a target virtual network.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["192.168.1.0/24"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["192.168.1.224/27"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_bastion_host = azure.compute.BastionHost("exampleBastionHost",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configuration=azure.compute.BastionHostIpConfigurationArgs(
                name="configuration",
                subnet_id=example_subnet.id,
                public_ip_address_id=example_public_ip.id,
            ))
        ```

        ## Import

        Bastion Hosts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/bastionHost:BastionHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BastionHostIpConfigurationArgs']] ip_configuration: A `ip_configuration` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        :param pulumi.Input[str] name: Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bastion Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BastionHostArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Bastion Host.

        ## Example Usage

        This example deploys an Azure Bastion Host Instance to a target virtual network.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["192.168.1.0/24"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["192.168.1.224/27"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_bastion_host = azure.compute.BastionHost("exampleBastionHost",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configuration=azure.compute.BastionHostIpConfigurationArgs(
                name="configuration",
                subnet_id=example_subnet.id,
                public_ip_address_id=example_public_ip.id,
            ))
        ```

        ## Import

        Bastion Hosts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/bastionHost:BastionHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
        ```

        :param str resource_name: The name of the resource.
        :param BastionHostArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BastionHostArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_configuration: Optional[pulumi.Input[pulumi.InputType['BastionHostIpConfigurationArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
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
            __props__ = BastionHostArgs.__new__(BastionHostArgs)

            __props__.__dict__["ip_configuration"] = ip_configuration
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["dns_name"] = None
        super(BastionHost, __self__).__init__(
            'azure:compute/bastionHost:BastionHost',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            ip_configuration: Optional[pulumi.Input[pulumi.InputType['BastionHostIpConfigurationArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'BastionHost':
        """
        Get an existing BastionHost resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_name: The FQDN for the Bastion Host.
        :param pulumi.Input[pulumi.InputType['BastionHostIpConfigurationArgs']] ip_configuration: A `ip_configuration` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        :param pulumi.Input[str] name: Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bastion Host.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BastionHostState.__new__(_BastionHostState)

        __props__.__dict__["dns_name"] = dns_name
        __props__.__dict__["ip_configuration"] = ip_configuration
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        return BastionHost(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        """
        The FQDN for the Bastion Host.
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="ipConfiguration")
    def ip_configuration(self) -> pulumi.Output[Optional['outputs.BastionHostIpConfiguration']]:
        """
        A `ip_configuration` block as defined below.
        """
        return pulumi.get(self, "ip_configuration")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Bastion Host.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

