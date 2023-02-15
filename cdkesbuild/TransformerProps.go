// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild


type TransformerProps struct {
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
	TransformProvider ITransformProvider `field:"optional" json:"transformProvider" yaml:"transformProvider"`
}

