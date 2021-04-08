# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ServiceCorArgs',
    'ServiceFeatureArgs',
    'ServiceSkuArgs',
    'ServiceUpstreamEndpointArgs',
]

@pulumi.input_type
class ServiceCorArgs:
    def __init__(__self__, *,
                 allowed_origins: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_origins: A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
        """
        pulumi.set(__self__, "allowed_origins", allowed_origins)

    @property
    @pulumi.getter(name="allowedOrigins")
    def allowed_origins(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
        """
        return pulumi.get(self, "allowed_origins")

    @allowed_origins.setter
    def allowed_origins(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_origins", value)


@pulumi.input_type
class ServiceFeatureArgs:
    def __init__(__self__, *,
                 flag: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] flag: The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
        :param pulumi.Input[str] value: A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
        """
        pulumi.set(__self__, "flag", flag)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def flag(self) -> pulumi.Input[str]:
        """
        The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
        """
        return pulumi.get(self, "flag")

    @flag.setter
    def flag(self, value: pulumi.Input[str]):
        pulumi.set(self, "flag", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceSkuArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[int],
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[int] capacity: Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
        :param pulumi.Input[str] name: Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[int]:
        """
        Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ServiceUpstreamEndpointArgs:
    def __init__(__self__, *,
                 category_patterns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 event_patterns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 hub_patterns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 url_template: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] category_patterns: The categories to match on, or `*` for all.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_patterns: The events to match on, or `*` for all.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hub_patterns: The hubs to match on, or `*` for all.
        :param pulumi.Input[str] url_template: The upstream URL Template. This can be a url or a template such as `http://host.com/{hub}/api/{category}/{event}`.
        """
        pulumi.set(__self__, "category_patterns", category_patterns)
        pulumi.set(__self__, "event_patterns", event_patterns)
        pulumi.set(__self__, "hub_patterns", hub_patterns)
        pulumi.set(__self__, "url_template", url_template)

    @property
    @pulumi.getter(name="categoryPatterns")
    def category_patterns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The categories to match on, or `*` for all.
        """
        return pulumi.get(self, "category_patterns")

    @category_patterns.setter
    def category_patterns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "category_patterns", value)

    @property
    @pulumi.getter(name="eventPatterns")
    def event_patterns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The events to match on, or `*` for all.
        """
        return pulumi.get(self, "event_patterns")

    @event_patterns.setter
    def event_patterns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "event_patterns", value)

    @property
    @pulumi.getter(name="hubPatterns")
    def hub_patterns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The hubs to match on, or `*` for all.
        """
        return pulumi.get(self, "hub_patterns")

    @hub_patterns.setter
    def hub_patterns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "hub_patterns", value)

    @property
    @pulumi.getter(name="urlTemplate")
    def url_template(self) -> pulumi.Input[str]:
        """
        The upstream URL Template. This can be a url or a template such as `http://host.com/{hub}/api/{category}/{event}`.
        """
        return pulumi.get(self, "url_template")

    @url_template.setter
    def url_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_template", value)


