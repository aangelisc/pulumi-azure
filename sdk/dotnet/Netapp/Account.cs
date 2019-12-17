// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp
{
    /// <summary>
    /// Manages a NetApp Account.
    /// 
    /// &gt; **NOTE:** Azure allows only one active directory can be joined to a single subscription at a time for NetApp Account.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/netapp_account.html.markdown.
    /// </summary>
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// A `active_directory` block as defined below.
        /// </summary>
        [Output("activeDirectory")]
        public Output<Outputs.AccountActiveDirectory?> ActiveDirectory { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the NetApp Account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azure:netapp/account:Account", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:netapp/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `active_directory` block as defined below.
        /// </summary>
        [Input("activeDirectory")]
        public Input<Inputs.AccountActiveDirectoryArgs>? ActiveDirectory { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the NetApp Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AccountArgs()
        {
        }
    }

    public sealed class AccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `active_directory` block as defined below.
        /// </summary>
        [Input("activeDirectory")]
        public Input<Inputs.AccountActiveDirectoryGetArgs>? ActiveDirectory { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the NetApp Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group where the NetApp Account should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public AccountState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AccountActiveDirectoryArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers", required: true)]
        private InputList<string>? _dnsServers;
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("smbServerName", required: true)]
        public Input<string> SmbServerName { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public AccountActiveDirectoryArgs()
        {
        }
    }

    public sealed class AccountActiveDirectoryGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers", required: true)]
        private InputList<string>? _dnsServers;
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("smbServerName", required: true)]
        public Input<string> SmbServerName { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public AccountActiveDirectoryGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AccountActiveDirectory
    {
        public readonly ImmutableArray<string> DnsServers;
        public readonly string Domain;
        public readonly string? OrganizationalUnit;
        public readonly string Password;
        public readonly string SmbServerName;
        public readonly string Username;

        [OutputConstructor]
        private AccountActiveDirectory(
            ImmutableArray<string> dnsServers,
            string domain,
            string? organizationalUnit,
            string password,
            string smbServerName,
            string username)
        {
            DnsServers = dnsServers;
            Domain = domain;
            OrganizationalUnit = organizationalUnit;
            Password = password;
            SmbServerName = smbServerName;
            Username = username;
        }
    }
    }
}