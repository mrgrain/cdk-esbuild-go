// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v4/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v4/internal"
)

type JavaScriptSource interface {
	awss3deployment.ISource
	Asset() interface{}
	SetAsset(val interface{})
	AssetClass() JavaScriptAsset
	SetAssetClass(val JavaScriptAsset)
	Props() *AssetProps
	SetProps(val *AssetProps)
	// Binds the source to a bucket deployment.
	Bind(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) *awss3deployment.SourceConfig
}

// The jsii proxy struct for JavaScriptSource
type jsiiProxy_JavaScriptSource struct {
	internal.Type__awss3deploymentISource
}

func (j *jsiiProxy_JavaScriptSource) Asset() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaScriptSource) AssetClass() JavaScriptAsset {
	var returns JavaScriptAsset
	_jsii_.Get(
		j,
		"assetClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaScriptSource) Props() *AssetProps {
	var returns *AssetProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewJavaScriptSource(entryPoints interface{}, props *JavaScriptSourceProps) JavaScriptSource {
	_init_.Initialize()

	if err := validateNewJavaScriptSourceParameters(entryPoints, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaScriptSource{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.JavaScriptSource",
		[]interface{}{entryPoints, props},
		&j,
	)

	return &j
}

func NewJavaScriptSource_Override(j JavaScriptSource, entryPoints interface{}, props *JavaScriptSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.JavaScriptSource",
		[]interface{}{entryPoints, props},
		j,
	)
}

func (j *jsiiProxy_JavaScriptSource)SetAsset(val interface{}) {
	if err := j.validateSetAssetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asset",
		val,
	)
}

func (j *jsiiProxy_JavaScriptSource)SetAssetClass(val JavaScriptAsset) {
	if err := j.validateSetAssetClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetClass",
		val,
	)
}

func (j *jsiiProxy_JavaScriptSource)SetProps(val *AssetProps) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_JavaScriptSource) Bind(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) *awss3deployment.SourceConfig {
	if err := j.validateBindParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *awss3deployment.SourceConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

