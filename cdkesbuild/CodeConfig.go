package cdkesbuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Result of binding `Code` into a `Function`.
type CodeConfig struct {
	// Docker image configuration (mutually exclusive with `s3Location` and `inlineCode`).
	// Default: - code is not an ECR container image.
	//
	Image *awslambda.CodeImageConfig `field:"optional" json:"image" yaml:"image"`
	// Inline code (mutually exclusive with `s3Location` and `image`).
	// Default: - code is not inline code.
	//
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode` and `image`).
	// Default: - code is not an s3 location.
	//
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

