# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
]

@pulumi.output_type
class GetServiceResult:
    """
    A collection of values returned by getService.
    """
    def __init__(__self__, additional_locations=None, developer_portal_url=None, gateway_regional_url=None, gateway_url=None, hostname_configurations=None, id=None, identities=None, location=None, management_api_url=None, name=None, notification_sender_email=None, portal_url=None, public_ip_addresses=None, publisher_email=None, publisher_name=None, resource_group_name=None, scm_url=None, sku_name=None, tags=None):
        if additional_locations and not isinstance(additional_locations, list):
            raise TypeError("Expected argument 'additional_locations' to be a list")
        pulumi.set(__self__, "additional_locations", additional_locations)
        if developer_portal_url and not isinstance(developer_portal_url, str):
            raise TypeError("Expected argument 'developer_portal_url' to be a str")
        pulumi.set(__self__, "developer_portal_url", developer_portal_url)
        if gateway_regional_url and not isinstance(gateway_regional_url, str):
            raise TypeError("Expected argument 'gateway_regional_url' to be a str")
        pulumi.set(__self__, "gateway_regional_url", gateway_regional_url)
        if gateway_url and not isinstance(gateway_url, str):
            raise TypeError("Expected argument 'gateway_url' to be a str")
        pulumi.set(__self__, "gateway_url", gateway_url)
        if hostname_configurations and not isinstance(hostname_configurations, list):
            raise TypeError("Expected argument 'hostname_configurations' to be a list")
        pulumi.set(__self__, "hostname_configurations", hostname_configurations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identities and not isinstance(identities, list):
            raise TypeError("Expected argument 'identities' to be a list")
        pulumi.set(__self__, "identities", identities)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if management_api_url and not isinstance(management_api_url, str):
            raise TypeError("Expected argument 'management_api_url' to be a str")
        pulumi.set(__self__, "management_api_url", management_api_url)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notification_sender_email and not isinstance(notification_sender_email, str):
            raise TypeError("Expected argument 'notification_sender_email' to be a str")
        pulumi.set(__self__, "notification_sender_email", notification_sender_email)
        if portal_url and not isinstance(portal_url, str):
            raise TypeError("Expected argument 'portal_url' to be a str")
        pulumi.set(__self__, "portal_url", portal_url)
        if public_ip_addresses and not isinstance(public_ip_addresses, list):
            raise TypeError("Expected argument 'public_ip_addresses' to be a list")
        pulumi.set(__self__, "public_ip_addresses", public_ip_addresses)
        if publisher_email and not isinstance(publisher_email, str):
            raise TypeError("Expected argument 'publisher_email' to be a str")
        pulumi.set(__self__, "publisher_email", publisher_email)
        if publisher_name and not isinstance(publisher_name, str):
            raise TypeError("Expected argument 'publisher_name' to be a str")
        pulumi.set(__self__, "publisher_name", publisher_name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if scm_url and not isinstance(scm_url, str):
            raise TypeError("Expected argument 'scm_url' to be a str")
        pulumi.set(__self__, "scm_url", scm_url)
        if sku_name and not isinstance(sku_name, str):
            raise TypeError("Expected argument 'sku_name' to be a str")
        pulumi.set(__self__, "sku_name", sku_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="additionalLocations")
    def additional_locations(self) -> List['outputs.GetServiceAdditionalLocationResult']:
        """
        One or more `additional_location` blocks as defined below
        """
        return pulumi.get(self, "additional_locations")

    @property
    @pulumi.getter(name="developerPortalUrl")
    def developer_portal_url(self) -> str:
        """
        The URL for the Developer Portal associated with this API Management service.
        """
        return pulumi.get(self, "developer_portal_url")

    @property
    @pulumi.getter(name="gatewayRegionalUrl")
    def gateway_regional_url(self) -> str:
        """
        Gateway URL of the API Management service in the Region.
        """
        return pulumi.get(self, "gateway_regional_url")

    @property
    @pulumi.getter(name="gatewayUrl")
    def gateway_url(self) -> str:
        """
        The URL for the API Management Service's Gateway.
        """
        return pulumi.get(self, "gateway_url")

    @property
    @pulumi.getter(name="hostnameConfigurations")
    def hostname_configurations(self) -> List['outputs.GetServiceHostnameConfigurationResult']:
        """
        A `hostname_configuration` block as defined below.
        """
        return pulumi.get(self, "hostname_configurations")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identities(self) -> List['outputs.GetServiceIdentityResult']:
        """
        (Optional) An `identity` block as defined below.
        """
        return pulumi.get(self, "identities")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location name of the additional region among Azure Data center regions.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementApiUrl")
    def management_api_url(self) -> str:
        """
        The URL for the Management API.
        """
        return pulumi.get(self, "management_api_url")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the plan's pricing tier.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationSenderEmail")
    def notification_sender_email(self) -> str:
        """
        The email address from which the notification will be sent.
        """
        return pulumi.get(self, "notification_sender_email")

    @property
    @pulumi.getter(name="portalUrl")
    def portal_url(self) -> str:
        """
        The URL of the Publisher Portal.
        """
        return pulumi.get(self, "portal_url")

    @property
    @pulumi.getter(name="publicIpAddresses")
    def public_ip_addresses(self) -> List[str]:
        """
        Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
        """
        return pulumi.get(self, "public_ip_addresses")

    @property
    @pulumi.getter(name="publisherEmail")
    def publisher_email(self) -> str:
        """
        The email of Publisher/Company of the API Management Service.
        """
        return pulumi.get(self, "publisher_email")

    @property
    @pulumi.getter(name="publisherName")
    def publisher_name(self) -> str:
        """
        The name of the Publisher/Company of the API Management Service.
        """
        return pulumi.get(self, "publisher_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="scmUrl")
    def scm_url(self) -> str:
        """
        The SCM (Source Code Management) endpoint.
        """
        return pulumi.get(self, "scm_url")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            additional_locations=self.additional_locations,
            developer_portal_url=self.developer_portal_url,
            gateway_regional_url=self.gateway_regional_url,
            gateway_url=self.gateway_url,
            hostname_configurations=self.hostname_configurations,
            id=self.id,
            identities=self.identities,
            location=self.location,
            management_api_url=self.management_api_url,
            name=self.name,
            notification_sender_email=self.notification_sender_email,
            portal_url=self.portal_url,
            public_ip_addresses=self.public_ip_addresses,
            publisher_email=self.publisher_email,
            publisher_name=self.publisher_name,
            resource_group_name=self.resource_group_name,
            scm_url=self.scm_url,
            sku_name=self.sku_name,
            tags=self.tags)


def get_service(name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Use this data source to access information about an existing API Management Service.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.apimanagement.get_service(name="search-api",
        resource_group_name="search-service")
    pulumi.export("apiManagementId", example.id)
    ```


    :param str name: The name of the API Management service.
    :param str resource_group_name: The Name of the Resource Group in which the API Management Service exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:apimanagement/getService:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        additional_locations=__ret__.additional_locations,
        developer_portal_url=__ret__.developer_portal_url,
        gateway_regional_url=__ret__.gateway_regional_url,
        gateway_url=__ret__.gateway_url,
        hostname_configurations=__ret__.hostname_configurations,
        id=__ret__.id,
        identities=__ret__.identities,
        location=__ret__.location,
        management_api_url=__ret__.management_api_url,
        name=__ret__.name,
        notification_sender_email=__ret__.notification_sender_email,
        portal_url=__ret__.portal_url,
        public_ip_addresses=__ret__.public_ip_addresses,
        publisher_email=__ret__.publisher_email,
        publisher_name=__ret__.publisher_name,
        resource_group_name=__ret__.resource_group_name,
        scm_url=__ret__.scm_url,
        sku_name=__ret__.sku_name,
        tags=__ret__.tags)
