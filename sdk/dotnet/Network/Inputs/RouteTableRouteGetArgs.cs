// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class RouteTableRouteGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR to which the route applies, such as 10.1.0.0/16. Tags such as `VirtualNetwork`, `AzureLoadBalancer` or `Internet` can also be used.
        /// </summary>
        [Input("addressPrefix", required: true)]
        public Input<string> AddressPrefix { get; set; } = null!;

        /// <summary>
        /// The name of the route.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        /// </summary>
        [Input("nextHopInIpAddress")]
        public Input<string>? NextHopInIpAddress { get; set; }

        /// <summary>
        /// The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`.
        /// </summary>
        [Input("nextHopType", required: true)]
        public Input<string> NextHopType { get; set; } = null!;

        public RouteTableRouteGetArgs()
        {
        }
    }
}
