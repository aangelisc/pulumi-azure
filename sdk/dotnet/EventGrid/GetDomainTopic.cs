// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid
{
    public static class GetDomainTopic
    {
        /// <summary>
        /// Use this data source to access information about an existing EventGrid Domain Topic
        /// </summary>
        public static Task<GetDomainTopicResult> InvokeAsync(GetDomainTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainTopicResult>("azure:eventgrid/getDomainTopic:getDomainTopic", args ?? new GetDomainTopicArgs(), options.WithVersion());
    }


    public sealed class GetDomainTopicArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EventGrid Domain Topic domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the EventGrid Domain Topic resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain Topic exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDomainTopicArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainTopicResult
    {
        /// <summary>
        /// The EventGrid Domain Topic Domain name.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetDomainTopicResult(
            string domainName,

            string id,

            string name,

            string resourceGroupName)
        {
            DomainName = domainName;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
        }
    }
}