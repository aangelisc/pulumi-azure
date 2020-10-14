// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Outputs
{

    [OutputType]
    public sealed class MeshApplicationService
    {
        /// <summary>
        /// Any number `code_package` block as described below.
        /// </summary>
        public readonly ImmutableArray<Outputs.MeshApplicationServiceCodePackage> CodePackages;
        /// <summary>
        /// The name of the service resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The operating system required by the code in service. Valid values are `Linux` or `Windows`.
        /// </summary>
        public readonly string OsType;

        [OutputConstructor]
        private MeshApplicationService(
            ImmutableArray<Outputs.MeshApplicationServiceCodePackage> codePackages,

            string name,

            string osType)
        {
            CodePackages = codePackages;
            Name = name;
            OsType = osType;
        }
    }
}
