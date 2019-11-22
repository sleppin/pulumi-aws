// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES receipt filter resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_receipt_filter.html.markdown.
type ReceiptFilter struct {
	s *pulumi.ResourceState
}

// NewReceiptFilter registers a new resource with the given unique name, arguments, and options.
func NewReceiptFilter(ctx *pulumi.Context,
	name string, args *ReceiptFilterArgs, opts ...pulumi.ResourceOpt) (*ReceiptFilter, error) {
	if args == nil || args.Cidr == nil {
		return nil, errors.New("missing required argument 'Cidr'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cidr"] = nil
		inputs["name"] = nil
		inputs["policy"] = nil
	} else {
		inputs["cidr"] = args.Cidr
		inputs["name"] = args.Name
		inputs["policy"] = args.Policy
	}
	s, err := ctx.RegisterResource("aws:ses/receiptFilter:ReceiptFilter", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReceiptFilter{s: s}, nil
}

// GetReceiptFilter gets an existing ReceiptFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptFilter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ReceiptFilterState, opts ...pulumi.ResourceOpt) (*ReceiptFilter, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cidr"] = state.Cidr
		inputs["name"] = state.Name
		inputs["policy"] = state.Policy
	}
	s, err := ctx.ReadResource("aws:ses/receiptFilter:ReceiptFilter", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReceiptFilter{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ReceiptFilter) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ReceiptFilter) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The IP address or address range to filter, in CIDR notation
func (r *ReceiptFilter) Cidr() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cidr"])
}

// The name of the filter
func (r *ReceiptFilter) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Block or Allow
func (r *ReceiptFilter) Policy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policy"])
}

// Input properties used for looking up and filtering ReceiptFilter resources.
type ReceiptFilterState struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr interface{}
	// The name of the filter
	Name interface{}
	// Block or Allow
	Policy interface{}
}

// The set of arguments for constructing a ReceiptFilter resource.
type ReceiptFilterArgs struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr interface{}
	// The name of the filter
	Name interface{}
	// Block or Allow
	Policy interface{}
}
