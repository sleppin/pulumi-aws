// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM policy attached to a group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_group_policy.html.markdown.
type GroupPolicy struct {
	pulumi.CustomResourceState

	// The IAM group to attach to the policy.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicy(ctx *pulumi.Context,
	name string, args *GroupPolicyArgs, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	if args == nil || args.Group == nil {
		return nil, errors.New("missing required argument 'Group'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil {
		args = &GroupPolicyArgs{}
	}
	var resource GroupPolicy
	err := ctx.RegisterResource("aws:iam/groupPolicy:GroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicy gets an existing GroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyState, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	var resource GroupPolicy
	err := ctx.ReadResource("aws:iam/groupPolicy:GroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicy resources.
type groupPolicyState struct {
	// The IAM group to attach to the policy.
	Group *string `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	Policy *string `pulumi:"policy"`
}

type GroupPolicyState struct {
	// The IAM group to attach to the policy.
	Group pulumi.StringPtrInput
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	Policy pulumi.StringPtrInput
}

func (GroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyState)(nil)).Elem()
}

type groupPolicyArgs struct {
	// The IAM group to attach to the policy.
	Group string `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a GroupPolicy resource.
type GroupPolicyArgs struct {
	// The IAM group to attach to the policy.
	Group pulumi.StringInput
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	Policy pulumi.Input
}

func (GroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyArgs)(nil)).Elem()
}

