// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class ScaleSetIdentity
    {
        /// <summary>
        /// Specifies a list of user managed identity ids to be assigned to the VMSS. Required if `type` is `UserAssigned`.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        public readonly string? PrincipalId;
        /// <summary>
        /// Specifies the identity type to be assigned to the scale set. Allowable values are `SystemAssigned` and `UserAssigned`. For the `SystemAssigned` identity the scale set's Service Principal ID (SPN) can be retrieved after the scale set has been created. See [documentation](https://docs.microsoft.com/en-us/azure/active-directory/managed-service-identity/overview) for more information.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ScaleSetIdentity(
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