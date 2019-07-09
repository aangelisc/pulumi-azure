// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Recovery Protected VM.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West US",
 *     name: "tfex-recovery_vault",
 * });
 * const exampleVault = new azure.recoveryservices.Vault("example", {
 *     location: exampleResourceGroup.location,
 *     name: "tfex-recovery-vault",
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 * });
 * const exampleProtectionPolicyVM = new azure.recoveryservices.ProtectionPolicyVM("example", {
 *     backup: {
 *         frequency: "Daily",
 *         time: "23:00",
 *     },
 *     name: "tfex-recovery-vault-policy",
 *     recoveryVaultName: exampleVault.name,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const vm1 = new azure.recoveryservices.ProtectedVM("vm1", {
 *     backupPolicyId: exampleProtectionPolicyVM.id,
 *     recoveryVaultName: exampleVault.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sourceVmId: azurerm_virtual_machine_example.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/recovery_services_protected_vm.html.markdown.
 */
export class ProtectedVM extends pulumi.CustomResource {
    /**
     * Get an existing ProtectedVM resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProtectedVMState, opts?: pulumi.CustomResourceOptions): ProtectedVM {
        return new ProtectedVM(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:recoveryservices/protectedVM:ProtectedVM';

    /**
     * Returns true if the given object is an instance of ProtectedVM.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProtectedVM {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProtectedVM.__pulumiType;
    }

    /**
     * Specifies the id of the backup policy to use. Changing this forces a new resource to be created.
     */
    public readonly backupPolicyId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
     */
    public readonly recoveryVaultName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
     */
    public readonly sourceVmId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a ProtectedVM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProtectedVMArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProtectedVMArgs | ProtectedVMState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProtectedVMState | undefined;
            inputs["backupPolicyId"] = state ? state.backupPolicyId : undefined;
            inputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sourceVmId"] = state ? state.sourceVmId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ProtectedVMArgs | undefined;
            if (!args || args.backupPolicyId === undefined) {
                throw new Error("Missing required property 'backupPolicyId'");
            }
            if (!args || args.recoveryVaultName === undefined) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sourceVmId === undefined) {
                throw new Error("Missing required property 'sourceVmId'");
            }
            inputs["backupPolicyId"] = args ? args.backupPolicyId : undefined;
            inputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceVmId"] = args ? args.sourceVmId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        super(ProtectedVM.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProtectedVM resources.
 */
export interface ProtectedVMState {
    /**
     * Specifies the id of the backup policy to use. Changing this forces a new resource to be created.
     */
    readonly backupPolicyId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
     */
    readonly recoveryVaultName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
     */
    readonly sourceVmId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ProtectedVM resource.
 */
export interface ProtectedVMArgs {
    /**
     * Specifies the id of the backup policy to use. Changing this forces a new resource to be created.
     */
    readonly backupPolicyId: pulumi.Input<string>;
    /**
     * Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
     */
    readonly recoveryVaultName: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
     */
    readonly sourceVmId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
