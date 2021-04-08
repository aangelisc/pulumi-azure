# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SlotVirtualNetworkSwiftConnectionArgs', 'SlotVirtualNetworkSwiftConnection']

@pulumi.input_type
class SlotVirtualNetworkSwiftConnectionArgs:
    def __init__(__self__, *,
                 app_service_id: pulumi.Input[str],
                 slot_name: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a SlotVirtualNetworkSwiftConnection resource.
        :param pulumi.Input[str] app_service_id: The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] slot_name: The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        pulumi.set(__self__, "app_service_id", app_service_id)
        pulumi.set(__self__, "slot_name", slot_name)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="appServiceId")
    def app_service_id(self) -> pulumi.Input[str]:
        """
        The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_service_id")

    @app_service_id.setter
    def app_service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_service_id", value)

    @property
    @pulumi.getter(name="slotName")
    def slot_name(self) -> pulumi.Input[str]:
        """
        The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "slot_name")

    @slot_name.setter
    def slot_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "slot_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class _SlotVirtualNetworkSwiftConnectionState:
    def __init__(__self__, *,
                 app_service_id: Optional[pulumi.Input[str]] = None,
                 slot_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SlotVirtualNetworkSwiftConnection resources.
        :param pulumi.Input[str] app_service_id: The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] slot_name: The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        if app_service_id is not None:
            pulumi.set(__self__, "app_service_id", app_service_id)
        if slot_name is not None:
            pulumi.set(__self__, "slot_name", slot_name)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="appServiceId")
    def app_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_service_id")

    @app_service_id.setter
    def app_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_service_id", value)

    @property
    @pulumi.getter(name="slotName")
    def slot_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "slot_name")

    @slot_name.setter
    def slot_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slot_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class SlotVirtualNetworkSwiftConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service_id: Optional[pulumi.Input[str]] = None,
                 slot_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an App Service Slot's Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"],
            delegations=[azure.network.SubnetDelegationArgs(
                name="example-delegation",
                service_delegation=azure.network.SubnetDelegationServiceDelegationArgs(
                    name="Microsoft.Web/serverFarms",
                    actions=["Microsoft.Network/virtualNetworks/subnets/action"],
                ),
            )])
        example_plan = azure.appservice.Plan("examplePlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=azure.appservice.PlanSkuArgs(
                tier="Standard",
                size="S1",
            ))
        example_app_service = azure.appservice.AppService("exampleAppService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id)
        example_staging = azure.appservice.Slot("example-staging",
            app_service_name=example_app_service.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id)
        example_slot_virtual_network_swift_connection = azure.appservice.SlotVirtualNetworkSwiftConnection("exampleSlotVirtualNetworkSwiftConnection",
            slot_name=example_staging.name,
            app_service_id=example_app_service.id,
            subnet_id=example_subnet.id)
        ```

        ## Import

        App Service Slot Virtual Network Associations can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection myassociation /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/stageing/config/virtualNetwork
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_id: The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] slot_name: The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SlotVirtualNetworkSwiftConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an App Service Slot's Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration) which is still in preview).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"],
            delegations=[azure.network.SubnetDelegationArgs(
                name="example-delegation",
                service_delegation=azure.network.SubnetDelegationServiceDelegationArgs(
                    name="Microsoft.Web/serverFarms",
                    actions=["Microsoft.Network/virtualNetworks/subnets/action"],
                ),
            )])
        example_plan = azure.appservice.Plan("examplePlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=azure.appservice.PlanSkuArgs(
                tier="Standard",
                size="S1",
            ))
        example_app_service = azure.appservice.AppService("exampleAppService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id)
        example_staging = azure.appservice.Slot("example-staging",
            app_service_name=example_app_service.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id)
        example_slot_virtual_network_swift_connection = azure.appservice.SlotVirtualNetworkSwiftConnection("exampleSlotVirtualNetworkSwiftConnection",
            slot_name=example_staging.name,
            app_service_id=example_app_service.id,
            subnet_id=example_subnet.id)
        ```

        ## Import

        App Service Slot Virtual Network Associations can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection myassociation /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/stageing/config/virtualNetwork
        ```

        :param str resource_name: The name of the resource.
        :param SlotVirtualNetworkSwiftConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SlotVirtualNetworkSwiftConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service_id: Optional[pulumi.Input[str]] = None,
                 slot_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SlotVirtualNetworkSwiftConnectionArgs.__new__(SlotVirtualNetworkSwiftConnectionArgs)

            if app_service_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_service_id'")
            __props__.__dict__["app_service_id"] = app_service_id
            if slot_name is None and not opts.urn:
                raise TypeError("Missing required property 'slot_name'")
            __props__.__dict__["slot_name"] = slot_name
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
        super(SlotVirtualNetworkSwiftConnection, __self__).__init__(
            'azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service_id: Optional[pulumi.Input[str]] = None,
            slot_name: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'SlotVirtualNetworkSwiftConnection':
        """
        Get an existing SlotVirtualNetworkSwiftConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_id: The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        :param pulumi.Input[str] slot_name: The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SlotVirtualNetworkSwiftConnectionState.__new__(_SlotVirtualNetworkSwiftConnectionState)

        __props__.__dict__["app_service_id"] = app_service_id
        __props__.__dict__["slot_name"] = slot_name
        __props__.__dict__["subnet_id"] = subnet_id
        return SlotVirtualNetworkSwiftConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appServiceId")
    def app_service_id(self) -> pulumi.Output[str]:
        """
        The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_service_id")

    @property
    @pulumi.getter(name="slotName")
    def slot_name(self) -> pulumi.Output[str]:
        """
        The name of the App Service Slot or Function App Slot. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "slot_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the subnet the app service will be associated to (the subnet must have a `service_delegation` configured for `Microsoft.Web/serverFarms`).
        """
        return pulumi.get(self, "subnet_id")

