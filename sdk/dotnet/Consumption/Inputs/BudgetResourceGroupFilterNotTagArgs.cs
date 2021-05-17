// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Consumption.Inputs
{

    public sealed class BudgetResourceGroupFilterNotTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the tag to use for the filter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The operator to use for comparison. The allowed values are `In`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Specifies a list of values for the tag.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public BudgetResourceGroupFilterNotTagArgs()
        {
        }
    }
}
