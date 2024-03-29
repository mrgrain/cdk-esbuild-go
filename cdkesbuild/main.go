// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.BuildOptions",
		reflect.TypeOf((*BuildOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.BundlerProps",
		reflect.TypeOf((*BundlerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.CompilerOptions",
		reflect.TypeOf((*CompilerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.EsbuildBundler",
		reflect.TypeOf((*EsbuildBundler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entryPoints", GoGetter: "EntryPoints"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "local", GoGetter: "Local"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_EsbuildBundler{}
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.EsbuildProvider",
		reflect.TypeOf((*EsbuildProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildSync", GoMethod: "BuildSync"},
			_jsii_.MemberMethod{JsiiMethod: "transformSync", GoMethod: "TransformSync"},
		},
		func() interface{} {
			j := jsiiProxy_EsbuildProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBuildProvider)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.EsbuildProviderProps",
		reflect.TypeOf((*EsbuildProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		reflect.TypeOf((*EsbuildSource)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EsbuildSource{}
		},
	)
	_jsii_.RegisterInterface(
		"@mrgrain/cdk-esbuild.IBuildProvider",
		reflect.TypeOf((*IBuildProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildSync", GoMethod: "BuildSync"},
		},
		func() interface{} {
			return &jsiiProxy_IBuildProvider{}
		},
	)
	_jsii_.RegisterInterface(
		"@mrgrain/cdk-esbuild.IEsbuildProvider",
		reflect.TypeOf((*IEsbuildProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildSync", GoMethod: "BuildSync"},
			_jsii_.MemberMethod{JsiiMethod: "transformSync", GoMethod: "TransformSync"},
		},
		func() interface{} {
			j := jsiiProxy_IEsbuildProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBuildProvider)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransformProvider)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@mrgrain/cdk-esbuild.ITransformProvider",
		reflect.TypeOf((*ITransformProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "transformSync", GoMethod: "TransformSync"},
		},
		func() interface{} {
			return &jsiiProxy_ITransformProvider{}
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.InlineJavaScriptCode",
		reflect.TypeOf((*InlineJavaScriptCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
		},
		func() interface{} {
			j := jsiiProxy_InlineJavaScriptCode{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaInlineCode)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.InlineTypeScriptCode",
		reflect.TypeOf((*InlineTypeScriptCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
		},
		func() interface{} {
			j := jsiiProxy_InlineTypeScriptCode{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaInlineCode)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.ProviderBuildOptions",
		reflect.TypeOf((*ProviderBuildOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.ProviderTransformOptions",
		reflect.TypeOf((*ProviderTransformOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.TransformOptions",
		reflect.TypeOf((*TransformOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.TransformerProps",
		reflect.TypeOf((*TransformerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.TsconfigRaw",
		reflect.TypeOf((*TsconfigRaw)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.TypeScriptAsset",
		reflect.TypeOf((*TypeScriptAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addResourceMetadata", GoMethod: "AddResourceMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "assetPath", GoGetter: "AssetPath"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "httpUrl", GoGetter: "HttpUrl"},
			_jsii_.MemberProperty{JsiiProperty: "isFile", GoGetter: "IsFile"},
			_jsii_.MemberProperty{JsiiProperty: "isZipArchive", GoGetter: "IsZipArchive"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketName", GoGetter: "S3BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "s3ObjectKey", GoGetter: "S3ObjectKey"},
			_jsii_.MemberProperty{JsiiProperty: "s3ObjectUrl", GoGetter: "S3ObjectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TypeScriptAsset{}
			_jsii_.InitJsiiProxy(&j.Type__awss3assetsAsset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.TypeScriptAssetProps",
		reflect.TypeOf((*TypeScriptAssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		reflect.TypeOf((*TypeScriptCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
		},
		func() interface{} {
			j := jsiiProxy_TypeScriptCode{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaCode)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.TypeScriptCodeProps",
		reflect.TypeOf((*TypeScriptCodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.TypeScriptSource",
		reflect.TypeOf((*TypeScriptSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_TypeScriptSource{}
			_jsii_.InitJsiiProxy(&j.Type__awss3deploymentISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.TypeScriptSourceProps",
		reflect.TypeOf((*TypeScriptSourceProps)(nil)).Elem(),
	)
}
