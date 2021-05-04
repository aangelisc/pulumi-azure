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
    public sealed class GetVirtualMachineScaleSetNetworkInterfaceResult
    {
        /// <summary>
        /// The dns servers in use.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        public readonly bool EnableAcceleratedNetworking;
        public readonly bool EnableIpForwarding;
        /// <summary>
        /// An ip_configuration block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineScaleSetNetworkInterfaceIpConfigurationResult> IpConfigurations;
        /// <summary>
        /// The name of this Virtual Machine Scale Set.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The identifier for the network security group.
        /// </summary>
        public readonly string NetworkSecurityGroupId;
        /// <summary>
        /// If this ip_configuration is the primary one.
        /// </summary>
        public readonly bool Primary;

        [OutputConstructor]
        private GetVirtualMachineScaleSetNetworkInterfaceResult(
            ImmutableArray<string> dnsServers,

            bool enableAcceleratedNetworking,

            bool enableIpForwarding,

            ImmutableArray<Outputs.GetVirtualMachineScaleSetNetworkInterfaceIpConfigurationResult> ipConfigurations,

            string name,

            string networkSecurityGroupId,

            bool primary)
        {
            DnsServers = dnsServers;
            EnableAcceleratedNetworking = enableAcceleratedNetworking;
            EnableIpForwarding = enableIpForwarding;
            IpConfigurations = ipConfigurations;
            Name = name;
            NetworkSecurityGroupId = networkSecurityGroupId;
            Primary = primary;
        }
    }
}
