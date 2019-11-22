// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AWS Storage Gateway SMB File Share.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/storagegateway_smb_file_share.html.markdown.
type SmbFileShare struct {
	s *pulumi.ResourceState
}

// NewSmbFileShare registers a new resource with the given unique name, arguments, and options.
func NewSmbFileShare(ctx *pulumi.Context,
	name string, args *SmbFileShareArgs, opts ...pulumi.ResourceOpt) (*SmbFileShare, error) {
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	if args == nil || args.LocationArn == nil {
		return nil, errors.New("missing required argument 'LocationArn'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["authentication"] = nil
		inputs["defaultStorageClass"] = nil
		inputs["gatewayArn"] = nil
		inputs["guessMimeTypeEnabled"] = nil
		inputs["invalidUserLists"] = nil
		inputs["kmsEncrypted"] = nil
		inputs["kmsKeyArn"] = nil
		inputs["locationArn"] = nil
		inputs["objectAcl"] = nil
		inputs["readOnly"] = nil
		inputs["requesterPays"] = nil
		inputs["roleArn"] = nil
		inputs["tags"] = nil
		inputs["validUserLists"] = nil
	} else {
		inputs["authentication"] = args.Authentication
		inputs["defaultStorageClass"] = args.DefaultStorageClass
		inputs["gatewayArn"] = args.GatewayArn
		inputs["guessMimeTypeEnabled"] = args.GuessMimeTypeEnabled
		inputs["invalidUserLists"] = args.InvalidUserLists
		inputs["kmsEncrypted"] = args.KmsEncrypted
		inputs["kmsKeyArn"] = args.KmsKeyArn
		inputs["locationArn"] = args.LocationArn
		inputs["objectAcl"] = args.ObjectAcl
		inputs["readOnly"] = args.ReadOnly
		inputs["requesterPays"] = args.RequesterPays
		inputs["roleArn"] = args.RoleArn
		inputs["tags"] = args.Tags
		inputs["validUserLists"] = args.ValidUserLists
	}
	inputs["arn"] = nil
	inputs["fileshareId"] = nil
	s, err := ctx.RegisterResource("aws:storagegateway/smbFileShare:SmbFileShare", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SmbFileShare{s: s}, nil
}

// GetSmbFileShare gets an existing SmbFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmbFileShare(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SmbFileShareState, opts ...pulumi.ResourceOpt) (*SmbFileShare, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["authentication"] = state.Authentication
		inputs["defaultStorageClass"] = state.DefaultStorageClass
		inputs["fileshareId"] = state.FileshareId
		inputs["gatewayArn"] = state.GatewayArn
		inputs["guessMimeTypeEnabled"] = state.GuessMimeTypeEnabled
		inputs["invalidUserLists"] = state.InvalidUserLists
		inputs["kmsEncrypted"] = state.KmsEncrypted
		inputs["kmsKeyArn"] = state.KmsKeyArn
		inputs["locationArn"] = state.LocationArn
		inputs["objectAcl"] = state.ObjectAcl
		inputs["readOnly"] = state.ReadOnly
		inputs["requesterPays"] = state.RequesterPays
		inputs["roleArn"] = state.RoleArn
		inputs["tags"] = state.Tags
		inputs["validUserLists"] = state.ValidUserLists
	}
	s, err := ctx.ReadResource("aws:storagegateway/smbFileShare:SmbFileShare", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SmbFileShare{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SmbFileShare) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SmbFileShare) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN) of the SMB File Share.
func (r *SmbFileShare) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
func (r *SmbFileShare) Authentication() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["authentication"])
}

// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
func (r *SmbFileShare) DefaultStorageClass() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultStorageClass"])
}

// ID of the SMB File Share.
func (r *SmbFileShare) FileshareId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["fileshareId"])
}

// Amazon Resource Name (ARN) of the file gateway.
func (r *SmbFileShare) GatewayArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["gatewayArn"])
}

// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
func (r *SmbFileShare) GuessMimeTypeEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["guessMimeTypeEnabled"])
}

// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
func (r *SmbFileShare) InvalidUserLists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["invalidUserLists"])
}

// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
func (r *SmbFileShare) KmsEncrypted() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["kmsEncrypted"])
}

// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
func (r *SmbFileShare) KmsKeyArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kmsKeyArn"])
}

// The ARN of the backed storage used for storing file data.
func (r *SmbFileShare) LocationArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["locationArn"])
}

// Access Control List permission for S3 bucket objects. Defaults to `private`.
func (r *SmbFileShare) ObjectAcl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["objectAcl"])
}

// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
func (r *SmbFileShare) ReadOnly() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["readOnly"])
}

// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
func (r *SmbFileShare) RequesterPays() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["requesterPays"])
}

// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
func (r *SmbFileShare) RoleArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleArn"])
}

// Key-value mapping of resource tags
func (r *SmbFileShare) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
func (r *SmbFileShare) ValidUserLists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["validUserLists"])
}

// Input properties used for looking up and filtering SmbFileShare resources.
type SmbFileShareState struct {
	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn interface{}
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication interface{}
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass interface{}
	// ID of the SMB File Share.
	FileshareId interface{}
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn interface{}
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled interface{}
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists interface{}
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted interface{}
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn interface{}
	// The ARN of the backed storage used for storing file data.
	LocationArn interface{}
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl interface{}
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly interface{}
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays interface{}
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn interface{}
	// Key-value mapping of resource tags
	Tags interface{}
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists interface{}
}

// The set of arguments for constructing a SmbFileShare resource.
type SmbFileShareArgs struct {
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication interface{}
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass interface{}
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn interface{}
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled interface{}
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists interface{}
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted interface{}
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn interface{}
	// The ARN of the backed storage used for storing file data.
	LocationArn interface{}
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl interface{}
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly interface{}
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays interface{}
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn interface{}
	// Key-value mapping of resource tags
	Tags interface{}
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists interface{}
}
