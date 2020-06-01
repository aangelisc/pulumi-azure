// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid.Outputs
{

    [OutputType]
    public sealed class TopicInputMappingDefaultValues
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? DataVersion;
        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? EventType;
        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Subject;

        [OutputConstructor]
        private TopicInputMappingDefaultValues(
            string? dataVersion,

            string? eventType,

            string? subject)
        {
            DataVersion = dataVersion;
            EventType = eventType;
            Subject = subject;
        }
    }
}
