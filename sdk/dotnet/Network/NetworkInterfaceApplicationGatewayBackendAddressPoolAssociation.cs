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
    /// Manages the association between a Network Interface and a Application Gateway's Backend Address Pool.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_interface_application_gateway_backend_address_pool_association.html.markdown.
    /// </summary>
    public partial class NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("backendAddressPoolId")]
        public Output<string> BackendAddressPoolId { get; private set; } = null!;

        /// <summary>
        /// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        /// </summary>
        [Output("ipConfigurationName")]
        public Output<string> IpConfigurationName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(string name, NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs args, CustomResourceOptions? options = null)
            : base("azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(string name, Input<string> id, NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation Get(string name, Input<string> id, NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("backendAddressPoolId", required: true)]
        public Input<string> BackendAddressPoolId { get; set; } = null!;

        /// <summary>
        /// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ipConfigurationName", required: true)]
        public Input<string> IpConfigurationName { get; set; } = null!;

        /// <summary>
        /// The ID of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        public NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs()
        {
        }
    }

    public sealed class NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("backendAddressPoolId")]
        public Input<string>? BackendAddressPoolId { get; set; }

        /// <summary>
        /// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ipConfigurationName")]
        public Input<string>? IpConfigurationName { get; set; }

        /// <summary>
        /// The ID of the Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        public NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState()
        {
        }
    }
}
