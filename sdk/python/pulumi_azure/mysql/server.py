# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Server(pulumi.CustomResource):
    """
    Manages a MySQL Server.
    """
    def __init__(__self__, __name__, __opts__=None, administrator_login=None, administrator_login_password=None, location=None, name=None, resource_group_name=None, sku=None, ssl_enforcement=None, storage_profile=None, tags=None, version=None):
        """Create a Server resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not administrator_login:
            raise TypeError('Missing required property administrator_login')
        elif not isinstance(administrator_login, basestring):
            raise TypeError('Expected property administrator_login to be a basestring')
        __self__.administrator_login = administrator_login
        """
        The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
        """
        __props__['administratorLogin'] = administrator_login

        if not administrator_login_password:
            raise TypeError('Missing required property administrator_login_password')
        elif not isinstance(administrator_login_password, basestring):
            raise TypeError('Expected property administrator_login_password to be a basestring')
        __self__.administrator_login_password = administrator_login_password
        """
        The Password associated with the `administrator_login` for the MySQL Server.
        """
        __props__['administratorLoginPassword'] = administrator_login_password

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the MySQL Server.
        """
        __props__['resourceGroupName'] = resource_group_name

        if not sku:
            raise TypeError('Missing required property sku')
        elif not isinstance(sku, dict):
            raise TypeError('Expected property sku to be a dict')
        __self__.sku = sku
        """
        A `sku` block as defined below.
        """
        __props__['sku'] = sku

        if not ssl_enforcement:
            raise TypeError('Missing required property ssl_enforcement')
        elif not isinstance(ssl_enforcement, basestring):
            raise TypeError('Expected property ssl_enforcement to be a basestring')
        __self__.ssl_enforcement = ssl_enforcement
        """
        Specifies if SSL should be enforced on connections. Possible values are `Enforced` and `Disabled`.
        """
        __props__['sslEnforcement'] = ssl_enforcement

        if not storage_profile:
            raise TypeError('Missing required property storage_profile')
        elif not isinstance(storage_profile, dict):
            raise TypeError('Expected property storage_profile to be a dict')
        __self__.storage_profile = storage_profile
        """
        A `storage_profile` block as defined below.
        """
        __props__['storageProfile'] = storage_profile

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        if not version:
            raise TypeError('Missing required property version')
        elif not isinstance(version, basestring):
            raise TypeError('Expected property version to be a basestring')
        __self__.version = version
        """
        Specifies the version of MySQL to use. Valid values are `5.6` and `5.7`. Changing this forces a new resource to be created.
        """
        __props__['version'] = version

        __self__.fqdn = pulumi.runtime.UNKNOWN
        """
        The FQDN of the MySQL Server.
        """

        super(Server, __self__).__init__(
            'azure:mysql/server:Server',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'administratorLogin' in outs:
            self.administrator_login = outs['administratorLogin']
        if 'administratorLoginPassword' in outs:
            self.administrator_login_password = outs['administratorLoginPassword']
        if 'fqdn' in outs:
            self.fqdn = outs['fqdn']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'sku' in outs:
            self.sku = outs['sku']
        if 'sslEnforcement' in outs:
            self.ssl_enforcement = outs['sslEnforcement']
        if 'storageProfile' in outs:
            self.storage_profile = outs['storageProfile']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'version' in outs:
            self.version = outs['version']
