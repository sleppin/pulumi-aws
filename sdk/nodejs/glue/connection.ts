// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Glue Connection resource.
 *
 * ## Example Usage
 *
 * ### Non-VPC Connection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Connection("example", {
 *     connectionProperties: {
 *         JDBC_CONNECTION_URL: "jdbc:mysql://example.com/exampledatabase",
 *         PASSWORD: "examplepassword",
 *         USERNAME: "exampleusername",
 *     },
 * });
 * ```
 *
 * ### VPC Connection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Connection("example", {
 *     connectionProperties: {
 *         JDBC_CONNECTION_URL: pulumi.interpolate`jdbc:mysql://${aws_rds_cluster_example.endpoint}/exampledatabase`,
 *         PASSWORD: "examplepassword",
 *         USERNAME: "exampleusername",
 *     },
 *     physicalConnectionRequirements: {
 *         availabilityZone: aws_subnet_example.availabilityZone,
 *         securityGroupIdLists: [aws_security_group_example.id],
 *         subnetId: aws_subnet_example.id,
 *     },
 * });
 * ```
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The ARN of the Glue Connection.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * A map of key-value pairs used as parameters for this connection.
     */
    public readonly connectionProperties!: pulumi.Output<{[key: string]: any}>;
    /**
     * The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`. Defaults to `JBDC`.
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * Description of the connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of criteria that can be used in selecting this connection.
     */
    public readonly matchCriterias!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the connection.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
     */
    public readonly physicalConnectionRequirements!: pulumi.Output<outputs.glue.ConnectionPhysicalConnectionRequirements | undefined>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["connectionProperties"] = state ? state.connectionProperties : undefined;
            inputs["connectionType"] = state ? state.connectionType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["matchCriterias"] = state ? state.matchCriterias : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["physicalConnectionRequirements"] = state ? state.physicalConnectionRequirements : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if (!args || args.connectionProperties === undefined) {
                throw new Error("Missing required property 'connectionProperties'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["connectionProperties"] = args ? args.connectionProperties : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["matchCriterias"] = args ? args.matchCriterias : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["physicalConnectionRequirements"] = args ? args.physicalConnectionRequirements : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Connection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    /**
     * The ARN of the Glue Connection.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
     */
    readonly catalogId?: pulumi.Input<string>;
    /**
     * A map of key-value pairs used as parameters for this connection.
     */
    readonly connectionProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`. Defaults to `JBDC`.
     */
    readonly connectionType?: pulumi.Input<string>;
    /**
     * Description of the connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of criteria that can be used in selecting this connection.
     */
    readonly matchCriterias?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the connection.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
     */
    readonly physicalConnectionRequirements?: pulumi.Input<inputs.glue.ConnectionPhysicalConnectionRequirements>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
     */
    readonly catalogId?: pulumi.Input<string>;
    /**
     * A map of key-value pairs used as parameters for this connection.
     */
    readonly connectionProperties: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`. Defaults to `JBDC`.
     */
    readonly connectionType?: pulumi.Input<string>;
    /**
     * Description of the connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of criteria that can be used in selecting this connection.
     */
    readonly matchCriterias?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the connection.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
     */
    readonly physicalConnectionRequirements?: pulumi.Input<inputs.glue.ConnectionPhysicalConnectionRequirements>;
}
