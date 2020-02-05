// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ad

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/azuread_application.html.markdown.
type Application struct {
	pulumi.CustomResourceState

	// The ID of the Azure Active Directory Application.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrOutput `pulumi:"availableToOtherTenants"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringOutput `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayOutput `pulumi:"identifierUris"`
	// The display name for the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrOutput `pulumi:"oauth2AllowImplicitFlow"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayOutput `pulumi:"replyUrls"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("azure:ad/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure:ad/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The ID of the Azure Active Directory Application.
	ApplicationId *string `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants *bool `pulumi:"availableToOtherTenants"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage *string `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// The display name for the application.
	Name *string `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow *bool `pulumi:"oauth2AllowImplicitFlow"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls []string `pulumi:"replyUrls"`
}

type ApplicationState struct {
	// The ID of the Azure Active Directory Application.
	ApplicationId pulumi.StringPtrInput
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrInput
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringPtrInput
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// The display name for the application.
	Name pulumi.StringPtrInput
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrInput
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants *bool `pulumi:"availableToOtherTenants"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage *string `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// The display name for the application.
	Name *string `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow *bool `pulumi:"oauth2AllowImplicitFlow"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls []string `pulumi:"replyUrls"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrInput
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringPtrInput
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// The display name for the application.
	Name pulumi.StringPtrInput
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrInput
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

