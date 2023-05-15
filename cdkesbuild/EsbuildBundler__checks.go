//go:build !no_runtime_type_checking

package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewEsbuildBundlerParameters(entryPoints interface{}, props *BundlerProps) error {
	if entryPoints == nil {
		return fmt.Errorf("parameter entryPoints is required, but nil was provided")
	}
	switch entryPoints.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	case *map[string]*string:
		// ok
	case map[string]*string:
		// ok
	default:
		return fmt.Errorf("parameter entryPoints must be one of the allowed types: *string, *[]*string, *map[string]*string; received %#v (a %T)", entryPoints, entryPoints)
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

