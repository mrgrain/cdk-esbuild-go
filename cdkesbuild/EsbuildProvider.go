package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v4/jsii"
)

// Default esbuild implementation calling esbuild's JavaScript API.
type EsbuildProvider interface {
	IBuildProvider
	ITransformProvider
	// A method implementing the code build.
	//
	// During synth time, the method will receive all computed `BuildOptions` from the bundler.
	//
	// It MUST implement any output options to integrate correctly and MAY use any other options.
	// On failure, it SHOULD print any warnings & errors to stderr and throw a `BuildFailure` to inform the bundler.
	BuildSync(options *ProviderBuildOptions)
	// A method implementing the inline code transformation.
	//
	// During synth time, the method will receive the inline code and all computed `TransformOptions` from the bundler.
	//
	// MUST return the transformed code as a string to integrate correctly.
	// It MAY use these options to do so.
	// On failure, it SHOULD print any warnings & errors to stderr and throw a `TransformFailure` to inform the bundler.
	TransformSync(input *string, options *ProviderTransformOptions) *string
}

// The jsii proxy struct for EsbuildProvider
type jsiiProxy_EsbuildProvider struct {
	jsiiProxy_IBuildProvider
	jsiiProxy_ITransformProvider
}

func NewEsbuildProvider(props *EsbuildProviderProps) EsbuildProvider {
	_init_.Initialize()

	if err := validateNewEsbuildProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EsbuildProvider{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEsbuildProvider_Override(e EsbuildProvider, props *EsbuildProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		[]interface{}{props},
		e,
	)
}

// Get the default implementation for the Build API.
func EsbuildProvider_DefaultBuildProvider() IBuildProvider {
	_init_.Initialize()

	var returns IBuildProvider

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		"defaultBuildProvider",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Get the default implementation for the Transformation API.
func EsbuildProvider_DefaultTransformationProvider() ITransformProvider {
	_init_.Initialize()

	var returns ITransformProvider

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		"defaultTransformationProvider",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Set the default implementation for the Build API.
func EsbuildProvider_OverrideDefaultBuildProvider(provider IBuildProvider) {
	_init_.Initialize()

	if err := validateEsbuildProvider_OverrideDefaultBuildProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		"overrideDefaultBuildProvider",
		[]interface{}{provider},
	)
}

// Set the default implementation for both Build and Transformation API.
func EsbuildProvider_OverrideDefaultProvider(provider IEsbuildProvider) {
	_init_.Initialize()

	if err := validateEsbuildProvider_OverrideDefaultProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		"overrideDefaultProvider",
		[]interface{}{provider},
	)
}

// Set the default implementation for the Transformation API.
func EsbuildProvider_OverrideDefaultTransformationProvider(provider ITransformProvider) {
	_init_.Initialize()

	if err := validateEsbuildProvider_OverrideDefaultTransformationProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		"overrideDefaultTransformationProvider",
		[]interface{}{provider},
	)
}

func (e *jsiiProxy_EsbuildProvider) BuildSync(options *ProviderBuildOptions) {
	if err := e.validateBuildSyncParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"buildSync",
		[]interface{}{options},
	)
}

func (e *jsiiProxy_EsbuildProvider) TransformSync(input *string, options *ProviderTransformOptions) *string {
	if err := e.validateTransformSyncParameters(input, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"transformSync",
		[]interface{}{input, options},
		&returns,
	)

	return returns
}

