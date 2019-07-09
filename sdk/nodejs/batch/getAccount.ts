// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Batch Account.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.output(azure.batch.getAccount({
 *     name: "testbatchaccount",
 *     resourceGroupName: "test",
 * }));
 * 
 * export const poolAllocationMode = test.poolAllocationMode;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/batch_account.html.markdown.
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    return pulumi.runtime.invoke("azure:batch/getAccount:getAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * The name of the Batch account.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group where this Batch account exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * The account endpoint used to interact with the Batch service.
     */
    readonly accountEndpoint: string;
    /**
     * The Azure Region in which this Batch account exists.
     */
    readonly location: string;
    /**
     * The Batch account name.
     */
    readonly name: string;
    /**
     * The pool allocation mode configured for this Batch account.
     */
    readonly poolAllocationMode: string;
    /**
     * The Batch account primary access key.
     */
    readonly primaryAccessKey: string;
    readonly resourceGroupName: string;
    /**
     * The Batch account secondary access key.
     */
    readonly secondaryAccessKey: string;
    /**
     * The ID of the Storage Account used for this Batch account.
     */
    readonly storageAccountId: string;
    /**
     * A map of tags assigned to the Batch account.
     */
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
