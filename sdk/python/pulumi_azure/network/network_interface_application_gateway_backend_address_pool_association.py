# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs', 'NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation']

@pulumi.input_type
class NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs:
    def __init__(__self__, *,
                 backend_address_pool_id: pulumi.Input[str],
                 ip_configuration_name: pulumi.Input[str],
                 network_interface_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resource.
        :param pulumi.Input[str] backend_address_pool_id: The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ip_configuration_name: The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_interface_id: The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "backend_address_pool_id", backend_address_pool_id)
        pulumi.set(__self__, "ip_configuration_name", ip_configuration_name)
        pulumi.set(__self__, "network_interface_id", network_interface_id)

    @property
    @pulumi.getter(name="backendAddressPoolId")
    def backend_address_pool_id(self) -> pulumi.Input[str]:
        """
        The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_address_pool_id")

    @backend_address_pool_id.setter
    def backend_address_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend_address_pool_id", value)

    @property
    @pulumi.getter(name="ipConfigurationName")
    def ip_configuration_name(self) -> pulumi.Input[str]:
        """
        The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ip_configuration_name")

    @ip_configuration_name.setter
    def ip_configuration_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_configuration_name", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Input[str]:
        """
        The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_interface_id", value)


@pulumi.input_type
class _NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState:
    def __init__(__self__, *,
                 backend_address_pool_id: Optional[pulumi.Input[str]] = None,
                 ip_configuration_name: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resources.
        :param pulumi.Input[str] backend_address_pool_id: The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ip_configuration_name: The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_interface_id: The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        if backend_address_pool_id is not None:
            pulumi.set(__self__, "backend_address_pool_id", backend_address_pool_id)
        if ip_configuration_name is not None:
            pulumi.set(__self__, "ip_configuration_name", ip_configuration_name)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)

    @property
    @pulumi.getter(name="backendAddressPoolId")
    def backend_address_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_address_pool_id")

    @backend_address_pool_id.setter
    def backend_address_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend_address_pool_id", value)

    @property
    @pulumi.getter(name="ipConfigurationName")
    def ip_configuration_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ip_configuration_name")

    @ip_configuration_name.setter
    def ip_configuration_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_configuration_name", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)


class NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_address_pool_id: Optional[pulumi.Input[str]] = None,
                 ip_configuration_name: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages the association between a Network Interface and a Application Gateway's Backend Address Pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        frontend = azure.network.Subnet("frontend",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"])
        backend = azure.network.Subnet("backend",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Dynamic")
        backend_address_pool_name = example_virtual_network.name.apply(lambda name: f"{name}-beap")
        frontend_port_name = example_virtual_network.name.apply(lambda name: f"{name}-feport")
        frontend_ip_configuration_name = example_virtual_network.name.apply(lambda name: f"{name}-feip")
        http_setting_name = example_virtual_network.name.apply(lambda name: f"{name}-be-htst")
        listener_name = example_virtual_network.name.apply(lambda name: f"{name}-httplstn")
        request_routing_rule_name = example_virtual_network.name.apply(lambda name: f"{name}-rqrt")
        network = azure.network.ApplicationGateway("network",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.network.ApplicationGatewaySkuArgs(
                name="Standard_Small",
                tier="Standard",
                capacity=2,
            ),
            gateway_ip_configurations=[azure.network.ApplicationGatewayGatewayIpConfigurationArgs(
                name="my-gateway-ip-configuration",
                subnet_id=frontend.id,
            )],
            frontend_ports=[azure.network.ApplicationGatewayFrontendPortArgs(
                name=frontend_port_name,
                port=80,
            )],
            frontend_ip_configurations=[azure.network.ApplicationGatewayFrontendIpConfigurationArgs(
                name=frontend_ip_configuration_name,
                public_ip_address_id=example_public_ip.id,
            )],
            backend_address_pools=[azure.network.ApplicationGatewayBackendAddressPoolArgs(
                name=backend_address_pool_name,
            )],
            backend_http_settings=[azure.network.ApplicationGatewayBackendHttpSettingArgs(
                name=http_setting_name,
                cookie_based_affinity="Disabled",
                port=80,
                protocol="Http",
                request_timeout=1,
            )],
            http_listeners=[azure.network.ApplicationGatewayHttpListenerArgs(
                name=listener_name,
                frontend_ip_configuration_name=frontend_ip_configuration_name,
                frontend_port_name=frontend_port_name,
                protocol="Http",
            )],
            request_routing_rules=[azure.network.ApplicationGatewayRequestRoutingRuleArgs(
                name=request_routing_rule_name,
                rule_type="Basic",
                http_listener_name=listener_name,
                backend_address_pool_name=backend_address_pool_name,
                backend_http_settings_name=http_setting_name,
            )])
        example_network_interface = azure.network.NetworkInterface("exampleNetworkInterface",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.NetworkInterfaceIpConfigurationArgs(
                name="testconfiguration1",
                subnet_id=frontend.id,
                private_ip_address_allocation="Dynamic",
            )])
        example_network_interface_application_gateway_backend_address_pool_association = azure.network.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation("exampleNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation",
            network_interface_id=example_network_interface.id,
            ip_configuration_name="testconfiguration1",
            backend_address_pool_id=network.backend_address_pools[0].id)
        ```

        ## Import

        Associations between Network Interfaces and Application Gateway Backend Address Pools can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation association1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationGateways/gateway1/backendAddressPools/pool1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend_address_pool_id: The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ip_configuration_name: The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_interface_id: The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the association between a Network Interface and a Application Gateway's Backend Address Pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        frontend = azure.network.Subnet("frontend",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"])
        backend = azure.network.Subnet("backend",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"])
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Dynamic")
        backend_address_pool_name = example_virtual_network.name.apply(lambda name: f"{name}-beap")
        frontend_port_name = example_virtual_network.name.apply(lambda name: f"{name}-feport")
        frontend_ip_configuration_name = example_virtual_network.name.apply(lambda name: f"{name}-feip")
        http_setting_name = example_virtual_network.name.apply(lambda name: f"{name}-be-htst")
        listener_name = example_virtual_network.name.apply(lambda name: f"{name}-httplstn")
        request_routing_rule_name = example_virtual_network.name.apply(lambda name: f"{name}-rqrt")
        network = azure.network.ApplicationGateway("network",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.network.ApplicationGatewaySkuArgs(
                name="Standard_Small",
                tier="Standard",
                capacity=2,
            ),
            gateway_ip_configurations=[azure.network.ApplicationGatewayGatewayIpConfigurationArgs(
                name="my-gateway-ip-configuration",
                subnet_id=frontend.id,
            )],
            frontend_ports=[azure.network.ApplicationGatewayFrontendPortArgs(
                name=frontend_port_name,
                port=80,
            )],
            frontend_ip_configurations=[azure.network.ApplicationGatewayFrontendIpConfigurationArgs(
                name=frontend_ip_configuration_name,
                public_ip_address_id=example_public_ip.id,
            )],
            backend_address_pools=[azure.network.ApplicationGatewayBackendAddressPoolArgs(
                name=backend_address_pool_name,
            )],
            backend_http_settings=[azure.network.ApplicationGatewayBackendHttpSettingArgs(
                name=http_setting_name,
                cookie_based_affinity="Disabled",
                port=80,
                protocol="Http",
                request_timeout=1,
            )],
            http_listeners=[azure.network.ApplicationGatewayHttpListenerArgs(
                name=listener_name,
                frontend_ip_configuration_name=frontend_ip_configuration_name,
                frontend_port_name=frontend_port_name,
                protocol="Http",
            )],
            request_routing_rules=[azure.network.ApplicationGatewayRequestRoutingRuleArgs(
                name=request_routing_rule_name,
                rule_type="Basic",
                http_listener_name=listener_name,
                backend_address_pool_name=backend_address_pool_name,
                backend_http_settings_name=http_setting_name,
            )])
        example_network_interface = azure.network.NetworkInterface("exampleNetworkInterface",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.NetworkInterfaceIpConfigurationArgs(
                name="testconfiguration1",
                subnet_id=frontend.id,
                private_ip_address_allocation="Dynamic",
            )])
        example_network_interface_application_gateway_backend_address_pool_association = azure.network.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation("exampleNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation",
            network_interface_id=example_network_interface.id,
            ip_configuration_name="testconfiguration1",
            backend_address_pool_id=network.backend_address_pools[0].id)
        ```

        ## Import

        Associations between Network Interfaces and Application Gateway Backend Address Pools can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation association1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationGateways/gateway1/backendAddressPools/pool1
        ```

        :param str resource_name: The name of the resource.
        :param NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_address_pool_id: Optional[pulumi.Input[str]] = None,
                 ip_configuration_name: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs.__new__(NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs)

            if backend_address_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'backend_address_pool_id'")
            __props__.__dict__["backend_address_pool_id"] = backend_address_pool_id
            if ip_configuration_name is None and not opts.urn:
                raise TypeError("Missing required property 'ip_configuration_name'")
            __props__.__dict__["ip_configuration_name"] = ip_configuration_name
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__.__dict__["network_interface_id"] = network_interface_id
        super(NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, __self__).__init__(
            'azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_address_pool_id: Optional[pulumi.Input[str]] = None,
            ip_configuration_name: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None) -> 'NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation':
        """
        Get an existing NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend_address_pool_id: The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ip_configuration_name: The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_interface_id: The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState.__new__(_NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState)

        __props__.__dict__["backend_address_pool_id"] = backend_address_pool_id
        __props__.__dict__["ip_configuration_name"] = ip_configuration_name
        __props__.__dict__["network_interface_id"] = network_interface_id
        return NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendAddressPoolId")
    def backend_address_pool_id(self) -> pulumi.Output[str]:
        """
        The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "backend_address_pool_id")

    @property
    @pulumi.getter(name="ipConfigurationName")
    def ip_configuration_name(self) -> pulumi.Output[str]:
        """
        The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ip_configuration_name")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The ID of the Network Interface. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_interface_id")

