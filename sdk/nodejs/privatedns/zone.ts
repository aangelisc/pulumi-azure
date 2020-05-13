// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Enables you to manage Private DNS zones within Azure DNS. These zones are hosted on Azure's name servers.
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
 * const exampleZone = new azure.privatedns.Zone("exampleZone", {resourceGroupName: exampleResourceGroup.name});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/private_dns_zone.html.markdown.
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatedns/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * The maximum number of record sets that can be created in this Private DNS zone.
     */
    public /*out*/ readonly maxNumberOfRecordSets!: pulumi.Output<number>;
    /**
     * The maximum number of virtual networks that can be linked to this Private DNS zone.
     */
    public /*out*/ readonly maxNumberOfVirtualNetworkLinks!: pulumi.Output<number>;
    /**
     * The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
     */
    public /*out*/ readonly maxNumberOfVirtualNetworkLinksWithRegistration!: pulumi.Output<number>;
    /**
     * The name of the Private DNS Zone. Must be a valid domain name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current number of record sets in this Private DNS zone.
     */
    public /*out*/ readonly numberOfRecordSets!: pulumi.Output<number>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ZoneState | undefined;
            inputs["maxNumberOfRecordSets"] = state ? state.maxNumberOfRecordSets : undefined;
            inputs["maxNumberOfVirtualNetworkLinks"] = state ? state.maxNumberOfVirtualNetworkLinks : undefined;
            inputs["maxNumberOfVirtualNetworkLinksWithRegistration"] = state ? state.maxNumberOfVirtualNetworkLinksWithRegistration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["numberOfRecordSets"] = state ? state.numberOfRecordSets : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["maxNumberOfRecordSets"] = undefined /*out*/;
            inputs["maxNumberOfVirtualNetworkLinks"] = undefined /*out*/;
            inputs["maxNumberOfVirtualNetworkLinksWithRegistration"] = undefined /*out*/;
            inputs["numberOfRecordSets"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Zone.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * The maximum number of record sets that can be created in this Private DNS zone.
     */
    readonly maxNumberOfRecordSets?: pulumi.Input<number>;
    /**
     * The maximum number of virtual networks that can be linked to this Private DNS zone.
     */
    readonly maxNumberOfVirtualNetworkLinks?: pulumi.Input<number>;
    /**
     * The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
     */
    readonly maxNumberOfVirtualNetworkLinksWithRegistration?: pulumi.Input<number>;
    /**
     * The name of the Private DNS Zone. Must be a valid domain name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The current number of record sets in this Private DNS zone.
     */
    readonly numberOfRecordSets?: pulumi.Input<number>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * The name of the Private DNS Zone. Must be a valid domain name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
