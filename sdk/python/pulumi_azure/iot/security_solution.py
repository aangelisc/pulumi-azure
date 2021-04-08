# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecuritySolutionArgs', 'SecuritySolution']

@pulumi.input_type
class SecuritySolutionArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 iothub_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 resource_group_name: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_to_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 log_unmasked_ips_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_for_resources: Optional[pulumi.Input[str]] = None,
                 query_subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recommendations_enabled: Optional[pulumi.Input['SecuritySolutionRecommendationsEnabledArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecuritySolution resource.
        :param pulumi.Input[str] display_name: Specifies the Display Name for this Iot Security Solution.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] iothub_ids: Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Is the Iot Security Solution enabled? Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_to_exports: A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the Log Analytics Workspace ID to which the security data will be sent.
        :param pulumi.Input[bool] log_unmasked_ips_enabled: Should ip addressed be unmasked in the log? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query_for_resources: An Azure Resource Graph query used to set the resources monitored.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_subscription_ids: A list of subscription Ids on which the user defined resources query should be executed.
        :param pulumi.Input['SecuritySolutionRecommendationsEnabledArgs'] recommendations_enabled: A `recommendations_enabled` block of options to enable or disable as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "iothub_ids", iothub_ids)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if events_to_exports is not None:
            pulumi.set(__self__, "events_to_exports", events_to_exports)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_analytics_workspace_id is not None:
            pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        if log_unmasked_ips_enabled is not None:
            pulumi.set(__self__, "log_unmasked_ips_enabled", log_unmasked_ips_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query_for_resources is not None:
            pulumi.set(__self__, "query_for_resources", query_for_resources)
        if query_subscription_ids is not None:
            pulumi.set(__self__, "query_subscription_ids", query_subscription_ids)
        if recommendations_enabled is not None:
            pulumi.set(__self__, "recommendations_enabled", recommendations_enabled)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Specifies the Display Name for this Iot Security Solution.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="iothubIds")
    def iothub_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        """
        return pulumi.get(self, "iothub_ids")

    @iothub_ids.setter
    def iothub_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "iothub_ids", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the Iot Security Solution enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventsToExports")
    def events_to_exports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        """
        return pulumi.get(self, "events_to_exports")

    @events_to_exports.setter
    def events_to_exports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "events_to_exports", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Log Analytics Workspace ID to which the security data will be sent.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @log_analytics_workspace_id.setter
    def log_analytics_workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_analytics_workspace_id", value)

    @property
    @pulumi.getter(name="logUnmaskedIpsEnabled")
    def log_unmasked_ips_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should ip addressed be unmasked in the log? Defaults to `false`.
        """
        return pulumi.get(self, "log_unmasked_ips_enabled")

    @log_unmasked_ips_enabled.setter
    def log_unmasked_ips_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "log_unmasked_ips_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="queryForResources")
    def query_for_resources(self) -> Optional[pulumi.Input[str]]:
        """
        An Azure Resource Graph query used to set the resources monitored.
        """
        return pulumi.get(self, "query_for_resources")

    @query_for_resources.setter
    def query_for_resources(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_for_resources", value)

    @property
    @pulumi.getter(name="querySubscriptionIds")
    def query_subscription_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of subscription Ids on which the user defined resources query should be executed.
        """
        return pulumi.get(self, "query_subscription_ids")

    @query_subscription_ids.setter
    def query_subscription_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "query_subscription_ids", value)

    @property
    @pulumi.getter(name="recommendationsEnabled")
    def recommendations_enabled(self) -> Optional[pulumi.Input['SecuritySolutionRecommendationsEnabledArgs']]:
        """
        A `recommendations_enabled` block of options to enable or disable as defined below.
        """
        return pulumi.get(self, "recommendations_enabled")

    @recommendations_enabled.setter
    def recommendations_enabled(self, value: Optional[pulumi.Input['SecuritySolutionRecommendationsEnabledArgs']]):
        pulumi.set(self, "recommendations_enabled", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SecuritySolutionState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_to_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iothub_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 log_unmasked_ips_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_for_resources: Optional[pulumi.Input[str]] = None,
                 query_subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recommendations_enabled: Optional[pulumi.Input['SecuritySolutionRecommendationsEnabledArgs']] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecuritySolution resources.
        :param pulumi.Input[str] display_name: Specifies the Display Name for this Iot Security Solution.
        :param pulumi.Input[bool] enabled: Is the Iot Security Solution enabled? Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_to_exports: A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] iothub_ids: Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the Log Analytics Workspace ID to which the security data will be sent.
        :param pulumi.Input[bool] log_unmasked_ips_enabled: Should ip addressed be unmasked in the log? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query_for_resources: An Azure Resource Graph query used to set the resources monitored.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_subscription_ids: A list of subscription Ids on which the user defined resources query should be executed.
        :param pulumi.Input['SecuritySolutionRecommendationsEnabledArgs'] recommendations_enabled: A `recommendations_enabled` block of options to enable or disable as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if events_to_exports is not None:
            pulumi.set(__self__, "events_to_exports", events_to_exports)
        if iothub_ids is not None:
            pulumi.set(__self__, "iothub_ids", iothub_ids)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_analytics_workspace_id is not None:
            pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        if log_unmasked_ips_enabled is not None:
            pulumi.set(__self__, "log_unmasked_ips_enabled", log_unmasked_ips_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query_for_resources is not None:
            pulumi.set(__self__, "query_for_resources", query_for_resources)
        if query_subscription_ids is not None:
            pulumi.set(__self__, "query_subscription_ids", query_subscription_ids)
        if recommendations_enabled is not None:
            pulumi.set(__self__, "recommendations_enabled", recommendations_enabled)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Display Name for this Iot Security Solution.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the Iot Security Solution enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventsToExports")
    def events_to_exports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        """
        return pulumi.get(self, "events_to_exports")

    @events_to_exports.setter
    def events_to_exports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "events_to_exports", value)

    @property
    @pulumi.getter(name="iothubIds")
    def iothub_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        """
        return pulumi.get(self, "iothub_ids")

    @iothub_ids.setter
    def iothub_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "iothub_ids", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Log Analytics Workspace ID to which the security data will be sent.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @log_analytics_workspace_id.setter
    def log_analytics_workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_analytics_workspace_id", value)

    @property
    @pulumi.getter(name="logUnmaskedIpsEnabled")
    def log_unmasked_ips_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should ip addressed be unmasked in the log? Defaults to `false`.
        """
        return pulumi.get(self, "log_unmasked_ips_enabled")

    @log_unmasked_ips_enabled.setter
    def log_unmasked_ips_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "log_unmasked_ips_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="queryForResources")
    def query_for_resources(self) -> Optional[pulumi.Input[str]]:
        """
        An Azure Resource Graph query used to set the resources monitored.
        """
        return pulumi.get(self, "query_for_resources")

    @query_for_resources.setter
    def query_for_resources(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_for_resources", value)

    @property
    @pulumi.getter(name="querySubscriptionIds")
    def query_subscription_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of subscription Ids on which the user defined resources query should be executed.
        """
        return pulumi.get(self, "query_subscription_ids")

    @query_subscription_ids.setter
    def query_subscription_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "query_subscription_ids", value)

    @property
    @pulumi.getter(name="recommendationsEnabled")
    def recommendations_enabled(self) -> Optional[pulumi.Input['SecuritySolutionRecommendationsEnabledArgs']]:
        """
        A `recommendations_enabled` block of options to enable or disable as defined below.
        """
        return pulumi.get(self, "recommendations_enabled")

    @recommendations_enabled.setter
    def recommendations_enabled(self, value: Optional[pulumi.Input['SecuritySolutionRecommendationsEnabledArgs']]):
        pulumi.set(self, "recommendations_enabled", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class SecuritySolution(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_to_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iothub_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 log_unmasked_ips_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_for_resources: Optional[pulumi.Input[str]] = None,
                 query_subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recommendations_enabled: Optional[pulumi.Input[pulumi.InputType['SecuritySolutionRecommendationsEnabledArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an iot security solution.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.iot.IoTHubSkuArgs(
                name="S1",
                capacity=1,
            ))
        example_security_solution = azure.iot.SecuritySolution("exampleSecuritySolution",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            display_name="Iot Security Solution",
            iothub_ids=[example_io_t_hub.id])
        ```

        ## Import

        Iot Security Solution can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/securitySolution:SecuritySolution example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Security/IoTSecuritySolutions/solution1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Specifies the Display Name for this Iot Security Solution.
        :param pulumi.Input[bool] enabled: Is the Iot Security Solution enabled? Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_to_exports: A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] iothub_ids: Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the Log Analytics Workspace ID to which the security data will be sent.
        :param pulumi.Input[bool] log_unmasked_ips_enabled: Should ip addressed be unmasked in the log? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query_for_resources: An Azure Resource Graph query used to set the resources monitored.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_subscription_ids: A list of subscription Ids on which the user defined resources query should be executed.
        :param pulumi.Input[pulumi.InputType['SecuritySolutionRecommendationsEnabledArgs']] recommendations_enabled: A `recommendations_enabled` block of options to enable or disable as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecuritySolutionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an iot security solution.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.iot.IoTHubSkuArgs(
                name="S1",
                capacity=1,
            ))
        example_security_solution = azure.iot.SecuritySolution("exampleSecuritySolution",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            display_name="Iot Security Solution",
            iothub_ids=[example_io_t_hub.id])
        ```

        ## Import

        Iot Security Solution can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/securitySolution:SecuritySolution example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Security/IoTSecuritySolutions/solution1
        ```

        :param str resource_name: The name of the resource.
        :param SecuritySolutionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecuritySolutionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_to_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iothub_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 log_unmasked_ips_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_for_resources: Optional[pulumi.Input[str]] = None,
                 query_subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recommendations_enabled: Optional[pulumi.Input[pulumi.InputType['SecuritySolutionRecommendationsEnabledArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecuritySolutionArgs.__new__(SecuritySolutionArgs)

            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["events_to_exports"] = events_to_exports
            if iothub_ids is None and not opts.urn:
                raise TypeError("Missing required property 'iothub_ids'")
            __props__.__dict__["iothub_ids"] = iothub_ids
            __props__.__dict__["location"] = location
            __props__.__dict__["log_analytics_workspace_id"] = log_analytics_workspace_id
            __props__.__dict__["log_unmasked_ips_enabled"] = log_unmasked_ips_enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["query_for_resources"] = query_for_resources
            __props__.__dict__["query_subscription_ids"] = query_subscription_ids
            __props__.__dict__["recommendations_enabled"] = recommendations_enabled
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
        super(SecuritySolution, __self__).__init__(
            'azure:iot/securitySolution:SecuritySolution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            events_to_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            iothub_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
            log_unmasked_ips_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query_for_resources: Optional[pulumi.Input[str]] = None,
            query_subscription_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            recommendations_enabled: Optional[pulumi.Input[pulumi.InputType['SecuritySolutionRecommendationsEnabledArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SecuritySolution':
        """
        Get an existing SecuritySolution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Specifies the Display Name for this Iot Security Solution.
        :param pulumi.Input[bool] enabled: Is the Iot Security Solution enabled? Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_to_exports: A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] iothub_ids: Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the Log Analytics Workspace ID to which the security data will be sent.
        :param pulumi.Input[bool] log_unmasked_ips_enabled: Should ip addressed be unmasked in the log? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[str] query_for_resources: An Azure Resource Graph query used to set the resources monitored.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_subscription_ids: A list of subscription Ids on which the user defined resources query should be executed.
        :param pulumi.Input[pulumi.InputType['SecuritySolutionRecommendationsEnabledArgs']] recommendations_enabled: A `recommendations_enabled` block of options to enable or disable as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecuritySolutionState.__new__(_SecuritySolutionState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["events_to_exports"] = events_to_exports
        __props__.__dict__["iothub_ids"] = iothub_ids
        __props__.__dict__["location"] = location
        __props__.__dict__["log_analytics_workspace_id"] = log_analytics_workspace_id
        __props__.__dict__["log_unmasked_ips_enabled"] = log_unmasked_ips_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["query_for_resources"] = query_for_resources
        __props__.__dict__["query_subscription_ids"] = query_subscription_ids
        __props__.__dict__["recommendations_enabled"] = recommendations_enabled
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        return SecuritySolution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Specifies the Display Name for this Iot Security Solution.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the Iot Security Solution enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventsToExports")
    def events_to_exports(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        """
        return pulumi.get(self, "events_to_exports")

    @property
    @pulumi.getter(name="iothubIds")
    def iothub_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        """
        return pulumi.get(self, "iothub_ids")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Log Analytics Workspace ID to which the security data will be sent.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter(name="logUnmaskedIpsEnabled")
    def log_unmasked_ips_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Should ip addressed be unmasked in the log? Defaults to `false`.
        """
        return pulumi.get(self, "log_unmasked_ips_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryForResources")
    def query_for_resources(self) -> pulumi.Output[str]:
        """
        An Azure Resource Graph query used to set the resources monitored.
        """
        return pulumi.get(self, "query_for_resources")

    @property
    @pulumi.getter(name="querySubscriptionIds")
    def query_subscription_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of subscription Ids on which the user defined resources query should be executed.
        """
        return pulumi.get(self, "query_subscription_ids")

    @property
    @pulumi.getter(name="recommendationsEnabled")
    def recommendations_enabled(self) -> pulumi.Output['outputs.SecuritySolutionRecommendationsEnabled']:
        """
        A `recommendations_enabled` block of options to enable or disable as defined below.
        """
        return pulumi.get(self, "recommendations_enabled")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

