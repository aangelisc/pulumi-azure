// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel
{
    public static class GetAlertRuleTemplate
    {
        /// <summary>
        /// Use this data source to access information about an existing Sentinel Alert Rule Template.
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
        ///         var example = Output.Create(Azure.Sentinel.GetAlertRuleTemplate.InvokeAsync(new Azure.Sentinel.GetAlertRuleTemplateArgs
        ///         {
        ///             LogAnalyticsWorkspaceId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1",
        ///             DisplayName = "Create incidents based on Azure Security Center for IoT alerts",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAlertRuleTemplateResult> InvokeAsync(GetAlertRuleTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlertRuleTemplateResult>("azure:sentinel/getAlertRuleTemplate:getAlertRuleTemplate", args ?? new GetAlertRuleTemplateArgs(), options.WithVersion());
    }


    public sealed class GetAlertRuleTemplateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name of this Sentinel Alert Rule Template. Either `display_name` or `name` have to be specified.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The ID of the Log Analytics Workspace.
        /// </summary>
        [Input("logAnalyticsWorkspaceId", required: true)]
        public string LogAnalyticsWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name of this Sentinel Alert Rule Template. Either `display_name` or `name` have to be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetAlertRuleTemplateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlertRuleTemplateResult
    {
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LogAnalyticsWorkspaceId;
        public readonly string Name;
        /// <summary>
        /// A `scheduled_template` block as defined below. This only applies to Sentinel Scheduled Alert Rule Template.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlertRuleTemplateScheduledTemplateResult> ScheduledTemplates;
        /// <summary>
        /// A `security_incident_template` block as defined below. This only applies to Sentinel MS Security Incident Alert Rule Template.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlertRuleTemplateSecurityIncidentTemplateResult> SecurityIncidentTemplates;

        [OutputConstructor]
        private GetAlertRuleTemplateResult(
            string displayName,

            string id,

            string logAnalyticsWorkspaceId,

            string name,

            ImmutableArray<Outputs.GetAlertRuleTemplateScheduledTemplateResult> scheduledTemplates,

            ImmutableArray<Outputs.GetAlertRuleTemplateSecurityIncidentTemplateResult> securityIncidentTemplates)
        {
            DisplayName = displayName;
            Id = id;
            LogAnalyticsWorkspaceId = logAnalyticsWorkspaceId;
            Name = name;
            ScheduledTemplates = scheduledTemplates;
            SecurityIncidentTemplates = securityIncidentTemplates;
        }
    }
}