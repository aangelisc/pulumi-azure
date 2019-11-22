// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a template deployment of resources
// 
// > **Note on ARM Template Deployments:** Due to the way the underlying Azure API is designed, this provider can only manage the deployment of the ARM Template - and not any resources which are created by it.
// This means that when deleting the `core.TemplateDeployment` resource, this provider will only remove the reference to the deployment, whilst leaving any resources created by that ARM Template Deployment.
// One workaround for this is to use a unique Resource Group for each ARM Template Deployment, which means deleting the Resource Group would contain any resources created within it - however this isn't ideal. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).
// 
// ## Note
// 
// This provider does not know about the individual resources created by Azure using a deployment template and therefore cannot delete these resources during a destroy. Destroying a template deployment removes the associated deployment operations, but will not delete the Azure resources created by the deployment. In order to delete these resources, the containing resource group must also be destroyed. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/template_deployment.html.markdown.
type TemplateDeployment struct {
	s *pulumi.ResourceState
}

// NewTemplateDeployment registers a new resource with the given unique name, arguments, and options.
func NewTemplateDeployment(ctx *pulumi.Context,
	name string, args *TemplateDeploymentArgs, opts ...pulumi.ResourceOpt) (*TemplateDeployment, error) {
	if args == nil || args.DeploymentMode == nil {
		return nil, errors.New("missing required argument 'DeploymentMode'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["deploymentMode"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["parametersBody"] = nil
		inputs["resourceGroupName"] = nil
		inputs["templateBody"] = nil
	} else {
		inputs["deploymentMode"] = args.DeploymentMode
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["parametersBody"] = args.ParametersBody
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["templateBody"] = args.TemplateBody
	}
	inputs["outputs"] = nil
	s, err := ctx.RegisterResource("azure:core/templateDeployment:TemplateDeployment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TemplateDeployment{s: s}, nil
}

// GetTemplateDeployment gets an existing TemplateDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateDeployment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TemplateDeploymentState, opts ...pulumi.ResourceOpt) (*TemplateDeployment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["deploymentMode"] = state.DeploymentMode
		inputs["name"] = state.Name
		inputs["outputs"] = state.Outputs
		inputs["parameters"] = state.Parameters
		inputs["parametersBody"] = state.ParametersBody
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["templateBody"] = state.TemplateBody
	}
	s, err := ctx.ReadResource("azure:core/templateDeployment:TemplateDeployment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TemplateDeployment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TemplateDeployment) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TemplateDeployment) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
// specified within the template, and this provider will not be aware of this.
func (r *TemplateDeployment) DeploymentMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["deploymentMode"])
}

// Specifies the name of the template deployment. Changing this forces a
// new resource to be created.
func (r *TemplateDeployment) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
func (r *TemplateDeployment) Outputs() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["outputs"])
}

// Specifies the name and value pairs that define the deployment parameters for the template.
func (r *TemplateDeployment) Parameters() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["parameters"])
}

// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
func (r *TemplateDeployment) ParametersBody() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["parametersBody"])
}

// The name of the resource group in which to
// create the template deployment.
func (r *TemplateDeployment) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the JSON definition for the template.
func (r *TemplateDeployment) TemplateBody() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["templateBody"])
}

// Input properties used for looking up and filtering TemplateDeployment resources.
type TemplateDeploymentState struct {
	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode interface{}
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name interface{}
	// A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
	Outputs interface{}
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters interface{}
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody interface{}
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName interface{}
	// Specifies the JSON definition for the template.
	TemplateBody interface{}
}

// The set of arguments for constructing a TemplateDeployment resource.
type TemplateDeploymentArgs struct {
	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode interface{}
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name interface{}
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters interface{}
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody interface{}
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName interface{}
	// Specifies the JSON definition for the template.
	TemplateBody interface{}
}
