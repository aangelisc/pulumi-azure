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
    /// Configures a Network Connection Monitor to monitor communication between a Virtual Machine and an endpoint using a Network Watcher.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_connection_monitor.html.markdown.
    /// </summary>
    public partial class NetworkConnectionMonitor : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether the connection monitor will start automatically once created. Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("autoStart")]
        public Output<bool?> AutoStart { get; private set; } = null!;

        /// <summary>
        /// A `destination` block as defined below.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.NetworkConnectionMonitorDestination> Destination { get; private set; } = null!;

        /// <summary>
        /// Monitoring interval in seconds. Defaults to `60`.
        /// </summary>
        [Output("intervalInSeconds")]
        public Output<int?> IntervalInSeconds { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Network Connection Monitor. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Network Watcher. Changing this forces a new resource to be created.
        /// </summary>
        [Output("networkWatcherName")]
        public Output<string> NetworkWatcherName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Connection Monitor. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `source` block as defined below.
        /// </summary>
        [Output("source")]
        public Output<Outputs.NetworkConnectionMonitorSource> Source { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkConnectionMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkConnectionMonitor(string name, NetworkConnectionMonitorArgs args, CustomResourceOptions? options = null)
            : base("azure:network/networkConnectionMonitor:NetworkConnectionMonitor", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkConnectionMonitor(string name, Input<string> id, NetworkConnectionMonitorState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/networkConnectionMonitor:NetworkConnectionMonitor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkConnectionMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkConnectionMonitor Get(string name, Input<string> id, NetworkConnectionMonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkConnectionMonitor(name, id, state, options);
        }
    }

    public sealed class NetworkConnectionMonitorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the connection monitor will start automatically once created. Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("autoStart")]
        public Input<bool>? AutoStart { get; set; }

        /// <summary>
        /// A `destination` block as defined below.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.NetworkConnectionMonitorDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Monitoring interval in seconds. Defaults to `60`.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Network Connection Monitor. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Network Watcher. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Connection Monitor. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `source` block as defined below.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.NetworkConnectionMonitorSourceArgs> Source { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public NetworkConnectionMonitorArgs()
        {
        }
    }

    public sealed class NetworkConnectionMonitorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the connection monitor will start automatically once created. Defaults to `true`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("autoStart")]
        public Input<bool>? AutoStart { get; set; }

        /// <summary>
        /// A `destination` block as defined below.
        /// </summary>
        [Input("destination")]
        public Input<Inputs.NetworkConnectionMonitorDestinationGetArgs>? Destination { get; set; }

        /// <summary>
        /// Monitoring interval in seconds. Defaults to `60`.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Network Connection Monitor. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Network Watcher. Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkWatcherName")]
        public Input<string>? NetworkWatcherName { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Connection Monitor. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `source` block as defined below.
        /// </summary>
        [Input("source")]
        public Input<Inputs.NetworkConnectionMonitorSourceGetArgs>? Source { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public NetworkConnectionMonitorState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NetworkConnectionMonitorDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        public NetworkConnectionMonitorDestinationArgs()
        {
        }
    }

    public sealed class NetworkConnectionMonitorDestinationGetArgs : Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        public NetworkConnectionMonitorDestinationGetArgs()
        {
        }
    }

    public sealed class NetworkConnectionMonitorSourceArgs : Pulumi.ResourceArgs
    {
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("virtualMachineId", required: true)]
        public Input<string> VirtualMachineId { get; set; } = null!;

        public NetworkConnectionMonitorSourceArgs()
        {
        }
    }

    public sealed class NetworkConnectionMonitorSourceGetArgs : Pulumi.ResourceArgs
    {
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("virtualMachineId", required: true)]
        public Input<string> VirtualMachineId { get; set; } = null!;

        public NetworkConnectionMonitorSourceGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NetworkConnectionMonitorDestination
    {
        public readonly string? Address;
        public readonly int Port;
        public readonly string? VirtualMachineId;

        [OutputConstructor]
        private NetworkConnectionMonitorDestination(
            string? address,
            int port,
            string? virtualMachineId)
        {
            Address = address;
            Port = port;
            VirtualMachineId = virtualMachineId;
        }
    }

    [OutputType]
    public sealed class NetworkConnectionMonitorSource
    {
        public readonly int? Port;
        public readonly string VirtualMachineId;

        [OutputConstructor]
        private NetworkConnectionMonitorSource(
            int? port,
            string virtualMachineId)
        {
            Port = port;
            VirtualMachineId = virtualMachineId;
        }
    }
    }
}
