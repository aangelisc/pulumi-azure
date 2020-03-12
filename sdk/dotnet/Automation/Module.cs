// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation
{
    /// <summary>
    /// Manages a Automation Module.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_module.html.markdown.
    /// </summary>
    public partial class Module : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("automationAccountName")]
        public Output<string> AutomationAccountName { get; private set; } = null!;

        /// <summary>
        /// The published Module link.
        /// </summary>
        [Output("moduleLink")]
        public Output<Outputs.ModuleModuleLink> ModuleLink { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Module. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Module is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a Module resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Module(string name, ModuleArgs args, CustomResourceOptions? options = null)
            : base("azure:automation/module:Module", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Module(string name, Input<string> id, ModuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:automation/module:Module", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Module resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Module Get(string name, Input<string> id, ModuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Module(name, id, state, options);
        }
    }

    public sealed class ModuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The published Module link.
        /// </summary>
        [Input("moduleLink", required: true)]
        public Input<Inputs.ModuleModuleLinkArgs> ModuleLink { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Module. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Module is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ModuleArgs()
        {
        }
    }

    public sealed class ModuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account in which the Module is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("automationAccountName")]
        public Input<string>? AutomationAccountName { get; set; }

        /// <summary>
        /// The published Module link.
        /// </summary>
        [Input("moduleLink")]
        public Input<Inputs.ModuleModuleLinkGetArgs>? ModuleLink { get; set; }

        /// <summary>
        /// Specifies the name of the Module. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Module is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ModuleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ModuleModuleLinkArgs : Pulumi.ResourceArgs
    {
        [Input("hash")]
        public Input<ModuleModuleLinkHashArgs>? Hash { get; set; }

        /// <summary>
        /// The uri of the module content (zip or nupkg).
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public ModuleModuleLinkArgs()
        {
        }
    }

    public sealed class ModuleModuleLinkGetArgs : Pulumi.ResourceArgs
    {
        [Input("hash")]
        public Input<ModuleModuleLinkHashGetArgs>? Hash { get; set; }

        /// <summary>
        /// The uri of the module content (zip or nupkg).
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public ModuleModuleLinkGetArgs()
        {
        }
    }

    public sealed class ModuleModuleLinkHashArgs : Pulumi.ResourceArgs
    {
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ModuleModuleLinkHashArgs()
        {
        }
    }

    public sealed class ModuleModuleLinkHashGetArgs : Pulumi.ResourceArgs
    {
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ModuleModuleLinkHashGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ModuleModuleLink
    {
        public readonly ModuleModuleLinkHash? Hash;
        /// <summary>
        /// The uri of the module content (zip or nupkg).
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private ModuleModuleLink(
            ModuleModuleLinkHash? hash,
            string uri)
        {
            Hash = hash;
            Uri = uri;
        }
    }

    [OutputType]
    public sealed class ModuleModuleLinkHash
    {
        public readonly string Algorithm;
        public readonly string Value;

        [OutputConstructor]
        private ModuleModuleLinkHash(
            string algorithm,
            string value)
        {
            Algorithm = algorithm;
            Value = value;
        }
    }
    }
}
