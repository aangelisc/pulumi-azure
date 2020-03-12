// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages a backend within an API Management Service.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_backend.html.markdown.
    /// </summary>
    public partial class Backend : Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// A `credentials` block as documented below.
        /// </summary>
        [Output("credentials")]
        public Output<Outputs.BackendCredentials?> Credentials { get; private set; } = null!;

        /// <summary>
        /// The description of the backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the API Management backend. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The protocol used by the backend host. Possible values are `http` or `soap`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// A `proxy` block as documented below.
        /// </summary>
        [Output("proxy")]
        public Output<Outputs.BackendProxy?> Proxy { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;

        /// <summary>
        /// A `service_fabric_cluster` block as documented below.
        /// </summary>
        [Output("serviceFabricCluster")]
        public Output<Outputs.BackendServiceFabricCluster?> ServiceFabricCluster { get; private set; } = null!;

        /// <summary>
        /// The title of the backend.
        /// </summary>
        [Output("title")]
        public Output<string?> Title { get; private set; } = null!;

        /// <summary>
        /// A `tls` block as documented below.
        /// </summary>
        [Output("tls")]
        public Output<Outputs.BackendTls?> Tls { get; private set; } = null!;

        /// <summary>
        /// The URL of the backend host.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Backend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backend(string name, BackendArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/backend:Backend", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Backend(string name, Input<string> id, BackendState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/backend:Backend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backend Get(string name, Input<string> id, BackendState? state = null, CustomResourceOptions? options = null)
        {
            return new Backend(name, id, state, options);
        }
    }

    public sealed class BackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// A `credentials` block as documented below.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.BackendCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// The description of the backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the API Management backend. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol used by the backend host. Possible values are `http` or `soap`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// A `proxy` block as documented below.
        /// </summary>
        [Input("proxy")]
        public Input<Inputs.BackendProxyArgs>? Proxy { get; set; }

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// A `service_fabric_cluster` block as documented below.
        /// </summary>
        [Input("serviceFabricCluster")]
        public Input<Inputs.BackendServiceFabricClusterArgs>? ServiceFabricCluster { get; set; }

        /// <summary>
        /// The title of the backend.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// A `tls` block as documented below.
        /// </summary>
        [Input("tls")]
        public Input<Inputs.BackendTlsArgs>? Tls { get; set; }

        /// <summary>
        /// The URL of the backend host.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public BackendArgs()
        {
        }
    }

    public sealed class BackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// A `credentials` block as documented below.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.BackendCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// The description of the backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the API Management backend. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol used by the backend host. Possible values are `http` or `soap`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// A `proxy` block as documented below.
        /// </summary>
        [Input("proxy")]
        public Input<Inputs.BackendProxyGetArgs>? Proxy { get; set; }

        /// <summary>
        /// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// A `service_fabric_cluster` block as documented below.
        /// </summary>
        [Input("serviceFabricCluster")]
        public Input<Inputs.BackendServiceFabricClusterGetArgs>? ServiceFabricCluster { get; set; }

        /// <summary>
        /// The title of the backend.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// A `tls` block as documented below.
        /// </summary>
        [Input("tls")]
        public Input<Inputs.BackendTlsGetArgs>? Tls { get; set; }

        /// <summary>
        /// The URL of the backend host.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public BackendState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class BackendCredentialsArgs : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        public Input<BackendCredentialsAuthorizationArgs>? Authorization { get; set; }

        [Input("certificates")]
        private InputList<string>? _certificates;
        public InputList<string> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<string>());
            set => _certificates = value;
        }

        [Input("header")]
        private InputMap<string>? _header;
        public InputMap<string> Header
        {
            get => _header ?? (_header = new InputMap<string>());
            set => _header = value;
        }

        [Input("query")]
        private InputMap<string>? _query;
        public InputMap<string> Query
        {
            get => _query ?? (_query = new InputMap<string>());
            set => _query = value;
        }

        public BackendCredentialsArgs()
        {
        }
    }

    public sealed class BackendCredentialsAuthorizationArgs : Pulumi.ResourceArgs
    {
        [Input("parameter")]
        public Input<string>? Parameter { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        public BackendCredentialsAuthorizationArgs()
        {
        }
    }

    public sealed class BackendCredentialsAuthorizationGetArgs : Pulumi.ResourceArgs
    {
        [Input("parameter")]
        public Input<string>? Parameter { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        public BackendCredentialsAuthorizationGetArgs()
        {
        }
    }

    public sealed class BackendCredentialsGetArgs : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        public Input<BackendCredentialsAuthorizationGetArgs>? Authorization { get; set; }

        [Input("certificates")]
        private InputList<string>? _certificates;
        public InputList<string> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<string>());
            set => _certificates = value;
        }

        [Input("header")]
        private InputMap<string>? _header;
        public InputMap<string> Header
        {
            get => _header ?? (_header = new InputMap<string>());
            set => _header = value;
        }

        [Input("query")]
        private InputMap<string>? _query;
        public InputMap<string> Query
        {
            get => _query ?? (_query = new InputMap<string>());
            set => _query = value;
        }

        public BackendCredentialsGetArgs()
        {
        }
    }

    public sealed class BackendProxyArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The URL of the backend host.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BackendProxyArgs()
        {
        }
    }

    public sealed class BackendProxyGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The URL of the backend host.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BackendProxyGetArgs()
        {
        }
    }

    public sealed class BackendServiceFabricClusterArgs : Pulumi.ResourceArgs
    {
        [Input("clientCertificateThumbprint", required: true)]
        public Input<string> ClientCertificateThumbprint { get; set; } = null!;

        [Input("managementEndpoints", required: true)]
        private InputList<string>? _managementEndpoints;
        public InputList<string> ManagementEndpoints
        {
            get => _managementEndpoints ?? (_managementEndpoints = new InputList<string>());
            set => _managementEndpoints = value;
        }

        [Input("maxPartitionResolutionRetries", required: true)]
        public Input<int> MaxPartitionResolutionRetries { get; set; } = null!;

        [Input("serverCertificateThumbprints")]
        private InputList<string>? _serverCertificateThumbprints;
        public InputList<string> ServerCertificateThumbprints
        {
            get => _serverCertificateThumbprints ?? (_serverCertificateThumbprints = new InputList<string>());
            set => _serverCertificateThumbprints = value;
        }

        [Input("serverX509Names")]
        private InputList<BackendServiceFabricClusterServerX509NamesArgs>? _serverX509Names;
        public InputList<BackendServiceFabricClusterServerX509NamesArgs> ServerX509Names
        {
            get => _serverX509Names ?? (_serverX509Names = new InputList<BackendServiceFabricClusterServerX509NamesArgs>());
            set => _serverX509Names = value;
        }

        public BackendServiceFabricClusterArgs()
        {
        }
    }

    public sealed class BackendServiceFabricClusterGetArgs : Pulumi.ResourceArgs
    {
        [Input("clientCertificateThumbprint", required: true)]
        public Input<string> ClientCertificateThumbprint { get; set; } = null!;

        [Input("managementEndpoints", required: true)]
        private InputList<string>? _managementEndpoints;
        public InputList<string> ManagementEndpoints
        {
            get => _managementEndpoints ?? (_managementEndpoints = new InputList<string>());
            set => _managementEndpoints = value;
        }

        [Input("maxPartitionResolutionRetries", required: true)]
        public Input<int> MaxPartitionResolutionRetries { get; set; } = null!;

        [Input("serverCertificateThumbprints")]
        private InputList<string>? _serverCertificateThumbprints;
        public InputList<string> ServerCertificateThumbprints
        {
            get => _serverCertificateThumbprints ?? (_serverCertificateThumbprints = new InputList<string>());
            set => _serverCertificateThumbprints = value;
        }

        [Input("serverX509Names")]
        private InputList<BackendServiceFabricClusterServerX509NamesGetArgs>? _serverX509Names;
        public InputList<BackendServiceFabricClusterServerX509NamesGetArgs> ServerX509Names
        {
            get => _serverX509Names ?? (_serverX509Names = new InputList<BackendServiceFabricClusterServerX509NamesGetArgs>());
            set => _serverX509Names = value;
        }

        public BackendServiceFabricClusterGetArgs()
        {
        }
    }

    public sealed class BackendServiceFabricClusterServerX509NamesArgs : Pulumi.ResourceArgs
    {
        [Input("issuerCertificateThumbprint", required: true)]
        public Input<string> IssuerCertificateThumbprint { get; set; } = null!;

        /// <summary>
        /// The name of the API Management backend. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public BackendServiceFabricClusterServerX509NamesArgs()
        {
        }
    }

    public sealed class BackendServiceFabricClusterServerX509NamesGetArgs : Pulumi.ResourceArgs
    {
        [Input("issuerCertificateThumbprint", required: true)]
        public Input<string> IssuerCertificateThumbprint { get; set; } = null!;

        /// <summary>
        /// The name of the API Management backend. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public BackendServiceFabricClusterServerX509NamesGetArgs()
        {
        }
    }

    public sealed class BackendTlsArgs : Pulumi.ResourceArgs
    {
        [Input("validateCertificateChain")]
        public Input<bool>? ValidateCertificateChain { get; set; }

        [Input("validateCertificateName")]
        public Input<bool>? ValidateCertificateName { get; set; }

        public BackendTlsArgs()
        {
        }
    }

    public sealed class BackendTlsGetArgs : Pulumi.ResourceArgs
    {
        [Input("validateCertificateChain")]
        public Input<bool>? ValidateCertificateChain { get; set; }

        [Input("validateCertificateName")]
        public Input<bool>? ValidateCertificateName { get; set; }

        public BackendTlsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class BackendCredentials
    {
        public readonly BackendCredentialsAuthorization? Authorization;
        public readonly ImmutableArray<string> Certificates;
        public readonly ImmutableDictionary<string, string>? Header;
        public readonly ImmutableDictionary<string, string>? Query;

        [OutputConstructor]
        private BackendCredentials(
            BackendCredentialsAuthorization? authorization,
            ImmutableArray<string> certificates,
            ImmutableDictionary<string, string>? header,
            ImmutableDictionary<string, string>? query)
        {
            Authorization = authorization;
            Certificates = certificates;
            Header = header;
            Query = query;
        }
    }

    [OutputType]
    public sealed class BackendCredentialsAuthorization
    {
        public readonly string? Parameter;
        public readonly string? Scheme;

        [OutputConstructor]
        private BackendCredentialsAuthorization(
            string? parameter,
            string? scheme)
        {
            Parameter = parameter;
            Scheme = scheme;
        }
    }

    [OutputType]
    public sealed class BackendProxy
    {
        public readonly string? Password;
        /// <summary>
        /// The URL of the backend host.
        /// </summary>
        public readonly string Url;
        public readonly string Username;

        [OutputConstructor]
        private BackendProxy(
            string? password,
            string url,
            string username)
        {
            Password = password;
            Url = url;
            Username = username;
        }
    }

    [OutputType]
    public sealed class BackendServiceFabricCluster
    {
        public readonly string ClientCertificateThumbprint;
        public readonly ImmutableArray<string> ManagementEndpoints;
        public readonly int MaxPartitionResolutionRetries;
        public readonly ImmutableArray<string> ServerCertificateThumbprints;
        public readonly ImmutableArray<BackendServiceFabricClusterServerX509Names> ServerX509Names;

        [OutputConstructor]
        private BackendServiceFabricCluster(
            string clientCertificateThumbprint,
            ImmutableArray<string> managementEndpoints,
            int maxPartitionResolutionRetries,
            ImmutableArray<string> serverCertificateThumbprints,
            ImmutableArray<BackendServiceFabricClusterServerX509Names> serverX509Names)
        {
            ClientCertificateThumbprint = clientCertificateThumbprint;
            ManagementEndpoints = managementEndpoints;
            MaxPartitionResolutionRetries = maxPartitionResolutionRetries;
            ServerCertificateThumbprints = serverCertificateThumbprints;
            ServerX509Names = serverX509Names;
        }
    }

    [OutputType]
    public sealed class BackendServiceFabricClusterServerX509Names
    {
        public readonly string IssuerCertificateThumbprint;
        /// <summary>
        /// The name of the API Management backend. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private BackendServiceFabricClusterServerX509Names(
            string issuerCertificateThumbprint,
            string name)
        {
            IssuerCertificateThumbprint = issuerCertificateThumbprint;
            Name = name;
        }
    }

    [OutputType]
    public sealed class BackendTls
    {
        public readonly bool? ValidateCertificateChain;
        public readonly bool? ValidateCertificateName;

        [OutputConstructor]
        private BackendTls(
            bool? validateCertificateChain,
            bool? validateCertificateName)
        {
            ValidateCertificateChain = validateCertificateChain;
            ValidateCertificateName = validateCertificateName;
        }
    }
    }
}
