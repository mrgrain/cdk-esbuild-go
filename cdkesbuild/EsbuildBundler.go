// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v3/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Low-level construct that can be used where `BundlingOptions` are required.
//
// This class directly interfaces with esbuild and provides almost no configuration safeguards.
// Experimental.
type EsbuildBundler interface {
	// A path or list or map of paths to the entry points of your code.
	//
	// Relative paths are by default resolved from the current working directory.
	// To change the working directory, see `buildOptions.absWorkingDir`.
	//
	// Absolute paths can be used if files are part of the working directory.
	//
	// Examples:
	//   - `'src/index.ts'`
	//   - `require.resolve('./lambda')`
	//   - `['src/index.ts', 'src/util.ts']`
	//   - `{one: 'src/two.ts', two: 'src/one.ts'}`
	// Experimental.
	EntryPoints() interface{}
	// Deprecated: This value is ignored since the bundler is always using a locally installed version of esbuild. However the property is required to comply with the `BundlingOptions` interface.
	Image() awscdk.DockerImage
	// Implementation of `ILocalBundling` interface, responsible for calling esbuild functions.
	// Experimental.
	Local() awscdk.ILocalBundling
	// Props to change the behaviour of the bundler.
	// Experimental.
	Props() *BundlerProps
}

// The jsii proxy struct for EsbuildBundler
type jsiiProxy_EsbuildBundler struct {
	_ byte // padding
}

func (j *jsiiProxy_EsbuildBundler) EntryPoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entryPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildBundler) Image() awscdk.DockerImage {
	var returns awscdk.DockerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildBundler) Local() awscdk.ILocalBundling {
	var returns awscdk.ILocalBundling
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EsbuildBundler) Props() *BundlerProps {
	var returns *BundlerProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewEsbuildBundler(entryPoints interface{}, props *BundlerProps) EsbuildBundler {
	_init_.Initialize()

	if err := validateNewEsbuildBundlerParameters(entryPoints, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EsbuildBundler{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildBundler",
		[]interface{}{entryPoints, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEsbuildBundler_Override(e EsbuildBundler, entryPoints interface{}, props *BundlerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildBundler",
		[]interface{}{entryPoints, props},
		e,
	)
}

