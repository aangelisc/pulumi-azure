// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages a Data Lake Gen2 File System within an Azure Storage Account.
    /// 
    /// &gt; **NOTE:** This Resource requires using Azure Active Directory to connect to Azure Storage, which in turn requires the `Storage` specific roles - which are not granted by default.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_data_lake_gen2_filesystem.html.markdown.
    /// </summary>
    public partial class DataLakeGen2Filesystem : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
        /// </summary>
        [Output("properties")]
        public Output<ImmutableDictionary<string, string>?> Properties { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string> StorageAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a DataLakeGen2Filesystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataLakeGen2Filesystem(string name, DataLakeGen2FilesystemArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/dataLakeGen2Filesystem:DataLakeGen2Filesystem", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DataLakeGen2Filesystem(string name, Input<string> id, DataLakeGen2FilesystemState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/dataLakeGen2Filesystem:DataLakeGen2Filesystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataLakeGen2Filesystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataLakeGen2Filesystem Get(string name, Input<string> id, DataLakeGen2FilesystemState? state = null, CustomResourceOptions? options = null)
        {
            return new DataLakeGen2Filesystem(name, id, state, options);
        }
    }

    public sealed class DataLakeGen2FilesystemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public DataLakeGen2FilesystemArgs()
        {
        }
    }

    public sealed class DataLakeGen2FilesystemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        public DataLakeGen2FilesystemState()
        {
        }
    }
}
