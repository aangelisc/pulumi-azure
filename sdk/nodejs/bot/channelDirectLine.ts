// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Directline integration for a Bot Channel
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const current = azure.core.getClientConfig();
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "northeurope",
 * });
 * const exampleChannelsRegistration = new azure.bot.ChannelsRegistration("example", {
 *     location: "global",
 *     microsoftAppId: current.servicePrincipalApplicationId,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "F0",
 * });
 * const exampleChannelDirectLine = new azure.bot.ChannelDirectLine("example", {
 *     botName: exampleChannelsRegistration.name,
 *     location: exampleChannelsRegistration.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sites: [{
 *         enabled: true,
 *         name: "default",
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channel_directline.markdown.
 */
export class ChannelDirectLine extends pulumi.CustomResource {
    /**
     * Get an existing ChannelDirectLine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChannelDirectLineState, opts?: pulumi.CustomResourceOptions): ChannelDirectLine {
        return new ChannelDirectLine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:bot/channelDirectLine:ChannelDirectLine';

    /**
     * Returns true if the given object is an instance of ChannelDirectLine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChannelDirectLine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChannelDirectLine.__pulumiType;
    }

    public readonly botName!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    public readonly resourceGroupName!: pulumi.Output<string>;
    public readonly sites!: pulumi.Output<outputs.bot.ChannelDirectLineSite[]>;

    /**
     * Create a ChannelDirectLine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelDirectLineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChannelDirectLineArgs | ChannelDirectLineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ChannelDirectLineState | undefined;
            inputs["botName"] = state ? state.botName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sites"] = state ? state.sites : undefined;
        } else {
            const args = argsOrState as ChannelDirectLineArgs | undefined;
            if (!args || args.botName === undefined) {
                throw new Error("Missing required property 'botName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sites === undefined) {
                throw new Error("Missing required property 'sites'");
            }
            inputs["botName"] = args ? args.botName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sites"] = args ? args.sites : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ChannelDirectLine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChannelDirectLine resources.
 */
export interface ChannelDirectLineState {
    readonly botName?: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    readonly resourceGroupName?: pulumi.Input<string>;
    readonly sites?: pulumi.Input<pulumi.Input<inputs.bot.ChannelDirectLineSite>[]>;
}

/**
 * The set of arguments for constructing a ChannelDirectLine resource.
 */
export interface ChannelDirectLineArgs {
    readonly botName: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    readonly resourceGroupName: pulumi.Input<string>;
    readonly sites: pulumi.Input<pulumi.Input<inputs.bot.ChannelDirectLineSite>[]>;
}
