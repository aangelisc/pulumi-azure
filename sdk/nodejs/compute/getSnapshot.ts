// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Snapshot.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.compute.getSnapshot({
 *     name: "my-snapshot",
 *     resourceGroupName: "my-resource-group",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/snapshot.html.markdown.
 */
export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:compute/getSnapshot:getSnapshot", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshot.
 */
export interface GetSnapshotArgs {
    /**
     * Specifies the name of the Snapshot.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Snapshot is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getSnapshot.
 */
export interface GetSnapshotResult {
    readonly creationOption: string;
    /**
     * The size of the Snapshotted Disk in GB.
     */
    readonly diskSizeGb: number;
    readonly encryptionSettings: outputs.compute.GetSnapshotEncryptionSetting[];
    readonly name: string;
    readonly osType: string;
    readonly resourceGroupName: string;
    /**
     * The reference to an existing snapshot.
     */
    readonly sourceResourceId: string;
    /**
     * The URI to a Managed or Unmanaged Disk.
     */
    readonly sourceUri: string;
    /**
     * The ID of an storage account.
     */
    readonly storageAccountId: string;
    readonly timeCreated: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
