// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkGatewayConnectionTrafficSelectorPolicy
    {
        public readonly ImmutableArray<string> LocalAddressCidrs;
        public readonly ImmutableArray<string> RemoteAddressCidrs;

        [OutputConstructor]
        private VirtualNetworkGatewayConnectionTrafficSelectorPolicy(
            ImmutableArray<string> localAddressCidrs,

            ImmutableArray<string> remoteAddressCidrs)
        {
            LocalAddressCidrs = localAddressCidrs;
            RemoteAddressCidrs = remoteAddressCidrs;
        }
    }
}
