// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ExpressRoutePortLink2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether enable administration state on the Express Route Port Link? Defaults to `false`.
        /// </summary>
        [Input("adminEnabled")]
        public Input<bool>? AdminEnabled { get; set; }

        /// <summary>
        /// The connector type of the Express Route Port Link.
        /// </summary>
        [Input("connectorType")]
        public Input<string>? ConnectorType { get; set; }

        /// <summary>
        /// The ID of this Express Route Port Link.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The interface name of the Azure router associated with the Express Route Port Link.
        /// </summary>
        [Input("interfaceName")]
        public Input<string>? InterfaceName { get; set; }

        /// <summary>
        /// The ID of the Key Vault Secret that contains the Mac security CAK key for this Express Route Port Link.
        /// </summary>
        [Input("macsecCakKeyvaultSecretId")]
        public Input<string>? MacsecCakKeyvaultSecretId { get; set; }

        /// <summary>
        /// The MACSec cipher used for this Express Route Port Link. Possible values are `GcmAes128` and `GcmAes256`. Defaults to `GcmAes128`.
        /// </summary>
        [Input("macsecCipher")]
        public Input<string>? MacsecCipher { get; set; }

        /// <summary>
        /// The ID of the Key Vault Secret that contains the MACSec CKN key for this Express Route Port Link.
        /// </summary>
        [Input("macsecCknKeyvaultSecretId")]
        public Input<string>? MacsecCknKeyvaultSecretId { get; set; }

        /// <summary>
        /// The ID that maps from the Express Route Port Link to the patch panel port.
        /// </summary>
        [Input("patchPanelId")]
        public Input<string>? PatchPanelId { get; set; }

        /// <summary>
        /// The ID that maps from the patch panel port to the rack.
        /// </summary>
        [Input("rackId")]
        public Input<string>? RackId { get; set; }

        /// <summary>
        /// The name of the Azure router associated with the Express Route Port Link.
        /// </summary>
        [Input("routerName")]
        public Input<string>? RouterName { get; set; }

        public ExpressRoutePortLink2Args()
        {
        }
    }
}
