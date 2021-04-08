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

__all__ = ['SnapshotArgs', 'Snapshot']

@pulumi.input_type
class SnapshotArgs:
    def __init__(__self__, *,
                 create_option: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 encryption_settings: Optional[pulumi.Input['SnapshotEncryptionSettingsArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Snapshot resource.
        :param pulumi.Input[str] create_option: Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[int] disk_size_gb: The size of the Snapshotted Disk in GB.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_resource_id: Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_uri: Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "create_option", create_option)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if disk_size_gb is not None:
            pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if encryption_settings is not None:
            pulumi.set(__self__, "encryption_settings", encryption_settings)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)
        if source_uri is not None:
            pulumi.set(__self__, "source_uri", source_uri)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> pulumi.Input[str]:
        """
        Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "create_option")

    @create_option.setter
    def create_option(self, value: pulumi.Input[str]):
        pulumi.set(self, "create_option", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the Snapshotted Disk in GB.
        """
        return pulumi.get(self, "disk_size_gb")

    @disk_size_gb.setter
    def disk_size_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size_gb", value)

    @property
    @pulumi.getter(name="encryptionSettings")
    def encryption_settings(self) -> Optional[pulumi.Input['SnapshotEncryptionSettingsArgs']]:
        return pulumi.get(self, "encryption_settings")

    @encryption_settings.setter
    def encryption_settings(self, value: Optional[pulumi.Input['SnapshotEncryptionSettingsArgs']]):
        pulumi.set(self, "encryption_settings", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_resource_id")

    @source_resource_id.setter
    def source_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_resource_id", value)

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_uri")

    @source_uri.setter
    def source_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_uri", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)

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
class _SnapshotState:
    def __init__(__self__, *,
                 create_option: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 encryption_settings: Optional[pulumi.Input['SnapshotEncryptionSettingsArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Snapshot resources.
        :param pulumi.Input[str] create_option: Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] disk_size_gb: The size of the Snapshotted Disk in GB.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_resource_id: Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_uri: Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if create_option is not None:
            pulumi.set(__self__, "create_option", create_option)
        if disk_size_gb is not None:
            pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if encryption_settings is not None:
            pulumi.set(__self__, "encryption_settings", encryption_settings)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)
        if source_uri is not None:
            pulumi.set(__self__, "source_uri", source_uri)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "create_option")

    @create_option.setter
    def create_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_option", value)

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the Snapshotted Disk in GB.
        """
        return pulumi.get(self, "disk_size_gb")

    @disk_size_gb.setter
    def disk_size_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size_gb", value)

    @property
    @pulumi.getter(name="encryptionSettings")
    def encryption_settings(self) -> Optional[pulumi.Input['SnapshotEncryptionSettingsArgs']]:
        return pulumi.get(self, "encryption_settings")

    @encryption_settings.setter
    def encryption_settings(self, value: Optional[pulumi.Input['SnapshotEncryptionSettingsArgs']]):
        pulumi.set(self, "encryption_settings", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_resource_id")

    @source_resource_id.setter
    def source_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_resource_id", value)

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_uri")

    @source_uri.setter
    def source_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_uri", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)

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


class Snapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_option: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 encryption_settings: Optional[pulumi.Input[pulumi.InputType['SnapshotEncryptionSettingsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Disk Snapshot.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_managed_disk = azure.compute.ManagedDisk("exampleManagedDisk",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb=10)
        example_snapshot = azure.compute.Snapshot("exampleSnapshot",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            create_option="Copy",
            source_uri=example_managed_disk.id)
        ```

        ## Import

        Snapshots can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/snapshot:Snapshot example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/snapshots/snapshot1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] disk_size_gb: The size of the Snapshotted Disk in GB.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_resource_id: Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_uri: Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Disk Snapshot.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_managed_disk = azure.compute.ManagedDisk("exampleManagedDisk",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_account_type="Standard_LRS",
            create_option="Empty",
            disk_size_gb=10)
        example_snapshot = azure.compute.Snapshot("exampleSnapshot",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            create_option="Copy",
            source_uri=example_managed_disk.id)
        ```

        ## Import

        Snapshots can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/snapshot:Snapshot example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/snapshots/snapshot1
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_option: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 encryption_settings: Optional[pulumi.Input[pulumi.InputType['SnapshotEncryptionSettingsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SnapshotArgs.__new__(SnapshotArgs)

            if create_option is None and not opts.urn:
                raise TypeError("Missing required property 'create_option'")
            __props__.__dict__["create_option"] = create_option
            __props__.__dict__["disk_size_gb"] = disk_size_gb
            __props__.__dict__["encryption_settings"] = encryption_settings
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["source_resource_id"] = source_resource_id
            __props__.__dict__["source_uri"] = source_uri
            __props__.__dict__["storage_account_id"] = storage_account_id
            __props__.__dict__["tags"] = tags
        super(Snapshot, __self__).__init__(
            'azure:compute/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_option: Optional[pulumi.Input[str]] = None,
            disk_size_gb: Optional[pulumi.Input[int]] = None,
            encryption_settings: Optional[pulumi.Input[pulumi.InputType['SnapshotEncryptionSettingsArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            source_resource_id: Optional[pulumi.Input[str]] = None,
            source_uri: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] disk_size_gb: The size of the Snapshotted Disk in GB.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_resource_id: Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] source_uri: Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotState.__new__(_SnapshotState)

        __props__.__dict__["create_option"] = create_option
        __props__.__dict__["disk_size_gb"] = disk_size_gb
        __props__.__dict__["encryption_settings"] = encryption_settings
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["source_resource_id"] = source_resource_id
        __props__.__dict__["source_uri"] = source_uri
        __props__.__dict__["storage_account_id"] = storage_account_id
        __props__.__dict__["tags"] = tags
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> pulumi.Output[str]:
        """
        Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "create_option")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> pulumi.Output[int]:
        """
        The size of the Snapshotted Disk in GB.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="encryptionSettings")
    def encryption_settings(self) -> pulumi.Output[Optional['outputs.SnapshotEncryptionSettings']]:
        return pulumi.get(self, "encryption_settings")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_resource_id")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

