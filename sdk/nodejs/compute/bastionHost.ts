// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Bastion Host.
 * 
 * > **Note:** Bastion Hosts are a preview feature in Azure, and therefore are only supported in a select number of regions. [Read more](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq).
 * 
 * ## Example Usage
 * 
 * This example deploys an Azure Bastion Host Instance to a target virtual network.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("example", {
 *     addressSpaces: ["192.168.1.0/24"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("example", {
 *     addressPrefix: "192.168.1.224/27",
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 * });
 * const examplePublicIp = new azure.network.PublicIp("example", {
 *     allocationMethod: "Static",
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 * });
 * const exampleBastionHost = new azure.compute.BastionHost("example", {
 *     ipConfiguration: {
 *         name: "configuration",
 *         publicIpAddressId: examplePublicIp.id,
 *         subnetId: exampleSubnet.id,
 *     },
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bastion_host.html.markdown.
 */
export class BastionHost extends pulumi.CustomResource {
    /**
     * Get an existing BastionHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BastionHostState, opts?: pulumi.CustomResourceOptions): BastionHost {
        return new BastionHost(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/bastionHost:BastionHost';

    /**
     * Returns true if the given object is an instance of BastionHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BastionHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BastionHost.__pulumiType;
    }

    /**
     * The FQDN for the Bastion Host.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * A `ipConfiguration` block as defined below.
     */
    public readonly ipConfiguration!: pulumi.Output<outputs.compute.BastionHostIpConfiguration | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Bastion Host.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a BastionHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BastionHostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BastionHostArgs | BastionHostState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BastionHostState | undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["ipConfiguration"] = state ? state.ipConfiguration : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as BastionHostArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["ipConfiguration"] = args ? args.ipConfiguration : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["dnsName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BastionHost.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BastionHost resources.
 */
export interface BastionHostState {
    /**
     * The FQDN for the Bastion Host.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * A `ipConfiguration` block as defined below.
     */
    readonly ipConfiguration?: pulumi.Input<inputs.compute.BastionHostIpConfiguration>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bastion Host.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a BastionHost resource.
 */
export interface BastionHostArgs {
    /**
     * A `ipConfiguration` block as defined below.
     */
    readonly ipConfiguration?: pulumi.Input<inputs.compute.BastionHostIpConfiguration>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bastion Host.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
