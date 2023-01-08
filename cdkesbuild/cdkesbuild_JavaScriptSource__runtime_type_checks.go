//go:build !no_runtime_type_checking

// CDK constructs for esbuild, an extremely fast JavaScript bundler
package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/constructs-go/constructs/v10"
)

func (j *jsiiProxy_JavaScriptSource) validateBindParameters(scope constructs.Construct, context *awss3deployment.DeploymentSourceContext) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(context, func() string { return "parameter context" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JavaScriptSource) validateSetAssetParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case JavaScriptAsset:
		// ok
	case TypeScriptAsset:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: JavaScriptAsset, TypeScriptAsset; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_JavaScriptSource) validateSetAssetClassParameters(val JavaScriptAsset) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JavaScriptSource) validateSetPropsParameters(val *AssetProps) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewJavaScriptSourceParameters(entryPoints interface{}, props *JavaScriptSourceProps) error {
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

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

