// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a directory registration in AWS WorkSpaces Service
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const mainVpc = new aws.ec2.Vpc("main", {
 *     cidrBlock: "10.0.0.0/16",
 * });
 * const privateA = new aws.ec2.Subnet("private-a", {
 *     availabilityZone: "us-east-1a",
 *     cidrBlock: "10.0.0.0/24",
 *     vpcId: mainVpc.id,
 * });
 * const privateB = new aws.ec2.Subnet("private-b", {
 *     availabilityZone: "us-east-1b",
 *     cidrBlock: "10.0.1.0/24",
 *     vpcId: mainVpc.id,
 * });
 * const mainDirectory = new aws.directoryservice.Directory("main", {
 *     password: "#S1ncerely",
 *     size: "Small",
 *     vpcSettings: {
 *         subnetIds: [
 *             private_a.id,
 *             private_b.id,
 *         ],
 *         vpcId: mainVpc.id,
 *     },
 * });
 * const mainWorkspacesDirectory = new aws.workspaces.Directory("main", {
 *     directoryId: mainDirectory.id,
 *     selfServicePermissions: {
 *         increaseVolumeSize: true,
 *         rebuildWorkspace: true,
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/workspaces_directory.html.markdown.
 */
export class Directory extends pulumi.CustomResource {
    /**
     * Get an existing Directory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DirectoryState, opts?: pulumi.CustomResourceOptions): Directory {
        return new Directory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:workspaces/directory:Directory';

    /**
     * Returns true if the given object is an instance of Directory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Directory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Directory.__pulumiType;
    }

    /**
     * The directory identifier for registration in WorkSpaces service.
     */
    public readonly directoryId!: pulumi.Output<string>;
    /**
     * The permissions to enable or disable self-service capabilities.
     */
    public readonly selfServicePermissions!: pulumi.Output<outputs.workspaces.DirectorySelfServicePermissions>;
    /**
     * The identifiers of the subnets where the directory resides.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A mapping of tags assigned to the WorkSpaces directory.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Directory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DirectoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DirectoryArgs | DirectoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DirectoryState | undefined;
            inputs["directoryId"] = state ? state.directoryId : undefined;
            inputs["selfServicePermissions"] = state ? state.selfServicePermissions : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DirectoryArgs | undefined;
            if (!args || args.directoryId === undefined) {
                throw new Error("Missing required property 'directoryId'");
            }
            inputs["directoryId"] = args ? args.directoryId : undefined;
            inputs["selfServicePermissions"] = args ? args.selfServicePermissions : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Directory.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Directory resources.
 */
export interface DirectoryState {
    /**
     * The directory identifier for registration in WorkSpaces service.
     */
    readonly directoryId?: pulumi.Input<string>;
    /**
     * The permissions to enable or disable self-service capabilities.
     */
    readonly selfServicePermissions?: pulumi.Input<inputs.workspaces.DirectorySelfServicePermissions>;
    /**
     * The identifiers of the subnets where the directory resides.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags assigned to the WorkSpaces directory.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Directory resource.
 */
export interface DirectoryArgs {
    /**
     * The directory identifier for registration in WorkSpaces service.
     */
    readonly directoryId: pulumi.Input<string>;
    /**
     * The permissions to enable or disable self-service capabilities.
     */
    readonly selfServicePermissions?: pulumi.Input<inputs.workspaces.DirectorySelfServicePermissions>;
    /**
     * The identifiers of the subnets where the directory resides.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags assigned to the WorkSpaces directory.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
