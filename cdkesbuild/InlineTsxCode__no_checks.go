//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineTsxCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineTsxCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions) error {
	return nil
}

func validateInlineTsxCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineTsxCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateInlineTsxCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineTsxCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateInlineTsxCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateInlineTsxCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateInlineTsxCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineTsxCodeParameters(code *string, props interface{}) error {
	return nil
}

