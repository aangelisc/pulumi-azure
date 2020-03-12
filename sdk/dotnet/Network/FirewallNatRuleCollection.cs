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
    /// Manages a NAT Rule Collection within an Azure Firewall.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/firewall_nat_rule_collection.html.markdown.
    /// </summary>
    public partial class FirewallNatRuleCollection : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Dnat` and `Snat`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("azureFirewallName")]
        public Output<string> AzureFirewallName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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
        public Output<ImmutableArray<Outputs.FirewallNatRuleCollectionRules>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallNatRuleCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallNatRuleCollection(string name, FirewallNatRuleCollectionArgs args, CustomResourceOptions? options = null)
            : base("azure:network/firewallNatRuleCollection:FirewallNatRuleCollection", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private FirewallNatRuleCollection(string name, Input<string> id, FirewallNatRuleCollectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/firewallNatRuleCollection:FirewallNatRuleCollection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallNatRuleCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallNatRuleCollection Get(string name, Input<string> id, FirewallNatRuleCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallNatRuleCollection(name, id, state, options);
        }
    }

    public sealed class FirewallNatRuleCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Dnat` and `Snat`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("azureFirewallName", required: true)]
        public Input<string> AzureFirewallName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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
        private InputList<Inputs.FirewallNatRuleCollectionRulesArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallNatRuleCollectionRulesArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallNatRuleCollectionRulesArgs>());
            set => _rules = value;
        }

        public FirewallNatRuleCollectionArgs()
        {
        }
    }

    public sealed class FirewallNatRuleCollectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action the rule will apply to matching traffic. Possible values are `Dnat` and `Snat`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("azureFirewallName")]
        public Input<string>? AzureFirewallName { get; set; }

        /// <summary>
        /// Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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
        private InputList<Inputs.FirewallNatRuleCollectionRulesGetArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallNatRuleCollectionRulesGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallNatRuleCollectionRulesGetArgs>());
            set => _rules = value;
        }

        public FirewallNatRuleCollectionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class FirewallNatRuleCollectionRulesArgs : Pulumi.ResourceArgs
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
        /// Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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

        [Input("translatedAddress", required: true)]
        public Input<string> TranslatedAddress { get; set; } = null!;

        [Input("translatedPort", required: true)]
        public Input<string> TranslatedPort { get; set; } = null!;

        public FirewallNatRuleCollectionRulesArgs()
        {
        }
    }

    public sealed class FirewallNatRuleCollectionRulesGetArgs : Pulumi.ResourceArgs
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
        /// Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
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

        [Input("translatedAddress", required: true)]
        public Input<string> TranslatedAddress { get; set; } = null!;

        [Input("translatedPort", required: true)]
        public Input<string> TranslatedPort { get; set; } = null!;

        public FirewallNatRuleCollectionRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class FirewallNatRuleCollectionRules
    {
        public readonly string? Description;
        public readonly ImmutableArray<string> DestinationAddresses;
        public readonly ImmutableArray<string> DestinationPorts;
        /// <summary>
        /// Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> Protocols;
        public readonly ImmutableArray<string> SourceAddresses;
        public readonly string TranslatedAddress;
        public readonly string TranslatedPort;

        [OutputConstructor]
        private FirewallNatRuleCollectionRules(
            string? description,
            ImmutableArray<string> destinationAddresses,
            ImmutableArray<string> destinationPorts,
            string name,
            ImmutableArray<string> protocols,
            ImmutableArray<string> sourceAddresses,
            string translatedAddress,
            string translatedPort)
        {
            Description = description;
            DestinationAddresses = destinationAddresses;
            DestinationPorts = destinationPorts;
            Name = name;
            Protocols = protocols;
            SourceAddresses = sourceAddresses;
            TranslatedAddress = translatedAddress;
            TranslatedPort = translatedPort;
        }
    }
    }
}
