//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineTypeScriptCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineTypeScriptCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *awslambda.ResourceBindOptions) error {
	return nil
}

func validateInlineTypeScriptCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineTypeScriptCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	return nil
}

func validateInlineTypeScriptCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineTypeScriptCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	return nil
}

func validateInlineTypeScriptCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	return nil
}

func validateInlineTypeScriptCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	return nil
}

func validateInlineTypeScriptCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineTypeScriptCodeParameters(code *string, props *TransformerProps) error {
	return nil
}

