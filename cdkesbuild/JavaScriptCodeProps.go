package cdkesbuild


type JavaScriptCodeProps struct {
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
	// Copy additional files to the code [asset staging directory](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.AssetStaging.html#absolutestagedpath), before the build runs. Files copied like this will be overwritten by esbuild if they share the same name as any of the outputs.
	//
	// * When provided with a `string` or `array`, all files are copied to the root of asset staging directory.
	// * When given a `map`, the key indicates the destination relative to the asset staging directory and the value is a list of all sources to be copied.
	//
	// Therefore the following values for `copyDir` are all equivalent:
	// ```
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
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it can be used in construct IDs in order to enforce creation of a new resource when the content hash has changed.
	//
	// Defaults to a hash of all files in the resulting bundle.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
}

