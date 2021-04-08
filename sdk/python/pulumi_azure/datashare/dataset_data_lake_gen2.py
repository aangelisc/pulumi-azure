# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DatasetDataLakeGen2Args', 'DatasetDataLakeGen2']

@pulumi.input_type
class DatasetDataLakeGen2Args:
    def __init__(__self__, *,
                 file_system_name: pulumi.Input[str],
                 share_id: pulumi.Input[str],
                 storage_account_id: pulumi.Input[str],
                 file_path: Optional[pulumi.Input[str]] = None,
                 folder_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatasetDataLakeGen2 resource.
        :param pulumi.Input[str] file_system_name: The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] share_id: The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] storage_account_id: The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] file_path: The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] folder_path: The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        pulumi.set(__self__, "file_system_name", file_system_name)
        pulumi.set(__self__, "share_id", share_id)
        pulumi.set(__self__, "storage_account_id", storage_account_id)
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)
        if folder_path is not None:
            pulumi.set(__self__, "folder_path", folder_path)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> pulumi.Input[str]:
        """
        The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "file_system_name")

    @file_system_name.setter
    def file_system_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_name", value)

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> pulumi.Input[str]:
        """
        The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "share_id")

    @share_id.setter
    def share_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "share_id", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Input[str]:
        """
        The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_id", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="folderPath")
    def folder_path(self) -> Optional[pulumi.Input[str]]:
        """
        The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "folder_path")

    @folder_path.setter
    def folder_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_path", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DatasetDataLakeGen2State:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 folder_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 share_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatasetDataLakeGen2 resources.
        :param pulumi.Input[str] display_name: The name of the Data Share Dataset.
        :param pulumi.Input[str] file_path: The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] file_system_name: The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] folder_path: The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] share_id: The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] storage_account_id: The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)
        if file_system_name is not None:
            pulumi.set(__self__, "file_system_name", file_system_name)
        if folder_path is not None:
            pulumi.set(__self__, "folder_path", folder_path)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if share_id is not None:
            pulumi.set(__self__, "share_id", share_id)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Data Share Dataset.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "file_system_name")

    @file_system_name.setter
    def file_system_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_name", value)

    @property
    @pulumi.getter(name="folderPath")
    def folder_path(self) -> Optional[pulumi.Input[str]]:
        """
        The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "folder_path")

    @folder_path.setter
    def folder_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_path", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "share_id")

    @share_id.setter
    def share_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_id", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)


class DatasetDataLakeGen2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 folder_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 share_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Data Share Data Lake Gen2 Dataset.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuread as azuread

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.datashare.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            identity=azure.datashare.AccountIdentityArgs(
                type="SystemAssigned",
            ))
        example_share = azure.datashare.Share("exampleShare",
            account_id=example_account.id,
            kind="CopyBased")
        example_storage_account_account = azure.storage.Account("exampleStorage/accountAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_kind="BlobStorage",
            account_tier="Standard",
            account_replication_type="LRS")
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_storage / account_account["id"])
        example_service_principal = example_account.name.apply(lambda name: azuread.get_service_principal(display_name=name))
        example_assignment = azure.authorization.Assignment("exampleAssignment",
            scope=example_storage / account_account["id"],
            role_definition_name="Storage Blob Data Reader",
            principal_id=example_service_principal.object_id)
        example_dataset_data_lake_gen2 = azure.datashare.DatasetDataLakeGen2("exampleDatasetDataLakeGen2",
            share_id=example_share.id,
            storage_account_id=example_storage / account_account["id"],
            file_system_name=example_data_lake_gen2_filesystem.name,
            file_path="myfile.txt",
            opts=pulumi.ResourceOptions(depends_on=[example_assignment]))
        ```

        ## Import

        Data Share Data Lake Gen2 Datasets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_path: The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] file_system_name: The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] folder_path: The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] share_id: The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] storage_account_id: The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetDataLakeGen2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Data Share Data Lake Gen2 Dataset.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuread as azuread

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.datashare.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            identity=azure.datashare.AccountIdentityArgs(
                type="SystemAssigned",
            ))
        example_share = azure.datashare.Share("exampleShare",
            account_id=example_account.id,
            kind="CopyBased")
        example_storage_account_account = azure.storage.Account("exampleStorage/accountAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_kind="BlobStorage",
            account_tier="Standard",
            account_replication_type="LRS")
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_storage / account_account["id"])
        example_service_principal = example_account.name.apply(lambda name: azuread.get_service_principal(display_name=name))
        example_assignment = azure.authorization.Assignment("exampleAssignment",
            scope=example_storage / account_account["id"],
            role_definition_name="Storage Blob Data Reader",
            principal_id=example_service_principal.object_id)
        example_dataset_data_lake_gen2 = azure.datashare.DatasetDataLakeGen2("exampleDatasetDataLakeGen2",
            share_id=example_share.id,
            storage_account_id=example_storage / account_account["id"],
            file_system_name=example_data_lake_gen2_filesystem.name,
            file_path="myfile.txt",
            opts=pulumi.ResourceOptions(depends_on=[example_assignment]))
        ```

        ## Import

        Data Share Data Lake Gen2 Datasets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
        ```

        :param str resource_name: The name of the resource.
        :param DatasetDataLakeGen2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetDataLakeGen2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 folder_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 share_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DatasetDataLakeGen2Args.__new__(DatasetDataLakeGen2Args)

            __props__.__dict__["file_path"] = file_path
            if file_system_name is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_name'")
            __props__.__dict__["file_system_name"] = file_system_name
            __props__.__dict__["folder_path"] = folder_path
            __props__.__dict__["name"] = name
            if share_id is None and not opts.urn:
                raise TypeError("Missing required property 'share_id'")
            __props__.__dict__["share_id"] = share_id
            if storage_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_id'")
            __props__.__dict__["storage_account_id"] = storage_account_id
            __props__.__dict__["display_name"] = None
        super(DatasetDataLakeGen2, __self__).__init__(
            'azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            file_path: Optional[pulumi.Input[str]] = None,
            file_system_name: Optional[pulumi.Input[str]] = None,
            folder_path: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            share_id: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None) -> 'DatasetDataLakeGen2':
        """
        Get an existing DatasetDataLakeGen2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The name of the Data Share Dataset.
        :param pulumi.Input[str] file_path: The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] file_system_name: The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] folder_path: The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] name: The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] share_id: The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        :param pulumi.Input[str] storage_account_id: The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatasetDataLakeGen2State.__new__(_DatasetDataLakeGen2State)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["file_path"] = file_path
        __props__.__dict__["file_system_name"] = file_system_name
        __props__.__dict__["folder_path"] = folder_path
        __props__.__dict__["name"] = name
        __props__.__dict__["share_id"] = share_id
        __props__.__dict__["storage_account_id"] = storage_account_id
        return DatasetDataLakeGen2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The name of the Data Share Dataset.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> pulumi.Output[Optional[str]]:
        """
        The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folder_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "file_path")

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> pulumi.Output[str]:
        """
        The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "file_system_name")

    @property
    @pulumi.getter(name="folderPath")
    def folder_path(self) -> pulumi.Output[Optional[str]]:
        """
        The folder path in the data lake file system to be shared with the receiver. Conflicts with `file_path` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "folder_path")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "share_id")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[str]:
        """
        The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
        """
        return pulumi.get(self, "storage_account_id")

