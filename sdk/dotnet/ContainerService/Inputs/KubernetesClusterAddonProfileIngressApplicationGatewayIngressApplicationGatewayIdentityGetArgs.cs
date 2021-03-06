// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentityGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID for the Service Principal.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The Object ID of the user-defined Managed Identity used by the OMS Agents.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The ID of a user assigned identity.
        /// </summary>
        [Input("userAssignedIdentityId")]
        public Input<string>? UserAssignedIdentityId { get; set; }

        public KubernetesClusterAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentityGetArgs()
        {
        }
    }
}
