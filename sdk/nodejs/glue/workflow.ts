// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Glue Workflow resource.
 * The workflow graph (DAG) can be build using the `aws.glue.Trigger` resource.
 * See the example below for creating a graph with four nodes (two triggers and two jobs).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Workflow("example", {});
 * const example_start = new aws.glue.Trigger("example-start", {
 *     type: "ON_DEMAND",
 *     workflowName: example.name,
 *     actions: [{
 *         jobName: "example-job",
 *     }],
 * });
 * const example_inner = new aws.glue.Trigger("example-inner", {
 *     type: "CONDITIONAL",
 *     workflowName: example.name,
 *     predicate: {
 *         conditions: [{
 *             jobName: "example-job",
 *             state: "SUCCEEDED",
 *         }],
 *     },
 *     actions: [{
 *         jobName: "another-example-job",
 *     }],
 * });
 * ```
 */
export class Workflow extends pulumi.CustomResource {
    /**
     * Get an existing Workflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowState, opts?: pulumi.CustomResourceOptions): Workflow {
        return new Workflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/workflow:Workflow';

    /**
     * Returns true if the given object is an instance of Workflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workflow.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of Glue Workflow
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
     */
    public readonly defaultRunProperties!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Description of the workflow.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
     */
    public readonly maxConcurrentRuns!: pulumi.Output<number | undefined>;
    /**
     * The name you assign to this workflow.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Workflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowArgs | WorkflowState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WorkflowState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["defaultRunProperties"] = state ? state.defaultRunProperties : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maxConcurrentRuns"] = state ? state.maxConcurrentRuns : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as WorkflowArgs | undefined;
            inputs["defaultRunProperties"] = args ? args.defaultRunProperties : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maxConcurrentRuns"] = args ? args.maxConcurrentRuns : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Workflow.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workflow resources.
 */
export interface WorkflowState {
    /**
     * Amazon Resource Name (ARN) of Glue Workflow
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
     */
    readonly defaultRunProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Description of the workflow.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
     */
    readonly maxConcurrentRuns?: pulumi.Input<number>;
    /**
     * The name you assign to this workflow.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Workflow resource.
 */
export interface WorkflowArgs {
    /**
     * A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
     */
    readonly defaultRunProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Description of the workflow.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
     */
    readonly maxConcurrentRuns?: pulumi.Input<number>;
    /**
     * The name you assign to this workflow.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
