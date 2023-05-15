//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EsbuildAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func (e *jsiiProxy_EsbuildAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEsbuildAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEsbuildAssetParameters(scope constructs.Construct, id *string, props *AssetProps) error {
	return nil
}

