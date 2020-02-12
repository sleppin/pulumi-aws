// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package inspector

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Amazon Inspector resource group resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/inspector_resource_group.html.markdown.
type ResourceGroup struct {
	pulumi.CustomResourceState

	// The resource group ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil || args.Tags == nil {
		return nil, errors.New("missing required argument 'Tags'")
	}
	if args == nil {
		args = &ResourceGroupArgs{}
	}
	var resource ResourceGroup
	err := ctx.RegisterResource("aws:inspector/resourceGroup:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("aws:inspector/resourceGroup:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroup resources.
type resourceGroupState struct {
	// The resource group ARN.
	Arn *string `pulumi:"arn"`
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags map[string]interface{} `pulumi:"tags"`
}

type ResourceGroupState struct {
	// The resource group ARN.
	Arn pulumi.StringPtrInput
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags pulumi.MapInput
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags pulumi.MapInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

