// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("destinationAddresses")]
        private InputList<string>? _destinationAddresses;

        /// <summary>
        /// Specifies a list of destination IP addresses (including CIDR and `*`) or Service Tags.
        /// </summary>
        public InputList<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new InputList<string>());
            set => _destinationAddresses = value;
        }

        [Input("destinationFqdns")]
        private InputList<string>? _destinationFqdns;

        /// <summary>
        /// Specifies a list of destination FQDNs.
        /// </summary>
        public InputList<string> DestinationFqdns
        {
            get => _destinationFqdns ?? (_destinationFqdns = new InputList<string>());
            set => _destinationFqdns = value;
        }

        [Input("destinationIpGroups")]
        private InputList<string>? _destinationIpGroups;

        /// <summary>
        /// Specifies a list of destination IP groups.
        /// </summary>
        public InputList<string> DestinationIpGroups
        {
            get => _destinationIpGroups ?? (_destinationIpGroups = new InputList<string>());
            set => _destinationIpGroups = value;
        }

        [Input("destinationPorts", required: true)]
        private InputList<string>? _destinationPorts;

        /// <summary>
        /// Specifies a list of destination ports.
        /// </summary>
        public InputList<string> DestinationPorts
        {
            get => _destinationPorts ?? (_destinationPorts = new InputList<string>());
            set => _destinationPorts = value;
        }

        /// <summary>
        /// The name which should be used for this rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("protocols", required: true)]
        private InputList<string>? _protocols;

        /// <summary>
        /// Specifies a list of network protocols this rule applies to. Possible values are `TCP`, `UDP`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("sourceAddresses")]
        private InputList<string>? _sourceAddresses;

        /// <summary>
        /// Specifies a list of source IP addresses (including CIDR and `*`).
        /// </summary>
        public InputList<string> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<string>());
            set => _sourceAddresses = value;
        }

        [Input("sourceIpGroups")]
        private InputList<string>? _sourceIpGroups;

        /// <summary>
        /// Specifies a list of source IP groups.
        /// </summary>
        public InputList<string> SourceIpGroups
        {
            get => _sourceIpGroups ?? (_sourceIpGroups = new InputList<string>());
            set => _sourceIpGroups = value;
        }

        public FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleGetArgs()
        {
        }
    }
}
