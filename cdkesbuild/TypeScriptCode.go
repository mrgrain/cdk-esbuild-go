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

// Represents the deployed TypeScript Code.
type TypeScriptCode interface {
	awslambda.Code
	// Determines whether this Code is inline code or not.
	// Deprecated: this value is ignored since inline is now determined based on the the inlineCode field of CodeConfig returned from bind().
	IsInline() *bool
	// Deprecated: this value is ignored since inline is now determined based on the the inlineCode field of CodeConfig returned from bind().
	SetIsInline(val *bool)
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *awslambda.CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions)
}

// The jsii proxy struct for TypeScriptCode
type jsiiProxy_TypeScriptCode struct {
	internal.Type__awslambdaCode
}

func (j *jsiiProxy_TypeScriptCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewTypeScriptCode(entryPoints interface{}, props *TypeScriptCodeProps) TypeScriptCode {
	_init_.Initialize()

	if err := validateNewTypeScriptCodeParameters(entryPoints, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TypeScriptCode{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		[]interface{}{entryPoints, props},
		&j,
	)

	return &j
}

func NewTypeScriptCode_Override(t TypeScriptCode, entryPoints interface{}, props *TypeScriptCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		[]interface{}{entryPoints, props},
		t,
	)
}

func (j *jsiiProxy_TypeScriptCode)SetIsInline(val *bool) {
	if err := j.validateSetIsInlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isInline",
		val,
	)
}

// Loads the function code from a local disk path.
func TypeScriptCode_FromAsset(path *string, options *awss3assets.AssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func TypeScriptCode_FromAssetImage(directory *string, props *awslambda.AssetImageCodeProps) awslambda.AssetImageCode {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns awslambda.AssetImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func TypeScriptCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) awslambda.S3Code {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns awslambda.S3Code

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func TypeScriptCode_FromCfnParameters(props *awslambda.CfnParametersCodeProps) awslambda.CfnParametersCode {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns awslambda.CfnParametersCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
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
func TypeScriptCode_FromDockerBuild(path *string, options *awslambda.DockerBuildAssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func TypeScriptCode_FromEcrImage(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) awslambda.EcrImageCode {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns awslambda.EcrImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func TypeScriptCode_FromInline(code *string) awslambda.InlineCode {
	_init_.Initialize()

	if err := validateTypeScriptCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns awslambda.InlineCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptCode) Bind(scope constructs.Construct) *awslambda.CodeConfig {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awslambda.CodeConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptCode) BindToResource(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions) {
	if err := t.validateBindToResourceParameters(resource, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"bindToResource",
		[]interface{}{resource, options},
	)
}

