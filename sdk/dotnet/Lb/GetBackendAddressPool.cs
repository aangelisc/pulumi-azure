// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Load Balancer's Backend Address Pool.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/lb_backend_address_pool.html.markdown.
        /// </summary>
        public static Task<GetBackendAddressPoolResult> GetBackendAddressPool(GetBackendAddressPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackendAddressPoolResult>("azure:lb/getBackendAddressPool:getBackendAddressPool", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetBackendAddressPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Load Balancer in which the Backend Address Pool exists.
        /// </summary>
        [Input("loadbalancerId", required: true)]
        public string LoadbalancerId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Backend Address Pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBackendAddressPoolArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetBackendAddressPoolResult
    {
        /// <summary>
        /// An array of references to IP addresses defined in network interfaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackendAddressPoolBackendIpConfigurationsResult> BackendIpConfigurations;
        public readonly string LoadbalancerId;
        /// <summary>
        /// The name of the Backend Address Pool.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetBackendAddressPoolResult(
            ImmutableArray<Outputs.GetBackendAddressPoolBackendIpConfigurationsResult> backendIpConfigurations,
            string loadbalancerId,
            string name,
            string id)
        {
            BackendIpConfigurations = backendIpConfigurations;
            LoadbalancerId = loadbalancerId;
            Name = name;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetBackendAddressPoolBackendIpConfigurationsResult
    {
        /// <summary>
        /// The ID of the Backend Address Pool.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetBackendAddressPoolBackendIpConfigurationsResult(string id)
        {
            Id = id;
        }
    }
    }
}
