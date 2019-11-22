// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ECS cluster.
// 
// ## setting
// 
// The `setting` configuration block supports the following:
// 
// * `name` - (Required) Name of the setting to manage. Valid values: `containerInsights`.
// * `value` -  (Required) The value to assign to the setting. Value values are `enabled` and `disabled`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_cluster.html.markdown.
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["settings"] = nil
		inputs["tags"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["settings"] = args.Settings
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:ecs/cluster:Cluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["settings"] = state.Settings
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:ecs/cluster:Cluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Cluster) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Cluster) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) that identifies the cluster
func (r *Cluster) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
func (r *Cluster) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
func (r *Cluster) Settings() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["settings"])
}

// Key-value mapping of resource tags
func (r *Cluster) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// The Amazon Resource Name (ARN) that identifies the cluster
	Arn interface{}
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name interface{}
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings interface{}
	// Key-value mapping of resource tags
	Tags interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name interface{}
	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
	Settings interface{}
	// Key-value mapping of resource tags
	Tags interface{}
}
