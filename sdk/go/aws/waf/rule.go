// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Rule Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_rule.html.markdown.
type Rule struct {
	s *pulumi.ResourceState
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOpt) (*Rule, error) {
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["metricName"] = nil
		inputs["name"] = nil
		inputs["predicates"] = nil
		inputs["tags"] = nil
	} else {
		inputs["metricName"] = args.MetricName
		inputs["name"] = args.Name
		inputs["predicates"] = args.Predicates
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:waf/rule:Rule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Rule{s: s}, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RuleState, opts ...pulumi.ResourceOpt) (*Rule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["metricName"] = state.MetricName
		inputs["name"] = state.Name
		inputs["predicates"] = state.Predicates
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:waf/rule:Rule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Rule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Rule) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Rule) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the WAF rule.
func (r *Rule) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
func (r *Rule) MetricName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["metricName"])
}

// The name or description of the rule.
func (r *Rule) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The objects to include in a rule (documented below).
func (r *Rule) Predicates() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["predicates"])
}

// Key-value mapping of resource tags
func (r *Rule) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Rule resources.
type RuleState struct {
	// The ARN of the WAF rule.
	Arn interface{}
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName interface{}
	// The name or description of the rule.
	Name interface{}
	// The objects to include in a rule (documented below).
	Predicates interface{}
	// Key-value mapping of resource tags
	Tags interface{}
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName interface{}
	// The name or description of the rule.
	Name interface{}
	// The objects to include in a rule (documented below).
	Predicates interface{}
	// Key-value mapping of resource tags
	Tags interface{}
}
