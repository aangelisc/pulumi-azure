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
    public sealed class GetClusterNodePoolUpgradeSettingResult
    {
        /// <summary>
        /// The maximum number or percentage of nodes which will be added to the Node Pool size during an upgrade.
        /// </summary>
        public readonly string MaxSurge;

        [OutputConstructor]
        private GetClusterNodePoolUpgradeSettingResult(string maxSurge)
        {
            MaxSurge = maxSurge;
        }
    }
}
