// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Outputs
{

    [OutputType]
    public sealed class AadDiagnosticSettingLogRetentionPolicy
    {
        /// <summary>
        /// The number of days for which this Retention Policy should apply. Defaults to `0`.
        /// </summary>
        public readonly int? Days;
        /// <summary>
        /// Is this Retention Policy enabled? Defaults to `false`.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private AadDiagnosticSettingLogRetentionPolicy(
            int? days,

            bool? enabled)
        {
            Days = days;
            Enabled = enabled;
        }
    }
}
