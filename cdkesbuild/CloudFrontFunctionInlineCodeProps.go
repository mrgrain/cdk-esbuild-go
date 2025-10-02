package cdkesbuild


// Properties for CloudFront Function inline code.
type CloudFrontFunctionInlineCodeProps struct {
	// Transform options passed on to esbuild.
	//
	// Please refer to the esbuild Transform API docs for details.
	// See: https://esbuild.github.io/api/#transform-api
	//
	TransformOptions *TransformOptions `field:"optional" json:"transformOptions" yaml:"transformOptions"`
	// The esbuild Transform API implementation to be used.
	//
	// Configure the default `EsbuildProvider` for more options or
	// provide a custom `ITransformProvider` as an escape hatch.
	// Default: new DefaultEsbuildProvider().
	//
	TransformProvider ITransformProvider `field:"optional" json:"transformProvider" yaml:"transformProvider"`
	// CloudFront Functions JavaScript runtime environment version to build for.
	// Default: CloudFrontFunctionRuntime.JS_1_0
	//
	Runtime CloudFrontFunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

