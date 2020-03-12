// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Public IP Address.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/public_ip.html.markdown.
        /// </summary>
        public static Task<GetPublicIPResult> GetPublicIP(GetPublicIPArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicIPResult>("azure:network/getPublicIP:getPublicIP", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetPublicIPArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the public IP address.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        [Input("zones")]
        private List<string>? _zones;
        public List<string> Zones
        {
            get => _zones ?? (_zones = new List<string>());
            set => _zones = value;
        }

        public GetPublicIPArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPublicIPResult
    {
        public readonly string AllocationMethod;
        /// <summary>
        /// The label for the Domain Name.
        /// </summary>
        public readonly string DomainNameLabel;
        /// <summary>
        /// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// Specifies the timeout for the TCP idle connection.
        /// </summary>
        public readonly int IdleTimeoutInMinutes;
        /// <summary>
        /// The IP address value that was allocated.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The IP version being used, for example `IPv4` or `IPv6`.
        /// </summary>
        public readonly string IpVersion;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly string ReverseFqdn;
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly ImmutableArray<string> Zones;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPublicIPResult(
            string allocationMethod,
            string domainNameLabel,
            string fqdn,
            int idleTimeoutInMinutes,
            string ipAddress,
            string ipVersion,
            string location,
            string name,
            string resourceGroupName,
            string reverseFqdn,
            string sku,
            ImmutableDictionary<string, string>? tags,
            ImmutableArray<string> zones,
            string id)
        {
            AllocationMethod = allocationMethod;
            DomainNameLabel = domainNameLabel;
            Fqdn = fqdn;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpAddress = ipAddress;
            IpVersion = ipVersion;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ReverseFqdn = reverseFqdn;
            Sku = sku;
            Tags = tags;
            Zones = zones;
            Id = id;
        }
    }
}
