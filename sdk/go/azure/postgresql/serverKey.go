// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Customer Managed Key for a PostgreSQL Server.
//
// ## Import
//
// A PostgreSQL Server Key can be imported using the `resource id` of the PostgreSQL Server Key, e.g.
//
// ```sh
//  $ pulumi import azure:postgresql/serverKey:ServerKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
// ```
type ServerKey struct {
	pulumi.CustomResourceState

	// The URL to a Key Vault Key.
	KeyVaultKeyId pulumi.StringOutput `pulumi:"keyVaultKeyId"`
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
}

// NewServerKey registers a new resource with the given unique name, arguments, and options.
func NewServerKey(ctx *pulumi.Context,
	name string, args *ServerKeyArgs, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultKeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultKeyId'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	var resource ServerKey
	err := ctx.RegisterResource("azure:postgresql/serverKey:ServerKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerKey gets an existing ServerKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerKeyState, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	var resource ServerKey
	err := ctx.ReadResource("azure:postgresql/serverKey:ServerKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerKey resources.
type serverKeyState struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId *string `pulumi:"keyVaultKeyId"`
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId *string `pulumi:"serverId"`
}

type ServerKeyState struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId pulumi.StringPtrInput
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId pulumi.StringPtrInput
}

func (ServerKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyState)(nil)).Elem()
}

type serverKeyArgs struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId string `pulumi:"keyVaultKeyId"`
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId string `pulumi:"serverId"`
}

// The set of arguments for constructing a ServerKey resource.
type ServerKeyArgs struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId pulumi.StringInput
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId pulumi.StringInput
}

func (ServerKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyArgs)(nil)).Elem()
}

type ServerKeyInput interface {
	pulumi.Input

	ToServerKeyOutput() ServerKeyOutput
	ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput
}

func (*ServerKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKey)(nil))
}

func (i *ServerKey) ToServerKeyOutput() ServerKeyOutput {
	return i.ToServerKeyOutputWithContext(context.Background())
}

func (i *ServerKey) ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyOutput)
}

func (i *ServerKey) ToServerKeyPtrOutput() ServerKeyPtrOutput {
	return i.ToServerKeyPtrOutputWithContext(context.Background())
}

func (i *ServerKey) ToServerKeyPtrOutputWithContext(ctx context.Context) ServerKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyPtrOutput)
}

type ServerKeyPtrInput interface {
	pulumi.Input

	ToServerKeyPtrOutput() ServerKeyPtrOutput
	ToServerKeyPtrOutputWithContext(ctx context.Context) ServerKeyPtrOutput
}

type serverKeyPtrType ServerKeyArgs

func (*serverKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKey)(nil))
}

func (i *serverKeyPtrType) ToServerKeyPtrOutput() ServerKeyPtrOutput {
	return i.ToServerKeyPtrOutputWithContext(context.Background())
}

func (i *serverKeyPtrType) ToServerKeyPtrOutputWithContext(ctx context.Context) ServerKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyPtrOutput)
}

// ServerKeyArrayInput is an input type that accepts ServerKeyArray and ServerKeyArrayOutput values.
// You can construct a concrete instance of `ServerKeyArrayInput` via:
//
//          ServerKeyArray{ ServerKeyArgs{...} }
type ServerKeyArrayInput interface {
	pulumi.Input

	ToServerKeyArrayOutput() ServerKeyArrayOutput
	ToServerKeyArrayOutputWithContext(context.Context) ServerKeyArrayOutput
}

type ServerKeyArray []ServerKeyInput

func (ServerKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServerKey)(nil))
}

func (i ServerKeyArray) ToServerKeyArrayOutput() ServerKeyArrayOutput {
	return i.ToServerKeyArrayOutputWithContext(context.Background())
}

func (i ServerKeyArray) ToServerKeyArrayOutputWithContext(ctx context.Context) ServerKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyArrayOutput)
}

// ServerKeyMapInput is an input type that accepts ServerKeyMap and ServerKeyMapOutput values.
// You can construct a concrete instance of `ServerKeyMapInput` via:
//
//          ServerKeyMap{ "key": ServerKeyArgs{...} }
type ServerKeyMapInput interface {
	pulumi.Input

	ToServerKeyMapOutput() ServerKeyMapOutput
	ToServerKeyMapOutputWithContext(context.Context) ServerKeyMapOutput
}

type ServerKeyMap map[string]ServerKeyInput

func (ServerKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServerKey)(nil))
}

func (i ServerKeyMap) ToServerKeyMapOutput() ServerKeyMapOutput {
	return i.ToServerKeyMapOutputWithContext(context.Background())
}

func (i ServerKeyMap) ToServerKeyMapOutputWithContext(ctx context.Context) ServerKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyMapOutput)
}

type ServerKeyOutput struct {
	*pulumi.OutputState
}

func (ServerKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKey)(nil))
}

func (o ServerKeyOutput) ToServerKeyOutput() ServerKeyOutput {
	return o
}

func (o ServerKeyOutput) ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput {
	return o
}

func (o ServerKeyOutput) ToServerKeyPtrOutput() ServerKeyPtrOutput {
	return o.ToServerKeyPtrOutputWithContext(context.Background())
}

func (o ServerKeyOutput) ToServerKeyPtrOutputWithContext(ctx context.Context) ServerKeyPtrOutput {
	return o.ApplyT(func(v ServerKey) *ServerKey {
		return &v
	}).(ServerKeyPtrOutput)
}

type ServerKeyPtrOutput struct {
	*pulumi.OutputState
}

func (ServerKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKey)(nil))
}

func (o ServerKeyPtrOutput) ToServerKeyPtrOutput() ServerKeyPtrOutput {
	return o
}

func (o ServerKeyPtrOutput) ToServerKeyPtrOutputWithContext(ctx context.Context) ServerKeyPtrOutput {
	return o
}

type ServerKeyArrayOutput struct{ *pulumi.OutputState }

func (ServerKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerKey)(nil))
}

func (o ServerKeyArrayOutput) ToServerKeyArrayOutput() ServerKeyArrayOutput {
	return o
}

func (o ServerKeyArrayOutput) ToServerKeyArrayOutputWithContext(ctx context.Context) ServerKeyArrayOutput {
	return o
}

func (o ServerKeyArrayOutput) Index(i pulumi.IntInput) ServerKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerKey {
		return vs[0].([]ServerKey)[vs[1].(int)]
	}).(ServerKeyOutput)
}

type ServerKeyMapOutput struct{ *pulumi.OutputState }

func (ServerKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServerKey)(nil))
}

func (o ServerKeyMapOutput) ToServerKeyMapOutput() ServerKeyMapOutput {
	return o
}

func (o ServerKeyMapOutput) ToServerKeyMapOutputWithContext(ctx context.Context) ServerKeyMapOutput {
	return o
}

func (o ServerKeyMapOutput) MapIndex(k pulumi.StringInput) ServerKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServerKey {
		return vs[0].(map[string]ServerKey)[vs[1].(string)]
	}).(ServerKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerKeyOutput{})
	pulumi.RegisterOutputType(ServerKeyPtrOutput{})
	pulumi.RegisterOutputType(ServerKeyArrayOutput{})
	pulumi.RegisterOutputType(ServerKeyMapOutput{})
}
