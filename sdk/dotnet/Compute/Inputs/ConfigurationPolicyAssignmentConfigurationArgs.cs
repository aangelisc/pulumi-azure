// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ConfigurationPolicyAssignmentConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Guest Configuration that will be assigned in this Guest Configuration Assignment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("parameters")]
        private InputList<Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs>? _parameters;

        /// <summary>
        /// One or more `parameter` blocks which define what configuration parameters and values against.
        /// </summary>
        public InputList<Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ConfigurationPolicyAssignmentConfigurationParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The version of the Guest Configuration that will be assigned in this Guest Configuration Assignment.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ConfigurationPolicyAssignmentConfigurationArgs()
        {
        }
    }
}
