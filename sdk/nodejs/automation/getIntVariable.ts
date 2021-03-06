// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Automation Int Variable.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.automation.getIntVariable({
 *     name: "tfex-example-var",
 *     resourceGroupName: "tfex-example-rg",
 *     automationAccountName: "tfex-example-account",
 * });
 * export const variableId = example.then(example => example.id);
 * ```
 */
export function getIntVariable(args: GetIntVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetIntVariableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:automation/getIntVariable:getIntVariable", {
        "automationAccountName": args.automationAccountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getIntVariable.
 */
export interface GetIntVariableArgs {
    /**
     * The name of the automation account in which the Automation Variable exists.
     */
    automationAccountName: string;
    /**
     * The name of the Automation Variable.
     */
    name: string;
    /**
     * The Name of the Resource Group where the automation account exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getIntVariable.
 */
export interface GetIntVariableResult {
    readonly automationAccountName: string;
    /**
     * The description of the Automation Variable.
     */
    readonly description: string;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The value of the Automation Variable as a `integer`.
     */
    readonly value: number;
}
