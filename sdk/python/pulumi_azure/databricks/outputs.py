# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'WorkspaceCustomParameters',
]

@pulumi.output_type
class WorkspaceCustomParameters(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "noPublicIp":
            suggest = "no_public_ip"
        elif key == "privateSubnetName":
            suggest = "private_subnet_name"
        elif key == "publicSubnetName":
            suggest = "public_subnet_name"
        elif key == "virtualNetworkId":
            suggest = "virtual_network_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkspaceCustomParameters. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkspaceCustomParameters.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkspaceCustomParameters.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 no_public_ip: Optional[bool] = None,
                 private_subnet_name: Optional[str] = None,
                 public_subnet_name: Optional[str] = None,
                 virtual_network_id: Optional[str] = None):
        """
        :param bool no_public_ip: Are public IP Addresses not allowed?
        :param str private_subnet_name: The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        :param str public_subnet_name: The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        :param str virtual_network_id: The ID of a Virtual Network where this Databricks Cluster should be created.
        """
        if no_public_ip is not None:
            pulumi.set(__self__, "no_public_ip", no_public_ip)
        if private_subnet_name is not None:
            pulumi.set(__self__, "private_subnet_name", private_subnet_name)
        if public_subnet_name is not None:
            pulumi.set(__self__, "public_subnet_name", public_subnet_name)
        if virtual_network_id is not None:
            pulumi.set(__self__, "virtual_network_id", virtual_network_id)

    @property
    @pulumi.getter(name="noPublicIp")
    def no_public_ip(self) -> Optional[bool]:
        """
        Are public IP Addresses not allowed?
        """
        return pulumi.get(self, "no_public_ip")

    @property
    @pulumi.getter(name="privateSubnetName")
    def private_subnet_name(self) -> Optional[str]:
        """
        The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        """
        return pulumi.get(self, "private_subnet_name")

    @property
    @pulumi.getter(name="publicSubnetName")
    def public_subnet_name(self) -> Optional[str]:
        """
        The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        """
        return pulumi.get(self, "public_subnet_name")

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[str]:
        """
        The ID of a Virtual Network where this Databricks Cluster should be created.
        """
        return pulumi.get(self, "virtual_network_id")


