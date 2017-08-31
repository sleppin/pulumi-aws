// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Recorder extends lumi.NamedResource implements RecorderArgs {
    public readonly name: string;
    public readonly recordingGroup: { allSupported?: boolean, includeGlobalResourceTypes?: boolean, resourceTypes?: string[] }[];
    public readonly roleArn: string;

    public static get(id: lumi.ID): Recorder {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Recorder[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(urnName: string, args: RecorderArgs) {
        super(urnName);
        this.name = <any>args.name;
        this.recordingGroup = <any>args.recordingGroup;
        if (lumirt.defaultIfComputed(args.roleArn, "") === undefined) {
            throw new Error("Property argument 'roleArn' is required, but was missing");
        }
        this.roleArn = <any>args.roleArn;
    }
}

export interface RecorderArgs {
    readonly name?: string;
    readonly recordingGroup?: { allSupported?: boolean, includeGlobalResourceTypes?: boolean, resourceTypes?: string[] }[];
    readonly roleArn: string;
}
