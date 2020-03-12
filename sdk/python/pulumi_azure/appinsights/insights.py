# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Insights(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    The App ID associated with this Application Insights component.
    """
    application_type: pulumi.Output[str]
    """
    Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
    """
    daily_data_cap_in_gb: pulumi.Output[float]
    """
    Specifies the Application Insights component daily data volume cap in GB.
    """
    daily_data_cap_notifications_disabled: pulumi.Output[bool]
    """
    Specifies if a notification email will be send when the daily data volume cap is met.
    """
    instrumentation_key: pulumi.Output[str]
    """
    The Instrumentation Key for this Application Insights component.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Application Insights component. Changing this forces a
    new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the Application Insights component.
    """
    retention_in_days: pulumi.Output[float]
    """
    Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`.
    """
    sampling_percentage: pulumi.Output[float]
    """
    Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, application_type=None, daily_data_cap_in_gb=None, daily_data_cap_notifications_disabled=None, location=None, name=None, resource_group_name=None, retention_in_days=None, sampling_percentage=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Application Insights component.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/application_insights.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_type: Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
        :param pulumi.Input[float] daily_data_cap_in_gb: Specifies the Application Insights component daily data volume cap in GB.
        :param pulumi.Input[bool] daily_data_cap_notifications_disabled: Specifies if a notification email will be send when the daily data volume cap is met.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Application Insights component. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the Application Insights component.
        :param pulumi.Input[float] retention_in_days: Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`.
        :param pulumi.Input[float] sampling_percentage: Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if application_type is None:
                raise TypeError("Missing required property 'application_type'")
            __props__['application_type'] = application_type
            __props__['daily_data_cap_in_gb'] = daily_data_cap_in_gb
            __props__['daily_data_cap_notifications_disabled'] = daily_data_cap_notifications_disabled
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retention_in_days'] = retention_in_days
            __props__['sampling_percentage'] = sampling_percentage
            __props__['tags'] = tags
            __props__['app_id'] = None
            __props__['instrumentation_key'] = None
        super(Insights, __self__).__init__(
            'azure:appinsights/insights:Insights',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, application_type=None, daily_data_cap_in_gb=None, daily_data_cap_notifications_disabled=None, instrumentation_key=None, location=None, name=None, resource_group_name=None, retention_in_days=None, sampling_percentage=None, tags=None):
        """
        Get an existing Insights resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The App ID associated with this Application Insights component.
        :param pulumi.Input[str] application_type: Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
        :param pulumi.Input[float] daily_data_cap_in_gb: Specifies the Application Insights component daily data volume cap in GB.
        :param pulumi.Input[bool] daily_data_cap_notifications_disabled: Specifies if a notification email will be send when the daily data volume cap is met.
        :param pulumi.Input[str] instrumentation_key: The Instrumentation Key for this Application Insights component.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Application Insights component. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the Application Insights component.
        :param pulumi.Input[float] retention_in_days: Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`.
        :param pulumi.Input[float] sampling_percentage: Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["application_type"] = application_type
        __props__["daily_data_cap_in_gb"] = daily_data_cap_in_gb
        __props__["daily_data_cap_notifications_disabled"] = daily_data_cap_notifications_disabled
        __props__["instrumentation_key"] = instrumentation_key
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["retention_in_days"] = retention_in_days
        __props__["sampling_percentage"] = sampling_percentage
        __props__["tags"] = tags
        return Insights(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

