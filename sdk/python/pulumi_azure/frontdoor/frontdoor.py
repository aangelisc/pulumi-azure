# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Frontdoor']


class Frontdoor(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_pool_health_probes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolHealthProbeArgs']]]]] = None,
                 backend_pool_load_balancings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolLoadBalancingArgs']]]]] = None,
                 backend_pools: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolArgs']]]]] = None,
                 backend_pools_send_receive_timeout_seconds: Optional[pulumi.Input[float]] = None,
                 enforce_backend_pools_certificate_name_check: Optional[pulumi.Input[bool]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 frontend_endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorFrontendEndpointArgs']]]]] = None,
                 load_balancer_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 routing_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorRoutingRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Front Door instance.

        Azure Front Door Service is Microsoft's highly available and scalable web application acceleration platform and global HTTP(s) load balancer. It provides built-in DDoS protection and application layer security and caching. Front Door enables you to build applications that maximize and automate high-availability and performance for your end-users. Use Front Door with Azure services including Web/Mobile Apps, Cloud Services and Virtual Machines – or combine it with on-premises services for hybrid deployments and smooth cloud migration.

        Below are some of the key scenarios that Azure Front Door Service addresses:
        * Use Front Door to improve application scale and availability with instant multi-region failover
        * Use Front Door to improve application performance with SSL offload and routing requests to the fastest available application backend.
        * Use Front Door for application layer security and DDoS protection for your application.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="EastUS2")
        example_frontdoor = azure.frontdoor.Frontdoor("exampleFrontdoor",
            resource_group_name=example_resource_group.name,
            enforce_backend_pools_certificate_name_check=False,
            routing_rules=[azure.frontdoor.FrontdoorRoutingRuleArgs(
                name="exampleRoutingRule1",
                accepted_protocols=[
                    "Http",
                    "Https",
                ],
                patterns_to_matches=["/*"],
                frontend_endpoints=["exampleFrontendEndpoint1"],
                forwarding_configuration=azure.frontdoor.FrontdoorRoutingRuleForwardingConfigurationArgs(
                    forwarding_protocol="MatchRequest",
                    backend_pool_name="exampleBackendBing",
                ),
            )],
            backend_pool_load_balancings=[azure.frontdoor.FrontdoorBackendPoolLoadBalancingArgs(
                name="exampleLoadBalancingSettings1",
            )],
            backend_pool_health_probes=[azure.frontdoor.FrontdoorBackendPoolHealthProbeArgs(
                name="exampleHealthProbeSetting1",
            )],
            backend_pools=[azure.frontdoor.FrontdoorBackendPoolArgs(
                name="exampleBackendBing",
                backends=[azure.frontdoor.FrontdoorBackendPoolBackendArgs(
                    host_header="www.bing.com",
                    address="www.bing.com",
                    http_port=80,
                    https_port=443,
                )],
                load_balancing_name="exampleLoadBalancingSettings1",
                health_probe_name="exampleHealthProbeSetting1",
            )],
            frontend_endpoints=[azure.frontdoor.FrontdoorFrontendEndpointArgs(
                name="exampleFrontendEndpoint1",
                host_name="example-FrontDoor.azurefd.net",
                custom_https_provisioning_enabled=False,
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolHealthProbeArgs']]]] backend_pool_health_probes: A `backend_pool_health_probe` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolLoadBalancingArgs']]]] backend_pool_load_balancings: A `backend_pool_load_balancing` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolArgs']]]] backend_pools: A `backend_pool` block as defined below.
        :param pulumi.Input[float] backend_pools_send_receive_timeout_seconds: Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
        :param pulumi.Input[bool] enforce_backend_pools_certificate_name_check: Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
        :param pulumi.Input[str] friendly_name: A friendly name for the Front Door service.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorFrontendEndpointArgs']]]] frontend_endpoints: A `frontend_endpoint` block as defined below.
        :param pulumi.Input[bool] load_balancer_enabled: Should the Front Door Load Balancer be Enabled? Defaults to `true`.
        :param pulumi.Input[str] name: Specifies the name of the Front Door service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorRoutingRuleArgs']]]] routing_rules: A `routing_rule` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if backend_pool_health_probes is None:
                raise TypeError("Missing required property 'backend_pool_health_probes'")
            __props__['backend_pool_health_probes'] = backend_pool_health_probes
            if backend_pool_load_balancings is None:
                raise TypeError("Missing required property 'backend_pool_load_balancings'")
            __props__['backend_pool_load_balancings'] = backend_pool_load_balancings
            if backend_pools is None:
                raise TypeError("Missing required property 'backend_pools'")
            __props__['backend_pools'] = backend_pools
            __props__['backend_pools_send_receive_timeout_seconds'] = backend_pools_send_receive_timeout_seconds
            if enforce_backend_pools_certificate_name_check is None:
                raise TypeError("Missing required property 'enforce_backend_pools_certificate_name_check'")
            __props__['enforce_backend_pools_certificate_name_check'] = enforce_backend_pools_certificate_name_check
            __props__['friendly_name'] = friendly_name
            if frontend_endpoints is None:
                raise TypeError("Missing required property 'frontend_endpoints'")
            __props__['frontend_endpoints'] = frontend_endpoints
            __props__['load_balancer_enabled'] = load_balancer_enabled
            if location is not None:
                warnings.warn("Due to the service's API changing 'location' must now always be set to 'Global' for new resources, however if the Front Door service was created prior 2020/03/10 it may continue to exist in a specific current location", DeprecationWarning)
                pulumi.log.warn("location is deprecated: Due to the service's API changing 'location' must now always be set to 'Global' for new resources, however if the Front Door service was created prior 2020/03/10 it may continue to exist in a specific current location")
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if routing_rules is None:
                raise TypeError("Missing required property 'routing_rules'")
            __props__['routing_rules'] = routing_rules
            __props__['tags'] = tags
            __props__['cname'] = None
            __props__['header_frontdoor_id'] = None
        super(Frontdoor, __self__).__init__(
            'azure:frontdoor/frontdoor:Frontdoor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_pool_health_probes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolHealthProbeArgs']]]]] = None,
            backend_pool_load_balancings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolLoadBalancingArgs']]]]] = None,
            backend_pools: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolArgs']]]]] = None,
            backend_pools_send_receive_timeout_seconds: Optional[pulumi.Input[float]] = None,
            cname: Optional[pulumi.Input[str]] = None,
            enforce_backend_pools_certificate_name_check: Optional[pulumi.Input[bool]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            frontend_endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorFrontendEndpointArgs']]]]] = None,
            header_frontdoor_id: Optional[pulumi.Input[str]] = None,
            load_balancer_enabled: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            routing_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorRoutingRuleArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Frontdoor':
        """
        Get an existing Frontdoor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolHealthProbeArgs']]]] backend_pool_health_probes: A `backend_pool_health_probe` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolLoadBalancingArgs']]]] backend_pool_load_balancings: A `backend_pool_load_balancing` block as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorBackendPoolArgs']]]] backend_pools: A `backend_pool` block as defined below.
        :param pulumi.Input[float] backend_pools_send_receive_timeout_seconds: Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
        :param pulumi.Input[str] cname: The host that each frontendEndpoint must CNAME to.
        :param pulumi.Input[bool] enforce_backend_pools_certificate_name_check: Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
        :param pulumi.Input[str] friendly_name: A friendly name for the Front Door service.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorFrontendEndpointArgs']]]] frontend_endpoints: A `frontend_endpoint` block as defined below.
        :param pulumi.Input[str] header_frontdoor_id: The unique ID of the Front Door which is embedded into the incoming headers `X-Azure-FDID` attribute and maybe used to filter traffic sent by the Front Door to your backend.
        :param pulumi.Input[bool] load_balancer_enabled: Should the Front Door Load Balancer be Enabled? Defaults to `true`.
        :param pulumi.Input[str] name: Specifies the name of the Front Door service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FrontdoorRoutingRuleArgs']]]] routing_rules: A `routing_rule` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend_pool_health_probes"] = backend_pool_health_probes
        __props__["backend_pool_load_balancings"] = backend_pool_load_balancings
        __props__["backend_pools"] = backend_pools
        __props__["backend_pools_send_receive_timeout_seconds"] = backend_pools_send_receive_timeout_seconds
        __props__["cname"] = cname
        __props__["enforce_backend_pools_certificate_name_check"] = enforce_backend_pools_certificate_name_check
        __props__["friendly_name"] = friendly_name
        __props__["frontend_endpoints"] = frontend_endpoints
        __props__["header_frontdoor_id"] = header_frontdoor_id
        __props__["load_balancer_enabled"] = load_balancer_enabled
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["routing_rules"] = routing_rules
        __props__["tags"] = tags
        return Frontdoor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendPoolHealthProbes")
    def backend_pool_health_probes(self) -> List['outputs.FrontdoorBackendPoolHealthProbe']:
        """
        A `backend_pool_health_probe` block as defined below.
        """
        return pulumi.get(self, "backend_pool_health_probes")

    @property
    @pulumi.getter(name="backendPoolLoadBalancings")
    def backend_pool_load_balancings(self) -> List['outputs.FrontdoorBackendPoolLoadBalancing']:
        """
        A `backend_pool_load_balancing` block as defined below.
        """
        return pulumi.get(self, "backend_pool_load_balancings")

    @property
    @pulumi.getter(name="backendPools")
    def backend_pools(self) -> List['outputs.FrontdoorBackendPool']:
        """
        A `backend_pool` block as defined below.
        """
        return pulumi.get(self, "backend_pools")

    @property
    @pulumi.getter(name="backendPoolsSendReceiveTimeoutSeconds")
    def backend_pools_send_receive_timeout_seconds(self) -> Optional[float]:
        """
        Specifies the send and receive timeout on forwarding request to the backend. When the timeout is reached, the request fails and returns. Possible values are between `0` - `240`. Defaults to `60`.
        """
        return pulumi.get(self, "backend_pools_send_receive_timeout_seconds")

    @property
    @pulumi.getter
    def cname(self) -> str:
        """
        The host that each frontendEndpoint must CNAME to.
        """
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="enforceBackendPoolsCertificateNameCheck")
    def enforce_backend_pools_certificate_name_check(self) -> bool:
        """
        Enforce certificate name check on `HTTPS` requests to all backend pools, this setting will have no effect on `HTTP` requests. Permitted values are `true` or `false`.
        """
        return pulumi.get(self, "enforce_backend_pools_certificate_name_check")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[str]:
        """
        A friendly name for the Front Door service.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="frontendEndpoints")
    def frontend_endpoints(self) -> List['outputs.FrontdoorFrontendEndpoint']:
        """
        A `frontend_endpoint` block as defined below.
        """
        return pulumi.get(self, "frontend_endpoints")

    @property
    @pulumi.getter(name="headerFrontdoorId")
    def header_frontdoor_id(self) -> str:
        """
        The unique ID of the Front Door which is embedded into the incoming headers `X-Azure-FDID` attribute and maybe used to filter traffic sent by the Front Door to your backend.
        """
        return pulumi.get(self, "header_frontdoor_id")

    @property
    @pulumi.getter(name="loadBalancerEnabled")
    def load_balancer_enabled(self) -> Optional[bool]:
        """
        Should the Front Door Load Balancer be Enabled? Defaults to `true`.
        """
        return pulumi.get(self, "load_balancer_enabled")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the Front Door service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        Specifies the name of the Resource Group in which the Front Door service should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="routingRules")
    def routing_rules(self) -> List['outputs.FrontdoorRoutingRule']:
        """
        A `routing_rule` block as defined below.
        """
        return pulumi.get(self, "routing_rules")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

