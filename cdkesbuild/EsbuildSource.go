package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v4/jsii"
)

type EsbuildSource interface {
}

// The jsii proxy struct for EsbuildSource
type jsiiProxy_EsbuildSource struct {
	_ byte // padding
}

// Try to find the module in most common paths.
func EsbuildSource_Anywhere() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"anywhere",
		nil, // no parameters
		&returns,
	)

	return returns
}

// First try to find to module, then install it to a temporary location.
func EsbuildSource_Auto() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Try to find the module in common global installation paths.
func EsbuildSource_GlobalPaths() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"globalPaths",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Install the module to a temporary location.
func EsbuildSource_Install() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"install",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Require module by name, do not attempt to find it anywhere else.
func EsbuildSource_NodeJs() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"nodeJs",
		nil, // no parameters
		&returns,
	)

	return returns
}

// `EsbuildSource.nodeJs()` for NodeJs, `EsbuildSource.auto()` for all other languages.
func EsbuildSource_PlatformDefault() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"platformDefault",
		nil, // no parameters
		&returns,
	)

	return returns
}

