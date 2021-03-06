// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MediaServices.Outputs
{

    [OutputType]
    public sealed class AccountKeyDeliveryAccessControl
    {
        /// <summary>
        /// The Default Action to use when no rules match from `ip_allow_list`. Possible values are `Allow` and `Deny`.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// One or more IP Addresses, or CIDR Blocks which should be able to access the Key Delivery.
        /// </summary>
        public readonly ImmutableArray<string> IpAllowLists;

        [OutputConstructor]
        private AccountKeyDeliveryAccessControl(
            string? defaultAction,

            ImmutableArray<string> ipAllowLists)
        {
            DefaultAction = defaultAction;
            IpAllowLists = ipAllowLists;
        }
    }
}
