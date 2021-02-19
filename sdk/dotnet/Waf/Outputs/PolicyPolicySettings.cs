// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Waf.Outputs
{

    [OutputType]
    public sealed class PolicyPolicySettings
    {
        /// <summary>
        /// Describes if the policy is in enabled state or disabled state. Defaults to `Enabled`.
        /// </summary>
        public readonly bool? Enabled;
        public readonly int? FileUploadLimitInMb;
        /// <summary>
        /// The Maximum Request Body Size in KB.  Accepted values are in the range `8` to `128`. Defaults to `128`.
        /// </summary>
        public readonly int? MaxRequestBodySizeInKb;
        /// <summary>
        /// Describes if it is in detection mode or prevention mode at the policy level. Defaults to `Prevention`.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// Is Request Body Inspection enabled? Defaults to `true`.
        /// </summary>
        public readonly bool? RequestBodyCheck;

        [OutputConstructor]
        private PolicyPolicySettings(
            bool? enabled,

            int? fileUploadLimitInMb,

            int? maxRequestBodySizeInKb,

            string? mode,

            bool? requestBodyCheck)
        {
            Enabled = enabled;
            FileUploadLimitInMb = fileUploadLimitInMb;
            MaxRequestBodySizeInKb = maxRequestBodySizeInKb;
            Mode = mode;
            RequestBodyCheck = requestBodyCheck;
        }
    }
}
