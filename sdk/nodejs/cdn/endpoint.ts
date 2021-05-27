// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * CDN Endpoints can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:cdn/endpoint:Endpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
 * ```
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:cdn/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
     */
    public readonly contentTypesToCompresses!: pulumi.Output<string[]>;
    /**
     * Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
     */
    public readonly deliveryRules!: pulumi.Output<outputs.cdn.EndpointDeliveryRule[] | undefined>;
    /**
     * A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
     */
    public readonly geoFilters!: pulumi.Output<outputs.cdn.EndpointGeoFilter[] | undefined>;
    /**
     * Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
     */
    public readonly globalDeliveryRule!: pulumi.Output<outputs.cdn.EndpointGlobalDeliveryRule | undefined>;
    /**
     * A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
     */
    public /*out*/ readonly hostName!: pulumi.Output<string>;
    /**
     * Indicates whether compression is to be enabled.
     */
    public readonly isCompressionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to `true`.
     */
    public readonly isHttpAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Defaults to `true`.
     */
    public readonly isHttpsAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
     */
    public readonly optimizationType!: pulumi.Output<string | undefined>;
    /**
     * The host header CDN provider will send along with content requests to origins.
     */
    public readonly originHostHeader!: pulumi.Output<string | undefined>;
    /**
     * The path used at for origin requests.
     */
    public readonly originPath!: pulumi.Output<string>;
    /**
     * The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
     */
    public readonly origins!: pulumi.Output<outputs.cdn.EndpointOrigin[]>;
    /**
     * the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
     */
    public readonly probePath!: pulumi.Output<string>;
    /**
     * The CDN Profile to which to attach the CDN Endpoint.
     */
    public readonly profileName!: pulumi.Output<string>;
    /**
     * Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
     */
    public readonly querystringCachingBehaviour!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the CDN Endpoint.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["contentTypesToCompresses"] = state ? state.contentTypesToCompresses : undefined;
            inputs["deliveryRules"] = state ? state.deliveryRules : undefined;
            inputs["geoFilters"] = state ? state.geoFilters : undefined;
            inputs["globalDeliveryRule"] = state ? state.globalDeliveryRule : undefined;
            inputs["hostName"] = state ? state.hostName : undefined;
            inputs["isCompressionEnabled"] = state ? state.isCompressionEnabled : undefined;
            inputs["isHttpAllowed"] = state ? state.isHttpAllowed : undefined;
            inputs["isHttpsAllowed"] = state ? state.isHttpsAllowed : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["optimizationType"] = state ? state.optimizationType : undefined;
            inputs["originHostHeader"] = state ? state.originHostHeader : undefined;
            inputs["originPath"] = state ? state.originPath : undefined;
            inputs["origins"] = state ? state.origins : undefined;
            inputs["probePath"] = state ? state.probePath : undefined;
            inputs["profileName"] = state ? state.profileName : undefined;
            inputs["querystringCachingBehaviour"] = state ? state.querystringCachingBehaviour : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if ((!args || args.origins === undefined) && !opts.urn) {
                throw new Error("Missing required property 'origins'");
            }
            if ((!args || args.profileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profileName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["contentTypesToCompresses"] = args ? args.contentTypesToCompresses : undefined;
            inputs["deliveryRules"] = args ? args.deliveryRules : undefined;
            inputs["geoFilters"] = args ? args.geoFilters : undefined;
            inputs["globalDeliveryRule"] = args ? args.globalDeliveryRule : undefined;
            inputs["isCompressionEnabled"] = args ? args.isCompressionEnabled : undefined;
            inputs["isHttpAllowed"] = args ? args.isHttpAllowed : undefined;
            inputs["isHttpsAllowed"] = args ? args.isHttpsAllowed : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["optimizationType"] = args ? args.optimizationType : undefined;
            inputs["originHostHeader"] = args ? args.originHostHeader : undefined;
            inputs["originPath"] = args ? args.originPath : undefined;
            inputs["origins"] = args ? args.origins : undefined;
            inputs["probePath"] = args ? args.probePath : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["querystringCachingBehaviour"] = args ? args.querystringCachingBehaviour : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["hostName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
     */
    contentTypesToCompresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
     */
    deliveryRules?: pulumi.Input<pulumi.Input<inputs.cdn.EndpointDeliveryRule>[]>;
    /**
     * A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
     */
    geoFilters?: pulumi.Input<pulumi.Input<inputs.cdn.EndpointGeoFilter>[]>;
    /**
     * Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
     */
    globalDeliveryRule?: pulumi.Input<inputs.cdn.EndpointGlobalDeliveryRule>;
    /**
     * A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
     */
    hostName?: pulumi.Input<string>;
    /**
     * Indicates whether compression is to be enabled.
     */
    isCompressionEnabled?: pulumi.Input<boolean>;
    /**
     * Defaults to `true`.
     */
    isHttpAllowed?: pulumi.Input<boolean>;
    /**
     * Defaults to `true`.
     */
    isHttpsAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
     */
    optimizationType?: pulumi.Input<string>;
    /**
     * The host header CDN provider will send along with content requests to origins.
     */
    originHostHeader?: pulumi.Input<string>;
    /**
     * The path used at for origin requests.
     */
    originPath?: pulumi.Input<string>;
    /**
     * The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
     */
    origins?: pulumi.Input<pulumi.Input<inputs.cdn.EndpointOrigin>[]>;
    /**
     * the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
     */
    probePath?: pulumi.Input<string>;
    /**
     * The CDN Profile to which to attach the CDN Endpoint.
     */
    profileName?: pulumi.Input<string>;
    /**
     * Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
     */
    querystringCachingBehaviour?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the CDN Endpoint.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
     */
    contentTypesToCompresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
     */
    deliveryRules?: pulumi.Input<pulumi.Input<inputs.cdn.EndpointDeliveryRule>[]>;
    /**
     * A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
     */
    geoFilters?: pulumi.Input<pulumi.Input<inputs.cdn.EndpointGeoFilter>[]>;
    /**
     * Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
     */
    globalDeliveryRule?: pulumi.Input<inputs.cdn.EndpointGlobalDeliveryRule>;
    /**
     * Indicates whether compression is to be enabled.
     */
    isCompressionEnabled?: pulumi.Input<boolean>;
    /**
     * Defaults to `true`.
     */
    isHttpAllowed?: pulumi.Input<boolean>;
    /**
     * Defaults to `true`.
     */
    isHttpsAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
     */
    optimizationType?: pulumi.Input<string>;
    /**
     * The host header CDN provider will send along with content requests to origins.
     */
    originHostHeader?: pulumi.Input<string>;
    /**
     * The path used at for origin requests.
     */
    originPath?: pulumi.Input<string>;
    /**
     * The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
     */
    origins: pulumi.Input<pulumi.Input<inputs.cdn.EndpointOrigin>[]>;
    /**
     * the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
     */
    probePath?: pulumi.Input<string>;
    /**
     * The CDN Profile to which to attach the CDN Endpoint.
     */
    profileName: pulumi.Input<string>;
    /**
     * Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
     */
    querystringCachingBehaviour?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the CDN Endpoint.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
