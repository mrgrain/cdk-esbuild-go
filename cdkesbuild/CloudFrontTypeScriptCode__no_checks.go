//go:build no_runtime_type_checking

package cdkesbuild

// Building without runtime type checking enabled, so all the below just return nil

func validateCloudFrontTypeScriptCode_FromFileParameters(entryPoint *string, props *CloudFrontFunctionCodeProps) error {
	return nil
}

func validateCloudFrontTypeScriptCode_FromInlineParameters(code *string, props *CloudFrontFunctionInlineCodeProps) error {
	return nil
}

