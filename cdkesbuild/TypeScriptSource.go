package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v5/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v5/internal"
)

type TypeScriptSource interface {
	awss3deployment.ISource
	// Binds the source to a bucket deployment.
	Bind(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) *awss3deployment.SourceConfig
}

// The jsii proxy struct for TypeScriptSource
type jsiiProxy_TypeScriptSource struct {
	internal.Type__awss3deploymentISource
}

func NewTypeScriptSource(entryPoints interface{}, props *TypeScriptSourceProps) TypeScriptSource {
	_init_.Initialize()

	if err := validateNewTypeScriptSourceParameters(entryPoints, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TypeScriptSource{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.TypeScriptSource",
		[]interface{}{entryPoints, props},
		&j,
	)

	return &j
}

func NewTypeScriptSource_Override(t TypeScriptSource, entryPoints interface{}, props *TypeScriptSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.TypeScriptSource",
		[]interface{}{entryPoints, props},
		t,
	)
}

func (t *jsiiProxy_TypeScriptSource) Bind(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) *awss3deployment.SourceConfig {
	if err := t.validateBindParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *awss3deployment.SourceConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

