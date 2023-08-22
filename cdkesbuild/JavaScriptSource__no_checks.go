//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaScriptSource) validateBindParameters(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) error {
	return nil
}

func (j *jsiiProxy_JavaScriptSource) validateSetAssetParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JavaScriptSource) validateSetAssetClassParameters(val JavaScriptAsset) error {
	return nil
}

func (j *jsiiProxy_JavaScriptSource) validateSetPropsParameters(val *AssetProps) error {
	return nil
}

func validateNewJavaScriptSourceParameters(entryPoints interface{}, props *JavaScriptSourceProps) error {
	return nil
}

