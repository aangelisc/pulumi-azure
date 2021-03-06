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
    public sealed class MeshApplicationServiceCodePackage
    {
        /// <summary>
        /// The Container image the code package will use.
        /// </summary>
        public readonly string ImageName;
        /// <summary>
        /// The name of the code package.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A `resources` block as defined below.
        /// </summary>
        public readonly Outputs.MeshApplicationServiceCodePackageResources Resources;

        [OutputConstructor]
        private MeshApplicationServiceCodePackage(
            string imageName,

            string name,

            Outputs.MeshApplicationServiceCodePackageResources resources)
        {
            ImageName = imageName;
            Name = name;
            Resources = resources;
        }
    }
}
