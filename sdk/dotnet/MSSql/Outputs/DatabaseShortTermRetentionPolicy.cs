// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Outputs
{

    [OutputType]
    public sealed class DatabaseShortTermRetentionPolicy
    {
        /// <summary>
        /// Point In Time Restore configuration. Value has to be between `7` and `35`.
        /// </summary>
        public readonly int RetentionDays;

        [OutputConstructor]
        private DatabaseShortTermRetentionPolicy(int retentionDays)
        {
            RetentionDays = retentionDays;
        }
    }
}
