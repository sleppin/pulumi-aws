// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 integration response.
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
 * const example = new aws.apigatewayv2.IntegrationResponse("example", {
 *     apiId: aws_apigatewayv2_api_example.id,
 *     integrationId: aws_apigatewayv2_integration_example.id,
 *     integrationResponseKey: "/200/",
 * });
 * ```
 */
export class IntegrationResponse extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationResponseState, opts?: pulumi.CustomResourceOptions): IntegrationResponse {
        return new IntegrationResponse(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/integrationResponse:IntegrationResponse';

    /**
     * Returns true if the given object is an instance of IntegrationResponse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationResponse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationResponse.__pulumiType;
    }

    /**
     * The API identifier.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
     */
    public readonly contentHandlingStrategy!: pulumi.Output<string | undefined>;
    /**
     * The identifier of the `aws.apigatewayv2.Integration`.
     */
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * The integration response key.
     */
    public readonly integrationResponseKey!: pulumi.Output<string>;
    /**
     * A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
     */
    public readonly responseTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
     */
    public readonly templateSelectionExpression!: pulumi.Output<string | undefined>;

    /**
     * Create a IntegrationResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationResponseArgs | IntegrationResponseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationResponseState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["contentHandlingStrategy"] = state ? state.contentHandlingStrategy : undefined;
            inputs["integrationId"] = state ? state.integrationId : undefined;
            inputs["integrationResponseKey"] = state ? state.integrationResponseKey : undefined;
            inputs["responseTemplates"] = state ? state.responseTemplates : undefined;
            inputs["templateSelectionExpression"] = state ? state.templateSelectionExpression : undefined;
        } else {
            const args = argsOrState as IntegrationResponseArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.integrationId === undefined) {
                throw new Error("Missing required property 'integrationId'");
            }
            if (!args || args.integrationResponseKey === undefined) {
                throw new Error("Missing required property 'integrationResponseKey'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["contentHandlingStrategy"] = args ? args.contentHandlingStrategy : undefined;
            inputs["integrationId"] = args ? args.integrationId : undefined;
            inputs["integrationResponseKey"] = args ? args.integrationResponseKey : undefined;
            inputs["responseTemplates"] = args ? args.responseTemplates : undefined;
            inputs["templateSelectionExpression"] = args ? args.templateSelectionExpression : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IntegrationResponse.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationResponse resources.
 */
export interface IntegrationResponseState {
    /**
     * The API identifier.
     */
    readonly apiId?: pulumi.Input<string>;
    /**
     * How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
     */
    readonly contentHandlingStrategy?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.apigatewayv2.Integration`.
     */
    readonly integrationId?: pulumi.Input<string>;
    /**
     * The integration response key.
     */
    readonly integrationResponseKey?: pulumi.Input<string>;
    /**
     * A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
     */
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
     */
    readonly templateSelectionExpression?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationResponse resource.
 */
export interface IntegrationResponseArgs {
    /**
     * The API identifier.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
     */
    readonly contentHandlingStrategy?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.apigatewayv2.Integration`.
     */
    readonly integrationId: pulumi.Input<string>;
    /**
     * The integration response key.
     */
    readonly integrationResponseKey: pulumi.Input<string>;
    /**
     * A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
     */
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
     */
    readonly templateSelectionExpression?: pulumi.Input<string>;
}
