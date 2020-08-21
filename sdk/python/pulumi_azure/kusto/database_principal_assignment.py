# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['DatabasePrincipalAssignment']


class DatabasePrincipalAssignment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Database Principal Assignment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        rg = azure.core.ResourceGroup("rg", location="East US")
        example_cluster = azure.kusto.Cluster("exampleCluster",
            location=rg.location,
            resource_group_name=rg.name,
            sku=azure.kusto.ClusterSkuArgs(
                name="Standard_D13_v2",
                capacity=2,
            ))
        example_database = azure.kusto.Database("exampleDatabase",
            resource_group_name=rg.name,
            location=rg.location,
            cluster_name=example_cluster.name,
            hot_cache_period="P7D",
            soft_delete_period="P31D")
        example_database_principal_assignment = azure.kusto.DatabasePrincipalAssignment("exampleDatabasePrincipalAssignment",
            resource_group_name=rg.name,
            cluster_name=example_cluster.name,
            database_name=example_database.name,
            tenant_id=current.tenant_id,
            principal_id=current.client_id,
            principal_type="App",
            role="Viewer")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the database in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_id: The object id of the principal. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_type: The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The tenant id in which the principal resides. Changing this forces a new resource to be created.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['name'] = name
            if principal_id is None:
                raise TypeError("Missing required property 'principal_id'")
            __props__['principal_id'] = principal_id
            if principal_type is None:
                raise TypeError("Missing required property 'principal_type'")
            __props__['principal_type'] = principal_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if tenant_id is None:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
            __props__['principal_name'] = None
            __props__['tenant_name'] = None
        super(DatabasePrincipalAssignment, __self__).__init__(
            'azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            principal_id: Optional[pulumi.Input[str]] = None,
            principal_name: Optional[pulumi.Input[str]] = None,
            principal_type: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            tenant_name: Optional[pulumi.Input[str]] = None) -> 'DatabasePrincipalAssignment':
        """
        Get an existing DatabasePrincipalAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] database_name: The name of the database in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_id: The object id of the principal. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_name: The name of the principal.
        :param pulumi.Input[str] principal_type: The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role: The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The tenant id in which the principal resides. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_name: The name of the tenant.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_name"] = cluster_name
        __props__["database_name"] = database_name
        __props__["name"] = name
        __props__["principal_id"] = principal_id
        __props__["principal_name"] = principal_name
        __props__["principal_type"] = principal_type
        __props__["resource_group_name"] = resource_group_name
        __props__["role"] = role
        __props__["tenant_id"] = tenant_id
        __props__["tenant_name"] = tenant_name
        return DatabasePrincipalAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        The name of the database in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The object id of the principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> str:
        """
        The name of the principal.
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> str:
        """
        The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The database role assigned to the principal. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User` and `Viewer`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant id in which the principal resides. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> str:
        """
        The name of the tenant.
        """
        return pulumi.get(self, "tenant_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

