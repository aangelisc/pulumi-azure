// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Create an Azure Storage Account.
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState): Account {
        return new Account(name, <any>state, { id });
    }

    /**
     * Defines the access tier for `BlobStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
     */
    public readonly accessTier: pulumi.Output<string>;
    /**
     * The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
     */
    public readonly accountEncryptionSource: pulumi.Output<string | undefined>;
    /**
     * Defines the Kind of account. Valid options are `Storage`,
     * `StorageV2` and `BlobStorage`. Changing this forces a new resource to be created.
     * Defaults to `Storage`.
     */
    public readonly accountKind: pulumi.Output<string | undefined>;
    /**
     * Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
     */
    public readonly accountReplicationType: pulumi.Output<string>;
    /**
     * Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. Changing this forces a new resource to be created
     */
    public readonly accountTier: pulumi.Output<string>;
    public readonly accountType: pulumi.Output<string>;
    /**
     * A `custom_domain` block as documented below.
     */
    public readonly customDomain: pulumi.Output<{ name: string, useSubdomain?: boolean } | undefined>;
    /**
     * Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
     */
    public readonly enableBlobEncryption: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
     */
    public readonly enableFileEncryption: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
     * for more information.
     */
    public readonly enableHttpsTrafficOnly: pulumi.Output<boolean | undefined>;
    /**
     * A Managed Service Identity block as defined below.
     */
    public readonly identity: pulumi.Output<{ principalId: string, tenantId: string, type: string }>;
    /**
     * Specifies the supported Azure location where the
     * resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * The Custom Domain Name to use for the Storage Account, which will be validated by Azure.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A `network_rules` block as documented below.
     */
    public readonly networkRules: pulumi.Output<{ bypasses: string[], ipRules?: string[], virtualNetworkSubnetIds?: string[] } | undefined>;
    /**
     * The primary access key for the storage account
     */
    public /*out*/ readonly primaryAccessKey: pulumi.Output<string>;
    /**
     * The connection string associated with the primary blob location
     */
    public /*out*/ readonly primaryBlobConnectionString: pulumi.Output<string>;
    /**
     * The endpoint URL for blob storage in the primary location.
     */
    public /*out*/ readonly primaryBlobEndpoint: pulumi.Output<string>;
    /**
     * The connection string associated with the primary location
     */
    public /*out*/ readonly primaryConnectionString: pulumi.Output<string>;
    /**
     * The endpoint URL for file storage in the primary location.
     */
    public /*out*/ readonly primaryFileEndpoint: pulumi.Output<string>;
    /**
     * The primary location of the storage account.
     */
    public /*out*/ readonly primaryLocation: pulumi.Output<string>;
    /**
     * The endpoint URL for queue storage in the primary location.
     */
    public /*out*/ readonly primaryQueueEndpoint: pulumi.Output<string>;
    /**
     * The endpoint URL for table storage in the primary location.
     */
    public /*out*/ readonly primaryTableEndpoint: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the storage account. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The secondary access key for the storage account
     */
    public /*out*/ readonly secondaryAccessKey: pulumi.Output<string>;
    /**
     * The connection string associated with the secondary blob location
     */
    public /*out*/ readonly secondaryBlobConnectionString: pulumi.Output<string>;
    /**
     * The endpoint URL for blob storage in the secondary location.
     */
    public /*out*/ readonly secondaryBlobEndpoint: pulumi.Output<string>;
    /**
     * The connection string associated with the secondary location
     */
    public /*out*/ readonly secondaryConnectionString: pulumi.Output<string>;
    /**
     * The secondary location of the storage account.
     */
    public /*out*/ readonly secondaryLocation: pulumi.Output<string>;
    /**
     * The endpoint URL for queue storage in the secondary location.
     */
    public /*out*/ readonly secondaryQueueEndpoint: pulumi.Output<string>;
    /**
     * The endpoint URL for table storage in the secondary location.
     */
    public /*out*/ readonly secondaryTableEndpoint: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AccountState = argsOrState as AccountState | undefined;
            inputs["accessTier"] = state ? state.accessTier : undefined;
            inputs["accountEncryptionSource"] = state ? state.accountEncryptionSource : undefined;
            inputs["accountKind"] = state ? state.accountKind : undefined;
            inputs["accountReplicationType"] = state ? state.accountReplicationType : undefined;
            inputs["accountTier"] = state ? state.accountTier : undefined;
            inputs["accountType"] = state ? state.accountType : undefined;
            inputs["customDomain"] = state ? state.customDomain : undefined;
            inputs["enableBlobEncryption"] = state ? state.enableBlobEncryption : undefined;
            inputs["enableFileEncryption"] = state ? state.enableFileEncryption : undefined;
            inputs["enableHttpsTrafficOnly"] = state ? state.enableHttpsTrafficOnly : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkRules"] = state ? state.networkRules : undefined;
            inputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            inputs["primaryBlobConnectionString"] = state ? state.primaryBlobConnectionString : undefined;
            inputs["primaryBlobEndpoint"] = state ? state.primaryBlobEndpoint : undefined;
            inputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            inputs["primaryFileEndpoint"] = state ? state.primaryFileEndpoint : undefined;
            inputs["primaryLocation"] = state ? state.primaryLocation : undefined;
            inputs["primaryQueueEndpoint"] = state ? state.primaryQueueEndpoint : undefined;
            inputs["primaryTableEndpoint"] = state ? state.primaryTableEndpoint : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            inputs["secondaryBlobConnectionString"] = state ? state.secondaryBlobConnectionString : undefined;
            inputs["secondaryBlobEndpoint"] = state ? state.secondaryBlobEndpoint : undefined;
            inputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            inputs["secondaryLocation"] = state ? state.secondaryLocation : undefined;
            inputs["secondaryQueueEndpoint"] = state ? state.secondaryQueueEndpoint : undefined;
            inputs["secondaryTableEndpoint"] = state ? state.secondaryTableEndpoint : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.accountReplicationType === undefined) {
                throw new Error("Missing required property 'accountReplicationType'");
            }
            if (!args || args.accountTier === undefined) {
                throw new Error("Missing required property 'accountTier'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accessTier"] = args ? args.accessTier : undefined;
            inputs["accountEncryptionSource"] = args ? args.accountEncryptionSource : undefined;
            inputs["accountKind"] = args ? args.accountKind : undefined;
            inputs["accountReplicationType"] = args ? args.accountReplicationType : undefined;
            inputs["accountTier"] = args ? args.accountTier : undefined;
            inputs["accountType"] = args ? args.accountType : undefined;
            inputs["customDomain"] = args ? args.customDomain : undefined;
            inputs["enableBlobEncryption"] = args ? args.enableBlobEncryption : undefined;
            inputs["enableFileEncryption"] = args ? args.enableFileEncryption : undefined;
            inputs["enableHttpsTrafficOnly"] = args ? args.enableHttpsTrafficOnly : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkRules"] = args ? args.networkRules : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["primaryAccessKey"] = undefined /*out*/;
            inputs["primaryBlobConnectionString"] = undefined /*out*/;
            inputs["primaryBlobEndpoint"] = undefined /*out*/;
            inputs["primaryConnectionString"] = undefined /*out*/;
            inputs["primaryFileEndpoint"] = undefined /*out*/;
            inputs["primaryLocation"] = undefined /*out*/;
            inputs["primaryQueueEndpoint"] = undefined /*out*/;
            inputs["primaryTableEndpoint"] = undefined /*out*/;
            inputs["secondaryAccessKey"] = undefined /*out*/;
            inputs["secondaryBlobConnectionString"] = undefined /*out*/;
            inputs["secondaryBlobEndpoint"] = undefined /*out*/;
            inputs["secondaryConnectionString"] = undefined /*out*/;
            inputs["secondaryLocation"] = undefined /*out*/;
            inputs["secondaryQueueEndpoint"] = undefined /*out*/;
            inputs["secondaryTableEndpoint"] = undefined /*out*/;
        }
        super("azure:storage/account:Account", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * Defines the access tier for `BlobStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
     */
    readonly accessTier?: pulumi.Input<string>;
    /**
     * The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
     */
    readonly accountEncryptionSource?: pulumi.Input<string>;
    /**
     * Defines the Kind of account. Valid options are `Storage`,
     * `StorageV2` and `BlobStorage`. Changing this forces a new resource to be created.
     * Defaults to `Storage`.
     */
    readonly accountKind?: pulumi.Input<string>;
    /**
     * Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
     */
    readonly accountReplicationType?: pulumi.Input<string>;
    /**
     * Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. Changing this forces a new resource to be created
     */
    readonly accountTier?: pulumi.Input<string>;
    readonly accountType?: pulumi.Input<string>;
    /**
     * A `custom_domain` block as documented below.
     */
    readonly customDomain?: pulumi.Input<{ name: pulumi.Input<string>, useSubdomain?: pulumi.Input<boolean> }>;
    /**
     * Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
     */
    readonly enableBlobEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
     */
    readonly enableFileEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
     * for more information.
     */
    readonly enableHttpsTrafficOnly?: pulumi.Input<boolean>;
    /**
     * A Managed Service Identity block as defined below.
     */
    readonly identity?: pulumi.Input<{ principalId?: pulumi.Input<string>, tenantId?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * Specifies the supported Azure location where the
     * resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Custom Domain Name to use for the Storage Account, which will be validated by Azure.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `network_rules` block as documented below.
     */
    readonly networkRules?: pulumi.Input<{ bypasses?: pulumi.Input<pulumi.Input<string>[]>, ipRules?: pulumi.Input<pulumi.Input<string>[]>, virtualNetworkSubnetIds?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * The primary access key for the storage account
     */
    readonly primaryAccessKey?: pulumi.Input<string>;
    /**
     * The connection string associated with the primary blob location
     */
    readonly primaryBlobConnectionString?: pulumi.Input<string>;
    /**
     * The endpoint URL for blob storage in the primary location.
     */
    readonly primaryBlobEndpoint?: pulumi.Input<string>;
    /**
     * The connection string associated with the primary location
     */
    readonly primaryConnectionString?: pulumi.Input<string>;
    /**
     * The endpoint URL for file storage in the primary location.
     */
    readonly primaryFileEndpoint?: pulumi.Input<string>;
    /**
     * The primary location of the storage account.
     */
    readonly primaryLocation?: pulumi.Input<string>;
    /**
     * The endpoint URL for queue storage in the primary location.
     */
    readonly primaryQueueEndpoint?: pulumi.Input<string>;
    /**
     * The endpoint URL for table storage in the primary location.
     */
    readonly primaryTableEndpoint?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the storage account. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The secondary access key for the storage account
     */
    readonly secondaryAccessKey?: pulumi.Input<string>;
    /**
     * The connection string associated with the secondary blob location
     */
    readonly secondaryBlobConnectionString?: pulumi.Input<string>;
    /**
     * The endpoint URL for blob storage in the secondary location.
     */
    readonly secondaryBlobEndpoint?: pulumi.Input<string>;
    /**
     * The connection string associated with the secondary location
     */
    readonly secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The secondary location of the storage account.
     */
    readonly secondaryLocation?: pulumi.Input<string>;
    /**
     * The endpoint URL for queue storage in the secondary location.
     */
    readonly secondaryQueueEndpoint?: pulumi.Input<string>;
    /**
     * The endpoint URL for table storage in the secondary location.
     */
    readonly secondaryTableEndpoint?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * Defines the access tier for `BlobStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
     */
    readonly accessTier?: pulumi.Input<string>;
    /**
     * The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
     */
    readonly accountEncryptionSource?: pulumi.Input<string>;
    /**
     * Defines the Kind of account. Valid options are `Storage`,
     * `StorageV2` and `BlobStorage`. Changing this forces a new resource to be created.
     * Defaults to `Storage`.
     */
    readonly accountKind?: pulumi.Input<string>;
    /**
     * Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
     */
    readonly accountReplicationType: pulumi.Input<string>;
    /**
     * Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. Changing this forces a new resource to be created
     */
    readonly accountTier: pulumi.Input<string>;
    readonly accountType?: pulumi.Input<string>;
    /**
     * A `custom_domain` block as documented below.
     */
    readonly customDomain?: pulumi.Input<{ name: pulumi.Input<string>, useSubdomain?: pulumi.Input<boolean> }>;
    /**
     * Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
     */
    readonly enableBlobEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
     */
    readonly enableFileEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
     * for more information.
     */
    readonly enableHttpsTrafficOnly?: pulumi.Input<boolean>;
    /**
     * A Managed Service Identity block as defined below.
     */
    readonly identity?: pulumi.Input<{ principalId?: pulumi.Input<string>, tenantId?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * Specifies the supported Azure location where the
     * resource exists. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The Custom Domain Name to use for the Storage Account, which will be validated by Azure.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `network_rules` block as documented below.
     */
    readonly networkRules?: pulumi.Input<{ bypasses?: pulumi.Input<pulumi.Input<string>[]>, ipRules?: pulumi.Input<pulumi.Input<string>[]>, virtualNetworkSubnetIds?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * The name of the resource group in which to
     * create the storage account. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
