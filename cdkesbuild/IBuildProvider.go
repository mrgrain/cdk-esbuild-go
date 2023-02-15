// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provides an implementation of the esbuild Build API.
type IBuildProvider interface {
	// A method implementing the code build.
	//
	// During synth time, the method will receive all computed `BuildOptions` from the bundler.
	//
	// It MUST implement any output options to integrate correctly and MAY use any other options.
	// On failure, it SHOULD print any warnings & errors to stderr and throw a `BuildFailure` to inform the bundler.
	BuildSync(options *ProviderBuildOptions)
}

// The jsii proxy for IBuildProvider
type jsiiProxy_IBuildProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IBuildProvider) BuildSync(options *ProviderBuildOptions) {
	if err := i.validateBuildSyncParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"buildSync",
		[]interface{}{options},
	)
}

