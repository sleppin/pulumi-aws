// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Amazon Inspector resource group resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/inspector_resource_group.html.markdown.
type ResourceGroup struct {
	s *pulumi.ResourceState
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOpt) (*ResourceGroup, error) {
	if args == nil || args.Tags == nil {
		return nil, errors.New("missing required argument 'Tags'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["tags"] = nil
	} else {
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:inspector/resourceGroup:ResourceGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ResourceGroup{s: s}, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ResourceGroupState, opts ...pulumi.ResourceOpt) (*ResourceGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:inspector/resourceGroup:ResourceGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ResourceGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ResourceGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ResourceGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The resource group ARN.
func (r *ResourceGroup) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
func (r *ResourceGroup) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering ResourceGroup resources.
type ResourceGroupState struct {
	// The resource group ARN.
	Arn interface{}
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags interface{}
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// Key-value map of tags that are used to select the EC2 instances to be included in an [Amazon Inspector assessment target](https://www.terraform.io/docs/providers/aws/r/inspector_assessment_target.html).
	Tags interface{}
}
