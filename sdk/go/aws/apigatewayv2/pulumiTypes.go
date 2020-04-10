// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AuthorizerJwtConfiguration struct {
	// A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
	Audiences []string `pulumi:"audiences"`
	// The base domain of the identity provider that issues JSON Web Tokens, such as the `endpoint` attribute of the [`cognito.UserPool`](https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html) resource.
	Issuer *string `pulumi:"issuer"`
}

type AuthorizerJwtConfigurationInput interface {
	pulumi.Input

	ToAuthorizerJwtConfigurationOutput() AuthorizerJwtConfigurationOutput
	ToAuthorizerJwtConfigurationOutputWithContext(context.Context) AuthorizerJwtConfigurationOutput
}

type AuthorizerJwtConfigurationArgs struct {
	// A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
	Audiences pulumi.StringArrayInput `pulumi:"audiences"`
	// The base domain of the identity provider that issues JSON Web Tokens, such as the `endpoint` attribute of the [`cognito.UserPool`](https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html) resource.
	Issuer pulumi.StringPtrInput `pulumi:"issuer"`
}

func (AuthorizerJwtConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizerJwtConfiguration)(nil)).Elem()
}

func (i AuthorizerJwtConfigurationArgs) ToAuthorizerJwtConfigurationOutput() AuthorizerJwtConfigurationOutput {
	return i.ToAuthorizerJwtConfigurationOutputWithContext(context.Background())
}

func (i AuthorizerJwtConfigurationArgs) ToAuthorizerJwtConfigurationOutputWithContext(ctx context.Context) AuthorizerJwtConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerJwtConfigurationOutput)
}

func (i AuthorizerJwtConfigurationArgs) ToAuthorizerJwtConfigurationPtrOutput() AuthorizerJwtConfigurationPtrOutput {
	return i.ToAuthorizerJwtConfigurationPtrOutputWithContext(context.Background())
}

func (i AuthorizerJwtConfigurationArgs) ToAuthorizerJwtConfigurationPtrOutputWithContext(ctx context.Context) AuthorizerJwtConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerJwtConfigurationOutput).ToAuthorizerJwtConfigurationPtrOutputWithContext(ctx)
}

type AuthorizerJwtConfigurationPtrInput interface {
	pulumi.Input

	ToAuthorizerJwtConfigurationPtrOutput() AuthorizerJwtConfigurationPtrOutput
	ToAuthorizerJwtConfigurationPtrOutputWithContext(context.Context) AuthorizerJwtConfigurationPtrOutput
}

type authorizerJwtConfigurationPtrType AuthorizerJwtConfigurationArgs

func AuthorizerJwtConfigurationPtr(v *AuthorizerJwtConfigurationArgs) AuthorizerJwtConfigurationPtrInput {
	return (*authorizerJwtConfigurationPtrType)(v)
}

func (*authorizerJwtConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizerJwtConfiguration)(nil)).Elem()
}

func (i *authorizerJwtConfigurationPtrType) ToAuthorizerJwtConfigurationPtrOutput() AuthorizerJwtConfigurationPtrOutput {
	return i.ToAuthorizerJwtConfigurationPtrOutputWithContext(context.Background())
}

func (i *authorizerJwtConfigurationPtrType) ToAuthorizerJwtConfigurationPtrOutputWithContext(ctx context.Context) AuthorizerJwtConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerJwtConfigurationPtrOutput)
}

type AuthorizerJwtConfigurationOutput struct{ *pulumi.OutputState }

func (AuthorizerJwtConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizerJwtConfiguration)(nil)).Elem()
}

func (o AuthorizerJwtConfigurationOutput) ToAuthorizerJwtConfigurationOutput() AuthorizerJwtConfigurationOutput {
	return o
}

func (o AuthorizerJwtConfigurationOutput) ToAuthorizerJwtConfigurationOutputWithContext(ctx context.Context) AuthorizerJwtConfigurationOutput {
	return o
}

func (o AuthorizerJwtConfigurationOutput) ToAuthorizerJwtConfigurationPtrOutput() AuthorizerJwtConfigurationPtrOutput {
	return o.ToAuthorizerJwtConfigurationPtrOutputWithContext(context.Background())
}

func (o AuthorizerJwtConfigurationOutput) ToAuthorizerJwtConfigurationPtrOutputWithContext(ctx context.Context) AuthorizerJwtConfigurationPtrOutput {
	return o.ApplyT(func(v AuthorizerJwtConfiguration) *AuthorizerJwtConfiguration {
		return &v
	}).(AuthorizerJwtConfigurationPtrOutput)
}

// A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
func (o AuthorizerJwtConfigurationOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthorizerJwtConfiguration) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

// The base domain of the identity provider that issues JSON Web Tokens, such as the `endpoint` attribute of the [`cognito.UserPool`](https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html) resource.
func (o AuthorizerJwtConfigurationOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizerJwtConfiguration) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

type AuthorizerJwtConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AuthorizerJwtConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizerJwtConfiguration)(nil)).Elem()
}

func (o AuthorizerJwtConfigurationPtrOutput) ToAuthorizerJwtConfigurationPtrOutput() AuthorizerJwtConfigurationPtrOutput {
	return o
}

func (o AuthorizerJwtConfigurationPtrOutput) ToAuthorizerJwtConfigurationPtrOutputWithContext(ctx context.Context) AuthorizerJwtConfigurationPtrOutput {
	return o
}

func (o AuthorizerJwtConfigurationPtrOutput) Elem() AuthorizerJwtConfigurationOutput {
	return o.ApplyT(func(v *AuthorizerJwtConfiguration) AuthorizerJwtConfiguration { return *v }).(AuthorizerJwtConfigurationOutput)
}

// A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
func (o AuthorizerJwtConfigurationPtrOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthorizerJwtConfiguration) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

// The base domain of the identity provider that issues JSON Web Tokens, such as the `endpoint` attribute of the [`cognito.UserPool`](https://www.terraform.io/docs/providers/aws/r/cognito_user_pool.html) resource.
func (o AuthorizerJwtConfigurationPtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizerJwtConfiguration) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizerJwtConfigurationOutput{})
	pulumi.RegisterOutputType(AuthorizerJwtConfigurationPtrOutput{})
}
