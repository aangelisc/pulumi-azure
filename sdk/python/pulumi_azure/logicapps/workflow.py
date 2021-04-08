# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WorkflowArgs', 'Workflow']

@pulumi.input_type
class WorkflowArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 integration_service_environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workflow_schema: Optional[pulumi.Input[str]] = None,
                 workflow_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Workflow resource.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] integration_service_environment_id: The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] logic_app_integration_account_id: The ID of the integration account linked by this Logic App Workflow.
        :param pulumi.Input[str] name: Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of Key-Value pairs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workflow_schema: Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workflow_version: Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if integration_service_environment_id is not None:
            pulumi.set(__self__, "integration_service_environment_id", integration_service_environment_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if logic_app_integration_account_id is not None:
            pulumi.set(__self__, "logic_app_integration_account_id", logic_app_integration_account_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if workflow_schema is not None:
            pulumi.set(__self__, "workflow_schema", workflow_schema)
        if workflow_version is not None:
            pulumi.set(__self__, "workflow_version", workflow_version)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="integrationServiceEnvironmentId")
    def integration_service_environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        """
        return pulumi.get(self, "integration_service_environment_id")

    @integration_service_environment_id.setter
    def integration_service_environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_service_environment_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logicAppIntegrationAccountId")
    def logic_app_integration_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the integration account linked by this Logic App Workflow.
        """
        return pulumi.get(self, "logic_app_integration_account_id")

    @logic_app_integration_account_id.setter
    def logic_app_integration_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic_app_integration_account_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of Key-Value pairs.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

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

    @property
    @pulumi.getter(name="workflowSchema")
    def workflow_schema(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_schema")

    @workflow_schema.setter
    def workflow_schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workflow_schema", value)

    @property
    @pulumi.getter(name="workflowVersion")
    def workflow_version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_version")

    @workflow_version.setter
    def workflow_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workflow_version", value)


@pulumi.input_type
class _WorkflowState:
    def __init__(__self__, *,
                 access_endpoint: Optional[pulumi.Input[str]] = None,
                 connector_endpoint_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connector_outbound_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 integration_service_environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workflow_endpoint_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workflow_outbound_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workflow_schema: Optional[pulumi.Input[str]] = None,
                 workflow_version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Workflow resources.
        :param pulumi.Input[str] access_endpoint: The Access Endpoint for the Logic App Workflow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connector_endpoint_ip_addresses: The list of access endpoint ip addresses of connector.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connector_outbound_ip_addresses: The list of outgoing ip addresses of connector.
        :param pulumi.Input[str] integration_service_environment_id: The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] logic_app_integration_account_id: The ID of the integration account linked by this Logic App Workflow.
        :param pulumi.Input[str] name: Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of Key-Value pairs.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workflow_endpoint_ip_addresses: The list of access endpoint ip addresses of workflow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workflow_outbound_ip_addresses: The list of outgoing ip addresses of workflow.
        :param pulumi.Input[str] workflow_schema: Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workflow_version: Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        if access_endpoint is not None:
            pulumi.set(__self__, "access_endpoint", access_endpoint)
        if connector_endpoint_ip_addresses is not None:
            pulumi.set(__self__, "connector_endpoint_ip_addresses", connector_endpoint_ip_addresses)
        if connector_outbound_ip_addresses is not None:
            pulumi.set(__self__, "connector_outbound_ip_addresses", connector_outbound_ip_addresses)
        if integration_service_environment_id is not None:
            pulumi.set(__self__, "integration_service_environment_id", integration_service_environment_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if logic_app_integration_account_id is not None:
            pulumi.set(__self__, "logic_app_integration_account_id", logic_app_integration_account_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if workflow_endpoint_ip_addresses is not None:
            pulumi.set(__self__, "workflow_endpoint_ip_addresses", workflow_endpoint_ip_addresses)
        if workflow_outbound_ip_addresses is not None:
            pulumi.set(__self__, "workflow_outbound_ip_addresses", workflow_outbound_ip_addresses)
        if workflow_schema is not None:
            pulumi.set(__self__, "workflow_schema", workflow_schema)
        if workflow_version is not None:
            pulumi.set(__self__, "workflow_version", workflow_version)

    @property
    @pulumi.getter(name="accessEndpoint")
    def access_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The Access Endpoint for the Logic App Workflow.
        """
        return pulumi.get(self, "access_endpoint")

    @access_endpoint.setter
    def access_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_endpoint", value)

    @property
    @pulumi.getter(name="connectorEndpointIpAddresses")
    def connector_endpoint_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of access endpoint ip addresses of connector.
        """
        return pulumi.get(self, "connector_endpoint_ip_addresses")

    @connector_endpoint_ip_addresses.setter
    def connector_endpoint_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "connector_endpoint_ip_addresses", value)

    @property
    @pulumi.getter(name="connectorOutboundIpAddresses")
    def connector_outbound_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of outgoing ip addresses of connector.
        """
        return pulumi.get(self, "connector_outbound_ip_addresses")

    @connector_outbound_ip_addresses.setter
    def connector_outbound_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "connector_outbound_ip_addresses", value)

    @property
    @pulumi.getter(name="integrationServiceEnvironmentId")
    def integration_service_environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        """
        return pulumi.get(self, "integration_service_environment_id")

    @integration_service_environment_id.setter
    def integration_service_environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_service_environment_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logicAppIntegrationAccountId")
    def logic_app_integration_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the integration account linked by this Logic App Workflow.
        """
        return pulumi.get(self, "logic_app_integration_account_id")

    @logic_app_integration_account_id.setter
    def logic_app_integration_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic_app_integration_account_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of Key-Value pairs.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
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

    @property
    @pulumi.getter(name="workflowEndpointIpAddresses")
    def workflow_endpoint_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of access endpoint ip addresses of workflow.
        """
        return pulumi.get(self, "workflow_endpoint_ip_addresses")

    @workflow_endpoint_ip_addresses.setter
    def workflow_endpoint_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "workflow_endpoint_ip_addresses", value)

    @property
    @pulumi.getter(name="workflowOutboundIpAddresses")
    def workflow_outbound_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of outgoing ip addresses of workflow.
        """
        return pulumi.get(self, "workflow_outbound_ip_addresses")

    @workflow_outbound_ip_addresses.setter
    def workflow_outbound_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "workflow_outbound_ip_addresses", value)

    @property
    @pulumi.getter(name="workflowSchema")
    def workflow_schema(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_schema")

    @workflow_schema.setter
    def workflow_schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workflow_schema", value)

    @property
    @pulumi.getter(name="workflowVersion")
    def workflow_version(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_version")

    @workflow_version.setter
    def workflow_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workflow_version", value)


class Workflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 integration_service_environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workflow_schema: Optional[pulumi.Input[str]] = None,
                 workflow_version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Logic App Workflow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        ```

        ## Import

        Logic App Workflows can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:logicapps/workflow:Workflow workflow1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_service_environment_id: The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] logic_app_integration_account_id: The ID of the integration account linked by this Logic App Workflow.
        :param pulumi.Input[str] name: Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of Key-Value pairs.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workflow_schema: Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workflow_version: Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkflowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Logic App Workflow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_workflow = azure.logicapps.Workflow("exampleWorkflow",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        ```

        ## Import

        Logic App Workflows can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:logicapps/workflow:Workflow workflow1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
        ```

        :param str resource_name: The name of the resource.
        :param WorkflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 integration_service_environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workflow_schema: Optional[pulumi.Input[str]] = None,
                 workflow_version: Optional[pulumi.Input[str]] = None,
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
            __props__ = WorkflowArgs.__new__(WorkflowArgs)

            __props__.__dict__["integration_service_environment_id"] = integration_service_environment_id
            __props__.__dict__["location"] = location
            __props__.__dict__["logic_app_integration_account_id"] = logic_app_integration_account_id
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["workflow_schema"] = workflow_schema
            __props__.__dict__["workflow_version"] = workflow_version
            __props__.__dict__["access_endpoint"] = None
            __props__.__dict__["connector_endpoint_ip_addresses"] = None
            __props__.__dict__["connector_outbound_ip_addresses"] = None
            __props__.__dict__["workflow_endpoint_ip_addresses"] = None
            __props__.__dict__["workflow_outbound_ip_addresses"] = None
        super(Workflow, __self__).__init__(
            'azure:logicapps/workflow:Workflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_endpoint: Optional[pulumi.Input[str]] = None,
            connector_endpoint_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            connector_outbound_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            integration_service_environment_id: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            logic_app_integration_account_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            workflow_endpoint_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            workflow_outbound_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            workflow_schema: Optional[pulumi.Input[str]] = None,
            workflow_version: Optional[pulumi.Input[str]] = None) -> 'Workflow':
        """
        Get an existing Workflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_endpoint: The Access Endpoint for the Logic App Workflow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connector_endpoint_ip_addresses: The list of access endpoint ip addresses of connector.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connector_outbound_ip_addresses: The list of outgoing ip addresses of connector.
        :param pulumi.Input[str] integration_service_environment_id: The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] logic_app_integration_account_id: The ID of the integration account linked by this Logic App Workflow.
        :param pulumi.Input[str] name: Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of Key-Value pairs.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workflow_endpoint_ip_addresses: The list of access endpoint ip addresses of workflow.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] workflow_outbound_ip_addresses: The list of outgoing ip addresses of workflow.
        :param pulumi.Input[str] workflow_schema: Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] workflow_version: Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkflowState.__new__(_WorkflowState)

        __props__.__dict__["access_endpoint"] = access_endpoint
        __props__.__dict__["connector_endpoint_ip_addresses"] = connector_endpoint_ip_addresses
        __props__.__dict__["connector_outbound_ip_addresses"] = connector_outbound_ip_addresses
        __props__.__dict__["integration_service_environment_id"] = integration_service_environment_id
        __props__.__dict__["location"] = location
        __props__.__dict__["logic_app_integration_account_id"] = logic_app_integration_account_id
        __props__.__dict__["name"] = name
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["workflow_endpoint_ip_addresses"] = workflow_endpoint_ip_addresses
        __props__.__dict__["workflow_outbound_ip_addresses"] = workflow_outbound_ip_addresses
        __props__.__dict__["workflow_schema"] = workflow_schema
        __props__.__dict__["workflow_version"] = workflow_version
        return Workflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessEndpoint")
    def access_endpoint(self) -> pulumi.Output[str]:
        """
        The Access Endpoint for the Logic App Workflow.
        """
        return pulumi.get(self, "access_endpoint")

    @property
    @pulumi.getter(name="connectorEndpointIpAddresses")
    def connector_endpoint_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of access endpoint ip addresses of connector.
        """
        return pulumi.get(self, "connector_endpoint_ip_addresses")

    @property
    @pulumi.getter(name="connectorOutboundIpAddresses")
    def connector_outbound_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of outgoing ip addresses of connector.
        """
        return pulumi.get(self, "connector_outbound_ip_addresses")

    @property
    @pulumi.getter(name="integrationServiceEnvironmentId")
    def integration_service_environment_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        """
        return pulumi.get(self, "integration_service_environment_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logicAppIntegrationAccountId")
    def logic_app_integration_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the integration account linked by this Logic App Workflow.
        """
        return pulumi.get(self, "logic_app_integration_account_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of Key-Value pairs.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workflowEndpointIpAddresses")
    def workflow_endpoint_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of access endpoint ip addresses of workflow.
        """
        return pulumi.get(self, "workflow_endpoint_ip_addresses")

    @property
    @pulumi.getter(name="workflowOutboundIpAddresses")
    def workflow_outbound_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of outgoing ip addresses of workflow.
        """
        return pulumi.get(self, "workflow_outbound_ip_addresses")

    @property
    @pulumi.getter(name="workflowSchema")
    def workflow_schema(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_schema")

    @property
    @pulumi.getter(name="workflowVersion")
    def workflow_version(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "workflow_version")

