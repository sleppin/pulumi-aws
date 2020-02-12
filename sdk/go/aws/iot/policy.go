// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IoT policy.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_policy.html.markdown.
type Policy struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The default version of this policy.
	DefaultVersionId pulumi.StringOutput `pulumi:"defaultVersionId"`
	// The name of the policy.
	Name pulumi.StringOutput `pulumi:"name"`
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("aws:iot/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws:iot/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// The ARN assigned by AWS to this policy.
	Arn *string `pulumi:"arn"`
	// The default version of this policy.
	DefaultVersionId *string `pulumi:"defaultVersionId"`
	// The name of the policy.
	Name *string `pulumi:"name"`
	Policy *string `pulumi:"policy"`
}

type PolicyState struct {
	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringPtrInput
	// The default version of this policy.
	DefaultVersionId pulumi.StringPtrInput
	// The name of the policy.
	Name pulumi.StringPtrInput
	Policy pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// The name of the policy.
	Name *string `pulumi:"name"`
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// The name of the policy.
	Name pulumi.StringPtrInput
	Policy pulumi.StringInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

