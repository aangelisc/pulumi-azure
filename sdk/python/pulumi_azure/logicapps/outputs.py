# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ActionHttpRunAfter',
    'TriggerRecurrenceSchedule',
]

@pulumi.output_type
class ActionHttpRunAfter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionName":
            suggest = "action_name"
        elif key == "actionResult":
            suggest = "action_result"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ActionHttpRunAfter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ActionHttpRunAfter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ActionHttpRunAfter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action_name: str,
                 action_result: str):
        """
        :param str action_name: Specifies the name of the precedent HTTP Action.
        :param str action_result: Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
        """
        pulumi.set(__self__, "action_name", action_name)
        pulumi.set(__self__, "action_result", action_result)

    @property
    @pulumi.getter(name="actionName")
    def action_name(self) -> str:
        """
        Specifies the name of the precedent HTTP Action.
        """
        return pulumi.get(self, "action_name")

    @property
    @pulumi.getter(name="actionResult")
    def action_result(self) -> str:
        """
        Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
        """
        return pulumi.get(self, "action_result")


@pulumi.output_type
class TriggerRecurrenceSchedule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "atTheseHours":
            suggest = "at_these_hours"
        elif key == "atTheseMinutes":
            suggest = "at_these_minutes"
        elif key == "onTheseDays":
            suggest = "on_these_days"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerRecurrenceSchedule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerRecurrenceSchedule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerRecurrenceSchedule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 at_these_hours: Optional[Sequence[int]] = None,
                 at_these_minutes: Optional[Sequence[int]] = None,
                 on_these_days: Optional[Sequence[str]] = None):
        """
        :param Sequence[int] at_these_hours: Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
        :param Sequence[int] at_these_minutes: Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
        :param Sequence[str] on_these_days: Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
        """
        if at_these_hours is not None:
            pulumi.set(__self__, "at_these_hours", at_these_hours)
        if at_these_minutes is not None:
            pulumi.set(__self__, "at_these_minutes", at_these_minutes)
        if on_these_days is not None:
            pulumi.set(__self__, "on_these_days", on_these_days)

    @property
    @pulumi.getter(name="atTheseHours")
    def at_these_hours(self) -> Optional[Sequence[int]]:
        """
        Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
        """
        return pulumi.get(self, "at_these_hours")

    @property
    @pulumi.getter(name="atTheseMinutes")
    def at_these_minutes(self) -> Optional[Sequence[int]]:
        """
        Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
        """
        return pulumi.get(self, "at_these_minutes")

    @property
    @pulumi.getter(name="onTheseDays")
    def on_these_days(self) -> Optional[Sequence[str]]:
        """
        Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
        """
        return pulumi.get(self, "on_these_days")


