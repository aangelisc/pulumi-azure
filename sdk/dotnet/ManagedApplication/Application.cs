// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ManagedApplication
{
    /// <summary>
    /// Manages a Managed Application.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var builtin = Output.Create(Azure.Authorization.GetRoleDefinition.InvokeAsync(new Azure.Authorization.GetRoleDefinitionArgs
    ///         {
    ///             Name = "Contributor",
    ///         }));
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleDefinition = new Azure.ManagedApplication.Definition("exampleDefinition", new Azure.ManagedApplication.DefinitionArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             LockLevel = "ReadOnly",
    ///             PackageFileUri = "https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip",
    ///             DisplayName = "TestManagedAppDefinition",
    ///             Description = "Test Managed App Definition",
    ///             Authorizations = 
    ///             {
    ///                 new Azure.ManagedApplication.Inputs.DefinitionAuthorizationArgs
    ///                 {
    ///                     ServicePrincipalId = current.Apply(current =&gt; current.ObjectId),
    ///                     RoleDefinitionId = Output.Tuple(builtin.Apply(builtin =&gt; builtin.Id.Split("/")), builtin.Apply(builtin =&gt; builtin.Id.Split("/")).Length).Apply(values =&gt;
    ///                     {
    ///                         var split = values.Item1;
    ///                         var length = values.Item2;
    ///                         return split[length - 1];
    ///                     }),
    ///                 },
    ///             },
    ///         });
    ///         var exampleApplication = new Azure.ManagedApplication.Application("exampleApplication", new Azure.ManagedApplication.ApplicationArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Kind = "ServiceCatalog",
    ///             ManagedResourceGroupName = "infrastructureGroup",
    ///             ApplicationDefinitionId = exampleDefinition.Id,
    ///             Parameters = 
    ///             {
    ///                 { "location", exampleResourceGroup.Location },
    ///                 { "storageAccountNamePrefix", "storeNamePrefix" },
    ///                 { "storageAccountType", "Standard_LRS" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Managed Application can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:managedapplication/application:Application example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applications/app1
    /// ```
    /// </summary>
    [AzureResourceType("azure:managedapplication/application:Application")]
    public partial class Application : Pulumi.CustomResource
    {
        /// <summary>
        /// The application definition ID to deploy.
        /// </summary>
        [Output("applicationDefinitionId")]
        public Output<string?> ApplicationDefinitionId { get; private set; } = null!;

        /// <summary>
        /// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
        /// </summary>
        [Output("managedResourceGroupName")]
        public Output<string> ManagedResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name and value pairs that define the managed application outputs.
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableDictionary<string, string>> Outputs { get; private set; } = null!;

        /// <summary>
        /// The parameter values to pass to the Managed Application. This field is a json object that allows you to assign parameters to this Managed Application.
        /// </summary>
        [Output("parameterValues")]
        public Output<string> ParameterValues { get; private set; } = null!;

        /// <summary>
        /// A mapping of name and value pairs to pass to the managed application as parameters.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>> Parameters { get; private set; } = null!;

        /// <summary>
        /// One `plan` block as defined below.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.ApplicationPlan?> Plan { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("azure:managedapplication/application:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
            : base("azure:managedapplication/application:Application", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Application(name, id, state, options);
        }
    }

    public sealed class ApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application definition ID to deploy.
        /// </summary>
        [Input("applicationDefinitionId")]
        public Input<string>? ApplicationDefinitionId { get; set; }

        /// <summary>
        /// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedResourceGroupName", required: true)]
        public Input<string> ManagedResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parameter values to pass to the Managed Application. This field is a json object that allows you to assign parameters to this Managed Application.
        /// </summary>
        [Input("parameterValues")]
        public Input<string>? ParameterValues { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A mapping of name and value pairs to pass to the managed application as parameters.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// One `plan` block as defined below.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.ApplicationPlanArgs>? Plan { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
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

        public ApplicationArgs()
        {
        }
    }

    public sealed class ApplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application definition ID to deploy.
        /// </summary>
        [Input("applicationDefinitionId")]
        public Input<string>? ApplicationDefinitionId { get; set; }

        /// <summary>
        /// The kind of the managed application to deploy. Possible values are `MarketPlace` and `ServiceCatalog`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the target resource group where all the resources deployed by the managed application will reside. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedResourceGroupName")]
        public Input<string>? ManagedResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outputs")]
        private InputMap<string>? _outputs;

        /// <summary>
        /// The name and value pairs that define the managed application outputs.
        /// </summary>
        public InputMap<string> Outputs
        {
            get => _outputs ?? (_outputs = new InputMap<string>());
            set => _outputs = value;
        }

        /// <summary>
        /// The parameter values to pass to the Managed Application. This field is a json object that allows you to assign parameters to this Managed Application.
        /// </summary>
        [Input("parameterValues")]
        public Input<string>? ParameterValues { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A mapping of name and value pairs to pass to the managed application as parameters.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// One `plan` block as defined below.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.ApplicationPlanGetArgs>? Plan { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Managed Application should exist. Changing this forces a new resource to be created.
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

        public ApplicationState()
        {
        }
    }
}
