// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages a Container within an Azure Storage Account.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_container.html.markdown.
    /// </summary>
    public partial class Container : Pulumi.CustomResource
    {
        /// <summary>
        /// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
        /// </summary>
        [Output("containerAccessType")]
        public Output<string?> ContainerAccessType { get; private set; } = null!;

        /// <summary>
        /// Is there an Immutability Policy configured on this Storage Container?
        /// </summary>
        [Output("hasImmutabilityPolicy")]
        public Output<bool> HasImmutabilityPolicy { get; private set; } = null!;

        /// <summary>
        /// Is there a Legal Hold configured on this Storage Container?
        /// </summary>
        [Output("hasLegalHold")]
        public Output<bool> HasLegalHold { get; private set; } = null!;

        /// <summary>
        /// A mapping of MetaData for this Container.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the Container which should be created within the Storage Account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Storage Account where the Container should be created.
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;


        /// <summary>
        /// Create a Container resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Container(string name, ContainerArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/container:Container", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Container(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/container:Container", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Container resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Container Get(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
        {
            return new Container(name, id, state, options);
        }
    }

    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
        /// </summary>
        [Input("containerAccessType")]
        public Input<string>? ContainerAccessType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A mapping of MetaData for this Container.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the Container which should be created within the Storage Account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Storage Account where the Container should be created.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        public ContainerArgs()
        {
        }
    }

    public sealed class ContainerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Level configured for this Container. Possible values are `blob`, `container` or `private`. Defaults to `private`.
        /// </summary>
        [Input("containerAccessType")]
        public Input<string>? ContainerAccessType { get; set; }

        /// <summary>
        /// Is there an Immutability Policy configured on this Storage Container?
        /// </summary>
        [Input("hasImmutabilityPolicy")]
        public Input<bool>? HasImmutabilityPolicy { get; set; }

        /// <summary>
        /// Is there a Legal Hold configured on this Storage Container?
        /// </summary>
        [Input("hasLegalHold")]
        public Input<bool>? HasLegalHold { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A mapping of MetaData for this Container.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the Container which should be created within the Storage Account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Storage Account where the Container should be created.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        public ContainerState()
        {
        }
    }
}
