//go:build no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EsbuildProvider) validateBuildSyncParameters(options *ProviderBuildOptions) error {
	return nil
}

func (e *jsiiProxy_EsbuildProvider) validateTransformSyncParameters(input *string, options *ProviderTransformOptions) error {
	return nil
}

func validateEsbuildProvider_OverrideDefaultBuildProviderParameters(provider IBuildProvider) error {
	return nil
}

func validateEsbuildProvider_OverrideDefaultProviderParameters(provider IEsbuildProvider) error {
	return nil
}

func validateEsbuildProvider_OverrideDefaultTransformationProviderParameters(provider ITransformProvider) error {
	return nil
}

func validateNewEsbuildProviderParameters(props *EsbuildProviderProps) error {
	return nil
}

