// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Inputs
{

    public sealed class VirtualMachineAutoPatchingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of week to apply the patch on.
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<string> DayOfWeek { get; set; } = null!;

        /// <summary>
        /// The size of the Maintenance Window in minutes.
        /// </summary>
        [Input("maintenanceWindowDurationInMinutes", required: true)]
        public Input<int> MaintenanceWindowDurationInMinutes { get; set; } = null!;

        /// <summary>
        /// The Hour, in the Virtual Machine Time-Zone when the patching maintenance window should begin.
        /// </summary>
        [Input("maintenanceWindowStartingHour", required: true)]
        public Input<int> MaintenanceWindowStartingHour { get; set; } = null!;

        public VirtualMachineAutoPatchingArgs()
        {
        }
    }
}