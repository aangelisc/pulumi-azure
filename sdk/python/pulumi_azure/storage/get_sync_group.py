# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSyncGroupResult',
    'AwaitableGetSyncGroupResult',
    'get_sync_group',
]

@pulumi.output_type
class GetSyncGroupResult:
    """
    A collection of values returned by getSyncGroup.
    """
    def __init__(__self__, id=None, name=None, storage_sync_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if storage_sync_id and not isinstance(storage_sync_id, str):
            raise TypeError("Expected argument 'storage_sync_id' to be a str")
        pulumi.set(__self__, "storage_sync_id", storage_sync_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="storageSyncId")
    def storage_sync_id(self) -> str:
        return pulumi.get(self, "storage_sync_id")


class AwaitableGetSyncGroupResult(GetSyncGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSyncGroupResult(
            id=self.id,
            name=self.name,
            storage_sync_id=self.storage_sync_id)


def get_sync_group(name: Optional[str] = None,
                   storage_sync_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSyncGroupResult:
    """
    Use this data source to access information about an existing Storage Sync Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.storage.get_sync_group(name="existing-ss-group",
        storage_sync_id="existing-ss-id")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of this Storage Sync Group.
    :param str storage_sync_id: The resource ID of the Storage Sync where this Storage Sync Group is.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['storageSyncId'] = storage_sync_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:storage/getSyncGroup:getSyncGroup', __args__, opts=opts, typ=GetSyncGroupResult).value

    return AwaitableGetSyncGroupResult(
        id=__ret__.id,
        name=__ret__.name,
        storage_sync_id=__ret__.storage_sync_id)
