// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Consumption.Inputs
{

    public sealed class BudgetResourceGroupTimePeriodGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end date for the budget. If not set this will be 10 years after the start date.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Resource Group Consumption Budget to be created.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public BudgetResourceGroupTimePeriodGetArgs()
        {
        }
    }
}
