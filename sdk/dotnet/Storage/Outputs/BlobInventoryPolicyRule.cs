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
    public sealed class BlobInventoryPolicyRule
    {
        /// <summary>
        /// A `filter` block as defined above.
        /// </summary>
        public readonly Outputs.BlobInventoryPolicyRuleFilter Filter;
        /// <summary>
        /// The name which should be used for this Blob Inventory Policy Rule.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private BlobInventoryPolicyRule(
            Outputs.BlobInventoryPolicyRuleFilter filter,

            string name)
        {
            Filter = filter;
            Name = name;
        }
    }
}
