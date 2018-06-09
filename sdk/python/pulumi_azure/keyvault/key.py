# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Key(pulumi.CustomResource):
    """
    Manages a Key Vault Key.
    """
    def __init__(__self__, __name__, __opts__=None, key_opts=None, key_size=None, key_type=None, name=None, tags=None, vault_uri=None):
        """Create a Key resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not key_opts:
            raise TypeError('Missing required property key_opts')
        elif not isinstance(key_opts, list):
            raise TypeError('Expected property key_opts to be a list')
        __self__.key_opts = key_opts
        """
        A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
        """
        __props__['keyOpts'] = key_opts

        if not key_size:
            raise TypeError('Missing required property key_size')
        elif not isinstance(key_size, int):
            raise TypeError('Expected property key_size to be a int')
        __self__.key_size = key_size
        """
        Specifies the Size of the Key to create in bytes. For example, 1024 or 2048. Changing this forces a new resource to be created.
        """
        __props__['keySize'] = key_size

        if not key_type:
            raise TypeError('Missing required property key_type')
        elif not isinstance(key_type, basestring):
            raise TypeError('Expected property key_type to be a basestring')
        __self__.key_type = key_type
        """
        Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
        """
        __props__['keyType'] = key_type

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        if not vault_uri:
            raise TypeError('Missing required property vault_uri')
        elif not isinstance(vault_uri, basestring):
            raise TypeError('Expected property vault_uri to be a basestring')
        __self__.vault_uri = vault_uri
        """
        Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
        """
        __props__['vaultUri'] = vault_uri

        __self__.e = pulumi.runtime.UNKNOWN
        """
        The RSA public exponent of this Key Vault Key.
        """
        __self__.n = pulumi.runtime.UNKNOWN
        """
        The RSA modulus of this Key Vault Key.
        """
        __self__.version = pulumi.runtime.UNKNOWN
        """
        The current version of the Key Vault Key.
        """

        super(Key, __self__).__init__(
            'azure:keyvault/key:Key',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'e' in outs:
            self.e = outs['e']
        if 'keyOpts' in outs:
            self.key_opts = outs['keyOpts']
        if 'keySize' in outs:
            self.key_size = outs['keySize']
        if 'keyType' in outs:
            self.key_type = outs['keyType']
        if 'n' in outs:
            self.n = outs['n']
        if 'name' in outs:
            self.name = outs['name']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'vaultUri' in outs:
            self.vault_uri = outs['vaultUri']
        if 'version' in outs:
            self.version = outs['version']
