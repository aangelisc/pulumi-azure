// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Vmware Cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const examplePrivateCloud = new azure.avs.PrivateCloud("examplePrivateCloud", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     skuName: "av36",
 *     managementCluster: {
 *         size: 3,
 *     },
 *     networkSubnetCidr: "192.168.48.0/22",
 *     internetConnectionEnabled: false,
 *     nsxtPassword: `QazWsx13$Edc`,
 *     vcenterPassword: `WsxEdc23$Rfv`,
 * });
 * const exampleCluster = new azure.avs.Cluster("exampleCluster", {
 *     vmwareCloudId: examplePrivateCloud.id,
 *     clusterNodeCount: 3,
 *     skuName: "av36",
 * });
 * ```
 *
 * ## Import
 *
 * Vmware Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:avs/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/cluster1
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:avs/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The count of the Vmware Cluster nodes.
     */
    public readonly clusterNodeCount!: pulumi.Output<number>;
    /**
     * A number that identifies this Vmware Cluster in its Vmware Private Cloud.
     */
    public /*out*/ readonly clusterNumber!: pulumi.Output<number>;
    /**
     * A list of host of the Vmware Cluster.
     */
    public /*out*/ readonly hosts!: pulumi.Output<string[]>;
    /**
     * The name which should be used for this Vmware Cluster. Changing this forces a new Vmware Cluster to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The cluster sku to use. Possible values are `av20`, `av36`, and `av36t`. Changing this forces a new Vmware Cluster to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * The ID of the Vmware Private Cloud in which to create this Vmware Cluster. Changing this forces a new Vmware Cluster to be created.
     */
    public readonly vmwareCloudId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["clusterNodeCount"] = state ? state.clusterNodeCount : undefined;
            inputs["clusterNumber"] = state ? state.clusterNumber : undefined;
            inputs["hosts"] = state ? state.hosts : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["vmwareCloudId"] = state ? state.vmwareCloudId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.clusterNodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterNodeCount'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            if ((!args || args.vmwareCloudId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmwareCloudId'");
            }
            inputs["clusterNodeCount"] = args ? args.clusterNodeCount : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["vmwareCloudId"] = args ? args.vmwareCloudId : undefined;
            inputs["clusterNumber"] = undefined /*out*/;
            inputs["hosts"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The count of the Vmware Cluster nodes.
     */
    readonly clusterNodeCount?: pulumi.Input<number>;
    /**
     * A number that identifies this Vmware Cluster in its Vmware Private Cloud.
     */
    readonly clusterNumber?: pulumi.Input<number>;
    /**
     * A list of host of the Vmware Cluster.
     */
    readonly hosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name which should be used for this Vmware Cluster. Changing this forces a new Vmware Cluster to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The cluster sku to use. Possible values are `av20`, `av36`, and `av36t`. Changing this forces a new Vmware Cluster to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * The ID of the Vmware Private Cloud in which to create this Vmware Cluster. Changing this forces a new Vmware Cluster to be created.
     */
    readonly vmwareCloudId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The count of the Vmware Cluster nodes.
     */
    readonly clusterNodeCount: pulumi.Input<number>;
    /**
     * The name which should be used for this Vmware Cluster. Changing this forces a new Vmware Cluster to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The cluster sku to use. Possible values are `av20`, `av36`, and `av36t`. Changing this forces a new Vmware Cluster to be created.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * The ID of the Vmware Private Cloud in which to create this Vmware Cluster. Changing this forces a new Vmware Cluster to be created.
     */
    readonly vmwareCloudId: pulumi.Input<string>;
}