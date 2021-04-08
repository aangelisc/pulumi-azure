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

__all__ = ['AssessmentArgs', 'Assessment']

@pulumi.input_type
class AssessmentArgs:
    def __init__(__self__, *,
                 assessment_policy_id: pulumi.Input[str],
                 status: pulumi.Input['AssessmentStatusArgs'],
                 target_resource_id: pulumi.Input[str],
                 additional_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Assessment resource.
        :param pulumi.Input[str] assessment_policy_id: The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        :param pulumi.Input['AssessmentStatusArgs'] status: A `status` block as defined below.
        :param pulumi.Input[str] target_resource_id: The ID of the target resource. Changing this forces a new security Assessment to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_data: A map of additional data to associate with the assessment.
        """
        pulumi.set(__self__, "assessment_policy_id", assessment_policy_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "target_resource_id", target_resource_id)
        if additional_data is not None:
            pulumi.set(__self__, "additional_data", additional_data)

    @property
    @pulumi.getter(name="assessmentPolicyId")
    def assessment_policy_id(self) -> pulumi.Input[str]:
        """
        The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        """
        return pulumi.get(self, "assessment_policy_id")

    @assessment_policy_id.setter
    def assessment_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "assessment_policy_id", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input['AssessmentStatusArgs']:
        """
        A `status` block as defined below.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input['AssessmentStatusArgs']):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the target resource. Changing this forces a new security Assessment to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_id", value)

    @property
    @pulumi.getter(name="additionalData")
    def additional_data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional data to associate with the assessment.
        """
        return pulumi.get(self, "additional_data")

    @additional_data.setter
    def additional_data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_data", value)


@pulumi.input_type
class _AssessmentState:
    def __init__(__self__, *,
                 additional_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 assessment_policy_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['AssessmentStatusArgs']] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Assessment resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_data: A map of additional data to associate with the assessment.
        :param pulumi.Input[str] assessment_policy_id: The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        :param pulumi.Input['AssessmentStatusArgs'] status: A `status` block as defined below.
        :param pulumi.Input[str] target_resource_id: The ID of the target resource. Changing this forces a new security Assessment to be created.
        """
        if additional_data is not None:
            pulumi.set(__self__, "additional_data", additional_data)
        if assessment_policy_id is not None:
            pulumi.set(__self__, "assessment_policy_id", assessment_policy_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if target_resource_id is not None:
            pulumi.set(__self__, "target_resource_id", target_resource_id)

    @property
    @pulumi.getter(name="additionalData")
    def additional_data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional data to associate with the assessment.
        """
        return pulumi.get(self, "additional_data")

    @additional_data.setter
    def additional_data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_data", value)

    @property
    @pulumi.getter(name="assessmentPolicyId")
    def assessment_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        """
        return pulumi.get(self, "assessment_policy_id")

    @assessment_policy_id.setter
    def assessment_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "assessment_policy_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['AssessmentStatusArgs']]:
        """
        A `status` block as defined below.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['AssessmentStatusArgs']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the target resource. Changing this forces a new security Assessment to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_resource_id", value)


class Assessment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 assessment_policy_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['AssessmentStatusArgs']]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages the Security Center Assessment for Azure Security Center.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            address_spaces=["10.0.0.0/16"])
        internal = azure.network.Subnet("internal",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"])
        example_linux_virtual_machine_scale_set = azure.compute.LinuxVirtualMachineScaleSet("exampleLinuxVirtualMachineScaleSet",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku="Standard_F2",
            instances=1,
            admin_username="adminuser",
            admin_ssh_keys=[azure.compute.LinuxVirtualMachineScaleSetAdminSshKeyArgs(
                username="adminuser",
                public_key=(lambda path: open(path).read())("~/.ssh/id_rsa.pub"),
            )],
            source_image_reference=azure.compute.LinuxVirtualMachineScaleSetSourceImageReferenceArgs(
                publisher="Canonical",
                offer="UbuntuServer",
                sku="16.04-LTS",
                version="latest",
            ),
            os_disk=azure.compute.LinuxVirtualMachineScaleSetOsDiskArgs(
                storage_account_type="Standard_LRS",
                caching="ReadWrite",
            ),
            network_interfaces=[azure.compute.LinuxVirtualMachineScaleSetNetworkInterfaceArgs(
                name="example",
                primary=True,
                ip_configurations=[{
                    "name": "internal",
                    "primary": True,
                    "subnet_id": internal.id,
                }],
            )])
        example_assessment_policy = azure.securitycenter.AssessmentPolicy("exampleAssessmentPolicy",
            display_name="Test Display Name",
            severity="Medium",
            description="Test Description")
        example_assessment = azure.securitycenter.Assessment("exampleAssessment",
            assessment_policy_id=example_assessment_policy.id,
            target_resource_id=example_linux_virtual_machine_scale_set.id,
            status=azure.securitycenter.AssessmentStatusArgs(
                code="Healthy",
            ))
        ```

        ## Import

        Security Assessment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:securitycenter/assessment:Assessment example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_data: A map of additional data to associate with the assessment.
        :param pulumi.Input[str] assessment_policy_id: The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        :param pulumi.Input[pulumi.InputType['AssessmentStatusArgs']] status: A `status` block as defined below.
        :param pulumi.Input[str] target_resource_id: The ID of the target resource. Changing this forces a new security Assessment to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssessmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the Security Center Assessment for Azure Security Center.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            address_spaces=["10.0.0.0/16"])
        internal = azure.network.Subnet("internal",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"])
        example_linux_virtual_machine_scale_set = azure.compute.LinuxVirtualMachineScaleSet("exampleLinuxVirtualMachineScaleSet",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku="Standard_F2",
            instances=1,
            admin_username="adminuser",
            admin_ssh_keys=[azure.compute.LinuxVirtualMachineScaleSetAdminSshKeyArgs(
                username="adminuser",
                public_key=(lambda path: open(path).read())("~/.ssh/id_rsa.pub"),
            )],
            source_image_reference=azure.compute.LinuxVirtualMachineScaleSetSourceImageReferenceArgs(
                publisher="Canonical",
                offer="UbuntuServer",
                sku="16.04-LTS",
                version="latest",
            ),
            os_disk=azure.compute.LinuxVirtualMachineScaleSetOsDiskArgs(
                storage_account_type="Standard_LRS",
                caching="ReadWrite",
            ),
            network_interfaces=[azure.compute.LinuxVirtualMachineScaleSetNetworkInterfaceArgs(
                name="example",
                primary=True,
                ip_configurations=[{
                    "name": "internal",
                    "primary": True,
                    "subnet_id": internal.id,
                }],
            )])
        example_assessment_policy = azure.securitycenter.AssessmentPolicy("exampleAssessmentPolicy",
            display_name="Test Display Name",
            severity="Medium",
            description="Test Description")
        example_assessment = azure.securitycenter.Assessment("exampleAssessment",
            assessment_policy_id=example_assessment_policy.id,
            target_resource_id=example_linux_virtual_machine_scale_set.id,
            status=azure.securitycenter.AssessmentStatusArgs(
                code="Healthy",
            ))
        ```

        ## Import

        Security Assessment can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:securitycenter/assessment:Assessment example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AssessmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssessmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 assessment_policy_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['AssessmentStatusArgs']]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AssessmentArgs.__new__(AssessmentArgs)

            __props__.__dict__["additional_data"] = additional_data
            if assessment_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'assessment_policy_id'")
            __props__.__dict__["assessment_policy_id"] = assessment_policy_id
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            if target_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__.__dict__["target_resource_id"] = target_resource_id
        super(Assessment, __self__).__init__(
            'azure:securitycenter/assessment:Assessment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            assessment_policy_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[pulumi.InputType['AssessmentStatusArgs']]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None) -> 'Assessment':
        """
        Get an existing Assessment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_data: A map of additional data to associate with the assessment.
        :param pulumi.Input[str] assessment_policy_id: The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        :param pulumi.Input[pulumi.InputType['AssessmentStatusArgs']] status: A `status` block as defined below.
        :param pulumi.Input[str] target_resource_id: The ID of the target resource. Changing this forces a new security Assessment to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AssessmentState.__new__(_AssessmentState)

        __props__.__dict__["additional_data"] = additional_data
        __props__.__dict__["assessment_policy_id"] = assessment_policy_id
        __props__.__dict__["status"] = status
        __props__.__dict__["target_resource_id"] = target_resource_id
        return Assessment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalData")
    def additional_data(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of additional data to associate with the assessment.
        """
        return pulumi.get(self, "additional_data")

    @property
    @pulumi.getter(name="assessmentPolicyId")
    def assessment_policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
        """
        return pulumi.get(self, "assessment_policy_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.AssessmentStatus']:
        """
        A `status` block as defined below.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the target resource. Changing this forces a new security Assessment to be created.
        """
        return pulumi.get(self, "target_resource_id")

