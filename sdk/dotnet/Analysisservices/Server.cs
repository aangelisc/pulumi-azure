// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AnalysisServices
{
    /// <summary>
    /// Manages an Analysis Services Server.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/analysis_services_server.html.markdown.
    /// </summary>
    public partial class Server : Pulumi.CustomResource
    {
        /// <summary>
        /// List of email addresses of admin users.
        /// </summary>
        [Output("adminUsers")]
        public Output<ImmutableArray<string>> AdminUsers { get; private set; } = null!;

        /// <summary>
        /// URI and SAS token for a blob container to store backups.
        /// </summary>
        [Output("backupBlobContainerUri")]
        public Output<string?> BackupBlobContainerUri { get; private set; } = null!;

        /// <summary>
        /// Indicates if the Power BI service is allowed to access or not.
        /// </summary>
        [Output("enablePowerBiService")]
        public Output<bool?> EnablePowerBiService { get; private set; } = null!;

        /// <summary>
        /// One or more `ipv4_firewall_rule` block(s) as defined below.
        /// </summary>
        [Output("ipv4FirewallRules")]
        public Output<ImmutableArray<Outputs.ServerIpv4FirewallRules>> Ipv4FirewallRules { get; private set; } = null!;

        /// <summary>
        /// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Controls how the read-write server is used in the query pool. If this values is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
        /// </summary>
        [Output("querypoolConnectionMode")]
        public Output<string> QuerypoolConnectionMode { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The full name of the Analysis Services Server.
        /// </summary>
        [Output("serverFullName")]
        public Output<string> ServerFullName { get; private set; } = null!;

        /// <summary>
        /// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8` and `S9`
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs args, CustomResourceOptions? options = null)
            : base("azure:analysisservices/server:Server", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
            : base("azure:analysisservices/server:Server", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Server(name, id, state, options);
        }
    }

    public sealed class ServerArgs : Pulumi.ResourceArgs
    {
        [Input("adminUsers")]
        private InputList<string>? _adminUsers;

        /// <summary>
        /// List of email addresses of admin users.
        /// </summary>
        public InputList<string> AdminUsers
        {
            get => _adminUsers ?? (_adminUsers = new InputList<string>());
            set => _adminUsers = value;
        }

        /// <summary>
        /// URI and SAS token for a blob container to store backups.
        /// </summary>
        [Input("backupBlobContainerUri")]
        public Input<string>? BackupBlobContainerUri { get; set; }

        /// <summary>
        /// Indicates if the Power BI service is allowed to access or not.
        /// </summary>
        [Input("enablePowerBiService")]
        public Input<bool>? EnablePowerBiService { get; set; }

        [Input("ipv4FirewallRules")]
        private InputList<Inputs.ServerIpv4FirewallRulesArgs>? _ipv4FirewallRules;

        /// <summary>
        /// One or more `ipv4_firewall_rule` block(s) as defined below.
        /// </summary>
        public InputList<Inputs.ServerIpv4FirewallRulesArgs> Ipv4FirewallRules
        {
            get => _ipv4FirewallRules ?? (_ipv4FirewallRules = new InputList<Inputs.ServerIpv4FirewallRulesArgs>());
            set => _ipv4FirewallRules = value;
        }

        /// <summary>
        /// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Controls how the read-write server is used in the query pool. If this values is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
        /// </summary>
        [Input("querypoolConnectionMode")]
        public Input<string>? QuerypoolConnectionMode { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8` and `S9`
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerArgs()
        {
        }
    }

    public sealed class ServerState : Pulumi.ResourceArgs
    {
        [Input("adminUsers")]
        private InputList<string>? _adminUsers;

        /// <summary>
        /// List of email addresses of admin users.
        /// </summary>
        public InputList<string> AdminUsers
        {
            get => _adminUsers ?? (_adminUsers = new InputList<string>());
            set => _adminUsers = value;
        }

        /// <summary>
        /// URI and SAS token for a blob container to store backups.
        /// </summary>
        [Input("backupBlobContainerUri")]
        public Input<string>? BackupBlobContainerUri { get; set; }

        /// <summary>
        /// Indicates if the Power BI service is allowed to access or not.
        /// </summary>
        [Input("enablePowerBiService")]
        public Input<bool>? EnablePowerBiService { get; set; }

        [Input("ipv4FirewallRules")]
        private InputList<Inputs.ServerIpv4FirewallRulesGetArgs>? _ipv4FirewallRules;

        /// <summary>
        /// One or more `ipv4_firewall_rule` block(s) as defined below.
        /// </summary>
        public InputList<Inputs.ServerIpv4FirewallRulesGetArgs> Ipv4FirewallRules
        {
            get => _ipv4FirewallRules ?? (_ipv4FirewallRules = new InputList<Inputs.ServerIpv4FirewallRulesGetArgs>());
            set => _ipv4FirewallRules = value;
        }

        /// <summary>
        /// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Controls how the read-write server is used in the query pool. If this values is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
        /// </summary>
        [Input("querypoolConnectionMode")]
        public Input<string>? QuerypoolConnectionMode { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The full name of the Analysis Services Server.
        /// </summary>
        [Input("serverFullName")]
        public Input<string>? ServerFullName { get; set; }

        /// <summary>
        /// SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8` and `S9`
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ServerIpv4FirewallRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// End of the firewall rule range as IPv4 address.
        /// </summary>
        [Input("rangeEnd", required: true)]
        public Input<string> RangeEnd { get; set; } = null!;

        /// <summary>
        /// Start of the firewall rule range as IPv4 address.
        /// </summary>
        [Input("rangeStart", required: true)]
        public Input<string> RangeStart { get; set; } = null!;

        public ServerIpv4FirewallRulesArgs()
        {
        }
    }

    public sealed class ServerIpv4FirewallRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// End of the firewall rule range as IPv4 address.
        /// </summary>
        [Input("rangeEnd", required: true)]
        public Input<string> RangeEnd { get; set; } = null!;

        /// <summary>
        /// Start of the firewall rule range as IPv4 address.
        /// </summary>
        [Input("rangeStart", required: true)]
        public Input<string> RangeStart { get; set; } = null!;

        public ServerIpv4FirewallRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ServerIpv4FirewallRules
    {
        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// End of the firewall rule range as IPv4 address.
        /// </summary>
        public readonly string RangeEnd;
        /// <summary>
        /// Start of the firewall rule range as IPv4 address.
        /// </summary>
        public readonly string RangeStart;

        [OutputConstructor]
        private ServerIpv4FirewallRules(
            string name,
            string rangeEnd,
            string rangeStart)
        {
            Name = name;
            RangeEnd = rangeEnd;
            RangeStart = rangeStart;
        }
    }
    }
}
