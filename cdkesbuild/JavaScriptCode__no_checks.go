//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaScriptCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (j *jsiiProxy_JavaScriptCode) validateBindToResourceParameters(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions) error {
	return nil
}

func (j *jsiiProxy_JavaScriptCode) validateGetAssetParameters(scope constructs.Construct) error {
	return nil
}

func validateJavaScriptCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateJavaScriptCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateJavaScriptCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateJavaScriptCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateJavaScriptCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateJavaScriptCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateJavaScriptCode_FromInlineParameters(code *string) error {
	return nil
}

func (j *jsiiProxy_JavaScriptCode) validateSetAssetParameters(val EsbuildAsset) error {
	return nil
}

func (j *jsiiProxy_JavaScriptCode) validateSetIsInlineParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_JavaScriptCode) validateSetPropsParameters(val *AssetProps) error {
	return nil
}

func validateNewJavaScriptCodeParameters(entryPoints interface{}, props *JavaScriptCodeProps) error {
	return nil
}

