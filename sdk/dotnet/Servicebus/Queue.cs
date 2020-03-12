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
    /// Manages a ServiceBus Queue.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_queue.html.markdown.
    /// </summary>
    public partial class Queue : Pulumi.CustomResource
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the
        /// Queue is automatically deleted, minimum of 5 minutes.
        /// </summary>
        [Output("autoDeleteOnIdle")]
        public Output<string> AutoDeleteOnIdle { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        /// </summary>
        [Output("deadLetteringOnMessageExpiration")]
        public Output<bool?> DeadLetteringOnMessageExpiration { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration of the TTL of messages sent to this
        /// queue. This is the default value used when TTL is not set on message itself.
        /// </summary>
        [Output("defaultMessageTtl")]
        public Output<string> DefaultMessageTtl { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration during which
        /// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
        /// </summary>
        [Output("duplicateDetectionHistoryTimeWindow")]
        public Output<string> DuplicateDetectionHistoryTimeWindow { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether Express Entities
        /// are enabled. An express queue holds a message in memory temporarily before writing
        /// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
        /// be set to `false`.
        /// </summary>
        [Output("enableExpress")]
        public Output<bool?> EnableExpress { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether to enable
        /// the queue to be partitioned across multiple message brokers. Changing this forces
        /// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
        /// be set to `true`.
        /// </summary>
        [Output("enablePartitioning")]
        public Output<bool?> EnablePartitioning { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
        /// </summary>
        [Output("lockDuration")]
        public Output<string> LockDuration { get; private set; } = null!;

        /// <summary>
        /// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
        /// </summary>
        [Output("maxDeliveryCount")]
        public Output<int?> MaxDeliveryCount { get; private set; } = null!;

        /// <summary>
        /// Integer value which controls the size of
        /// memory allocated for the queue. For supported values see the "Queue/topic size"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        [Output("maxSizeInMegabytes")]
        public Output<int> MaxSizeInMegabytes { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Queue resource. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace to create
        /// this queue in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether
        /// the Queue requires duplicate detection. Changing this forces
        /// a new resource to be created. Defaults to `false`.
        /// </summary>
        [Output("requiresDuplicateDetection")]
        public Output<bool?> RequiresDuplicateDetection { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether the Queue requires sessions.
        /// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
        /// a queue can guarantee first-in-first-out delivery of messages.
        /// Changing this forces a new resource to be created. Defaults to `false`.
        /// </summary>
        [Output("requiresSession")]
        public Output<bool?> RequiresSession { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs args, CustomResourceOptions? options = null)
            : base("azure:servicebus/queue:Queue", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
            : base("azure:servicebus/queue:Queue", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,                Aliases = { new Alias { Type = "azure:eventhub/queue:Queue" } },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Queue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Queue Get(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
        {
            return new Queue(name, id, state, options);
        }
    }

    public sealed class QueueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the
        /// Queue is automatically deleted, minimum of 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        /// </summary>
        [Input("deadLetteringOnMessageExpiration")]
        public Input<bool>? DeadLetteringOnMessageExpiration { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration of the TTL of messages sent to this
        /// queue. This is the default value used when TTL is not set on message itself.
        /// </summary>
        [Input("defaultMessageTtl")]
        public Input<string>? DefaultMessageTtl { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration during which
        /// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
        /// </summary>
        [Input("duplicateDetectionHistoryTimeWindow")]
        public Input<string>? DuplicateDetectionHistoryTimeWindow { get; set; }

        /// <summary>
        /// Boolean flag which controls whether Express Entities
        /// are enabled. An express queue holds a message in memory temporarily before writing
        /// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
        /// be set to `false`.
        /// </summary>
        [Input("enableExpress")]
        public Input<bool>? EnableExpress { get; set; }

        /// <summary>
        /// Boolean flag which controls whether to enable
        /// the queue to be partitioned across multiple message brokers. Changing this forces
        /// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
        /// be set to `true`.
        /// </summary>
        [Input("enablePartitioning")]
        public Input<bool>? EnablePartitioning { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
        /// </summary>
        [Input("lockDuration")]
        public Input<string>? LockDuration { get; set; }

        /// <summary>
        /// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        /// <summary>
        /// Integer value which controls the size of
        /// memory allocated for the queue. For supported values see the "Queue/topic size"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        [Input("maxSizeInMegabytes")]
        public Input<int>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Queue resource. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace to create
        /// this queue in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether
        /// the Queue requires duplicate detection. Changing this forces
        /// a new resource to be created. Defaults to `false`.
        /// </summary>
        [Input("requiresDuplicateDetection")]
        public Input<bool>? RequiresDuplicateDetection { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Queue requires sessions.
        /// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
        /// a queue can guarantee first-in-first-out delivery of messages.
        /// Changing this forces a new resource to be created. Defaults to `false`.
        /// </summary>
        [Input("requiresSession")]
        public Input<bool>? RequiresSession { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public QueueArgs()
        {
        }
    }

    public sealed class QueueState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the
        /// Queue is automatically deleted, minimum of 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        /// </summary>
        [Input("deadLetteringOnMessageExpiration")]
        public Input<bool>? DeadLetteringOnMessageExpiration { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration of the TTL of messages sent to this
        /// queue. This is the default value used when TTL is not set on message itself.
        /// </summary>
        [Input("defaultMessageTtl")]
        public Input<string>? DefaultMessageTtl { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration during which
        /// duplicates can be detected. Default value is 10 minutes. (`PT10M`)
        /// </summary>
        [Input("duplicateDetectionHistoryTimeWindow")]
        public Input<string>? DuplicateDetectionHistoryTimeWindow { get; set; }

        /// <summary>
        /// Boolean flag which controls whether Express Entities
        /// are enabled. An express queue holds a message in memory temporarily before writing
        /// it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST
        /// be set to `false`.
        /// </summary>
        [Input("enableExpress")]
        public Input<bool>? EnableExpress { get; set; }

        /// <summary>
        /// Boolean flag which controls whether to enable
        /// the queue to be partitioned across multiple message brokers. Changing this forces
        /// a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST
        /// be set to `true`.
        /// </summary>
        [Input("enablePartitioning")]
        public Input<bool>? EnablePartitioning { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute. (`PT1M`)
        /// </summary>
        [Input("lockDuration")]
        public Input<string>? LockDuration { get; set; }

        /// <summary>
        /// Integer value which controls when a message is automatically deadlettered. Defaults to `10`.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        /// <summary>
        /// Integer value which controls the size of
        /// memory allocated for the queue. For supported values see the "Queue/topic size"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        [Input("maxSizeInMegabytes")]
        public Input<int>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Queue resource. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace to create
        /// this queue in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// Boolean flag which controls whether
        /// the Queue requires duplicate detection. Changing this forces
        /// a new resource to be created. Defaults to `false`.
        /// </summary>
        [Input("requiresDuplicateDetection")]
        public Input<bool>? RequiresDuplicateDetection { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Queue requires sessions.
        /// This will allow ordered handling of unbounded sequences of related messages. With sessions enabled
        /// a queue can guarantee first-in-first-out delivery of messages.
        /// Changing this forces a new resource to be created. Defaults to `false`.
        /// </summary>
        [Input("requiresSession")]
        public Input<bool>? RequiresSession { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public QueueState()
        {
        }
    }
}
