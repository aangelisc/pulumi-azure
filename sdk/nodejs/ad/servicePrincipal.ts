// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ServicePrincipal extends pulumi.CustomResource {
    /**
     * Get an existing ServicePrincipal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePrincipalState, opts?: pulumi.CustomResourceOptions): ServicePrincipal {
        return new ServicePrincipal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:ad/servicePrincipal:ServicePrincipal';

    /**
     * Returns true if the given object is an instance of ServicePrincipal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePrincipal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePrincipal.__pulumiType;
    }

    /**
     * The ID of the Azure AD Application for which to create a Service Principal.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The Display Name of the Azure Active Directory Application associated with this Service Principal.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;

    /**
     * Create a ServicePrincipal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePrincipalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePrincipalArgs | ServicePrincipalState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServicePrincipalState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
        } else {
            const args = argsOrState as ServicePrincipalArgs | undefined;
            if (!args || args.applicationId === undefined) {
                throw new Error("Missing required property 'applicationId'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["displayName"] = undefined /*out*/;
        }
        super(ServicePrincipal.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipal resources.
 */
export interface ServicePrincipalState {
    /**
     * The ID of the Azure AD Application for which to create a Service Principal.
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * The Display Name of the Azure Active Directory Application associated with this Service Principal.
     */
    readonly displayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServicePrincipal resource.
 */
export interface ServicePrincipalArgs {
    /**
     * The ID of the Azure AD Application for which to create a Service Principal.
     */
    readonly applicationId: pulumi.Input<string>;
}
