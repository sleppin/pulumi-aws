// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway Deployment.
//
// > **Note:** Depends on having `apigateway.Integration` inside your rest api (which in turn depends on `apigateway.Method`). To avoid race conditions
// you might need to add an explicit `dependsOn = ["aws_api_gateway_integration.name"]`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_deployment.html.markdown.
type Deployment struct {
	pulumi.CustomResourceState

	// The creation date of the deployment
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the deployment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The execution ARN to be used in [`lambdaPermission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringOutput `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The description of the stage
	StageDescription pulumi.StringPtrOutput `pulumi:"stageDescription"`
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName pulumi.StringPtrOutput `pulumi:"stageName"`
	// A map that defines variables for the stage
	Variables pulumi.StringMapOutput `pulumi:"variables"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil {
		args = &DeploymentArgs{}
	}
	var resource Deployment
	err := ctx.RegisterResource("aws:apigateway/deployment:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("aws:apigateway/deployment:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
	// The creation date of the deployment
	CreatedDate *string `pulumi:"createdDate"`
	// The description of the deployment
	Description *string `pulumi:"description"`
	// The execution ARN to be used in [`lambdaPermission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn *string `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl *string `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The description of the stage
	StageDescription *string `pulumi:"stageDescription"`
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName *string `pulumi:"stageName"`
	// A map that defines variables for the stage
	Variables map[string]string `pulumi:"variables"`
}

type DeploymentState struct {
	// The creation date of the deployment
	CreatedDate pulumi.StringPtrInput
	// The description of the deployment
	Description pulumi.StringPtrInput
	// The execution ARN to be used in [`lambdaPermission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringPtrInput
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The description of the stage
	StageDescription pulumi.StringPtrInput
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName pulumi.StringPtrInput
	// A map that defines variables for the stage
	Variables pulumi.StringMapInput
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The description of the deployment
	Description *string `pulumi:"description"`
	// The ID of the associated REST API
	RestApi string `pulumi:"restApi"`
	// The description of the stage
	StageDescription *string `pulumi:"stageDescription"`
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName *string `pulumi:"stageName"`
	// A map that defines variables for the stage
	Variables map[string]string `pulumi:"variables"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The description of the deployment
	Description pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringInput
	// The description of the stage
	StageDescription pulumi.StringPtrInput
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName pulumi.StringPtrInput
	// A map that defines variables for the stage
	Variables pulumi.StringMapInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

