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
    /// Manages the association between a Network Interface and a Load Balancer's Backend Address Pool.
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
    ///             Location = "West Europe",
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
    ///                 "10.0.2.0/24",
    ///             },
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Static",
    ///         });
    ///         var exampleLoadBalancer = new Azure.Lb.LoadBalancer("exampleLoadBalancer", new Azure.Lb.LoadBalancerArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             FrontendIpConfigurations = 
    ///             {
    ///                 new Azure.Lb.Inputs.LoadBalancerFrontendIpConfigurationArgs
    ///                 {
    ///                     Name = "primary",
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                 },
    ///             },
    ///         });
    ///         var exampleBackendAddressPool = new Azure.Lb.BackendAddressPool("exampleBackendAddressPool", new Azure.Lb.BackendAddressPoolArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             LoadbalancerId = exampleLoadBalancer.Id,
    ///         });
    ///         var exampleNetworkInterface = new Azure.Network.NetworkInterface("exampleNetworkInterface", new Azure.Network.NetworkInterfaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
    ///                 {
    ///                     Name = "testconfiguration1",
    ///                     SubnetId = exampleSubnet.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                 },
    ///             },
    ///         });
    ///         var exampleNetworkInterfaceBackendAddressPoolAssociation = new Azure.Network.NetworkInterfaceBackendAddressPoolAssociation("exampleNetworkInterfaceBackendAddressPoolAssociation", new Azure.Network.NetworkInterfaceBackendAddressPoolAssociationArgs
    ///         {
    ///             NetworkInterfaceId = exampleNetworkInterface.Id,
    ///             IpConfigurationName = "testconfiguration1",
    ///             BackendAddressPoolId = exampleBackendAddressPool.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NetworkInterfaceBackendAddressPoolAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
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
        /// Create a NetworkInterfaceBackendAddressPoolAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceBackendAddressPoolAssociation(string name, NetworkInterfaceBackendAddressPoolAssociationArgs args, CustomResourceOptions? options = null)
            : base("azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation", name, args ?? new NetworkInterfaceBackendAddressPoolAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceBackendAddressPoolAssociation(string name, Input<string> id, NetworkInterfaceBackendAddressPoolAssociationState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceBackendAddressPoolAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceBackendAddressPoolAssociation Get(string name, Input<string> id, NetworkInterfaceBackendAddressPoolAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceBackendAddressPoolAssociation(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceBackendAddressPoolAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
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

        public NetworkInterfaceBackendAddressPoolAssociationArgs()
        {
        }
    }

    public sealed class NetworkInterfaceBackendAddressPoolAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
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

        public NetworkInterfaceBackendAddressPoolAssociationState()
        {
        }
    }
}
