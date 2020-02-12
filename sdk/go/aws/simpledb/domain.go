// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package simpledb

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SimpleDB domain resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/simpledb_domain.html.markdown.
type Domain struct {
	pulumi.CustomResourceState

	// The name of the SimpleDB domain
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		args = &DomainArgs{}
	}
	var resource Domain
	err := ctx.RegisterResource("aws:simpledb/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws:simpledb/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The name of the SimpleDB domain
	Name *string `pulumi:"name"`
}

type DomainState struct {
	// The name of the SimpleDB domain
	Name pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The name of the SimpleDB domain
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The name of the SimpleDB domain
	Name pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

