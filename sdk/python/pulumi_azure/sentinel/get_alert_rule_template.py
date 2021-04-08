# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetAlertRuleTemplateResult',
    'AwaitableGetAlertRuleTemplateResult',
    'get_alert_rule_template',
]

@pulumi.output_type
class GetAlertRuleTemplateResult:
    """
    A collection of values returned by getAlertRuleTemplate.
    """
    def __init__(__self__, display_name=None, id=None, log_analytics_workspace_id=None, name=None, scheduled_templates=None, security_incident_templates=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_analytics_workspace_id and not isinstance(log_analytics_workspace_id, str):
            raise TypeError("Expected argument 'log_analytics_workspace_id' to be a str")
        pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if scheduled_templates and not isinstance(scheduled_templates, list):
            raise TypeError("Expected argument 'scheduled_templates' to be a list")
        pulumi.set(__self__, "scheduled_templates", scheduled_templates)
        if security_incident_templates and not isinstance(security_incident_templates, list):
            raise TypeError("Expected argument 'security_incident_templates' to be a list")
        pulumi.set(__self__, "security_incident_templates", security_incident_templates)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> str:
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scheduledTemplates")
    def scheduled_templates(self) -> Sequence['outputs.GetAlertRuleTemplateScheduledTemplateResult']:
        """
        A `scheduled_template` block as defined below. This only applies to Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "scheduled_templates")

    @property
    @pulumi.getter(name="securityIncidentTemplates")
    def security_incident_templates(self) -> Sequence['outputs.GetAlertRuleTemplateSecurityIncidentTemplateResult']:
        """
        A `security_incident_template` block as defined below. This only applies to Sentinel MS Security Incident Alert Rule Template.
        """
        return pulumi.get(self, "security_incident_templates")


class AwaitableGetAlertRuleTemplateResult(GetAlertRuleTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlertRuleTemplateResult(
            display_name=self.display_name,
            id=self.id,
            log_analytics_workspace_id=self.log_analytics_workspace_id,
            name=self.name,
            scheduled_templates=self.scheduled_templates,
            security_incident_templates=self.security_incident_templates)


def get_alert_rule_template(display_name: Optional[str] = None,
                            log_analytics_workspace_id: Optional[str] = None,
                            name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlertRuleTemplateResult:
    """
    Use this data source to access information about an existing Sentinel Alert Rule Template.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.sentinel.get_alert_rule_template(log_analytics_workspace_id="/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1",
        display_name="Create incidents based on Azure Security Center for IoT alerts")
    pulumi.export("id", example.id)
    ```


    :param str display_name: The display name of this Sentinel Alert Rule Template. Either `display_name` or `name` have to be specified.
    :param str log_analytics_workspace_id: The ID of the Log Analytics Workspace.
    :param str name: The name of this Sentinel Alert Rule Template. Either `display_name` or `name` have to be specified.
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['logAnalyticsWorkspaceId'] = log_analytics_workspace_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:sentinel/getAlertRuleTemplate:getAlertRuleTemplate', __args__, opts=opts, typ=GetAlertRuleTemplateResult).value

    return AwaitableGetAlertRuleTemplateResult(
        display_name=__ret__.display_name,
        id=__ret__.id,
        log_analytics_workspace_id=__ret__.log_analytics_workspace_id,
        name=__ret__.name,
        scheduled_templates=__ret__.scheduled_templates,
        security_incident_templates=__ret__.security_incident_templates)
