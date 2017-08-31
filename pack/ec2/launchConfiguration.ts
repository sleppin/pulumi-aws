// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class LaunchConfiguration extends lumi.NamedResource implements LaunchConfigurationArgs {
    public readonly associatePublicIpAddress?: boolean;
    public readonly ebsBlockDevice: { deleteOnTermination?: boolean, deviceName: string, encrypted: boolean, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[];
    public readonly ebsOptimized: boolean;
    public readonly enableMonitoring?: boolean;
    public readonly ephemeralBlockDevice?: { deviceName: string, virtualName: string }[];
    public readonly iamInstanceProfile?: string;
    public readonly imageId: string;
    public readonly instanceType: string;
    public readonly keyName: string;
    public readonly name: string;
    public readonly namePrefix?: string;
    public readonly placementTenancy?: string;
    public readonly rootBlockDevice: { deleteOnTermination?: boolean, iops: number, volumeSize: number, volumeType: string }[];
    public readonly securityGroups?: string[];
    public readonly spotPrice?: string;
    public readonly userData?: string;
    public readonly vpcClassicLinkId?: string;
    public readonly vpcClassicLinkSecurityGroups?: string[];

    public static get(id: lumi.ID): LaunchConfiguration {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): LaunchConfiguration[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(urnName: string, args: LaunchConfigurationArgs) {
        super(urnName);
        this.associatePublicIpAddress = <any>args.associatePublicIpAddress;
        this.ebsBlockDevice = <any>args.ebsBlockDevice;
        this.ebsOptimized = <any>args.ebsOptimized;
        this.enableMonitoring = <any>args.enableMonitoring;
        this.ephemeralBlockDevice = <any>args.ephemeralBlockDevice;
        this.iamInstanceProfile = <any>args.iamInstanceProfile;
        if (lumirt.defaultIfComputed(args.imageId, "") === undefined) {
            throw new Error("Property argument 'imageId' is required, but was missing");
        }
        this.imageId = <any>args.imageId;
        if (lumirt.defaultIfComputed(args.instanceType, "") === undefined) {
            throw new Error("Property argument 'instanceType' is required, but was missing");
        }
        this.instanceType = <any>args.instanceType;
        this.keyName = <any>args.keyName;
        this.name = <any>args.name;
        this.namePrefix = <any>args.namePrefix;
        this.placementTenancy = <any>args.placementTenancy;
        this.rootBlockDevice = <any>args.rootBlockDevice;
        this.securityGroups = <any>args.securityGroups;
        this.spotPrice = <any>args.spotPrice;
        this.userData = <any>args.userData;
        this.vpcClassicLinkId = <any>args.vpcClassicLinkId;
        this.vpcClassicLinkSecurityGroups = <any>args.vpcClassicLinkSecurityGroups;
    }
}

export interface LaunchConfigurationArgs {
    readonly associatePublicIpAddress?: boolean;
    readonly ebsBlockDevice?: { deleteOnTermination?: boolean, deviceName: string, encrypted?: boolean, iops?: number, snapshotId?: string, volumeSize?: number, volumeType?: string }[];
    readonly ebsOptimized?: boolean;
    readonly enableMonitoring?: boolean;
    readonly ephemeralBlockDevice?: { deviceName: string, virtualName: string }[];
    readonly iamInstanceProfile?: string;
    readonly imageId: string;
    readonly instanceType: string;
    readonly keyName?: string;
    readonly name?: string;
    readonly namePrefix?: string;
    readonly placementTenancy?: string;
    readonly rootBlockDevice?: { deleteOnTermination?: boolean, iops?: number, volumeSize?: number, volumeType?: string }[];
    readonly securityGroups?: string[];
    readonly spotPrice?: string;
    readonly userData?: string;
    readonly vpcClassicLinkId?: string;
    readonly vpcClassicLinkSecurityGroups?: string[];
}
