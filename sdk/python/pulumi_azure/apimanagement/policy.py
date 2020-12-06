# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Policy']


class Policy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_id: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a API Management service Policy.

        > **NOTE:** This resource will, upon creation, **overwrite any existing policy in the API Management service**, as there is no feasible way to test whether the policy has been modified from the default. Similarly, when this resource is destroyed, the API Management service will revert to its default policy.

        ## Import

        API Management service Policys can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
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

            if api_management_id is None:
                raise TypeError("Missing required property 'api_management_id'")
            __props__['api_management_id'] = api_management_id
            __props__['xml_content'] = xml_content
            __props__['xml_link'] = xml_link
        super(Policy, __self__).__init__(
            'azure:apimanagement/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_id: Optional[pulumi.Input[str]] = None,
            xml_content: Optional[pulumi.Input[str]] = None,
            xml_link: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_id"] = api_management_id
        __props__["xml_content"] = xml_content
        __props__["xml_link"] = xml_link
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementId")
    def api_management_id(self) -> pulumi.Output[str]:
        """
        The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        """
        return pulumi.get(self, "api_management_id")

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> pulumi.Output[str]:
        return pulumi.get(self, "xml_content")

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> pulumi.Output[Optional[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

