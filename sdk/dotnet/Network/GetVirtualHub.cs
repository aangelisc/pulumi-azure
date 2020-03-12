// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static partial class Invokes
    {
        /// <summary>
        /// Uses this data source to access information about an existing Virtual Hub.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/virtual_hub.html.markdown.
        /// </summary>
        public static Task<GetVirtualHubResult> GetVirtualHub(GetVirtualHubArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualHubResult>("azure:network/getVirtualHub:getVirtualHub", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVirtualHubArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Virtual Hub.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the Virtual Hub exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualHubArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVirtualHubResult
    {
        /// <summary>
        /// The Address Prefix used for this Virtual Hub.
        /// </summary>
        public readonly string AddressPrefix;
        /// <summary>
        /// The Azure Region where the Virtual Hub exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the Virtual Hub.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The ID of the Virtual WAN within which the Virtual Hub exists.
        /// </summary>
        public readonly string VirtualWanId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVirtualHubResult(
            string addressPrefix,
            string location,
            string name,
            string resourceGroupName,
            ImmutableDictionary<string, string> tags,
            string virtualWanId,
            string id)
        {
            AddressPrefix = addressPrefix;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            VirtualWanId = virtualWanId;
            Id = id;
        }
    }
}
