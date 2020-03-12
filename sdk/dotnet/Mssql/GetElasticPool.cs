// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing SQL elastic pool.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/mssql_elasticpool.html.markdown.
        /// </summary>
        public static Task<GetElasticPoolResult> GetElasticPool(GetElasticPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetElasticPoolResult>("azure:mssql/getElasticPool:getElasticPool", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetElasticPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the elastic pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group which contains the elastic pool.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SQL Server which contains the elastic pool.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetElasticPoolArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetElasticPoolResult
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The max data size of the elastic pool in bytes.
        /// </summary>
        public readonly int MaxSizeBytes;
        /// <summary>
        /// The max data size of the elastic pool in gigabytes.
        /// </summary>
        public readonly double MaxSizeGb;
        public readonly string Name;
        /// <summary>
        /// The maximum capacity any one database can consume.
        /// </summary>
        public readonly int PerDbMaxCapacity;
        /// <summary>
        /// The minimum capacity all databases are guaranteed.
        /// </summary>
        public readonly int PerDbMinCapacity;
        public readonly string ResourceGroupName;
        public readonly string ServerName;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Whether or not this elastic pool is zone redundant.
        /// </summary>
        public readonly bool ZoneRedundant;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetElasticPoolResult(
            string location,
            int maxSizeBytes,
            double maxSizeGb,
            string name,
            int perDbMaxCapacity,
            int perDbMinCapacity,
            string resourceGroupName,
            string serverName,
            ImmutableDictionary<string, string> tags,
            bool zoneRedundant,
            string id)
        {
            Location = location;
            MaxSizeBytes = maxSizeBytes;
            MaxSizeGb = maxSizeGb;
            Name = name;
            PerDbMaxCapacity = perDbMaxCapacity;
            PerDbMinCapacity = perDbMinCapacity;
            ResourceGroupName = resourceGroupName;
            ServerName = serverName;
            Tags = tags;
            ZoneRedundant = zoneRedundant;
            Id = id;
        }
    }
}
