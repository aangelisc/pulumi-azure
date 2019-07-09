// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Stream Analytics Stream Input Blob.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleAccount = new azure.storage.Account("example", {
 *     accountReplicationType: "LRS",
 *     accountTier: "Standard",
 *     location: azurerm_resource_group_example.location,
 *     name: "examplestoracc",
 *     resourceGroupName: azurerm_resource_group_example.name,
 * });
 * const exampleResourceGroup = pulumi.output(azure.core.getResourceGroup({
 *     name: "example-resources",
 * }));
 * const exampleJob = azurerm_resource_group_example.name.apply(name => azure.streamanalytics.getJob({
 *     name: "example-job",
 *     resourceGroupName: name,
 * }));
 * const exampleContainer = new azure.storage.Container("example", {
 *     containerAccessType: "private",
 *     name: "example",
 *     resourceGroupName: azurerm_resource_group_example.name,
 *     storageAccountName: exampleAccount.name,
 * });
 * const test = new azure.streamanalytics.StreamInputEventHub("test", {
 *     dateFormat: "yyyy/MM/dd",
 *     name: "eventhub-stream-input",
 *     pathPattern: "some-random-pattern",
 *     resourceGroupName: exampleJob.resourceGroupName,
 *     serialization: {
 *         encoding: "UTF8",
 *         type: "Json",
 *     },
 *     storageAccountKey: exampleAccount.primaryAccessKey,
 *     storageAccountName: exampleAccount.name,
 *     storageContainerName: exampleContainer.name,
 *     streamAnalyticsJobName: exampleJob.name,
 *     timeFormat: "HH",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_blob.html.markdown.
 */
export class StreamInputBlob extends pulumi.CustomResource {
    /**
     * Get an existing StreamInputBlob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamInputBlobState, opts?: pulumi.CustomResourceOptions): StreamInputBlob {
        return new StreamInputBlob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:streamanalytics/streamInputBlob:StreamInputBlob';

    /**
     * Returns true if the given object is an instance of StreamInputBlob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamInputBlob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamInputBlob.__pulumiType;
    }

    /**
     * The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
     */
    public readonly dateFormat!: pulumi.Output<string>;
    /**
     * The name of the Stream Input Blob. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
     */
    public readonly pathPattern!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `serialization` block as defined below.
     */
    public readonly serialization!: pulumi.Output<{ encoding?: string, fieldDelimiter?: string, type: string }>;
    /**
     * The Access Key which should be used to connect to this Storage Account.
     */
    public readonly storageAccountKey!: pulumi.Output<string>;
    /**
     * The name of the Storage Account.
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * The name of the Container within the Storage Account.
     */
    public readonly storageContainerName!: pulumi.Output<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    public readonly streamAnalyticsJobName!: pulumi.Output<string>;
    /**
     * The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
     */
    public readonly timeFormat!: pulumi.Output<string>;

    /**
     * Create a StreamInputBlob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamInputBlobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamInputBlobArgs | StreamInputBlobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StreamInputBlobState | undefined;
            inputs["dateFormat"] = state ? state.dateFormat : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pathPattern"] = state ? state.pathPattern : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serialization"] = state ? state.serialization : undefined;
            inputs["storageAccountKey"] = state ? state.storageAccountKey : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            inputs["storageContainerName"] = state ? state.storageContainerName : undefined;
            inputs["streamAnalyticsJobName"] = state ? state.streamAnalyticsJobName : undefined;
            inputs["timeFormat"] = state ? state.timeFormat : undefined;
        } else {
            const args = argsOrState as StreamInputBlobArgs | undefined;
            if (!args || args.dateFormat === undefined) {
                throw new Error("Missing required property 'dateFormat'");
            }
            if (!args || args.pathPattern === undefined) {
                throw new Error("Missing required property 'pathPattern'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serialization === undefined) {
                throw new Error("Missing required property 'serialization'");
            }
            if (!args || args.storageAccountKey === undefined) {
                throw new Error("Missing required property 'storageAccountKey'");
            }
            if (!args || args.storageAccountName === undefined) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            if (!args || args.storageContainerName === undefined) {
                throw new Error("Missing required property 'storageContainerName'");
            }
            if (!args || args.streamAnalyticsJobName === undefined) {
                throw new Error("Missing required property 'streamAnalyticsJobName'");
            }
            if (!args || args.timeFormat === undefined) {
                throw new Error("Missing required property 'timeFormat'");
            }
            inputs["dateFormat"] = args ? args.dateFormat : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pathPattern"] = args ? args.pathPattern : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serialization"] = args ? args.serialization : undefined;
            inputs["storageAccountKey"] = args ? args.storageAccountKey : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            inputs["storageContainerName"] = args ? args.storageContainerName : undefined;
            inputs["streamAnalyticsJobName"] = args ? args.streamAnalyticsJobName : undefined;
            inputs["timeFormat"] = args ? args.timeFormat : undefined;
        }
        super(StreamInputBlob.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamInputBlob resources.
 */
export interface StreamInputBlobState {
    /**
     * The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
     */
    readonly dateFormat?: pulumi.Input<string>;
    /**
     * The name of the Stream Input Blob. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
     */
    readonly pathPattern?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `serialization` block as defined below.
     */
    readonly serialization?: pulumi.Input<{ encoding?: pulumi.Input<string>, fieldDelimiter?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The Access Key which should be used to connect to this Storage Account.
     */
    readonly storageAccountKey?: pulumi.Input<string>;
    /**
     * The name of the Storage Account.
     */
    readonly storageAccountName?: pulumi.Input<string>;
    /**
     * The name of the Container within the Storage Account.
     */
    readonly storageContainerName?: pulumi.Input<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    readonly streamAnalyticsJobName?: pulumi.Input<string>;
    /**
     * The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
     */
    readonly timeFormat?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StreamInputBlob resource.
 */
export interface StreamInputBlobArgs {
    /**
     * The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
     */
    readonly dateFormat: pulumi.Input<string>;
    /**
     * The name of the Stream Input Blob. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
     */
    readonly pathPattern: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `serialization` block as defined below.
     */
    readonly serialization: pulumi.Input<{ encoding?: pulumi.Input<string>, fieldDelimiter?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The Access Key which should be used to connect to this Storage Account.
     */
    readonly storageAccountKey: pulumi.Input<string>;
    /**
     * The name of the Storage Account.
     */
    readonly storageAccountName: pulumi.Input<string>;
    /**
     * The name of the Container within the Storage Account.
     */
    readonly storageContainerName: pulumi.Input<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    readonly streamAnalyticsJobName: pulumi.Input<string>;
    /**
     * The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
     */
    readonly timeFormat: pulumi.Input<string>;
}
