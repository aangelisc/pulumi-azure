// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage.Inputs
{

    public sealed class AccountNetworkRulesPrivateLinkAccessRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource id of the resource access rule to be granted access.
        /// </summary>
        [Input("endpointResourceId", required: true)]
        public Input<string> EndpointResourceId { get; set; } = null!;

        /// <summary>
        /// The tenant id of the resource of the resource access rule to be granted access. Defaults to the current tenant id.
        /// </summary>
        [Input("endpointTenantId")]
        public Input<string>? EndpointTenantId { get; set; }

        public AccountNetworkRulesPrivateLinkAccessRuleArgs()
        {
        }
    }
}
