package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mrgrain/cdk-esbuild-go/cdkesbuild/internal"
)

// An implementation of `lambda.InlineCode` using the esbuild Transform API. Inline function code is limited to 4 KiB after transformation.
type InlineJavaScriptCode interface {
	awslambda.InlineCode
	// Determines whether this Code is inline code or not.
	IsInline() *bool
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *awslambda.CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions)
}

// The jsii proxy struct for InlineJavaScriptCode
type jsiiProxy_InlineJavaScriptCode struct {
	internal.Type__awslambdaInlineCode
}

func (j *jsiiProxy_InlineJavaScriptCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewInlineJavaScriptCode(code *string, props *TransformerProps) InlineJavaScriptCode {
	_init_.Initialize()

	if err := validateNewInlineJavaScriptCodeParameters(code, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineJavaScriptCode{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		[]interface{}{code, props},
		&j,
	)

	return &j
}

func NewInlineJavaScriptCode_Override(i InlineJavaScriptCode, code *string, props *TransformerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		[]interface{}{code, props},
		i,
	)
}

// Loads the function code from a local disk path.
func InlineJavaScriptCode_FromAsset(path *string, options *awss3assets.AssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func InlineJavaScriptCode_FromAssetImage(directory *string, props *awslambda.AssetImageCodeProps) awslambda.AssetImageCode {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns awslambda.AssetImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func InlineJavaScriptCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) awslambda.S3Code {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns awslambda.S3Code

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func InlineJavaScriptCode_FromCfnParameters(props *awslambda.CfnParametersCodeProps) awslambda.CfnParametersCode {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns awslambda.CfnParametersCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func InlineJavaScriptCode_FromDockerBuild(path *string, options *awslambda.DockerBuildAssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func InlineJavaScriptCode_FromEcrImage(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) awslambda.EcrImageCode {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns awslambda.EcrImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func InlineJavaScriptCode_FromInline(code *string) awslambda.InlineCode {
	_init_.Initialize()

	if err := validateInlineJavaScriptCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns awslambda.InlineCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineJavaScriptCode) Bind(scope constructs.Construct) *awslambda.CodeConfig {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awslambda.CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineJavaScriptCode) BindToResource(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions) {
	if err := i.validateBindToResourceParameters(_resource, _options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

