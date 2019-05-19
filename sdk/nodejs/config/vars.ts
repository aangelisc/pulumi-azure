// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("azure");

export let clientCertificatePassword: string | undefined = __config.get("clientCertificatePassword") || (utilities.getEnv("ARM_CLIENT_CERTIFICATE_PASSWORD") || "");
export let clientCertificatePath: string | undefined = __config.get("clientCertificatePath") || (utilities.getEnv("ARM_CLIENT_CERTIFICATE_PATH") || "");
export let clientId: string | undefined = __config.get("clientId") || (utilities.getEnv("ARM_CLIENT_ID") || "");
export let clientSecret: string | undefined = __config.get("clientSecret") || (utilities.getEnv("ARM_CLIENT_SECRET") || "");
export let environment: string | undefined = __config.get("environment") || (utilities.getEnv("ARM_ENVIRONMENT") || "public");
export let msiEndpoint: string | undefined = __config.get("msiEndpoint") || (utilities.getEnv("ARM_MSI_ENDPOINT") || "");
export let partnerId: string | undefined = __config.get("partnerId") || (utilities.getEnv("ARM_PARTNER_ID") || "");
export let skipCredentialsValidation: boolean | undefined = __config.getObject<boolean>("skipCredentialsValidation") || (utilities.getEnvBoolean("ARM_SKIP_CREDENTIALS_VALIDATION") || false);
export let skipProviderRegistration: boolean | undefined = __config.getObject<boolean>("skipProviderRegistration") || (utilities.getEnvBoolean("ARM_SKIP_PROVIDER_REGISTRATION") || false);
export let subscriptionId: string | undefined = __config.get("subscriptionId") || (utilities.getEnv("ARM_SUBSCRIPTION_ID") || "");
export let tenantId: string | undefined = __config.get("tenantId") || (utilities.getEnv("ARM_TENANT_ID") || "");
export let useMsi: boolean | undefined = __config.getObject<boolean>("useMsi") || (utilities.getEnvBoolean("ARM_USE_MSI") || false);
export let location: string | undefined = __config.get("location") || utilities.getEnv("ARM_LOCATION");
