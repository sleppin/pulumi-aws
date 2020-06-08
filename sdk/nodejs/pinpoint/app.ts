// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Pinpoint App resource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.pinpoint.App("example", {
 *     limits: {
 *         maximumDuration: 600,
 *     },
 *     quietTime: {
 *         end: "06:00",
 *         start: "00:00",
 *     },
 * });
 * ```
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppState, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:pinpoint/app:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * The Application ID of the Pinpoint App.
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the PinPoint Application
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    public readonly campaignHook!: pulumi.Output<outputs.pinpoint.AppCampaignHook | undefined>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    public readonly limits!: pulumi.Output<outputs.pinpoint.AppLimits | undefined>;
    /**
     * The application name. By default generated by this provider
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Pinpoint application. Conflicts with `name`
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
     */
    public readonly quietTime!: pulumi.Output<outputs.pinpoint.AppQuietTime | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppArgs | AppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AppState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["campaignHook"] = state ? state.campaignHook : undefined;
            inputs["limits"] = state ? state.limits : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["quietTime"] = state ? state.quietTime : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AppArgs | undefined;
            inputs["campaignHook"] = args ? args.campaignHook : undefined;
            inputs["limits"] = args ? args.limits : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["quietTime"] = args ? args.quietTime : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["applicationId"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(App.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering App resources.
 */
export interface AppState {
    /**
     * The Application ID of the Pinpoint App.
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the PinPoint Application
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    readonly campaignHook?: pulumi.Input<inputs.pinpoint.AppCampaignHook>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    readonly limits?: pulumi.Input<inputs.pinpoint.AppLimits>;
    /**
     * The application name. By default generated by this provider
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Pinpoint application. Conflicts with `name`
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
     */
    readonly quietTime?: pulumi.Input<inputs.pinpoint.AppQuietTime>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    readonly campaignHook?: pulumi.Input<inputs.pinpoint.AppCampaignHook>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    readonly limits?: pulumi.Input<inputs.pinpoint.AppLimits>;
    /**
     * The application name. By default generated by this provider
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Pinpoint application. Conflicts with `name`
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
     */
    readonly quietTime?: pulumi.Input<inputs.pinpoint.AppQuietTime>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
