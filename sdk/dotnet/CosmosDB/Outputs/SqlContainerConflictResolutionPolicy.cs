// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Outputs
{

    [OutputType]
    public sealed class SqlContainerConflictResolutionPolicy
    {
        /// <summary>
        /// The conflict resolution path in the case of `LastWriterWins` mode.
        /// </summary>
        public readonly string? ConflictResolutionPath;
        /// <summary>
        /// The procedure to resolve conflicts in the case of `Custom` mode.
        /// </summary>
        public readonly string? ConflictResolutionProcedure;
        /// <summary>
        /// Indicates the conflict resolution mode. Possible values include: `LastWriterWins`, `Custom`.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private SqlContainerConflictResolutionPolicy(
            string? conflictResolutionPath,

            string? conflictResolutionProcedure,

            string mode)
        {
            ConflictResolutionPath = conflictResolutionPath;
            ConflictResolutionProcedure = conflictResolutionProcedure;
            Mode = mode;
        }
    }
}
