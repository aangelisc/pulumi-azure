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
    /// Manages a Data Share.
    /// </summary>
    public partial class Share : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The Data Share's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The kind of the Data Share. Possible values are `CopyBased` and `InPlace`. Changing this forces a new Data Share to be created.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Data Share. Changing this forces a new Data Share to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `snapshot_schedule` block as defined below.
        /// </summary>
        [Output("snapshotSchedule")]
        public Output<Outputs.ShareSnapshotSchedule?> SnapshotSchedule { get; private set; } = null!;

        /// <summary>
        /// The terms of the Data Share.
        /// </summary>
        [Output("terms")]
        public Output<string?> Terms { get; private set; } = null!;


        /// <summary>
        /// Create a Share resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Share(string name, ShareArgs args, CustomResourceOptions? options = null)
            : base("azure:datashare/share:Share", name, args ?? new ShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Share(string name, Input<string> id, ShareState? state = null, CustomResourceOptions? options = null)
            : base("azure:datashare/share:Share", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Share resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Share Get(string name, Input<string> id, ShareState? state = null, CustomResourceOptions? options = null)
        {
            return new Share(name, id, state, options);
        }
    }

    public sealed class ShareArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The Data Share's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The kind of the Data Share. Possible values are `CopyBased` and `InPlace`. Changing this forces a new Data Share to be created.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Data Share. Changing this forces a new Data Share to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `snapshot_schedule` block as defined below.
        /// </summary>
        [Input("snapshotSchedule")]
        public Input<Inputs.ShareSnapshotScheduleArgs>? SnapshotSchedule { get; set; }

        /// <summary>
        /// The terms of the Data Share.
        /// </summary>
        [Input("terms")]
        public Input<string>? Terms { get; set; }

        public ShareArgs()
        {
        }
    }

    public sealed class ShareState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The Data Share's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The kind of the Data Share. Possible values are `CopyBased` and `InPlace`. Changing this forces a new Data Share to be created.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The name which should be used for this Data Share. Changing this forces a new Data Share to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `snapshot_schedule` block as defined below.
        /// </summary>
        [Input("snapshotSchedule")]
        public Input<Inputs.ShareSnapshotScheduleGetArgs>? SnapshotSchedule { get; set; }

        /// <summary>
        /// The terms of the Data Share.
        /// </summary>
        [Input("terms")]
        public Input<string>? Terms { get; set; }

        public ShareState()
        {
        }
    }
}