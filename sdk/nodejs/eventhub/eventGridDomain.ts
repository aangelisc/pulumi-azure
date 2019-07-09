// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EventGrid Domain
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West US 2",
 *     name: "resourceGroup1",
 * });
 * const testEventGridDomain = new azure.eventhub.EventGridDomain("test", {
 *     location: testResourceGroup.location,
 *     name: "my-eventgrid-domain",
 *     resourceGroupName: testResourceGroup.name,
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventgrid_domain.html.markdown.
 */
export class EventGridDomain extends pulumi.CustomResource {
    /**
     * Get an existing EventGridDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventGridDomainState, opts?: pulumi.CustomResourceOptions): EventGridDomain {
        return new EventGridDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventhub/eventGridDomain:EventGridDomain';

    /**
     * Returns true if the given object is an instance of EventGridDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventGridDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventGridDomain.__pulumiType;
    }

    /**
     * The Endpoint associated with the EventGrid Domain.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * A `input_mapping_default_values` block as defined below.
     */
    public readonly inputMappingDefaultValues!: pulumi.Output<{ dataVersion?: string, eventType?: string, subject?: string } | undefined>;
    /**
     * A `input_mapping_fields` block as defined below.
     */
    public readonly inputMappingFields!: pulumi.Output<{ dataVersion?: string, eventTime?: string, eventType?: string, id?: string, subject?: string, topic?: string } | undefined>;
    /**
     * Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
     */
    public readonly inputSchema!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a EventGridDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventGridDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventGridDomainArgs | EventGridDomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventGridDomainState | undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["inputMappingDefaultValues"] = state ? state.inputMappingDefaultValues : undefined;
            inputs["inputMappingFields"] = state ? state.inputMappingFields : undefined;
            inputs["inputSchema"] = state ? state.inputSchema : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventGridDomainArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["inputMappingDefaultValues"] = args ? args.inputMappingDefaultValues : undefined;
            inputs["inputMappingFields"] = args ? args.inputMappingFields : undefined;
            inputs["inputSchema"] = args ? args.inputSchema : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["endpoint"] = undefined /*out*/;
        }
        super(EventGridDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventGridDomain resources.
 */
export interface EventGridDomainState {
    /**
     * The Endpoint associated with the EventGrid Domain.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * A `input_mapping_default_values` block as defined below.
     */
    readonly inputMappingDefaultValues?: pulumi.Input<{ dataVersion?: pulumi.Input<string>, eventType?: pulumi.Input<string>, subject?: pulumi.Input<string> }>;
    /**
     * A `input_mapping_fields` block as defined below.
     */
    readonly inputMappingFields?: pulumi.Input<{ dataVersion?: pulumi.Input<string>, eventTime?: pulumi.Input<string>, eventType?: pulumi.Input<string>, id?: pulumi.Input<string>, subject?: pulumi.Input<string>, topic?: pulumi.Input<string> }>;
    /**
     * Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
     */
    readonly inputSchema?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EventGridDomain resource.
 */
export interface EventGridDomainArgs {
    /**
     * A `input_mapping_default_values` block as defined below.
     */
    readonly inputMappingDefaultValues?: pulumi.Input<{ dataVersion?: pulumi.Input<string>, eventType?: pulumi.Input<string>, subject?: pulumi.Input<string> }>;
    /**
     * A `input_mapping_fields` block as defined below.
     */
    readonly inputMappingFields?: pulumi.Input<{ dataVersion?: pulumi.Input<string>, eventTime?: pulumi.Input<string>, eventType?: pulumi.Input<string>, id?: pulumi.Input<string>, subject?: pulumi.Input<string>, topic?: pulumi.Input<string> }>;
    /**
     * Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
     */
    readonly inputSchema?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
