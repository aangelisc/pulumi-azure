// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hpc.Inputs
{

    public sealed class CacheDirectoryActiveDirectoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server.
        /// </summary>
        [Input("cacheNetbiosName", required: true)]
        public Input<string> CacheNetbiosName { get; set; } = null!;

        /// <summary>
        /// The primary DNS IP address used to resolve the Active Directory domain controller's FQDN.
        /// </summary>
        [Input("dnsPrimaryIp", required: true)]
        public Input<string> DnsPrimaryIp { get; set; } = null!;

        /// <summary>
        /// The secondary DNS IP address used to resolve the Active Directory domain controller's FQDN.
        /// </summary>
        [Input("dnsSecondaryIp")]
        public Input<string>? DnsSecondaryIp { get; set; }

        /// <summary>
        /// The fully qualified domain name of the Active Directory domain controller.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The Active Directory domain's NetBIOS name.
        /// </summary>
        [Input("domainNetbiosName", required: true)]
        public Input<string> DomainNetbiosName { get; set; } = null!;

        /// <summary>
        /// The password of the Active Directory domain administrator.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username of the Active Directory domain administrator.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public CacheDirectoryActiveDirectoryArgs()
        {
        }
    }
}
