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
    public sealed class AccountBlobProperties
    {
        /// <summary>
        /// Is the blob service properties for change feed events enabled? Default to `false`.
        /// </summary>
        public readonly bool? ChangeFeedEnabled;
        /// <summary>
        /// A `container_delete_retention_policy` block as defined below.
        /// </summary>
        public readonly Outputs.AccountBlobPropertiesContainerDeleteRetentionPolicy? ContainerDeleteRetentionPolicy;
        /// <summary>
        /// A `cors_rule` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccountBlobPropertiesCorsRule> CorsRules;
        /// <summary>
        /// The API Version which should be used by default for requests to the Data Plane API if an incoming request doesn't specify an API Version. Defaults to `2020-06-12`.
        /// </summary>
        public readonly string? DefaultServiceVersion;
        /// <summary>
        /// A `delete_retention_policy` block as defined below.
        /// </summary>
        public readonly Outputs.AccountBlobPropertiesDeleteRetentionPolicy? DeleteRetentionPolicy;
        /// <summary>
        /// Is the last access time based tracking enabled? Default to `false`.
        /// </summary>
        public readonly bool? LastAccessTimeEnabled;
        /// <summary>
        /// Is versioning enabled? Default to `false`.
        /// </summary>
        public readonly bool? VersioningEnabled;

        [OutputConstructor]
        private AccountBlobProperties(
            bool? changeFeedEnabled,

            Outputs.AccountBlobPropertiesContainerDeleteRetentionPolicy? containerDeleteRetentionPolicy,

            ImmutableArray<Outputs.AccountBlobPropertiesCorsRule> corsRules,

            string? defaultServiceVersion,

            Outputs.AccountBlobPropertiesDeleteRetentionPolicy? deleteRetentionPolicy,

            bool? lastAccessTimeEnabled,

            bool? versioningEnabled)
        {
            ChangeFeedEnabled = changeFeedEnabled;
            ContainerDeleteRetentionPolicy = containerDeleteRetentionPolicy;
            CorsRules = corsRules;
            DefaultServiceVersion = defaultServiceVersion;
            DeleteRetentionPolicy = deleteRetentionPolicy;
            LastAccessTimeEnabled = lastAccessTimeEnabled;
            VersioningEnabled = versioningEnabled;
        }
    }
}
