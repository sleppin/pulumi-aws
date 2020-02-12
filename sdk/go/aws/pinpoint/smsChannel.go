// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pinpoint

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Pinpoint SMS Channel resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_sms_channel.html.markdown.
type SmsChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond pulumi.IntOutput `pulumi:"promotionalMessagesPerSecond"`
	// Sender identifier of your messages.
	SenderId pulumi.StringPtrOutput `pulumi:"senderId"`
	// The Short Code registered with the phone provider.
	ShortCode pulumi.StringPtrOutput `pulumi:"shortCode"`
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond pulumi.IntOutput `pulumi:"transactionalMessagesPerSecond"`
}

// NewSmsChannel registers a new resource with the given unique name, arguments, and options.
func NewSmsChannel(ctx *pulumi.Context,
	name string, args *SmsChannelArgs, opts ...pulumi.ResourceOption) (*SmsChannel, error) {
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil {
		args = &SmsChannelArgs{}
	}
	var resource SmsChannel
	err := ctx.RegisterResource("aws:pinpoint/smsChannel:SmsChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmsChannel gets an existing SmsChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmsChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmsChannelState, opts ...pulumi.ResourceOption) (*SmsChannel, error) {
	var resource SmsChannel
	err := ctx.ReadResource("aws:pinpoint/smsChannel:SmsChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmsChannel resources.
type smsChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond *int `pulumi:"promotionalMessagesPerSecond"`
	// Sender identifier of your messages.
	SenderId *string `pulumi:"senderId"`
	// The Short Code registered with the phone provider.
	ShortCode *string `pulumi:"shortCode"`
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond *int `pulumi:"transactionalMessagesPerSecond"`
}

type SmsChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond pulumi.IntPtrInput
	// Sender identifier of your messages.
	SenderId pulumi.StringPtrInput
	// The Short Code registered with the phone provider.
	ShortCode pulumi.StringPtrInput
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond pulumi.IntPtrInput
}

func (SmsChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*smsChannelState)(nil)).Elem()
}

type smsChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Sender identifier of your messages.
	SenderId *string `pulumi:"senderId"`
	// The Short Code registered with the phone provider.
	ShortCode *string `pulumi:"shortCode"`
}

// The set of arguments for constructing a SmsChannel resource.
type SmsChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Sender identifier of your messages.
	SenderId pulumi.StringPtrInput
	// The Short Code registered with the phone provider.
	ShortCode pulumi.StringPtrInput
}

func (SmsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smsChannelArgs)(nil)).Elem()
}

