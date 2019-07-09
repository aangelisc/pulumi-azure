// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Network Profile.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 *     name: "examplegroup",
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("example", {
 *     addressSpaces: ["10.1.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     name: "examplevnet",
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("example", {
 *     addressPrefix: "10.1.0.0/24",
 *     delegations: [{
 *         name: "delegation",
 *         serviceDelegation: {
 *             actions: ["Microsoft.Network/virtualNetworks/subnets/action"],
 *             name: "Microsoft.ContainerInstance/containerGroups",
 *         },
 *     }],
 *     name: "examplesubnet",
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 * });
 * const exampleProfile = new azure.network.Profile("example", {
 *     containerNetworkInterfaceConfiguration: [{
 *         ipConfiguration: [{
 *             name: "exampleipconfig",
 *             subnetId: exampleSubnet.id,
 *         }],
 *         name: "examplecnic",
 *     }],
 *     location: exampleResourceGroup.location,
 *     name: "examplenetprofile",
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_profile.html.markdown.
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileState, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/profile:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    public readonly containerNetworkInterface!: pulumi.Output<{ ipConfigurations: { name: string, subnetId: string }[], name: string }>;
    /**
     * One or more Resource IDs of Azure Container Network Interfaces.
     */
    public /*out*/ readonly containerNetworkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Network Profile. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileArgs | ProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileState | undefined;
            inputs["containerNetworkInterface"] = state ? state.containerNetworkInterface : undefined;
            inputs["containerNetworkInterfaceIds"] = state ? state.containerNetworkInterfaceIds : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ProfileArgs | undefined;
            if (!args || args.containerNetworkInterface === undefined) {
                throw new Error("Missing required property 'containerNetworkInterface'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["containerNetworkInterface"] = args ? args.containerNetworkInterface : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["containerNetworkInterfaceIds"] = undefined /*out*/;
        }
        super(Profile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Profile resources.
 */
export interface ProfileState {
    readonly containerNetworkInterface?: pulumi.Input<{ ipConfigurations: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, subnetId: pulumi.Input<string> }>[]>, name: pulumi.Input<string> }>;
    /**
     * One or more Resource IDs of Azure Container Network Interfaces.
     */
    readonly containerNetworkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Network Profile. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    readonly containerNetworkInterface: pulumi.Input<{ ipConfigurations: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, subnetId: pulumi.Input<string> }>[]>, name: pulumi.Input<string> }>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Network Profile. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
