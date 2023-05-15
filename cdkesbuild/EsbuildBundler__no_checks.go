//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEsbuildBundlerParameters(entryPoints interface{}, props *BundlerProps) error {
	return nil
}

