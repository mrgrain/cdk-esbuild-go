//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBuildProvider) validateBuildSyncParameters(options *ProviderBuildOptions) error {
	return nil
}

