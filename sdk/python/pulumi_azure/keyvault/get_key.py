# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetKeyResult',
    'AwaitableGetKeyResult',
    'get_key',
]

@pulumi.output_type
class GetKeyResult:
    """
    A collection of values returned by getKey.
    """
    def __init__(__self__, e=None, id=None, key_opts=None, key_size=None, key_type=None, key_vault_id=None, n=None, name=None, tags=None, version=None, versionless_id=None):
        if e and not isinstance(e, str):
            raise TypeError("Expected argument 'e' to be a str")
        pulumi.set(__self__, "e", e)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_opts and not isinstance(key_opts, list):
            raise TypeError("Expected argument 'key_opts' to be a list")
        pulumi.set(__self__, "key_opts", key_opts)
        if key_size and not isinstance(key_size, int):
            raise TypeError("Expected argument 'key_size' to be a int")
        pulumi.set(__self__, "key_size", key_size)
        if key_type and not isinstance(key_type, str):
            raise TypeError("Expected argument 'key_type' to be a str")
        pulumi.set(__self__, "key_type", key_type)
        if key_vault_id and not isinstance(key_vault_id, str):
            raise TypeError("Expected argument 'key_vault_id' to be a str")
        pulumi.set(__self__, "key_vault_id", key_vault_id)
        if n and not isinstance(n, str):
            raise TypeError("Expected argument 'n' to be a str")
        pulumi.set(__self__, "n", n)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if versionless_id and not isinstance(versionless_id, str):
            raise TypeError("Expected argument 'versionless_id' to be a str")
        pulumi.set(__self__, "versionless_id", versionless_id)

    @property
    @pulumi.getter
    def e(self) -> str:
        """
        The RSA public exponent of this Key Vault Key.
        """
        return pulumi.get(self, "e")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyOpts")
    def key_opts(self) -> Sequence[str]:
        """
        A list of JSON web key operations assigned to this Key Vault Key
        """
        return pulumi.get(self, "key_opts")

    @property
    @pulumi.getter(name="keySize")
    def key_size(self) -> int:
        """
        Specifies the Size of this Key Vault Key.
        """
        return pulumi.get(self, "key_size")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> str:
        """
        Specifies the Key Type of this Key Vault Key
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> str:
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def n(self) -> str:
        """
        The RSA modulus of this Key Vault Key.
        """
        return pulumi.get(self, "n")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to this Key Vault Key.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The current version of the Key Vault Key.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionlessId")
    def versionless_id(self) -> str:
        """
        The Base ID of the Key Vault Key.
        """
        return pulumi.get(self, "versionless_id")


class AwaitableGetKeyResult(GetKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyResult(
            e=self.e,
            id=self.id,
            key_opts=self.key_opts,
            key_size=self.key_size,
            key_type=self.key_type,
            key_vault_id=self.key_vault_id,
            n=self.n,
            name=self.name,
            tags=self.tags,
            version=self.version,
            versionless_id=self.versionless_id)


def get_key(key_vault_id: Optional[str] = None,
            name: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyResult:
    """
    Use this data source to access information about an existing Key Vault Key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.keyvault.get_key(name="secret-sauce",
        key_vault_id=data["azurerm_key_vault"]["existing"]["id"])
    pulumi.export("keyType", example.key_type)
    ```


    :param str key_vault_id: Specifies the ID of the Key Vault instance where the Secret resides, available on the `keyvault.KeyVault` Data Source / Resource.
    :param str name: Specifies the name of the Key Vault Key.
    """
    __args__ = dict()
    __args__['keyVaultId'] = key_vault_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:keyvault/getKey:getKey', __args__, opts=opts, typ=GetKeyResult).value

    return AwaitableGetKeyResult(
        e=__ret__.e,
        id=__ret__.id,
        key_opts=__ret__.key_opts,
        key_size=__ret__.key_size,
        key_type=__ret__.key_type,
        key_vault_id=__ret__.key_vault_id,
        n=__ret__.n,
        name=__ret__.name,
        tags=__ret__.tags,
        version=__ret__.version,
        versionless_id=__ret__.versionless_id)
