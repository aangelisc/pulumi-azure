// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class MLServicesClusterRolesEdgeNode
    {
        /// <summary>
        /// The Password associated with the local administrator for the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// A list of SSH Keys which should be used for the local administrator on the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> SshKeys;
        /// <summary>
        /// The ID of the Subnet within the Virtual Network where the Edge Node should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The Username of the local administrator for the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The ID of the Virtual Network where the Edge Node should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? VirtualNetworkId;
        /// <summary>
        /// The Size of the Virtual Machine which should be used as the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string VmSize;

        [OutputConstructor]
        private MLServicesClusterRolesEdgeNode(
            string? password,

            ImmutableArray<string> sshKeys,

            string? subnetId,

            string username,

            string? virtualNetworkId,

            string vmSize)
        {
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }
}
