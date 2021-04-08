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

__all__ = ['ZoneArgs', 'Zone']

@pulumi.input_type
class ZoneArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 soa_record: Optional[pulumi.Input['ZoneSoaRecordArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Zone resource.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the DNS Zone. Must be a valid domain name.
        :param pulumi.Input['ZoneSoaRecordArgs'] soa_record: An `soa_record` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Record Set.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if soa_record is not None:
            pulumi.set(__self__, "soa_record", soa_record)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the DNS Zone. Must be a valid domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="soaRecord")
    def soa_record(self) -> Optional[pulumi.Input['ZoneSoaRecordArgs']]:
        """
        An `soa_record` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "soa_record")

    @soa_record.setter
    def soa_record(self, value: Optional[pulumi.Input['ZoneSoaRecordArgs']]):
        pulumi.set(self, "soa_record", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the Record Set.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ZoneState:
    def __init__(__self__, *,
                 max_number_of_record_sets: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 number_of_record_sets: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 soa_record: Optional[pulumi.Input['ZoneSoaRecordArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Zone resources.
        :param pulumi.Input[int] max_number_of_record_sets: (Optional) Maximum number of Records in the zone. Defaults to `1000`.
        :param pulumi.Input[str] name: The name of the DNS Zone. Must be a valid domain name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] name_servers: (Optional) A list of values that make up the NS record for the zone.
        :param pulumi.Input[int] number_of_record_sets: (Optional) The number of records already in the zone.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input['ZoneSoaRecordArgs'] soa_record: An `soa_record` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Record Set.
        """
        if max_number_of_record_sets is not None:
            pulumi.set(__self__, "max_number_of_record_sets", max_number_of_record_sets)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_servers is not None:
            pulumi.set(__self__, "name_servers", name_servers)
        if number_of_record_sets is not None:
            pulumi.set(__self__, "number_of_record_sets", number_of_record_sets)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if soa_record is not None:
            pulumi.set(__self__, "soa_record", soa_record)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="maxNumberOfRecordSets")
    def max_number_of_record_sets(self) -> Optional[pulumi.Input[int]]:
        """
        (Optional) Maximum number of Records in the zone. Defaults to `1000`.
        """
        return pulumi.get(self, "max_number_of_record_sets")

    @max_number_of_record_sets.setter
    def max_number_of_record_sets(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_number_of_record_sets", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the DNS Zone. Must be a valid domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (Optional) A list of values that make up the NS record for the zone.
        """
        return pulumi.get(self, "name_servers")

    @name_servers.setter
    def name_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "name_servers", value)

    @property
    @pulumi.getter(name="numberOfRecordSets")
    def number_of_record_sets(self) -> Optional[pulumi.Input[int]]:
        """
        (Optional) The number of records already in the zone.
        """
        return pulumi.get(self, "number_of_record_sets")

    @number_of_record_sets.setter
    def number_of_record_sets(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "number_of_record_sets", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="soaRecord")
    def soa_record(self) -> Optional[pulumi.Input['ZoneSoaRecordArgs']]:
        """
        An `soa_record` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "soa_record")

    @soa_record.setter
    def soa_record(self, value: Optional[pulumi.Input['ZoneSoaRecordArgs']]):
        pulumi.set(self, "soa_record", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the Record Set.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Zone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 soa_record: Optional[pulumi.Input[pulumi.InputType['ZoneSoaRecordArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Enables you to manage DNS zones within Azure DNS. These zones are hosted on Azure's name servers to which you can delegate the zone from the parent domain.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West Europe")
        example_public = azure.dns.Zone("example-public", resource_group_name=example.name)
        example_private = azure.privatedns.Zone("example-private", resource_group_name=example.name)
        ```

        ## Import

        DNS Zones can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dns/zone:Zone zone1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the DNS Zone. Must be a valid domain name.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ZoneSoaRecordArgs']] soa_record: An `soa_record` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Record Set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enables you to manage DNS zones within Azure DNS. These zones are hosted on Azure's name servers to which you can delegate the zone from the parent domain.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West Europe")
        example_public = azure.dns.Zone("example-public", resource_group_name=example.name)
        example_private = azure.privatedns.Zone("example-private", resource_group_name=example.name)
        ```

        ## Import

        DNS Zones can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dns/zone:Zone zone1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1
        ```

        :param str resource_name: The name of the resource.
        :param ZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 soa_record: Optional[pulumi.Input[pulumi.InputType['ZoneSoaRecordArgs']]] = None,
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
            __props__ = ZoneArgs.__new__(ZoneArgs)

            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["soa_record"] = soa_record
            __props__.__dict__["tags"] = tags
            __props__.__dict__["max_number_of_record_sets"] = None
            __props__.__dict__["name_servers"] = None
            __props__.__dict__["number_of_record_sets"] = None
        super(Zone, __self__).__init__(
            'azure:dns/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            max_number_of_record_sets: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            number_of_record_sets: Optional[pulumi.Input[int]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            soa_record: Optional[pulumi.Input[pulumi.InputType['ZoneSoaRecordArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] max_number_of_record_sets: (Optional) Maximum number of Records in the zone. Defaults to `1000`.
        :param pulumi.Input[str] name: The name of the DNS Zone. Must be a valid domain name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] name_servers: (Optional) A list of values that make up the NS record for the zone.
        :param pulumi.Input[int] number_of_record_sets: (Optional) The number of records already in the zone.
        :param pulumi.Input[str] resource_group_name: Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['ZoneSoaRecordArgs']] soa_record: An `soa_record` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the Record Set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneState.__new__(_ZoneState)

        __props__.__dict__["max_number_of_record_sets"] = max_number_of_record_sets
        __props__.__dict__["name"] = name
        __props__.__dict__["name_servers"] = name_servers
        __props__.__dict__["number_of_record_sets"] = number_of_record_sets
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["soa_record"] = soa_record
        __props__.__dict__["tags"] = tags
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="maxNumberOfRecordSets")
    def max_number_of_record_sets(self) -> pulumi.Output[int]:
        """
        (Optional) Maximum number of Records in the zone. Defaults to `1000`.
        """
        return pulumi.get(self, "max_number_of_record_sets")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the DNS Zone. Must be a valid domain name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> pulumi.Output[Sequence[str]]:
        """
        (Optional) A list of values that make up the NS record for the zone.
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter(name="numberOfRecordSets")
    def number_of_record_sets(self) -> pulumi.Output[int]:
        """
        (Optional) The number of records already in the zone.
        """
        return pulumi.get(self, "number_of_record_sets")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="soaRecord")
    def soa_record(self) -> pulumi.Output['outputs.ZoneSoaRecord']:
        """
        An `soa_record` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "soa_record")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the Record Set.
        """
        return pulumi.get(self, "tags")

