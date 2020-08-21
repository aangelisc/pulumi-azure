# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GlobalVMShutdownScheduleNotificationSettings',
    'LinuxVirtualMachineGalleryImageReference',
    'LinuxVirtualMachineInboundNatRule',
    'ScheduleDailyRecurrence',
    'ScheduleHourlyRecurrence',
    'ScheduleNotificationSettings',
    'ScheduleWeeklyRecurrence',
    'VirtualNetworkSubnet',
    'WindowsVirtualMachineGalleryImageReference',
    'WindowsVirtualMachineInboundNatRule',
    'GetVirtualNetworkAllowedSubnetResult',
    'GetVirtualNetworkSubnetOverrideResult',
]

@pulumi.output_type
class GlobalVMShutdownScheduleNotificationSettings(dict):
    def __init__(__self__, *,
                 enabled: bool,
                 time_in_minutes: Optional[float] = None,
                 webhook_url: Optional[str] = None):
        """
        :param bool enabled: Whether to enable pre-shutdown notifications. Possible values are `true` and `false`. Defaults to `false`
        :param float time_in_minutes: Time in minutes between 15 and 120 before a shutdown event at which a notification will be sent. Defaults to `30`.
        :param str webhook_url: The webhook URL to which the notification will be sent. Required if `enabled` is `true`. Optional otherwise.
        """
        pulumi.set(__self__, "enabled", enabled)
        if time_in_minutes is not None:
            pulumi.set(__self__, "time_in_minutes", time_in_minutes)
        if webhook_url is not None:
            pulumi.set(__self__, "webhook_url", webhook_url)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether to enable pre-shutdown notifications. Possible values are `true` and `false`. Defaults to `false`
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="timeInMinutes")
    def time_in_minutes(self) -> Optional[float]:
        """
        Time in minutes between 15 and 120 before a shutdown event at which a notification will be sent. Defaults to `30`.
        """
        return pulumi.get(self, "time_in_minutes")

    @property
    @pulumi.getter(name="webhookUrl")
    def webhook_url(self) -> Optional[str]:
        """
        The webhook URL to which the notification will be sent. Required if `enabled` is `true`. Optional otherwise.
        """
        return pulumi.get(self, "webhook_url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LinuxVirtualMachineGalleryImageReference(dict):
    def __init__(__self__, *,
                 offer: str,
                 publisher: str,
                 sku: str,
                 version: str):
        """
        :param str offer: The Offer of the Gallery Image. Changing this forces a new resource to be created.
        :param str publisher: The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        :param str sku: The SKU of the Gallery Image. Changing this forces a new resource to be created.
        :param str version: The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "offer", offer)
        pulumi.set(__self__, "publisher", publisher)
        pulumi.set(__self__, "sku", sku)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def offer(self) -> str:
        """
        The Offer of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "offer")

    @property
    @pulumi.getter
    def publisher(self) -> str:
        """
        The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter
    def sku(self) -> str:
        """
        The SKU of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LinuxVirtualMachineInboundNatRule(dict):
    def __init__(__self__, *,
                 backend_port: float,
                 protocol: str,
                 frontend_port: Optional[float] = None):
        """
        :param float backend_port: The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        :param str protocol: The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        :param float frontend_port: The frontend port associated with this Inbound NAT Rule.
        """
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "protocol", protocol)
        if frontend_port is not None:
            pulumi.set(__self__, "frontend_port", frontend_port)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> float:
        """
        The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> Optional[float]:
        """
        The frontend port associated with this Inbound NAT Rule.
        """
        return pulumi.get(self, "frontend_port")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleDailyRecurrence(dict):
    def __init__(__self__, *,
                 time: str):
        """
        :param str time: The time each day when the schedule takes effect.
        """
        pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def time(self) -> str:
        """
        The time each day when the schedule takes effect.
        """
        return pulumi.get(self, "time")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleHourlyRecurrence(dict):
    def __init__(__self__, *,
                 minute: float):
        pulumi.set(__self__, "minute", minute)

    @property
    @pulumi.getter
    def minute(self) -> float:
        return pulumi.get(self, "minute")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleNotificationSettings(dict):
    def __init__(__self__, *,
                 status: Optional[str] = None,
                 time_in_minutes: Optional[float] = None,
                 webhook_url: Optional[str] = None):
        """
        :param str status: The status of the notification. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`
        :param float time_in_minutes: Time in minutes before event at which notification will be sent.
        :param str webhook_url: The webhook URL to which the notification will be sent.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if time_in_minutes is not None:
            pulumi.set(__self__, "time_in_minutes", time_in_minutes)
        if webhook_url is not None:
            pulumi.set(__self__, "webhook_url", webhook_url)

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the notification. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="timeInMinutes")
    def time_in_minutes(self) -> Optional[float]:
        """
        Time in minutes before event at which notification will be sent.
        """
        return pulumi.get(self, "time_in_minutes")

    @property
    @pulumi.getter(name="webhookUrl")
    def webhook_url(self) -> Optional[str]:
        """
        The webhook URL to which the notification will be sent.
        """
        return pulumi.get(self, "webhook_url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleWeeklyRecurrence(dict):
    def __init__(__self__, *,
                 time: str,
                 week_days: Optional[List[str]] = None):
        """
        :param str time: The time when the schedule takes effect.
        :param List[str] week_days: A list of days that this schedule takes effect . Possible values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        pulumi.set(__self__, "time", time)
        if week_days is not None:
            pulumi.set(__self__, "week_days", week_days)

    @property
    @pulumi.getter
    def time(self) -> str:
        """
        The time when the schedule takes effect.
        """
        return pulumi.get(self, "time")

    @property
    @pulumi.getter(name="weekDays")
    def week_days(self) -> Optional[List[str]]:
        """
        A list of days that this schedule takes effect . Possible values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        return pulumi.get(self, "week_days")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VirtualNetworkSubnet(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 use_in_virtual_machine_creation: Optional[str] = None,
                 use_public_ip_address: Optional[str] = None):
        """
        :param str name: Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        :param str use_in_virtual_machine_creation: Can this subnet be used for creating Virtual Machines? Possible values are `Allow`, `Default` and `Deny`.
        :param str use_public_ip_address: Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are `Allow`, `Default` and `Deny`.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if use_in_virtual_machine_creation is not None:
            pulumi.set(__self__, "use_in_virtual_machine_creation", use_in_virtual_machine_creation)
        if use_public_ip_address is not None:
            pulumi.set(__self__, "use_public_ip_address", use_public_ip_address)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="useInVirtualMachineCreation")
    def use_in_virtual_machine_creation(self) -> Optional[str]:
        """
        Can this subnet be used for creating Virtual Machines? Possible values are `Allow`, `Default` and `Deny`.
        """
        return pulumi.get(self, "use_in_virtual_machine_creation")

    @property
    @pulumi.getter(name="usePublicIpAddress")
    def use_public_ip_address(self) -> Optional[str]:
        """
        Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are `Allow`, `Default` and `Deny`.
        """
        return pulumi.get(self, "use_public_ip_address")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WindowsVirtualMachineGalleryImageReference(dict):
    def __init__(__self__, *,
                 offer: str,
                 publisher: str,
                 sku: str,
                 version: str):
        """
        :param str offer: The Offer of the Gallery Image. Changing this forces a new resource to be created.
        :param str publisher: The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        :param str sku: The SKU of the Gallery Image. Changing this forces a new resource to be created.
        :param str version: The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "offer", offer)
        pulumi.set(__self__, "publisher", publisher)
        pulumi.set(__self__, "sku", sku)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def offer(self) -> str:
        """
        The Offer of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "offer")

    @property
    @pulumi.getter
    def publisher(self) -> str:
        """
        The Publisher of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter
    def sku(self) -> str:
        """
        The SKU of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The Version of the Gallery Image. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WindowsVirtualMachineInboundNatRule(dict):
    def __init__(__self__, *,
                 backend_port: float,
                 protocol: str,
                 frontend_port: Optional[float] = None):
        """
        :param float backend_port: The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        :param str protocol: The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        :param float frontend_port: The frontend port associated with this Inbound NAT Rule.
        """
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "protocol", protocol)
        if frontend_port is not None:
            pulumi.set(__self__, "frontend_port", frontend_port)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> float:
        """
        The Backend Port associated with this NAT Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The Protocol used for this NAT Rule. Possible values are `Tcp` and `Udp`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> Optional[float]:
        """
        The frontend port associated with this Inbound NAT Rule.
        """
        return pulumi.get(self, "frontend_port")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetVirtualNetworkAllowedSubnetResult(dict):
    def __init__(__self__, *,
                 allow_public_ip: str,
                 lab_subnet_name: str,
                 resource_id: str):
        """
        :param str allow_public_ip: Indicates if this subnet allows public IP addresses. Possible values are `Allow`, `Default` and `Deny`.
        :param str lab_subnet_name: The name of the subnet.
        :param str resource_id: The resource identifier for the subnet.
        """
        pulumi.set(__self__, "allow_public_ip", allow_public_ip)
        pulumi.set(__self__, "lab_subnet_name", lab_subnet_name)
        pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="allowPublicIp")
    def allow_public_ip(self) -> str:
        """
        Indicates if this subnet allows public IP addresses. Possible values are `Allow`, `Default` and `Deny`.
        """
        return pulumi.get(self, "allow_public_ip")

    @property
    @pulumi.getter(name="labSubnetName")
    def lab_subnet_name(self) -> str:
        """
        The name of the subnet.
        """
        return pulumi.get(self, "lab_subnet_name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The resource identifier for the subnet.
        """
        return pulumi.get(self, "resource_id")


@pulumi.output_type
class GetVirtualNetworkSubnetOverrideResult(dict):
    def __init__(__self__, *,
                 lab_subnet_name: str,
                 resource_id: str,
                 use_in_vm_creation_permission: str,
                 use_public_ip_address_permission: str,
                 virtual_network_pool_name: str):
        """
        :param str lab_subnet_name: The name of the subnet.
        :param str resource_id: The resource identifier for the subnet.
        :param str use_in_vm_creation_permission: Indicates if the subnet can be used for VM creation.  Possible values are `Allow`, `Default` and `Deny`.
        :param str virtual_network_pool_name: The virtual network pool associated with this subnet.
        """
        pulumi.set(__self__, "lab_subnet_name", lab_subnet_name)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "use_in_vm_creation_permission", use_in_vm_creation_permission)
        pulumi.set(__self__, "use_public_ip_address_permission", use_public_ip_address_permission)
        pulumi.set(__self__, "virtual_network_pool_name", virtual_network_pool_name)

    @property
    @pulumi.getter(name="labSubnetName")
    def lab_subnet_name(self) -> str:
        """
        The name of the subnet.
        """
        return pulumi.get(self, "lab_subnet_name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The resource identifier for the subnet.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="useInVmCreationPermission")
    def use_in_vm_creation_permission(self) -> str:
        """
        Indicates if the subnet can be used for VM creation.  Possible values are `Allow`, `Default` and `Deny`.
        """
        return pulumi.get(self, "use_in_vm_creation_permission")

    @property
    @pulumi.getter(name="usePublicIpAddressPermission")
    def use_public_ip_address_permission(self) -> str:
        return pulumi.get(self, "use_public_ip_address_permission")

    @property
    @pulumi.getter(name="virtualNetworkPoolName")
    def virtual_network_pool_name(self) -> str:
        """
        The virtual network pool associated with this subnet.
        """
        return pulumi.get(self, "virtual_network_pool_name")


