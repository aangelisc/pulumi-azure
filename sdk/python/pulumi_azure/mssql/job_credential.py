# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['JobCredentialArgs', 'JobCredential']

@pulumi.input_type
class JobCredentialArgs:
    def __init__(__self__, *,
                 job_agent_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a JobCredential resource.
        :param pulumi.Input[str] job_agent_id: The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] password: The password part of the credential.
        :param pulumi.Input[str] username: The username part of the credential.
        :param pulumi.Input[str] name: The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        """
        pulumi.set(__self__, "job_agent_id", job_agent_id)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="jobAgentId")
    def job_agent_id(self) -> pulumi.Input[str]:
        """
        The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        """
        return pulumi.get(self, "job_agent_id")

    @job_agent_id.setter
    def job_agent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_agent_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The password part of the credential.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The username part of the credential.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _JobCredentialState:
    def __init__(__self__, *,
                 job_agent_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering JobCredential resources.
        :param pulumi.Input[str] job_agent_id: The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] name: The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] password: The password part of the credential.
        :param pulumi.Input[str] username: The username part of the credential.
        """
        if job_agent_id is not None:
            pulumi.set(__self__, "job_agent_id", job_agent_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="jobAgentId")
    def job_agent_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        """
        return pulumi.get(self, "job_agent_id")

    @job_agent_id.setter
    def job_agent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_agent_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password part of the credential.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username part of the credential.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class JobCredential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_agent_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Elastic Job Credential.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_server = azure.mssql.Server("exampleServer",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            version="12.0",
            administrator_login="4dm1n157r470r",
            administrator_login_password="4-v3ry-53cr37-p455w0rd")
        example_database = azure.mssql.Database("exampleDatabase",
            server_id=example_server.id,
            collation="SQL_Latin1_General_CP1_CI_AS",
            sku_name="S1")
        example_job_agent = azure.mssql.JobAgent("exampleJobAgent",
            location=example_resource_group.location,
            database_id=example_database.id)
        example_job_credential = azure.mssql.JobCredential("exampleJobCredential",
            job_agent_id=example_job_agent.id,
            username="my-username",
            password="MyP4ssw0rd!!!")
        ```

        ## Import

        Elastic Job Credentials can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mssql/jobCredential:JobCredential example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/credentials/credential1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_agent_id: The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] name: The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] password: The password part of the credential.
        :param pulumi.Input[str] username: The username part of the credential.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Elastic Job Credential.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_server = azure.mssql.Server("exampleServer",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            version="12.0",
            administrator_login="4dm1n157r470r",
            administrator_login_password="4-v3ry-53cr37-p455w0rd")
        example_database = azure.mssql.Database("exampleDatabase",
            server_id=example_server.id,
            collation="SQL_Latin1_General_CP1_CI_AS",
            sku_name="S1")
        example_job_agent = azure.mssql.JobAgent("exampleJobAgent",
            location=example_resource_group.location,
            database_id=example_database.id)
        example_job_credential = azure.mssql.JobCredential("exampleJobCredential",
            job_agent_id=example_job_agent.id,
            username="my-username",
            password="MyP4ssw0rd!!!")
        ```

        ## Import

        Elastic Job Credentials can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mssql/jobCredential:JobCredential example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/credentials/credential1
        ```

        :param str resource_name: The name of the resource.
        :param JobCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_agent_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobCredentialArgs.__new__(JobCredentialArgs)

            if job_agent_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_agent_id'")
            __props__.__dict__["job_agent_id"] = job_agent_id
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(JobCredential, __self__).__init__(
            'azure:mssql/jobCredential:JobCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            job_agent_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'JobCredential':
        """
        Get an existing JobCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_agent_id: The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] name: The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        :param pulumi.Input[str] password: The password part of the credential.
        :param pulumi.Input[str] username: The username part of the credential.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobCredentialState.__new__(_JobCredentialState)

        __props__.__dict__["job_agent_id"] = job_agent_id
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["username"] = username
        return JobCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="jobAgentId")
    def job_agent_id(self) -> pulumi.Output[str]:
        """
        The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
        """
        return pulumi.get(self, "job_agent_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Elastic Job Credential. Changing this forces a new Elastic Job Credential to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password part of the credential.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The username part of the credential.
        """
        return pulumi.get(self, "username")

