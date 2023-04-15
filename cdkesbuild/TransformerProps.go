// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild


// Experimental.
type TransformerProps struct {
	// Path to the binary used by esbuild.
	//
	// This is the same as setting the ESBUILD_BINARY_PATH environment variable.
	// Experimental.
	EsbuildBinaryPath *string `field:"optional" json:"esbuildBinaryPath" yaml:"esbuildBinaryPath"`
	// Absolute path to the esbuild module JS file.
	//
	// E.g. "/home/user/.npm/node_modules/esbuild/lib/main.js"
	//
	// If not set, the module path will be determined in the following order:
	//
	// - Use a path from the `CDK_ESBUILD_MODULE_PATH` environment variable
	// - In TypeScript, fallback to the default Node.js package resolution mechanism
	// - All other languages (Python, Go, .NET, Java) use an automatic "best effort" resolution mechanism. \
	//    The exact algorithm of this mechanism is considered an implementation detail and should not be relied on.
	//    If `esbuild` cannot be found, it might be installed dynamically to a temporary location.
	//    To opt-out of this behavior, set either `esbuildModulePath` or `CDK_ESBUILD_MODULE_PATH` env variable.
	// Experimental.
	EsbuildModulePath *string `field:"optional" json:"esbuildModulePath" yaml:"esbuildModulePath"`
	// Escape hatch to provide the bundler with a custom transform function.
	//
	// The function will receive the computed options from the bundler. It can use with these options as it wishes, however a TransformResult must be returned to integrate with CDK.
	// Must throw a `TransformFailure` on failure to correctly inform the bundler.
	//
	// Returns: esbuild.TransformResult
	// Experimental.
	TransformFn interface{} `field:"optional" json:"transformFn" yaml:"transformFn"`
	// Transform options passed on to esbuild.
	//
	// Please refer to the esbuild Transform API docs for details.
	// See: https://esbuild.github.io/api/#transform-api
	//
	// Experimental.
	TransformOptions *TransformOptions `field:"optional" json:"transformOptions" yaml:"transformOptions"`
}

