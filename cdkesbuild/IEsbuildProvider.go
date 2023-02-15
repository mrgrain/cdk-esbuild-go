// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild


// Provides an implementation of the esbuild Build & Transform API.
type IEsbuildProvider interface {
	IBuildProvider
	ITransformProvider
}

// The jsii proxy for IEsbuildProvider
type jsiiProxy_IEsbuildProvider struct {
	jsiiProxy_IBuildProvider
	jsiiProxy_ITransformProvider
}

