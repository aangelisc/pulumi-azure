// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class HBaseClusterMetastores
    {
        /// <summary>
        /// An `ambari` block as defined below.
        /// </summary>
        public readonly Outputs.HBaseClusterMetastoresAmbari? Ambari;
        /// <summary>
        /// A `hive` block as defined below.
        /// </summary>
        public readonly Outputs.HBaseClusterMetastoresHive? Hive;
        /// <summary>
        /// An `oozie` block as defined below.
        /// </summary>
        public readonly Outputs.HBaseClusterMetastoresOozie? Oozie;

        [OutputConstructor]
        private HBaseClusterMetastores(
            Outputs.HBaseClusterMetastoresAmbari? ambari,

            Outputs.HBaseClusterMetastoresHive? hive,

            Outputs.HBaseClusterMetastoresOozie? oozie)
        {
            Ambari = ambari;
            Hive = hive;
            Oozie = oozie;
        }
    }
}