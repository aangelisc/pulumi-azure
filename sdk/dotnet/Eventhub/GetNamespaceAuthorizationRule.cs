// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an Authorization Rule for an Event Hub Namespace.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/eventhub_namespace_authorization_rule.html.markdown.
        /// </summary>
        public static Task<GetNamespaceAuthorizationRuleResult> GetNamespaceAuthorizationRule(GetNamespaceAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceAuthorizationRuleResult>("azure:eventhub/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetNamespaceAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EventHub Authorization Rule resource. 
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceAuthorizationRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetNamespaceAuthorizationRuleResult
    {
        /// <summary>
        /// Does this Authorization Rule have permissions to Listen to the Event Hub?
        /// </summary>
        public readonly bool Listen;
        /// <summary>
        /// Does this Authorization Rule have permissions to Manage to the Event Hub?
        /// </summary>
        public readonly bool Manage;
        public readonly string Name;
        /// <summary>
        /// The name of the EventHub Namespace. 
        /// </summary>
        public readonly string NamespaceName;
        /// <summary>
        /// The Primary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The Primary Key for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Connection String for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The Secondary Key for the Event Hubs authorization Rule.
        /// </summary>
        public readonly string SecondaryKey;
        /// <summary>
        /// Does this Authorization Rule have permissions to Send to the Event Hub?
        /// </summary>
        public readonly bool Send;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetNamespaceAuthorizationRuleResult(
            bool listen,
            bool manage,
            string name,
            string namespaceName,
            string primaryConnectionString,
            string primaryKey,
            string resourceGroupName,
            string secondaryConnectionString,
            string secondaryKey,
            bool send,
            string id)
        {
            Listen = listen;
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
