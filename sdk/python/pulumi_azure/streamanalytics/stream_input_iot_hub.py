# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['StreamInputIotHubArgs', 'StreamInputIotHub']

@pulumi.input_type
class StreamInputIotHubArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 eventhub_consumer_group_name: pulumi.Input[str],
                 iothub_namespace: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 serialization: pulumi.Input['StreamInputIotHubSerializationArgs'],
                 shared_access_policy_key: pulumi.Input[str],
                 shared_access_policy_name: pulumi.Input[str],
                 stream_analytics_job_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StreamInputIotHub resource.
        :param pulumi.Input[str] endpoint: The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] iothub_namespace: The name or the URI of the IoT Hub.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input['StreamInputIotHubSerializationArgs'] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "eventhub_consumer_group_name", eventhub_consumer_group_name)
        pulumi.set(__self__, "iothub_namespace", iothub_namespace)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "serialization", serialization)
        pulumi.set(__self__, "shared_access_policy_key", shared_access_policy_key)
        pulumi.set(__self__, "shared_access_policy_name", shared_access_policy_name)
        pulumi.set(__self__, "stream_analytics_job_name", stream_analytics_job_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="eventhubConsumerGroupName")
    def eventhub_consumer_group_name(self) -> pulumi.Input[str]:
        """
        The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        """
        return pulumi.get(self, "eventhub_consumer_group_name")

    @eventhub_consumer_group_name.setter
    def eventhub_consumer_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "eventhub_consumer_group_name", value)

    @property
    @pulumi.getter(name="iothubNamespace")
    def iothub_namespace(self) -> pulumi.Input[str]:
        """
        The name or the URI of the IoT Hub.
        """
        return pulumi.get(self, "iothub_namespace")

    @iothub_namespace.setter
    def iothub_namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "iothub_namespace", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def serialization(self) -> pulumi.Input['StreamInputIotHubSerializationArgs']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @serialization.setter
    def serialization(self, value: pulumi.Input['StreamInputIotHubSerializationArgs']):
        pulumi.set(self, "serialization", value)

    @property
    @pulumi.getter(name="sharedAccessPolicyKey")
    def shared_access_policy_key(self) -> pulumi.Input[str]:
        """
        The shared access policy key for the specified shared access policy.
        """
        return pulumi.get(self, "shared_access_policy_key")

    @shared_access_policy_key.setter
    def shared_access_policy_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "shared_access_policy_key", value)

    @property
    @pulumi.getter(name="sharedAccessPolicyName")
    def shared_access_policy_name(self) -> pulumi.Input[str]:
        """
        The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "shared_access_policy_name")

    @shared_access_policy_name.setter
    def shared_access_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "shared_access_policy_name", value)

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Input[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @stream_analytics_job_name.setter
    def stream_analytics_job_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_analytics_job_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _StreamInputIotHubState:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
                 iothub_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input['StreamInputIotHubSerializationArgs']] = None,
                 shared_access_policy_key: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StreamInputIotHub resources.
        :param pulumi.Input[str] endpoint: The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] iothub_namespace: The name or the URI of the IoT Hub.
        :param pulumi.Input[str] name: The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input['StreamInputIotHubSerializationArgs'] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if eventhub_consumer_group_name is not None:
            pulumi.set(__self__, "eventhub_consumer_group_name", eventhub_consumer_group_name)
        if iothub_namespace is not None:
            pulumi.set(__self__, "iothub_namespace", iothub_namespace)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if serialization is not None:
            pulumi.set(__self__, "serialization", serialization)
        if shared_access_policy_key is not None:
            pulumi.set(__self__, "shared_access_policy_key", shared_access_policy_key)
        if shared_access_policy_name is not None:
            pulumi.set(__self__, "shared_access_policy_name", shared_access_policy_name)
        if stream_analytics_job_name is not None:
            pulumi.set(__self__, "stream_analytics_job_name", stream_analytics_job_name)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="eventhubConsumerGroupName")
    def eventhub_consumer_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        """
        return pulumi.get(self, "eventhub_consumer_group_name")

    @eventhub_consumer_group_name.setter
    def eventhub_consumer_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eventhub_consumer_group_name", value)

    @property
    @pulumi.getter(name="iothubNamespace")
    def iothub_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The name or the URI of the IoT Hub.
        """
        return pulumi.get(self, "iothub_namespace")

    @iothub_namespace.setter
    def iothub_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iothub_namespace", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def serialization(self) -> Optional[pulumi.Input['StreamInputIotHubSerializationArgs']]:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @serialization.setter
    def serialization(self, value: Optional[pulumi.Input['StreamInputIotHubSerializationArgs']]):
        pulumi.set(self, "serialization", value)

    @property
    @pulumi.getter(name="sharedAccessPolicyKey")
    def shared_access_policy_key(self) -> Optional[pulumi.Input[str]]:
        """
        The shared access policy key for the specified shared access policy.
        """
        return pulumi.get(self, "shared_access_policy_key")

    @shared_access_policy_key.setter
    def shared_access_policy_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_access_policy_key", value)

    @property
    @pulumi.getter(name="sharedAccessPolicyName")
    def shared_access_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "shared_access_policy_name")

    @shared_access_policy_name.setter
    def shared_access_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_access_policy_name", value)

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @stream_analytics_job_name.setter
    def stream_analytics_job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_analytics_job_name", value)


class StreamInputIotHub(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
                 iothub_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputIotHubSerializationArgs']]] = None,
                 shared_access_policy_key: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Stream Analytics Stream Input IoTHub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=azurerm_resource_group["example"]["name"],
            location=azurerm_resource_group["example"]["location"],
            sku=azure.iot.IoTHubSkuArgs(
                name="S1",
                capacity=1,
            ))
        example_stream_input_iot_hub = azure.streamanalytics.StreamInputIotHub("exampleStreamInputIotHub",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            endpoint="messages/events",
            eventhub_consumer_group_name="$Default",
            iothub_namespace=example_io_t_hub.name,
            shared_access_policy_key=example_io_t_hub.shared_access_policies[0].primary_key,
            shared_access_policy_name="iothubowner",
            serialization=azure.streamanalytics.StreamInputIotHubSerializationArgs(
                type="Json",
                encoding="UTF8",
            ))
        ```

        ## Import

        Stream Analytics Stream Input IoTHub's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/streamInputIotHub:StreamInputIotHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] iothub_namespace: The name or the URI of the IoT Hub.
        :param pulumi.Input[str] name: The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['StreamInputIotHubSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamInputIotHubArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Stream Analytics Stream Input IoTHub.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=azurerm_resource_group["example"]["name"],
            location=azurerm_resource_group["example"]["location"],
            sku=azure.iot.IoTHubSkuArgs(
                name="S1",
                capacity=1,
            ))
        example_stream_input_iot_hub = azure.streamanalytics.StreamInputIotHub("exampleStreamInputIotHub",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            endpoint="messages/events",
            eventhub_consumer_group_name="$Default",
            iothub_namespace=example_io_t_hub.name,
            shared_access_policy_key=example_io_t_hub.shared_access_policies[0].primary_key,
            shared_access_policy_name="iothubowner",
            serialization=azure.streamanalytics.StreamInputIotHubSerializationArgs(
                type="Json",
                encoding="UTF8",
            ))
        ```

        ## Import

        Stream Analytics Stream Input IoTHub's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/streamInputIotHub:StreamInputIotHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
        ```

        :param str resource_name: The name of the resource.
        :param StreamInputIotHubArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamInputIotHubArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
                 iothub_namespace: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputIotHubSerializationArgs']]] = None,
                 shared_access_policy_key: Optional[pulumi.Input[str]] = None,
                 shared_access_policy_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = StreamInputIotHubArgs.__new__(StreamInputIotHubArgs)

            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            if eventhub_consumer_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'eventhub_consumer_group_name'")
            __props__.__dict__["eventhub_consumer_group_name"] = eventhub_consumer_group_name
            if iothub_namespace is None and not opts.urn:
                raise TypeError("Missing required property 'iothub_namespace'")
            __props__.__dict__["iothub_namespace"] = iothub_namespace
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if serialization is None and not opts.urn:
                raise TypeError("Missing required property 'serialization'")
            __props__.__dict__["serialization"] = serialization
            if shared_access_policy_key is None and not opts.urn:
                raise TypeError("Missing required property 'shared_access_policy_key'")
            __props__.__dict__["shared_access_policy_key"] = shared_access_policy_key
            if shared_access_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'shared_access_policy_name'")
            __props__.__dict__["shared_access_policy_name"] = shared_access_policy_name
            if stream_analytics_job_name is None and not opts.urn:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__.__dict__["stream_analytics_job_name"] = stream_analytics_job_name
        super(StreamInputIotHub, __self__).__init__(
            'azure:streamanalytics/streamInputIotHub:StreamInputIotHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            eventhub_consumer_group_name: Optional[pulumi.Input[str]] = None,
            iothub_namespace: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            serialization: Optional[pulumi.Input[pulumi.InputType['StreamInputIotHubSerializationArgs']]] = None,
            shared_access_policy_key: Optional[pulumi.Input[str]] = None,
            shared_access_policy_name: Optional[pulumi.Input[str]] = None,
            stream_analytics_job_name: Optional[pulumi.Input[str]] = None) -> 'StreamInputIotHub':
        """
        Get an existing StreamInputIotHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        :param pulumi.Input[str] eventhub_consumer_group_name: The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        :param pulumi.Input[str] iothub_namespace: The name or the URI of the IoT Hub.
        :param pulumi.Input[str] name: The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['StreamInputIotHubSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StreamInputIotHubState.__new__(_StreamInputIotHubState)

        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["eventhub_consumer_group_name"] = eventhub_consumer_group_name
        __props__.__dict__["iothub_namespace"] = iothub_namespace
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["serialization"] = serialization
        __props__.__dict__["shared_access_policy_key"] = shared_access_policy_key
        __props__.__dict__["shared_access_policy_name"] = shared_access_policy_name
        __props__.__dict__["stream_analytics_job_name"] = stream_analytics_job_name
        return StreamInputIotHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="eventhubConsumerGroupName")
    def eventhub_consumer_group_name(self) -> pulumi.Output[str]:
        """
        The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
        """
        return pulumi.get(self, "eventhub_consumer_group_name")

    @property
    @pulumi.getter(name="iothubNamespace")
    def iothub_namespace(self) -> pulumi.Output[str]:
        """
        The name or the URI of the IoT Hub.
        """
        return pulumi.get(self, "iothub_namespace")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def serialization(self) -> pulumi.Output['outputs.StreamInputIotHubSerialization']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @property
    @pulumi.getter(name="sharedAccessPolicyKey")
    def shared_access_policy_key(self) -> pulumi.Output[str]:
        """
        The shared access policy key for the specified shared access policy.
        """
        return pulumi.get(self, "shared_access_policy_key")

    @property
    @pulumi.getter(name="sharedAccessPolicyName")
    def shared_access_policy_name(self) -> pulumi.Output[str]:
        """
        The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        """
        return pulumi.get(self, "shared_access_policy_name")

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

