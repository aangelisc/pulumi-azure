// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Inputs
{

    public sealed class VirtualMachineStorageConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `storage_settings` as defined below.
        /// </summary>
        [Input("dataSettings")]
        public Input<Inputs.VirtualMachineStorageConfigurationDataSettingsArgs>? DataSettings { get; set; }

        /// <summary>
        /// The type of disk configuration to apply to the SQL Server. Valid values include `NEW`, `EXTEND`, or `ADD`.
        /// </summary>
        [Input("diskType", required: true)]
        public Input<string> DiskType { get; set; } = null!;

        /// <summary>
        /// An `storage_settings` as defined below.
        /// </summary>
        [Input("logSettings")]
        public Input<Inputs.VirtualMachineStorageConfigurationLogSettingsArgs>? LogSettings { get; set; }

        /// <summary>
        /// The type of storage workload. Valid values include `GENERAL`, `OLTP`, or `DW`.
        /// </summary>
        [Input("storageWorkloadType", required: true)]
        public Input<string> StorageWorkloadType { get; set; } = null!;

        /// <summary>
        /// An `storage_settings` as defined below.
        /// </summary>
        [Input("tempDbSettings")]
        public Input<Inputs.VirtualMachineStorageConfigurationTempDbSettingsArgs>? TempDbSettings { get; set; }

        public VirtualMachineStorageConfigurationArgs()
        {
        }
    }
}
