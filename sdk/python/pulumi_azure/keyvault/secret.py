# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretArgs', 'Secret']

@pulumi.input_type
class SecretArgs:
    def __init__(__self__, *,
                 key_vault_id: pulumi.Input[str],
                 value: pulumi.Input[str],
                 content_type: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before_date: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Secret resource.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] value: Specifies the value of the Key Vault Secret.
        :param pulumi.Input[str] content_type: Specifies the content type for the Key Vault Secret.
        :param pulumi.Input[str] expiration_date: Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] not_before_date: Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "key_vault_id", key_vault_id)
        pulumi.set(__self__, "value", value)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if not_before_date is not None:
            pulumi.set(__self__, "not_before_date", not_before_date)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Input[str]:
        """
        The ID of the Key Vault where the Secret should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Specifies the value of the Key Vault Secret.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the content type for the Key Vault Secret.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notBeforeDate")
    def not_before_date(self) -> Optional[pulumi.Input[str]]:
        """
        Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "not_before_date")

    @not_before_date.setter
    def not_before_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_before_date", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SecretState:
    def __init__(__self__, *,
                 content_type: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before_date: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 versionless_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Secret resources.
        :param pulumi.Input[str] content_type: Specifies the content type for the Key Vault Secret.
        :param pulumi.Input[str] expiration_date: Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] not_before_date: Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] value: Specifies the value of the Key Vault Secret.
        :param pulumi.Input[str] version: The current version of the Key Vault Secret.
        :param pulumi.Input[str] versionless_id: The Base ID of the Key Vault Secret.
        """
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if not_before_date is not None:
            pulumi.set(__self__, "not_before_date", not_before_date)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if versionless_id is not None:
            pulumi.set(__self__, "versionless_id", versionless_id)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the content type for the Key Vault Secret.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Key Vault where the Secret should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notBeforeDate")
    def not_before_date(self) -> Optional[pulumi.Input[str]]:
        """
        Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "not_before_date")

    @not_before_date.setter
    def not_before_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_before_date", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the value of the Key Vault Secret.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The current version of the Key Vault Secret.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="versionlessId")
    def versionless_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Base ID of the Key Vault Secret.
        """
        return pulumi.get(self, "versionless_id")

    @versionless_id.setter
    def versionless_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "versionless_id", value)


class Secret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before_date: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Key Vault Secret.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            soft_delete_retention_days=7,
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                key_permissions=[
                    "create",
                    "get",
                ],
                secret_permissions=[
                    "set",
                    "get",
                    "delete",
                    "purge",
                    "recover",
                ],
            )])
        example_secret = azure.keyvault.Secret("exampleSecret",
            value="szechuan",
            key_vault_id=example_key_vault.id)
        ```

        ## Import

        Key Vault Secrets which are Enabled can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:keyvault/secret:Secret example "https://example-keyvault.vault.azure.net/secrets/example/fdf067c93bbb4b22bff4d8b7a9a56217"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: Specifies the content type for the Key Vault Secret.
        :param pulumi.Input[str] expiration_date: Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] not_before_date: Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] value: Specifies the value of the Key Vault Secret.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Key Vault Secret.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            soft_delete_retention_days=7,
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                key_permissions=[
                    "create",
                    "get",
                ],
                secret_permissions=[
                    "set",
                    "get",
                    "delete",
                    "purge",
                    "recover",
                ],
            )])
        example_secret = azure.keyvault.Secret("exampleSecret",
            value="szechuan",
            key_vault_id=example_key_vault.id)
        ```

        ## Import

        Key Vault Secrets which are Enabled can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:keyvault/secret:Secret example "https://example-keyvault.vault.azure.net/secrets/example/fdf067c93bbb4b22bff4d8b7a9a56217"
        ```

        :param str resource_name: The name of the resource.
        :param SecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before_date: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 value: Optional[pulumi.Input[str]] = None,
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
            __props__ = SecretArgs.__new__(SecretArgs)

            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["expiration_date"] = expiration_date
            if key_vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_vault_id'")
            __props__.__dict__["key_vault_id"] = key_vault_id
            __props__.__dict__["name"] = name
            __props__.__dict__["not_before_date"] = not_before_date
            __props__.__dict__["tags"] = tags
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["version"] = None
            __props__.__dict__["versionless_id"] = None
        super(Secret, __self__).__init__(
            'azure:keyvault/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            not_before_date: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            value: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None,
            versionless_id: Optional[pulumi.Input[str]] = None) -> 'Secret':
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: Specifies the content type for the Key Vault Secret.
        :param pulumi.Input[str] expiration_date: Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] not_before_date: Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] value: Specifies the value of the Key Vault Secret.
        :param pulumi.Input[str] version: The current version of the Key Vault Secret.
        :param pulumi.Input[str] versionless_id: The Base ID of the Key Vault Secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretState.__new__(_SecretState)

        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["key_vault_id"] = key_vault_id
        __props__.__dict__["name"] = name
        __props__.__dict__["not_before_date"] = not_before_date
        __props__.__dict__["tags"] = tags
        __props__.__dict__["value"] = value
        __props__.__dict__["version"] = version
        __props__.__dict__["versionless_id"] = versionless_id
        return Secret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the content type for the Key Vault Secret.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the Key Vault where the Secret should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notBeforeDate")
    def not_before_date(self) -> pulumi.Output[Optional[str]]:
        """
        Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "not_before_date")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Specifies the value of the Key Vault Secret.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The current version of the Key Vault Secret.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionlessId")
    def versionless_id(self) -> pulumi.Output[str]:
        """
        The Base ID of the Key Vault Secret.
        """
        return pulumi.get(self, "versionless_id")

