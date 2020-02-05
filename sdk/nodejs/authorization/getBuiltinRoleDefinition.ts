// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about a built-in Role Definition. To access information about a custom Role Definition, please see the `azure.authorization.RoleDefinition` data source instead.
 * 
 * > **Note:** The this datasource has been deprecated in favour of `azure.authorization.RoleDefinition` that now can look up role definitions by name. As such this data source will be removed in version 2.0 of the AzureRM Provider.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const contributor = azure.authorization.getBuiltinRoleDefinition({
 *     name: "Contributor",
 * });
 * 
 * export const contributorRoleDefinitionId = contributor.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/builtin_role_definition.html.markdown.
 */
export function getBuiltinRoleDefinition(args: GetBuiltinRoleDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetBuiltinRoleDefinitionResult> & GetBuiltinRoleDefinitionResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetBuiltinRoleDefinitionResult> = pulumi.runtime.invoke("azure:authorization/getBuiltinRoleDefinition:getBuiltinRoleDefinition", {
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getBuiltinRoleDefinition.
 */
export interface GetBuiltinRoleDefinitionArgs {
    /**
     * Specifies the name of the built-in Role Definition. Possible values are: `Contributor`, `Owner`, `Reader` and `VirtualMachineContributor`.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getBuiltinRoleDefinition.
 */
export interface GetBuiltinRoleDefinitionResult {
    /**
     * One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
     */
    readonly assignableScopes: string[];
    /**
     * the Description of the built-in Role.
     */
    readonly description: string;
    readonly name: string;
    /**
     * a `permissions` block as documented below.
     */
    readonly permissions: outputs.authorization.GetBuiltinRoleDefinitionPermission[];
    /**
     * the Type of the Role.
     */
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
