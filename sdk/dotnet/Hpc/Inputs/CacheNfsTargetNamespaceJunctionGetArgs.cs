// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hpc.Inputs
{

    public sealed class CacheNfsTargetNamespaceJunctionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the access policy applied to this target. Defaults to `default`.
        /// </summary>
        [Input("accessPolicyName")]
        public Input<string>? AccessPolicyName { get; set; }

        /// <summary>
        /// The client-facing file path of this NFS target within the HPC Cache NFS Target.
        /// </summary>
        [Input("namespacePath", required: true)]
        public Input<string> NamespacePath { get; set; } = null!;

        /// <summary>
        /// The NFS export of this NFS target within the HPC Cache NFS Target.
        /// </summary>
        [Input("nfsExport", required: true)]
        public Input<string> NfsExport { get; set; } = null!;

        /// <summary>
        /// The relative subdirectory path from the `nfs_export` to map to the `namespace_path`. Defaults to `""`, in which case the whole `nfs_export` is exported.
        /// </summary>
        [Input("targetPath")]
        public Input<string>? TargetPath { get; set; }

        public CacheNfsTargetNamespaceJunctionGetArgs()
        {
        }
    }
}
