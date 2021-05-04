// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class GetKubernetesClusterAddonProfileIngressApplicationGatewayResult
    {
        /// <summary>
        /// The ID of the Application Gateway associated with the ingress controller deployed to this Kubernetes Cluster.
        /// </summary>
        public readonly string EffectiveGatewayId;
        /// <summary>
        /// Is Role Based Access Control enabled?
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID of the Application Gateway integrated with the ingress controller of this Kubernetes Cluster. This attribute is only set when gateway_id is specified when configuring the `ingress_application_gateway` addon.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The subnet CIDR used to create an Application Gateway, which in turn will be integrated with the ingress controller of this Kubernetes Cluster. This attribute is only set when `subnet_cidr` is specified when configuring the `ingress_application_gateway` addon.
        /// </summary>
        public readonly string SubnetCidr;
        /// <summary>
        /// The ID of the subnet on which to create an Application Gateway, which in turn will be integrated with the ingress controller of this Kubernetes Cluster. This attribute is only set when `subnet_id` is specified when configuring the `ingress_application_gateway` addon.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private GetKubernetesClusterAddonProfileIngressApplicationGatewayResult(
            string effectiveGatewayId,

            bool enabled,

            string gatewayId,

            string subnetCidr,

            string subnetId)
        {
            EffectiveGatewayId = effectiveGatewayId;
            Enabled = enabled;
            GatewayId = gatewayId;
            SubnetCidr = subnetCidr;
            SubnetId = subnetId;
        }
    }
}
