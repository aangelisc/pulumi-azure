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

__all__ = ['EnterpriseDatabaseArgs', 'EnterpriseDatabase']

@pulumi.input_type
class EnterpriseDatabaseArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 client_protocol: Optional[pulumi.Input[str]] = None,
                 clustering_policy: Optional[pulumi.Input[str]] = None,
                 eviction_policy: Optional[pulumi.Input[str]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a EnterpriseDatabase resource.
        :param pulumi.Input[str] cluster_id: The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] client_protocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] clustering_policy: Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] eviction_policy: Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]] modules: A `module` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[int] port: TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if client_protocol is not None:
            pulumi.set(__self__, "client_protocol", client_protocol)
        if clustering_policy is not None:
            pulumi.set(__self__, "clustering_policy", clustering_policy)
        if eviction_policy is not None:
            pulumi.set(__self__, "eviction_policy", eviction_policy)
        if modules is not None:
            pulumi.set(__self__, "modules", modules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "client_protocol")

    @client_protocol.setter
    def client_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_protocol", value)

    @property
    @pulumi.getter(name="clusteringPolicy")
    def clustering_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "clustering_policy")

    @clustering_policy.setter
    def clustering_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clustering_policy", value)

    @property
    @pulumi.getter(name="evictionPolicy")
    def eviction_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "eviction_policy")

    @eviction_policy.setter
    def eviction_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eviction_policy", value)

    @property
    @pulumi.getter
    def modules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]]]:
        """
        A `module` block as defined below.
        """
        return pulumi.get(self, "modules")

    @modules.setter
    def modules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]]]):
        pulumi.set(self, "modules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class _EnterpriseDatabaseState:
    def __init__(__self__, *,
                 client_protocol: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 clustering_policy: Optional[pulumi.Input[str]] = None,
                 eviction_policy: Optional[pulumi.Input[str]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnterpriseDatabase resources.
        :param pulumi.Input[str] client_protocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] cluster_id: The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] clustering_policy: Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] eviction_policy: Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]] modules: A `module` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[int] port: TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        """
        if client_protocol is not None:
            pulumi.set(__self__, "client_protocol", client_protocol)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if clustering_policy is not None:
            pulumi.set(__self__, "clustering_policy", clustering_policy)
        if eviction_policy is not None:
            pulumi.set(__self__, "eviction_policy", eviction_policy)
        if modules is not None:
            pulumi.set(__self__, "modules", modules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "client_protocol")

    @client_protocol.setter
    def client_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_protocol", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="clusteringPolicy")
    def clustering_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "clustering_policy")

    @clustering_policy.setter
    def clustering_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clustering_policy", value)

    @property
    @pulumi.getter(name="evictionPolicy")
    def eviction_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "eviction_policy")

    @eviction_policy.setter
    def eviction_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eviction_policy", value)

    @property
    @pulumi.getter
    def modules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]]]:
        """
        A `module` block as defined below.
        """
        return pulumi.get(self, "modules")

    @modules.setter
    def modules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnterpriseDatabaseModuleArgs']]]]):
        pulumi.set(self, "modules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)


class EnterpriseDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_protocol: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 clustering_policy: Optional[pulumi.Input[str]] = None,
                 eviction_policy: Optional[pulumi.Input[str]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnterpriseDatabaseModuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Redis Enterprise Database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_enterprise_cluster = azure.redis.EnterpriseCluster("exampleEnterpriseCluster",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku_name="Enterprise_E20-4")
        example_enterprise_database = azure.redis.EnterpriseDatabase("exampleEnterpriseDatabase",
            resource_group_name=example_resource_group.name,
            cluster_id=example_enterprise_cluster.id)
        ```

        ## Import

        Redis Enterprise Databases can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:redis/enterpriseDatabase:EnterpriseDatabase example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1/databases/database1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_protocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] cluster_id: The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] clustering_policy: Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] eviction_policy: Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnterpriseDatabaseModuleArgs']]]] modules: A `module` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[int] port: TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnterpriseDatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Redis Enterprise Database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_enterprise_cluster = azure.redis.EnterpriseCluster("exampleEnterpriseCluster",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku_name="Enterprise_E20-4")
        example_enterprise_database = azure.redis.EnterpriseDatabase("exampleEnterpriseDatabase",
            resource_group_name=example_resource_group.name,
            cluster_id=example_enterprise_cluster.id)
        ```

        ## Import

        Redis Enterprise Databases can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:redis/enterpriseDatabase:EnterpriseDatabase example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1/databases/database1
        ```

        :param str resource_name: The name of the resource.
        :param EnterpriseDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnterpriseDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_protocol: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 clustering_policy: Optional[pulumi.Input[str]] = None,
                 eviction_policy: Optional[pulumi.Input[str]] = None,
                 modules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnterpriseDatabaseModuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = EnterpriseDatabaseArgs.__new__(EnterpriseDatabaseArgs)

            __props__.__dict__["client_protocol"] = client_protocol
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["clustering_policy"] = clustering_policy
            __props__.__dict__["eviction_policy"] = eviction_policy
            __props__.__dict__["modules"] = modules
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
        super(EnterpriseDatabase, __self__).__init__(
            'azure:redis/enterpriseDatabase:EnterpriseDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_protocol: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            clustering_policy: Optional[pulumi.Input[str]] = None,
            eviction_policy: Optional[pulumi.Input[str]] = None,
            modules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnterpriseDatabaseModuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'EnterpriseDatabase':
        """
        Get an existing EnterpriseDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_protocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] cluster_id: The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] clustering_policy: Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] eviction_policy: Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnterpriseDatabaseModuleArgs']]]] modules: A `module` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[int] port: TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnterpriseDatabaseState.__new__(_EnterpriseDatabaseState)

        __props__.__dict__["client_protocol"] = client_protocol
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["clustering_policy"] = clustering_policy
        __props__.__dict__["eviction_policy"] = eviction_policy
        __props__.__dict__["modules"] = modules
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["resource_group_name"] = resource_group_name
        return EnterpriseDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "client_protocol")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusteringPolicy")
    def clustering_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "clustering_policy")

    @property
    @pulumi.getter(name="evictionPolicy")
    def eviction_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "eviction_policy")

    @property
    @pulumi.getter
    def modules(self) -> pulumi.Output[Optional[Sequence['outputs.EnterpriseDatabaseModule']]]:
        """
        A `module` block as defined below.
        """
        return pulumi.get(self, "modules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "resource_group_name")

