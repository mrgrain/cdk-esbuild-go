//go:build !no_runtime_type_checking

package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EsbuildCode) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EsbuildCode) validateBindToResourceParameters(resource awscdk.CfnResource, options *awslambda.ResourceBindOptions) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_EsbuildCode) validateGetAssetParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateEsbuildCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEsbuildCode_FromAssetImageParameters(directory *string, props *awslambda.AssetImageCodeProps) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateEsbuildCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateEsbuildCode_FromCfnParametersParameters(props *awslambda.CfnParametersCodeProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateEsbuildCode_FromDockerBuildParameters(path *string, options *awslambda.DockerBuildAssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEsbuildCode_FromEcrImageParameters(repository awsecr.IRepository, props *awslambda.EcrImageCodeProps) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateEsbuildCode_FromInlineParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EsbuildCode) validateSetAssetParameters(val EsbuildAsset) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EsbuildCode) validateSetIsInlineParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EsbuildCode) validateSetPropsParameters(val *AssetProps) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewEsbuildCodeParameters(entryPoints interface{}, props interface{}) error {
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
	switch props.(type) {
	case *JavaScriptCodeProps:
		props := props.(*JavaScriptCodeProps)
		if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
			return err
		}
	case JavaScriptCodeProps:
		props_ := props.(JavaScriptCodeProps)
		props := &props_
		if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
			return err
		}
	case *TypeScriptCodeProps:
		props := props.(*TypeScriptCodeProps)
		if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
			return err
		}
	case TypeScriptCodeProps:
		props_ := props.(TypeScriptCodeProps)
		props := &props_
		if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(props) {
			return fmt.Errorf("parameter props must be one of the allowed types: *JavaScriptCodeProps, *TypeScriptCodeProps; received %#v (a %T)", props, props)
		}
	}

	return nil
}

