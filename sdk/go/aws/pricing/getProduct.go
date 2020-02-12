// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pricing

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the pricing information of all products in AWS.
// This data source is only available in a us-east-1 or ap-south-1 provider.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/pricing_product.html.markdown.
func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	var rv GetProductResult
	err := ctx.Invoke("aws:pricing/getProduct:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProduct.
type GetProductArgs struct {
	// A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
	Filters []GetProductFilter `pulumi:"filters"`
	// The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.
	ServiceCode string `pulumi:"serviceCode"`
}


// A collection of values returned by getProduct.
type GetProductResult struct {
	Filters []GetProductFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set to the product returned from the API.
	Result string `pulumi:"result"`
	ServiceCode string `pulumi:"serviceCode"`
}

