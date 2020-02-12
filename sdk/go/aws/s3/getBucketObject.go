// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package s3

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The S3 object data source allows access to the metadata and
// _optionally_ (see below) content of an object stored inside S3 bucket.
//
// > **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/s3_bucket_object.html.markdown.
func LookupBucketObject(ctx *pulumi.Context, args *LookupBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupBucketObjectResult, error) {
	var rv LookupBucketObjectResult
	err := ctx.Invoke("aws:s3/getBucketObject:getBucketObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObject.
type LookupBucketObjectArgs struct {
	// The name of the bucket to read the object from
	Bucket string `pulumi:"bucket"`
	// The full path to the object inside the bucket
	Key string `pulumi:"key"`
	Range *string `pulumi:"range"`
	Tags map[string]interface{} `pulumi:"tags"`
	// Specific version ID of the object returned (defaults to latest version)
	VersionId *string `pulumi:"versionId"`
}


// A collection of values returned by getBucketObject.
type LookupBucketObjectResult struct {
	// Object data (see **limitations above** to understand cases in which this field is actually available)
	Body string `pulumi:"body"`
	Bucket string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain.
	CacheControl string `pulumi:"cacheControl"`
	// Specifies presentational information for the object.
	ContentDisposition string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding string `pulumi:"contentEncoding"`
	// The language the content is in.
	ContentLanguage string `pulumi:"contentLanguage"`
	// Size of the body in bytes.
	ContentLength int `pulumi:"contentLength"`
	// A standard MIME type describing the format of the object data.
	ContentType string `pulumi:"contentType"`
	// [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
	Etag string `pulumi:"etag"`
	// If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
	Expiration string `pulumi:"expiration"`
	// The date and time at which the object is no longer cacheable.
	Expires string `pulumi:"expires"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Key string `pulumi:"key"`
	// Last modified date of the object in RFC1123 format (e.g. `Mon, 02 Jan 2006 15:04:05 MST`)
	LastModified string `pulumi:"lastModified"`
	// A map of metadata stored with the object in S3
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Indicates whether this object has an active [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds). This field is only returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus string `pulumi:"objectLockLegalHoldStatus"`
	// The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) currently in place for this object.
	ObjectLockMode string `pulumi:"objectLockMode"`
	// The date and time when this object's object lock will expire.
	ObjectLockRetainUntilDate string `pulumi:"objectLockRetainUntilDate"`
	Range *string `pulumi:"range"`
	// If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.
	ServerSideEncryption string `pulumi:"serverSideEncryption"`
	// If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.
	SseKmsKeyId string `pulumi:"sseKmsKeyId"`
	// [Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for `Standard` storage class objects.
	StorageClass string `pulumi:"storageClass"`
	// A mapping of tags assigned to the object.
	Tags map[string]interface{} `pulumi:"tags"`
	// The latest version ID of the object returned.
	VersionId string `pulumi:"versionId"`
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.
	WebsiteRedirectLocation string `pulumi:"websiteRedirectLocation"`
}

