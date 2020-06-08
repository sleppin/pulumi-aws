// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 integration.
 * More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.Integration("example", {
 *     apiId: aws_apigatewayv2_api_example.id,
 *     integrationType: "MOCK",
 * });
 * ```
 *
 * ### Lambda Integration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleFunction = new aws.lambda.Function("example", {
 *     code: new pulumi.asset.FileArchive("example.zip"),
 *     handler: "index.handler",
 *     role: aws_iam_role_example.arn,
 *     runtime: "nodejs10.x",
 * });
 * const exampleIntegration = new aws.apigatewayv2.Integration("example", {
 *     apiId: aws_apigatewayv2_api_example.id,
 *     connectionType: "INTERNET",
 *     contentHandlingStrategy: "CONVERT_TO_TEXT",
 *     description: "Lambda example",
 *     integrationMethod: "POST",
 *     integrationType: "AWS",
 *     integrationUri: exampleFunction.invokeArn,
 *     passthroughBehavior: "WHEN_NO_MATCH",
 * });
 * ```
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationState, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/integration:Integration';

    /**
     * Returns true if the given object is an instance of Integration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Integration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Integration.__pulumiType;
    }

    /**
     * The API identifier.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The ID of the VPC link for a private integration. Supported only for HTTP APIs.
     */
    public readonly connectionId!: pulumi.Output<string | undefined>;
    /**
     * The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
     */
    public readonly contentHandlingStrategy!: pulumi.Output<string | undefined>;
    /**
     * The credentials required for the integration, if any.
     */
    public readonly credentialsArn!: pulumi.Output<string | undefined>;
    /**
     * The description of the integration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
     */
    public readonly integrationMethod!: pulumi.Output<string | undefined>;
    /**
     * The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
     */
    public /*out*/ readonly integrationResponseSelectionExpression!: pulumi.Output<string>;
    /**
     * The integration type of an integration.
     * Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
     */
    public readonly integrationType!: pulumi.Output<string>;
    /**
     * The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
     * For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
     */
    public readonly integrationUri!: pulumi.Output<string | undefined>;
    /**
     * The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
     * Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
     */
    public readonly passthroughBehavior!: pulumi.Output<string | undefined>;
    /**
     * The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
     */
    public readonly payloadFormatVersion!: pulumi.Output<string | undefined>;
    /**
     * A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
     */
    public readonly requestTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
     */
    public readonly templateSelectionExpression!: pulumi.Output<string | undefined>;
    /**
     * Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
     */
    public readonly timeoutMilliseconds!: pulumi.Output<number | undefined>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["connectionType"] = state ? state.connectionType : undefined;
            inputs["contentHandlingStrategy"] = state ? state.contentHandlingStrategy : undefined;
            inputs["credentialsArn"] = state ? state.credentialsArn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["integrationMethod"] = state ? state.integrationMethod : undefined;
            inputs["integrationResponseSelectionExpression"] = state ? state.integrationResponseSelectionExpression : undefined;
            inputs["integrationType"] = state ? state.integrationType : undefined;
            inputs["integrationUri"] = state ? state.integrationUri : undefined;
            inputs["passthroughBehavior"] = state ? state.passthroughBehavior : undefined;
            inputs["payloadFormatVersion"] = state ? state.payloadFormatVersion : undefined;
            inputs["requestTemplates"] = state ? state.requestTemplates : undefined;
            inputs["templateSelectionExpression"] = state ? state.templateSelectionExpression : undefined;
            inputs["timeoutMilliseconds"] = state ? state.timeoutMilliseconds : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.integrationType === undefined) {
                throw new Error("Missing required property 'integrationType'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["connectionId"] = args ? args.connectionId : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["contentHandlingStrategy"] = args ? args.contentHandlingStrategy : undefined;
            inputs["credentialsArn"] = args ? args.credentialsArn : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["integrationMethod"] = args ? args.integrationMethod : undefined;
            inputs["integrationType"] = args ? args.integrationType : undefined;
            inputs["integrationUri"] = args ? args.integrationUri : undefined;
            inputs["passthroughBehavior"] = args ? args.passthroughBehavior : undefined;
            inputs["payloadFormatVersion"] = args ? args.payloadFormatVersion : undefined;
            inputs["requestTemplates"] = args ? args.requestTemplates : undefined;
            inputs["templateSelectionExpression"] = args ? args.templateSelectionExpression : undefined;
            inputs["timeoutMilliseconds"] = args ? args.timeoutMilliseconds : undefined;
            inputs["integrationResponseSelectionExpression"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Integration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * The API identifier.
     */
    readonly apiId?: pulumi.Input<string>;
    /**
     * The ID of the VPC link for a private integration. Supported only for HTTP APIs.
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
     */
    readonly connectionType?: pulumi.Input<string>;
    /**
     * How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
     */
    readonly contentHandlingStrategy?: pulumi.Input<string>;
    /**
     * The credentials required for the integration, if any.
     */
    readonly credentialsArn?: pulumi.Input<string>;
    /**
     * The description of the integration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
     */
    readonly integrationMethod?: pulumi.Input<string>;
    /**
     * The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
     */
    readonly integrationResponseSelectionExpression?: pulumi.Input<string>;
    /**
     * The integration type of an integration.
     * Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
     */
    readonly integrationType?: pulumi.Input<string>;
    /**
     * The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
     * For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
     */
    readonly integrationUri?: pulumi.Input<string>;
    /**
     * The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
     * Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
     */
    readonly passthroughBehavior?: pulumi.Input<string>;
    /**
     * The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
     */
    readonly payloadFormatVersion?: pulumi.Input<string>;
    /**
     * A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
     */
    readonly requestTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
     */
    readonly templateSelectionExpression?: pulumi.Input<string>;
    /**
     * Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
     */
    readonly timeoutMilliseconds?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * The API identifier.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * The ID of the VPC link for a private integration. Supported only for HTTP APIs.
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
     */
    readonly connectionType?: pulumi.Input<string>;
    /**
     * How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
     */
    readonly contentHandlingStrategy?: pulumi.Input<string>;
    /**
     * The credentials required for the integration, if any.
     */
    readonly credentialsArn?: pulumi.Input<string>;
    /**
     * The description of the integration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
     */
    readonly integrationMethod?: pulumi.Input<string>;
    /**
     * The integration type of an integration.
     * Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
     */
    readonly integrationType: pulumi.Input<string>;
    /**
     * The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
     * For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
     */
    readonly integrationUri?: pulumi.Input<string>;
    /**
     * The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
     * Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
     */
    readonly passthroughBehavior?: pulumi.Input<string>;
    /**
     * The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
     */
    readonly payloadFormatVersion?: pulumi.Input<string>;
    /**
     * A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
     */
    readonly requestTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
     */
    readonly templateSelectionExpression?: pulumi.Input<string>;
    /**
     * Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
     */
    readonly timeoutMilliseconds?: pulumi.Input<number>;
}
