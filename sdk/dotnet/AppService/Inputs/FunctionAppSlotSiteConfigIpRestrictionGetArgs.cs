// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class FunctionAppSlotSiteConfigIpRestrictionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this restriction `Allow` or `Deny` access for this IP range. Defaults to `Allow`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The headers for this specific `ip_restriction` as defined below.
        /// </summary>
        [Input("headers")]
        public Input<Inputs.FunctionAppSlotSiteConfigIpRestrictionHeadersGetArgs>? Headers { get; set; }

        /// <summary>
        /// The IP Address used for this IP Restriction in CIDR notation.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The name for this IP Restriction.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority for this IP Restriction. Restrictions are enforced in priority order. By default, priority is set to 65000 if not specified.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The Service Tag used for this IP Restriction.
        /// </summary>
        [Input("serviceTag")]
        public Input<string>? ServiceTag { get; set; }

        /// <summary>
        /// The Virtual Network Subnet ID used for this IP Restriction.
        /// </summary>
        [Input("virtualNetworkSubnetId")]
        public Input<string>? VirtualNetworkSubnetId { get; set; }

        public FunctionAppSlotSiteConfigIpRestrictionGetArgs()
        {
        }
    }
}
