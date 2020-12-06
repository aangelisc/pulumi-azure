// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This certificate can be used to secure custom domains on App Services (Windows and Linux) hosted on an App Service Plan of Basic and above (free and shared tiers are not supported).
 *
 * > NOTE: A certificate is valid for six months, and about a month before the certificate’s expiration date, App Services renews/rotates the certificate. This is managed by Azure and doesn't requre this resource to be changed or reprovisioned. It will change the `thumbprint` computed attribute the next time the resource is refreshed after rotation occurs, so keep that in mind if you have any dependencies on this attribute directly.
 *
 * ## Import
 *
 * App Service Managed Certificates can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appservice/managedCertificate:ManagedCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Web/certificates/customhost.contoso.com
 * ```
 */
export class ManagedCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ManagedCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedCertificateState, opts?: pulumi.CustomResourceOptions): ManagedCertificate {
        return new ManagedCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/managedCertificate:ManagedCertificate';

    /**
     * Returns true if the given object is an instance of ManagedCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedCertificate.__pulumiType;
    }

    /**
     * The Canonical Name of the Certificate.
     */
    public /*out*/ readonly canonicalName!: pulumi.Output<string>;
    /**
     * The ID of the App Service Custom Hostname Binding for the Certificate. Changing this forces a new App Service Managed Certificate to be created.
     */
    public readonly customHostnameBindingId!: pulumi.Output<string>;
    /**
     * The expiration date of the Certificate.
     */
    public /*out*/ readonly expirationDate!: pulumi.Output<string>;
    /**
     * The friendly name of the Certificate.
     */
    public /*out*/ readonly friendlyName!: pulumi.Output<string>;
    /**
     * The list of Host Names for the Certificate.
     */
    public /*out*/ readonly hostNames!: pulumi.Output<string[]>;
    /**
     * The Start date for the Certificate.
     */
    public /*out*/ readonly issueDate!: pulumi.Output<string>;
    /**
     * The issuer of the Certificate.
     */
    public /*out*/ readonly issuer!: pulumi.Output<string>;
    /**
     * The Subject Name for the Certificate.
     */
    public /*out*/ readonly subjectName!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the App Service Managed Certificate.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Certificate Thumbprint.
     */
    public /*out*/ readonly thumbprint!: pulumi.Output<string>;

    /**
     * Create a ManagedCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedCertificateArgs | ManagedCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ManagedCertificateState | undefined;
            inputs["canonicalName"] = state ? state.canonicalName : undefined;
            inputs["customHostnameBindingId"] = state ? state.customHostnameBindingId : undefined;
            inputs["expirationDate"] = state ? state.expirationDate : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["hostNames"] = state ? state.hostNames : undefined;
            inputs["issueDate"] = state ? state.issueDate : undefined;
            inputs["issuer"] = state ? state.issuer : undefined;
            inputs["subjectName"] = state ? state.subjectName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["thumbprint"] = state ? state.thumbprint : undefined;
        } else {
            const args = argsOrState as ManagedCertificateArgs | undefined;
            if (!args || args.customHostnameBindingId === undefined) {
                throw new Error("Missing required property 'customHostnameBindingId'");
            }
            inputs["customHostnameBindingId"] = args ? args.customHostnameBindingId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["canonicalName"] = undefined /*out*/;
            inputs["expirationDate"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["hostNames"] = undefined /*out*/;
            inputs["issueDate"] = undefined /*out*/;
            inputs["issuer"] = undefined /*out*/;
            inputs["subjectName"] = undefined /*out*/;
            inputs["thumbprint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagedCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedCertificate resources.
 */
export interface ManagedCertificateState {
    /**
     * The Canonical Name of the Certificate.
     */
    readonly canonicalName?: pulumi.Input<string>;
    /**
     * The ID of the App Service Custom Hostname Binding for the Certificate. Changing this forces a new App Service Managed Certificate to be created.
     */
    readonly customHostnameBindingId?: pulumi.Input<string>;
    /**
     * The expiration date of the Certificate.
     */
    readonly expirationDate?: pulumi.Input<string>;
    /**
     * The friendly name of the Certificate.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The list of Host Names for the Certificate.
     */
    readonly hostNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Start date for the Certificate.
     */
    readonly issueDate?: pulumi.Input<string>;
    /**
     * The issuer of the Certificate.
     */
    readonly issuer?: pulumi.Input<string>;
    /**
     * The Subject Name for the Certificate.
     */
    readonly subjectName?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the App Service Managed Certificate.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Certificate Thumbprint.
     */
    readonly thumbprint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedCertificate resource.
 */
export interface ManagedCertificateArgs {
    /**
     * The ID of the App Service Custom Hostname Binding for the Certificate. Changing this forces a new App Service Managed Certificate to be created.
     */
    readonly customHostnameBindingId: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the App Service Managed Certificate.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
