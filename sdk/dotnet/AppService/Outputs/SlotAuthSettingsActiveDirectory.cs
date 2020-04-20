// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class SlotAuthSettingsActiveDirectory
    {
        /// <summary>
        /// Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
        /// </summary>
        public readonly ImmutableArray<string> AllowedAudiences;
        /// <summary>
        /// The Client ID of this relying party application. Enables OpenIDConnection authentication with Azure Active Directory.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The Client Secret of this relying party application. If no secret is provided, implicit flow will be used.
        /// </summary>
        public readonly string? ClientSecret;

        [OutputConstructor]
        private SlotAuthSettingsActiveDirectory(
            ImmutableArray<string> allowedAudiences,

            string clientId,

            string? clientSecret)
        {
            AllowedAudiences = allowedAudiences;
            ClientId = clientId;
            ClientSecret = clientSecret;
        }
    }
}