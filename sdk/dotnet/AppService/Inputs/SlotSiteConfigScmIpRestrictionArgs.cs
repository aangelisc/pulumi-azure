// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class SlotSiteConfigScmIpRestrictionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this restriction `Allow` or `Deny` access for this IP range. Defaults to `Allow`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The IP Address used for this IP Restriction in CIDR notation.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority for this IP Restriction. Restrictions are enforced in priority order. By default, priority is set to 65000 if not specified.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The Virtual Network Subnet ID used for this IP Restriction.
        /// </summary>
        [Input("virtualNetworkSubnetId")]
        public Input<string>? VirtualNetworkSubnetId { get; set; }

        public SlotSiteConfigScmIpRestrictionArgs()
        {
        }
    }
}
