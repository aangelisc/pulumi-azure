// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage.Outputs
{

    [OutputType]
    public sealed class AccountRouting
    {
        /// <summary>
        /// Specifies the kind of network routing opted by the user. Possible values are `InternetRouting` and `MicrosoftRouting`. Defaults to `MicrosoftRouting`.
        /// </summary>
        public readonly string? Choice;
        /// <summary>
        /// Should internet routing storage endpoints be published? Defaults to `false`.
        /// </summary>
        public readonly bool? PublishInternetEndpoints;
        /// <summary>
        /// Should microsoft routing storage endpoints be published? Defaults to `false`.
        /// </summary>
        public readonly bool? PublishMicrosoftEndpoints;

        [OutputConstructor]
        private AccountRouting(
            string? choice,

            bool? publishInternetEndpoints,

            bool? publishMicrosoftEndpoints)
        {
            Choice = choice;
            PublishInternetEndpoints = publishInternetEndpoints;
            PublishMicrosoftEndpoints = publishMicrosoftEndpoints;
        }
    }
}