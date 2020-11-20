// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare
{
    /// <summary>
    /// Manages a Data Share Kusto Database Dataset.
    /// </summary>
    public partial class DatasetKustoDatabase : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Data Share Dataset.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The location of the Kusto Cluster.
        /// </summary>
        [Output("kustoClusterLocation")]
        public Output<string> KustoClusterLocation { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Output("kustoDatabaseId")]
        public Output<string> KustoDatabaseId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Output("shareId")]
        public Output<string> ShareId { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetKustoDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetKustoDatabase(string name, DatasetKustoDatabaseArgs args, CustomResourceOptions? options = null)
            : base("azure:datashare/datasetKustoDatabase:DatasetKustoDatabase", name, args ?? new DatasetKustoDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetKustoDatabase(string name, Input<string> id, DatasetKustoDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("azure:datashare/datasetKustoDatabase:DatasetKustoDatabase", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetKustoDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetKustoDatabase Get(string name, Input<string> id, DatasetKustoDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetKustoDatabase(name, id, state, options);
        }
    }

    public sealed class DatasetKustoDatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Input("kustoDatabaseId", required: true)]
        public Input<string> KustoDatabaseId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Input("shareId", required: true)]
        public Input<string> ShareId { get; set; } = null!;

        public DatasetKustoDatabaseArgs()
        {
        }
    }

    public sealed class DatasetKustoDatabaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Share Dataset.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The location of the Kusto Cluster.
        /// </summary>
        [Input("kustoClusterLocation")]
        public Input<string>? KustoClusterLocation { get; set; }

        /// <summary>
        /// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Input("kustoDatabaseId")]
        public Input<string>? KustoDatabaseId { get; set; }

        /// <summary>
        /// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
        /// </summary>
        [Input("shareId")]
        public Input<string>? ShareId { get; set; }

        public DatasetKustoDatabaseState()
        {
        }
    }
}