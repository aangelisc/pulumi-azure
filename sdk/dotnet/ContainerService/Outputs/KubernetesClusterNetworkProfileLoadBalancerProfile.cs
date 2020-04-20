// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class KubernetesClusterNetworkProfileLoadBalancerProfile
    {
        /// <summary>
        /// The outcome (resource IDs) of the specified arguments.
        /// </summary>
        public readonly ImmutableArray<string> EffectiveOutboundIps;
        /// <summary>
        /// Count of desired managed outbound IPs for the cluster load balancer. Must be in the range of [1, 100].
        /// </summary>
        public readonly int? ManagedOutboundIpCount;
        /// <summary>
        /// The ID of the Public IP Addresses which should be used for outbound communication for the cluster load balancer.
        /// </summary>
        public readonly ImmutableArray<string> OutboundIpAddressIds;
        /// <summary>
        /// The ID of the outbound Public IP Address Prefixes which should be used for the cluster load balancer.
        /// </summary>
        public readonly ImmutableArray<string> OutboundIpPrefixIds;

        [OutputConstructor]
        private KubernetesClusterNetworkProfileLoadBalancerProfile(
            ImmutableArray<string> effectiveOutboundIps,

            int? managedOutboundIpCount,

            ImmutableArray<string> outboundIpAddressIds,

            ImmutableArray<string> outboundIpPrefixIds)
        {
            EffectiveOutboundIps = effectiveOutboundIps;
            ManagedOutboundIpCount = managedOutboundIpCount;
            OutboundIpAddressIds = outboundIpAddressIds;
            OutboundIpPrefixIds = outboundIpPrefixIds;
        }
    }
}