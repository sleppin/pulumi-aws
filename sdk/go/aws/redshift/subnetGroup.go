// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package redshift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new Amazon Redshift subnet group. You must provide a list of one or more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when creating Amazon Redshift subnet group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_subnet_group.html.markdown.
type SubnetGroup struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Redshift Subnet group name
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the Redshift Subnet group.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of VPC subnet IDs.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	if args == nil {
		args = &SubnetGroupArgs{}
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:redshift/subnetGroup:SubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetGroup gets an existing SubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetGroupState, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	var resource SubnetGroup
	err := ctx.ReadResource("aws:redshift/subnetGroup:SubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetGroup resources.
type subnetGroupState struct {
	// Amazon Resource Name (ARN) of the Redshift Subnet group name
	Arn *string `pulumi:"arn"`
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The name of the Redshift Subnet group.
	Name *string `pulumi:"name"`
	// An array of VPC subnet IDs.
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type SubnetGroupState struct {
	// Amazon Resource Name (ARN) of the Redshift Subnet group name
	Arn pulumi.StringPtrInput
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The name of the Redshift Subnet group.
	Name pulumi.StringPtrInput
	// An array of VPC subnet IDs.
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (SubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupState)(nil)).Elem()
}

type subnetGroupArgs struct {
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The name of the Redshift Subnet group.
	Name *string `pulumi:"name"`
	// An array of VPC subnet IDs.
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The name of the Redshift Subnet group.
	Name pulumi.StringPtrInput
	// An array of VPC subnet IDs.
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (SubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupArgs)(nil)).Elem()
}

