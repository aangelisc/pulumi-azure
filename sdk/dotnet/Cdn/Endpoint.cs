// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Cdn
{
    /// <summary>
    /// ## Import
    /// 
    /// CDN Endpoints can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:cdn/endpoint:Endpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
    /// ```
    /// </summary>
    [AzureResourceType("azure:cdn/endpoint:Endpoint")]
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
        /// </summary>
        [Output("contentTypesToCompresses")]
        public Output<ImmutableArray<string>> ContentTypesToCompresses { get; private set; } = null!;

        /// <summary>
        /// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `delivery_rule` blocks as defined below.
        /// </summary>
        [Output("deliveryRules")]
        public Output<ImmutableArray<Outputs.EndpointDeliveryRule>> DeliveryRules { get; private set; } = null!;

        /// <summary>
        /// A set of Geo Filters for this CDN Endpoint. Each `geo_filter` block supports fields documented below.
        /// </summary>
        [Output("geoFilters")]
        public Output<ImmutableArray<Outputs.EndpointGeoFilter>> GeoFilters { get; private set; } = null!;

        /// <summary>
        /// Actions that are valid for all resources regardless of any conditions. A `global_delivery_rule` block as defined below.
        /// </summary>
        [Output("globalDeliveryRule")]
        public Output<Outputs.EndpointGlobalDeliveryRule?> GlobalDeliveryRule { get; private set; } = null!;

        /// <summary>
        /// A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether compression is to be enabled.
        /// </summary>
        [Output("isCompressionEnabled")]
        public Output<bool?> IsCompressionEnabled { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("isHttpAllowed")]
        public Output<bool?> IsHttpAllowed { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("isHttpsAllowed")]
        public Output<bool?> IsHttpsAllowed { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
        /// </summary>
        [Output("optimizationType")]
        public Output<string?> OptimizationType { get; private set; } = null!;

        /// <summary>
        /// The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
        /// </summary>
        [Output("originHostHeader")]
        public Output<string?> OriginHostHeader { get; private set; } = null!;

        /// <summary>
        /// The path used at for origin requests.
        /// </summary>
        [Output("originPath")]
        public Output<string> OriginPath { get; private set; } = null!;

        /// <summary>
        /// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
        /// </summary>
        [Output("origins")]
        public Output<ImmutableArray<Outputs.EndpointOrigin>> Origins { get; private set; } = null!;

        /// <summary>
        /// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `origin_path`.
        /// </summary>
        [Output("probePath")]
        public Output<string> ProbePath { get; private set; } = null!;

        /// <summary>
        /// The CDN Profile to which to attach the CDN Endpoint.
        /// </summary>
        [Output("profileName")]
        public Output<string> ProfileName { get; private set; } = null!;

        /// <summary>
        /// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
        /// </summary>
        [Output("querystringCachingBehaviour")]
        public Output<string?> QuerystringCachingBehaviour { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the CDN Endpoint.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("azure:cdn/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("azure:cdn/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        [Input("contentTypesToCompresses")]
        private InputList<string>? _contentTypesToCompresses;

        /// <summary>
        /// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
        /// </summary>
        public InputList<string> ContentTypesToCompresses
        {
            get => _contentTypesToCompresses ?? (_contentTypesToCompresses = new InputList<string>());
            set => _contentTypesToCompresses = value;
        }

        [Input("deliveryRules")]
        private InputList<Inputs.EndpointDeliveryRuleArgs>? _deliveryRules;

        /// <summary>
        /// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `delivery_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.EndpointDeliveryRuleArgs> DeliveryRules
        {
            get => _deliveryRules ?? (_deliveryRules = new InputList<Inputs.EndpointDeliveryRuleArgs>());
            set => _deliveryRules = value;
        }

        [Input("geoFilters")]
        private InputList<Inputs.EndpointGeoFilterArgs>? _geoFilters;

        /// <summary>
        /// A set of Geo Filters for this CDN Endpoint. Each `geo_filter` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.EndpointGeoFilterArgs> GeoFilters
        {
            get => _geoFilters ?? (_geoFilters = new InputList<Inputs.EndpointGeoFilterArgs>());
            set => _geoFilters = value;
        }

        /// <summary>
        /// Actions that are valid for all resources regardless of any conditions. A `global_delivery_rule` block as defined below.
        /// </summary>
        [Input("globalDeliveryRule")]
        public Input<Inputs.EndpointGlobalDeliveryRuleArgs>? GlobalDeliveryRule { get; set; }

        /// <summary>
        /// Indicates whether compression is to be enabled.
        /// </summary>
        [Input("isCompressionEnabled")]
        public Input<bool>? IsCompressionEnabled { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("isHttpAllowed")]
        public Input<bool>? IsHttpAllowed { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("isHttpsAllowed")]
        public Input<bool>? IsHttpsAllowed { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
        /// </summary>
        [Input("optimizationType")]
        public Input<string>? OptimizationType { get; set; }

        /// <summary>
        /// The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
        /// </summary>
        [Input("originHostHeader")]
        public Input<string>? OriginHostHeader { get; set; }

        /// <summary>
        /// The path used at for origin requests.
        /// </summary>
        [Input("originPath")]
        public Input<string>? OriginPath { get; set; }

        [Input("origins", required: true)]
        private InputList<Inputs.EndpointOriginArgs>? _origins;

        /// <summary>
        /// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.EndpointOriginArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.EndpointOriginArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `origin_path`.
        /// </summary>
        [Input("probePath")]
        public Input<string>? ProbePath { get; set; }

        /// <summary>
        /// The CDN Profile to which to attach the CDN Endpoint.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
        /// </summary>
        [Input("querystringCachingBehaviour")]
        public Input<string>? QuerystringCachingBehaviour { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the CDN Endpoint.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        [Input("contentTypesToCompresses")]
        private InputList<string>? _contentTypesToCompresses;

        /// <summary>
        /// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
        /// </summary>
        public InputList<string> ContentTypesToCompresses
        {
            get => _contentTypesToCompresses ?? (_contentTypesToCompresses = new InputList<string>());
            set => _contentTypesToCompresses = value;
        }

        [Input("deliveryRules")]
        private InputList<Inputs.EndpointDeliveryRuleGetArgs>? _deliveryRules;

        /// <summary>
        /// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `delivery_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.EndpointDeliveryRuleGetArgs> DeliveryRules
        {
            get => _deliveryRules ?? (_deliveryRules = new InputList<Inputs.EndpointDeliveryRuleGetArgs>());
            set => _deliveryRules = value;
        }

        [Input("geoFilters")]
        private InputList<Inputs.EndpointGeoFilterGetArgs>? _geoFilters;

        /// <summary>
        /// A set of Geo Filters for this CDN Endpoint. Each `geo_filter` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.EndpointGeoFilterGetArgs> GeoFilters
        {
            get => _geoFilters ?? (_geoFilters = new InputList<Inputs.EndpointGeoFilterGetArgs>());
            set => _geoFilters = value;
        }

        /// <summary>
        /// Actions that are valid for all resources regardless of any conditions. A `global_delivery_rule` block as defined below.
        /// </summary>
        [Input("globalDeliveryRule")]
        public Input<Inputs.EndpointGlobalDeliveryRuleGetArgs>? GlobalDeliveryRule { get; set; }

        /// <summary>
        /// A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// Indicates whether compression is to be enabled.
        /// </summary>
        [Input("isCompressionEnabled")]
        public Input<bool>? IsCompressionEnabled { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("isHttpAllowed")]
        public Input<bool>? IsHttpAllowed { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("isHttpsAllowed")]
        public Input<bool>? IsHttpsAllowed { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
        /// </summary>
        [Input("optimizationType")]
        public Input<string>? OptimizationType { get; set; }

        /// <summary>
        /// The host header CDN provider will send along with content requests to origins. Defaults to the host name of the origin.
        /// </summary>
        [Input("originHostHeader")]
        public Input<string>? OriginHostHeader { get; set; }

        /// <summary>
        /// The path used at for origin requests.
        /// </summary>
        [Input("originPath")]
        public Input<string>? OriginPath { get; set; }

        [Input("origins")]
        private InputList<Inputs.EndpointOriginGetArgs>? _origins;

        /// <summary>
        /// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.EndpointOriginGetArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.EndpointOriginGetArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `origin_path`.
        /// </summary>
        [Input("probePath")]
        public Input<string>? ProbePath { get; set; }

        /// <summary>
        /// The CDN Profile to which to attach the CDN Endpoint.
        /// </summary>
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        /// <summary>
        /// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
        /// </summary>
        [Input("querystringCachingBehaviour")]
        public Input<string>? QuerystringCachingBehaviour { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the CDN Endpoint.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EndpointState()
        {
        }
    }
}
