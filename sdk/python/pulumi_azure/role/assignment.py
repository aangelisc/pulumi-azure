# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class assignment(pulumi.CustomResource):
    """
    Assigns a given Principal (User or Application) to a given Role.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, principal_id=None, role_definition_id=None, role_definition_name=None, scope=None):
        """Create a assignment resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not principal_id:
            raise TypeError('Missing required property principal_id')
        elif not isinstance(principal_id, basestring):
            raise TypeError('Expected property principal_id to be a basestring')
        __self__.principal_id = principal_id
        """
        The ID of the Principal (User or Application) to assign the Role Definition to. Changing this forces a new resource to be created.
        """
        __props__['principalId'] = principal_id

        if role_definition_id and not isinstance(role_definition_id, basestring):
            raise TypeError('Expected property role_definition_id to be a basestring')
        __self__.role_definition_id = role_definition_id
        """
        The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `role_definition_name`.
        """
        __props__['roleDefinitionId'] = role_definition_id

        if role_definition_name and not isinstance(role_definition_name, basestring):
            raise TypeError('Expected property role_definition_name to be a basestring')
        __self__.role_definition_name = role_definition_name
        """
        The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `role_definition_id`.
        """
        __props__['roleDefinitionName'] = role_definition_name

        if not scope:
            raise TypeError('Missing required property scope')
        elif not isinstance(scope, basestring):
            raise TypeError('Expected property scope to be a basestring')
        __self__.scope = scope
        """
        The scope at which the Role Assignment applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
        """
        __props__['scope'] = scope

        super(assignment, __self__).__init__(
            'azure:role/assignment:assignment',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'name' in outs:
            self.name = outs['name']
        if 'principalId' in outs:
            self.principal_id = outs['principalId']
        if 'roleDefinitionId' in outs:
            self.role_definition_id = outs['roleDefinitionId']
        if 'roleDefinitionName' in outs:
            self.role_definition_name = outs['roleDefinitionName']
        if 'scope' in outs:
            self.scope = outs['scope']
