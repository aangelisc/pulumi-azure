// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing EventGrid Topic
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/eventgrid_topic.html.markdown.
        /// </summary>
        public static Task<GetTopicResult> GetTopic(GetTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicResult>("azure:eventgrid/getTopic:getTopic", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetTopicArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EventGrid Topic resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid Topic exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetTopicArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTopicResult
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Topic.
        /// </summary>
        public readonly string Endpoint;
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Topic.
        /// </summary>
        public readonly string PrimaryAccessKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Topic.
        /// </summary>
        public readonly string SecondaryAccessKey;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTopicResult(
            string endpoint,
            string location,
            string name,
            string primaryAccessKey,
            string resourceGroupName,
            string secondaryAccessKey,
            ImmutableDictionary<string, string>? tags,
            string id)
        {
            Endpoint = endpoint;
            Location = location;
            Name = name;
            PrimaryAccessKey = primaryAccessKey;
            ResourceGroupName = resourceGroupName;
            SecondaryAccessKey = secondaryAccessKey;
            Tags = tags;
            Id = id;
        }
    }
}
