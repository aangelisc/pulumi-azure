// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keyvault

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Key Vault Certificate.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/key_vault_certificate.html.markdown.
type Certificate struct {
	pulumi.CustomResourceState

	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertificateCertificatePtrOutput `pulumi:"certificate"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData pulumi.StringOutput `pulumi:"certificateData"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicyOutput `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the associated Key Vault Secret.
	SecretId pulumi.StringOutput `pulumi:"secretId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The current version of the Key Vault Certificate.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.CertificatePolicy == nil {
		return nil, errors.New("missing required argument 'CertificatePolicy'")
	}
	if args == nil || args.KeyVaultId == nil {
		return nil, errors.New("missing required argument 'KeyVaultId'")
	}
	if args == nil {
		args = &CertificateArgs{}
	}
	var resource Certificate
	err := ctx.RegisterResource("azure:keyvault/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure:keyvault/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate *CertificateCertificate `pulumi:"certificate"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData *string `pulumi:"certificateData"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy *CertificateCertificatePolicy `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the associated Key Vault Secret.
	SecretId *string `pulumi:"secretId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint *string `pulumi:"thumbprint"`
	// The current version of the Key Vault Certificate.
	Version *string `pulumi:"version"`
}

type CertificateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertificateCertificatePtrInput
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData pulumi.StringPtrInput
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicyPtrInput
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringPtrInput
	// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the associated Key Vault Secret.
	SecretId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint pulumi.StringPtrInput
	// The current version of the Key Vault Certificate.
	Version pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate *CertificateCertificate `pulumi:"certificate"`
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicy `pulumi:"certificatePolicy"`
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate CertificateCertificatePtrInput
	// A `certificatePolicy` block as defined below.
	CertificatePolicy CertificateCertificatePolicyInput
	// The ID of the Key Vault where the Certificate should be created.
	KeyVaultId pulumi.StringInput
	// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

