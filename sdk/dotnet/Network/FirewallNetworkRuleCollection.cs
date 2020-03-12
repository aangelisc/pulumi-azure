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
    /// Manages a Network Rule Collection within an Azure Firewall.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/firewall_network_rule_collection.html.markdown.
    /// </summary>
    public partial class FirewallNetworkRuleCollection : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("azureFirewallName")]
        public Output<string> AzureFirewallName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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
        public Output<ImmutableArray<Outputs.FirewallNetworkRuleCollectionRules>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallNetworkRuleCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallNetworkRuleCollection(string name, FirewallNetworkRuleCollectionArgs args, CustomResourceOptions? options = null)
            : base("azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private FirewallNetworkRuleCollection(string name, Input<string> id, FirewallNetworkRuleCollectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallNetworkRuleCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallNetworkRuleCollection Get(string name, Input<string> id, FirewallNetworkRuleCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallNetworkRuleCollection(name, id, state, options);
        }
    }

    public sealed class FirewallNetworkRuleCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("azureFirewallName", required: true)]
        public Input<string> AzureFirewallName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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
        private InputList<Inputs.FirewallNetworkRuleCollectionRulesArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallNetworkRuleCollectionRulesArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallNetworkRuleCollectionRulesArgs>());
            set => _rules = value;
        }

        public FirewallNetworkRuleCollectionArgs()
        {
        }
    }

    public sealed class FirewallNetworkRuleCollectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Specifies the name of the Firewall in which the Network Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("azureFirewallName")]
        public Input<string>? AzureFirewallName { get; set; }

        /// <summary>
        /// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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
        private InputList<Inputs.FirewallNetworkRuleCollectionRulesGetArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallNetworkRuleCollectionRulesGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallNetworkRuleCollectionRulesGetArgs>());
            set => _rules = value;
        }

        public FirewallNetworkRuleCollectionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class FirewallNetworkRuleCollectionRulesArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationAddresses", required: true)]
        private InputList<string>? _destinationAddresses;
        public InputList<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new InputList<string>());
            set => _destinationAddresses = value;
        }

        [Input("destinationPorts", required: true)]
        private InputList<string>? _destinationPorts;
        public InputList<string> DestinationPorts
        {
            get => _destinationPorts ?? (_destinationPorts = new InputList<string>());
            set => _destinationPorts = value;
        }

        /// <summary>
        /// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("protocols", required: true)]
        private InputList<string>? _protocols;
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("sourceAddresses", required: true)]
        private InputList<string>? _sourceAddresses;
        public InputList<string> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<string>());
            set => _sourceAddresses = value;
        }

        public FirewallNetworkRuleCollectionRulesArgs()
        {
        }
    }

    public sealed class FirewallNetworkRuleCollectionRulesGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationAddresses", required: true)]
        private InputList<string>? _destinationAddresses;
        public InputList<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new InputList<string>());
            set => _destinationAddresses = value;
        }

        [Input("destinationPorts", required: true)]
        private InputList<string>? _destinationPorts;
        public InputList<string> DestinationPorts
        {
            get => _destinationPorts ?? (_destinationPorts = new InputList<string>());
            set => _destinationPorts = value;
        }

        /// <summary>
        /// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("protocols", required: true)]
        private InputList<string>? _protocols;
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("sourceAddresses", required: true)]
        private InputList<string>? _sourceAddresses;
        public InputList<string> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<string>());
            set => _sourceAddresses = value;
        }

        public FirewallNetworkRuleCollectionRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class FirewallNetworkRuleCollectionRules
    {
        public readonly string? Description;
        public readonly ImmutableArray<string> DestinationAddresses;
        public readonly ImmutableArray<string> DestinationPorts;
        /// <summary>
        /// Specifies the name of the Network Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> Protocols;
        public readonly ImmutableArray<string> SourceAddresses;

        [OutputConstructor]
        private FirewallNetworkRuleCollectionRules(
            string? description,
            ImmutableArray<string> destinationAddresses,
            ImmutableArray<string> destinationPorts,
            string name,
            ImmutableArray<string> protocols,
            ImmutableArray<string> sourceAddresses)
        {
            Description = description;
            DestinationAddresses = destinationAddresses;
            DestinationPorts = destinationPorts;
            Name = name;
            Protocols = protocols;
            SourceAddresses = sourceAddresses;
        }
    }
    }
}
