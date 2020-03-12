# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OutboundRule(pulumi.CustomResource):
    allocated_outbound_ports: pulumi.Output[float]
    """
    The number of outbound ports to be used for NAT.
    """
    backend_address_pool_id: pulumi.Output[str]
    """
    The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
    """
    enable_tcp_reset: pulumi.Output[bool]
    """
    Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
    """
    frontend_ip_configurations: pulumi.Output[list]
    """
    One or more `frontend_ip_configuration` blocks as defined below.

      * `id` (`str`) - The ID of the Load Balancer Outbound Rule.
      * `name` (`str`) - Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
    """
    idle_timeout_in_minutes: pulumi.Output[float]
    """
    The timeout for the TCP idle connection
    """
    loadbalancer_id: pulumi.Output[str]
    """
    The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
    """
    protocol: pulumi.Output[str]
    """
    The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, allocated_outbound_ports=None, backend_address_pool_id=None, enable_tcp_reset=None, frontend_ip_configurations=None, idle_timeout_in_minutes=None, loadbalancer_id=None, name=None, protocol=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Load Balancer Outbound Rule.

        > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration and a Backend Address Pool Attached.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb_outbound_rule.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] allocated_outbound_ports: The number of outbound ports to be used for NAT.
        :param pulumi.Input[str] backend_address_pool_id: The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        :param pulumi.Input[bool] enable_tcp_reset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        :param pulumi.Input[list] frontend_ip_configurations: One or more `frontend_ip_configuration` blocks as defined below.
        :param pulumi.Input[float] idle_timeout_in_minutes: The timeout for the TCP idle connection
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protocol: The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.

        The **frontend_ip_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ID of the Load Balancer Outbound Rule.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
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

            __props__['allocated_outbound_ports'] = allocated_outbound_ports
            if backend_address_pool_id is None:
                raise TypeError("Missing required property 'backend_address_pool_id'")
            __props__['backend_address_pool_id'] = backend_address_pool_id
            __props__['enable_tcp_reset'] = enable_tcp_reset
            __props__['frontend_ip_configurations'] = frontend_ip_configurations
            __props__['idle_timeout_in_minutes'] = idle_timeout_in_minutes
            if loadbalancer_id is None:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__['loadbalancer_id'] = loadbalancer_id
            __props__['name'] = name
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(OutboundRule, __self__).__init__(
            'azure:lb/outboundRule:OutboundRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allocated_outbound_ports=None, backend_address_pool_id=None, enable_tcp_reset=None, frontend_ip_configurations=None, idle_timeout_in_minutes=None, loadbalancer_id=None, name=None, protocol=None, resource_group_name=None):
        """
        Get an existing OutboundRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] allocated_outbound_ports: The number of outbound ports to be used for NAT.
        :param pulumi.Input[str] backend_address_pool_id: The ID of the Backend Address Pool. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        :param pulumi.Input[bool] enable_tcp_reset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        :param pulumi.Input[list] frontend_ip_configurations: One or more `frontend_ip_configuration` blocks as defined below.
        :param pulumi.Input[float] idle_timeout_in_minutes: The timeout for the TCP idle connection
        :param pulumi.Input[str] loadbalancer_id: The ID of the Load Balancer in which to create the Outbound Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protocol: The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the resource. Changing this forces a new resource to be created.

        The **frontend_ip_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ID of the Load Balancer Outbound Rule.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the Outbound Rule. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocated_outbound_ports"] = allocated_outbound_ports
        __props__["backend_address_pool_id"] = backend_address_pool_id
        __props__["enable_tcp_reset"] = enable_tcp_reset
        __props__["frontend_ip_configurations"] = frontend_ip_configurations
        __props__["idle_timeout_in_minutes"] = idle_timeout_in_minutes
        __props__["loadbalancer_id"] = loadbalancer_id
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["resource_group_name"] = resource_group_name
        return OutboundRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

