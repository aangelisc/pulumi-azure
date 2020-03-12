// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Event Hubs Authorization Rule within an Event Hub.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/eventhub_authorization_rule.html.markdown.
        /// </summary>
        public static Task<GetAuthorizationRuleResult> GetAuthorizationRule(GetAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationRuleResult>("azure:eventhub/getAuthorizationRule:getAuthorizationRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub.
        /// </summary>
        [Input("eventhubName", required: true)]
        public string EventhubName { get; set; } = null!;

        [Input("listen")]
        public bool? Listen { get; set; }

        [Input("manage")]
        public bool? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Authorization Rule resource. be created.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Authorization Rule's grandparent Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("send")]
        public bool? Send { get; set; }

        public GetAuthorizationRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAuthorizationRuleResult
    {
        public readonly string EventhubName;
        public readonly bool? Listen;
        public readonly string Location;
        public readonly bool? Manage;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// The Primary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The Primary Key for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The Secondary Key for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string SecondaryKey;
        public readonly bool? Send;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAuthorizationRuleResult(
            string eventhubName,
            bool? listen,
            string location,
            bool? manage,
            string name,
            string namespaceName,
            string primaryConnectionString,
            string primaryKey,
            string resourceGroupName,
            string secondaryConnectionString,
            string secondaryKey,
            bool? send,
            string id)
        {
            EventhubName = eventhubName;
            Listen = listen;
            Location = location;
            Manage = manage;
            Name = name;
            NamespaceName = namespaceName;
            PrimaryConnectionString = primaryConnectionString;
            PrimaryKey = primaryKey;
            ResourceGroupName = resourceGroupName;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryKey = secondaryKey;
            Send = send;
            Id = id;
        }
    }
}
