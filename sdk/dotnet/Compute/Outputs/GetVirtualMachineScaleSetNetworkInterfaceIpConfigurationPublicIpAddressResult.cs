// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class GetVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressResult
    {
        /// <summary>
        /// The domain name label for the dns settings.
        /// </summary>
        public readonly string DomainNameLabel;
        public readonly int IdleTimeoutInMinutes;
        public readonly ImmutableArray<Outputs.GetVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagResult> IpTags;
        /// <summary>
        /// The name of this Virtual Machine Scale Set.
        /// </summary>
        public readonly string Name;
        public readonly string PublicIpPrefixId;

        [OutputConstructor]
        private GetVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressResult(
            string domainNameLabel,

            int idleTimeoutInMinutes,

            ImmutableArray<Outputs.GetVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagResult> ipTags,

            string name,

            string publicIpPrefixId)
        {
            DomainNameLabel = domainNameLabel;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpTags = ipTags;
            Name = name;
            PublicIpPrefixId = publicIpPrefixId;
        }
    }
}
