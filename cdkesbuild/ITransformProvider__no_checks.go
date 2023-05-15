//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITransformProvider) validateTransformSyncParameters(input *string, options *ProviderTransformOptions) error {
	return nil
}

