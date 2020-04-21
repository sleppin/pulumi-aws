// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PlanRule struct {
	// The amount of time AWS Backup attempts a backup before canceling the job and returning an error.
	CompletionWindow *int `pulumi:"completionWindow"`
	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle *PlanRuleLifecycle `pulumi:"lifecycle"`
	// Metadata that you can assign to help organize the resources that you create.
	RecoveryPointTags map[string]interface{} `pulumi:"recoveryPointTags"`
	// An display name for a backup rule.
	RuleName string `pulumi:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	Schedule *string `pulumi:"schedule"`
	// The amount of time in minutes before beginning a backup.
	StartWindow *int `pulumi:"startWindow"`
	// The name of a logical container where backups are stored.
	TargetVaultName string `pulumi:"targetVaultName"`
}

// PlanRuleInput is an input type that accepts PlanRuleArgs and PlanRuleOutput values.
// You can construct a concrete instance of `PlanRuleInput` via:
//
// 		 PlanRuleArgs{...}
//
type PlanRuleInput interface {
	pulumi.Input

	ToPlanRuleOutput() PlanRuleOutput
	ToPlanRuleOutputWithContext(context.Context) PlanRuleOutput
}

type PlanRuleArgs struct {
	// The amount of time AWS Backup attempts a backup before canceling the job and returning an error.
	CompletionWindow pulumi.IntPtrInput `pulumi:"completionWindow"`
	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle PlanRuleLifecyclePtrInput `pulumi:"lifecycle"`
	// Metadata that you can assign to help organize the resources that you create.
	RecoveryPointTags pulumi.MapInput `pulumi:"recoveryPointTags"`
	// An display name for a backup rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	Schedule pulumi.StringPtrInput `pulumi:"schedule"`
	// The amount of time in minutes before beginning a backup.
	StartWindow pulumi.IntPtrInput `pulumi:"startWindow"`
	// The name of a logical container where backups are stored.
	TargetVaultName pulumi.StringInput `pulumi:"targetVaultName"`
}

func (PlanRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanRule)(nil)).Elem()
}

func (i PlanRuleArgs) ToPlanRuleOutput() PlanRuleOutput {
	return i.ToPlanRuleOutputWithContext(context.Background())
}

func (i PlanRuleArgs) ToPlanRuleOutputWithContext(ctx context.Context) PlanRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanRuleOutput)
}

// PlanRuleArrayInput is an input type that accepts PlanRuleArray and PlanRuleArrayOutput values.
// You can construct a concrete instance of `PlanRuleArrayInput` via:
//
// 		 PlanRuleArray{ PlanRuleArgs{...} }
//
type PlanRuleArrayInput interface {
	pulumi.Input

	ToPlanRuleArrayOutput() PlanRuleArrayOutput
	ToPlanRuleArrayOutputWithContext(context.Context) PlanRuleArrayOutput
}

type PlanRuleArray []PlanRuleInput

func (PlanRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanRule)(nil)).Elem()
}

func (i PlanRuleArray) ToPlanRuleArrayOutput() PlanRuleArrayOutput {
	return i.ToPlanRuleArrayOutputWithContext(context.Background())
}

func (i PlanRuleArray) ToPlanRuleArrayOutputWithContext(ctx context.Context) PlanRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanRuleArrayOutput)
}

type PlanRuleOutput struct{ *pulumi.OutputState }

func (PlanRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanRule)(nil)).Elem()
}

func (o PlanRuleOutput) ToPlanRuleOutput() PlanRuleOutput {
	return o
}

func (o PlanRuleOutput) ToPlanRuleOutputWithContext(ctx context.Context) PlanRuleOutput {
	return o
}

// The amount of time AWS Backup attempts a backup before canceling the job and returning an error.
func (o PlanRuleOutput) CompletionWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlanRule) *int { return v.CompletionWindow }).(pulumi.IntPtrOutput)
}

// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
func (o PlanRuleOutput) Lifecycle() PlanRuleLifecyclePtrOutput {
	return o.ApplyT(func(v PlanRule) *PlanRuleLifecycle { return v.Lifecycle }).(PlanRuleLifecyclePtrOutput)
}

// Metadata that you can assign to help organize the resources that you create.
func (o PlanRuleOutput) RecoveryPointTags() pulumi.MapOutput {
	return o.ApplyT(func(v PlanRule) map[string]interface{} { return v.RecoveryPointTags }).(pulumi.MapOutput)
}

// An display name for a backup rule.
func (o PlanRuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v PlanRule) string { return v.RuleName }).(pulumi.StringOutput)
}

// A CRON expression specifying when AWS Backup initiates a backup job.
func (o PlanRuleOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanRule) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

// The amount of time in minutes before beginning a backup.
func (o PlanRuleOutput) StartWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlanRule) *int { return v.StartWindow }).(pulumi.IntPtrOutput)
}

// The name of a logical container where backups are stored.
func (o PlanRuleOutput) TargetVaultName() pulumi.StringOutput {
	return o.ApplyT(func(v PlanRule) string { return v.TargetVaultName }).(pulumi.StringOutput)
}

type PlanRuleArrayOutput struct{ *pulumi.OutputState }

func (PlanRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanRule)(nil)).Elem()
}

func (o PlanRuleArrayOutput) ToPlanRuleArrayOutput() PlanRuleArrayOutput {
	return o
}

func (o PlanRuleArrayOutput) ToPlanRuleArrayOutputWithContext(ctx context.Context) PlanRuleArrayOutput {
	return o
}

func (o PlanRuleArrayOutput) Index(i pulumi.IntInput) PlanRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlanRule {
		return vs[0].([]PlanRule)[vs[1].(int)]
	}).(PlanRuleOutput)
}

type PlanRuleLifecycle struct {
	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter *int `pulumi:"coldStorageAfter"`
	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `coldStorageAfter`.
	DeleteAfter *int `pulumi:"deleteAfter"`
}

// PlanRuleLifecycleInput is an input type that accepts PlanRuleLifecycleArgs and PlanRuleLifecycleOutput values.
// You can construct a concrete instance of `PlanRuleLifecycleInput` via:
//
// 		 PlanRuleLifecycleArgs{...}
//
type PlanRuleLifecycleInput interface {
	pulumi.Input

	ToPlanRuleLifecycleOutput() PlanRuleLifecycleOutput
	ToPlanRuleLifecycleOutputWithContext(context.Context) PlanRuleLifecycleOutput
}

type PlanRuleLifecycleArgs struct {
	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter pulumi.IntPtrInput `pulumi:"coldStorageAfter"`
	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `coldStorageAfter`.
	DeleteAfter pulumi.IntPtrInput `pulumi:"deleteAfter"`
}

func (PlanRuleLifecycleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanRuleLifecycle)(nil)).Elem()
}

func (i PlanRuleLifecycleArgs) ToPlanRuleLifecycleOutput() PlanRuleLifecycleOutput {
	return i.ToPlanRuleLifecycleOutputWithContext(context.Background())
}

func (i PlanRuleLifecycleArgs) ToPlanRuleLifecycleOutputWithContext(ctx context.Context) PlanRuleLifecycleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanRuleLifecycleOutput)
}

func (i PlanRuleLifecycleArgs) ToPlanRuleLifecyclePtrOutput() PlanRuleLifecyclePtrOutput {
	return i.ToPlanRuleLifecyclePtrOutputWithContext(context.Background())
}

func (i PlanRuleLifecycleArgs) ToPlanRuleLifecyclePtrOutputWithContext(ctx context.Context) PlanRuleLifecyclePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanRuleLifecycleOutput).ToPlanRuleLifecyclePtrOutputWithContext(ctx)
}

// PlanRuleLifecyclePtrInput is an input type that accepts PlanRuleLifecycleArgs, PlanRuleLifecyclePtr and PlanRuleLifecyclePtrOutput values.
// You can construct a concrete instance of `PlanRuleLifecyclePtrInput` via:
//
// 		 PlanRuleLifecycleArgs{...}
//
//  or:
//
// 		 nil
//
type PlanRuleLifecyclePtrInput interface {
	pulumi.Input

	ToPlanRuleLifecyclePtrOutput() PlanRuleLifecyclePtrOutput
	ToPlanRuleLifecyclePtrOutputWithContext(context.Context) PlanRuleLifecyclePtrOutput
}

type planRuleLifecyclePtrType PlanRuleLifecycleArgs

func PlanRuleLifecyclePtr(v *PlanRuleLifecycleArgs) PlanRuleLifecyclePtrInput {
	return (*planRuleLifecyclePtrType)(v)
}

func (*planRuleLifecyclePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanRuleLifecycle)(nil)).Elem()
}

func (i *planRuleLifecyclePtrType) ToPlanRuleLifecyclePtrOutput() PlanRuleLifecyclePtrOutput {
	return i.ToPlanRuleLifecyclePtrOutputWithContext(context.Background())
}

func (i *planRuleLifecyclePtrType) ToPlanRuleLifecyclePtrOutputWithContext(ctx context.Context) PlanRuleLifecyclePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanRuleLifecyclePtrOutput)
}

type PlanRuleLifecycleOutput struct{ *pulumi.OutputState }

func (PlanRuleLifecycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanRuleLifecycle)(nil)).Elem()
}

func (o PlanRuleLifecycleOutput) ToPlanRuleLifecycleOutput() PlanRuleLifecycleOutput {
	return o
}

func (o PlanRuleLifecycleOutput) ToPlanRuleLifecycleOutputWithContext(ctx context.Context) PlanRuleLifecycleOutput {
	return o
}

func (o PlanRuleLifecycleOutput) ToPlanRuleLifecyclePtrOutput() PlanRuleLifecyclePtrOutput {
	return o.ToPlanRuleLifecyclePtrOutputWithContext(context.Background())
}

func (o PlanRuleLifecycleOutput) ToPlanRuleLifecyclePtrOutputWithContext(ctx context.Context) PlanRuleLifecyclePtrOutput {
	return o.ApplyT(func(v PlanRuleLifecycle) *PlanRuleLifecycle {
		return &v
	}).(PlanRuleLifecyclePtrOutput)
}

// Specifies the number of days after creation that a recovery point is moved to cold storage.
func (o PlanRuleLifecycleOutput) ColdStorageAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlanRuleLifecycle) *int { return v.ColdStorageAfter }).(pulumi.IntPtrOutput)
}

// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `coldStorageAfter`.
func (o PlanRuleLifecycleOutput) DeleteAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlanRuleLifecycle) *int { return v.DeleteAfter }).(pulumi.IntPtrOutput)
}

type PlanRuleLifecyclePtrOutput struct{ *pulumi.OutputState }

func (PlanRuleLifecyclePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanRuleLifecycle)(nil)).Elem()
}

func (o PlanRuleLifecyclePtrOutput) ToPlanRuleLifecyclePtrOutput() PlanRuleLifecyclePtrOutput {
	return o
}

func (o PlanRuleLifecyclePtrOutput) ToPlanRuleLifecyclePtrOutputWithContext(ctx context.Context) PlanRuleLifecyclePtrOutput {
	return o
}

func (o PlanRuleLifecyclePtrOutput) Elem() PlanRuleLifecycleOutput {
	return o.ApplyT(func(v *PlanRuleLifecycle) PlanRuleLifecycle { return *v }).(PlanRuleLifecycleOutput)
}

// Specifies the number of days after creation that a recovery point is moved to cold storage.
func (o PlanRuleLifecyclePtrOutput) ColdStorageAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PlanRuleLifecycle) *int {
		if v == nil {
			return nil
		}
		return v.ColdStorageAfter
	}).(pulumi.IntPtrOutput)
}

// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `coldStorageAfter`.
func (o PlanRuleLifecyclePtrOutput) DeleteAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PlanRuleLifecycle) *int {
		if v == nil {
			return nil
		}
		return v.DeleteAfter
	}).(pulumi.IntPtrOutput)
}

type SelectionSelectionTag struct {
	// The key in a key-value pair.
	Key string `pulumi:"key"`
	// An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
	Type string `pulumi:"type"`
	// The value in a key-value pair.
	Value string `pulumi:"value"`
}

// SelectionSelectionTagInput is an input type that accepts SelectionSelectionTagArgs and SelectionSelectionTagOutput values.
// You can construct a concrete instance of `SelectionSelectionTagInput` via:
//
// 		 SelectionSelectionTagArgs{...}
//
type SelectionSelectionTagInput interface {
	pulumi.Input

	ToSelectionSelectionTagOutput() SelectionSelectionTagOutput
	ToSelectionSelectionTagOutputWithContext(context.Context) SelectionSelectionTagOutput
}

type SelectionSelectionTagArgs struct {
	// The key in a key-value pair.
	Key pulumi.StringInput `pulumi:"key"`
	// An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
	Type pulumi.StringInput `pulumi:"type"`
	// The value in a key-value pair.
	Value pulumi.StringInput `pulumi:"value"`
}

func (SelectionSelectionTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectionSelectionTag)(nil)).Elem()
}

func (i SelectionSelectionTagArgs) ToSelectionSelectionTagOutput() SelectionSelectionTagOutput {
	return i.ToSelectionSelectionTagOutputWithContext(context.Background())
}

func (i SelectionSelectionTagArgs) ToSelectionSelectionTagOutputWithContext(ctx context.Context) SelectionSelectionTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionSelectionTagOutput)
}

// SelectionSelectionTagArrayInput is an input type that accepts SelectionSelectionTagArray and SelectionSelectionTagArrayOutput values.
// You can construct a concrete instance of `SelectionSelectionTagArrayInput` via:
//
// 		 SelectionSelectionTagArray{ SelectionSelectionTagArgs{...} }
//
type SelectionSelectionTagArrayInput interface {
	pulumi.Input

	ToSelectionSelectionTagArrayOutput() SelectionSelectionTagArrayOutput
	ToSelectionSelectionTagArrayOutputWithContext(context.Context) SelectionSelectionTagArrayOutput
}

type SelectionSelectionTagArray []SelectionSelectionTagInput

func (SelectionSelectionTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectionSelectionTag)(nil)).Elem()
}

func (i SelectionSelectionTagArray) ToSelectionSelectionTagArrayOutput() SelectionSelectionTagArrayOutput {
	return i.ToSelectionSelectionTagArrayOutputWithContext(context.Background())
}

func (i SelectionSelectionTagArray) ToSelectionSelectionTagArrayOutputWithContext(ctx context.Context) SelectionSelectionTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionSelectionTagArrayOutput)
}

type SelectionSelectionTagOutput struct{ *pulumi.OutputState }

func (SelectionSelectionTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectionSelectionTag)(nil)).Elem()
}

func (o SelectionSelectionTagOutput) ToSelectionSelectionTagOutput() SelectionSelectionTagOutput {
	return o
}

func (o SelectionSelectionTagOutput) ToSelectionSelectionTagOutputWithContext(ctx context.Context) SelectionSelectionTagOutput {
	return o
}

// The key in a key-value pair.
func (o SelectionSelectionTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v SelectionSelectionTag) string { return v.Key }).(pulumi.StringOutput)
}

// An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
func (o SelectionSelectionTagOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SelectionSelectionTag) string { return v.Type }).(pulumi.StringOutput)
}

// The value in a key-value pair.
func (o SelectionSelectionTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SelectionSelectionTag) string { return v.Value }).(pulumi.StringOutput)
}

type SelectionSelectionTagArrayOutput struct{ *pulumi.OutputState }

func (SelectionSelectionTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectionSelectionTag)(nil)).Elem()
}

func (o SelectionSelectionTagArrayOutput) ToSelectionSelectionTagArrayOutput() SelectionSelectionTagArrayOutput {
	return o
}

func (o SelectionSelectionTagArrayOutput) ToSelectionSelectionTagArrayOutputWithContext(ctx context.Context) SelectionSelectionTagArrayOutput {
	return o
}

func (o SelectionSelectionTagArrayOutput) Index(i pulumi.IntInput) SelectionSelectionTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelectionSelectionTag {
		return vs[0].([]SelectionSelectionTag)[vs[1].(int)]
	}).(SelectionSelectionTagOutput)
}

func init() {
	pulumi.RegisterOutputType(PlanRuleOutput{})
	pulumi.RegisterOutputType(PlanRuleArrayOutput{})
	pulumi.RegisterOutputType(PlanRuleLifecycleOutput{})
	pulumi.RegisterOutputType(PlanRuleLifecyclePtrOutput{})
	pulumi.RegisterOutputType(SelectionSelectionTagOutput{})
	pulumi.RegisterOutputType(SelectionSelectionTagArrayOutput{})
}
