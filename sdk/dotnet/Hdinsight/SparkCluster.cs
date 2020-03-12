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
    /// Manages a HDInsight Spark Cluster.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/hdinsight_spark_cluster.html.markdown.
    /// </summary>
    public partial class SparkCluster : Pulumi.CustomResource
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
        public Output<Outputs.SparkClusterComponentVersion> ComponentVersion { get; private set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.SparkClusterGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
        /// </summary>
        [Output("httpsEndpoint")]
        public Output<string> HttpsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.SparkClusterRoles> Roles { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
        /// </summary>
        [Output("sshEndpoint")]
        public Output<string> SshEndpoint { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.SparkClusterStorageAccounts>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Output("storageAccountGen2")]
        public Output<Outputs.SparkClusterStorageAccountGen2?> StorageAccountGen2 { get; private set; } = null!;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Spark Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a SparkCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SparkCluster(string name, SparkClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:hdinsight/sparkCluster:SparkCluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SparkCluster(string name, Input<string> id, SparkClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:hdinsight/sparkCluster:SparkCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SparkCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SparkCluster Get(string name, Input<string> id, SparkClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new SparkCluster(name, id, state, options);
        }
    }

    public sealed class SparkClusterArgs : Pulumi.ResourceArgs
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
        public Input<Inputs.SparkClusterComponentVersionArgs> ComponentVersion { get; set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<Inputs.SparkClusterGatewayArgs> Gateway { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.SparkClusterRolesArgs> Roles { get; set; } = null!;

        [Input("storageAccounts")]
        private InputList<Inputs.SparkClusterStorageAccountsArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.SparkClusterStorageAccountsArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.SparkClusterStorageAccountsArgs>());
            set => _storageAccounts = value;
        }

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.SparkClusterStorageAccountGen2Args>? StorageAccountGen2 { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Spark Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public SparkClusterArgs()
        {
        }
    }

    public sealed class SparkClusterState : Pulumi.ResourceArgs
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
        public Input<Inputs.SparkClusterComponentVersionGetArgs>? ComponentVersion { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.SparkClusterGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
        /// </summary>
        [Input("httpsEndpoint")]
        public Input<string>? HttpsEndpoint { get; set; }

        /// <summary>
        /// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.SparkClusterRolesGetArgs>? Roles { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
        /// </summary>
        [Input("sshEndpoint")]
        public Input<string>? SshEndpoint { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.SparkClusterStorageAccountsGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.SparkClusterStorageAccountsGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.SparkClusterStorageAccountsGetArgs>());
            set => _storageAccounts = value;
        }

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.SparkClusterStorageAccountGen2GetArgs>? StorageAccountGen2 { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight Spark Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public SparkClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SparkClusterComponentVersionArgs : Pulumi.ResourceArgs
    {
        [Input("spark", required: true)]
        public Input<string> Spark { get; set; } = null!;

        public SparkClusterComponentVersionArgs()
        {
        }
    }

    public sealed class SparkClusterComponentVersionGetArgs : Pulumi.ResourceArgs
    {
        [Input("spark", required: true)]
        public Input<string> Spark { get; set; } = null!;

        public SparkClusterComponentVersionGetArgs()
        {
        }
    }

    public sealed class SparkClusterGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SparkClusterGatewayArgs()
        {
        }
    }

    public sealed class SparkClusterGatewayGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SparkClusterGatewayGetArgs()
        {
        }
    }

    public sealed class SparkClusterRolesArgs : Pulumi.ResourceArgs
    {
        [Input("headNode", required: true)]
        public Input<SparkClusterRolesHeadNodeArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<SparkClusterRolesWorkerNodeArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<SparkClusterRolesZookeeperNodeArgs> ZookeeperNode { get; set; } = null!;

        public SparkClusterRolesArgs()
        {
        }
    }

    public sealed class SparkClusterRolesGetArgs : Pulumi.ResourceArgs
    {
        [Input("headNode", required: true)]
        public Input<SparkClusterRolesHeadNodeGetArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<SparkClusterRolesWorkerNodeGetArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<SparkClusterRolesZookeeperNodeGetArgs> ZookeeperNode { get; set; } = null!;

        public SparkClusterRolesGetArgs()
        {
        }
    }

    public sealed class SparkClusterRolesHeadNodeArgs : Pulumi.ResourceArgs
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

        public SparkClusterRolesHeadNodeArgs()
        {
        }
    }

    public sealed class SparkClusterRolesHeadNodeGetArgs : Pulumi.ResourceArgs
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

        public SparkClusterRolesHeadNodeGetArgs()
        {
        }
    }

    public sealed class SparkClusterRolesWorkerNodeArgs : Pulumi.ResourceArgs
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

        public SparkClusterRolesWorkerNodeArgs()
        {
        }
    }

    public sealed class SparkClusterRolesWorkerNodeGetArgs : Pulumi.ResourceArgs
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

        public SparkClusterRolesWorkerNodeGetArgs()
        {
        }
    }

    public sealed class SparkClusterRolesZookeeperNodeArgs : Pulumi.ResourceArgs
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

        public SparkClusterRolesZookeeperNodeArgs()
        {
        }
    }

    public sealed class SparkClusterRolesZookeeperNodeGetArgs : Pulumi.ResourceArgs
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

        public SparkClusterRolesZookeeperNodeGetArgs()
        {
        }
    }

    public sealed class SparkClusterStorageAccountGen2Args : Pulumi.ResourceArgs
    {
        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("managedIdentityResourceId", required: true)]
        public Input<string> ManagedIdentityResourceId { get; set; } = null!;

        [Input("storageResourceId", required: true)]
        public Input<string> StorageResourceId { get; set; } = null!;

        public SparkClusterStorageAccountGen2Args()
        {
        }
    }

    public sealed class SparkClusterStorageAccountGen2GetArgs : Pulumi.ResourceArgs
    {
        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("managedIdentityResourceId", required: true)]
        public Input<string> ManagedIdentityResourceId { get; set; } = null!;

        [Input("storageResourceId", required: true)]
        public Input<string> StorageResourceId { get; set; } = null!;

        public SparkClusterStorageAccountGen2GetArgs()
        {
        }
    }

    public sealed class SparkClusterStorageAccountsArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public SparkClusterStorageAccountsArgs()
        {
        }
    }

    public sealed class SparkClusterStorageAccountsGetArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public SparkClusterStorageAccountsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SparkClusterComponentVersion
    {
        public readonly string Spark;

        [OutputConstructor]
        private SparkClusterComponentVersion(string spark)
        {
            Spark = spark;
        }
    }

    [OutputType]
    public sealed class SparkClusterGateway
    {
        public readonly bool Enabled;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private SparkClusterGateway(
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
    public sealed class SparkClusterRoles
    {
        public readonly SparkClusterRolesHeadNode HeadNode;
        public readonly SparkClusterRolesWorkerNode WorkerNode;
        public readonly SparkClusterRolesZookeeperNode ZookeeperNode;

        [OutputConstructor]
        private SparkClusterRoles(
            SparkClusterRolesHeadNode headNode,
            SparkClusterRolesWorkerNode workerNode,
            SparkClusterRolesZookeeperNode zookeeperNode)
        {
            HeadNode = headNode;
            WorkerNode = workerNode;
            ZookeeperNode = zookeeperNode;
        }
    }

    [OutputType]
    public sealed class SparkClusterRolesHeadNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private SparkClusterRolesHeadNode(
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
    public sealed class SparkClusterRolesWorkerNode
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
        private SparkClusterRolesWorkerNode(
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
    public sealed class SparkClusterRolesZookeeperNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private SparkClusterRolesZookeeperNode(
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
    public sealed class SparkClusterStorageAccountGen2
    {
        public readonly string FilesystemId;
        public readonly bool IsDefault;
        public readonly string ManagedIdentityResourceId;
        public readonly string StorageResourceId;

        [OutputConstructor]
        private SparkClusterStorageAccountGen2(
            string filesystemId,
            bool isDefault,
            string managedIdentityResourceId,
            string storageResourceId)
        {
            FilesystemId = filesystemId;
            IsDefault = isDefault;
            ManagedIdentityResourceId = managedIdentityResourceId;
            StorageResourceId = storageResourceId;
        }
    }

    [OutputType]
    public sealed class SparkClusterStorageAccounts
    {
        public readonly bool IsDefault;
        public readonly string StorageAccountKey;
        public readonly string StorageContainerId;

        [OutputConstructor]
        private SparkClusterStorageAccounts(
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
