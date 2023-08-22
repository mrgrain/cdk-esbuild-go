package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v3/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v3/internal"
)

// Represents a generic esbuild asset.
//
// You should always use `TypeScriptAsset` or `JavaScriptAsset`.
// Experimental.
type EsbuildAsset interface {
	awss3assets.Asset
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	// Experimental.
	AssetHash() *string
	// The path to the asset, relative to the current Cloud Assembly.
	//
	// If asset staging is disabled, this will just be the original path.
	// If asset staging is enabled it will be the staged path.
	// Experimental.
	AssetPath() *string
	// The S3 bucket in which this asset resides.
	// Experimental.
	Bucket() awss3.IBucket
	// Attribute which represents the S3 HTTP URL of this asset.
	//
	// Example:
	//   https://s3.us-west-1.amazonaws.com/bucket/key
	//
	// Experimental.
	HttpUrl() *string
	// Indicates if this asset is a single file.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	// Experimental.
	IsFile() *bool
	// Indicates if this asset is a zip archive.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	// Experimental.
	IsZipArchive() *bool
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Attribute that represents the name of the bucket this asset exists in.
	// Experimental.
	S3BucketName() *string
	// Attribute which represents the S3 object key of this asset.
	// Experimental.
	S3ObjectKey() *string
	// Attribute which represents the S3 URL of this asset.
	//
	// Example:
	//   s3://bucket/key
	//
	// Experimental.
	S3ObjectUrl() *string
	// Adds CloudFormation template metadata to the specified resource with information that indicates which resource property is mapped to this local asset.
	//
	// This can be used by tools such as SAM CLI to provide local
	// experience such as local invocation and debugging of Lambda functions.
	//
	// Asset metadata will only be included if the stack is synthesized with the
	// "aws:cdk:enable-asset-metadata" context key defined, which is the default
	// behavior when synthesizing via the CDK Toolkit.
	// See: https://github.com/aws/aws-cdk/issues/1432
	//
	// Experimental.
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	// Grants read permissions to the principal on the assets bucket.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EsbuildAsset
type jsiiProxy_EsbuildAsset struct {
	internal.Type__awss3assetsAsset
}

func (j *jsiiProxy_EsbuildAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) AssetPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) IsFile() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) IsZipArchive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isZipArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) S3ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildAsset) S3ObjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectUrl",
		&returns,
	)
	return returns
}


// Experimental.
func NewEsbuildAsset(scope constructs.Construct, id *string, props *AssetProps) EsbuildAsset {
	_init_.Initialize()

	if err := validateNewEsbuildAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EsbuildAsset{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEsbuildAsset_Override(e EsbuildAsset, scope constructs.Construct, id *string, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildAsset",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func EsbuildAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEsbuildAsset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EsbuildAsset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	if err := e.validateAddResourceMetadataParameters(resource, resourceProperty); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

func (e *jsiiProxy_EsbuildAsset) GrantRead(grantee awsiam.IGrantable) {
	if err := e.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"grantRead",
		[]interface{}{grantee},
	)
}

func (e *jsiiProxy_EsbuildAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

