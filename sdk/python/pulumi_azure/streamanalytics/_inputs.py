# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'FunctionJavaScriptUDFInputArgs',
    'FunctionJavaScriptUDFOutputArgs',
    'OutputBlobSerializationArgs',
    'OutputEventHubSerializationArgs',
    'OutputServiceBusQueueSerializationArgs',
    'OutputServicebusTopicSerializationArgs',
    'ReferenceInputBlobSerializationArgs',
    'StreamInputBlobSerializationArgs',
    'StreamInputEventHubSerializationArgs',
    'StreamInputIotHubSerializationArgs',
]

@pulumi.input_type
class FunctionJavaScriptUDFInputArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] type: The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class FunctionJavaScriptUDFOutputArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] type: The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class OutputBlobSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param pulumi.Input[str] format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)


@pulumi.input_type
class OutputEventHubSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param pulumi.Input[str] format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)


@pulumi.input_type
class OutputServiceBusQueueSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param pulumi.Input[str] format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)


@pulumi.input_type
class OutputServicebusTopicSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        :param pulumi.Input[str] format: Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)


@pulumi.input_type
class ReferenceInputBlobSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for the reference data. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `	` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for the reference data. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `	` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)


@pulumi.input_type
class StreamInputBlobSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)


@pulumi.input_type
class StreamInputEventHubSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)


@pulumi.input_type
class StreamInputIotHubSerializationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 field_delimiter: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        :param pulumi.Input[str] encoding: The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        :param pulumi.Input[str] field_delimiter: The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        pulumi.set(__self__, "type", type)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if field_delimiter is not None:
            pulumi.set(__self__, "field_delimiter", field_delimiter)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The serialization format used for incoming data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="fieldDelimiter")
    def field_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
        """
        return pulumi.get(self, "field_delimiter")

    @field_delimiter.setter
    def field_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_delimiter", value)


