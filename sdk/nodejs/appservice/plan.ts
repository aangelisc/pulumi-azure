// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Create an App Service Plan component.
 */
export class Plan extends pulumi.CustomResource {
    /**
     * Get an existing Plan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlanState): Plan {
        return new Plan(name, <any>state, { id });
    }

    /**
     * The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux` and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
     */
    public readonly kind: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Maximum number of instances that can be assigned to this App Service plan.
     */
    public /*out*/ readonly maximumNumberOfWorkers: pulumi.Output<number>;
    /**
     * Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A `properties` block as documented below.
     */
    public readonly properties: pulumi.Output<{ appServiceEnvironmentId?: string, perSiteScaling?: boolean, reserved?: boolean }>;
    /**
     * The name of the resource group in which to create the App Service Plan component.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * A `sku` block as documented below.
     */
    public readonly sku: pulumi.Output<{ capacity: number, size: string, tier: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Plan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlanArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: PlanArgs | PlanState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PlanState = argsOrState as PlanState | undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maximumNumberOfWorkers"] = state ? state.maximumNumberOfWorkers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as PlanArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["maximumNumberOfWorkers"] = undefined /*out*/;
        }
        super("azure:appservice/plan:Plan", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Plan resources.
 */
export interface PlanState {
    /**
     * The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux` and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Maximum number of instances that can be assigned to this App Service plan.
     */
    readonly maximumNumberOfWorkers?: pulumi.Input<number>;
    /**
     * Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `properties` block as documented below.
     */
    readonly properties?: pulumi.Input<{ appServiceEnvironmentId?: pulumi.Input<string>, perSiteScaling?: pulumi.Input<boolean>, reserved?: pulumi.Input<boolean> }>;
    /**
     * The name of the resource group in which to create the App Service Plan component.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    readonly sku?: pulumi.Input<{ capacity?: pulumi.Input<number>, size: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Plan resource.
 */
export interface PlanArgs {
    /**
     * The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux` and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `properties` block as documented below.
     */
    readonly properties?: pulumi.Input<{ appServiceEnvironmentId?: pulumi.Input<string>, perSiteScaling?: pulumi.Input<boolean>, reserved?: pulumi.Input<boolean> }>;
    /**
     * The name of the resource group in which to create the App Service Plan component.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    readonly sku: pulumi.Input<{ capacity?: pulumi.Input<number>, size: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
