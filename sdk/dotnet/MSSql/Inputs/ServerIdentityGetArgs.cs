// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Inputs
{

    public sealed class ServerIdentityGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Specifies the identity type of the Microsoft SQL Server. At this time the only allowed value is `SystemAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ServerIdentityGetArgs()
        {
        }
    }
}
