// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Consumption.Outputs
{

    [OutputType]
    public sealed class BudgetResourceGroupFilterTag
    {
        /// <summary>
        /// The name of the tag to use for the filter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The operator to use for comparison. The allowed values are `In`.
        /// </summary>
        public readonly string? Operator;
        /// <summary>
        /// Specifies a list of values for the tag.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private BudgetResourceGroupFilterTag(
            string name,

            string? @operator,

            ImmutableArray<string> values)
        {
            Name = name;
            Operator = @operator;
            Values = values;
        }
    }
}
