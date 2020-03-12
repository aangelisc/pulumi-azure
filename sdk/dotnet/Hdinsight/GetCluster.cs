// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing HDInsight Cluster.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/hdinsight_cluster.html.markdown.
        /// </summary>
        public static Task<GetClusterResult> GetCluster(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azure:hdinsight/getCluster:getCluster", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of this HDInsight Cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Cluster exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The version of HDInsights which is used on this HDInsight Cluster.
        /// </summary>
        public readonly string ClusterVersion;
        /// <summary>
        /// A map of versions of software used on this HDInsights Cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ComponentVersions;
        /// <summary>
        /// The SSH Endpoint of the Edge Node for this HDInsight Cluster, if an Edge Node exists.
        /// </summary>
        public readonly string EdgeSshEndpoint;
        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterGatewaysResult> Gateways;
        /// <summary>
        /// The HTTPS Endpoint for this HDInsight Cluster.
        /// </summary>
        public readonly string HttpsEndpoint;
        /// <summary>
        /// The kind of HDInsight Cluster this is, such as a Spark or Storm cluster.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The Azure Region in which this HDInsight Cluster exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SSH Endpoint for this HDInsight Cluster.
        /// </summary>
        public readonly string SshEndpoint;
        /// <summary>
        /// A map of tags assigned to the HDInsight Cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The SKU / Tier of this HDInsight Cluster.
        /// </summary>
        public readonly string Tier;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterResult(
            string clusterVersion,
            ImmutableDictionary<string, string> componentVersions,
            string edgeSshEndpoint,
            ImmutableArray<Outputs.GetClusterGatewaysResult> gateways,
            string httpsEndpoint,
            string kind,
            string location,
            string name,
            string resourceGroupName,
            string sshEndpoint,
            ImmutableDictionary<string, string> tags,
            string tier,
            string id)
        {
            ClusterVersion = clusterVersion;
            ComponentVersions = componentVersions;
            EdgeSshEndpoint = edgeSshEndpoint;
            Gateways = gateways;
            HttpsEndpoint = httpsEndpoint;
            Kind = kind;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SshEndpoint = sshEndpoint;
            Tags = tags;
            Tier = tier;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClusterGatewaysResult
    {
        /// <summary>
        /// Is the Ambari Portal enabled?
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The password used for the Ambari Portal.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The username used for the Ambari Portal.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetClusterGatewaysResult(
            bool enabled,
            string password,
            string username)
        {
            Enabled = enabled;
            Password = password;
            Username = username;
        }
    }
    }
}
