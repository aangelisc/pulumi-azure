// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform.Outputs
{

    [OutputType]
    public sealed class GetSpringCloudAppIdentityResult
    {
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The Type of Managed Identity assigned to the Spring Cloud Application.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSpringCloudAppIdentityResult(
            string principalId,

            string tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
