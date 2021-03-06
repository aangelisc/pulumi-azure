// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataboxEdge.Outputs
{

    [OutputType]
    public sealed class OrderStatus
    {
        /// <summary>
        /// Dictionary to hold generic information which is not stored by the already existing properties.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AdditionalDetails;
        /// <summary>
        /// Comments related to this status change.
        /// </summary>
        public readonly string? Comments;
        /// <summary>
        /// The current status of the order. Possible values include `Untracked`, `AwaitingFulfilment`, `AwaitingPreparation`, `AwaitingShipment`, `Shipped`, `Arriving`, `Delivered`, `ReplacementRequested`, `LostDevice`, `Declined`, `ReturnInitiated`, `AwaitingReturnShipment`, `ShippedBack` or `CollectedAtMicrosoft`.
        /// </summary>
        public readonly string? Info;
        /// <summary>
        /// Time of status update.
        /// </summary>
        public readonly string? LastUpdate;

        [OutputConstructor]
        private OrderStatus(
            ImmutableDictionary<string, string>? additionalDetails,

            string? comments,

            string? info,

            string? lastUpdate)
        {
            AdditionalDetails = additionalDetails;
            Comments = comments;
            Info = info;
            LastUpdate = lastUpdate;
        }
    }
}
