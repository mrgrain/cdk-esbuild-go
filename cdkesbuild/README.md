<img src="https://raw.githubusercontent.com/mrgrain/cdk-esbuild/main/images/wordmark-light.svg" alt="cdk-esbuild">

*CDK constructs for [esbuild](https://github.com/evanw/esbuild), an extremely fast JavaScript bundler*

[Getting started](#getting-started) |
[Documentation](#documentation) | [API Reference](#api-reference) | [Python & .NET](#python-and-net) | [FAQ](#faq)

[![View on Construct Hub](https://constructs.dev/badge?package=%40mrgrain%2Fcdk-esbuild)](https://constructs.dev/packages/@mrgrain/cdk-esbuild)

## Why?

*esbuild* is an extremely fast bundler and minifier for TypeScript and JavaScript.
This package makes *esbuild* available to deploy AWS Lambda Functions, static websites or publish assets for further usage.

AWS CDK [supports *esbuild* for AWS Lambda Functions](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-lambda-nodejs-readme.html), but the implementation cannot be used with other Constructs and doesn't expose all of *esbuild*'s API.

## Getting started

Install `cdk-esbuild` for Node.js using your favorite package manager:

```sh
# npm
npm install @mrgrain/cdk-esbuild@4
# Yarn
yarn add @mrgrain/cdk-esbuild@4
# pnpm
pnpm add @mrgrain/cdk-esbuild@4
```

For Python and .NET, use these commands:

```sh
# Python
pip install mrgrain.cdk-esbuild

# .NET
dotnet add package Mrgrain.CdkEsbuild
```

### AWS Lambda: Serverless function

> 💡 See [Lambda (TypeScript)](examples/typescript/lambda) and [Lambda (Python)](examples/typescript/lambda) for complete working examples of a how to deploy an AWS Lambda Function.

Use `TypeScriptCode` as the `code` of a [lambda function](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_lambda.Function.html#code):

```go
// Example automatically generated from non-compiling source. May contain errors.
bundledCode := cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"))

fn := lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime.nODEJS_18_X,
	handler: jsii.String("index.handler"),
	code: bundledCode,
})
```

### AWS S3: Static Website

> 💡 See [React App (TypeScript)](examples/typescript/website) for a complete working example of a how to deploy a React app to S3.

Use `TypeScriptSource` as one of the `sources` of a [static website deployment](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-s3-deployment-readme.html#roadmap):

```go
websiteBundle := cdkesbuild.NewTypeScriptSource(jsii.String("src/index.tsx"))

websiteBucket := s3.NewBucket(stack, jsii.String("WebsiteBucket"), &bucketProps{
	autoDeleteObjects: jsii.Boolean(true),
	publicReadAccess: jsii.Boolean(true),
	removalPolicy: cdk.removalPolicy_DESTROY,
	websiteIndexDocument: jsii.String("index.html"),
})

s3deploy.NewBucketDeployment(stack, jsii.String("DeployWebsite"), &bucketDeploymentProps{
	destinationBucket: websiteBucket,
	sources: []iSource{
		websiteBundle,
	},
})
```

### Amazon CloudWatch Synthetics: Canary monitoring

> 💡 See [Monitored Website (TypeScript)](examples/typescript/website) for a complete working example of a deployed and monitored website.

Synthetics runs a canary to produce traffic to an application for monitoring purposes. Use `TypeScriptCode` as the `code` of a Canary test:

> ℹ️ This feature depends on `@aws-cdk/aws-synthetics-alpha` which is in developer preview.
> Please install the package using the respective tool of your language.
> You may need to update your source code when upgrading to a newer version of this alpha package.

```go
bundledCode := cdkesbuild.NewTypeScriptCode(jsii.String("src/canary.ts"), &typeScriptCodeProps{
	buildOptions: &buildOptions{
		outdir: jsii.String("nodejs/node_modules"),
	},
})

canary := synthetics.NewCanary(stack, jsii.String("MyCanary"), &canaryProps{
	runtime: synthetics.runtime_SYNTHETICS_NODEJS_PUPPETEER_3_2(),
	test: synthetics.test.custom(&customTestOptions{
		code: bundledCode,
		handler: jsii.String("index.handler"),
	}),
})
```

## Documentation

The package exports various different constructs for use with AWS CDK features. The guiding design principal for this package is *"extend, don't replace"*. Expect constructs that you can provide as props, not complete replacements.

For use in **Lambda Functions** and **Synthetic Canaries**, the following classes implement `lambda.Code` ([reference](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_lambda.Code.html)) and `synthetics.Code` ([reference](https://docs.aws.amazon.com/cdk/api/v2/docs/@aws-cdk_aws-synthetics-alpha.Code.html)):

* `TypeScriptCode` & `JavaScriptCode`

Inline code is only supported by **Lambda**:

* `InlineTypeScriptCode` & `InlineJavaScriptCode`

For use with **S3 bucket deployments**, classes implementing `s3deploy.ISource` ([reference](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-s3-deployment-readme.html)):

* `TypeScriptSource` & `JavaScriptSource`

> *Code and Source constructs seamlessly plug-in to other high-level CDK constructs. They share the same set of parameters, props and build options.*

The following classes power the other features. You normally won't have to use them, but they are there if you need them:

* `TypeScriptAsset` & `JavaScriptAsset` implements `s3.Asset` ([reference](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_s3_assets.Asset.html)) \
  creates an asset uploaded to S3 which can be referenced by other constructs
* `EsbuildBundler` implements `core.BundlingOptions` ([reference](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.BundlingOptions.html)) \
  provides an interface for a *esbuild* bundler wherever needed
* `EsbuildProvider` implements `IBuildProvider` and `ITransformProvider` \
  provides the default *esbuild* API implementation and can be replaced with a custom implementation

### [API Reference](API.md)

Auto-generated reference for Constructs, Classes and Structs.
This information is also available as part of your IDE's code completion.

### Python and .NET

*Esbuild* requires a platform and architecture specific binary and currently has to be installed with a Node.js package manager like npm.

When using `cdk-esbuild` with Python or .NET, the package will automatically detect a local or global installation of the *esbuild* npm package.If none can be found, it will fall back to dynamically installing a copy into a temporary location.
The exact algorithm of this mechanism must be treated as an implementation detail and should not be relied on.
It can however be configured to a certain extent.
See the examples below for more details.

This "best effort" approach makes it easy to get started.
But is not always desirable, for example in environments with limited network access or when guaranteed repeatability of builds is a concern.
You have several options to opt-out of this behavior.

#### Provide a controlled installation of *esbuild*

The first step is to install a version of *esbuild* that is controlled by you.

I **strongly recommend** to install *esbuild* as a local package.
The easiest approach is to manage an additional Node.js project at the root of your AWS CDK project.
*esbuild* can then be added to the `package.json` file and it is your responsibility to ensure required setup steps are taken in all environments like development machines and CI/CD systems.

Instead of installing the *esbuild* package in a local project, it can also be **installed globally** with `npm install -g esbuild` or a similar command.
This approach might be preferred if a build container is prepared ahead of time, thus avoiding repeated package installation.

#### Change the automatic package detection

The second step is to make `cdk-esbuild` aware of your chosen install location.
This step is optional, but it's a good idea to have the location or at least the method explicitly defined.

To do this, you can set the `esbuildModulePath` prop on a `EsbuildProvider`.
Either pass a known, absolute or relative path as value, or use the `EsbuildSource` helper to set the detection method.
Please refer to the [`EsbuildSource`](API.md#esbuildsource) reference for a complete list of available methods.

```go
// Example automatically generated from non-compiling source. May contain errors.
// Use the standard Node.js algorithm to detect a locally installed package
// Use the standard Node.js algorithm to detect a locally installed package
cdkesbuild.NewEsbuildProvider(&esbuildProviderProps{
	esbuildModulePath: esbuildSource_NodeJs(),
})

// Provide an explicit path
// Provide an explicit path
cdkesbuild.NewEsbuildProvider(&esbuildProviderProps{
	esbuildModulePath: jsii.String("/home/user/node_modules/esbuild/lib/main.js"),
})
```

As a no-code approach, the `CDK_ESBUILD_MODULE_PATH` environment variable can be set in the same way.
An advantage of this is that the path can easily be changed for different systems.
Setting the env variable can be used with any installation approach, but is typically paired with a global installation of the *esbuild* package.
Note that `esbuildModulePath` takes precedence.

#### Override the default detection method

For an AWS CDK app with many instances of `TypeScriptCode` etc. it would be annoying to change the above for every single one of them.
Luckily, the default can be changed for all usages per app:

```go
// Example automatically generated from non-compiling source. May contain errors.
customModule := cdkesbuild.NewEsbuildProvider(&esbuildProviderProps{
	esbuildModulePath: esbuildSource_GlobalPaths(),
})
cdkesbuild.EsbuildProvider.overrideDefaultProvider(customModule)
```

### Customizing the Esbuild API

This package uses the *esbuild* JavaScript API.
In most situations the default API configuration will be suitable.
But sometimes it is required to configure *esbuild* differently or even provide a custom implementation.
Common reasons for this are:

* Using a pre-installed version of *esbuild* with Python or .NET
* If features not supported by the synchronous API are required, e.g. support for plugins
* If the default version constraints for *esbuild* are not suitable
* To use a version of esbuild that is installed by any other means than `npm`, including Docker

For these scenarios, this package offers customization options and an interface to provide a custom implementation:

```go
var myCustomBuildProvider iBuildProvider

var myCustomTransformProvider iTransformProvider


cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"), &typeScriptCodeProps{
	buildProvider: myCustomBuildProvider,
})

cdkesbuild.NewInlineTypeScriptCode(jsii.String("let x: number = 1"), &transformerProps{
	transformProvider: myCustomTransformProvider,
})
```

#### Esbuild binary path

It is possible to override the binary used by *esbuild* by setting a property on `EsbuildProvider`.
This is the same as setting the `ESBUILD_BINARY_PATH` environment variable.
Defining the `esbuildBinaryPath` prop takes precedence.

```go
buildProvider := cdkesbuild.NewEsbuildProvider(&esbuildProviderProps{
	esbuildBinaryPath: jsii.String("path/to/esbuild/binary"),
})

// This will use a different esbuild binary
// This will use a different esbuild binary
cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"), &typeScriptCodeProps{
	buildProvider: buildProvider,
})
```

#### Esbuild module path

The Node.js module discovery algorithm will normally be used to find the *esbuild* package.
It can be useful to use specify a different module path, for example if a globally installed package should be used instead of a local version.

```go
buildProvider := cdkesbuild.NewEsbuildProvider(&esbuildProviderProps{
	esbuildModulePath: jsii.String("/home/user/node_modules/esbuild/lib/main.js"),
})

// This will use a different esbuild module
// This will use a different esbuild module
cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"), &typeScriptCodeProps{
	buildProvider: buildProvider,
})
```

Alternatively supported by setting the `CDK_ESBUILD_MODULE_PATH` environment variable, which will apply to all uses.
Defining the `esbuildModulePath` prop takes precedence.

If you are a Python or .NET users, see the language specific guide for a more detailed expose on this feature.

#### Custom Build and Transform API implementations

> 💡 See [esbuild plugins w/ TypeScript](examples/typescript/esbuild-with-plugins) for a complete working example of a custom Build API implementation.

A custom implementation can be provided by implementing `IBuildProvider` or `ITransformProvider`:

```go
type customEsbuild struct {
}

func (this *customEsbuild) buildSync(options buildOptions) {}

func (this *customEsbuild) transformSync(code *string, options transformOptions) *string {
	// custom implementation goes here, return transformed code
	return jsii.String("transformed code")
}

// These will use the custom implementation
// These will use the custom implementation
cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"), &typeScriptCodeProps{
	buildProvider: NewCustomEsbuild(),
})
cdkesbuild.NewInlineTypeScriptCode(jsii.String("let x: number = 1"), &transformerProps{
	transformProvider: NewCustomEsbuild(),
})
```

Instead of *esbuild*, the custom methods will be invoked with all computed options.
Custom implementations can amend, change or discard any of the options.

The `IBuildProvider` integration with CDK relies on the `outdir`/`outfile` values and it's usually required to use them unchanged.

`ITransformProvider` must return the transformed code as a string.

Failures and warnings should be printed to stderr and thrown as the respective *esbuild* error.

#### Overriding the default implementation providers

The default implementation can also be set for all usages of `TypeScriptCode` etc. in an AWS CDK app.
You can change the default for both APIs at once or set a different implementation for each of them.

```go
myCustomEsbuildProvider := NewMyCustomEsbuildProvider()

cdkesbuild.EsbuildProvider.overrideDefaultProvider(myCustomEsbuildProvider)
cdkesbuild.EsbuildProvider.overrideDefaultBuildProvider(myCustomEsbuildProvider)
cdkesbuild.EsbuildProvider.overrideDefaultTransformationProvider(myCustomEsbuildProvider)

// This will use the custom provider without the need to define it as a prop
// This will use the custom provider without the need to define it as a prop
cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"))
```

### Roadmap & Contributions

[The project's roadmap is available on GitHub.](https://github.com/users/mrgrain/projects/1/views/7)

Please submit feature requests as issues to the repository.
All contributions are welcome, no matter if they are for already planned or completely new features.

## FAQ

### Should I use this package in production?

This package is stable and ready to be used in production, as many do.
However *esbuild* has not yet released a version 1.0.0 and its API is still in active development.
Please read the guide on [esbuild's production readiness](https://esbuild.github.io/faq/#production-readiness).

Note that *esbuild* minor version upgrades are also introduced in **minor versions** of this package.
Since *esbuild* is pre v1, these versions typically introduce breaking changes and this package will inherit them.
To avoid this behavior, add the desired *esbuild* version as a dependency to your package.

### How do I upgrade from `cdk-esbuild` v3?

Please refer to the [v4 release notes](https://github.com/mrgrain/cdk-esbuild/releases/tag/v4.0.0) for a list of backwards incompatible changes and respective upgrade instructions.

### [TS/JS] Why am I getting the error `Cannot find module 'esbuild'`?

This package depends on *esbuild* as an optional dependencies. If optional dependencies are not installed automatically on your system (e.g. when using npm v4-6), install *esbuild* explicitly:

```console
npm install esbuild
```

### [TS/JS] How can I use a different version of *esbuild*?

Use the [override](https://docs.npmjs.com/cli/v9/configuring-npm/package-json?v=true#overrides) instructions for your package manager to force a specific version, for example:

```json
{
  "overrides": {
    "esbuild": "0.14.7"
  }
}
```

Build and Transform interfaces are relatively stable across *esbuild* versions.
However if any incompatibilities occur, `buildOptions` / `transformOptions` can be cast to `any`:

```go
bundledCode := cdkesbuild.NewTypeScriptCode(jsii.String("src/handler.ts"), &typeScriptCodeProps{
	buildOptions: map[string]*string{
		"unsupportedOption": jsii.String("value"),
	},
})
```

### [Python/.NET] How can I use a different version of *esbuild*?

Install the desired version of *esbuild* locally or globally [as described in the documentation above](#python-and-dotnet).

Build and Transform interfaces are relatively stable across *esbuild* versions.
However if any incompatibilities occur, use the appropriate language features to cast any incompatible `buildOptions` / `transformOptions` to the correct types.

### Can I use this package in my published AWS CDK Construct?

It is possible to use `cdk-esbuild` in a published AWS CDK Construct library, but not recommended.
Always prefer to ship a compiled `.js` file or even bundle a zip archive in your package.
For AWS Lambda Functions, [projen provides an excellent solution](https://projen.io/awscdk.html#aws-lambda-functions).

If you do end up consuming `cdk-esbuild`, you will have to set `buildOptions.absWorkingDir`. The easiest way to do this, is to resolve the path based on the directory name of the calling file:

```js
// file: node_modules/construct-library/src/index.ts
const props = {
  buildOptions: {
    absWorkingDir: path.resolve(__dirname, ".."),
    // now: /user/local-app/node_modules/construct-library
  },
};
```

This will dynamically resolve to the correct path, wherever the package is installed. Please open an issue if you encounter any difficulties.

### Can I use this package with AWS CDK v1?

Yes, the `2.x.x` line of `cdk-esbuild` is compatible with AWS CDK v1. You can find the [documentation for it on the v2 branch](https://github.com/mrgrain/cdk-esbuild/tree/v2).

However, in line with AWS CDK v2, this version release now only receives security updates and critical bug fixes and support will fully end on June 1, 2023.
