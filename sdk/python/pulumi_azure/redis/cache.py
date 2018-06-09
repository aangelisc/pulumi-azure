# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Cache(pulumi.CustomResource):
    """
    Manages a Redis Cache.
    """
    def __init__(__self__, __name__, __opts__=None, capacity=None, enable_non_ssl_port=None, family=None, location=None, name=None, patch_schedules=None, private_static_ip_address=None, redis_configuration=None, resource_group_name=None, shard_count=None, sku_name=None, subnet_id=None, tags=None):
        """Create a Cache resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not capacity:
            raise TypeError('Missing required property capacity')
        elif not isinstance(capacity, int):
            raise TypeError('Expected property capacity to be a int')
        __self__.capacity = capacity
        """
        The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
        """
        __props__['capacity'] = capacity

        if enable_non_ssl_port and not isinstance(enable_non_ssl_port, bool):
            raise TypeError('Expected property enable_non_ssl_port to be a bool')
        __self__.enable_non_ssl_port = enable_non_ssl_port
        """
        Enable the non-SSL port (6789) - disabled by default.
        """
        __props__['enableNonSslPort'] = enable_non_ssl_port

        if not family:
            raise TypeError('Missing required property family')
        elif not isinstance(family, basestring):
            raise TypeError('Expected property family to be a basestring')
        __self__.family = family
        """
        The SKU family to use. Valid values are `C` and `P`, where C = Basic/Standard, P = Premium.
        """
        __props__['family'] = family

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The location of the resource group.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the Redis instance. Changing this forces a
        new resource to be created.
        """
        __props__['name'] = name

        if patch_schedules and not isinstance(patch_schedules, list):
            raise TypeError('Expected property patch_schedules to be a list')
        __self__.patch_schedules = patch_schedules
        """
        A list of `patch_schedule` blocks as defined below - only available for Premium SKU's.
        """
        __props__['patchSchedules'] = patch_schedules

        if private_static_ip_address and not isinstance(private_static_ip_address, basestring):
            raise TypeError('Expected property private_static_ip_address to be a basestring')
        __self__.private_static_ip_address = private_static_ip_address
        """
        The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
        """
        __props__['privateStaticIpAddress'] = private_static_ip_address

        if not redis_configuration:
            raise TypeError('Missing required property redis_configuration')
        elif not isinstance(redis_configuration, dict):
            raise TypeError('Expected property redis_configuration to be a dict')
        __self__.redis_configuration = redis_configuration
        """
        A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
        """
        __props__['redisConfiguration'] = redis_configuration

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to
        create the Redis instance.
        """
        __props__['resourceGroupName'] = resource_group_name

        if shard_count and not isinstance(shard_count, int):
            raise TypeError('Expected property shard_count to be a int')
        __self__.shard_count = shard_count
        """
        *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
        """
        __props__['shardCount'] = shard_count

        if not sku_name:
            raise TypeError('Missing required property sku_name')
        elif not isinstance(sku_name, basestring):
            raise TypeError('Expected property sku_name to be a basestring')
        __self__.sku_name = sku_name
        """
        The SKU of Redis to use - can be either Basic, Standard or Premium.
        """
        __props__['skuName'] = sku_name

        if subnet_id and not isinstance(subnet_id, basestring):
            raise TypeError('Expected property subnet_id to be a basestring')
        __self__.subnet_id = subnet_id
        """
        The ID of the Subnet within which the Redis Cache should be deployed. Changing this forces a new resource to be created.
        """
        __props__['subnetId'] = subnet_id

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        __props__['tags'] = tags

        __self__.hostname = pulumi.runtime.UNKNOWN
        """
        The Hostname of the Redis Instance
        """
        __self__.port = pulumi.runtime.UNKNOWN
        """
        The non-SSL Port of the Redis Instance
        """
        __self__.primary_access_key = pulumi.runtime.UNKNOWN
        """
        The Primary Access Key for the Redis Instance
        """
        __self__.secondary_access_key = pulumi.runtime.UNKNOWN
        """
        The Secondary Access Key for the Redis Instance
        """
        __self__.ssl_port = pulumi.runtime.UNKNOWN
        """
        The SSL Port of the Redis Instance
        """

        super(Cache, __self__).__init__(
            'azure:redis/cache:Cache',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'capacity' in outs:
            self.capacity = outs['capacity']
        if 'enableNonSslPort' in outs:
            self.enable_non_ssl_port = outs['enableNonSslPort']
        if 'family' in outs:
            self.family = outs['family']
        if 'hostname' in outs:
            self.hostname = outs['hostname']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'patchSchedules' in outs:
            self.patch_schedules = outs['patchSchedules']
        if 'port' in outs:
            self.port = outs['port']
        if 'primaryAccessKey' in outs:
            self.primary_access_key = outs['primaryAccessKey']
        if 'privateStaticIpAddress' in outs:
            self.private_static_ip_address = outs['privateStaticIpAddress']
        if 'redisConfiguration' in outs:
            self.redis_configuration = outs['redisConfiguration']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'secondaryAccessKey' in outs:
            self.secondary_access_key = outs['secondaryAccessKey']
        if 'shardCount' in outs:
            self.shard_count = outs['shardCount']
        if 'skuName' in outs:
            self.sku_name = outs['skuName']
        if 'sslPort' in outs:
            self.ssl_port = outs['sslPort']
        if 'subnetId' in outs:
            self.subnet_id = outs['subnetId']
        if 'tags' in outs:
            self.tags = outs['tags']
