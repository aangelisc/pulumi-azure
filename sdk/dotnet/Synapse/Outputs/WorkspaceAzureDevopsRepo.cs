// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Synapse.Outputs
{

    [OutputType]
    public sealed class WorkspaceAzureDevopsRepo
    {
        /// <summary>
        /// Specifies the Azure DevOps account name.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// Specifies the collaboration branch of the repository to get code from.
        /// </summary>
        public readonly string BranchName;
        /// <summary>
        /// Specifies the name of the Azure DevOps project.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// Specifies the name of the git repository.
        /// </summary>
        public readonly string RepositoryName;
        /// <summary>
        /// Specifies the root folder within the repository. Set to `/` for the top level.
        /// </summary>
        public readonly string RootFolder;

        [OutputConstructor]
        private WorkspaceAzureDevopsRepo(
            string accountName,

            string branchName,

            string projectName,

            string repositoryName,

            string rootFolder)
        {
            AccountName = accountName;
            BranchName = branchName;
            ProjectName = projectName;
            RepositoryName = repositoryName;
            RootFolder = rootFolder;
        }
    }
}
