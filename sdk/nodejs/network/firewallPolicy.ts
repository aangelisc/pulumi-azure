// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Firewall Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.network.FirewallPolicy("example", {
 *     location: "West Europe",
 *     resourceGroupName: "example",
 * });
 * ```
 */
export class FirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallPolicyState, opts?: pulumi.CustomResourceOptions): FirewallPolicy {
        return new FirewallPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/firewallPolicy:FirewallPolicy';

    /**
     * Returns true if the given object is an instance of FirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicy.__pulumiType;
    }

    /**
     * The ID of the base Firewall Policy.
     */
    public readonly basePolicyId!: pulumi.Output<string | undefined>;
    /**
     * A list of reference to child Firewall Policies of this Firewall Policy.
     */
    public /*out*/ readonly childPolicies!: pulumi.Output<string[]>;
    /**
     * A `dns` block as defined below.
     */
    public readonly dns!: pulumi.Output<outputs.network.FirewallPolicyDns | undefined>;
    /**
     * A list of references to Azure Firewalls that this Firewall Policy is associated with.
     */
    public /*out*/ readonly firewalls!: pulumi.Output<string[]>;
    /**
     * The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
     */
    public /*out*/ readonly ruleCollectionGroups!: pulumi.Output<string[]>;
    /**
     * A mapping of tags which should be assigned to the Firewall Policy.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `threatIntelligenceAllowlist` block as defined below.
     */
    public readonly threatIntelligenceAllowlist!: pulumi.Output<outputs.network.FirewallPolicyThreatIntelligenceAllowlist | undefined>;
    /**
     * The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
     */
    public readonly threatIntelligenceMode!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicyArgs | FirewallPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirewallPolicyState | undefined;
            inputs["basePolicyId"] = state ? state.basePolicyId : undefined;
            inputs["childPolicies"] = state ? state.childPolicies : undefined;
            inputs["dns"] = state ? state.dns : undefined;
            inputs["firewalls"] = state ? state.firewalls : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["ruleCollectionGroups"] = state ? state.ruleCollectionGroups : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["threatIntelligenceAllowlist"] = state ? state.threatIntelligenceAllowlist : undefined;
            inputs["threatIntelligenceMode"] = state ? state.threatIntelligenceMode : undefined;
        } else {
            const args = argsOrState as FirewallPolicyArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["basePolicyId"] = args ? args.basePolicyId : undefined;
            inputs["dns"] = args ? args.dns : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["threatIntelligenceAllowlist"] = args ? args.threatIntelligenceAllowlist : undefined;
            inputs["threatIntelligenceMode"] = args ? args.threatIntelligenceMode : undefined;
            inputs["childPolicies"] = undefined /*out*/;
            inputs["firewalls"] = undefined /*out*/;
            inputs["ruleCollectionGroups"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FirewallPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallPolicy resources.
 */
export interface FirewallPolicyState {
    /**
     * The ID of the base Firewall Policy.
     */
    readonly basePolicyId?: pulumi.Input<string>;
    /**
     * A list of reference to child Firewall Policies of this Firewall Policy.
     */
    readonly childPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `dns` block as defined below.
     */
    readonly dns?: pulumi.Input<inputs.network.FirewallPolicyDns>;
    /**
     * A list of references to Azure Firewalls that this Firewall Policy is associated with.
     */
    readonly firewalls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
     */
    readonly ruleCollectionGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags which should be assigned to the Firewall Policy.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `threatIntelligenceAllowlist` block as defined below.
     */
    readonly threatIntelligenceAllowlist?: pulumi.Input<inputs.network.FirewallPolicyThreatIntelligenceAllowlist>;
    /**
     * The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
     */
    readonly threatIntelligenceMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallPolicy resource.
 */
export interface FirewallPolicyArgs {
    /**
     * The ID of the base Firewall Policy.
     */
    readonly basePolicyId?: pulumi.Input<string>;
    /**
     * A `dns` block as defined below.
     */
    readonly dns?: pulumi.Input<inputs.network.FirewallPolicyDns>;
    /**
     * The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Firewall Policy.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `threatIntelligenceAllowlist` block as defined below.
     */
    readonly threatIntelligenceAllowlist?: pulumi.Input<inputs.network.FirewallPolicyThreatIntelligenceAllowlist>;
    /**
     * The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
     */
    readonly threatIntelligenceMode?: pulumi.Input<string>;
}