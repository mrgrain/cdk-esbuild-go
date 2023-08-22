//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TypeScriptSource) validateBindParameters(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) error {
	return nil
}

func (j *jsiiProxy_TypeScriptSource) validateSetAssetParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TypeScriptSource) validateSetAssetClassParameters(val TypeScriptAsset) error {
	return nil
}

func (j *jsiiProxy_TypeScriptSource) validateSetPropsParameters(val *AssetProps) error {
	return nil
}

func validateNewTypeScriptSourceParameters(entryPoints interface{}, props *TypeScriptSourceProps) error {
	return nil
}

