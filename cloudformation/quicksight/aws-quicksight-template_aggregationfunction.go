// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_AggregationFunction AWS CloudFormation Resource (AWS::QuickSight::Template.AggregationFunction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html
type Template_AggregationFunction struct {

	// CategoricalAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-categoricalaggregationfunction
	CategoricalAggregationFunction *string `json:"CategoricalAggregationFunction,omitempty"`

	// DateAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-dateaggregationfunction
	DateAggregationFunction *string `json:"DateAggregationFunction,omitempty"`

	// NumericalAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-numericalaggregationfunction
	NumericalAggregationFunction *Template_NumericalAggregationFunction `json:"NumericalAggregationFunction,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Template_AggregationFunction) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.AggregationFunction"
}
