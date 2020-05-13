// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Authorization Rule for a ServiceBus Queue.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     tags: {
 *         source: "example",
 *     },
 * });
 * const exampleQueue = new azure.servicebus.Queue("exampleQueue", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     enablePartitioning: true,
 * });
 * const exampleQueueAuthorizationRule = new azure.servicebus.QueueAuthorizationRule("exampleQueueAuthorizationRule", {
 *     namespaceName: exampleNamespace.name,
 *     queueName: exampleQueue.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     listen: true,
 *     send: true,
 *     manage: false,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_queue_authorization_rule.html.markdown.
 */
export class QueueAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing QueueAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueAuthorizationRuleState, opts?: pulumi.CustomResourceOptions): QueueAuthorizationRule {
        return new QueueAuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule';

    /**
     * Returns true if the given object is an instance of QueueAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueueAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueueAuthorizationRule.__pulumiType;
    }

    /**
     * Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
     */
    public readonly listen!: pulumi.Output<boolean | undefined>;
    /**
     * Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    public readonly manage!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The Primary Connection String for the Authorization Rule.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The Primary Key for the Authorization Rule.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
     */
    public readonly queueName!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Connection String for the Authorization Rule.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The Secondary Key for the Authorization Rule.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
     */
    public readonly send!: pulumi.Output<boolean | undefined>;

    /**
     * Create a QueueAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueAuthorizationRuleArgs | QueueAuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QueueAuthorizationRuleState | undefined;
            inputs["listen"] = state ? state.listen : undefined;
            inputs["manage"] = state ? state.manage : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            inputs["primaryKey"] = state ? state.primaryKey : undefined;
            inputs["queueName"] = state ? state.queueName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            inputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            inputs["send"] = state ? state.send : undefined;
        } else {
            const args = argsOrState as QueueAuthorizationRuleArgs | undefined;
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.queueName === undefined) {
                throw new Error("Missing required property 'queueName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["listen"] = args ? args.listen : undefined;
            inputs["manage"] = args ? args.manage : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["queueName"] = args ? args.queueName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["send"] = args ? args.send : undefined;
            inputs["primaryConnectionString"] = undefined /*out*/;
            inputs["primaryKey"] = undefined /*out*/;
            inputs["secondaryConnectionString"] = undefined /*out*/;
            inputs["secondaryKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure:eventhub/queueAuthorizationRule:QueueAuthorizationRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(QueueAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QueueAuthorizationRule resources.
 */
export interface QueueAuthorizationRuleState {
    /**
     * Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
     */
    readonly listen?: pulumi.Input<boolean>;
    /**
     * Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    readonly manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * The Primary Connection String for the Authorization Rule.
     */
    readonly primaryConnectionString?: pulumi.Input<string>;
    /**
     * The Primary Key for the Authorization Rule.
     */
    readonly primaryKey?: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
     */
    readonly queueName?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Connection String for the Authorization Rule.
     */
    readonly secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The Secondary Key for the Authorization Rule.
     */
    readonly secondaryKey?: pulumi.Input<string>;
    /**
     * Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
     */
    readonly send?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a QueueAuthorizationRule resource.
 */
export interface QueueAuthorizationRuleArgs {
    /**
     * Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
     */
    readonly listen?: pulumi.Input<boolean>;
    /**
     * Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    readonly manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
     */
    readonly queueName: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
     */
    readonly send?: pulumi.Input<boolean>;
}
