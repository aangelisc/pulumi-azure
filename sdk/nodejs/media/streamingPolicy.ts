// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Streaming Policy.
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
 *     accountReplicationType: "GRS",
 * });
 * const exampleServiceAccount = new azure.media.ServiceAccount("exampleServiceAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     storageAccounts: [{
 *         id: exampleAccount.id,
 *         isPrimary: true,
 *     }],
 * });
 * const exampleStreamingPolicy = new azure.media.StreamingPolicy("exampleStreamingPolicy", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     mediaServicesAccountName: exampleServiceAccount.name,
 *     commonEncryptionCenc: {
 *         enabledProtocols: {
 *             download: false,
 *             dash: true,
 *             hls: false,
 *             smoothStreaming: false,
 *         },
 *         drmPlayready: {
 *             customLicenseAcquisitionUrlTemplate: "https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}",
 *             customAttributes: "PlayReady CustomAttributes",
 *         },
 *         drmWidevineCustomLicenseAcquisitionUrlTemplate: "https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId}",
 *     },
 *     commonEncryptionCbcs: {
 *         enabledProtocols: {
 *             download: false,
 *             dash: true,
 *             hls: false,
 *             smoothStreaming: false,
 *         },
 *         drmFairplay: {
 *             customLicenseAcquisitionUrlTemplate: "https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}",
 *             allowPersistentLicense: true,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Streaming Policys can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:media/streamingPolicy:StreamingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/streamingpolicies/policy1
 * ```
 */
export class StreamingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing StreamingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamingPolicyState, opts?: pulumi.CustomResourceOptions): StreamingPolicy {
        return new StreamingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:media/streamingPolicy:StreamingPolicy';

    /**
     * Returns true if the given object is an instance of StreamingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamingPolicy.__pulumiType;
    }

    /**
     * A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    public readonly commonEncryptionCbcs!: pulumi.Output<outputs.media.StreamingPolicyCommonEncryptionCbcs | undefined>;
    /**
     * A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    public readonly commonEncryptionCenc!: pulumi.Output<outputs.media.StreamingPolicyCommonEncryptionCenc | undefined>;
    /**
     * Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
     */
    public readonly defaultContentKeyPolicyName!: pulumi.Output<string | undefined>;
    /**
     * The Media Services account name. Changing this forces a new Streaming Policy to be created.
     */
    public readonly mediaServicesAccountName!: pulumi.Output<string>;
    /**
     * The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    public readonly noEncryptionEnabledProtocols!: pulumi.Output<outputs.media.StreamingPolicyNoEncryptionEnabledProtocols | undefined>;
    /**
     * The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a StreamingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamingPolicyArgs | StreamingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StreamingPolicyState | undefined;
            inputs["commonEncryptionCbcs"] = state ? state.commonEncryptionCbcs : undefined;
            inputs["commonEncryptionCenc"] = state ? state.commonEncryptionCenc : undefined;
            inputs["defaultContentKeyPolicyName"] = state ? state.defaultContentKeyPolicyName : undefined;
            inputs["mediaServicesAccountName"] = state ? state.mediaServicesAccountName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["noEncryptionEnabledProtocols"] = state ? state.noEncryptionEnabledProtocols : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as StreamingPolicyArgs | undefined;
            if ((!args || args.mediaServicesAccountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'mediaServicesAccountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["commonEncryptionCbcs"] = args ? args.commonEncryptionCbcs : undefined;
            inputs["commonEncryptionCenc"] = args ? args.commonEncryptionCenc : undefined;
            inputs["defaultContentKeyPolicyName"] = args ? args.defaultContentKeyPolicyName : undefined;
            inputs["mediaServicesAccountName"] = args ? args.mediaServicesAccountName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["noEncryptionEnabledProtocols"] = args ? args.noEncryptionEnabledProtocols : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StreamingPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamingPolicy resources.
 */
export interface StreamingPolicyState {
    /**
     * A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    readonly commonEncryptionCbcs?: pulumi.Input<inputs.media.StreamingPolicyCommonEncryptionCbcs>;
    /**
     * A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    readonly commonEncryptionCenc?: pulumi.Input<inputs.media.StreamingPolicyCommonEncryptionCenc>;
    /**
     * Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
     */
    readonly defaultContentKeyPolicyName?: pulumi.Input<string>;
    /**
     * The Media Services account name. Changing this forces a new Streaming Policy to be created.
     */
    readonly mediaServicesAccountName?: pulumi.Input<string>;
    /**
     * The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    readonly noEncryptionEnabledProtocols?: pulumi.Input<inputs.media.StreamingPolicyNoEncryptionEnabledProtocols>;
    /**
     * The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StreamingPolicy resource.
 */
export interface StreamingPolicyArgs {
    /**
     * A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    readonly commonEncryptionCbcs?: pulumi.Input<inputs.media.StreamingPolicyCommonEncryptionCbcs>;
    /**
     * A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    readonly commonEncryptionCenc?: pulumi.Input<inputs.media.StreamingPolicyCommonEncryptionCenc>;
    /**
     * Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
     */
    readonly defaultContentKeyPolicyName?: pulumi.Input<string>;
    /**
     * The Media Services account name. Changing this forces a new Streaming Policy to be created.
     */
    readonly mediaServicesAccountName: pulumi.Input<string>;
    /**
     * The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
     */
    readonly noEncryptionEnabledProtocols?: pulumi.Input<inputs.media.StreamingPolicyNoEncryptionEnabledProtocols>;
    /**
     * The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}