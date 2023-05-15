//go:build !no_runtime_type_checking

package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_ITransformProvider) validateTransformSyncParameters(input *string, options *ProviderTransformOptions) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

