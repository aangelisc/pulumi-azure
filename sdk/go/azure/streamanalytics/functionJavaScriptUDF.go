// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a JavaScript UDF Function within Stream Analytics Streaming Job.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_function_javascript_udf.html.markdown.
type FunctionJavaScriptUDF struct {
	s *pulumi.ResourceState
}

// NewFunctionJavaScriptUDF registers a new resource with the given unique name, arguments, and options.
func NewFunctionJavaScriptUDF(ctx *pulumi.Context,
	name string, args *FunctionJavaScriptUDFArgs, opts ...pulumi.ResourceOpt) (*FunctionJavaScriptUDF, error) {
	if args == nil || args.Inputs == nil {
		return nil, errors.New("missing required argument 'Inputs'")
	}
	if args == nil || args.Output == nil {
		return nil, errors.New("missing required argument 'Output'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Script == nil {
		return nil, errors.New("missing required argument 'Script'")
	}
	if args == nil || args.StreamAnalyticsJobName == nil {
		return nil, errors.New("missing required argument 'StreamAnalyticsJobName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["inputs"] = nil
		inputs["name"] = nil
		inputs["output"] = nil
		inputs["resourceGroupName"] = nil
		inputs["script"] = nil
		inputs["streamAnalyticsJobName"] = nil
	} else {
		inputs["inputs"] = args.Inputs
		inputs["name"] = args.Name
		inputs["output"] = args.Output
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["script"] = args.Script
		inputs["streamAnalyticsJobName"] = args.StreamAnalyticsJobName
	}
	s, err := ctx.RegisterResource("azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FunctionJavaScriptUDF{s: s}, nil
}

// GetFunctionJavaScriptUDF gets an existing FunctionJavaScriptUDF resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionJavaScriptUDF(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FunctionJavaScriptUDFState, opts ...pulumi.ResourceOpt) (*FunctionJavaScriptUDF, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["inputs"] = state.Inputs
		inputs["name"] = state.Name
		inputs["output"] = state.Output
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["script"] = state.Script
		inputs["streamAnalyticsJobName"] = state.StreamAnalyticsJobName
	}
	s, err := ctx.ReadResource("azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FunctionJavaScriptUDF{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FunctionJavaScriptUDF) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FunctionJavaScriptUDF) ID() pulumi.IDOutput {
	return r.s.ID()
}

// One or more `input` blocks as defined below.
func (r *FunctionJavaScriptUDF) Inputs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["inputs"])
}

// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
func (r *FunctionJavaScriptUDF) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// An `output` blocks as defined below.
func (r *FunctionJavaScriptUDF) Output() pulumi.Output {
	return r.s.State["output"]
}

// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
func (r *FunctionJavaScriptUDF) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The JavaScript of this UDF Function.
func (r *FunctionJavaScriptUDF) Script() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["script"])
}

// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
func (r *FunctionJavaScriptUDF) StreamAnalyticsJobName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["streamAnalyticsJobName"])
}

// Input properties used for looking up and filtering FunctionJavaScriptUDF resources.
type FunctionJavaScriptUDFState struct {
	// One or more `input` blocks as defined below.
	Inputs interface{}
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name interface{}
	// An `output` blocks as defined below.
	Output interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The JavaScript of this UDF Function.
	Script interface{}
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName interface{}
}

// The set of arguments for constructing a FunctionJavaScriptUDF resource.
type FunctionJavaScriptUDFArgs struct {
	// One or more `input` blocks as defined below.
	Inputs interface{}
	// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
	Name interface{}
	// An `output` blocks as defined below.
	Output interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The JavaScript of this UDF Function.
	Script interface{}
	// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
	StreamAnalyticsJobName interface{}
}
