//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TypeScriptAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func (t *jsiiProxy_TypeScriptAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateTypeScriptAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTypeScriptAssetParameters(scope constructs.Construct, id *string, props *AssetProps) error {
	return nil
}

