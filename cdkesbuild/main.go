package cdkesbuild

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.AssetProps",
		reflect.TypeOf((*AssetProps)(nil)).Elem(),
	)
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
		"@mrgrain/cdk-esbuild.EsbuildAsset",
		reflect.TypeOf((*EsbuildAsset)(nil)).Elem(),
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
			j := jsiiProxy_EsbuildAsset{}
			_jsii_.InitJsiiProxy(&j.Type__awss3assetsAsset)
			return &j
		},
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
		"@mrgrain/cdk-esbuild.EsbuildCode",
		reflect.TypeOf((*EsbuildCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "asset", GoGetter: "Asset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "entryPoints", GoGetter: "EntryPoints"},
			_jsii_.MemberMethod{JsiiMethod: "getAsset", GoMethod: "GetAsset"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_EsbuildCode{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaCode)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.EsbuildSource",
		reflect.TypeOf((*EsbuildSource)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EsbuildSource{}
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
		"@mrgrain/cdk-esbuild.InlineJsxCode",
		reflect.TypeOf((*InlineJsxCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
		},
		func() interface{} {
			j := jsiiProxy_InlineJsxCode{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaInlineCode)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.InlineTsxCode",
		reflect.TypeOf((*InlineTsxCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
		},
		func() interface{} {
			j := jsiiProxy_InlineTsxCode{}
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
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.JavaScriptAsset",
		reflect.TypeOf((*JavaScriptAsset)(nil)).Elem(),
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
			j := jsiiProxy_JavaScriptAsset{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EsbuildAsset)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.JavaScriptCode",
		reflect.TypeOf((*JavaScriptCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "asset", GoGetter: "Asset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "entryPoints", GoGetter: "EntryPoints"},
			_jsii_.MemberMethod{JsiiMethod: "getAsset", GoMethod: "GetAsset"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_JavaScriptCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EsbuildCode)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.JavaScriptCodeProps",
		reflect.TypeOf((*JavaScriptCodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.JavaScriptSource",
		reflect.TypeOf((*JavaScriptSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "asset", GoGetter: "Asset"},
			_jsii_.MemberProperty{JsiiProperty: "assetClass", GoGetter: "AssetClass"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_JavaScriptSource{}
			_jsii_.InitJsiiProxy(&j.Type__awss3deploymentISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@mrgrain/cdk-esbuild.JavaScriptSourceProps",
		reflect.TypeOf((*JavaScriptSourceProps)(nil)).Elem(),
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
		"@mrgrain/cdk-esbuild.TsconfigOptions",
		reflect.TypeOf((*TsconfigOptions)(nil)).Elem(),
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
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EsbuildAsset)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@mrgrain/cdk-esbuild.TypeScriptCode",
		reflect.TypeOf((*TypeScriptCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "asset", GoGetter: "Asset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bindToResource", GoMethod: "BindToResource"},
			_jsii_.MemberProperty{JsiiProperty: "entryPoints", GoGetter: "EntryPoints"},
			_jsii_.MemberMethod{JsiiMethod: "getAsset", GoMethod: "GetAsset"},
			_jsii_.MemberProperty{JsiiProperty: "isInline", GoGetter: "IsInline"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_TypeScriptCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EsbuildCode)
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
			_jsii_.MemberProperty{JsiiProperty: "asset", GoGetter: "Asset"},
			_jsii_.MemberProperty{JsiiProperty: "assetClass", GoGetter: "AssetClass"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
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
