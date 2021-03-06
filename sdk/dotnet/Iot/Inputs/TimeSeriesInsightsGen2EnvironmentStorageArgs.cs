// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Inputs
{

    public sealed class TimeSeriesInsightsGen2EnvironmentStorageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access key of storage account for Azure IoT Time Series Insights Gen2 Environment
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Name of storage account for Azure IoT Time Series Insights Gen2 Environment
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TimeSeriesInsightsGen2EnvironmentStorageArgs()
        {
        }
    }
}
