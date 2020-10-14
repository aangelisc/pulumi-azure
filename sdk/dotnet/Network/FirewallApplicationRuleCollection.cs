// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages an Application Rule Collection within an Azure Firewall.
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
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "North Europe",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/24",
    ///             },
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Static",
    ///             Sku = "Standard",
    ///         });
    ///         var exampleFirewall = new Azure.Network.Firewall("exampleFirewall", new Azure.Network.FirewallArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.FirewallIpConfigurationArgs
    ///                 {
    ///                     Name = "configuration",
    ///                     SubnetId = exampleSubnet.Id,
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                 },
    ///             },
    ///         });
    ///         var exampleFirewallApplicationRuleCollection = new Azure.Network.FirewallApplicationRuleCollection("exampleFirewallApplicationRuleCollection", new Azure.Network.FirewallApplicationRuleCollectionArgs
    ///         {
    ///             AzureFirewallName = exampleFirewall.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Priority = 100,
    ///             Action = "Allow",
    ///             Rules = 
    ///             {
    ///                 new Azure.Network.Inputs.FirewallApplicationRuleCollectionRuleArgs
    ///                 {
    ///                     Name = "testrule",
    ///                     SourceAddresses = 
    ///                     {
    ///                         "10.0.0.0/16",
    ///                     },
    ///                     TargetFqdns = 
    ///                     {
    ///                         "*.google.com",
    ///                     },
    ///                     Protocols = 
    ///                     {
    ///                         new Azure.Network.Inputs.FirewallApplicationRuleCollectionRuleProtocolArgs
    ///                         {
    ///                             Port = 443,
    ///                             Type = "Https",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class FirewallApplicationRuleCollection : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("azureFirewallName")]
        public Output<string> AzureFirewallName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.FirewallApplicationRuleCollectionRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallApplicationRuleCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallApplicationRuleCollection(string name, FirewallApplicationRuleCollectionArgs args, CustomResourceOptions? options = null)
            : base("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, args ?? new FirewallApplicationRuleCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallApplicationRuleCollection(string name, Input<string> id, FirewallApplicationRuleCollectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallApplicationRuleCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallApplicationRuleCollection Get(string name, Input<string> id, FirewallApplicationRuleCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallApplicationRuleCollection(name, id, state, options);
        }
    }

    public sealed class FirewallApplicationRuleCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("azureFirewallName", required: true)]
        public Input<string> AzureFirewallName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.FirewallApplicationRuleCollectionRuleArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallApplicationRuleCollectionRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallApplicationRuleCollectionRuleArgs>());
            set => _rules = value;
        }

        public FirewallApplicationRuleCollectionArgs()
        {
        }
    }

    public sealed class FirewallApplicationRuleCollectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("azureFirewallName")]
        public Input<string>? AzureFirewallName { get; set; }

        /// <summary>
        /// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("rules")]
        private InputList<Inputs.FirewallApplicationRuleCollectionRuleGetArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallApplicationRuleCollectionRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallApplicationRuleCollectionRuleGetArgs>());
            set => _rules = value;
        }

        public FirewallApplicationRuleCollectionState()
        {
        }
    }
}
