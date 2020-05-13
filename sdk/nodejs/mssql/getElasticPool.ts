// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing SQL elastic pool.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.mssql.getElasticPool({
 *     name: "mssqlelasticpoolname",
 *     resourceGroupName: "example-resources",
 *     serverName: "example-sql-server",
 * });
 * export const elasticpoolId = example.then(example => example.id);
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/mssql_elasticpool.html.markdown.
 */
export function getElasticPool(args: GetElasticPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetElasticPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:mssql/getElasticPool:getElasticPool", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

/**
 * A collection of arguments for invoking getElasticPool.
 */
export interface GetElasticPoolArgs {
    /**
     * The name of the elastic pool.
     */
    readonly name: string;
    /**
     * The name of the resource group which contains the elastic pool.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the SQL Server which contains the elastic pool.
     */
    readonly serverName: string;
}

/**
 * A collection of values returned by getElasticPool.
 */
export interface GetElasticPoolResult {
    /**
     * The license type to apply for this database.
     */
    readonly licenseType: string;
    /**
     * Specifies the supported Azure location where the resource exists.
     */
    readonly location: string;
    /**
     * The max data size of the elastic pool in bytes.
     */
    readonly maxSizeBytes: number;
    /**
     * The max data size of the elastic pool in gigabytes.
     */
    readonly maxSizeGb: number;
    readonly name: string;
    /**
     * The maximum capacity any one database can consume.
     */
    readonly perDbMaxCapacity: number;
    /**
     * The minimum capacity all databases are guaranteed.
     */
    readonly perDbMinCapacity: number;
    readonly resourceGroupName: string;
    readonly serverName: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * Whether or not this elastic pool is zone redundant.
     */
    readonly zoneRedundant: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
