// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild


type JavaScriptCodeProps struct {
	// Escape hatch to provide the bundler with a custom build function.
	//
	// The function will receive the computed options from the bundler. It can use with these options as it wishes, however `outdir`/`outfile` must be respected to integrate with CDK.
	// Must throw a `BuildFailure` on failure to correctly inform the bundler.
	//
	// Returns: esbuild.BuildResult
	// Experimental.
	BuildFn interface{} `field:"optional" json:"buildFn" yaml:"buildFn"`
	// Build options passed on to esbuild. Please refer to the esbuild Build API docs for details.
	//
	// * `buildOptions.outdir: string`
	// The actual path for the output directory is defined by CDK. However setting this option allows to write files into a subdirectory. \
	// For example `{ outdir: 'js' }` will create an asset with a single directory called `js`, which contains all built files. This approach can be useful for static website deployments, where JavaScript code should be placed into a subdirectory. \
	// *Cannot be used together with `outfile`*.
	// * `buildOptions.outfile: string`
	// Relative path to a file inside the CDK asset output directory.
	// For example `{ outfile: 'js/index.js' }` will create an asset with a single directory called `js`, which contains a single file `index.js`. This can be useful to rename the entry point. \
	// *Cannot be used with multiple entryPoints or together with `outdir`.*
	// * `buildOptions.absWorkingDir: string`
	// Absolute path to the [esbuild working directory](https://esbuild.github.io/api/#working-directory) and defaults to the [current working directory](https://en.wikipedia.org/wiki/Working_directory). \
	// If paths cannot be found, a good starting point is to look at the concatenation of `absWorkingDir + entryPoint`. It must always be a valid absolute path pointing to the entry point. When needed, the probably easiest way to set absWorkingDir is to use a combination of `resolve` and `__dirname` (see "Library authors" section in the documentation).
	// See: https://esbuild.github.io/api/#build-api
	//
	BuildOptions *BuildOptions `field:"optional" json:"buildOptions" yaml:"buildOptions"`
	// Copy additional files to the code [asset staging directory](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.AssetStaging.html#absolutestagedpath), before the build runs. Files copied like this will be overwritten by esbuild if they share the same name as any of the outputs.
	//
	// * When provided with a `string` or `array`, all files are copied to the root of asset staging directory.
	// * When given a `map`, the key indicates the destination relative to the asset staging directory and the value is a list of all sources to be copied.
	//
	// Therefore the following values for `copyDir` are all equivalent:
	// ```ts
	// { copyDir: "path/to/source" }
	// { copyDir: ["path/to/source"] }
	// { copyDir: { ".": "path/to/source" } }
	// { copyDir: { ".": ["path/to/source"] } }
	// ```
	// The destination cannot be outside of the asset staging directory.
	// If you are receiving the error "Cannot copy files to outside of the asset staging directory."
	// you are likely using `..` or an absolute path as key on the `copyDir` map.
	// Instead use only relative paths and avoid `..`.
	CopyDir interface{} `field:"optional" json:"copyDir" yaml:"copyDir"`
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
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it can be used in construct IDs in order to enforce creation of a new resource when the content hash has changed.
	//
	// Defaults to a hash of all files in the resulting bundle.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
}

