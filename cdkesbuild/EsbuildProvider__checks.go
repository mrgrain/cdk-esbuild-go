//go:build !no_runtime_type_checking

package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (e *jsiiProxy_EsbuildProvider) validateBuildSyncParameters(options *ProviderBuildOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_EsbuildProvider) validateTransformSyncParameters(input *string, options *ProviderTransformOptions) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEsbuildProvider_OverrideDefaultBuildProviderParameters(provider IBuildProvider) error {
	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}

	return nil
}

func validateEsbuildProvider_OverrideDefaultProviderParameters(provider IEsbuildProvider) error {
	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}

	return nil
}

func validateEsbuildProvider_OverrideDefaultTransformationProviderParameters(provider ITransformProvider) error {
	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}

	return nil
}

func validateNewEsbuildProviderParameters(props *EsbuildProviderProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

