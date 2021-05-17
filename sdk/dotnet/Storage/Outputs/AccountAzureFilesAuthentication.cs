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
    public sealed class AccountAzureFilesAuthentication
    {
        /// <summary>
        /// A `active_directory` block as defined below. Required when `directory_type` is `AD`.
        /// </summary>
        public readonly Outputs.AccountAzureFilesAuthenticationActiveDirectory? ActiveDirectory;
        /// <summary>
        /// Specifies the directory service used. Possible values are `AADDS` and `AD`.
        /// </summary>
        public readonly string DirectoryType;

        [OutputConstructor]
        private AccountAzureFilesAuthentication(
            Outputs.AccountAzureFilesAuthenticationActiveDirectory? activeDirectory,

            string directoryType)
        {
            ActiveDirectory = activeDirectory;
            DirectoryType = directoryType;
        }
    }
}
