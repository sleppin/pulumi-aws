// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regional Regex Match Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRegexPatternSet, err := wafregional.NewRegexPatternSet(ctx, "exampleRegexPatternSet", &wafregional.RegexPatternSetArgs{
// 			RegexPatternStrings: pulumi.StringArray{
// 				pulumi.String("one"),
// 				pulumi.String("two"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = wafregional.NewRegexMatchSet(ctx, "exampleRegexMatchSet", &wafregional.RegexMatchSetArgs{
// 			RegexMatchTuples: wafregional.RegexMatchSetRegexMatchTupleArray{
// 				&wafregional.RegexMatchSetRegexMatchTupleArgs{
// 					FieldToMatch: &wafregional.RegexMatchSetRegexMatchTupleFieldToMatchArgs{
// 						Data: pulumi.String("User-Agent"),
// 						Type: pulumi.String("HEADER"),
// 					},
// 					RegexPatternSetId:  exampleRegexPatternSet.ID(),
// 					TextTransformation: pulumi.String("NONE"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// WAF Regional Regex Match Set can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:wafregional/regexMatchSet:RegexMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type RegexMatchSet struct {
	pulumi.CustomResourceState

	// The name or description of the Regex Match Set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayOutput `pulumi:"regexMatchTuples"`
}

// NewRegexMatchSet registers a new resource with the given unique name, arguments, and options.
func NewRegexMatchSet(ctx *pulumi.Context,
	name string, args *RegexMatchSetArgs, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	if args == nil {
		args = &RegexMatchSetArgs{}
	}
	var resource RegexMatchSet
	err := ctx.RegisterResource("aws:wafregional/regexMatchSet:RegexMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexMatchSet gets an existing RegexMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexMatchSetState, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	var resource RegexMatchSet
	err := ctx.ReadResource("aws:wafregional/regexMatchSet:RegexMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexMatchSet resources.
type regexMatchSetState struct {
	// The name or description of the Regex Match Set.
	Name *string `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

type RegexMatchSetState struct {
	// The name or description of the Regex Match Set.
	Name pulumi.StringPtrInput
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetState)(nil)).Elem()
}

type regexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name *string `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

// The set of arguments for constructing a RegexMatchSet resource.
type RegexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name pulumi.StringPtrInput
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetArgs)(nil)).Elem()
}

type RegexMatchSetInput interface {
	pulumi.Input

	ToRegexMatchSetOutput() RegexMatchSetOutput
	ToRegexMatchSetOutputWithContext(ctx context.Context) RegexMatchSetOutput
}

func (RegexMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexMatchSet)(nil)).Elem()
}

func (i RegexMatchSet) ToRegexMatchSetOutput() RegexMatchSetOutput {
	return i.ToRegexMatchSetOutputWithContext(context.Background())
}

func (i RegexMatchSet) ToRegexMatchSetOutputWithContext(ctx context.Context) RegexMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetOutput)
}

type RegexMatchSetOutput struct {
	*pulumi.OutputState
}

func (RegexMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexMatchSetOutput)(nil)).Elem()
}

func (o RegexMatchSetOutput) ToRegexMatchSetOutput() RegexMatchSetOutput {
	return o
}

func (o RegexMatchSetOutput) ToRegexMatchSetOutputWithContext(ctx context.Context) RegexMatchSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegexMatchSetOutput{})
}
