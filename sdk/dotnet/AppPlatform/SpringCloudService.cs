// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform
{
    /// <summary>
    /// Manages an Azure Spring Cloud Service.
    /// </summary>
    public partial class SpringCloudService : Pulumi.CustomResource
    {
        /// <summary>
        /// A `config_server_git_setting` block as defined below.
        /// </summary>
        [Output("configServerGitSetting")]
        public Output<Outputs.SpringCloudServiceConfigServerGitSetting?> ConfigServerGitSetting { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a SpringCloudService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SpringCloudService(string name, SpringCloudServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:appplatform/springCloudService:SpringCloudService", name, args ?? new SpringCloudServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SpringCloudService(string name, Input<string> id, SpringCloudServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:appplatform/springCloudService:SpringCloudService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SpringCloudService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SpringCloudService Get(string name, Input<string> id, SpringCloudServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new SpringCloudService(name, id, state, options);
        }
    }

    public sealed class SpringCloudServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `config_server_git_setting` block as defined below.
        /// </summary>
        [Input("configServerGitSetting")]
        public Input<Inputs.SpringCloudServiceConfigServerGitSettingArgs>? ConfigServerGitSetting { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SpringCloudServiceArgs()
        {
        }
    }

    public sealed class SpringCloudServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `config_server_git_setting` block as defined below.
        /// </summary>
        [Input("configServerGitSetting")]
        public Input<Inputs.SpringCloudServiceConfigServerGitSettingGetArgs>? ConfigServerGitSetting { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SpringCloudServiceState()
        {
        }
    }
}