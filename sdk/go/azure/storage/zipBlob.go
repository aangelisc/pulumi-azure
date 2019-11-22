// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ZipBlob struct {
	s *pulumi.ResourceState
}

// NewZipBlob registers a new resource with the given unique name, arguments, and options.
func NewZipBlob(ctx *pulumi.Context,
	name string, args *ZipBlobArgs, opts ...pulumi.ResourceOpt) (*ZipBlob, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	if args == nil || args.StorageContainerName == nil {
		return nil, errors.New("missing required argument 'StorageContainerName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessTier"] = nil
		inputs["attempts"] = nil
		inputs["contentType"] = nil
		inputs["metadata"] = nil
		inputs["name"] = nil
		inputs["parallelism"] = nil
		inputs["resourceGroupName"] = nil
		inputs["size"] = nil
		inputs["content"] = nil
		inputs["sourceContent"] = nil
		inputs["sourceUri"] = nil
		inputs["storageAccountName"] = nil
		inputs["storageContainerName"] = nil
		inputs["type"] = nil
	} else {
		inputs["accessTier"] = args.AccessTier
		inputs["attempts"] = args.Attempts
		inputs["contentType"] = args.ContentType
		inputs["metadata"] = args.Metadata
		inputs["name"] = args.Name
		inputs["parallelism"] = args.Parallelism
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["size"] = args.Size
		inputs["content"] = args.Content
		inputs["sourceContent"] = args.SourceContent
		inputs["sourceUri"] = args.SourceUri
		inputs["storageAccountName"] = args.StorageAccountName
		inputs["storageContainerName"] = args.StorageContainerName
		inputs["type"] = args.Type
	}
	inputs["url"] = nil
	s, err := ctx.RegisterResource("azure:storage/zipBlob:ZipBlob", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ZipBlob{s: s}, nil
}

// GetZipBlob gets an existing ZipBlob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZipBlob(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ZipBlobState, opts ...pulumi.ResourceOpt) (*ZipBlob, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessTier"] = state.AccessTier
		inputs["attempts"] = state.Attempts
		inputs["contentType"] = state.ContentType
		inputs["metadata"] = state.Metadata
		inputs["name"] = state.Name
		inputs["parallelism"] = state.Parallelism
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["size"] = state.Size
		inputs["content"] = state.Content
		inputs["sourceContent"] = state.SourceContent
		inputs["sourceUri"] = state.SourceUri
		inputs["storageAccountName"] = state.StorageAccountName
		inputs["storageContainerName"] = state.StorageContainerName
		inputs["type"] = state.Type
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("azure:storage/zipBlob:ZipBlob", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ZipBlob{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ZipBlob) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ZipBlob) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *ZipBlob) AccessTier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessTier"])
}

func (r *ZipBlob) Attempts() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["attempts"])
}

func (r *ZipBlob) ContentType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["contentType"])
}

func (r *ZipBlob) Metadata() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["metadata"])
}

func (r *ZipBlob) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *ZipBlob) Parallelism() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["parallelism"])
}

func (r *ZipBlob) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

func (r *ZipBlob) Size() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["size"])
}

func (r *ZipBlob) Content() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["content"])
}

func (r *ZipBlob) SourceContent() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceContent"])
}

func (r *ZipBlob) SourceUri() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceUri"])
}

func (r *ZipBlob) StorageAccountName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageAccountName"])
}

func (r *ZipBlob) StorageContainerName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageContainerName"])
}

func (r *ZipBlob) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

func (r *ZipBlob) Url() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering ZipBlob resources.
type ZipBlobState struct {
	AccessTier interface{}
	Attempts interface{}
	ContentType interface{}
	Metadata interface{}
	Name interface{}
	Parallelism interface{}
	ResourceGroupName interface{}
	Size interface{}
	Content interface{}
	SourceContent interface{}
	SourceUri interface{}
	StorageAccountName interface{}
	StorageContainerName interface{}
	Type interface{}
	Url interface{}
}

// The set of arguments for constructing a ZipBlob resource.
type ZipBlobArgs struct {
	AccessTier interface{}
	Attempts interface{}
	ContentType interface{}
	Metadata interface{}
	Name interface{}
	Parallelism interface{}
	ResourceGroupName interface{}
	Size interface{}
	Content interface{}
	SourceContent interface{}
	SourceUri interface{}
	StorageAccountName interface{}
	StorageContainerName interface{}
	Type interface{}
}
