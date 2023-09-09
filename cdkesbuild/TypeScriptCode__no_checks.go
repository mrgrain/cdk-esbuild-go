//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TypeScriptCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (t *jsiiProxy_TypeScriptCode) validateBindToResourceParameters(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions) error {
	return nil
}

func validateTypeScriptCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateTypeScriptCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateTypeScriptCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateTypeScriptCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateTypeScriptCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateTypeScriptCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateTypeScriptCode_FromInlineParameters(code *string) error {
	return nil
}

func (j *jsiiProxy_TypeScriptCode) validateSetIsInlineParameters(val *bool) error {
	return nil
}

func validateNewTypeScriptCodeParameters(entryPoints interface{}, props *TypeScriptCodeProps) error {
	return nil
}

