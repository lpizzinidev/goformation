// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_CustomParameterValues AWS CloudFormation Resource (AWS::QuickSight::Dashboard.CustomParameterValues)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customparametervalues.html
type Dashboard_CustomParameterValues struct {

	// DateTimeValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customparametervalues.html#cfn-quicksight-dashboard-customparametervalues-datetimevalues
	DateTimeValues []string `json:"DateTimeValues,omitempty"`

	// DecimalValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customparametervalues.html#cfn-quicksight-dashboard-customparametervalues-decimalvalues
	DecimalValues []float64 `json:"DecimalValues,omitempty"`

	// IntegerValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customparametervalues.html#cfn-quicksight-dashboard-customparametervalues-integervalues
	IntegerValues []float64 `json:"IntegerValues,omitempty"`

	// StringValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customparametervalues.html#cfn-quicksight-dashboard-customparametervalues-stringvalues
	StringValues []string `json:"StringValues,omitempty"`

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
func (r *Dashboard_CustomParameterValues) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.CustomParameterValues"
}
