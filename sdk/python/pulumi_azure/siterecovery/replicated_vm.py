# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ReplicatedVM(pulumi.CustomResource):
    managed_disks: pulumi.Output[list]
    """
    One or more `managed_disk` block.

      * `diskId` (`str`) - Id of disk that should be replicated.
      * `stagingStorageAccountId` (`str`) - Storage account that should be used for caching.
      * `targetDiskType` (`str`) - What type should the disk be when a failover is done.
      * `targetReplicaDiskType` (`str`) - What type should the disk be that holds the replication data.
      * `target_resource_group_id` (`str`) - Resource group disk should belong to when a failover is done.
    """
    name: pulumi.Output[str]
    """
    The name of the network mapping.
    """
    network_interfaces: pulumi.Output[list]
    """
    One or more `network_interface` block.

      * `sourceNetworkInterfaceId` (`str`) - Id source network interface.
      * `targetStaticIp` (`str`) - Static IP to assign when a failover is done.
      * `targetSubnetName` (`str`) - Name of the subnet to to use when a failover is done.
    """
    recovery_replication_policy_id: pulumi.Output[str]
    recovery_vault_name: pulumi.Output[str]
    """
    The name of the vault that should be updated.
    """
    resource_group_name: pulumi.Output[str]
    """
    Name of the resource group where the vault that should be updated is located.
    """
    source_recovery_fabric_name: pulumi.Output[str]
    """
    Name of fabric that should contains this replication.
    """
    source_recovery_protection_container_name: pulumi.Output[str]
    """
    Name of the protection container to use.
    """
    source_vm_id: pulumi.Output[str]
    """
    Id of the VM to replicate
    """
    target_availability_set_id: pulumi.Output[str]
    """
    Id of availability set that the new VM should belong to when a failover is done.
    """
    target_network_id: pulumi.Output[str]
    """
    Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
    """
    target_recovery_fabric_id: pulumi.Output[str]
    """
    Id of fabric where the VM replication should be handled when a failover is done.
    """
    target_recovery_protection_container_id: pulumi.Output[str]
    """
    Id of protection container where the VM replication should be created when a failover is done.
    """
    target_resource_group_id: pulumi.Output[str]
    """
    Id of resource group where the VM should be created when a failover is done.
    """
    def __init__(__self__, resource_name, opts=None, managed_disks=None, name=None, network_interfaces=None, recovery_replication_policy_id=None, recovery_vault_name=None, resource_group_name=None, source_recovery_fabric_name=None, source_recovery_protection_container_name=None, source_vm_id=None, target_availability_set_id=None, target_network_id=None, target_recovery_fabric_id=None, target_recovery_protection_container_id=None, target_resource_group_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a VM replicated using Azure Site Recovery (Azure to Azure only). A replicated VM keeps a copiously updated image of the VM in another region in order to be able to start the VM in that region in case of a disaster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] managed_disks: One or more `managed_disk` block.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[list] network_interfaces: One or more `network_interface` block.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_recovery_fabric_name: Name of fabric that should contains this replication.
        :param pulumi.Input[str] source_recovery_protection_container_name: Name of the protection container to use.
        :param pulumi.Input[str] source_vm_id: Id of the VM to replicate
        :param pulumi.Input[str] target_availability_set_id: Id of availability set that the new VM should belong to when a failover is done.
        :param pulumi.Input[str] target_network_id: Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        :param pulumi.Input[str] target_recovery_fabric_id: Id of fabric where the VM replication should be handled when a failover is done.
        :param pulumi.Input[str] target_recovery_protection_container_id: Id of protection container where the VM replication should be created when a failover is done.
        :param pulumi.Input[str] target_resource_group_id: Id of resource group where the VM should be created when a failover is done.

        The **managed_disks** object supports the following:

          * `diskId` (`pulumi.Input[str]`) - Id of disk that should be replicated.
          * `stagingStorageAccountId` (`pulumi.Input[str]`) - Storage account that should be used for caching.
          * `targetDiskType` (`pulumi.Input[str]`) - What type should the disk be when a failover is done.
          * `targetReplicaDiskType` (`pulumi.Input[str]`) - What type should the disk be that holds the replication data.
          * `target_resource_group_id` (`pulumi.Input[str]`) - Resource group disk should belong to when a failover is done.

        The **network_interfaces** object supports the following:

          * `sourceNetworkInterfaceId` (`pulumi.Input[str]`) - Id source network interface.
          * `targetStaticIp` (`pulumi.Input[str]`) - Static IP to assign when a failover is done.
          * `targetSubnetName` (`pulumi.Input[str]`) - Name of the subnet to to use when a failover is done.
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['managed_disks'] = managed_disks
            __props__['name'] = name
            __props__['network_interfaces'] = network_interfaces
            if recovery_replication_policy_id is None:
                raise TypeError("Missing required property 'recovery_replication_policy_id'")
            __props__['recovery_replication_policy_id'] = recovery_replication_policy_id
            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_recovery_fabric_name is None:
                raise TypeError("Missing required property 'source_recovery_fabric_name'")
            __props__['source_recovery_fabric_name'] = source_recovery_fabric_name
            if source_recovery_protection_container_name is None:
                raise TypeError("Missing required property 'source_recovery_protection_container_name'")
            __props__['source_recovery_protection_container_name'] = source_recovery_protection_container_name
            if source_vm_id is None:
                raise TypeError("Missing required property 'source_vm_id'")
            __props__['source_vm_id'] = source_vm_id
            __props__['target_availability_set_id'] = target_availability_set_id
            __props__['target_network_id'] = target_network_id
            if target_recovery_fabric_id is None:
                raise TypeError("Missing required property 'target_recovery_fabric_id'")
            __props__['target_recovery_fabric_id'] = target_recovery_fabric_id
            if target_recovery_protection_container_id is None:
                raise TypeError("Missing required property 'target_recovery_protection_container_id'")
            __props__['target_recovery_protection_container_id'] = target_recovery_protection_container_id
            if target_resource_group_id is None:
                raise TypeError("Missing required property 'target_resource_group_id'")
            __props__['target_resource_group_id'] = target_resource_group_id
        super(ReplicatedVM, __self__).__init__(
            'azure:siterecovery/replicatedVM:ReplicatedVM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, managed_disks=None, name=None, network_interfaces=None, recovery_replication_policy_id=None, recovery_vault_name=None, resource_group_name=None, source_recovery_fabric_name=None, source_recovery_protection_container_name=None, source_vm_id=None, target_availability_set_id=None, target_network_id=None, target_recovery_fabric_id=None, target_recovery_protection_container_id=None, target_resource_group_id=None):
        """
        Get an existing ReplicatedVM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] managed_disks: One or more `managed_disk` block.
        :param pulumi.Input[str] name: The name of the network mapping.
        :param pulumi.Input[list] network_interfaces: One or more `network_interface` block.
        :param pulumi.Input[str] recovery_vault_name: The name of the vault that should be updated.
        :param pulumi.Input[str] resource_group_name: Name of the resource group where the vault that should be updated is located.
        :param pulumi.Input[str] source_recovery_fabric_name: Name of fabric that should contains this replication.
        :param pulumi.Input[str] source_recovery_protection_container_name: Name of the protection container to use.
        :param pulumi.Input[str] source_vm_id: Id of the VM to replicate
        :param pulumi.Input[str] target_availability_set_id: Id of availability set that the new VM should belong to when a failover is done.
        :param pulumi.Input[str] target_network_id: Network to use when a failover is done (recommended to set if any network_interface is configured for failover).
        :param pulumi.Input[str] target_recovery_fabric_id: Id of fabric where the VM replication should be handled when a failover is done.
        :param pulumi.Input[str] target_recovery_protection_container_id: Id of protection container where the VM replication should be created when a failover is done.
        :param pulumi.Input[str] target_resource_group_id: Id of resource group where the VM should be created when a failover is done.

        The **managed_disks** object supports the following:

          * `diskId` (`pulumi.Input[str]`) - Id of disk that should be replicated.
          * `stagingStorageAccountId` (`pulumi.Input[str]`) - Storage account that should be used for caching.
          * `targetDiskType` (`pulumi.Input[str]`) - What type should the disk be when a failover is done.
          * `targetReplicaDiskType` (`pulumi.Input[str]`) - What type should the disk be that holds the replication data.
          * `target_resource_group_id` (`pulumi.Input[str]`) - Resource group disk should belong to when a failover is done.

        The **network_interfaces** object supports the following:

          * `sourceNetworkInterfaceId` (`pulumi.Input[str]`) - Id source network interface.
          * `targetStaticIp` (`pulumi.Input[str]`) - Static IP to assign when a failover is done.
          * `targetSubnetName` (`pulumi.Input[str]`) - Name of the subnet to to use when a failover is done.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["managed_disks"] = managed_disks
        __props__["name"] = name
        __props__["network_interfaces"] = network_interfaces
        __props__["recovery_replication_policy_id"] = recovery_replication_policy_id
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["source_recovery_fabric_name"] = source_recovery_fabric_name
        __props__["source_recovery_protection_container_name"] = source_recovery_protection_container_name
        __props__["source_vm_id"] = source_vm_id
        __props__["target_availability_set_id"] = target_availability_set_id
        __props__["target_network_id"] = target_network_id
        __props__["target_recovery_fabric_id"] = target_recovery_fabric_id
        __props__["target_recovery_protection_container_id"] = target_recovery_protection_container_id
        __props__["target_resource_group_id"] = target_resource_group_id
        return ReplicatedVM(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
