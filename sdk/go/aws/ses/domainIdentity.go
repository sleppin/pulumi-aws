// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES domain identity resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_domain_identity.html.markdown.
type DomainIdentity struct {
	pulumi.CustomResourceState

	// The ARN of the domain identity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The domain name to assign to SES
	Domain pulumi.StringOutput `pulumi:"domain"`
	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done. See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.  Find out
	// more about verifying domains in Amazon SES in the [AWS SES
	// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
	VerificationToken pulumi.StringOutput `pulumi:"verificationToken"`
}

// NewDomainIdentity registers a new resource with the given unique name, arguments, and options.
func NewDomainIdentity(ctx *pulumi.Context,
	name string, args *DomainIdentityArgs, opts ...pulumi.ResourceOption) (*DomainIdentity, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil {
		args = &DomainIdentityArgs{}
	}
	var resource DomainIdentity
	err := ctx.RegisterResource("aws:ses/domainIdentity:DomainIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainIdentity gets an existing DomainIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainIdentityState, opts ...pulumi.ResourceOption) (*DomainIdentity, error) {
	var resource DomainIdentity
	err := ctx.ReadResource("aws:ses/domainIdentity:DomainIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainIdentity resources.
type domainIdentityState struct {
	// The ARN of the domain identity.
	Arn *string `pulumi:"arn"`
	// The domain name to assign to SES
	Domain *string `pulumi:"domain"`
	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done. See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.  Find out
	// more about verifying domains in Amazon SES in the [AWS SES
	// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
	VerificationToken *string `pulumi:"verificationToken"`
}

type DomainIdentityState struct {
	// The ARN of the domain identity.
	Arn pulumi.StringPtrInput
	// The domain name to assign to SES
	Domain pulumi.StringPtrInput
	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done. See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.  Find out
	// more about verifying domains in Amazon SES in the [AWS SES
	// docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html).
	VerificationToken pulumi.StringPtrInput
}

func (DomainIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainIdentityState)(nil)).Elem()
}

type domainIdentityArgs struct {
	// The domain name to assign to SES
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a DomainIdentity resource.
type DomainIdentityArgs struct {
	// The domain name to assign to SES
	Domain pulumi.StringInput
}

func (DomainIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainIdentityArgs)(nil)).Elem()
}

