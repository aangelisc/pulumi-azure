// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault.Inputs
{

    public sealed class CertificateCertificatePolicyKeyPropertiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the curve to use when creating an `EC` key. Possible values are `P-256`, `P-256K`, `P-384`, and `P-521`. This field will be required in a future release if `key_type` is `EC` or `EC-HSM`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("curve")]
        public Input<string>? Curve { get; set; }

        /// <summary>
        /// Is this certificate exportable? Changing this forces a new resource to be created.
        /// </summary>
        [Input("exportable", required: true)]
        public Input<bool> Exportable { get; set; } = null!;

        /// <summary>
        /// The size of the key used in the certificate. Possible values include `2048`, `3072`, and `4096` for `RSA` keys, or `256`, `384`, and `521` for `EC` keys. This property is required when using RSA keys. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// Specifies the type of key, such as `RSA` or `EC`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        /// <summary>
        /// Is the key reusable? Changing this forces a new resource to be created.
        /// </summary>
        [Input("reuseKey", required: true)]
        public Input<bool> ReuseKey { get; set; } = null!;

        public CertificateCertificatePolicyKeyPropertiesGetArgs()
        {
        }
    }
}
