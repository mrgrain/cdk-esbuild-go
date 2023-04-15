//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EsbuildCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (e *jsiiProxy_EsbuildCode) validateBindToResourceParameters(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions) error {
	return nil
}

func (e *jsiiProxy_EsbuildCode) validateGetAssetParameters(scope constructs.Construct) error {
	return nil
}

func validateEsbuildCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateEsbuildCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateEsbuildCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateEsbuildCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateEsbuildCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateEsbuildCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateEsbuildCode_FromInlineParameters(code *string) error {
	return nil
}

func (j *jsiiProxy_EsbuildCode) validateSetAssetParameters(val EsbuildAsset) error {
	return nil
}

func (j *jsiiProxy_EsbuildCode) validateSetIsInlineParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_EsbuildCode) validateSetPropsParameters(val *AssetProps) error {
	return nil
}

func validateNewEsbuildCodeParameters(entryPoints interface{}, props interface{}) error {
	return nil
}

