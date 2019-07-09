# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DataDiskAttachment(pulumi.CustomResource):
    caching: pulumi.Output[str]
    """
    Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
    """
    create_option: pulumi.Output[str]
    """
    The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
    """
    lun: pulumi.Output[float]
    """
    The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
    """
    managed_disk_id: pulumi.Output[str]
    """
    The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
    """
    virtual_machine_id: pulumi.Output[str]
    """
    The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
    """
    write_accelerator_enabled: pulumi.Output[bool]
    """
    Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
    """
    def __init__(__self__, resource_name, opts=None, caching=None, create_option=None, lun=None, managed_disk_id=None, virtual_machine_id=None, write_accelerator_enabled=None, __name__=None, __opts__=None):
        """
        Manages attaching a Disk to a Virtual Machine.
        
        > **NOTE:** Data Disks can be attached either directly on the `azurerm_virtual_machine` resource, or using the `azurerm_virtual_machine_data_disk_attachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
        
        > **Please Note:** only Managed Disks are supported via this separate resource, Unmanaged Disks can be attached using the `storage_data_disk` block in the `azurerm_virtual_machine` resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] caching: Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
        :param pulumi.Input[str] create_option: The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
        :param pulumi.Input[float] lun: The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_disk_id: The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] write_accelerator_enabled: Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine_data_disk_attachment.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if caching is None:
            raise TypeError("Missing required property 'caching'")
        __props__['caching'] = caching

        __props__['create_option'] = create_option

        if lun is None:
            raise TypeError("Missing required property 'lun'")
        __props__['lun'] = lun

        if managed_disk_id is None:
            raise TypeError("Missing required property 'managed_disk_id'")
        __props__['managed_disk_id'] = managed_disk_id

        if virtual_machine_id is None:
            raise TypeError("Missing required property 'virtual_machine_id'")
        __props__['virtual_machine_id'] = virtual_machine_id

        __props__['write_accelerator_enabled'] = write_accelerator_enabled

        super(DataDiskAttachment, __self__).__init__(
            'azure:compute/dataDiskAttachment:DataDiskAttachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

