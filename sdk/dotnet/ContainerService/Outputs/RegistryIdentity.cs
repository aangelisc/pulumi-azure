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
    public sealed class RegistryIdentity
    {
        /// <summary>
        /// A list of User Managed Identity ID's which should be assigned to the Container Registry.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        public readonly string? PrincipalId;
        /// <summary>
        /// The type of Managed Identity which should be assigned to the Container Registry. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RegistryIdentity(
            ImmutableArray<string> identityIds,

            string? principalId,

            string type)
        {
            IdentityIds = identityIds;
            PrincipalId = principalId;
            Type = type;
        }
    }
}
