// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_AggregationFunction AWS CloudFormation Resource (AWS::QuickSight::Analysis.AggregationFunction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-aggregationfunction.html
type Analysis_AggregationFunction struct {

	// CategoricalAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-aggregationfunction.html#cfn-quicksight-analysis-aggregationfunction-categoricalaggregationfunction
	CategoricalAggregationFunction *string `json:"CategoricalAggregationFunction,omitempty"`

	// DateAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-aggregationfunction.html#cfn-quicksight-analysis-aggregationfunction-dateaggregationfunction
	DateAggregationFunction *string `json:"DateAggregationFunction,omitempty"`

	// NumericalAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-aggregationfunction.html#cfn-quicksight-analysis-aggregationfunction-numericalaggregationfunction
	NumericalAggregationFunction *Analysis_NumericalAggregationFunction `json:"NumericalAggregationFunction,omitempty"`

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
func (r *Analysis_AggregationFunction) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.AggregationFunction"
}
