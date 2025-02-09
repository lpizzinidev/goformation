// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_NumericRangeFilter AWS CloudFormation Resource (AWS::QuickSight::Dashboard.NumericRangeFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html
type Dashboard_NumericRangeFilter struct {

	// AggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-aggregationfunction
	AggregationFunction *Dashboard_AggregationFunction `json:"AggregationFunction,omitempty"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-column
	Column *Dashboard_ColumnIdentifier `json:"Column"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-filterid
	FilterId string `json:"FilterId"`

	// IncludeMaximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-includemaximum
	IncludeMaximum *bool `json:"IncludeMaximum,omitempty"`

	// IncludeMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-includeminimum
	IncludeMinimum *bool `json:"IncludeMinimum,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-nulloption
	NullOption string `json:"NullOption"`

	// RangeMaximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-rangemaximum
	RangeMaximum *Dashboard_NumericRangeFilterValue `json:"RangeMaximum,omitempty"`

	// RangeMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-rangeminimum
	RangeMinimum *Dashboard_NumericRangeFilterValue `json:"RangeMinimum,omitempty"`

	// SelectAllOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericrangefilter.html#cfn-quicksight-dashboard-numericrangefilter-selectalloptions
	SelectAllOptions *string `json:"SelectAllOptions,omitempty"`

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
func (r *Dashboard_NumericRangeFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.NumericRangeFilter"
}
