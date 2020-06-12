// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Assignment extends pulumi.CustomResource {
    /**
     * Get an existing Assignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssignmentState, opts?: pulumi.CustomResourceOptions): Assignment {
        return new Assignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:blueprint/assignment:Assignment';

    /**
     * Returns true if the given object is an instance of Assignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Assignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Assignment.__pulumiType;
    }

    /**
     * The name of the blueprint assigned
     */
    public /*out*/ readonly blueprintName!: pulumi.Output<string>;
    /**
     * The Description on the Blueprint
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * The display name of the blueprint
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    public readonly identity!: pulumi.Output<outputs.blueprint.AssignmentIdentity | undefined>;
    /**
     * The Azure location of the Assignment. 
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
     */
    public readonly lockExcludePrincipals!: pulumi.Output<string[] | undefined>;
    /**
     * The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
     */
    public readonly lockMode!: pulumi.Output<string | undefined>;
    /**
     * The name of the Blueprint Assignment
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * a JSON string to supply Blueprint Assignment parameter values.
     */
    public readonly parameterValues!: pulumi.Output<string | undefined>;
    /**
     * a JSON string to supply the Blueprint Resource Group information. 
     */
    public readonly resourceGroups!: pulumi.Output<string | undefined>;
    /**
     * The Subscription ID the Blueprint Published Version is to be applied to.
     */
    public readonly targetSubscriptionId!: pulumi.Output<string>;
    /**
     * The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The ID of the Published Version of the blueprint to be assigned. 
     */
    public readonly versionId!: pulumi.Output<string>;

    /**
     * Create a Assignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssignmentArgs | AssignmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AssignmentState | undefined;
            inputs["blueprintName"] = state ? state.blueprintName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["lockExcludePrincipals"] = state ? state.lockExcludePrincipals : undefined;
            inputs["lockMode"] = state ? state.lockMode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameterValues"] = state ? state.parameterValues : undefined;
            inputs["resourceGroups"] = state ? state.resourceGroups : undefined;
            inputs["targetSubscriptionId"] = state ? state.targetSubscriptionId : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as AssignmentArgs | undefined;
            if (!args || args.targetSubscriptionId === undefined) {
                throw new Error("Missing required property 'targetSubscriptionId'");
            }
            if (!args || args.versionId === undefined) {
                throw new Error("Missing required property 'versionId'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["lockExcludePrincipals"] = args ? args.lockExcludePrincipals : undefined;
            inputs["lockMode"] = args ? args.lockMode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameterValues"] = args ? args.parameterValues : undefined;
            inputs["resourceGroups"] = args ? args.resourceGroups : undefined;
            inputs["targetSubscriptionId"] = args ? args.targetSubscriptionId : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
            inputs["blueprintName"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Assignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Assignment resources.
 */
export interface AssignmentState {
    /**
     * The name of the blueprint assigned
     */
    readonly blueprintName?: pulumi.Input<string>;
    /**
     * The Description on the Blueprint
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the blueprint
     */
    readonly displayName?: pulumi.Input<string>;
    readonly identity?: pulumi.Input<inputs.blueprint.AssignmentIdentity>;
    /**
     * The Azure location of the Assignment. 
     */
    readonly location?: pulumi.Input<string>;
    /**
     * a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
     */
    readonly lockExcludePrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
     */
    readonly lockMode?: pulumi.Input<string>;
    /**
     * The name of the Blueprint Assignment
     */
    readonly name?: pulumi.Input<string>;
    /**
     * a JSON string to supply Blueprint Assignment parameter values.
     */
    readonly parameterValues?: pulumi.Input<string>;
    /**
     * a JSON string to supply the Blueprint Resource Group information. 
     */
    readonly resourceGroups?: pulumi.Input<string>;
    /**
     * The Subscription ID the Blueprint Published Version is to be applied to.
     */
    readonly targetSubscriptionId?: pulumi.Input<string>;
    /**
     * The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The ID of the Published Version of the blueprint to be assigned. 
     */
    readonly versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Assignment resource.
 */
export interface AssignmentArgs {
    readonly identity?: pulumi.Input<inputs.blueprint.AssignmentIdentity>;
    /**
     * The Azure location of the Assignment. 
     */
    readonly location?: pulumi.Input<string>;
    /**
     * a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
     */
    readonly lockExcludePrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
     */
    readonly lockMode?: pulumi.Input<string>;
    /**
     * The name of the Blueprint Assignment
     */
    readonly name?: pulumi.Input<string>;
    /**
     * a JSON string to supply Blueprint Assignment parameter values.
     */
    readonly parameterValues?: pulumi.Input<string>;
    /**
     * a JSON string to supply the Blueprint Resource Group information. 
     */
    readonly resourceGroups?: pulumi.Input<string>;
    /**
     * The Subscription ID the Blueprint Published Version is to be applied to.
     */
    readonly targetSubscriptionId: pulumi.Input<string>;
    /**
     * The ID of the Published Version of the blueprint to be assigned. 
     */
    readonly versionId: pulumi.Input<string>;
}