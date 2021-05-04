// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage.Outputs
{

    [OutputType]
    public sealed class GetPolicyRuleActionSnapshotResult
    {
        /// <summary>
        /// The age in days after creation to tier blob version to archive storage.
        /// </summary>
        public readonly int ChangeTierToArchiveAfterDaysSinceCreation;
        /// <summary>
        /// The age in days after creation to tier blob version to cool storage.
        /// </summary>
        public readonly int ChangeTierToCoolAfterDaysSinceCreation;
        /// <summary>
        /// The age in days after creation to delete the blob snapshot.
        /// </summary>
        public readonly int DeleteAfterDaysSinceCreationGreaterThan;

        [OutputConstructor]
        private GetPolicyRuleActionSnapshotResult(
            int changeTierToArchiveAfterDaysSinceCreation,

            int changeTierToCoolAfterDaysSinceCreation,

            int deleteAfterDaysSinceCreationGreaterThan)
        {
            ChangeTierToArchiveAfterDaysSinceCreation = changeTierToArchiveAfterDaysSinceCreation;
            ChangeTierToCoolAfterDaysSinceCreation = changeTierToCoolAfterDaysSinceCreation;
            DeleteAfterDaysSinceCreationGreaterThan = deleteAfterDaysSinceCreationGreaterThan;
        }
    }
}
