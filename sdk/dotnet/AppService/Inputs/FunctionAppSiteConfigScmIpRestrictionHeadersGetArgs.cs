// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class FunctionAppSiteConfigScmIpRestrictionHeadersGetArgs : Pulumi.ResourceArgs
    {
        [Input("xAzureFdids")]
        private InputList<string>? _xAzureFdids;

        /// <summary>
        /// A list of allowed Azure FrontDoor IDs in UUID notation with a maximum of 8.
        /// </summary>
        public InputList<string> XAzureFdids
        {
            get => _xAzureFdids ?? (_xAzureFdids = new InputList<string>());
            set => _xAzureFdids = value;
        }

        /// <summary>
        /// A list to allow the Azure FrontDoor health probe header. Only allowed value is "1".
        /// </summary>
        [Input("xFdHealthProbe")]
        public Input<string>? XFdHealthProbe { get; set; }

        [Input("xForwardedFors")]
        private InputList<string>? _xForwardedFors;

        /// <summary>
        /// A list of allowed 'X-Forwarded-For' IPs in CIDR notation with a maximum of 8
        /// </summary>
        public InputList<string> XForwardedFors
        {
            get => _xForwardedFors ?? (_xForwardedFors = new InputList<string>());
            set => _xForwardedFors = value;
        }

        [Input("xForwardedHosts")]
        private InputList<string>? _xForwardedHosts;

        /// <summary>
        /// A list of allowed 'X-Forwarded-Host' domains with a maximum of 8.
        /// </summary>
        public InputList<string> XForwardedHosts
        {
            get => _xForwardedHosts ?? (_xForwardedHosts = new InputList<string>());
            set => _xForwardedHosts = value;
        }

        public FunctionAppSiteConfigScmIpRestrictionHeadersGetArgs()
        {
        }
    }
}
