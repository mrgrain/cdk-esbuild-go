// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provides an implementation of the esbuild Transform API.
type ITransformProvider interface {
	// A method implementing the inline code transformation.
	//
	// During synth time, the method will receive the inline code and all computed `TransformOptions` from the bundler.
	//
	// MUST return the transformed code as a string to integrate correctly.
	// It MAY use these options to do so.
	// On failure, it SHOULD print any warnings & errors to stderr and throw a `TransformFailure` to inform the bundler.
	TransformSync(input *string, options *ProviderTransformOptions) *string
}

// The jsii proxy for ITransformProvider
type jsiiProxy_ITransformProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_ITransformProvider) TransformSync(input *string, options *ProviderTransformOptions) *string {
	if err := i.validateTransformSyncParameters(input, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"transformSync",
		[]interface{}{input, options},
		&returns,
	)

	return returns
}

