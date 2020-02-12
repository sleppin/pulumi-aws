// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package directoryservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Log subscription for AWS Directory Service that pushes logs to cloudwatch.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/directory_service_log_subscription.html.markdown.
type LogService struct {
	pulumi.CustomResourceState

	// The id of directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
}

// NewLogService registers a new resource with the given unique name, arguments, and options.
func NewLogService(ctx *pulumi.Context,
	name string, args *LogServiceArgs, opts ...pulumi.ResourceOption) (*LogService, error) {
	if args == nil || args.DirectoryId == nil {
		return nil, errors.New("missing required argument 'DirectoryId'")
	}
	if args == nil || args.LogGroupName == nil {
		return nil, errors.New("missing required argument 'LogGroupName'")
	}
	if args == nil {
		args = &LogServiceArgs{}
	}
	var resource LogService
	err := ctx.RegisterResource("aws:directoryservice/logService:LogService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogService gets an existing LogService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogServiceState, opts ...pulumi.ResourceOption) (*LogService, error) {
	var resource LogService
	err := ctx.ReadResource("aws:directoryservice/logService:LogService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogService resources.
type logServiceState struct {
	// The id of directory.
	DirectoryId *string `pulumi:"directoryId"`
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName *string `pulumi:"logGroupName"`
}

type LogServiceState struct {
	// The id of directory.
	DirectoryId pulumi.StringPtrInput
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName pulumi.StringPtrInput
}

func (LogServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*logServiceState)(nil)).Elem()
}

type logServiceArgs struct {
	// The id of directory.
	DirectoryId string `pulumi:"directoryId"`
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName string `pulumi:"logGroupName"`
}

// The set of arguments for constructing a LogService resource.
type LogServiceArgs struct {
	// The id of directory.
	DirectoryId pulumi.StringInput
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName pulumi.StringInput
}

func (LogServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logServiceArgs)(nil)).Elem()
}

