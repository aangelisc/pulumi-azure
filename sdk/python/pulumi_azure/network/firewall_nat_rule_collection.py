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

__all__ = ['FirewallNatRuleCollection']


class FirewallNatRuleCollection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 azure_firewall_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallNatRuleCollectionRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a NAT Rule Collection within an Azure Firewall.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="North Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefix="10.0.1.0/24")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            allocation_method="Static",
            sku="Standard")
        example_firewall = azure.network.Firewall("exampleFirewall",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            ip_configurations=[azure.network.FirewallIpConfigurationArgs(
                name="configuration",
                subnet_id=example_subnet.id,
                public_ip_address_id=example_public_ip.id,
            )])
        example_firewall_nat_rule_collection = azure.network.FirewallNatRuleCollection("exampleFirewallNatRuleCollection",
            azure_firewall_name=example_firewall.name,
            resource_group_name=example_resource_group.name,
            priority=100,
            action="Dnat",
            rules=[azure.network.FirewallNatRuleCollectionRuleArgs(
                name="testrule",
                source_addresses=["10.0.0.0/16"],
                destination_ports=["53"],
                destination_addresses=[example_public_ip.ip_address],
                translated_port="53",
                translated_address="8.8.8.8",
                protocols=[
                    "TCP",
                    "UDP",
                ],
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies the action the rule will apply to matching traffic. Possible values are `Dnat` and `Snat`.
        :param pulumi.Input[str] azure_firewall_name: Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[float] priority: Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallNatRuleCollectionRuleArgs']]]] rules: One or more `rule` blocks as defined below.
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

            if action is None:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            if azure_firewall_name is None:
                raise TypeError("Missing required property 'azure_firewall_name'")
            __props__['azure_firewall_name'] = azure_firewall_name
            __props__['name'] = name
            if priority is None:
                raise TypeError("Missing required property 'priority'")
            __props__['priority'] = priority
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if rules is None:
                raise TypeError("Missing required property 'rules'")
            __props__['rules'] = rules
        super(FirewallNatRuleCollection, __self__).__init__(
            'azure:network/firewallNatRuleCollection:FirewallNatRuleCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            azure_firewall_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallNatRuleCollectionRuleArgs']]]]] = None) -> 'FirewallNatRuleCollection':
        """
        Get an existing FirewallNatRuleCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Specifies the action the rule will apply to matching traffic. Possible values are `Dnat` and `Snat`.
        :param pulumi.Input[str] azure_firewall_name: Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        :param pulumi.Input[float] priority: Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallNatRuleCollectionRuleArgs']]]] rules: One or more `rule` blocks as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["azure_firewall_name"] = azure_firewall_name
        __props__["name"] = name
        __props__["priority"] = priority
        __props__["resource_group_name"] = resource_group_name
        __props__["rules"] = rules
        return FirewallNatRuleCollection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Specifies the action the rule will apply to matching traffic. Possible values are `Dnat` and `Snat`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="azureFirewallName")
    def azure_firewall_name(self) -> str:
        """
        Specifies the name of the Firewall in which the NAT Rule Collection should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "azure_firewall_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the NAT Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> float:
        """
        Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def rules(self) -> List['outputs.FirewallNatRuleCollectionRule']:
        """
        One or more `rule` blocks as defined below.
        """
        return pulumi.get(self, "rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

