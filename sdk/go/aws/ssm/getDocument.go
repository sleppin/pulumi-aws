// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ssm

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Gets the contents of the specified Systems Manager document.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ssm_document.html.markdown.
func LookupDocument(ctx *pulumi.Context, args *LookupDocumentArgs, opts ...pulumi.InvokeOption) (*LookupDocumentResult, error) {
	var rv LookupDocumentResult
	err := ctx.Invoke("aws:ssm/getDocument:getDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocument.
type LookupDocumentArgs struct {
	// Returns the document in the specified format. The document format can be either JSON or YAML. JSON is the default format.
	DocumentFormat *string `pulumi:"documentFormat"`
	// The document version for which you want information.
	DocumentVersion *string `pulumi:"documentVersion"`
	// The name of the Systems Manager document.
	Name string `pulumi:"name"`
}


// A collection of values returned by getDocument.
type LookupDocumentResult struct {
	// The ARN of the document.
	Arn string `pulumi:"arn"`
	// The contents of the document.
	Content string `pulumi:"content"`
	DocumentFormat *string `pulumi:"documentFormat"`
	// The type of the document.
	DocumentType string `pulumi:"documentType"`
	DocumentVersion *string `pulumi:"documentVersion"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
}

