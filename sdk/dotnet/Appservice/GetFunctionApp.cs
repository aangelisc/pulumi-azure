// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about a Function App.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/function_app.html.markdown.
        /// </summary>
        public static Task<GetFunctionAppResult> GetFunctionApp(GetFunctionAppArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionAppResult>("azure:appservice/getFunctionApp:getFunctionApp", args ?? InvokeArgs.Empty, options.WithVersion());
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
        /// An `connection_string` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAppConnectionStringsResult> ConnectionStrings;
        /// <summary>
        /// The default hostname associated with the Function App.
        /// </summary>
        public readonly string DefaultHostname;
        /// <summary>
        /// Is the Function App enabled?
        /// </summary>
        public readonly bool Enabled;
        public readonly string Location;
        /// <summary>
        /// The name of the Connection String.
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
        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAppSiteCredentialsResult> SiteCredentials;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFunctionAppResult(
            string appServicePlanId,
            ImmutableDictionary<string, string> appSettings,
            ImmutableArray<Outputs.GetFunctionAppConnectionStringsResult> connectionStrings,
            string defaultHostname,
            bool enabled,
            string location,
            string name,
            string osType,
            string outboundIpAddresses,
            string possibleOutboundIpAddresses,
            string resourceGroupName,
            ImmutableArray<Outputs.GetFunctionAppSiteCredentialsResult> siteCredentials,
            ImmutableDictionary<string, string>? tags,
            string id)
        {
            AppServicePlanId = appServicePlanId;
            AppSettings = appSettings;
            ConnectionStrings = connectionStrings;
            DefaultHostname = defaultHostname;
            Enabled = enabled;
            Location = location;
            Name = name;
            OsType = osType;
            OutboundIpAddresses = outboundIpAddresses;
            PossibleOutboundIpAddresses = possibleOutboundIpAddresses;
            ResourceGroupName = resourceGroupName;
            SiteCredentials = siteCredentials;
            Tags = tags;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetFunctionAppConnectionStringsResult
    {
        /// <summary>
        /// The name of the Function App resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the Connection String. 
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value for the Connection String.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetFunctionAppConnectionStringsResult(
            string name,
            string type,
            string value)
        {
            Name = name;
            Type = type;
            Value = value;
        }
    }

    [OutputType]
    public sealed class GetFunctionAppSiteCredentialsResult
    {
        /// <summary>
        /// The password associated with the username, which can be used to publish to this App Service.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The username which can be used to publish to this App Service
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetFunctionAppSiteCredentialsResult(
            string password,
            string username)
        {
            Password = password;
            Username = username;
        }
    }
    }
}
