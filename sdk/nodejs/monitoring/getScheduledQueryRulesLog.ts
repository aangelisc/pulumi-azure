// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access the properties of a LogToMetricAction scheduled query rule.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.monitoring.getScheduledQueryRulesLog({
 *     name: "tfex-queryrule",
 *     resourceGroupName: "example-rg",
 * }, { async: true }));
 *
 * export const queryRuleId = example.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/monitor_scheduled_query_rules_log.html.markdown.
 */
export function getScheduledQueryRulesLog(args: GetScheduledQueryRulesLogArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduledQueryRulesLogResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:monitoring/getScheduledQueryRulesLog:getScheduledQueryRulesLog", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getScheduledQueryRulesLog.
 */
export interface GetScheduledQueryRulesLogArgs {
    /**
     * Specifies the name of the scheduled query rule.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group where the scheduled query rule is located.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getScheduledQueryRulesLog.
 */
export interface GetScheduledQueryRulesLogResult {
    readonly authorizedResourceIds: string[];
    /**
     * A `criteria` block as defined below.
     */
    readonly criterias: outputs.monitoring.GetScheduledQueryRulesLogCriteria[];
    /**
     * The resource URI over which log search query is to be run.
     */
    readonly dataSourceId: string;
    /**
     * The description of the scheduled query rule.
     */
    readonly description: string;
    /**
     * Whether this scheduled query rule is enabled.
     */
    readonly enabled: boolean;
    readonly location: string;
    /**
     * Name of the dimension.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    readonly tags: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
