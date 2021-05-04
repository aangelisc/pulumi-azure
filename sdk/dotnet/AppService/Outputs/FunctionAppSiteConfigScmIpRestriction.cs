// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class FunctionAppSiteConfigScmIpRestriction
    {
        /// <summary>
        /// Allow or Deny access for this IP range. Defaults to Allow.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// The headers for this specific `scm_ip_restriction` as defined below.
        /// </summary>
        public readonly Outputs.FunctionAppSiteConfigScmIpRestrictionHeaders? Headers;
        /// <summary>
        /// The IP Address used for this IP Restriction in CIDR notation.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// The name for this IP Restriction.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The priority for this IP Restriction. Restrictions are enforced in priority order. By default, priority is set to 65000 if not specified.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The Service Tag used for this IP Restriction.
        /// </summary>
        public readonly string? ServiceTag;
        /// <summary>
        /// The Virtual Network Subnet ID used for this IP Restriction.
        /// </summary>
        public readonly string? VirtualNetworkSubnetId;

        [OutputConstructor]
        private FunctionAppSiteConfigScmIpRestriction(
            string? action,

            Outputs.FunctionAppSiteConfigScmIpRestrictionHeaders? headers,

            string? ipAddress,

            string? name,

            int? priority,

            string? serviceTag,

            string? virtualNetworkSubnetId)
        {
            Action = action;
            Headers = headers;
            IpAddress = ipAddress;
            Name = name;
            Priority = priority;
            ServiceTag = serviceTag;
            VirtualNetworkSubnetId = virtualNetworkSubnetId;
        }
    }
}
