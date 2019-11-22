// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a SES Identity Policy. More information about SES Sending Authorization Policies can be found in the [SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_identity_policy.html.markdown.
type IdentityPolicy struct {
	s *pulumi.ResourceState
}

// NewIdentityPolicy registers a new resource with the given unique name, arguments, and options.
func NewIdentityPolicy(ctx *pulumi.Context,
	name string, args *IdentityPolicyArgs, opts ...pulumi.ResourceOpt) (*IdentityPolicy, error) {
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["identity"] = nil
		inputs["name"] = nil
		inputs["policy"] = nil
	} else {
		inputs["identity"] = args.Identity
		inputs["name"] = args.Name
		inputs["policy"] = args.Policy
	}
	s, err := ctx.RegisterResource("aws:ses/identityPolicy:IdentityPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IdentityPolicy{s: s}, nil
}

// GetIdentityPolicy gets an existing IdentityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IdentityPolicyState, opts ...pulumi.ResourceOpt) (*IdentityPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["identity"] = state.Identity
		inputs["name"] = state.Name
		inputs["policy"] = state.Policy
	}
	s, err := ctx.ReadResource("aws:ses/identityPolicy:IdentityPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IdentityPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IdentityPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IdentityPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Name or Amazon Resource Name (ARN) of the SES Identity.
func (r *IdentityPolicy) Identity() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["identity"])
}

// Name of the policy.
func (r *IdentityPolicy) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *IdentityPolicy) Policy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policy"])
}

// Input properties used for looking up and filtering IdentityPolicy resources.
type IdentityPolicyState struct {
	// Name or Amazon Resource Name (ARN) of the SES Identity.
	Identity interface{}
	// Name of the policy.
	Name interface{}
	Policy interface{}
}

// The set of arguments for constructing a IdentityPolicy resource.
type IdentityPolicyArgs struct {
	// Name or Amazon Resource Name (ARN) of the SES Identity.
	Identity interface{}
	// Name of the policy.
	Name interface{}
	Policy interface{}
}
