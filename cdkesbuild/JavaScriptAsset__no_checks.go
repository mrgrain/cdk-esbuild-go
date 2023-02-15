//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaScriptAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func (j *jsiiProxy_JavaScriptAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateJavaScriptAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewJavaScriptAssetParameters(scope constructs.Construct, id *string, props *AssetProps) error {
	return nil
}

