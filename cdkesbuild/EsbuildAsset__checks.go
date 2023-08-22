//go:build !no_runtime_type_checking

package cdkesbuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EsbuildAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	if resourceProperty == nil {
		return fmt.Errorf("parameter resourceProperty is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EsbuildAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateEsbuildAsset_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewEsbuildAssetParameters(scope constructs.Construct, id *string, props *AssetProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

