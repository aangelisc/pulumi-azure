// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterDefaultNodePoolArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// A list of Availability Zones across which the Node Pool should be spread. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// Should [the Kubernetes Auto Scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler) be enabled for this Node Pool? Defaults to `false`.
        /// </summary>
        [Input("enableAutoScaling")]
        public Input<bool>? EnableAutoScaling { get; set; }

        /// <summary>
        /// Should the nodes in the Default Node Pool have host encryption enabled? Defaults to `false`.
        /// </summary>
        [Input("enableHostEncryption")]
        public Input<bool>? EnableHostEncryption { get; set; }

        /// <summary>
        /// Should nodes in this Node Pool have a Public IP Address? Defaults to `false`.
        /// </summary>
        [Input("enableNodePublicIp")]
        public Input<bool>? EnableNodePublicIp { get; set; }

        /// <summary>
        /// The maximum number of nodes which should exist in this Node Pool. If specified this must be between `1` and `1000`.
        /// </summary>
        [Input("maxCount")]
        public Input<int>? MaxCount { get; set; }

        /// <summary>
        /// The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
        /// </summary>
        [Input("maxPods")]
        public Input<int>? MaxPods { get; set; }

        /// <summary>
        /// The minimum number of nodes which should exist in this Node Pool. If specified this must be between `1` and `1000`.
        /// </summary>
        [Input("minCount")]
        public Input<int>? MinCount { get; set; }

        /// <summary>
        /// The name which should be used for the default Kubernetes Node Pool. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The initial number of nodes which should exist in this Node Pool. If specified this must be between `1` and `1000` and between `min_count` and `max_count`.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("nodeLabels")]
        private InputMap<string>? _nodeLabels;

        /// <summary>
        /// A map of Kubernetes labels which should be applied to nodes in the Default Node Pool. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> NodeLabels
        {
            get => _nodeLabels ?? (_nodeLabels = new InputMap<string>());
            set => _nodeLabels = value;
        }

        [Input("nodeTaints")]
        private InputList<string>? _nodeTaints;
        public InputList<string> NodeTaints
        {
            get => _nodeTaints ?? (_nodeTaints = new InputList<string>());
            set => _nodeTaints = value;
        }

        /// <summary>
        /// Enabling this option will taint default node pool with `CriticalAddonsOnly=true:NoSchedule` taint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("onlyCriticalAddonsEnabled")]
        public Input<bool>? OnlyCriticalAddonsEnabled { get; set; }

        /// <summary>
        /// Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
        /// </summary>
        [Input("orchestratorVersion")]
        public Input<string>? OrchestratorVersion { get; set; }

        /// <summary>
        /// The size of the OS Disk which should be used for each agent in the Node Pool. Changing this forces a new resource to be created.
        /// </summary>
        [Input("osDiskSizeGb")]
        public Input<int>? OsDiskSizeGb { get; set; }

        /// <summary>
        /// The type of disk which should be used for the Operating System. Possible values are `Ephemeral` and `Managed`. Defaults to `Managed`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("osDiskType")]
        public Input<string>? OsDiskType { get; set; }

        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Node Pool.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of Node Pool which should be created. Possible values are `AvailabilitySet` and `VirtualMachineScaleSets`. Defaults to `VirtualMachineScaleSets`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A `upgrade_settings` block as documented below.
        /// </summary>
        [Input("upgradeSettings")]
        public Input<Inputs.KubernetesClusterDefaultNodePoolUpgradeSettingsArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The size of the Virtual Machine, such as `Standard_DS2_v2`.
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        /// <summary>
        /// The ID of a Subnet where the Kubernetes Node Pool should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vnetSubnetId")]
        public Input<string>? VnetSubnetId { get; set; }

        public KubernetesClusterDefaultNodePoolArgs()
        {
        }
    }
}
