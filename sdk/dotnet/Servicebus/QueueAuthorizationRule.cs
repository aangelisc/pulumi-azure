// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    /// <summary>
    /// Manages an Authorization Rule for a ServiceBus Queue.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_queue_authorization_rule.html.markdown.
    /// </summary>
    public partial class QueueAuthorizationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
        /// </summary>
        [Output("listen")]
        public Output<bool?> Listen { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        /// </summary>
        [Output("manage")]
        public Output<bool?> Manage { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The Primary Connection String for the Authorization Rule.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The Primary Key for the Authorization Rule.
        /// </summary>
        [Output("primaryKey")]
        public Output<string> PrimaryKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
        /// </summary>
        [Output("queueName")]
        public Output<string> QueueName { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Connection String for the Authorization Rule.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The Secondary Key for the Authorization Rule.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string> SecondaryKey { get; private set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
        /// </summary>
        [Output("send")]
        public Output<bool?> Send { get; private set; } = null!;


        /// <summary>
        /// Create a QueueAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueueAuthorizationRule(string name, QueueAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private QueueAuthorizationRule(string name, Input<string> id, QueueAuthorizationRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,                Aliases = { new Alias { Type = "azure:eventhub/queueAuthorizationRule:QueueAuthorizationRule" } },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing QueueAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueueAuthorizationRule Get(string name, Input<string> id, QueueAuthorizationRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new QueueAuthorizationRule(name, id, state, options);
        }
    }

    public sealed class QueueAuthorizationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
        /// </summary>
        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        /// <summary>
        /// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        /// </summary>
        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
        /// </summary>
        [Input("send")]
        public Input<bool>? Send { get; set; }

        public QueueAuthorizationRuleArgs()
        {
        }
    }

    public sealed class QueueAuthorizationRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
        /// </summary>
        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        /// <summary>
        /// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        /// </summary>
        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The Primary Connection String for the Authorization Rule.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        /// <summary>
        /// The Primary Key for the Authorization Rule.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
        /// </summary>
        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Connection String for the Authorization Rule.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// The Secondary Key for the Authorization Rule.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        /// <summary>
        /// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
        /// </summary>
        [Input("send")]
        public Input<bool>? Send { get; set; }

        public QueueAuthorizationRuleState()
        {
        }
    }
}
