// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Resource Group Consumption Budget.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getSubscription({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "eastus"});
 * const test = new azure.monitoring.ActionGroup("test", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     shortName: "example",
 * });
 * const exampleBudgetResourceGroup = new azure.consumption.BudgetResourceGroup("exampleBudgetResourceGroup", {
 *     subscriptionId: current.then(current => current.subscriptionId),
 *     resourceGroupId: exampleResourceGroup.id,
 *     amount: 1000,
 *     timeGrain: "Monthly",
 *     timePeriod: {
 *         startDate: "2020-11-01T00:00:00Z",
 *         endDate: "2020-12-01T00:00:00Z",
 *     },
 *     filter: {
 *         dimensions: [{
 *             name: "ResourceId",
 *             values: [azurerm_monitor_action_group.example.id],
 *         }],
 *         tags: [{
 *             name: "foo",
 *             values: [
 *                 "bar",
 *                 "baz",
 *             ],
 *         }],
 *     },
 *     notifications: [
 *         {
 *             enabled: true,
 *             threshold: 90,
 *             operator: "EqualTo",
 *             contactEmails: [
 *                 "foo@example.com",
 *                 "bar@example.com",
 *             ],
 *             contactGroups: [azurerm_monitor_action_group.example.id],
 *             contactRoles: ["Owner"],
 *         },
 *         {
 *             enabled: false,
 *             threshold: 100,
 *             operator: "GreaterThan",
 *             contactEmails: [
 *                 "foo@example.com",
 *                 "bar@example.com",
 *             ],
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Resource Group Consumption Budgets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:consumption/budgetResourceGroup:BudgetResourceGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Consumption/budgets/resourceGroup1
 * ```
 */
export class BudgetResourceGroup extends pulumi.CustomResource {
    /**
     * Get an existing BudgetResourceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BudgetResourceGroupState, opts?: pulumi.CustomResourceOptions): BudgetResourceGroup {
        return new BudgetResourceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:consumption/budgetResourceGroup:BudgetResourceGroup';

    /**
     * Returns true if the given object is an instance of BudgetResourceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BudgetResourceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BudgetResourceGroup.__pulumiType;
    }

    /**
     * The total amount of cost to track with the budget.
     */
    public readonly amount!: pulumi.Output<number>;
    /**
     * A `filter` block as defined below.
     */
    public readonly filter!: pulumi.Output<outputs.consumption.BudgetResourceGroupFilter | undefined>;
    /**
     * The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more `notification` blocks as defined below.
     */
    public readonly notifications!: pulumi.Output<outputs.consumption.BudgetResourceGroupNotification[]>;
    /**
     * The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
     */
    public readonly timeGrain!: pulumi.Output<string | undefined>;
    /**
     * A `timePeriod` block as defined below.
     */
    public readonly timePeriod!: pulumi.Output<outputs.consumption.BudgetResourceGroupTimePeriod>;

    /**
     * Create a BudgetResourceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BudgetResourceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BudgetResourceGroupArgs | BudgetResourceGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BudgetResourceGroupState | undefined;
            inputs["amount"] = state ? state.amount : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifications"] = state ? state.notifications : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["timeGrain"] = state ? state.timeGrain : undefined;
            inputs["timePeriod"] = state ? state.timePeriod : undefined;
        } else {
            const args = argsOrState as BudgetResourceGroupArgs | undefined;
            if ((!args || args.amount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'amount'");
            }
            if ((!args || args.notifications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notifications'");
            }
            if ((!args || args.resourceGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupId'");
            }
            if ((!args || args.timePeriod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timePeriod'");
            }
            inputs["amount"] = args ? args.amount : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["timeGrain"] = args ? args.timeGrain : undefined;
            inputs["timePeriod"] = args ? args.timePeriod : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BudgetResourceGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BudgetResourceGroup resources.
 */
export interface BudgetResourceGroupState {
    /**
     * The total amount of cost to track with the budget.
     */
    readonly amount?: pulumi.Input<number>;
    /**
     * A `filter` block as defined below.
     */
    readonly filter?: pulumi.Input<inputs.consumption.BudgetResourceGroupFilter>;
    /**
     * The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `notification` blocks as defined below.
     */
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.consumption.BudgetResourceGroupNotification>[]>;
    /**
     * The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
     */
    readonly timeGrain?: pulumi.Input<string>;
    /**
     * A `timePeriod` block as defined below.
     */
    readonly timePeriod?: pulumi.Input<inputs.consumption.BudgetResourceGroupTimePeriod>;
}

/**
 * The set of arguments for constructing a BudgetResourceGroup resource.
 */
export interface BudgetResourceGroupArgs {
    /**
     * The total amount of cost to track with the budget.
     */
    readonly amount: pulumi.Input<number>;
    /**
     * A `filter` block as defined below.
     */
    readonly filter?: pulumi.Input<inputs.consumption.BudgetResourceGroupFilter>;
    /**
     * The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `notification` blocks as defined below.
     */
    readonly notifications: pulumi.Input<pulumi.Input<inputs.consumption.BudgetResourceGroupNotification>[]>;
    /**
     * The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
     */
    readonly resourceGroupId: pulumi.Input<string>;
    /**
     * The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
     */
    readonly timeGrain?: pulumi.Input<string>;
    /**
     * A `timePeriod` block as defined below.
     */
    readonly timePeriod: pulumi.Input<inputs.consumption.BudgetResourceGroupTimePeriod>;
}
