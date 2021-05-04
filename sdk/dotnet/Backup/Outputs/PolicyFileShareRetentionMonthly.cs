// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup.Outputs
{

    [OutputType]
    public sealed class PolicyFileShareRetentionMonthly
    {
        /// <summary>
        /// The number of yearly backups to keep. Must be between `1` and `10`
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        /// </summary>
        public readonly ImmutableArray<string> Weekdays;
        /// <summary>
        /// The weeks of the month to retain backups of. Must be one of `First`, `Second`, `Third`, `Fourth`, `Last`.
        /// </summary>
        public readonly ImmutableArray<string> Weeks;

        [OutputConstructor]
        private PolicyFileShareRetentionMonthly(
            int count,

            ImmutableArray<string> weekdays,

            ImmutableArray<string> weeks)
        {
            Count = count;
            Weekdays = weekdays;
            Weeks = weeks;
        }
    }
}
