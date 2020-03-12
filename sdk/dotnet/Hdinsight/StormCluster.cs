// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight
{
    /// <summary>
    /// Manages a HDInsight Storm Cluster.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/hdinsight_storm_cluster.html.markdown.
    /// </summary>
    public partial class StormCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterVersion")]
        public Output<string> ClusterVersion { get; private set; } = null!;

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Output("componentVersion")]
        public Output<Outputs.StormClusterComponentVersion> ComponentVersion { get; private set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.StormClusterGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Storm Cluster.
        /// </summary>
        [Output("httpsEndpoint")]
        public Output<string> HttpsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.StormClusterRoles> Roles { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Storm Cluster.
        /// </summary>
        [Output("sshEndpoint")]
        public Output<string> SshEndpoint { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.StormClusterStorageAccounts>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Storm Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a StormCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StormCluster(string name, StormClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:hdinsight/stormCluster:StormCluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private StormCluster(string name, Input<string> id, StormClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:hdinsight/stormCluster:StormCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StormCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StormCluster Get(string name, Input<string> id, StormClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new StormCluster(name, id, state, options);
        }
    }

    public sealed class StormClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion", required: true)]
        public Input<string> ClusterVersion { get; set; } = null!;

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Input("componentVersion", required: true)]
        public Input<Inputs.StormClusterComponentVersionArgs> ComponentVersion { get; set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<Inputs.StormClusterGatewayArgs> Gateway { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.StormClusterRolesArgs> Roles { get; set; } = null!;

        [Input("storageAccounts")]
        private InputList<Inputs.StormClusterStorageAccountsArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.StormClusterStorageAccountsArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.StormClusterStorageAccountsArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Storm Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public StormClusterArgs()
        {
        }
    }

    public sealed class StormClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Input("componentVersion")]
        public Input<Inputs.StormClusterComponentVersionGetArgs>? ComponentVersion { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.StormClusterGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Storm Cluster.
        /// </summary>
        [Input("httpsEndpoint")]
        public Input<string>? HttpsEndpoint { get; set; }

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.StormClusterRolesGetArgs>? Roles { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Storm Cluster.
        /// </summary>
        [Input("sshEndpoint")]
        public Input<string>? SshEndpoint { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.StormClusterStorageAccountsGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.StormClusterStorageAccountsGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.StormClusterStorageAccountsGetArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Storm Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public StormClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class StormClusterComponentVersionArgs : Pulumi.ResourceArgs
    {
        [Input("storm", required: true)]
        public Input<string> Storm { get; set; } = null!;

        public StormClusterComponentVersionArgs()
        {
        }
    }

    public sealed class StormClusterComponentVersionGetArgs : Pulumi.ResourceArgs
    {
        [Input("storm", required: true)]
        public Input<string> Storm { get; set; } = null!;

        public StormClusterComponentVersionGetArgs()
        {
        }
    }

    public sealed class StormClusterGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public StormClusterGatewayArgs()
        {
        }
    }

    public sealed class StormClusterGatewayGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public StormClusterGatewayGetArgs()
        {
        }
    }

    public sealed class StormClusterRolesArgs : Pulumi.ResourceArgs
    {
        [Input("headNode", required: true)]
        public Input<StormClusterRolesHeadNodeArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<StormClusterRolesWorkerNodeArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<StormClusterRolesZookeeperNodeArgs> ZookeeperNode { get; set; } = null!;

        public StormClusterRolesArgs()
        {
        }
    }

    public sealed class StormClusterRolesGetArgs : Pulumi.ResourceArgs
    {
        [Input("headNode", required: true)]
        public Input<StormClusterRolesHeadNodeGetArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<StormClusterRolesWorkerNodeGetArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<StormClusterRolesZookeeperNodeGetArgs> ZookeeperNode { get; set; } = null!;

        public StormClusterRolesGetArgs()
        {
        }
    }

    public sealed class StormClusterRolesHeadNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesHeadNodeArgs()
        {
        }
    }

    public sealed class StormClusterRolesHeadNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesHeadNodeGetArgs()
        {
        }
    }

    public sealed class StormClusterRolesWorkerNodeArgs : Pulumi.ResourceArgs
    {
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesWorkerNodeArgs()
        {
        }
    }

    public sealed class StormClusterRolesWorkerNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesWorkerNodeGetArgs()
        {
        }
    }

    public sealed class StormClusterRolesZookeeperNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesZookeeperNodeArgs()
        {
        }
    }

    public sealed class StormClusterRolesZookeeperNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesZookeeperNodeGetArgs()
        {
        }
    }

    public sealed class StormClusterStorageAccountsArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public StormClusterStorageAccountsArgs()
        {
        }
    }

    public sealed class StormClusterStorageAccountsGetArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public StormClusterStorageAccountsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class StormClusterComponentVersion
    {
        public readonly string Storm;

        [OutputConstructor]
        private StormClusterComponentVersion(string storm)
        {
            Storm = storm;
        }
    }

    [OutputType]
    public sealed class StormClusterGateway
    {
        public readonly bool Enabled;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private StormClusterGateway(
            bool enabled,
            string password,
            string username)
        {
            Enabled = enabled;
            Password = password;
            Username = username;
        }
    }

    [OutputType]
    public sealed class StormClusterRoles
    {
        public readonly StormClusterRolesHeadNode HeadNode;
        public readonly StormClusterRolesWorkerNode WorkerNode;
        public readonly StormClusterRolesZookeeperNode ZookeeperNode;

        [OutputConstructor]
        private StormClusterRoles(
            StormClusterRolesHeadNode headNode,
            StormClusterRolesWorkerNode workerNode,
            StormClusterRolesZookeeperNode zookeeperNode)
        {
            HeadNode = headNode;
            WorkerNode = workerNode;
            ZookeeperNode = zookeeperNode;
        }
    }

    [OutputType]
    public sealed class StormClusterRolesHeadNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private StormClusterRolesHeadNode(
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

    [OutputType]
    public sealed class StormClusterRolesWorkerNode
    {
        public readonly int? MinInstanceCount;
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly int TargetInstanceCount;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private StormClusterRolesWorkerNode(
            int? minInstanceCount,
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            int targetInstanceCount,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            MinInstanceCount = minInstanceCount;
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            TargetInstanceCount = targetInstanceCount;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class StormClusterRolesZookeeperNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private StormClusterRolesZookeeperNode(
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

    [OutputType]
    public sealed class StormClusterStorageAccounts
    {
        public readonly bool IsDefault;
        public readonly string StorageAccountKey;
        public readonly string StorageContainerId;

        [OutputConstructor]
        private StormClusterStorageAccounts(
            bool isDefault,
            string storageAccountKey,
            string storageContainerId)
        {
            IsDefault = isDefault;
            StorageAccountKey = storageAccountKey;
            StorageContainerId = storageContainerId;
        }
    }
    }
}
