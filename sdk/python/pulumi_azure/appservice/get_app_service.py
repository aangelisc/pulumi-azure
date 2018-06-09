# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetAppServiceResult(object):
    """
    A collection of values returned by getAppService.
    """
    def __init__(__self__, app_service_plan_id=None, app_settings=None, client_affinity_enabled=None, connection_strings=None, default_site_hostname=None, enabled=None, https_only=None, location=None, outbound_ip_addresses=None, site_config=None, site_credentials=None, source_controls=None, tags=None):
        if not app_service_plan_id:
            raise TypeError('Missing required argument app_service_plan_id')
        elif not isinstance(app_service_plan_id, basestring):
            raise TypeError('Expected argument app_service_plan_id to be a basestring')
        __self__.app_service_plan_id = app_service_plan_id
        """
        The ID of the App Service Plan within which the App Service exists.
        """
        if not app_settings:
            raise TypeError('Missing required argument app_settings')
        elif not isinstance(app_settings, dict):
            raise TypeError('Expected argument app_settings to be a dict')
        __self__.app_settings = app_settings
        """
        A key-value pair of App Settings for the App Service.
        """
        if not client_affinity_enabled:
            raise TypeError('Missing required argument client_affinity_enabled')
        elif not isinstance(client_affinity_enabled, bool):
            raise TypeError('Expected argument client_affinity_enabled to be a bool')
        __self__.client_affinity_enabled = client_affinity_enabled
        """
        Does the App Service send session affinity cookies, which route client requests in the same session to the same instance?
        """
        if not connection_strings:
            raise TypeError('Missing required argument connection_strings')
        elif not isinstance(connection_strings, list):
            raise TypeError('Expected argument connection_strings to be a list')
        __self__.connection_strings = connection_strings
        """
        An `connection_string` block as defined below.
        """
        if not default_site_hostname:
            raise TypeError('Missing required argument default_site_hostname')
        elif not isinstance(default_site_hostname, basestring):
            raise TypeError('Expected argument default_site_hostname to be a basestring')
        __self__.default_site_hostname = default_site_hostname
        if not enabled:
            raise TypeError('Missing required argument enabled')
        elif not isinstance(enabled, bool):
            raise TypeError('Expected argument enabled to be a bool')
        __self__.enabled = enabled
        """
        Is the App Service Enabled?
        """
        if not https_only:
            raise TypeError('Missing required argument https_only')
        elif not isinstance(https_only, bool):
            raise TypeError('Expected argument https_only to be a bool')
        __self__.https_only = https_only
        """
        Can the App Service only be accessed via HTTPS?
        """
        if not location:
            raise TypeError('Missing required argument location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected argument location to be a basestring')
        __self__.location = location
        """
        The Azure location where the App Service exists.
        """
        if not outbound_ip_addresses:
            raise TypeError('Missing required argument outbound_ip_addresses')
        elif not isinstance(outbound_ip_addresses, basestring):
            raise TypeError('Expected argument outbound_ip_addresses to be a basestring')
        __self__.outbound_ip_addresses = outbound_ip_addresses
        if not site_config:
            raise TypeError('Missing required argument site_config')
        elif not isinstance(site_config, dict):
            raise TypeError('Expected argument site_config to be a dict')
        __self__.site_config = site_config
        """
        A `site_config` block as defined below.
        """
        if not site_credentials:
            raise TypeError('Missing required argument site_credentials')
        elif not isinstance(site_credentials, list):
            raise TypeError('Expected argument site_credentials to be a list')
        __self__.site_credentials = site_credentials
        if not source_controls:
            raise TypeError('Missing required argument source_controls')
        elif not isinstance(source_controls, list):
            raise TypeError('Expected argument source_controls to be a list')
        __self__.source_controls = source_controls
        if not tags:
            raise TypeError('Missing required argument tags')
        elif not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """

def get_app_service(name=None, resource_group_name=None, site_config=None):
    """
    Use this data source to obtain information about an App Service.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['siteConfig'] = site_config
    __ret__ = pulumi.runtime.invoke('azure:appservice/getAppService:getAppService', __args__)

    return GetAppServiceResult(
        app_service_plan_id=__ret__['appServicePlanId'],
        app_settings=__ret__['appSettings'],
        client_affinity_enabled=__ret__['clientAffinityEnabled'],
        connection_strings=__ret__['connectionStrings'],
        default_site_hostname=__ret__['defaultSiteHostname'],
        enabled=__ret__['enabled'],
        https_only=__ret__['httpsOnly'],
        location=__ret__['location'],
        outbound_ip_addresses=__ret__['outboundIpAddresses'],
        site_config=__ret__['siteConfig'],
        site_credentials=__ret__['siteCredentials'],
        source_controls=__ret__['sourceControls'],
        tags=__ret__['tags'])
