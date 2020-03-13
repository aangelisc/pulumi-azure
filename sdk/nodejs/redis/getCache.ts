// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Redis Cache
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/redis_cache.html.markdown.
 */
export function getCache(args: GetCacheArgs, opts?: pulumi.InvokeOptions): Promise<GetCacheResult> & GetCacheResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetCacheResult> = pulumi.runtime.invoke("azure:redis/getCache:getCache", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "zones": args.zones,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCache.
 */
export interface GetCacheArgs {
    /**
     * The name of the Redis cache
     */
    readonly name: string;
    /**
     * The name of the resource group the Redis cache instance is located in.
     */
    readonly resourceGroupName: string;
    readonly zones?: string[];
}

/**
 * A collection of values returned by getCache.
 */
export interface GetCacheResult {
    /**
     * The size of the Redis Cache deployed.
     */
    readonly capacity: number;
    /**
     * Whether the SSL port is enabled.
     */
    readonly enableNonSslPort: boolean;
    /**
     * The SKU family/pricing group used. Possible values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    readonly family: string;
    /**
     * The Hostname of the Redis Instance
     */
    readonly hostname: string;
    /**
     * The location of the Redis Cache.
     */
    readonly location: string;
    /**
     * The minimum TLS version.
     */
    readonly minimumTlsVersion: string;
    readonly name: string;
    /**
     * A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
     */
    readonly patchSchedules: outputs.redis.GetCachePatchSchedule[];
    /**
     * The non-SSL Port of the Redis Instance
     */
    readonly port: number;
    /**
     * The Primary Access Key for the Redis Instance
     */
    readonly primaryAccessKey: string;
    /**
     * The primary connection string of the Redis Instance.
     */
    readonly primaryConnectionString: string;
    readonly privateStaticIpAddress: string;
    /**
     * A `redisConfiguration` block as defined below.
     */
    readonly redisConfigurations: outputs.redis.GetCacheRedisConfiguration[];
    readonly resourceGroupName: string;
    /**
     * The Secondary Access Key for the Redis Instance
     */
    readonly secondaryAccessKey: string;
    /**
     * The secondary connection string of the Redis Instance.
     */
    readonly secondaryConnectionString: string;
    readonly shardCount: number;
    /**
     * The SKU of Redis used. Possible values are `Basic`, `Standard` and `Premium`.
     */
    readonly skuName: string;
    /**
     * The SSL Port of the Redis Instance
     */
    readonly sslPort: number;
    readonly subnetId: string;
    readonly tags: {[key: string]: string};
    readonly zones: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
