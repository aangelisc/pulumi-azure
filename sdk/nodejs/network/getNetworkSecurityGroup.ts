// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Network Security Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getNetworkSecurityGroup({
 *     name: "example",
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * export const location = example.then(example => example.location);
 * ```
 */
export function getNetworkSecurityGroup(args: GetNetworkSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkSecurityGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getNetworkSecurityGroup:getNetworkSecurityGroup", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkSecurityGroup.
 */
export interface GetNetworkSecurityGroupArgs {
    /**
     * Specifies the Name of the Network Security Group.
     */
    name: string;
    /**
     * Specifies the Name of the Resource Group within which the Network Security Group exists
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getNetworkSecurityGroup.
 */
export interface GetNetworkSecurityGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The supported Azure location where the resource exists.
     */
    readonly location: string;
    /**
     * The name of the security rule.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * One or more `securityRule` blocks as defined below.
     */
    readonly securityRules: outputs.network.GetNetworkSecurityGroupSecurityRule[];
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}
