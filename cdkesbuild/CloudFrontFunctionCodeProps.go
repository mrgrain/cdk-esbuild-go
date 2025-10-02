package cdkesbuild


// Properties for CloudFront Function TypeScript code.
type CloudFrontFunctionCodeProps struct {
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
	// The esbuild Build API implementation to be used.
	//
	// Configure the default `EsbuildProvider` for more options or
	// provide a custom `IBuildProvider` as an escape hatch.
	// Default: new EsbuildProvider().
	//
	BuildProvider IBuildProvider `field:"optional" json:"buildProvider" yaml:"buildProvider"`
	// CloudFront Functions JavaScript runtime environment version to build for.
	// Default: CloudFrontFunctionRuntime.JS_1_0
	//
	Runtime CloudFrontFunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

