// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class AssetFilterPresentationTimeRangeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The absolute end time boundary. Applies to Video on Demand (VoD).
        /// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by `unit_timescale_in_miliseconds`, so an `end_in_units` of 180 would be for 3 minutes. Use `start_in_units` and `end_in_units` to trim the fragments that will be in the playlist (manifest). For example, `start_in_units` set to 20 and `end_in_units` set to 60 using `unit_timescale_in_miliseconds` in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
        /// </summary>
        [Input("endInUnits")]
        public Input<int>? EndInUnits { get; set; }

        /// <summary>
        /// Indicates whether the `end_in_units` property must be present. If true, `end_in_units` must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
        /// </summary>
        [Input("forceEnd")]
        public Input<bool>? ForceEnd { get; set; }

        /// <summary>
        /// The relative to end right edge. Applies to Live Streaming only.
        /// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by `unit_timescale_in_miliseconds`. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
        /// </summary>
        [Input("liveBackoffInUnits")]
        public Input<int>? LiveBackoffInUnits { get; set; }

        /// <summary>
        /// The relative to end sliding window. Applies to Live Streaming only. Use `presentation_window_in_units` to apply a sliding window of fragments to include in a playlist. The unit is defined by `unit_timescale_in_miliseconds`. For example, set  `presentation_window_in_units` to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
        /// </summary>
        [Input("presentationWindowInUnits")]
        public Input<int>? PresentationWindowInUnits { get; set; }

        /// <summary>
        /// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by `unit_timescale_in_miliseconds`, so a `start_in_units` of 15 would be for 15 seconds. Use `start_in_units` and `end_in_units` to trim the fragments that will be in the playlist (manifest). For example, `start_in_units` set to 20 and `end_in_units` set to 60 using `unit_timescale_in_miliseconds` in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
        /// </summary>
        [Input("startInUnits")]
        public Input<int>? StartInUnits { get; set; }

        /// <summary>
        /// Specified as the number of miliseconds in one unit timescale. For example, if you want to set a `start_in_units` at 30 seconds, you would use a value of 30 when using the `unit_timescale_in_miliseconds` in 1000. Or if you want to set `start_in_units` in 30 miliseconds, you would use a value of 30 when using the `unit_timescale_in_miliseconds` in 1.  Applies timescale to `start_in_units`, `start_timescale` and `presentation_window_in_timescale` and `live_backoff_in_timescale`.
        /// </summary>
        [Input("unitTimescaleInMiliseconds")]
        public Input<int>? UnitTimescaleInMiliseconds { get; set; }

        public AssetFilterPresentationTimeRangeGetArgs()
        {
        }
    }
}
