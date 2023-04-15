//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineJsxCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineJsxCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions) error {
	return nil
}

func validateInlineJsxCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineJsxCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateInlineJsxCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineJsxCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateInlineJsxCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateInlineJsxCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateInlineJsxCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineJsxCodeParameters(code *string, props interface{}) error {
	return nil
}

