package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v5/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v5/internal"
)

// An implementation of `lambda.InlineCode` using the esbuild Transform API. Inline function code is limited to 4 KiB after transformation.
type InlineTypeScriptCode interface {
	awslambda.InlineCode
	IsInline() *bool
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *awslambda.CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions)
}

// The jsii proxy struct for InlineTypeScriptCode
type jsiiProxy_InlineTypeScriptCode struct {
	internal.Type__awslambdaInlineCode
}

func (j *jsiiProxy_InlineTypeScriptCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewInlineTypeScriptCode(code *string, props *TransformerProps) InlineTypeScriptCode {
	_init_.Initialize()

	if err := validateNewInlineTypeScriptCodeParameters(code, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineTypeScriptCode{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		[]interface{}{code, props},
		&j,
	)

	return &j
}

func NewInlineTypeScriptCode_Override(i InlineTypeScriptCode, code *string, props *TransformerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		[]interface{}{code, props},
		i,
	)
}

// Loads the function code from a local disk path.
func InlineTypeScriptCode_FromAsset(path *string, options *awss3assets.AssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func InlineTypeScriptCode_FromAssetImage(directory *string, props *awslambda.AssetImageCodeProps) awslambda.AssetImageCode {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns awslambda.AssetImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func InlineTypeScriptCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) awslambda.S3Code {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns awslambda.S3Code

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func InlineTypeScriptCode_FromCfnParameters(props *awslambda.CfnParametersCodeProps) awslambda.CfnParametersCode {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns awslambda.CfnParametersCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
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
func InlineTypeScriptCode_FromDockerBuild(path *string, options *awslambda.DockerBuildAssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func InlineTypeScriptCode_FromEcrImage(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) awslambda.EcrImageCode {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns awslambda.EcrImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func InlineTypeScriptCode_FromInline(code *string) awslambda.InlineCode {
	_init_.Initialize()

	if err := validateInlineTypeScriptCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns awslambda.InlineCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineTypeScriptCode) Bind(scope constructs.Construct) *awslambda.CodeConfig {
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

func (i *jsiiProxy_InlineTypeScriptCode) BindToResource(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions) {
	if err := i.validateBindToResourceParameters(_resource, _options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

