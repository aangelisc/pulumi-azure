# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EncryptionScopeArgs', 'EncryptionScope']

@pulumi.input_type
class EncryptionScopeArgs:
    def __init__(__self__, *,
                 source: pulumi.Input[str],
                 storage_account_id: pulumi.Input[str],
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EncryptionScope resource.
        :param pulumi.Input[str] source: The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        :param pulumi.Input[str] key_vault_key_id: The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        :param pulumi.Input[str] name: The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        """
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "storage_account_id", storage_account_id)
        if key_vault_key_id is not None:
            pulumi.set(__self__, "key_vault_key_id", key_vault_key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Input[str]:
        """
        The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_id", value)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        """
        return pulumi.get(self, "key_vault_key_id")

    @key_vault_key_id.setter
    def key_vault_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _EncryptionScopeState:
    def __init__(__self__, *,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EncryptionScope resources.
        :param pulumi.Input[str] key_vault_key_id: The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        :param pulumi.Input[str] name: The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        :param pulumi.Input[str] source: The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        """
        if key_vault_key_id is not None:
            pulumi.set(__self__, "key_vault_key_id", key_vault_key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        """
        return pulumi.get(self, "key_vault_key_id")

    @key_vault_key_id.setter
    def key_vault_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)


class EncryptionScope(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Storage Encryption Scope.

        > **Note:** Storage Encryption Scopes are in Preview [more information can be found here](https://docs.microsoft.com/en-us/azure/storage/blobs/encryption-scope-manage).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            identity=azure.storage.AccountIdentityArgs(
                type="SystemAssigned",
            ))
        example_encryption_scope = azure.storage.EncryptionScope("exampleEncryptionScope",
            storage_account_id=example_account.id,
            source="Microsoft.Storage")
        ```

        ## Import

        Storage Encryption Scopes can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:storage/encryptionScope:EncryptionScope example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/encryptionScopes/scope1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_key_id: The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        :param pulumi.Input[str] name: The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        :param pulumi.Input[str] source: The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EncryptionScopeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Storage Encryption Scope.

        > **Note:** Storage Encryption Scopes are in Preview [more information can be found here](https://docs.microsoft.com/en-us/azure/storage/blobs/encryption-scope-manage).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            identity=azure.storage.AccountIdentityArgs(
                type="SystemAssigned",
            ))
        example_encryption_scope = azure.storage.EncryptionScope("exampleEncryptionScope",
            storage_account_id=example_account.id,
            source="Microsoft.Storage")
        ```

        ## Import

        Storage Encryption Scopes can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:storage/encryptionScope:EncryptionScope example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/encryptionScopes/scope1
        ```

        :param str resource_name: The name of the resource.
        :param EncryptionScopeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EncryptionScopeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = EncryptionScopeArgs.__new__(EncryptionScopeArgs)

            __props__.__dict__["key_vault_key_id"] = key_vault_key_id
            __props__.__dict__["name"] = name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            if storage_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_id'")
            __props__.__dict__["storage_account_id"] = storage_account_id
        super(EncryptionScope, __self__).__init__(
            'azure:storage/encryptionScope:EncryptionScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_vault_key_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None) -> 'EncryptionScope':
        """
        Get an existing EncryptionScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_key_id: The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        :param pulumi.Input[str] name: The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        :param pulumi.Input[str] source: The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EncryptionScopeState.__new__(_EncryptionScopeState)

        __props__.__dict__["key_vault_key_id"] = key_vault_key_id
        __props__.__dict__["name"] = name
        __props__.__dict__["source"] = source
        __props__.__dict__["storage_account_id"] = storage_account_id
        return EncryptionScope(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Key Vault Key. Required when `source` is `Microsoft.KeyVault`.
        """
        return pulumi.get(self, "key_vault_key_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Storage Encryption Scope. Changing this forces a new Storage Encryption Scope to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        The source of the Storage Encryption Scope. Possible values are `Microsoft.KeyVault` and `Microsoft.Storage`.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the Storage Account where this Storage Encryption Scope is created. Changing this forces a new Storage Encryption Scope to be created.
        """
        return pulumi.get(self, "storage_account_id")

