// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about RDS orderable DB instances and valid parameter combinations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "5.7.22"
// 		opt1 := "general-public-license"
// 		opt2 := "standard"
// 		_, err := rds.GetOrderableDbInstance(ctx, &rds.GetOrderableDbInstanceArgs{
// 			Engine:        "mysql",
// 			EngineVersion: &opt0,
// 			LicenseModel:  &opt1,
// 			PreferredInstanceClasses: []string{
// 				"db.r6.xlarge",
// 				"db.m4.large",
// 				"db.t3.small",
// 			},
// 			StorageType: &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Valid parameter combinations can also be found with `preferredEngineVersions` and/or `preferredInstanceClasses`.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "general-public-license"
// 		_, err := rds.GetOrderableDbInstance(ctx, &rds.GetOrderableDbInstanceArgs{
// 			Engine:       "mysql",
// 			LicenseModel: &opt0,
// 			PreferredEngineVersions: []string{
// 				"5.6.35",
// 				"5.6.41",
// 				"5.6.44",
// 			},
// 			PreferredInstanceClasses: []string{
// 				"db.t2.small",
// 				"db.t3.medium",
// 				"db.t3.large",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrderableDbInstance(ctx *pulumi.Context, args *GetOrderableDbInstanceArgs, opts ...pulumi.InvokeOption) (*GetOrderableDbInstanceResult, error) {
	var rv GetOrderableDbInstanceResult
	err := ctx.Invoke("aws:rds/getOrderableDbInstance:getOrderableDbInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderableDbInstance.
type GetOrderableDbInstanceArgs struct {
	// Availability zone group.
	AvailabilityZoneGroup *string `pulumi:"availabilityZoneGroup"`
	// DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
	Engine string `pulumi:"engine"`
	// Version of the DB engine. If none is provided, the AWS-defined default version will be used.
	EngineVersion *string `pulumi:"engineVersion"`
	// DB instance class. Examples of classes are `db.m3.2xlarge`, `db.t2.small`, and `db.m3.medium`.
	InstanceClass *string `pulumi:"instanceClass"`
	// License model. Examples of license models are `general-public-license`, `bring-your-own-license`, and `amazon-license`.
	LicenseModel *string `pulumi:"licenseModel"`
	// Ordered list of preferred RDS DB instance engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredEngineVersions []string `pulumi:"preferredEngineVersions"`
	// Ordered list of preferred RDS DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredInstanceClasses []string `pulumi:"preferredInstanceClasses"`
	// Storage types. Examples of storage types are `standard`, `io1`, `gp2`, and `aurora`.
	StorageType *string `pulumi:"storageType"`
	// Enable this to ensure a DB instance supports Enhanced Monitoring at intervals from 1 to 60 seconds.
	SupportsEnhancedMonitoring *bool `pulumi:"supportsEnhancedMonitoring"`
	// Enable this to ensure a DB instance supports Aurora global databases with a specific combination of other DB engine attributes.
	SupportsGlobalDatabases *bool `pulumi:"supportsGlobalDatabases"`
	// Enable this to ensure a DB instance supports IAM database authentication.
	SupportsIamDatabaseAuthentication *bool `pulumi:"supportsIamDatabaseAuthentication"`
	// Enable this to ensure a DB instance supports provisioned IOPS.
	SupportsIops *bool `pulumi:"supportsIops"`
	// Enable this to ensure a DB instance supports Kerberos Authentication.
	SupportsKerberosAuthentication *bool `pulumi:"supportsKerberosAuthentication"`
	// Enable this to ensure a DB instance supports Performance Insights.
	SupportsPerformanceInsights *bool `pulumi:"supportsPerformanceInsights"`
	// Enable this to ensure Amazon RDS can automatically scale storage for DB instances that use the specified DB instance class.
	SupportsStorageAutoscaling *bool `pulumi:"supportsStorageAutoscaling"`
	// Enable this to ensure a DB instance supports encrypted storage.
	SupportsStorageEncryption *bool `pulumi:"supportsStorageEncryption"`
	// Boolean that indicates whether to show only VPC or non-VPC offerings.
	Vpc *bool `pulumi:"vpc"`
}

// A collection of values returned by getOrderableDbInstance.
type GetOrderableDbInstanceResult struct {
	AvailabilityZoneGroup string `pulumi:"availabilityZoneGroup"`
	// Availability zones where the instance is available.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	Engine            string   `pulumi:"engine"`
	EngineVersion     string   `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	InstanceClass string `pulumi:"instanceClass"`
	LicenseModel  string `pulumi:"licenseModel"`
	// Maximum total provisioned IOPS for a DB instance.
	MaxIopsPerDbInstance int `pulumi:"maxIopsPerDbInstance"`
	// Maximum provisioned IOPS per GiB for a DB instance.
	MaxIopsPerGib float64 `pulumi:"maxIopsPerGib"`
	// Maximum storage size for a DB instance.
	MaxStorageSize int `pulumi:"maxStorageSize"`
	// Minimum total provisioned IOPS for a DB instance.
	MinIopsPerDbInstance int `pulumi:"minIopsPerDbInstance"`
	// Minimum provisioned IOPS per GiB for a DB instance.
	MinIopsPerGib float64 `pulumi:"minIopsPerGib"`
	// Minimum storage size for a DB instance.
	MinStorageSize int `pulumi:"minStorageSize"`
	// Whether a DB instance is Multi-AZ capable.
	MultiAzCapable bool `pulumi:"multiAzCapable"`
	// Whether a DB instance supports RDS on Outposts.
	OutpostCapable           bool     `pulumi:"outpostCapable"`
	PreferredEngineVersions  []string `pulumi:"preferredEngineVersions"`
	PreferredInstanceClasses []string `pulumi:"preferredInstanceClasses"`
	// Whether a DB instance can have a read replica.
	ReadReplicaCapable bool   `pulumi:"readReplicaCapable"`
	StorageType        string `pulumi:"storageType"`
	// A list of the supported DB engine modes.
	SupportedEngineModes              []string `pulumi:"supportedEngineModes"`
	SupportsEnhancedMonitoring        bool     `pulumi:"supportsEnhancedMonitoring"`
	SupportsGlobalDatabases           bool     `pulumi:"supportsGlobalDatabases"`
	SupportsIamDatabaseAuthentication bool     `pulumi:"supportsIamDatabaseAuthentication"`
	SupportsIops                      bool     `pulumi:"supportsIops"`
	SupportsKerberosAuthentication    bool     `pulumi:"supportsKerberosAuthentication"`
	SupportsPerformanceInsights       bool     `pulumi:"supportsPerformanceInsights"`
	SupportsStorageAutoscaling        bool     `pulumi:"supportsStorageAutoscaling"`
	SupportsStorageEncryption         bool     `pulumi:"supportsStorageEncryption"`
	Vpc                               bool     `pulumi:"vpc"`
}
