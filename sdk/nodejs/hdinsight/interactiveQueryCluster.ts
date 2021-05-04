// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a HDInsight Interactive Query Cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleContainer = new azure.storage.Container("exampleContainer", {
 *     storageAccountName: exampleAccount.name,
 *     containerAccessType: "private",
 * });
 * const exampleInteractiveQueryCluster = new azure.hdinsight.InteractiveQueryCluster("exampleInteractiveQueryCluster", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     clusterVersion: "3.6",
 *     tier: "Standard",
 *     componentVersion: {
 *         interactiveHive: "2.1",
 *     },
 *     gateway: {
 *         enabled: true,
 *         username: "acctestusrgw",
 *         password: "Password!",
 *     },
 *     storageAccounts: [{
 *         storageContainerId: exampleContainer.id,
 *         storageAccountKey: exampleAccount.primaryAccessKey,
 *         isDefault: true,
 *     }],
 *     roles: {
 *         headNode: {
 *             vmSize: "Standard_D13_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *         },
 *         workerNode: {
 *             vmSize: "Standard_D14_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *             targetInstanceCount: 3,
 *         },
 *         zookeeperNode: {
 *             vmSize: "Standard_A4_V2",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * HDInsight Interactive Query Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
 * ```
 */
export class InteractiveQueryCluster extends pulumi.CustomResource {
    /**
     * Get an existing InteractiveQueryCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InteractiveQueryClusterState, opts?: pulumi.CustomResourceOptions): InteractiveQueryCluster {
        return new InteractiveQueryCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster';

    /**
     * Returns true if the given object is an instance of InteractiveQueryCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InteractiveQueryCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InteractiveQueryCluster.__pulumiType;
    }

    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    public readonly clusterVersion!: pulumi.Output<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    public readonly componentVersion!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    public readonly gateway!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Interactive Query Cluster.
     */
    public /*out*/ readonly httpsEndpoint!: pulumi.Output<string>;
    /**
     * Specifies the Azure Region which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A `metastores` block as defined below.
     */
    public readonly metastores!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterMetastores | undefined>;
    /**
     * A `monitor` block as defined below.
     */
    public readonly monitor!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterMonitor | undefined>;
    /**
     * Specifies the name for this HDInsight Interactive Query Cluster. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `network` block as defined below.
     */
    public readonly network!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterNetwork | undefined>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `roles` block as defined below.
     */
    public readonly roles!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Interactive Query Cluster.
     */
    public /*out*/ readonly sshEndpoint!: pulumi.Output<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    public readonly storageAccountGen2!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterStorageAccountGen2 | undefined>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    public readonly storageAccounts!: pulumi.Output<outputs.hdinsight.InteractiveQueryClusterStorageAccount[] | undefined>;
    /**
     * A map of Tags which should be assigned to this HDInsight Interactive Query Cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the Tier which should be used for this HDInsight Interactive Query Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    public readonly tier!: pulumi.Output<string>;
    public readonly tlsMinVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a InteractiveQueryCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InteractiveQueryClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InteractiveQueryClusterArgs | InteractiveQueryClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InteractiveQueryClusterState | undefined;
            inputs["clusterVersion"] = state ? state.clusterVersion : undefined;
            inputs["componentVersion"] = state ? state.componentVersion : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["httpsEndpoint"] = state ? state.httpsEndpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["metastores"] = state ? state.metastores : undefined;
            inputs["monitor"] = state ? state.monitor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["sshEndpoint"] = state ? state.sshEndpoint : undefined;
            inputs["storageAccountGen2"] = state ? state.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = state ? state.storageAccounts : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
        } else {
            const args = argsOrState as InteractiveQueryClusterArgs | undefined;
            if ((!args || args.clusterVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterVersion'");
            }
            if ((!args || args.componentVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'componentVersion'");
            }
            if ((!args || args.gateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gateway'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            inputs["clusterVersion"] = args ? args.clusterVersion : undefined;
            inputs["componentVersion"] = args ? args.componentVersion : undefined;
            inputs["gateway"] = args ? args.gateway : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metastores"] = args ? args.metastores : undefined;
            inputs["monitor"] = args ? args.monitor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["storageAccountGen2"] = args ? args.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = args ? args.storageAccounts : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            inputs["httpsEndpoint"] = undefined /*out*/;
            inputs["sshEndpoint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InteractiveQueryCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InteractiveQueryCluster resources.
 */
export interface InteractiveQueryClusterState {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    readonly clusterVersion?: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    readonly componentVersion?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    readonly gateway?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Interactive Query Cluster.
     */
    readonly httpsEndpoint?: pulumi.Input<string>;
    /**
     * Specifies the Azure Region which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    readonly metastores?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    readonly monitor?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Interactive Query Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `network` block as defined below.
     */
    readonly network?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterNetwork>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `roles` block as defined below.
     */
    readonly roles?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Interactive Query Cluster.
     */
    readonly sshEndpoint?: pulumi.Input<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    readonly storageAccountGen2?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.InteractiveQueryClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Interactive Query Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Interactive Query Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    readonly tier?: pulumi.Input<string>;
    readonly tlsMinVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InteractiveQueryCluster resource.
 */
export interface InteractiveQueryClusterArgs {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    readonly clusterVersion: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    readonly componentVersion: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    readonly gateway: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterGateway>;
    /**
     * Specifies the Azure Region which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    readonly metastores?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    readonly monitor?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Interactive Query Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `network` block as defined below.
     */
    readonly network?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterNetwork>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Interactive Query Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `roles` block as defined below.
     */
    readonly roles: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterRoles>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    readonly storageAccountGen2?: pulumi.Input<inputs.hdinsight.InteractiveQueryClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.InteractiveQueryClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Interactive Query Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Interactive Query Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    readonly tier: pulumi.Input<string>;
    readonly tlsMinVersion?: pulumi.Input<string>;
}
