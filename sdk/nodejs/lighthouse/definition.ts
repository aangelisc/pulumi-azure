// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a [Lighthouse](https://docs.microsoft.com/en-us/azure/lighthouse) Definition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const contributor = azure.authorization.getRoleDefinition({
 *     roleDefinitionId: "b24988ac-6180-42a0-ab88-20f7382dd24c",
 * });
 * const example = new azure.lighthouse.Definition("example", {
 *     description: "This is a lighthouse definition created IaC",
 *     managingTenantId: "00000000-0000-0000-0000-000000000000",
 *     authorizations: [{
 *         principalId: "00000000-0000-0000-0000-000000000000",
 *         roleDefinitionId: contributor.then(contributor => contributor.roleDefinitionId),
 *         principalDisplayName: "Tier 1 Support",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Lighthouse Definitions can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:lighthouse/definition:Definition example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationDefinitions/00000000-0000-0000-0000-000000000000
 * ```
 */
export class Definition extends pulumi.CustomResource {
    /**
     * Get an existing Definition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefinitionState, opts?: pulumi.CustomResourceOptions): Definition {
        return new Definition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:lighthouse/definition:Definition';

    /**
     * Returns true if the given object is an instance of Definition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Definition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Definition.__pulumiType;
    }

    /**
     * An authorization block as defined below.
     */
    public readonly authorizations!: pulumi.Output<outputs.lighthouse.DefinitionAuthorization[]>;
    /**
     * A description of the Lighthouse Definition.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
     */
    public readonly lighthouseDefinitionId!: pulumi.Output<string>;
    /**
     * The ID of the managing tenant.
     */
    public readonly managingTenantId!: pulumi.Output<string>;
    /**
     * The name of the Lighthouse Definition.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly scope!: pulumi.Output<string>;

    /**
     * Create a Definition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefinitionArgs | DefinitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefinitionState | undefined;
            inputs["authorizations"] = state ? state.authorizations : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lighthouseDefinitionId"] = state ? state.lighthouseDefinitionId : undefined;
            inputs["managingTenantId"] = state ? state.managingTenantId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scope"] = state ? state.scope : undefined;
        } else {
            const args = argsOrState as DefinitionArgs | undefined;
            if ((!args || args.authorizations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizations'");
            }
            if ((!args || args.managingTenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managingTenantId'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["authorizations"] = args ? args.authorizations : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["lighthouseDefinitionId"] = args ? args.lighthouseDefinitionId : undefined;
            inputs["managingTenantId"] = args ? args.managingTenantId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scope"] = args ? args.scope : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Definition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Definition resources.
 */
export interface DefinitionState {
    /**
     * An authorization block as defined below.
     */
    readonly authorizations?: pulumi.Input<pulumi.Input<inputs.lighthouse.DefinitionAuthorization>[]>;
    /**
     * A description of the Lighthouse Definition.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
     */
    readonly lighthouseDefinitionId?: pulumi.Input<string>;
    /**
     * The ID of the managing tenant.
     */
    readonly managingTenantId?: pulumi.Input<string>;
    /**
     * The name of the Lighthouse Definition.
     */
    readonly name?: pulumi.Input<string>;
    readonly scope?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Definition resource.
 */
export interface DefinitionArgs {
    /**
     * An authorization block as defined below.
     */
    readonly authorizations: pulumi.Input<pulumi.Input<inputs.lighthouse.DefinitionAuthorization>[]>;
    /**
     * A description of the Lighthouse Definition.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique UUID/GUID which identifies this lighthouse definition - one will be generated if not specified. Changing this forces a new resource to be created.
     */
    readonly lighthouseDefinitionId?: pulumi.Input<string>;
    /**
     * The ID of the managing tenant.
     */
    readonly managingTenantId: pulumi.Input<string>;
    /**
     * The name of the Lighthouse Definition.
     */
    readonly name?: pulumi.Input<string>;
    readonly scope: pulumi.Input<string>;
}
