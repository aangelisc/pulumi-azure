// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup.Inputs
{

    public sealed class PolicyFileShareRetentionMonthlyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of yearly backups to keep. Must be between `1` and `10`
        /// </summary>
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("weekdays", required: true)]
        private InputList<string>? _weekdays;

        /// <summary>
        /// The weekday backups to retain . Must be one of `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday` or `Saturday`.
        /// </summary>
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        [Input("weeks", required: true)]
        private InputList<string>? _weeks;

        /// <summary>
        /// The weeks of the month to retain backups of. Must be one of `First`, `Second`, `Third`, `Fourth`, `Last`.
        /// </summary>
        public InputList<string> Weeks
        {
            get => _weeks ?? (_weeks = new InputList<string>());
            set => _weeks = value;
        }

        public PolicyFileShareRetentionMonthlyGetArgs()
        {
        }
    }
}
