// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an RDS DB cluster parameter group resource. Documentation of the available parameters for various Aurora engines can be found at:
//
// * [Aurora MySQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Reference.html)
// * [Aurora PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraPostgreSQL.Reference.html)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_cluster_parameter_group.html.markdown.
type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	// The ARN of the db cluster parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The family of the DB cluster parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the DB parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	if args == nil {
		args = &ClusterParameterGroupArgs{}
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws:rds/clusterParameterGroup:ClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterParameterGroup gets an existing ClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterParameterGroupState, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	var resource ClusterParameterGroup
	err := ctx.ReadResource("aws:rds/clusterParameterGroup:ClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterParameterGroup resources.
type clusterParameterGroupState struct {
	// The ARN of the db cluster parameter group.
	Arn *string `pulumi:"arn"`
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the DB cluster parameter group.
	Family *string `pulumi:"family"`
	// The name of the DB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ClusterParameterGroupState struct {
	// The ARN of the db cluster parameter group.
	Arn pulumi.StringPtrInput
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the DB cluster parameter group.
	Family pulumi.StringPtrInput
	// The name of the DB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters ClusterParameterGroupParameterArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the DB cluster parameter group.
	Family string `pulumi:"family"`
	// The name of the DB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the DB cluster parameter group.
	Family pulumi.StringInput
	// The name of the DB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters ClusterParameterGroupParameterArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}

