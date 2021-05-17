// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Inputs
{

    public sealed class AccountBackupGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The interval in minutes between two backups. This is configurable only when `type` is `Periodic`. Possible values are between 60 and 1440.
        /// </summary>
        [Input("intervalInMinutes")]
        public Input<int>? IntervalInMinutes { get; set; }

        /// <summary>
        /// The time in hours that each backup is retained. This is configurable only when `type` is `Periodic`. Possible values are between 8 and 720.
        /// </summary>
        [Input("retentionInHours")]
        public Input<int>? RetentionInHours { get; set; }

        /// <summary>
        /// The type of the `backup`. Possible values are `Continuous` and `Periodic`. Defaults to `Periodic`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AccountBackupGetArgs()
        {
        }
    }
}
