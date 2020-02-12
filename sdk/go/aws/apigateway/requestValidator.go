// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Gateway Request Validator.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_request_validator.html.markdown.
type RequestValidator struct {
	pulumi.CustomResourceState

	// The name of the request validator
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the associated Rest API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// Boolean whether to validate request body. Defaults to `false`.
	ValidateRequestBody pulumi.BoolPtrOutput `pulumi:"validateRequestBody"`
	// Boolean whether to validate request parameters. Defaults to `false`.
	ValidateRequestParameters pulumi.BoolPtrOutput `pulumi:"validateRequestParameters"`
}

// NewRequestValidator registers a new resource with the given unique name, arguments, and options.
func NewRequestValidator(ctx *pulumi.Context,
	name string, args *RequestValidatorArgs, opts ...pulumi.ResourceOption) (*RequestValidator, error) {
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil {
		args = &RequestValidatorArgs{}
	}
	var resource RequestValidator
	err := ctx.RegisterResource("aws:apigateway/requestValidator:RequestValidator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequestValidator gets an existing RequestValidator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestValidator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequestValidatorState, opts ...pulumi.ResourceOption) (*RequestValidator, error) {
	var resource RequestValidator
	err := ctx.ReadResource("aws:apigateway/requestValidator:RequestValidator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequestValidator resources.
type requestValidatorState struct {
	// The name of the request validator
	Name *string `pulumi:"name"`
	// The ID of the associated Rest API
	RestApi *string `pulumi:"restApi"`
	// Boolean whether to validate request body. Defaults to `false`.
	ValidateRequestBody *bool `pulumi:"validateRequestBody"`
	// Boolean whether to validate request parameters. Defaults to `false`.
	ValidateRequestParameters *bool `pulumi:"validateRequestParameters"`
}

type RequestValidatorState struct {
	// The name of the request validator
	Name pulumi.StringPtrInput
	// The ID of the associated Rest API
	RestApi pulumi.StringPtrInput
	// Boolean whether to validate request body. Defaults to `false`.
	ValidateRequestBody pulumi.BoolPtrInput
	// Boolean whether to validate request parameters. Defaults to `false`.
	ValidateRequestParameters pulumi.BoolPtrInput
}

func (RequestValidatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*requestValidatorState)(nil)).Elem()
}

type requestValidatorArgs struct {
	// The name of the request validator
	Name *string `pulumi:"name"`
	// The ID of the associated Rest API
	RestApi string `pulumi:"restApi"`
	// Boolean whether to validate request body. Defaults to `false`.
	ValidateRequestBody *bool `pulumi:"validateRequestBody"`
	// Boolean whether to validate request parameters. Defaults to `false`.
	ValidateRequestParameters *bool `pulumi:"validateRequestParameters"`
}

// The set of arguments for constructing a RequestValidator resource.
type RequestValidatorArgs struct {
	// The name of the request validator
	Name pulumi.StringPtrInput
	// The ID of the associated Rest API
	RestApi pulumi.StringInput
	// Boolean whether to validate request body. Defaults to `false`.
	ValidateRequestBody pulumi.BoolPtrInput
	// Boolean whether to validate request parameters. Defaults to `false`.
	ValidateRequestParameters pulumi.BoolPtrInput
}

func (RequestValidatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requestValidatorArgs)(nil)).Elem()
}

