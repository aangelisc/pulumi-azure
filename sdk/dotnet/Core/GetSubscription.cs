// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Subscription.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subscription.html.markdown.
        /// </summary>
        public static Task<GetSubscriptionResult> GetSubscription(GetSubscriptionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionResult>("azure:core/getSubscription:getSubscription", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSubscriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the ID of the subscription. If this argument is omitted, the subscription ID of the current Azure Resource Manager provider is used.
        /// </summary>
        [Input("subscriptionId")]
        public string? SubscriptionId { get; set; }

        public GetSubscriptionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSubscriptionResult
    {
        /// <summary>
        /// The subscription display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The subscription location placement ID.
        /// </summary>
        public readonly string LocationPlacementId;
        /// <summary>
        /// The subscription quota ID.
        /// </summary>
        public readonly string QuotaId;
        /// <summary>
        /// The subscription spending limit.
        /// </summary>
        public readonly string SpendingLimit;
        /// <summary>
        /// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The subscription GUID.
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// The subscription tenant ID.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSubscriptionResult(
            string displayName,
            string locationPlacementId,
            string quotaId,
            string spendingLimit,
            string state,
            string subscriptionId,
            string tenantId,
            string id)
        {
            DisplayName = displayName;
            LocationPlacementId = locationPlacementId;
            QuotaId = quotaId;
            SpendingLimit = spendingLimit;
            State = state;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
            Id = id;
        }
    }
}
