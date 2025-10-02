package cdkesbuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mrgrain/cdk-esbuild-go/cdkesbuild/v5/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

// TypeScript code for CloudFront Functions.
type CloudFrontTypeScriptCode interface {
}

// The jsii proxy struct for CloudFrontTypeScriptCode
type jsiiProxy_CloudFrontTypeScriptCode struct {
	_ byte // padding
}

func NewCloudFrontTypeScriptCode() CloudFrontTypeScriptCode {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontTypeScriptCode{}

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.CloudFrontTypeScriptCode",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewCloudFrontTypeScriptCode_Override(c CloudFrontTypeScriptCode) {
	_init_.Initialize()

	_jsii_.Create(
		"@mrgrain/cdk-esbuild.CloudFrontTypeScriptCode",
		nil, // no parameters
		c,
	)
}

// Create CloudFront Function code from a TypeScript file.
func CloudFrontTypeScriptCode_FromFile(entryPoint *string, props *CloudFrontFunctionCodeProps) awscloudfront.FunctionCode {
	_init_.Initialize()

	if err := validateCloudFrontTypeScriptCode_FromFileParameters(entryPoint, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.FunctionCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.CloudFrontTypeScriptCode",
		"fromFile",
		[]interface{}{entryPoint, props},
		&returns,
	)

	return returns
}

// Create CloudFront Function code from inline TypeScript code.
func CloudFrontTypeScriptCode_FromInline(code *string, props *CloudFrontFunctionInlineCodeProps) awscloudfront.FunctionCode {
	_init_.Initialize()

	if err := validateCloudFrontTypeScriptCode_FromInlineParameters(code, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.FunctionCode

	_jsii_.StaticInvoke(
		"@mrgrain/cdk-esbuild.CloudFrontTypeScriptCode",
		"fromInline",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

