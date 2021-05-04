// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp
{
    public static class GetVolume
    {
        /// <summary>
        /// Uses this data source to access information about an existing NetApp Volume.
        /// 
        /// ## NetApp Volume Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.NetApp.GetVolume.InvokeAsync(new Azure.NetApp.GetVolumeArgs
        ///         {
        ///             ResourceGroupName = "acctestRG",
        ///             AccountName = "acctestnetappaccount",
        ///             PoolName = "acctestnetapppool",
        ///             Name = "example-volume",
        ///         }));
        ///         this.NetappVolumeId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("netappVolumeId")]
        ///     public Output&lt;string&gt; NetappVolumeId { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("azure:netapp/getVolume:getVolume", args ?? new GetVolumeArgs(), options.WithVersion());
    }


    public sealed class GetVolumeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account where the NetApp pool exists.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Volume.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp pool where the NetApp volume exists.
        /// </summary>
        [Input("poolName", required: true)]
        public string PoolName { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Volume exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVolumeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        public readonly string AccountName;
        /// <summary>
        /// A `data_protection_replication` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVolumeDataProtectionReplicationResult> DataProtectionReplications;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the NetApp Volume exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// A list of IPv4 Addresses which should be used to mount the volume.
        /// </summary>
        public readonly ImmutableArray<string> MountIpAddresses;
        public readonly string Name;
        public readonly string PoolName;
        /// <summary>
        /// A list of protocol types.
        /// </summary>
        public readonly ImmutableArray<string> Protocols;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The service level of the file system.
        /// </summary>
        public readonly string ServiceLevel;
        /// <summary>
        /// The maximum Storage Quota in Gigabytes allowed for a file system.
        /// </summary>
        public readonly int StorageQuotaInGb;
        /// <summary>
        /// The ID of a Subnet in which the NetApp Volume resides.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The unique file path of the volume.
        /// </summary>
        public readonly string VolumePath;

        [OutputConstructor]
        private GetVolumeResult(
            string accountName,

            ImmutableArray<Outputs.GetVolumeDataProtectionReplicationResult> dataProtectionReplications,

            string id,

            string location,

            ImmutableArray<string> mountIpAddresses,

            string name,

            string poolName,

            ImmutableArray<string> protocols,

            string resourceGroupName,

            string serviceLevel,

            int storageQuotaInGb,

            string subnetId,

            string volumePath)
        {
            AccountName = accountName;
            DataProtectionReplications = dataProtectionReplications;
            Id = id;
            Location = location;
            MountIpAddresses = mountIpAddresses;
            Name = name;
            PoolName = poolName;
            Protocols = protocols;
            ResourceGroupName = resourceGroupName;
            ServiceLevel = serviceLevel;
            StorageQuotaInGb = storageQuotaInGb;
            SubnetId = subnetId;
            VolumePath = volumePath;
        }
    }
}
