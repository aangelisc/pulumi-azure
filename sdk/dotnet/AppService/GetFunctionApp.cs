// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static class GetFunctionApp
    {
        /// <summary>
        /// Use this data source to access information about a Function App.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.AppService.GetFunctionApp.InvokeAsync(new Azure.AppService.GetFunctionAppArgs
        ///         {
        ///             Name = "test-azure-functions",
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionAppResult> InvokeAsync(GetFunctionAppArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionAppResult>("azure:appservice/getFunctionApp:getFunctionApp", args ?? new GetFunctionAppArgs(), options.WithVersion());
    }


    public sealed class GetFunctionAppArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Function App resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Function App exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetFunctionAppArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionAppResult
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Function App.
        /// </summary>
        public readonly string AppServicePlanId;
        /// <summary>
        /// A key-value pair of App Settings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> AppSettings;
        /// <summary>
        /// The mode of the Function App's client certificates requirement for incoming requests.
        /// </summary>
        public readonly string ClientCertMode;
        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAppConnectionStringResult> ConnectionStrings;
        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        public readonly string CustomDomainVerificationId;
        /// <summary>
        /// The default hostname associated with the Function App.
        /// </summary>
        public readonly string DefaultHostname;
        /// <summary>
        /// Is the Function App enabled?
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAppIdentityResult> Identities;
        public readonly string Location;
        /// <summary>
        /// The name for this IP Restriction.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A string indicating the Operating System type for this function app.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// A comma separated list of outbound IP addresses.
        /// </summary>
        public readonly string OutboundIpAddresses;
        /// <summary>
        /// A comma separated list of outbound IP addresses, not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        /// </summary>
        public readonly string PossibleOutboundIpAddresses;
        public readonly string ResourceGroupName;
        public readonly ImmutableArray<Outputs.GetFunctionAppSiteConfigResult> SiteConfigs;
        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAppSiteCredentialResult> SiteCredentials;
        /// <summary>
        /// A `source_control` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAppSourceControlResult> SourceControls;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetFunctionAppResult(
            string appServicePlanId,

            ImmutableDictionary<string, string> appSettings,

            string clientCertMode,

            ImmutableArray<Outputs.GetFunctionAppConnectionStringResult> connectionStrings,

            string customDomainVerificationId,

            string defaultHostname,

            bool enabled,

            string id,

            ImmutableArray<Outputs.GetFunctionAppIdentityResult> identities,

            string location,

            string name,

            string osType,

            string outboundIpAddresses,

            string possibleOutboundIpAddresses,

            string resourceGroupName,

            ImmutableArray<Outputs.GetFunctionAppSiteConfigResult> siteConfigs,

            ImmutableArray<Outputs.GetFunctionAppSiteCredentialResult> siteCredentials,

            ImmutableArray<Outputs.GetFunctionAppSourceControlResult> sourceControls,

            ImmutableDictionary<string, string>? tags)
        {
            AppServicePlanId = appServicePlanId;
            AppSettings = appSettings;
            ClientCertMode = clientCertMode;
            ConnectionStrings = connectionStrings;
            CustomDomainVerificationId = customDomainVerificationId;
            DefaultHostname = defaultHostname;
            Enabled = enabled;
            Id = id;
            Identities = identities;
            Location = location;
            Name = name;
            OsType = osType;
            OutboundIpAddresses = outboundIpAddresses;
            PossibleOutboundIpAddresses = possibleOutboundIpAddresses;
            ResourceGroupName = resourceGroupName;
            SiteConfigs = siteConfigs;
            SiteCredentials = siteCredentials;
            SourceControls = sourceControls;
            Tags = tags;
        }
    }
}
