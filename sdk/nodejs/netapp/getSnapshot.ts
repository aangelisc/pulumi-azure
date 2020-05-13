// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Uses this data source to access information about an existing NetApp Snapshot.
 *
 * ## NetApp Snapshot Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const test = azure.netapp.getSnapshot({
 *     resourceGroupName: "acctestRG",
 *     name: "acctestnetappsnapshot",
 *     accountName: "acctestnetappaccount",
 *     poolName: "acctestnetapppool",
 *     volumeName: "acctestnetappvolume",
 * });
 * export const netappSnapshotId = data.azurerm_netapp_snapshot.example.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/netapp_snapshot.html.markdown.
 */
export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:netapp/getSnapshot:getSnapshot", {
        "accountName": args.accountName,
        "name": args.name,
        "poolName": args.poolName,
        "resourceGroupName": args.resourceGroupName,
        "volumeName": args.volumeName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshot.
 */
export interface GetSnapshotArgs {
    /**
     * The name of the NetApp Account where the NetApp Pool exists.
     */
    readonly accountName: string;
    /**
     * The name of the NetApp Snapshot.
     */
    readonly name: string;
    /**
     * The name of the NetApp Pool where the NetApp Volume exists.
     */
    readonly poolName: string;
    /**
     * The Name of the Resource Group where the NetApp Snapshot exists.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the NetApp Volume where the NetApp Snapshot exists.
     */
    readonly volumeName: string;
}

/**
 * A collection of values returned by getSnapshot.
 */
export interface GetSnapshotResult {
    readonly accountName: string;
    /**
     * The Azure Region where the NetApp Snapshot exists.
     */
    readonly location: string;
    readonly name: string;
    readonly poolName: string;
    readonly resourceGroupName: string;
    readonly volumeName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
