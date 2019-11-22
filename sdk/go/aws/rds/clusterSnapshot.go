// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a RDS database cluster snapshot for Aurora clusters. For managing RDS database instance snapshots, see the [`rds.Snapshot` resource](https://www.terraform.io/docs/providers/aws/r/db_snapshot.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_cluster_snapshot.html.markdown.
type ClusterSnapshot struct {
	s *pulumi.ResourceState
}

// NewClusterSnapshot registers a new resource with the given unique name, arguments, and options.
func NewClusterSnapshot(ctx *pulumi.Context,
	name string, args *ClusterSnapshotArgs, opts ...pulumi.ResourceOpt) (*ClusterSnapshot, error) {
	if args == nil || args.DbClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'DbClusterIdentifier'")
	}
	if args == nil || args.DbClusterSnapshotIdentifier == nil {
		return nil, errors.New("missing required argument 'DbClusterSnapshotIdentifier'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dbClusterIdentifier"] = nil
		inputs["dbClusterSnapshotIdentifier"] = nil
		inputs["tags"] = nil
	} else {
		inputs["dbClusterIdentifier"] = args.DbClusterIdentifier
		inputs["dbClusterSnapshotIdentifier"] = args.DbClusterSnapshotIdentifier
		inputs["tags"] = args.Tags
	}
	inputs["allocatedStorage"] = nil
	inputs["availabilityZones"] = nil
	inputs["dbClusterSnapshotArn"] = nil
	inputs["engine"] = nil
	inputs["engineVersion"] = nil
	inputs["kmsKeyId"] = nil
	inputs["licenseModel"] = nil
	inputs["port"] = nil
	inputs["snapshotType"] = nil
	inputs["sourceDbClusterSnapshotArn"] = nil
	inputs["status"] = nil
	inputs["storageEncrypted"] = nil
	inputs["vpcId"] = nil
	s, err := ctx.RegisterResource("aws:rds/clusterSnapshot:ClusterSnapshot", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterSnapshot{s: s}, nil
}

// GetClusterSnapshot gets an existing ClusterSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterSnapshot(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterSnapshotState, opts ...pulumi.ResourceOpt) (*ClusterSnapshot, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allocatedStorage"] = state.AllocatedStorage
		inputs["availabilityZones"] = state.AvailabilityZones
		inputs["dbClusterIdentifier"] = state.DbClusterIdentifier
		inputs["dbClusterSnapshotArn"] = state.DbClusterSnapshotArn
		inputs["dbClusterSnapshotIdentifier"] = state.DbClusterSnapshotIdentifier
		inputs["engine"] = state.Engine
		inputs["engineVersion"] = state.EngineVersion
		inputs["kmsKeyId"] = state.KmsKeyId
		inputs["licenseModel"] = state.LicenseModel
		inputs["port"] = state.Port
		inputs["snapshotType"] = state.SnapshotType
		inputs["sourceDbClusterSnapshotArn"] = state.SourceDbClusterSnapshotArn
		inputs["status"] = state.Status
		inputs["storageEncrypted"] = state.StorageEncrypted
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:rds/clusterSnapshot:ClusterSnapshot", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterSnapshot{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ClusterSnapshot) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ClusterSnapshot) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the allocated storage size in gigabytes (GB).
func (r *ClusterSnapshot) AllocatedStorage() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["allocatedStorage"])
}

// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
func (r *ClusterSnapshot) AvailabilityZones() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["availabilityZones"])
}

// The DB Cluster Identifier from which to take the snapshot.
func (r *ClusterSnapshot) DbClusterIdentifier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dbClusterIdentifier"])
}

// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
func (r *ClusterSnapshot) DbClusterSnapshotArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dbClusterSnapshotArn"])
}

// The Identifier for the snapshot.
func (r *ClusterSnapshot) DbClusterSnapshotIdentifier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dbClusterSnapshotIdentifier"])
}

// Specifies the name of the database engine.
func (r *ClusterSnapshot) Engine() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engine"])
}

// Version of the database engine for this DB cluster snapshot.
func (r *ClusterSnapshot) EngineVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineVersion"])
}

// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
func (r *ClusterSnapshot) KmsKeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kmsKeyId"])
}

// License model information for the restored DB cluster.
func (r *ClusterSnapshot) LicenseModel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["licenseModel"])
}

// Port that the DB cluster was listening on at the time of the snapshot.
func (r *ClusterSnapshot) Port() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["port"])
}

func (r *ClusterSnapshot) SnapshotType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["snapshotType"])
}

func (r *ClusterSnapshot) SourceDbClusterSnapshotArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceDbClusterSnapshotArn"])
}

// The status of this DB Cluster Snapshot.
func (r *ClusterSnapshot) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// Specifies whether the DB cluster snapshot is encrypted.
func (r *ClusterSnapshot) StorageEncrypted() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["storageEncrypted"])
}

// A mapping of tags to assign to the DB cluster.
func (r *ClusterSnapshot) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The VPC ID associated with the DB cluster snapshot.
func (r *ClusterSnapshot) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering ClusterSnapshot resources.
type ClusterSnapshotState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage interface{}
	// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
	AvailabilityZones interface{}
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier interface{}
	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	DbClusterSnapshotArn interface{}
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier interface{}
	// Specifies the name of the database engine.
	Engine interface{}
	// Version of the database engine for this DB cluster snapshot.
	EngineVersion interface{}
	// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
	KmsKeyId interface{}
	// License model information for the restored DB cluster.
	LicenseModel interface{}
	// Port that the DB cluster was listening on at the time of the snapshot.
	Port interface{}
	SnapshotType interface{}
	SourceDbClusterSnapshotArn interface{}
	// The status of this DB Cluster Snapshot.
	Status interface{}
	// Specifies whether the DB cluster snapshot is encrypted.
	StorageEncrypted interface{}
	// A mapping of tags to assign to the DB cluster.
	Tags interface{}
	// The VPC ID associated with the DB cluster snapshot.
	VpcId interface{}
}

// The set of arguments for constructing a ClusterSnapshot resource.
type ClusterSnapshotArgs struct {
	// The DB Cluster Identifier from which to take the snapshot.
	DbClusterIdentifier interface{}
	// The Identifier for the snapshot.
	DbClusterSnapshotIdentifier interface{}
	// A mapping of tags to assign to the DB cluster.
	Tags interface{}
}
