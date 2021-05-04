// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a NetApp Volume.
 *
 * ## NetApp Volume Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.2.0/24"],
 *     delegations: [{
 *         name: "netapp",
 *         serviceDelegation: {
 *             name: "Microsoft.Netapp/volumes",
 *             actions: [
 *                 "Microsoft.Network/networkinterfaces/*",
 *                 "Microsoft.Network/virtualNetworks/subnets/join/action",
 *             ],
 *         },
 *     }],
 * });
 * const exampleAccount = new azure.netapp.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const examplePool = new azure.netapp.Pool("examplePool", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     accountName: exampleAccount.name,
 *     serviceLevel: "Premium",
 *     sizeInTb: 4,
 * });
 * const exampleVolume = new azure.netapp.Volume("exampleVolume", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     accountName: exampleAccount.name,
 *     poolName: examplePool.name,
 *     volumePath: "my-unique-file-path",
 *     serviceLevel: "Premium",
 *     subnetId: exampleSubnet.id,
 *     protocols: ["NFSv4.1"],
 *     storageQuotaInGb: 100,
 *     createFromSnapshotResourceId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1",
 *     dataProtectionReplication: {
 *         endpointType: "dst",
 *         remoteVolumeLocation: azurerm_resource_group.example_primary.location,
 *         remoteVolumeResourceId: azurerm_netapp_volume.example_primary.id,
 *         replicationFrequency: "10minutes",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * NetApp Volumes can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:netapp/volume:Volume example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:netapp/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Creates volume from snapshot. Following properties must be the same as the original volume where the snapshot was taken from: `protocols`, `subnetId`, `location`, `serviceLevel`, `resourceGroupName`, `accountName` and `poolName`.
     */
    public readonly createFromSnapshotResourceId!: pulumi.Output<string>;
    /**
     * A `dataProtectionReplication` block as defined below.
     */
    public readonly dataProtectionReplication!: pulumi.Output<outputs.netapp.VolumeDataProtectionReplication | undefined>;
    /**
     * One or more `exportPolicyRule` block defined below.
     */
    public readonly exportPolicyRules!: pulumi.Output<outputs.netapp.VolumeExportPolicyRule[] | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A list of IPv4 Addresses which should be used to mount the volume.
     */
    public /*out*/ readonly mountIpAddresses!: pulumi.Output<string[]>;
    /**
     * The name of the NetApp Volume. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    public readonly poolName!: pulumi.Output<string>;
    /**
     * The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost. Dual protocol scenario is supported for CIFS and NFSv3, for more information, please refer to [Create a dual-protocol volume for Azure NetApp Files](https://docs.microsoft.com/en-us/azure/azure-netapp-files/create-volumes-dual-protocol) document.
     */
    public readonly protocols!: pulumi.Output<string[]>;
    /**
     * The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
     */
    public readonly serviceLevel!: pulumi.Output<string>;
    /**
     * The maximum Storage Quota allowed for a file system in Gigabytes.
     */
    public readonly storageQuotaInGb!: pulumi.Output<number>;
    /**
     * The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
     */
    public readonly volumePath!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["createFromSnapshotResourceId"] = state ? state.createFromSnapshotResourceId : undefined;
            inputs["dataProtectionReplication"] = state ? state.dataProtectionReplication : undefined;
            inputs["exportPolicyRules"] = state ? state.exportPolicyRules : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["mountIpAddresses"] = state ? state.mountIpAddresses : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["poolName"] = state ? state.poolName : undefined;
            inputs["protocols"] = state ? state.protocols : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serviceLevel"] = state ? state.serviceLevel : undefined;
            inputs["storageQuotaInGb"] = state ? state.storageQuotaInGb : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["volumePath"] = state ? state.volumePath : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.poolName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceLevel'");
            }
            if ((!args || args.storageQuotaInGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageQuotaInGb'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.volumePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumePath'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["createFromSnapshotResourceId"] = args ? args.createFromSnapshotResourceId : undefined;
            inputs["dataProtectionReplication"] = args ? args.dataProtectionReplication : undefined;
            inputs["exportPolicyRules"] = args ? args.exportPolicyRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["poolName"] = args ? args.poolName : undefined;
            inputs["protocols"] = args ? args.protocols : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceLevel"] = args ? args.serviceLevel : undefined;
            inputs["storageQuotaInGb"] = args ? args.storageQuotaInGb : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["volumePath"] = args ? args.volumePath : undefined;
            inputs["mountIpAddresses"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * Creates volume from snapshot. Following properties must be the same as the original volume where the snapshot was taken from: `protocols`, `subnetId`, `location`, `serviceLevel`, `resourceGroupName`, `accountName` and `poolName`.
     */
    readonly createFromSnapshotResourceId?: pulumi.Input<string>;
    /**
     * A `dataProtectionReplication` block as defined below.
     */
    readonly dataProtectionReplication?: pulumi.Input<inputs.netapp.VolumeDataProtectionReplication>;
    /**
     * One or more `exportPolicyRule` block defined below.
     */
    readonly exportPolicyRules?: pulumi.Input<pulumi.Input<inputs.netapp.VolumeExportPolicyRule>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A list of IPv4 Addresses which should be used to mount the volume.
     */
    readonly mountIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the NetApp Volume. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly poolName?: pulumi.Input<string>;
    /**
     * The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost. Dual protocol scenario is supported for CIFS and NFSv3, for more information, please refer to [Create a dual-protocol volume for Azure NetApp Files](https://docs.microsoft.com/en-us/azure/azure-netapp-files/create-volumes-dual-protocol) document.
     */
    readonly protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
     */
    readonly serviceLevel?: pulumi.Input<string>;
    /**
     * The maximum Storage Quota allowed for a file system in Gigabytes.
     */
    readonly storageQuotaInGb?: pulumi.Input<number>;
    /**
     * The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
     */
    readonly volumePath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Creates volume from snapshot. Following properties must be the same as the original volume where the snapshot was taken from: `protocols`, `subnetId`, `location`, `serviceLevel`, `resourceGroupName`, `accountName` and `poolName`.
     */
    readonly createFromSnapshotResourceId?: pulumi.Input<string>;
    /**
     * A `dataProtectionReplication` block as defined below.
     */
    readonly dataProtectionReplication?: pulumi.Input<inputs.netapp.VolumeDataProtectionReplication>;
    /**
     * One or more `exportPolicyRule` block defined below.
     */
    readonly exportPolicyRules?: pulumi.Input<pulumi.Input<inputs.netapp.VolumeExportPolicyRule>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the NetApp Volume. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly poolName: pulumi.Input<string>;
    /**
     * The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost. Dual protocol scenario is supported for CIFS and NFSv3, for more information, please refer to [Create a dual-protocol volume for Azure NetApp Files](https://docs.microsoft.com/en-us/azure/azure-netapp-files/create-volumes-dual-protocol) document.
     */
    readonly protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
     */
    readonly serviceLevel: pulumi.Input<string>;
    /**
     * The maximum Storage Quota allowed for a file system in Gigabytes.
     */
    readonly storageQuotaInGb: pulumi.Input<number>;
    /**
     * The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
     */
    readonly volumePath: pulumi.Input<string>;
}
