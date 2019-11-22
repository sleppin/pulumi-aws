// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a KMS Customer Master Key that uses external key material. To instead manage a KMS Customer Master Key where AWS automatically generates and potentially rotates key material, see the [`kms.Key` resource](https://www.terraform.io/docs/providers/aws/r/kms_key.html).
// 
// > **Note:** All arguments including the key material will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_external_key.html.markdown.
type ExternalKey struct {
	s *pulumi.ResourceState
}

// NewExternalKey registers a new resource with the given unique name, arguments, and options.
func NewExternalKey(ctx *pulumi.Context,
	name string, args *ExternalKeyArgs, opts ...pulumi.ResourceOpt) (*ExternalKey, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["deletionWindowInDays"] = nil
		inputs["description"] = nil
		inputs["enabled"] = nil
		inputs["keyMaterialBase64"] = nil
		inputs["policy"] = nil
		inputs["tags"] = nil
		inputs["validTo"] = nil
	} else {
		inputs["deletionWindowInDays"] = args.DeletionWindowInDays
		inputs["description"] = args.Description
		inputs["enabled"] = args.Enabled
		inputs["keyMaterialBase64"] = args.KeyMaterialBase64
		inputs["policy"] = args.Policy
		inputs["tags"] = args.Tags
		inputs["validTo"] = args.ValidTo
	}
	inputs["arn"] = nil
	inputs["expirationModel"] = nil
	inputs["keyState"] = nil
	inputs["keyUsage"] = nil
	s, err := ctx.RegisterResource("aws:kms/externalKey:ExternalKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExternalKey{s: s}, nil
}

// GetExternalKey gets an existing ExternalKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExternalKeyState, opts ...pulumi.ResourceOpt) (*ExternalKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["deletionWindowInDays"] = state.DeletionWindowInDays
		inputs["description"] = state.Description
		inputs["enabled"] = state.Enabled
		inputs["expirationModel"] = state.ExpirationModel
		inputs["keyMaterialBase64"] = state.KeyMaterialBase64
		inputs["keyState"] = state.KeyState
		inputs["keyUsage"] = state.KeyUsage
		inputs["policy"] = state.Policy
		inputs["tags"] = state.Tags
		inputs["validTo"] = state.ValidTo
	}
	s, err := ctx.ReadResource("aws:kms/externalKey:ExternalKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExternalKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ExternalKey) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ExternalKey) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) of the key.
func (r *ExternalKey) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
func (r *ExternalKey) DeletionWindowInDays() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["deletionWindowInDays"])
}

// Description of the key.
func (r *ExternalKey) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
func (r *ExternalKey) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
func (r *ExternalKey) ExpirationModel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["expirationModel"])
}

// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
func (r *ExternalKey) KeyMaterialBase64() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyMaterialBase64"])
}

// The state of the CMK.
func (r *ExternalKey) KeyState() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyState"])
}

// The cryptographic operations for which you can use the CMK.
func (r *ExternalKey) KeyUsage() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyUsage"])
}

// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
func (r *ExternalKey) Policy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policy"])
}

// A key-value map of tags to assign to the key.
func (r *ExternalKey) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
func (r *ExternalKey) ValidTo() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["validTo"])
}

// Input properties used for looking up and filtering ExternalKey resources.
type ExternalKeyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn interface{}
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays interface{}
	// Description of the key.
	Description interface{}
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled interface{}
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel interface{}
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 interface{}
	// The state of the CMK.
	KeyState interface{}
	// The cryptographic operations for which you can use the CMK.
	KeyUsage interface{}
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy interface{}
	// A key-value map of tags to assign to the key.
	Tags interface{}
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo interface{}
}

// The set of arguments for constructing a ExternalKey resource.
type ExternalKeyArgs struct {
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays interface{}
	// Description of the key.
	Description interface{}
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled interface{}
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 interface{}
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy interface{}
	// A key-value map of tags to assign to the key.
	Tags interface{}
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo interface{}
}
