# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetVirtualWanResult',
    'AwaitableGetVirtualWanResult',
    'get_virtual_wan',
]

@pulumi.output_type
class GetVirtualWanResult:
    """
    A collection of values returned by getVirtualWan.
    """
    def __init__(__self__, allow_branch_to_branch_traffic=None, disable_vpn_encryption=None, id=None, location=None, name=None, office365_local_breakout_category=None, resource_group_name=None, sku=None, tags=None, virtual_hub_ids=None, vpn_site_ids=None):
        if allow_branch_to_branch_traffic and not isinstance(allow_branch_to_branch_traffic, bool):
            raise TypeError("Expected argument 'allow_branch_to_branch_traffic' to be a bool")
        pulumi.set(__self__, "allow_branch_to_branch_traffic", allow_branch_to_branch_traffic)
        if disable_vpn_encryption and not isinstance(disable_vpn_encryption, bool):
            raise TypeError("Expected argument 'disable_vpn_encryption' to be a bool")
        pulumi.set(__self__, "disable_vpn_encryption", disable_vpn_encryption)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if office365_local_breakout_category and not isinstance(office365_local_breakout_category, str):
            raise TypeError("Expected argument 'office365_local_breakout_category' to be a str")
        pulumi.set(__self__, "office365_local_breakout_category", office365_local_breakout_category)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if virtual_hub_ids and not isinstance(virtual_hub_ids, list):
            raise TypeError("Expected argument 'virtual_hub_ids' to be a list")
        pulumi.set(__self__, "virtual_hub_ids", virtual_hub_ids)
        if vpn_site_ids and not isinstance(vpn_site_ids, list):
            raise TypeError("Expected argument 'vpn_site_ids' to be a list")
        pulumi.set(__self__, "vpn_site_ids", vpn_site_ids)

    @property
    @pulumi.getter(name="allowBranchToBranchTraffic")
    def allow_branch_to_branch_traffic(self) -> bool:
        """
        Is branch to branch traffic is allowed?
        """
        return pulumi.get(self, "allow_branch_to_branch_traffic")

    @property
    @pulumi.getter(name="disableVpnEncryption")
    def disable_vpn_encryption(self) -> bool:
        """
        Is VPN Encryption disabled?
        """
        return pulumi.get(self, "disable_vpn_encryption")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure Region where the Virtual Wan exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="office365LocalBreakoutCategory")
    def office365_local_breakout_category(self) -> str:
        """
        The Office365 Local Breakout Category.
        """
        return pulumi.get(self, "office365_local_breakout_category")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def sku(self) -> str:
        """
        Type of Virtual Wan (Basic or Standard).
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the Virtual Wan.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualHubIds")
    def virtual_hub_ids(self) -> Sequence[str]:
        """
        A list of Virtual Hubs ID's attached to this Virtual WAN.
        """
        return pulumi.get(self, "virtual_hub_ids")

    @property
    @pulumi.getter(name="vpnSiteIds")
    def vpn_site_ids(self) -> Sequence[str]:
        """
        A list of VPN Site ID's attached to this Virtual WAN.
        """
        return pulumi.get(self, "vpn_site_ids")


class AwaitableGetVirtualWanResult(GetVirtualWanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualWanResult(
            allow_branch_to_branch_traffic=self.allow_branch_to_branch_traffic,
            disable_vpn_encryption=self.disable_vpn_encryption,
            id=self.id,
            location=self.location,
            name=self.name,
            office365_local_breakout_category=self.office365_local_breakout_category,
            resource_group_name=self.resource_group_name,
            sku=self.sku,
            tags=self.tags,
            virtual_hub_ids=self.virtual_hub_ids,
            vpn_site_ids=self.vpn_site_ids)


def get_virtual_wan(name: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualWanResult:
    """
    Use this data source to access information about an existing Virtual Wan.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.network.get_virtual_wan(name="existing",
        resource_group_name="existing")
    pulumi.export("id", example.id)
    pulumi.export("allowBranchToBranchTraffic", example.allow_branch_to_branch_traffic)
    pulumi.export("disableVpnEncryption", data["azurerm_virtual_wan"]["exemple"]["disable_vpn_encryption"])
    pulumi.export("location", data["azurerm_virtual_wan"]["exemple"]["location"])
    pulumi.export("office365LocalBreakoutCategory", data["azurerm_virtual_wan"]["exemple"]["office365_local_breakout_category"])
    pulumi.export("sku", data["azurerm_virtual_wan"]["exemple"]["sku"])
    pulumi.export("tags", data["azurerm_virtual_wan"]["exemple"]["tags"])
    pulumi.export("virtualHubs", data["azurerm_virtual_wan"]["exemple"]["virtual_hubs"])
    pulumi.export("vpnSites", data["azurerm_virtual_wan"]["exemple"]["vpn_sites"])
    ```


    :param str name: The name of this Virtual Wan.
    :param str resource_group_name: The name of the Resource Group where the Virtual Wan exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getVirtualWan:getVirtualWan', __args__, opts=opts, typ=GetVirtualWanResult).value

    return AwaitableGetVirtualWanResult(
        allow_branch_to_branch_traffic=__ret__.allow_branch_to_branch_traffic,
        disable_vpn_encryption=__ret__.disable_vpn_encryption,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        office365_local_breakout_category=__ret__.office365_local_breakout_category,
        resource_group_name=__ret__.resource_group_name,
        sku=__ret__.sku,
        tags=__ret__.tags,
        virtual_hub_ids=__ret__.virtual_hub_ids,
        vpn_site_ids=__ret__.vpn_site_ids)
