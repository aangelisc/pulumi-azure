# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEventHubResult',
    'AwaitableGetEventHubResult',
    'get_event_hub',
]

@pulumi.output_type
class GetEventHubResult:
    """
    A collection of values returned by getEventHub.
    """
    def __init__(__self__, id=None, name=None, namespace_name=None, partition_count=None, partition_ids=None, resource_group_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        pulumi.set(__self__, "namespace_name", namespace_name)
        if partition_count and not isinstance(partition_count, int):
            raise TypeError("Expected argument 'partition_count' to be a int")
        pulumi.set(__self__, "partition_count", partition_count)
        if partition_ids and not isinstance(partition_ids, list):
            raise TypeError("Expected argument 'partition_ids' to be a list")
        pulumi.set(__self__, "partition_ids", partition_ids)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)

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
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> str:
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> int:
        """
        The number of partitions in the EventHub.
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="partitionIds")
    def partition_ids(self) -> Sequence[str]:
        """
        The identifiers for the partitions of this EventHub.
        """
        return pulumi.get(self, "partition_ids")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")


class AwaitableGetEventHubResult(GetEventHubResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventHubResult(
            id=self.id,
            name=self.name,
            namespace_name=self.namespace_name,
            partition_count=self.partition_count,
            partition_ids=self.partition_ids,
            resource_group_name=self.resource_group_name)


def get_event_hub(name: Optional[str] = None,
                  namespace_name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventHubResult:
    """
    Use this data source to access information about an existing EventHub.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.eventhub.get_event_hub(name="search-eventhub",
        resource_group_name="search-service",
        namespace_name="search-eventhubns")
    pulumi.export("eventhubId", example.id)
    ```


    :param str name: The name of this EventHub.
    :param str namespace_name: The name of the EventHub Namespace where the EventHub exists.
    :param str resource_group_name: The name of the Resource Group where the EventHub exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:eventhub/getEventHub:getEventHub', __args__, opts=opts, typ=GetEventHubResult).value

    return AwaitableGetEventHubResult(
        id=__ret__.id,
        name=__ret__.name,
        namespace_name=__ret__.namespace_name,
        partition_count=__ret__.partition_count,
        partition_ids=__ret__.partition_ids,
        resource_group_name=__ret__.resource_group_name)
