//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineJavaScriptCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineJavaScriptCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions) error {
	return nil
}

func validateInlineJavaScriptCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineJavaScriptCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateInlineJavaScriptCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineJavaScriptCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateInlineJavaScriptCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateInlineJavaScriptCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateInlineJavaScriptCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineJavaScriptCodeParameters(code *string, props interface{}) error {
	return nil
}

