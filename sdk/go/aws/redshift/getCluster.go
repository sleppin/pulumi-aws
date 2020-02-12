// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package redshift

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides details about a specific redshift cluster.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/redshift_cluster.html.markdown.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("aws:redshift/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The cluster identifier
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// Whether major version upgrades can be applied during maintenance period
	AllowVersionUpgrade bool `pulumi:"allowVersionUpgrade"`
	// The backup retention period
	AutomatedSnapshotRetentionPeriod int `pulumi:"automatedSnapshotRetentionPeriod"`
	// The availability zone of the cluster
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The name of the S3 bucket where the log files are to be stored
	BucketName string `pulumi:"bucketName"`
	// The cluster identifier
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// The name of the parameter group to be associated with this cluster
	ClusterParameterGroupName string `pulumi:"clusterParameterGroupName"`
	// The public key for the cluster
	ClusterPublicKey string `pulumi:"clusterPublicKey"`
	// The cluster revision number
	ClusterRevisionNumber string `pulumi:"clusterRevisionNumber"`
	// The security groups associated with the cluster
	ClusterSecurityGroups []string `pulumi:"clusterSecurityGroups"`
	// The name of a cluster subnet group to be associated with this cluster
	ClusterSubnetGroupName string `pulumi:"clusterSubnetGroupName"`
	// The cluster type
	ClusterType string `pulumi:"clusterType"`
	ClusterVersion string `pulumi:"clusterVersion"`
	// The name of the default database in the cluster
	DatabaseName string `pulumi:"databaseName"`
	// The Elastic IP of the cluster
	ElasticIp string `pulumi:"elasticIp"`
	// Whether cluster logging is enabled
	EnableLogging bool `pulumi:"enableLogging"`
	// Whether the cluster data is encrypted
	Encrypted bool `pulumi:"encrypted"`
	// The cluster endpoint
	Endpoint string `pulumi:"endpoint"`
	// Whether enhanced VPC routing is enabled
	EnhancedVpcRouting bool `pulumi:"enhancedVpcRouting"`
	// The IAM roles associated to the cluster
	IamRoles []string `pulumi:"iamRoles"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The KMS encryption key associated to the cluster
	KmsKeyId string `pulumi:"kmsKeyId"`
	// Username for the master DB user
	MasterUsername string `pulumi:"masterUsername"`
	// The cluster node type
	NodeType string `pulumi:"nodeType"`
	// The number of nodes in the cluster
	NumberOfNodes int `pulumi:"numberOfNodes"`
	// The port the cluster responds on
	Port int `pulumi:"port"`
	// The maintenance window
	PreferredMaintenanceWindow string `pulumi:"preferredMaintenanceWindow"`
	// Whether the cluster is publicly accessible
	PubliclyAccessible bool `pulumi:"publiclyAccessible"`
	// The folder inside the S3 bucket where the log files are stored
	S3KeyPrefix string `pulumi:"s3KeyPrefix"`
	// The tags associated to the cluster
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC Id associated with the cluster
	VpcId string `pulumi:"vpcId"`
	// The VPC security group Ids associated with the cluster
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

