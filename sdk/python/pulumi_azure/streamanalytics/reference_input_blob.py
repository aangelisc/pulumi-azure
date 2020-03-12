# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ReferenceInputBlob(pulumi.CustomResource):
    date_format: pulumi.Output[str]
    """
    The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
    """
    name: pulumi.Output[str]
    """
    The name of the Reference Input Blob. Changing this forces a new resource to be created.
    """
    path_pattern: pulumi.Output[str]
    """
    The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
    """
    serialization: pulumi.Output[dict]
    """
    A `serialization` block as defined below.

      * `encoding` (`str`)
      * `fieldDelimiter` (`str`)
      * `type` (`str`)
    """
    storage_account_key: pulumi.Output[str]
    """
    The Access Key which should be used to connect to this Storage Account.
    """
    storage_account_name: pulumi.Output[str]
    """
    The name of the Storage Account that has the blob container with reference data.
    """
    storage_container_name: pulumi.Output[str]
    """
    The name of the Container within the Storage Account.
    """
    stream_analytics_job_name: pulumi.Output[str]
    """
    The name of the Stream Analytics Job. Changing this forces a new resource to be created.
    """
    time_format: pulumi.Output[str]
    """
    The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
    """
    def __init__(__self__, resource_name, opts=None, date_format=None, name=None, path_pattern=None, resource_group_name=None, serialization=None, storage_account_key=None, storage_account_name=None, storage_container_name=None, stream_analytics_job_name=None, time_format=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Stream Analytics Reference Input Blob. Reference data (also known as a lookup table) is a finite data set that is static or slowly changing in nature, used to perform a lookup or to correlate with your data stream. Learn more [here](https://docs.microsoft.com/en-us/azure/stream-analytics/stream-analytics-use-reference-data#azure-blob-storage).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_reference_input_blob.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Reference Input Blob. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account that has the blob container with reference data.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.

        The **serialization** object supports the following:

          * `encoding` (`pulumi.Input[str]`)
          * `fieldDelimiter` (`pulumi.Input[str]`)
          * `type` (`pulumi.Input[str]`)
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

            if date_format is None:
                raise TypeError("Missing required property 'date_format'")
            __props__['date_format'] = date_format
            __props__['name'] = name
            if path_pattern is None:
                raise TypeError("Missing required property 'path_pattern'")
            __props__['path_pattern'] = path_pattern
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if serialization is None:
                raise TypeError("Missing required property 'serialization'")
            __props__['serialization'] = serialization
            if storage_account_key is None:
                raise TypeError("Missing required property 'storage_account_key'")
            __props__['storage_account_key'] = storage_account_key
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if storage_container_name is None:
                raise TypeError("Missing required property 'storage_container_name'")
            __props__['storage_container_name'] = storage_container_name
            if stream_analytics_job_name is None:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
            if time_format is None:
                raise TypeError("Missing required property 'time_format'")
            __props__['time_format'] = time_format
        super(ReferenceInputBlob, __self__).__init__(
            'azure:streamanalytics/referenceInputBlob:ReferenceInputBlob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, date_format=None, name=None, path_pattern=None, resource_group_name=None, serialization=None, storage_account_key=None, storage_account_name=None, storage_container_name=None, stream_analytics_job_name=None, time_format=None):
        """
        Get an existing ReferenceInputBlob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Reference Input Blob. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account that has the blob container with reference data.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.

        The **serialization** object supports the following:

          * `encoding` (`pulumi.Input[str]`)
          * `fieldDelimiter` (`pulumi.Input[str]`)
          * `type` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["date_format"] = date_format
        __props__["name"] = name
        __props__["path_pattern"] = path_pattern
        __props__["resource_group_name"] = resource_group_name
        __props__["serialization"] = serialization
        __props__["storage_account_key"] = storage_account_key
        __props__["storage_account_name"] = storage_account_name
        __props__["storage_container_name"] = storage_container_name
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        __props__["time_format"] = time_format
        return ReferenceInputBlob(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
