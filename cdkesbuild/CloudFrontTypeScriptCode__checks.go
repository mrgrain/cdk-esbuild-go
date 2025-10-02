//go:build !no_runtime_type_checking

package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCloudFrontTypeScriptCode_FromFileParameters(entryPoint *string, props *CloudFrontFunctionCodeProps) error {
	if entryPoint == nil {
		return fmt.Errorf("parameter entryPoint is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateCloudFrontTypeScriptCode_FromInlineParameters(code *string, props *CloudFrontFunctionInlineCodeProps) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

