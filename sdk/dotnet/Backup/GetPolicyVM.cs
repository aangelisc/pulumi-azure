// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing VM Backup Policy.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/backup_policy_vm.markdown.
        /// </summary>
        public static Task<GetPolicyVMResult> GetPolicyVM(GetPolicyVMArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyVMResult>("azure:backup/getPolicyVM:getPolicyVM", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetPolicyVMArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the VM Backup Policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public string RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the VM Backup Policy resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPolicyVMArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPolicyVMResult
    {
        public readonly string Name;
        public readonly string RecoveryVaultName;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPolicyVMResult(
            string name,
            string recoveryVaultName,
            string resourceGroupName,
            ImmutableDictionary<string, string> tags,
            string id)
        {
            Name = name;
            RecoveryVaultName = recoveryVaultName;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            Id = id;
        }
    }
}
