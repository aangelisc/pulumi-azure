# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InferenceClusterSslArgs',
    'WorkspaceIdentityArgs',
]

@pulumi.input_type
class InferenceClusterSslArgs:
    def __init__(__self__, *,
                 cert: Optional[pulumi.Input[str]] = None,
                 cname: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cert: The certificate for the ssl configuration. Changing this forces a new Machine Learning Inference Cluster to be created.
        :param pulumi.Input[str] cname: The cname of the ssl configuration. Changing this forces a new Machine Learning Inference Cluster to be created.
        :param pulumi.Input[str] key: The key content for the ssl configuration. Changing this forces a new Machine Learning Inference Cluster to be created.
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if cname is not None:
            pulumi.set(__self__, "cname", cname)
        if key is not None:
            pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate for the ssl configuration. Changing this forces a new Machine Learning Inference Cluster to be created.
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def cname(self) -> Optional[pulumi.Input[str]]:
        """
        The cname of the ssl configuration. Changing this forces a new Machine Learning Inference Cluster to be created.
        """
        return pulumi.get(self, "cname")

    @cname.setter
    def cname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cname", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The key content for the ssl configuration. Changing this forces a new Machine Learning Inference Cluster to be created.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)


@pulumi.input_type
class WorkspaceIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The (Client) ID of the Service Principal.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant the Service Principal is assigned in.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The (Client) ID of the Service Principal.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Tenant the Service Principal is assigned in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


