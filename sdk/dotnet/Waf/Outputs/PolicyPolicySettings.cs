// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Waf.Outputs
{

    [OutputType]
    public sealed class PolicyPolicySettings
    {
        /// <summary>
        /// Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private PolicyPolicySettings(
            bool? enabled,

            string? mode)
        {
            Enabled = enabled;
            Mode = mode;
        }
    }
}