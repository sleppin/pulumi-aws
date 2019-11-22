// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Resource for managing QuickSight Group
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/quicksight_group.html.markdown.
type Group struct {
	s *pulumi.ResourceState
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOpt) (*Group, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["awsAccountId"] = nil
		inputs["description"] = nil
		inputs["groupName"] = nil
		inputs["namespace"] = nil
	} else {
		inputs["awsAccountId"] = args.AwsAccountId
		inputs["description"] = args.Description
		inputs["groupName"] = args.GroupName
		inputs["namespace"] = args.Namespace
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:quicksight/group:Group", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Group{s: s}, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GroupState, opts ...pulumi.ResourceOpt) (*Group, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["awsAccountId"] = state.AwsAccountId
		inputs["description"] = state.Description
		inputs["groupName"] = state.GroupName
		inputs["namespace"] = state.Namespace
	}
	s, err := ctx.ReadResource("aws:quicksight/group:Group", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Group{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Group) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Group) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN) of group
func (r *Group) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
func (r *Group) AwsAccountId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["awsAccountId"])
}

// A description for the group.
func (r *Group) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// A name for the group.
func (r *Group) GroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["groupName"])
}

// The namespace. Currently, you should set this to `default`.
func (r *Group) Namespace() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namespace"])
}

// Input properties used for looking up and filtering Group resources.
type GroupState struct {
	// Amazon Resource Name (ARN) of group
	Arn interface{}
	// The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId interface{}
	// A description for the group.
	Description interface{}
	// A name for the group.
	GroupName interface{}
	// The namespace. Currently, you should set this to `default`.
	Namespace interface{}
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId interface{}
	// A description for the group.
	Description interface{}
	// A name for the group.
	GroupName interface{}
	// The namespace. Currently, you should set this to `default`.
	Namespace interface{}
}
