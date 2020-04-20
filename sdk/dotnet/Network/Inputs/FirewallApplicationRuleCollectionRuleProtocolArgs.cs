// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallApplicationRuleCollectionRuleProtocolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify a port for the connection.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Specifies the type of connection. Possible values are `Http`, `Https` and `Mssql`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FirewallApplicationRuleCollectionRuleProtocolArgs()
        {
        }
    }
}