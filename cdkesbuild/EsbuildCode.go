package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v4/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v4/internal"
)

// Represents a generic esbuild code bundle.
//
// You should always use `TypeScriptCode` or `JavaScriptCode`.
// Experimental.
type EsbuildCode interface {
	awslambda.Code
	// Experimental.
	Asset() EsbuildAsset
	// Experimental.
	SetAsset(val EsbuildAsset)
	// A path or list or map of paths to the entry points of your code.
	//
	// Relative paths are by default resolved from the current working directory.
	// To change the working directory, see `buildOptions.absWorkingDir`.
	//
	// Absolute paths can be used if files are part of the working directory.
	//
	// Examples:
	//   - `'src/index.ts'`
	//   - `require.resolve('./lambda')`
	//   - `['src/index.ts', 'src/util.ts']`
	//   - `{one: 'src/two.ts', two: 'src/one.ts'}`
	// Experimental.
	EntryPoints() interface{}
	// Determines whether this Code is inline code or not.
	// Deprecated: this value is ignored since inline is now determined based on the the inlineCode field of CodeConfig returned from bind().
	IsInline() *bool
	// Deprecated: this value is ignored since inline is now determined based on the the inlineCode field of CodeConfig returned from bind().
	SetIsInline(val *bool)
	// Experimental.
	Props() *AssetProps
	// Experimental.
	SetProps(val *AssetProps)
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(scope constructs.Construct) *awslambda.CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	// Experimental.
	BindToResource(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions)
	// Experimental.
	GetAsset(scope constructs.Construct) EsbuildAsset
}

// The jsii proxy struct for EsbuildCode
type jsiiProxy_EsbuildCode struct {
	internal.Type__awslambdaCode
}

func (j *jsiiProxy_EsbuildCode) Asset() EsbuildAsset {
	var returns EsbuildAsset
	_jsii_.Get(
		j,
		"asset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildCode) EntryPoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entryPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildCode) Props() *AssetProps {
	var returns *AssetProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewEsbuildCode(entryPoints interface{}, props interface{}) EsbuildCode {
	_init_.Initialize()

	if err := validateNewEsbuildCodeParameters(entryPoints, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EsbuildCode{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		[]interface{}{entryPoints, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEsbuildCode_Override(e EsbuildCode, entryPoints interface{}, props interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		[]interface{}{entryPoints, props},
		e,
	)
}

func (j *jsiiProxy_EsbuildCode)SetAsset(val EsbuildAsset) {
	if err := j.validateSetAssetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asset",
		val,
	)
}

func (j *jsiiProxy_EsbuildCode)SetIsInline(val *bool) {
	if err := j.validateSetIsInlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isInline",
		val,
	)
}

func (j *jsiiProxy_EsbuildCode)SetProps(val *AssetProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

// Loads the function code from a local disk path.
// Experimental.
func EsbuildCode_FromAsset(path *string, options *awss3assets.AssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateEsbuildCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
// Experimental.
func EsbuildCode_FromAssetImage(directory *string, props *awslambda.AssetImageCodeProps) awslambda.AssetImageCode {
	_init_.Initialize()

	if err := validateEsbuildCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns awslambda.AssetImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
// Experimental.
func EsbuildCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) awslambda.S3Code {
	_init_.Initialize()

	if err := validateEsbuildCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns awslambda.S3Code

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
// Experimental.
func EsbuildCode_FromCfnParameters(props *awslambda.CfnParametersCodeProps) awslambda.CfnParametersCode {
	_init_.Initialize()

	if err := validateEsbuildCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns awslambda.CfnParametersCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
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
// Experimental.
func EsbuildCode_FromDockerBuild(path *string, options *awslambda.DockerBuildAssetOptions) awslambda.AssetCode {
	_init_.Initialize()

	if err := validateEsbuildCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns awslambda.AssetCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
// Experimental.
func EsbuildCode_FromEcrImage(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) awslambda.EcrImageCode {
	_init_.Initialize()

	if err := validateEsbuildCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns awslambda.EcrImageCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
// Experimental.
func EsbuildCode_FromInline(code *string) awslambda.InlineCode {
	_init_.Initialize()

	if err := validateEsbuildCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns awslambda.InlineCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EsbuildCode) Bind(scope constructs.Construct) *awslambda.CodeConfig {
	if err := e.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awslambda.CodeConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EsbuildCode) BindToResource(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions) {
	if err := e.validateBindToResourceParameters(resource, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"bindToResource",
		[]interface{}{resource, options},
	)
}

func (e *jsiiProxy_EsbuildCode) GetAsset(scope constructs.Construct) EsbuildAsset {
	if err := e.validateGetAssetParameters(scope); err != nil {
		panic(err)
	}
	var returns EsbuildAsset

	_jsii_.Invoke(
		e,
		"getAsset",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

