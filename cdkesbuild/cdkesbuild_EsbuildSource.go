// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v3/jsii"
)

type EsbuildSource interface {
}

// The jsii proxy struct for EsbuildSource
type jsiiProxy_EsbuildSource struct {
	_ byte // padding
}

func NewEsbuildSource() EsbuildSource {
	_init_.Initialize()

	j := jsiiProxy_EsbuildSource{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEsbuildSource_Override(e EsbuildSource) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		nil, // no parameters
		e,
	)
}

func EsbuildSource_Anywhere() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"anywhere",
		&returns,
	)
	return returns
}

func EsbuildSource_Auto() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"auto",
		&returns,
	)
	return returns
}

func EsbuildSource_Default() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"default",
		&returns,
	)
	return returns
}

func EsbuildSource_SetDefault(val *string) {
	_init_.Initialize()
	_jsii_.StaticSet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"default",
		val,
	)
}

func EsbuildSource_GlobalPaths() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"globalPaths",
		&returns,
	)
	return returns
}

func EsbuildSource_Install() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"install",
		&returns,
	)
	return returns
}

func EsbuildSource_NodeJs() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"nodeJs",
		&returns,
	)
	return returns
}

func EsbuildSource_PlatformDefault() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		"platformDefault",
		&returns,
	)
	return returns
}

