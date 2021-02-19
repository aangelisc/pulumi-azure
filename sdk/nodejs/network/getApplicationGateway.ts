// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Application Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getApplicationGateway({
 *     name: "existing-app-gateway",
 *     resourceGroupName: "existing-resources",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getApplicationGateway(args: GetApplicationGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getApplicationGateway:getApplicationGateway", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationGateway.
 */
export interface GetApplicationGatewayArgs {
    /**
     * The name of this Application Gateway.
     */
    readonly name: string;
    /**
     * The name of the Resource Group where the Application Gateway exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getApplicationGateway.
 */
export interface GetApplicationGatewayResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A `identity` block as defined below.
     */
    readonly identities: outputs.network.GetApplicationGatewayIdentity[];
    /**
     * The Azure Region where the Application Gateway exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the Application Gateway.
     */
    readonly tags: {[key: string]: string};
}
