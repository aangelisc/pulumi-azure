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
    public sealed class SlotSiteConfigIpRestrictionHeaders
    {
        /// <summary>
        /// A list of allowed Azure FrontDoor IDs in UUID notation with a maximum of 8.
        /// </summary>
        public readonly ImmutableArray<string> XAzureFdids;
        /// <summary>
        /// A list to allow the Azure FrontDoor health probe header. Only allowed value is "1".
        /// </summary>
        public readonly string? XFdHealthProbe;
        /// <summary>
        /// A list of allowed 'X-Forwarded-For' IPs in CIDR notation with a maximum of 8
        /// </summary>
        public readonly ImmutableArray<string> XForwardedFors;
        /// <summary>
        /// A list of allowed 'X-Forwarded-Host' domains with a maximum of 8.
        /// </summary>
        public readonly ImmutableArray<string> XForwardedHosts;

        [OutputConstructor]
        private SlotSiteConfigIpRestrictionHeaders(
            ImmutableArray<string> xAzureFdids,

            string? xFdHealthProbe,

            ImmutableArray<string> xForwardedFors,

            ImmutableArray<string> xForwardedHosts)
        {
            XAzureFdids = xAzureFdids;
            XFdHealthProbe = xFdHealthProbe;
            XForwardedFors = xForwardedFors;
            XForwardedHosts = xForwardedHosts;
        }
    }
}
