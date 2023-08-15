package cdkesbuild


// Configure the default EsbuildProvider.
type EsbuildProviderProps struct {
	// Path to the binary used by esbuild.
	//
	// This is the same as setting the ESBUILD_BINARY_PATH environment variable.
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
	//
	// Use the static methods on `EsbuildSource` to customize the default behavior.
	// Default: - `CDK_ESBUILD_MODULE_PATH` or package resolution (see description).
	//
	EsbuildModulePath *string `field:"optional" json:"esbuildModulePath" yaml:"esbuildModulePath"`
}

