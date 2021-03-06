// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Certificate within an API Management Service.
//
// ## Example Usage
//
// ## Import
//
// API Management Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/certificate:Certificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/certificates/certificate1
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// The Client ID of the User Assigned Managed Identity to use for retrieving certificate.
	KeyVaultIdentityClientId pulumi.StringPtrOutput `pulumi:"keyVaultIdentityClientId"`
	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type `application/x-pkcs12`.
	KeyVaultSecretId pulumi.StringPtrOutput `pulumi:"keyVaultSecretId"`
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Subject of this Certificate.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// The Thumbprint of this Certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Certificate
	err := ctx.RegisterResource("azure:apimanagement/certificate:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:apimanagement/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data *string `pulumi:"data"`
	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration *string `pulumi:"expiration"`
	// The Client ID of the User Assigned Managed Identity to use for retrieving certificate.
	KeyVaultIdentityClientId *string `pulumi:"keyVaultIdentityClientId"`
	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type `application/x-pkcs12`.
	KeyVaultSecretId *string `pulumi:"keyVaultSecretId"`
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Subject of this Certificate.
	Subject *string `pulumi:"subject"`
	// The Thumbprint of this Certificate.
	Thumbprint *string `pulumi:"thumbprint"`
}

type CertificateState struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data pulumi.StringPtrInput
	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration pulumi.StringPtrInput
	// The Client ID of the User Assigned Managed Identity to use for retrieving certificate.
	KeyVaultIdentityClientId pulumi.StringPtrInput
	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type `application/x-pkcs12`.
	KeyVaultSecretId pulumi.StringPtrInput
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Subject of this Certificate.
	Subject pulumi.StringPtrInput
	// The Thumbprint of this Certificate.
	Thumbprint pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data *string `pulumi:"data"`
	// The Client ID of the User Assigned Managed Identity to use for retrieving certificate.
	KeyVaultIdentityClientId *string `pulumi:"keyVaultIdentityClientId"`
	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type `application/x-pkcs12`.
	KeyVaultSecretId *string `pulumi:"keyVaultSecretId"`
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data pulumi.StringPtrInput
	// The Client ID of the User Assigned Managed Identity to use for retrieving certificate.
	KeyVaultIdentityClientId pulumi.StringPtrInput
	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type `application/x-pkcs12`.
	KeyVaultSecretId pulumi.StringPtrInput
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

func (i *Certificate) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

type CertificatePtrInput interface {
	pulumi.Input

	ToCertificatePtrOutput() CertificatePtrOutput
	ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput
}

type certificatePtrType CertificateArgs

func (*certificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (i *certificatePtrType) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *certificatePtrType) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//          CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Certificate)(nil))
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//          CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Certificate)(nil))
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o.ToCertificatePtrOutputWithContext(context.Background())
}

func (o CertificateOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o.ApplyT(func(v Certificate) *Certificate {
		return &v
	}).(CertificatePtrOutput)
}

type CertificatePtrOutput struct {
	*pulumi.OutputState
}

func (CertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (o CertificatePtrOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil))
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].([]Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Certificate)(nil))
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].(map[string]Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificatePtrOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
