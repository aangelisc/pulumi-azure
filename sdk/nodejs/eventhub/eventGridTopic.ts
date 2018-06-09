// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages an EventGrid Topic
 * 
 * ~> **Note:** at this time EventGrid Topic's are only available in a limited number of regions.
 */
export class EventGridTopic extends pulumi.CustomResource {
    /**
     * Get an existing EventGridTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventGridTopicState): EventGridTopic {
        return new EventGridTopic(name, <any>state, { id });
    }

    /**
     * The Endpoint associated with the EventGrid Topic.
     */
    public /*out*/ readonly endpoint: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The Primary Shared Access Key associated with the EventGrid Topic.
     */
    public /*out*/ readonly primaryAccessKey: pulumi.Output<string>;
    /**
     * The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The Secondary Shared Access Key associated with the EventGrid Topic.
     */
    public /*out*/ readonly secondaryAccessKey: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a EventGridTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventGridTopicArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: EventGridTopicArgs | EventGridTopicState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: EventGridTopicState = argsOrState as EventGridTopicState | undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventGridTopicArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["endpoint"] = undefined /*out*/;
            inputs["primaryAccessKey"] = undefined /*out*/;
            inputs["secondaryAccessKey"] = undefined /*out*/;
        }
        super("azure:eventhub/eventGridTopic:EventGridTopic", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventGridTopic resources.
 */
export interface EventGridTopicState {
    /**
     * The Endpoint associated with the EventGrid Topic.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Primary Shared Access Key associated with the EventGrid Topic.
     */
    readonly primaryAccessKey?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Shared Access Key associated with the EventGrid Topic.
     */
    readonly secondaryAccessKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EventGridTopic resource.
 */
export interface EventGridTopicArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
