// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb
{
    /// <summary>
    /// Manages a Load Balancer Outbound Rule.
    /// 
    /// &gt; **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration and a Backend Address Pool Attached.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb_outbound_rule.html.markdown.
    /// </summary>
    public partial class OutboundRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The number of outbound ports to be used for NAT.
        /// </summary>
        [Output("allocatedOutboundPorts")]
        public Output<int?> AllocatedOutboundPorts { get; private set; } = null!;

        /// <summary>
        /// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        [Output("backendAddressPoolId")]
        public Output<string> BackendAddressPoolId { get; private set; } = null!;

        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Output("enableTcpReset")]
        public Output<bool?> EnableTcpReset { get; private set; } = null!;

        /// <summary>
        /// One or more `frontend_ip_configuration` blocks as defined below.
        /// </summary>
        [Output("frontendIpConfigurations")]
        public Output<ImmutableArray<Outputs.OutboundRuleFrontendIpConfigurations>> FrontendIpConfigurations { get; private set; } = null!;

        /// <summary>
        /// The timeout for the TCP idle connection
        /// </summary>
        [Output("idleTimeoutInMinutes")]
        public Output<int?> IdleTimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("loadbalancerId")]
        public Output<string> LoadbalancerId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a OutboundRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutboundRule(string name, OutboundRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:lb/outboundRule:OutboundRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private OutboundRule(string name, Input<string> id, OutboundRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:lb/outboundRule:OutboundRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OutboundRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutboundRule Get(string name, Input<string> id, OutboundRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new OutboundRule(name, id, state, options);
        }
    }

    public sealed class OutboundRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of outbound ports to be used for NAT.
        /// </summary>
        [Input("allocatedOutboundPorts")]
        public Input<int>? AllocatedOutboundPorts { get; set; }

        /// <summary>
        /// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        [Input("backendAddressPoolId", required: true)]
        public Input<string> BackendAddressPoolId { get; set; } = null!;

        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Input("enableTcpReset")]
        public Input<bool>? EnableTcpReset { get; set; }

        [Input("frontendIpConfigurations")]
        private InputList<Inputs.OutboundRuleFrontendIpConfigurationsArgs>? _frontendIpConfigurations;

        /// <summary>
        /// One or more `frontend_ip_configuration` blocks as defined below.
        /// </summary>
        public InputList<Inputs.OutboundRuleFrontendIpConfigurationsArgs> FrontendIpConfigurations
        {
            get => _frontendIpConfigurations ?? (_frontendIpConfigurations = new InputList<Inputs.OutboundRuleFrontendIpConfigurationsArgs>());
            set => _frontendIpConfigurations = value;
        }

        /// <summary>
        /// The timeout for the TCP idle connection
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("loadbalancerId", required: true)]
        public Input<string> LoadbalancerId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public OutboundRuleArgs()
        {
        }
    }

    public sealed class OutboundRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of outbound ports to be used for NAT.
        /// </summary>
        [Input("allocatedOutboundPorts")]
        public Input<int>? AllocatedOutboundPorts { get; set; }

        /// <summary>
        /// The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        [Input("backendAddressPoolId")]
        public Input<string>? BackendAddressPoolId { get; set; }

        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Input("enableTcpReset")]
        public Input<bool>? EnableTcpReset { get; set; }

        [Input("frontendIpConfigurations")]
        private InputList<Inputs.OutboundRuleFrontendIpConfigurationsGetArgs>? _frontendIpConfigurations;

        /// <summary>
        /// One or more `frontend_ip_configuration` blocks as defined below.
        /// </summary>
        public InputList<Inputs.OutboundRuleFrontendIpConfigurationsGetArgs> FrontendIpConfigurations
        {
            get => _frontendIpConfigurations ?? (_frontendIpConfigurations = new InputList<Inputs.OutboundRuleFrontendIpConfigurationsGetArgs>());
            set => _frontendIpConfigurations = value;
        }

        /// <summary>
        /// The timeout for the TCP idle connection
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        /// <summary>
        /// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public OutboundRuleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class OutboundRuleFrontendIpConfigurationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Load Balancer Outbound Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public OutboundRuleFrontendIpConfigurationsArgs()
        {
        }
    }

    public sealed class OutboundRuleFrontendIpConfigurationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Load Balancer Outbound Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public OutboundRuleFrontendIpConfigurationsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class OutboundRuleFrontendIpConfigurations
    {
        /// <summary>
        /// The ID of the Load Balancer Outbound Rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private OutboundRuleFrontendIpConfigurations(
            string id,
            string name)
        {
            Id = id;
            Name = name;
        }
    }
    }
}
