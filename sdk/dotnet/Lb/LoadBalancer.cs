// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb
{
    /// <summary>
    /// Manages a Load Balancer Resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb.html.markdown.
    /// </summary>
    public partial class LoadBalancer : Pulumi.CustomResource
    {
        /// <summary>
        /// One or multiple `frontend_ip_configuration` blocks as documented below.
        /// </summary>
        [Output("frontendIpConfigurations")]
        public Output<ImmutableArray<Outputs.LoadBalancerFrontendIpConfigurations>> FrontendIpConfigurations { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure Region where the Load Balancer should be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the frontend ip configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The list of private IP address assigned to the load balancer in `frontend_ip_configuration` blocks, if any.
        /// </summary>
        [Output("privateIpAddresses")]
        public Output<ImmutableArray<string>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which to create the Load Balancer.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs args, CustomResourceOptions? options = null)
            : base("azure:lb/loadBalancer:LoadBalancer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
            : base("azure:lb/loadBalancer:LoadBalancer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, state, options);
        }
    }

    public sealed class LoadBalancerArgs : Pulumi.ResourceArgs
    {
        [Input("frontendIpConfigurations")]
        private InputList<Inputs.LoadBalancerFrontendIpConfigurationsArgs>? _frontendIpConfigurations;

        /// <summary>
        /// One or multiple `frontend_ip_configuration` blocks as documented below.
        /// </summary>
        public InputList<Inputs.LoadBalancerFrontendIpConfigurationsArgs> FrontendIpConfigurations
        {
            get => _frontendIpConfigurations ?? (_frontendIpConfigurations = new InputList<Inputs.LoadBalancerFrontendIpConfigurationsArgs>());
            set => _frontendIpConfigurations = value;
        }

        /// <summary>
        /// Specifies the supported Azure Region where the Load Balancer should be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the frontend ip configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group in which to create the Load Balancer.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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

        public LoadBalancerArgs()
        {
        }
    }

    public sealed class LoadBalancerState : Pulumi.ResourceArgs
    {
        [Input("frontendIpConfigurations")]
        private InputList<Inputs.LoadBalancerFrontendIpConfigurationsGetArgs>? _frontendIpConfigurations;

        /// <summary>
        /// One or multiple `frontend_ip_configuration` blocks as documented below.
        /// </summary>
        public InputList<Inputs.LoadBalancerFrontendIpConfigurationsGetArgs> FrontendIpConfigurations
        {
            get => _frontendIpConfigurations ?? (_frontendIpConfigurations = new InputList<Inputs.LoadBalancerFrontendIpConfigurationsGetArgs>());
            set => _frontendIpConfigurations = value;
        }

        /// <summary>
        /// Specifies the supported Azure Region where the Load Balancer should be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the frontend ip configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;

        /// <summary>
        /// The list of private IP address assigned to the load balancer in `frontend_ip_configuration` blocks, if any.
        /// </summary>
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        /// <summary>
        /// The name of the Resource Group in which to create the Load Balancer.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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

        public LoadBalancerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class LoadBalancerFrontendIpConfigurationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the Frontend IP Configuration.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("inboundNatRules")]
        private InputList<string>? _inboundNatRules;
        public InputList<string> InboundNatRules
        {
            get => _inboundNatRules ?? (_inboundNatRules = new InputList<string>());
            set => _inboundNatRules = value;
        }

        [Input("loadBalancerRules")]
        private InputList<string>? _loadBalancerRules;
        public InputList<string> LoadBalancerRules
        {
            get => _loadBalancerRules ?? (_loadBalancerRules = new InputList<string>());
            set => _loadBalancerRules = value;
        }

        /// <summary>
        /// Specifies the name of the frontend ip configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("outboundRules")]
        private InputList<string>? _outboundRules;
        public InputList<string> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<string>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
        /// </summary>
        [Input("privateIpAddressAllocation")]
        public Input<string>? PrivateIpAddressAllocation { get; set; }

        /// <summary>
        /// The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
        /// </summary>
        [Input("privateIpAddressVersion")]
        public Input<string>? PrivateIpAddressVersion { get; set; }

        /// <summary>
        /// The ID of a Public IP Address which should be associated with the Load Balancer.
        /// </summary>
        [Input("publicIpAddressId")]
        public Input<string>? PublicIpAddressId { get; set; }

        /// <summary>
        /// The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
        /// </summary>
        [Input("publicIpPrefixId")]
        public Input<string>? PublicIpPrefixId { get; set; }

        /// <summary>
        /// The ID of the Subnet which should be associated with the IP Configuration.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public LoadBalancerFrontendIpConfigurationsArgs()
        {
        }
    }

    public sealed class LoadBalancerFrontendIpConfigurationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the Frontend IP Configuration.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("inboundNatRules")]
        private InputList<string>? _inboundNatRules;
        public InputList<string> InboundNatRules
        {
            get => _inboundNatRules ?? (_inboundNatRules = new InputList<string>());
            set => _inboundNatRules = value;
        }

        [Input("loadBalancerRules")]
        private InputList<string>? _loadBalancerRules;
        public InputList<string> LoadBalancerRules
        {
            get => _loadBalancerRules ?? (_loadBalancerRules = new InputList<string>());
            set => _loadBalancerRules = value;
        }

        /// <summary>
        /// Specifies the name of the frontend ip configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("outboundRules")]
        private InputList<string>? _outboundRules;
        public InputList<string> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<string>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
        /// </summary>
        [Input("privateIpAddressAllocation")]
        public Input<string>? PrivateIpAddressAllocation { get; set; }

        /// <summary>
        /// The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
        /// </summary>
        [Input("privateIpAddressVersion")]
        public Input<string>? PrivateIpAddressVersion { get; set; }

        /// <summary>
        /// The ID of a Public IP Address which should be associated with the Load Balancer.
        /// </summary>
        [Input("publicIpAddressId")]
        public Input<string>? PublicIpAddressId { get; set; }

        /// <summary>
        /// The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
        /// </summary>
        [Input("publicIpPrefixId")]
        public Input<string>? PublicIpPrefixId { get; set; }

        /// <summary>
        /// The ID of the Subnet which should be associated with the IP Configuration.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public LoadBalancerFrontendIpConfigurationsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class LoadBalancerFrontendIpConfigurations
    {
        /// <summary>
        /// The id of the Frontend IP Configuration.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InboundNatRules;
        public readonly ImmutableArray<string> LoadBalancerRules;
        /// <summary>
        /// Specifies the name of the frontend ip configuration.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> OutboundRules;
        /// <summary>
        /// Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
        /// </summary>
        public readonly string PrivateIpAddress;
        /// <summary>
        /// The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
        /// </summary>
        public readonly string PrivateIpAddressAllocation;
        /// <summary>
        /// The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
        /// </summary>
        public readonly string? PrivateIpAddressVersion;
        /// <summary>
        /// The ID of a Public IP Address which should be associated with the Load Balancer.
        /// </summary>
        public readonly string PublicIpAddressId;
        /// <summary>
        /// The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
        /// </summary>
        public readonly string PublicIpPrefixId;
        /// <summary>
        /// The ID of the Subnet which should be associated with the IP Configuration.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
        /// </summary>
        public readonly string? Zones;

        [OutputConstructor]
        private LoadBalancerFrontendIpConfigurations(
            string id,
            ImmutableArray<string> inboundNatRules,
            ImmutableArray<string> loadBalancerRules,
            string name,
            ImmutableArray<string> outboundRules,
            string privateIpAddress,
            string privateIpAddressAllocation,
            string? privateIpAddressVersion,
            string publicIpAddressId,
            string publicIpPrefixId,
            string subnetId,
            string? zones)
        {
            Id = id;
            InboundNatRules = inboundNatRules;
            LoadBalancerRules = loadBalancerRules;
            Name = name;
            OutboundRules = outboundRules;
            PrivateIpAddress = privateIpAddress;
            PrivateIpAddressAllocation = privateIpAddressAllocation;
            PrivateIpAddressVersion = privateIpAddressVersion;
            PublicIpAddressId = publicIpAddressId;
            PublicIpPrefixId = publicIpPrefixId;
            SubnetId = subnetId;
            Zones = zones;
        }
    }
    }
}
