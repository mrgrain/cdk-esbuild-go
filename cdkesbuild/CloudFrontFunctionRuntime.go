package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v5/jsii"
)

// CloudFront Functions JavaScript runtime environment version.
type CloudFrontFunctionRuntime interface {
	Value() *string
}

// The jsii proxy struct for CloudFrontFunctionRuntime
type jsiiProxy_CloudFrontFunctionRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_CloudFrontFunctionRuntime) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CloudFrontFunctionRuntime_JS_1_0() CloudFrontFunctionRuntime {
	_init_.Initialize()
	var returns CloudFrontFunctionRuntime
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.CloudFrontFunctionRuntime",
		"JS_1_0",
		&returns,
	)
	return returns
}

func CloudFrontFunctionRuntime_JS_2_0() CloudFrontFunctionRuntime {
	_init_.Initialize()
	var returns CloudFrontFunctionRuntime
	_jsii_.StaticGet(
		"@mrgrain/cdk-esbuild.CloudFrontFunctionRuntime",
		"JS_2_0",
		&returns,
	)
	return returns
}

