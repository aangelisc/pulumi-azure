// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class AssetFilterTrackSelectionConditionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition operation to test a track property against. Supported values are `Equal` and `NotEqual`.
        /// </summary>
        [Input("operation")]
        public Input<string>? Operation { get; set; }

        /// <summary>
        /// The track property to compare. Supported values are `Bitrate`, `FourCC`, `Language`, `Name` and `Type`. Check [documentation](https://docs.microsoft.com/en-us/azure/media-services/latest/filters-concept) for more details.
        /// </summary>
        [Input("property")]
        public Input<string>? Property { get; set; }

        /// <summary>
        /// The track property value to match or not match.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public AssetFilterTrackSelectionConditionGetArgs()
        {
        }
    }
}
